package routes

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

func (route *Route) ToMap() map[string]any {
	routeMap := route.ToAuditMap()
	routeMap["name"] = helper.NilIfEmpty(route.Name)
	routeMap["description"] = helper.NilIfEmpty(route.Description)
	routeMap["tags"] = helper.ParseToMaps(route.Tags)
	routeMap["definition"] = route.Definition.ToMap()
	routeMap["route_instances"] = helper.ParseToMaps(route.RouteInstances)
	routeMap["processor_ids"] = helper.NilIfEmpty(route.ProcessorIds)
	return routeMap
}

func (definition *Definition) ToMap() map[string]any {
	definitionMap := make(map[string]any)
	definitionMap["configuration"] = helper.NilIfEmpty(definition.Configuration)
	definitionMap["log_level"] = helper.NilIfEmpty(definition.LogLevel)
	definitionMap["valid"] = helper.NilIfEmpty(definition.Valid)
	definitionMap["service_account_id"] = helper.NilIfEmpty(definition.ServiceAccountId)
	definitionMap["errors"] = helper.ParseToMaps(definition.Errors)

	return definitionMap
}

func (routeInstance *RouteInstance) ToMap() map[string]any {
	routeInstanceMap := make(map[string]any)
	routeInstanceMap["status"] = helper.NilIfEmpty(routeInstance.Status)
	routeInstanceMap["last_status_at"] = helper.NilIfEmpty(routeInstance.LastStatusAt)
	routeInstanceMap["container_id"] = helper.NilIfEmpty(routeInstance.ContainerId)
	routeInstanceMap["last_message_start_process_at"] = helper.NilIfEmpty(routeInstance.LastMessageStartProcessAt)
	routeInstanceMap["last_message_end_process_at"] = helper.NilIfEmpty(routeInstance.LastMessageEndProcessAt)
	routeInstanceMap["number_of_messages_processed"] = helper.NilIfEmpty(routeInstance.NumberOfMessagesProcessed)
	routeInstanceMap["camel_route_id"] = helper.NilIfEmpty(routeInstance.CamelRouteId)
	return routeInstanceMap
}

func (err *Error) ToMap() map[string]any {
	errMap := make(map[string]any)
	errMap["code"] = helper.NilIfEmpty(err.Code)
	errMap["message"] = helper.NilIfEmpty(err.Message)
	return errMap
}

func (route *Route) FromMap(routeMap map[string]any) error {
	route.FromAuditMap(routeMap)
	route.Name = helper.CastString(routeMap, "name")
	route.Description = helper.CastString(routeMap, "description")
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](helper.CastSlice(routeMap, "tags")); err != nil {
		return err
	} else {
		route.Tags = tags
	}
	if err := route.Definition.FromMap(helper.CastMapAny(routeMap, "definition")); err != nil {
		return err
	}
	route.ProcessorIds = make([]string, len(helper.CastSlice(routeMap, "processor_ids")))
	for i, processorId := range helper.CastSlice(routeMap, "processor_ids") {
		route.ProcessorIds[i] = processorId.(string)
	}
	return nil
}

func (definition *Definition) FromMap(definitionMap map[string]any) error {
	definition.Configuration = helper.CastString(definitionMap, "configuration")
	definition.LogLevel = helper.CastString(definitionMap, "log_level")
	definition.Valid = helper.CastBool(definitionMap, "valid")
	definition.ServiceAccountId = helper.CastString(definitionMap, "service_account_id")
	if errors, err := helper.ParseFromMaps[Error](helper.CastSlice(definitionMap, "errors")); err != nil {
		return err
	} else {
		definition.Errors = errors
	}

	return nil
}

func (routeInstance *RouteInstance) FromMap(routeInstanceMap map[string]any) error {
	routeInstance.Status = helper.CastString(routeInstanceMap, "status")
	routeInstance.LastStatusAt = helper.CastString(routeInstanceMap, "last_status_at")
	routeInstance.ContainerId = helper.CastString(routeInstanceMap, "container_id")
	routeInstance.LastMessageStartProcessAt = helper.CastString(routeInstanceMap, "last_message_start_process_at")
	routeInstance.LastMessageEndProcessAt = helper.CastString(routeInstanceMap, "last_message_end_process_at")
	routeInstance.NumberOfMessagesProcessed = helper.CastInt(routeInstanceMap, "number_of_messages_processed")
	routeInstance.CamelRouteId = helper.CastString(routeInstanceMap, "camel_route_id")

	return nil
}

func (err *Error) FromMap(errorMap map[string]any) error {
	err.Code = helper.CastString(errorMap, "code")
	err.Message = helper.CastString(errorMap, "message")
	return nil
}
