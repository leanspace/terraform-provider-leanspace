package helper

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// A condition is an interface that can be used to evaluate a map[string]any,
// and decide if it is valid or not.
type Condition interface {
	// Will evaluate against the given object and return if it is valid or not.
	eval(v map[string]any) bool
	// A human readable formatting of what this condition expects.
	printExpected() string
	// A human readable formatting of what this condition received.
	printActual(v map[string]any) string
}

// A slice of Conditions. It can be evaluated against a map[string]any, and all
// errors will be aggregated together.
type Validators []Condition

// Will check these conditions, and ensure they all evaluate to true for the given object.
// If all tests pass, returns nil.
// Otherwise returns a human readable error with all failed conditions aggregated in a
// human readable format.
func (validators Validators) Check(obj map[string]any) error {
	errorMsg := ""
	for _, validator := range validators {
		if !validator.eval(obj) {
			errorMsg += fmt.Sprintf(
				"  - %v\n"+
					"    got %v\n",
				validator.printExpected(),
				validator.printActual(obj),
			)
		}
	}
	if errorMsg != "" {
		return fmt.Errorf("validation error(s):\n%v", errorMsg)
	}
	return nil
}

func GetValue(key string, v map[string]any) any {
	keys := strings.Split(key, ".")
	valueToCheck := v[keys[0]]
	for i := 1; i < len(keys); i++ {
		valueToCheck = valueToCheck.([]interface{})[0].(map[string]any)[keys[i]]
	}

	return valueToCheck
}

// Conditions

type ifCondition struct {
	if_  Condition
	then Condition
}

func (c ifCondition) eval(v map[string]any) bool {
	return !c.if_.eval(v) || c.then.eval(v)
}
func (c ifCondition) printExpected() string {
	return fmt.Sprintf("if %v then %v", c.if_.printExpected(), c.then.printExpected())
}
func (c ifCondition) printActual(v map[string]any) string {
	return fmt.Sprintf("%v / %v", c.if_.printActual(v), c.then.printActual(v))
}

// A standard if ... then ... condition, that only evaluates to false if
// the "if" condition is true while "then" isn't.
func If(if_ Condition, then Condition) Condition {
	return ifCondition{if_, then}
}

type equivCondition struct {
	equivA Condition
	equivB Condition
}

func (c equivCondition) eval(v map[string]any) bool {
	return c.equivA.eval(v) == c.equivB.eval(v)
}
func (c equivCondition) printExpected() string {
	return fmt.Sprintf("if and only if %v then %v", c.equivA.printExpected(), c.equivB.printExpected())
}
func (c equivCondition) printActual(v map[string]any) string {
	return fmt.Sprintf("%v / %v", c.equivA.printActual(v), c.equivB.printActual(v))
}

// A "if and only if" (ie. equivalence) condition, that evaluates to true
// If both conditionals evaluate to the same.
func Equivalence(equivA Condition, equivB Condition) Condition {
	return equivCondition{equivA, equivB}
}

type notCondition struct {
	cond Condition
}

func (c notCondition) eval(v map[string]any) bool {
	return !c.cond.eval(v)
}
func (c notCondition) printExpected() string {
	return fmt.Sprintf("not %v", c.cond.printExpected())
}
func (c notCondition) printActual(v map[string]any) string {
	return c.cond.printActual(v)
}

// Will inverse the given condition
func Not(cond Condition) Condition {
	return notCondition{cond}
}

type andCondition struct {
	conds []Condition
}

func (c andCondition) eval(v map[string]any) bool {
	for _, cond := range c.conds {
		if !cond.eval(v) {
			return false
		}
	}
	return true
}
func (c andCondition) printExpected() string {
	base := "("
	for i, cond := range c.conds {
		if i > 0 {
			base += " & "
		}
		base += cond.printExpected()
	}
	base += ")"
	return base
}
func (c andCondition) printActual(v map[string]any) string {
	base := ""
	isEmpty := true
	for _, cond := range c.conds {
		d := cond.printActual(v)
		if !strings.Contains(base, d) {
			if !isEmpty {
				base += ", "
			}
			base += d
			isEmpty = false
		}
	}
	return base
}

// Will only evaluate to true if all conditions evaluate to true.
func And(conds ...Condition) Condition {
	return andCondition{conds}
}

type orCondition struct {
	conds []Condition
}

func (c orCondition) eval(v map[string]any) bool {
	for _, cond := range c.conds {
		if cond.eval(v) {
			return true
		}
	}
	return false
}
func (c orCondition) printExpected() string {
	base := "("
	for i, cond := range c.conds {
		if i > 0 {
			base += " | "
		}
		base += cond.printExpected()
	}
	base += ")"
	return base
}
func (c orCondition) printActual(v map[string]any) string {
	base := ""
	isEmpty := true
	for _, cond := range c.conds {
		d := cond.printActual(v)
		if !strings.Contains(base, d) {
			if !isEmpty {
				base += ", "
			}
			base += d
			isEmpty = false
		}
	}
	return base
}

// Will evaluate to true if any of the conditions evaluates to true.
func Or(conds ...Condition) Condition {
	return orCondition{conds}
}

type isSetCondition struct {
	key string
}

func (c isSetCondition) eval(v map[string]any) bool {
	return v[c.key] != nil && v[c.key] != "" && v[c.key] != 0 && v[c.key] != '\x00' && v[c.key] != 0.0
}
func (c isSetCondition) printExpected() string {
	return fmt.Sprintf("%q is set", c.key)
}
func (c isSetCondition) printActual(v map[string]any) string {
	return fmt.Sprintf("%q = %v", c.key, v[c.key])
}

// Will evaluate to true if the given key is set (ie. non-nil, not empty string)
func IsSet(key string) Condition {
	return isSetCondition{key}
}

type isEqualsCondition struct {
	key   string
	value any
}

func (c isEqualsCondition) eval(v map[string]any) bool {
	return GetValue(c.key, v) == c.value
}
func (c isEqualsCondition) printExpected() string {
	return fmt.Sprintf("%q = %v", c.key, c.value)
}
func (c isEqualsCondition) printActual(v map[string]any) string {
	return fmt.Sprintf("%q = %v", c.key, GetValue(c.key, v))
}

// Will evaluate to true if the value at the given key equals the given value
func Equals(key string, value any) Condition {
	return isEqualsCondition{key, value}
}

type hasLengthCondition struct {
	key    string
	length int
}

func (c hasLengthCondition) eval(v map[string]any) bool {
	if list, isList := v[c.key].([]any); isList {
		return len(list) == c.length
	}
	if map_, isMap := v[c.key].(map[string]any); isMap {
		return len(map_) == c.length
	}
	if set, isSet := v[c.key].(*schema.Set); isSet {
		return set.Len() == c.length
	}
	panic(fmt.Sprintf("Tried checking length of %#v (only accepts lists, maps and sets)", v[c.key]))
}
func (c hasLengthCondition) printExpected() string {
	return fmt.Sprintf("length(%q) = %v", c.key, c.length)
}
func (c hasLengthCondition) printActual(v map[string]any) string {
	return fmt.Sprintf("%q = %v", c.key, v[c.key])
}

// Will evaluate to true if the list/set/map at the given key is empty.
// Panics if something other than a list/set/map is found.
func IsEmpty(key string) Condition {
	return hasLengthCondition{key, 0}
}

// Will evaluate to true if the list/set/map at the given key has the given
// length. Panics if something other than a list/set/map is found.
func HasLength(key string, length int) Condition {
	return hasLengthCondition{key, length}
}

type regexCondition struct {
	key   string
	regex regexp.Regexp
}

func (c regexCondition) eval(v map[string]any) bool {
	if str, isString := v[c.key].(string); isString {
		return c.regex.MatchString(str)
	}
	panic(fmt.Sprintf("Tried regexing non-string: %v", v[c.key]))
}
func (c regexCondition) printExpected() string {
	return fmt.Sprintf("%q matches %v", c.key, c.regex)
}
func (c regexCondition) printActual(v map[string]any) string {
	return fmt.Sprintf("%q = %v", c.key, v[c.key])
}

func Matches(key string, regex regexp.Regexp) Condition {
	return regexCondition{key, regex}
}

type Number interface {
	int | int32 | int64 | float32 | float64
}

type compareCondition[T Number] struct {
	key   string
	value T
	op    string
}

func (c compareCondition[T]) eval(v map[string]any) bool {
	switch c.op {
	case "<":
		return GetValue(c.key, v).(T) < c.value
	case ">":
		return GetValue(c.key, v).(T) > c.value
	case "<=":
		return GetValue(c.key, v).(T) <= c.value
	case ">=":
		return GetValue(c.key, v).(T) >= c.value
	default:
		panic(fmt.Sprintf("unrecognized operator %q", c.op))
	}
}

func (c compareCondition[T]) printExpected() string {
	return fmt.Sprintf("%q %v %v", c.key, c.op, c.value)
}
func (c compareCondition[T]) printActual(v map[string]any) string {
	return fmt.Sprintf("%q = %v", c.key, GetValue(c.key, v))
}

// Evaluates to true if the value at the given key is less than the given value.
func LessThan[T Number](key string, value T) Condition {
	return compareCondition[T]{key, value, "<"}
}

// Evaluates to true if the value at the given key is greater than the given value.
func GreaterThan[T Number](key string, value T) Condition {
	return compareCondition[T]{key, value, ">"}
}

// Evaluates to true if the value at the given key is less than or equal to the given value.
func LessThanEq[T Number](key string, value T) Condition {
	return compareCondition[T]{key, value, "<="}
}

// Evaluates to true if the value at the given key is greater than or equal to the given value.
func GreaterThanEq[T Number](key string, value T) Condition {
	return compareCondition[T]{key, value, ">="}
}

func IsValidTimeDateOrTimestamp(i interface{}, k string) (warnings []string, errorsOnField []error) {
	v, ok := i.(string)
	if !ok {
		errorsOnField = append(errorsOnField, fmt.Errorf("expected type of %q to be string", k))
		return warnings, errorsOnField
	}
	const dateLayoutReference = "2006-01-02"
	const timeLayoutReference = "15:04:05"
	const timestampLayoutReference = time.RFC3339

	_, errTimestamp := time.Parse(timestampLayoutReference, v)
	_, errDate := time.Parse(dateLayoutReference, v)
	_, errTime := time.Parse(timeLayoutReference, v)

	if errTimestamp != nil && errDate != nil && errTime != nil {
		errorsOnField = append(errorsOnField, fmt.Errorf("expected %q to be a valid date, time or timestamp, got %q: \n %+v \n %+v \n %+v", k, i, errTimestamp, errDate, errTime))
	}
	return warnings, errorsOnField
}

func IsValidSemVer(i interface{}, fieldName string) (warnings []string, errorsOnField []error) {
	var semVerValues []string = strings.Split(i.(string), ".")

	if len(semVerValues) != 3 {
		errorsOnField = append(errorsOnField, fmt.Errorf("expected %q to be a valid semantic version MAJOR.MINOR.PATCH, got %q", fieldName, i.(string)))
		return warnings, errorsOnField
	}

	if major, ok := isValidVersion(semVerValues[0]); !ok {
		errorsOnField = append(errorsOnField, fmt.Errorf("expected %q to have minor version between 0 and 9, got %q", fieldName, strconv.FormatInt(major, 10)))
	}

	if minor, ok := isValidVersion(semVerValues[1]); !ok {
		errorsOnField = append(errorsOnField, fmt.Errorf("expected %q to have minor version between 0 and 9, got %q", fieldName, strconv.FormatInt(minor, 10)))
	}

	if patch, ok := isValidVersion(semVerValues[2]); !ok {
		errorsOnField = append(errorsOnField, fmt.Errorf("expected %q to have patch version between 0 and 9, got %q", fieldName, strconv.FormatInt(patch, 10)))
	}

	return warnings, errorsOnField
}

func isValidVersion(value string) (v int64, ok bool) {
	if version, err := strconv.ParseInt(value, 10, 64); err == nil {
		ok = true
		if version < 0 || version > 9 {
			ok = false
		}
		return version, ok
	}
	return 0, false
}

func FloatAtLeastAndLessThan(min, maxExclusive float64) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		v, ok := i.(float64)
		if !ok {
			es = append(es, fmt.Errorf("expected type of %s to be float64", k))
			return
		}

		if v < min || v >= maxExclusive {
			es = append(es, fmt.Errorf("expected %s to be at equal or greater than %f and strictly less than %f, got %f", k, min, maxExclusive, v))
			return
		}

		return
	}
}