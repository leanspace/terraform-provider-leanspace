package nodes

import (
	"fmt"
	"regexp"
)

var tle1stLineRegex = regexp.MustCompile(`^1 (?P<noradId>[ 0-9]{5})[A-Z] [ 0-9]{5}[ A-Z]{3} [ 0-9]{5}[.][ 0-9]{8} (?:(?:[ 0+-][.][ 0-9]{8})|(?: [ +-][.][ 0-9]{7})) [ +-][ 0-9]{5}[+-][ 0-9] [ +-][ 0-9]{5}[+-][ 0-9] [ 0-9] [ 0-9]{4}[ 0-9]$`)
var tle2ndLineRegex = regexp.MustCompile(`^2 (?P<noradId>[ 0-9]{5}) [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{7} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{3}[.][ 0-9]{4} [ 0-9]{2}[.][ 0-9]{13}[ 0-9]$`)

func (node *Node) Validate() error {
	if node.Type == "ASSET" && !(node.Kind == "GENERIC" || node.Kind == "SATELLITE" || node.Kind == "GROUND_STATION") {
		return fmt.Errorf("kind must be either GENERIC, SATELLITE ou GROUND_STATION, got: %q", node.Kind)
	}

	if node.Kind != "SATELLITE" && node.Tle != nil && len(node.Tle) != 0 {
		return fmt.Errorf("TLE must only be specified for satellites, got %q", node.Tle)
	}

	if node.Kind == "SATELLITE" && node.Tle != nil {

		if len(node.Tle) != 2 {
			return fmt.Errorf("if a TLE is specified, it must contain only two strings, got %q", node.Tle)
		}

		if !tle1stLineRegex.MatchString(node.Tle[0]) {
			return fmt.Errorf("TLE first line mutch match %q, got: %q", tle1stLineRegex, node.Tle[0])
		}

		if !tle2ndLineRegex.MatchString(node.Tle[1]) {
			return fmt.Errorf("TLE first line mutch match %q, got: %q", tle2ndLineRegex, node.Tle[1])
		}
	}

	return nil
}
