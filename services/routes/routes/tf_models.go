package routes

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

type RouteTF struct {
	general_objects.AuditModelTF
	Name           types.String                 `tfsdk:"name"`
	Description    types.String                 `tfsdk:"description"`
	Tags           []general_objects.KeyValueTF `tfsdk:"tags"`
	Definition     *DefinitionTF                `tfsdk:"definition"`
	RouteInstances []RouteInstanceTF            `tfsdk:"route_instances"`
	ProcessorIds   []types.String               `tfsdk:"processor_ids"`
}

type DefinitionTF struct {
	Configuration    types.String `tfsdk:"configuration"`
	LogLevel         types.String `tfsdk:"log_level"`
	Valid            types.Bool   `tfsdk:"valid"`
	ServiceAccountId types.String `tfsdk:"service_account_id"`
	Errors           []ErrorTF    `tfsdk:"errors"`
}

type ErrorTF struct {
	Code    types.String `tfsdk:"code"`
	Message types.String `tfsdk:"message"`
}

type RouteInstanceTF struct {
	Status                    types.String `tfsdk:"status"`
	LastStatusAt              types.String `tfsdk:"last_status_at"`
	ContainerId               types.String `tfsdk:"container_id"`
	LastMessageStartProcessAt types.String `tfsdk:"last_message_start_process_at"`
	LastMessageEndProcessAt   types.String `tfsdk:"last_message_end_process_at"`
	NumberOfMessagesProcessed types.Int64  `tfsdk:"number_of_messages_processed"`
	CamelRouteId              types.String `tfsdk:"camel_route_id"`
}

func (x *Route) ToTF() any {
	errs := make([]ErrorTF, len(x.Definition.Errors))
	for i, e := range x.Definition.Errors {
		errs[i] = ErrorTF{
			Code:    helper.TFStringValue(e.Code),
			Message: helper.TFStringValue(e.Message),
		}
	}

	routeInstances := make([]RouteInstanceTF, len(x.RouteInstances))
	for i, ri := range x.RouteInstances {
		routeInstances[i] = RouteInstanceTF{
			Status:                    helper.TFStringValue(ri.Status),
			LastStatusAt:              helper.TFStringValue(ri.LastStatusAt),
			ContainerId:               helper.TFStringValue(ri.ContainerId),
			LastMessageStartProcessAt: helper.TFStringValue(ri.LastMessageStartProcessAt),
			LastMessageEndProcessAt:   helper.TFStringValue(ri.LastMessageEndProcessAt),
			NumberOfMessagesProcessed: helper.TFInt64Value(ri.NumberOfMessagesProcessed),
			CamelRouteId:              helper.TFStringValue(ri.CamelRouteId),
		}
	}

	return &RouteTF{
		AuditModelTF: general_objects.AuditModelToTF(&x.AuditModel),
		Name:         helper.TFStringValue(x.Name),
		Description:  helper.TFStringValue(x.Description),
		Tags:         general_objects.KeyValuesToTF(x.Tags),
		Definition: &DefinitionTF{
			Configuration:    helper.TFStringValue(x.Definition.Configuration),
			LogLevel:         helper.TFStringValue(x.Definition.LogLevel),
			Valid:            helper.TFBoolValue(x.Definition.Valid),
			ServiceAccountId: helper.TFStringValue(x.Definition.ServiceAccountId),
			Errors:           errs,
		},
		RouteInstances: routeInstances,
		ProcessorIds:   helper.TFStringsValue(x.ProcessorIds),
	}
}

func (tf *RouteTF) ToAPI() any {
	var def Definition
	if tf.Definition != nil {
		errs := make([]Error, len(tf.Definition.Errors))
		for i, e := range tf.Definition.Errors {
			errs[i] = Error{
				Code:    helper.FromTFString(e.Code),
				Message: helper.FromTFString(e.Message),
			}
		}
		def = Definition{
			Configuration:    helper.FromTFString(tf.Definition.Configuration),
			LogLevel:         helper.FromTFString(tf.Definition.LogLevel),
			Valid:            helper.FromTFBool(tf.Definition.Valid),
			ServiceAccountId: helper.FromTFString(tf.Definition.ServiceAccountId),
			Errors:           errs,
		}
	}

	routeInstances := make([]RouteInstance, len(tf.RouteInstances))
	for i, ri := range tf.RouteInstances {
		routeInstances[i] = RouteInstance{
			Status:                    helper.FromTFString(ri.Status),
			LastStatusAt:              helper.FromTFString(ri.LastStatusAt),
			ContainerId:               helper.FromTFString(ri.ContainerId),
			LastMessageStartProcessAt: helper.FromTFString(ri.LastMessageStartProcessAt),
			LastMessageEndProcessAt:   helper.FromTFString(ri.LastMessageEndProcessAt),
			NumberOfMessagesProcessed: helper.FromTFInt64(ri.NumberOfMessagesProcessed),
			CamelRouteId:              helper.FromTFString(ri.CamelRouteId),
		}
	}

	return &Route{
		AuditModel:     general_objects.AuditModelFromTF(tf.AuditModelTF),
		Name:           helper.FromTFString(tf.Name),
		Description:    helper.FromTFString(tf.Description),
		Tags:           general_objects.KeyValuesFromTF(tf.Tags),
		Definition:     def,
		RouteInstances: routeInstances,
		ProcessorIds:   helper.FromTFStrings(tf.ProcessorIds),
	}
}
