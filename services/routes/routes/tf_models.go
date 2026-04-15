package routes

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

var errorAttrTypes = map[string]attr.Type{
	"code":    types.StringType,
	"message": types.StringType,
}

var routeInstanceAttrTypes = map[string]attr.Type{
	"status":                        types.StringType,
	"last_status_at":                types.StringType,
	"container_id":                  types.StringType,
	"last_message_start_process_at": types.StringType,
	"last_message_end_process_at":   types.StringType,
	"number_of_messages_processed":  types.Int64Type,
	"camel_route_id":                types.StringType,
}

type RouteTF struct {
	general_objects.AuditModelTF
	Name           types.String                 `tfsdk:"name"`
	Description    types.String                 `tfsdk:"description"`
	Tags           []general_objects.KeyValueTF `tfsdk:"tags"`
	Definition     *DefinitionTF                `tfsdk:"definition"`
	RouteInstances types.List                   `tfsdk:"route_instances"`
	ProcessorIds   []types.String               `tfsdk:"processor_ids"`
}

type DefinitionTF struct {
	Configuration    types.String `tfsdk:"configuration"`
	LogLevel         types.String `tfsdk:"log_level"`
	Valid            types.Bool   `tfsdk:"valid"`
	ServiceAccountId types.String `tfsdk:"service_account_id"`
	Errors           types.List   `tfsdk:"errors"`
}

func (x *Route) ToTF() any {
	errElems := make([]attr.Value, len(x.Definition.Errors))
	for i, e := range x.Definition.Errors {
		eObj, _ := types.ObjectValue(errorAttrTypes, map[string]attr.Value{
		"code":    helper.TFStringValue(e.Code),
			"message": helper.TFStringValue(e.Message),
		})
		errElems[i] = eObj
	}
	errors := types.ListValueMust(types.ObjectType{AttrTypes: errorAttrTypes}, errElems)

	riElems := make([]attr.Value, len(x.RouteInstances))
	for i, ri := range x.RouteInstances {
		riObj, _ := types.ObjectValue(routeInstanceAttrTypes, map[string]attr.Value{
			"status":                        helper.TFStringValue(ri.Status),
			"last_status_at":                helper.TFStringPtrValue(ri.LastStatusAt),
			"container_id":                  helper.TFStringValue(ri.ContainerId),
			"last_message_start_process_at": helper.TFStringPtrValue(ri.LastMessageStartProcessAt),
			"last_message_end_process_at":   helper.TFStringPtrValue(ri.LastMessageEndProcessAt),
			"number_of_messages_processed":  helper.TFInt64Value(ri.NumberOfMessagesProcessed),
			"camel_route_id":                helper.TFStringPtrValue(ri.CamelRouteId),
		})
		riElems[i] = riObj
	}
	routeInstances := types.ListValueMust(types.ObjectType{AttrTypes: routeInstanceAttrTypes}, riElems)

	return &RouteTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		Name:         helper.TFStringValue(x.Name),
		Description:  helper.TFStringPtrValue(x.Description),
		Tags:         general_objects.KeyValuesToTF(x.Tags),
		Definition: &DefinitionTF{
			Configuration:    helper.TFStringValue(x.Definition.Configuration),
			LogLevel:         helper.TFStringValue(x.Definition.LogLevel),
			Valid:            helper.TFBoolValue(x.Definition.Valid),
			ServiceAccountId: helper.TFStringPtrValue(x.Definition.ServiceAccountId),
			Errors:           errors,
		},
		RouteInstances: routeInstances,
		ProcessorIds:   helper.TFStringsValue(x.ProcessorIds),
	}
}

func (tf *RouteTF) ToAPI() any {
	var def Definition
	if tf.Definition != nil {
		def = Definition{
			Configuration:    helper.FromTFString(tf.Definition.Configuration),
			LogLevel:         helper.FromTFString(tf.Definition.LogLevel),
			Valid:            helper.FromTFBool(tf.Definition.Valid),
			ServiceAccountId: helper.FromTFStringPtr(tf.Definition.ServiceAccountId),
		}
	}

	return &Route{
		AuditModel:   general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:         helper.FromTFString(tf.Name),
		Description:  helper.FromTFStringPtr(tf.Description),
		Tags:         general_objects.KeyValuesFromTF(tf.Tags),
		Definition:   def,
		ProcessorIds: helper.FromTFStrings(tf.ProcessorIds),
	}
}
