package helper

import (
	"fmt"
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
	return ifCondition{if_: if_, then: then}
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
	return equivCondition{equivA: equivA, equivB: equivB}
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
	return notCondition{cond: cond}
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
	return andCondition{conds: conds}
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
	return orCondition{conds: conds}
}

type isSetCondition struct {
	key string
}

func (c isSetCondition) eval(v map[string]any) bool {
	return v[c.key] != nil && v[c.key] != ""
}
func (c isSetCondition) message() string {
	return fmt.Sprintf("%q is set", c.key)
}
func (c isSetCondition) debug(v map[string]any) string {
	return fmt.Sprintf("%q = %q", c.key, v[c.key])
}

// Will evaluate to true if the given key is set (ie. non-nil, not empty string)
func IsSet(key string) Condition {
	return isSetCondition{key: key}
}

type isEqualsCondition struct {
	key   string
	value any
}

func (c isEqualsCondition) eval(v map[string]any) bool {
	return v[c.key] == c.value
}
func (c isEqualsCondition) message() string {
	return fmt.Sprintf("%q = %q", c.key, c.value)
}
func (c isEqualsCondition) debug(v map[string]any) string {
	return fmt.Sprintf("%q = %q", c.key, v[c.key])
}

// Will evaluate to true if the value at the given key equals the given value
func Equals(key string, value any) Condition {
	return isEqualsCondition{key: key, value: value}
}

type isEmptyCondition struct {
	key string
}

func (c isEmptyCondition) eval(v map[string]any) bool {
	if list, isList := v[c.key].([]any); isList {
		return len(list) == 0
	}
	if set, isSet := v[c.key].(*schema.Set); isSet {
		return set.Len() == 0
	}
	panic(fmt.Sprintf("Tried checking if %#v is empty", v[c.key]))
}
func (c isEmptyCondition) message() string {
	return fmt.Sprintf("%q is empty", c.key)
}
func (c isEmptyCondition) debug(v map[string]any) string {
	return fmt.Sprintf("%q = %q", c.key, v[c.key])
}

// Will evaluate to true if the list/set at the given key is empty.
// Panics if something other than a list/set is found.
func IsEmpty(key string) Condition {
	return isEmptyCondition{key: key}
}
