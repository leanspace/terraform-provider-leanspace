package plugins

import (
	"fmt"
	"strings"

	"github.com/leanspace/terraform-provider-leanspace/helper"
)

func isValidSemVerForPlugins(i interface{}, fieldName string) (warnings []string, errorsOnField []error) {

	warnings, errors := helper.IsValidSemVer(i, fieldName)
	if errors != nil {
		for _, error := range errors {
			errorsOnField = append(errorsOnField, error)
		}
	}

	var semVerValues []string = strings.Split(i.(string), ".")
	if semVerValues[0] != "1" && semVerValues[0] != "2" {
		errorsOnField = append(errorsOnField, fmt.Errorf("expected %q to have major version between 1 and 2, got %q", fieldName, semVerValues[0]))
	}

	return warnings, errorsOnField
}
