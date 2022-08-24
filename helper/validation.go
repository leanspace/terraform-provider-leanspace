package helper

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Condition interface {
	Eval(v map[string]any) bool
	Message() string
	Debug(v map[string]any) string
}

type Validator struct {
	If   Condition
	Iff  Condition
	Then Condition
}

type Validators []Validator

func (validators Validators) Check(obj map[string]any) error {
	errorMsg := ""
	for _, validator := range validators {
		switch {
		case validator.Iff != nil && validator.Then != nil:
			if validator.Iff.Eval(obj) != validator.Then.Eval(obj) {
				errorMsg += fmt.Sprintf(
					"  - if and only if %v, then %v\n"+
						"    got %v / %v\n",
					validator.Iff.Message(),
					validator.Then.Message(),
					validator.Iff.Debug(obj),
					validator.Then.Debug(obj),
				)
			}
		case validator.If != nil && validator.Then != nil:
			if validator.If.Eval(obj) && !validator.Then.Eval(obj) {
				errorMsg += fmt.Sprintf(
					"  - if %v, then %v\n"+
						"    got %v / %v\n",
					validator.If.Message(),
					validator.Then.Message(),
					validator.If.Debug(obj),
					validator.Then.Debug(obj),
				)
			}
		}
	}
	if errorMsg != "" {
		return fmt.Errorf("validation error(s):\n%v", errorMsg)
	}
	return nil
}

// Validators

type notCondition struct {
	cond Condition
}

func (c notCondition) Eval(v map[string]any) bool {
	return !c.cond.Eval(v)
}
func (c notCondition) Message() string {
	return fmt.Sprintf("not %v", c.cond.Message())
}
func (c notCondition) Debug(v map[string]any) string {
	return c.cond.Debug(v)
}
func Not(cond Condition) Condition {
	return notCondition{cond: cond}
}

type andCondition struct {
	conds []Condition
}

func (c andCondition) Eval(v map[string]any) bool {
	for _, cond := range c.conds {
		if !cond.Eval(v) {
			return false
		}
	}
	return true
}
func (c andCondition) Message() string {
	base := "("
	for i, cond := range c.conds {
		if i > 0 {
			base += " & "
		}
		base += cond.Message()
	}
	base += ")"
	return base
}
func (c andCondition) Debug(v map[string]any) string {
	base := ""
	isEmpty := true
	for _, cond := range c.conds {
		d := cond.Debug(v)
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
func And(conds ...Condition) Condition {
	return andCondition{conds: conds}
}

type orCondition struct {
	conds []Condition
}

func (c orCondition) Eval(v map[string]any) bool {
	for _, cond := range c.conds {
		if cond.Eval(v) {
			return true
		}
	}
	return false
}
func (c orCondition) Message() string {
	base := "("
	for i, cond := range c.conds {
		if i > 0 {
			base += " | "
		}
		base += cond.Message()
	}
	base += ")"
	return base
}
func (c orCondition) Debug(v map[string]any) string {
	base := ""
	isEmpty := true
	for _, cond := range c.conds {
		d := cond.Debug(v)
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
func Or(conds ...Condition) Condition {
	return orCondition{conds: conds}
}

type isSetCondition struct {
	key string
}

func (c isSetCondition) Eval(v map[string]any) bool {
	return v[c.key] != nil && v[c.key] != ""
}
func (c isSetCondition) Message() string {
	return fmt.Sprintf("%q is set", c.key)
}
func (c isSetCondition) Debug(v map[string]any) string {
	return fmt.Sprintf("%q = %q", c.key, v[c.key])
}
func IsSet(key string) Condition {
	return isSetCondition{key: key}
}

type isEqualsCondition struct {
	key   string
	value any
}

func (c isEqualsCondition) Eval(v map[string]any) bool {
	return v[c.key] == c.value
}
func (c isEqualsCondition) Message() string {
	return fmt.Sprintf("%q = %q", c.key, c.value)
}
func (c isEqualsCondition) Debug(v map[string]any) string {
	return fmt.Sprintf("%q = %q", c.key, v[c.key])
}
func Equals(key string, value any) Condition {
	return isEqualsCondition{key: key, value: value}
}

type isEmptyCondition struct {
	key string
}

func (c isEmptyCondition) Eval(v map[string]any) bool {
	if list, isList := v[c.key].([]any); isList {
		return len(list) == 0
	}
	if set, isSet := v[c.key].(*schema.Set); isSet {
		return set.Len() == 0
	}
	panic(fmt.Sprintf("Tried checking if %#v is empty", v[c.key]))
}
func (c isEmptyCondition) Message() string {
	return fmt.Sprintf("%q is empty", c.key)
}
func (c isEmptyCondition) Debug(v map[string]any) string {
	return fmt.Sprintf("%q = %q", c.key, v[c.key])
}
func IsEmpty(key string) Condition {
	return isEmptyCondition{key: key}
}
