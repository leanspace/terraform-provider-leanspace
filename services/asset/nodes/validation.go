package nodes

import (
	"fmt"
	"regexp"

	. "github.com/leanspace/terraform-provider-leanspace/helper"
)

var tle1stLineRegex = regexp.MustCompile(`^1 (?P<noradId>[ 0-9]{5})[A-Z] [ 0-9]{5}[ A-Z]{3} [ 0-9]{5}[.][ 0-9]{8} (?:(?:[ 0+-][.][ 0-9]{8})|(?: [ +-][.][ 0-9]{7})) [ +-][ 0-9]{5}[+-][ 0-9] [ +-][ 0-9]{5}[+-][ 0-9] [ 0-9] [ 0-9]{4}[ 0-9]$`)
var tle2ndLineRegex = regexp.MustCompile(`^2 (?P<noradId>[ 0-9]{5}) [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{7} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{2}[.][ 0-9]{13}[ 0-9]$`)

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
