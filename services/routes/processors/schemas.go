package processors

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	resourceschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/leanspace/terraform-provider-leanspace/helper"
)

var processorSchema = map[string]resourceschema.Attribute{
	"id": resourceschema.StringAttribute{
		Computed: true,
	},
	"name": resourceschema.StringAttribute{
		Required: true,
	},
	"description": resourceschema.StringAttribute{
		Optional: true,
	},
	"version": resourceschema.StringAttribute{
		Required:      true,
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
	},
	"type": resourceschema.StringAttribute{
		Computed: true,
	},
	"file_path": resourceschema.StringAttribute{
		Required:      true,
		Description:   "It must be a valid path to a .jar file",
		PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
		Validators:    []validator.String{stringvalidator.RegexMatches(helper.PathToJarFileRegex, "Must be a valid file path")},
	},
	"created_at": resourceschema.StringAttribute{
		Computed:    true,
		Description: "When it was created",
	},
	"created_by": resourceschema.StringAttribute{
		Computed:    true,
		Description: "Who created it",
	},
	"last_modified_at": resourceschema.StringAttribute{
		Computed:    true,
		Description: "When it was last modified",
	},
	"last_modified_by": resourceschema.StringAttribute{
		Computed:    true,
		Description: "Who modified it the last",
	},
	"file_sha": resourceschema.StringAttribute{
		Computed:    true,
		Description: "Unique identifier of the processor file",
	},
}
