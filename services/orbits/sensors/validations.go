package sensors

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func validateMax(maximum float64) schema.SchemaValidateFunc {
	return validation.FloatBetween(0, maximum)
}
