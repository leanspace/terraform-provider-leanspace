package helper

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

// CheckValue is like Check but accepts any struct and converts it to a map[string]any
// automatically using reflection. PascalCase field names are mapped to snake_case keys,
// pointer values are dereferenced (nil pointers become nil), and anonymous embedded
// struct fields are flattened into the top-level map.
func (validators Validators) CheckValue(v any) error {
	return validators.Check(structToMap(v))
}

// camelToSnakeCase converts a PascalCase or camelCase identifier to snake_case.
func camelToSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) && i > 0 {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}

// derefValue dereferences a reflect.Value through any number of pointer indirections,
// returning nil if any pointer in the chain is nil.
func derefValue(rv reflect.Value) any {
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	return rv.Interface()
}

// structToMap converts a struct (or pointer to struct) to a map[string]any.
// Field names are converted from PascalCase to snake_case. Anonymous embedded
// struct fields are flattened into the top-level map.
func structToMap(v any) map[string]any {
	result := make(map[string]any)
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return result
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return result
	}
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		fieldVal := rv.Field(i)
		if field.Anonymous {
			for k, val := range structToMap(fieldVal.Interface()) {
				result[k] = val
			}
			continue
		}
		result[camelToSnakeCase(field.Name)] = derefValue(fieldVal)
	}
	return result
}

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
	val := v[c.key]
	if val == nil {
		return false
	}
	// Handle typed nils (e.g. (*int)(nil) stored as a non-nil interface).
	rv := reflect.ValueOf(val)
	if rv.Kind() == reflect.Ptr && rv.IsNil() {
		return false
	}
	return val != "" && val != 0 && val != '\x00' && val != 0.0
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
	val := v[c.key]
	// Handle untyped nil and typed nils (e.g. (*map[string]any)(nil)).
	if val == nil {
		return c.length == 0
	}
	rv := reflect.ValueOf(val)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return c.length == 0
		}
		rv = rv.Elem()
		val = rv.Interface()
	}
	if list, isList := val.([]any); isList {
		return len(list) == c.length
	}
	if list, isList := val.([]string); isList {
		return len(list) == c.length
	}
	if mapObj, isMap := val.(map[string]any); isMap {
		return len(mapObj) == c.length
	}
	if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Map {
		return rv.Len() == c.length
	}
	if rv.Kind() == reflect.Struct {
		// Structs have no "length"; treat a zero struct as empty (0) and any non-zero struct as present (1).
		if rv.IsZero() {
			return c.length == 0
		}
		return c.length == 1
	}
	panic(fmt.Sprintf("Tried checking length of %#v (only accepts lists and maps)", v[c.key]))
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

type timeDateOrTimestampValidator struct{}

func (v timeDateOrTimestampValidator) Description(_ context.Context) string {
	return "must be a valid date (2006-01-02), time (15:04:05), or RFC3339 timestamp"
}

func (v timeDateOrTimestampValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v timeDateOrTimestampValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	value := req.ConfigValue.ValueString()
	const dateLayoutReference = "2006-01-02"
	const timeLayoutReference = "15:04:05"
	const timestampLayoutReference = time.RFC3339

	_, errTimestamp := time.Parse(timestampLayoutReference, value)
	_, errDate := time.Parse(dateLayoutReference, value)
	_, errTime := time.Parse(timeLayoutReference, value)

	if errTimestamp != nil && errDate != nil && errTime != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid date, time, or timestamp",
			fmt.Sprintf("expected a valid date, time or timestamp, got %q:\n %+v\n %+v\n %+v", value, errTimestamp, errDate, errTime),
		)
	}
}

func IsValidTimeDateOrTimestamp() []validator.String {
	return []validator.String{timeDateOrTimestampValidator{}}
}

type semVerValidator struct{}

func (v semVerValidator) Description(_ context.Context) string {
	return "must be a valid semantic version in MAJOR.MINOR.PATCH format (each part 0-9)"
}

func (v semVerValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v semVerValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	value := req.ConfigValue.ValueString()
	semVerValues := strings.Split(value, ".")
	if len(semVerValues) != 3 {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid semantic version",
			fmt.Sprintf("expected a valid semantic version MAJOR.MINOR.PATCH, got %q", value),
		)
		return
	}
	for _, part := range semVerValues {
		if _, ok := isValidVersion(part); !ok {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid semantic version",
				fmt.Sprintf("expected each part to be between 0 and 9, got %q", part),
			)
		}
	}
}

func IsValidSemVer() []validator.String {
	return []validator.String{semVerValidator{}}
}

func isValidVersion(value string) (int64, bool) {
	version, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, false
	}
	return version, version >= 0 && version <= 9
}

type floatAtLeastAndLessThanValidator struct {
	min          float64
	maxExclusive float64
}

func (v floatAtLeastAndLessThanValidator) Description(_ context.Context) string {
	return fmt.Sprintf("must be at least %g and strictly less than %g", v.min, v.maxExclusive)
}

func (v floatAtLeastAndLessThanValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v floatAtLeastAndLessThanValidator) ValidateFloat64(_ context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	val := req.ConfigValue.ValueFloat64()
	if val < v.min || val >= v.maxExclusive {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid float value",
			fmt.Sprintf("expected value to be at least %g and strictly less than %g, got %g", v.min, v.maxExclusive, val),
		)
	}
}

func FloatAtLeastAndLessThan(min, maxExclusive float64) []validator.Float64 {
	return []validator.Float64{floatAtLeastAndLessThanValidator{min, maxExclusive}}
}

const stateNameRegexStr = `^[A-Z](?:[A-Z_]*[A-Z])?$`

var stateNameRegex = regexp.MustCompile(stateNameRegexStr)

const nameRegexStr = `^[ a-zA-Z0-9_-]*$`

var nameRegex = regexp.MustCompile(nameRegexStr)

func ValidStateName() []validator.String {
	return []validator.String{stringvalidator.RegexMatches(stateNameRegex, "")}
}

func ValidName() []validator.String {
	return []validator.String{stringvalidator.RegexMatches(nameRegex, "")}
}

const uuidRegexStr = `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`

var uuidRegex = regexp.MustCompile(uuidRegexStr)

func ValidUUID() []validator.String {
	return []validator.String{stringvalidator.RegexMatches(uuidRegex, "must be a valid UUID")}
}

// parentPathOf reconstructs the parent path by iterating path steps, dropping the last one.
func parentPathOf(p path.Path) path.Path {
	steps := p.Steps()
	if len(steps) == 0 {
		return path.Empty()
	}
	parent := path.Empty()
	for _, step := range steps[:len(steps)-1] {
		switch s := step.(type) {
		case path.PathStepAttributeName:
			parent = parent.AtName(string(s))
		case path.PathStepElementKeyInt:
			parent = parent.AtListIndex(int(s))
		case path.PathStepElementKeyString:
			parent = parent.AtMapKey(string(s))
		case path.PathStepElementKeyValue:
			parent = parent.AtSetValue(s.Value)
		}
	}
	return parent
}

// RequiredIfParentConfigured returns a String validator that mimics Required behaviour
// but only when the immediate parent block is actually configured (non-null).
//
// This is needed because terraform-plugin-framework evaluates Required inside
// SingleNestedBlock even when the block is absent (children are null). Making a
// field Optional + RequiredIfParentConfigured() preserves the "required within the
// block" semantics while not failing when the block itself is omitted.
func RequiredIfParentConfigured() validator.String {
	return requiredStringIfParentConfigured{}
}

type requiredStringIfParentConfigured struct{}

func (v requiredStringIfParentConfigured) Description(_ context.Context) string {
	return "Required when the parent block is configured."
}

func (v requiredStringIfParentConfigured) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v requiredStringIfParentConfigured) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	// Value is set — no issue.
	if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() {
		return
	}

	// Value is null/unknown — check whether the parent block is configured.
	parentPath := parentPathOf(req.Path)
	var parentVal attr.Value
	diags := req.Config.GetAttribute(ctx, parentPath, &parentVal)
	if diags.HasError() {
		return // can't determine parent state; skip
	}

	// Parent is absent or unknown — the block is not written by the user; no error.
	if parentVal == nil || parentVal.IsNull() || parentVal.IsUnknown() {
		return
	}

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Missing Required Value",
		"This attribute is required when the parent block is configured.",
	)
}

// RequiredFloat64IfParentConfigured returns a Float64 validator with the same
// "required within its optional parent block" semantics as RequiredIfParentConfigured.
func RequiredFloat64IfParentConfigured() validator.Float64 {
	return requiredFloat64IfParentConfigured{}
}

type requiredFloat64IfParentConfigured struct{}

func (v requiredFloat64IfParentConfigured) Description(_ context.Context) string {
	return "Required when the parent block is configured."
}

func (v requiredFloat64IfParentConfigured) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v requiredFloat64IfParentConfigured) ValidateFloat64(ctx context.Context, req validator.Float64Request, resp *validator.Float64Response) {
	if !req.ConfigValue.IsNull() && !req.ConfigValue.IsUnknown() {
		return
	}

	parentPath := parentPathOf(req.Path)
	var parentVal attr.Value
	diags := req.Config.GetAttribute(ctx, parentPath, &parentVal)
	if diags.HasError() {
		return
	}

	if parentVal == nil || parentVal.IsNull() || parentVal.IsUnknown() {
		return
	}

	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Missing Required Value",
		"This attribute is required when the parent block is configured.",
	)
}

// Ensure the types import is used (types.StringValue is referenced by path validators elsewhere).
var _ = types.StringNull
