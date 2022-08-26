package helper

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Condition interface {
	eval(v map[string]any) bool
	message() string
	debug(v map[string]any) string
}

type Validators []Condition

func (validators Validators) Check(obj map[string]any) error {
	errorMsg := ""
	for _, validator := range validators {
		if !validator.eval(obj) {
			errorMsg += fmt.Sprintf(
				"  - %v\n"+
					"    got %v\n",
				validator.message(),
				validator.debug(obj),
			)
		}
	}
	if errorMsg != "" {
		return fmt.Errorf("validation error(s):\n%v", errorMsg)
	}
	return nil
}

// Conditions

type ifCondition struct {
	if_  Condition
	then Condition
}

func (c ifCondition) eval(v map[string]any) bool {
	return !c.if_.eval(v) || c.then.eval(v)
}
func (c ifCondition) message() string {
	return fmt.Sprintf("if %v then %v", c.if_.message(), c.then.message())
}
func (c ifCondition) debug(v map[string]any) string {
	return fmt.Sprintf("%v / %v", c.if_.debug(v), c.then.debug(v))
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
func (c equivCondition) message() string {
	return fmt.Sprintf("if and only if %v then %v", c.equivA.message(), c.equivB.message())
}
func (c equivCondition) debug(v map[string]any) string {
	return fmt.Sprintf("%v / %v", c.equivA.debug(v), c.equivB.debug(v))
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
func (c notCondition) message() string {
	return fmt.Sprintf("not %v", c.cond.message())
}
func (c notCondition) debug(v map[string]any) string {
	return c.cond.debug(v)
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
func (c andCondition) message() string {
	base := "("
	for i, cond := range c.conds {
		if i > 0 {
			base += " & "
		}
		base += cond.message()
	}
	base += ")"
	return base
}
func (c andCondition) debug(v map[string]any) string {
	base := ""
	isEmpty := true
	for _, cond := range c.conds {
		d := cond.debug(v)
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
func (c orCondition) message() string {
	base := "("
	for i, cond := range c.conds {
		if i > 0 {
			base += " | "
		}
		base += cond.message()
	}
	base += ")"
	return base
}
func (c orCondition) debug(v map[string]any) string {
	base := ""
	isEmpty := true
	for _, cond := range c.conds {
		d := cond.debug(v)
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
func (c isSetCondition) message() string {
	return fmt.Sprintf("%q is set", c.key)
}
func (c isSetCondition) debug(v map[string]any) string {
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
	return v[c.key] == c.value
}
func (c isEqualsCondition) message() string {
	return fmt.Sprintf("%q = %v", c.key, c.value)
}
func (c isEqualsCondition) debug(v map[string]any) string {
	return fmt.Sprintf("%q = %v", c.key, v[c.key])
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
func (c hasLengthCondition) message() string {
	return fmt.Sprintf("length(%q) = %v", c.key, c.length)
}
func (c hasLengthCondition) debug(v map[string]any) string {
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
func (c regexCondition) message() string {
	return fmt.Sprintf("%q matches %v", c.key, c.regex)
}
func (c regexCondition) debug(v map[string]any) string {
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
		return v[c.key].(T) < c.value
	case ">":
		return v[c.key].(T) > c.value
	case "<=":
		return v[c.key].(T) <= c.value
	case ">=":
		return v[c.key].(T) >= c.value
	default:
		panic(fmt.Sprintf("unrecognized operator %q", c.op))
	}
}

func (c compareCondition[T]) message() string {
	return fmt.Sprintf("%q %v %v", c.key, c.op, c.value)
}
func (c compareCondition[T]) debug(v map[string]any) string {
	return fmt.Sprintf("%q = %v", c.key, v[c.key])
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
