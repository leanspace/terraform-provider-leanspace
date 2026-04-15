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
		And(Or(Not(IsSet("tle")), IsEmpty("tle")), Not(IsSet("norad_id")), Not(IsSet("international_designator"))),
	),
	If(
		Not(Equals("kind", "GROUND_STATION")),
		And(Not(IsSet("latitude")), Not(IsSet("longitude")), Not(IsSet("elevation"))),
	),
	If(
		And(IsSet("tle"), Not(IsEmpty("tle"))),
		HasLength("tle", 2),
	),
}

func (node *Node) Validate() error {
	if err := nodeValidators.CheckValue(node); err != nil {
		return err
	}
	if node.Kind != nil && *node.Kind == "SATELLITE" && node.Tle != nil && len(node.Tle) >= 2 {
		if !tle1stLineRegex.MatchString(node.Tle[0]) {
			return fmt.Errorf("TLE first line must match %q, got: %q", tle1stLineRegex, node.Tle[0])
		}
		if !tle2ndLineRegex.MatchString(node.Tle[1]) {
			return fmt.Errorf("TLE second line must match %q, got: %q", tle2ndLineRegex, node.Tle[1])
		}
	}
	return nil
}
