package nodes

import (
	"fmt"
	. "github.com/leanspace/terraform-provider-leanspace/helper"
)

var nodeValidators = Validators{
	Equivalence(
		Equals("type", "ASSET"),
		IsSet("kind"),
	),
	If(
		Not(Equals("kind", "SATELLITE")),
		Or(Not(IsSet("tle")), IsEmpty("tle")),
	),
	If(
		And(IsSet("tle"), Not(IsEmpty("tle"))),
		HasLength("tle", 2),
	),
}

func (node *Node) Validate(obj map[string]any) error {
	if err := nodeValidators.Check(obj); err != nil {
		return err
	}
	if node.Kind == "SATELLITE" && node.Tle != nil {
		if !tle1stLineRegex.MatchString(node.Tle[0]) {
			return fmt.Errorf("TLE first line mutch match %q, got: %q", tle1stLineRegex, node.Tle[0])
		}

		if !tle2ndLineRegex.MatchString(node.Tle[1]) {
			return fmt.Errorf("TLE first line mutch match %q, got: %q", tle2ndLineRegex, node.Tle[1])
		}
	}

	return nil
}
