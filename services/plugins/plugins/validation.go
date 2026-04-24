package plugins

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

type semVerForPluginsValidator struct{}

func (v semVerForPluginsValidator) Description(_ context.Context) string {
	return "must be a valid semantic version MAJOR.MINOR.PATCH with major version 1 or 2"
}

func (v semVerForPluginsValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v semVerForPluginsValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	for _, v := range helper.IsValidSemVer() {
		v.ValidateString(ctx, req, resp)
	}
	if resp.Diagnostics.HasError() || req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	major := strings.Split(req.ConfigValue.ValueString(), ".")[0]
	if major != "1" && major != "2" {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid semantic version",
			fmt.Sprintf("expected major version to be 1 or 2, got %q", major),
		)
	}
}

func isValidSemVerForPlugins() []validator.String {
	return []validator.String{semVerForPluginsValidator{}}
}
