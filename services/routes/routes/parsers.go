package routes

import (
	"github.com/leanspace/terraform-provider-leanspace/helper"
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (route *Route) ToMap() map[string]any {
	routeMap := make(map[string]any)
	routeMap["id"] = route.ID
	routeMap["name"] = route.Name
	routeMap["description"] = route.Description
	routeMap["tags"] = helper.ParseToMaps(route.Tags)
	routeMap["definition"] = []any{route.Definition.ToMap()}
	routeMap["route_instances"] = helper.ParseToMaps(route.RouteInstances)
	routeMap["processor_ids"] = route.ProcessorIds

	routeMap["created_at"] = route.CreatedAt
	routeMap["created_by"] = route.CreatedBy
	routeMap["last_modified_at"] = route.LastModifiedAt
	routeMap["last_modified_by"] = route.LastModifiedBy

	return routeMap
}

func (definition *Definition) ToMap() map[string]any {
	definitionMap := make(map[string]any)
	definitionMap["configuration"] = definition.Configuration
	definitionMap["log_level"] = definition.LogLevel
	definitionMap["valid"] = definition.Valid
	definitionMap["service_account_id"] = definition.ServiceAccountId
	definitionMap["errors"] = helper.ParseToMaps(definition.Errors)

	return definitionMap
}

func (routeInstance *RouteInstance) ToMap() map[string]any {
	routeInstanceMap := make(map[string]any)
	routeInstanceMap["status"] = routeInstance.Status
	routeInstanceMap["last_status_at"] = routeInstance.LastStatusAt
	routeInstanceMap["container_id"] = routeInstance.ContainerId
	routeInstanceMap["last_message_start_process_at"] = routeInstance.LastMessageStartProcessAt
	routeInstanceMap["last_message_end_process_at"] = routeInstance.LastMessageEndProcessAt
	routeInstanceMap["number_of_messages_processed"] = routeInstance.NumberOfMessagesProcessed
	routeInstanceMap["camel_route_id"] = routeInstance.CamelRouteId
	return routeInstanceMap
}

func (err *Error) ToMap() map[string]any {
	errMap := make(map[string]any)
	errMap["code"] = err.Code
	errMap["message"] = err.Message
	return errMap
}

func (route *Route) FromMap(routeMap map[string]any) error {
	route.ID = routeMap["id"].(string)
	route.Name = routeMap["name"].(string)
	route.Description = routeMap["description"].(string)
	if tags, err := helper.ParseFromMaps[general_objects.KeyValue](routeMap["tags"].(*schema.Set).List()); err != nil {
		return err
	} else {
		route.Tags = tags
	}
	if len(routeMap["definition"].([]any)) > 0 {
		if err := route.Definition.FromMap(routeMap["definition"].([]any)[0].(map[string]any)); err != nil {
			return err
		}
	}
	route.ProcessorIds = make([]string, len(routeMap["processor_ids"].(*schema.Set).List()))
	for i, processorId := range routeMap["processor_ids"].(*schema.Set).List() {
		route.ProcessorIds[i] = processorId.(string)
	}
	route.CreatedAt = routeMap["created_at"].(string)
	route.CreatedBy = routeMap["created_by"].(string)
	route.LastModifiedAt = routeMap["last_modified_at"].(string)
	route.LastModifiedBy = routeMap["last_modified_by"].(string)

	return nil
}

func (definition *Definition) FromMap(definitionMap map[string]any) error {
	definition.Configuration = definitionMap["configuration"].(string)
	definition.LogLevel = definitionMap["log_level"].(string)
	definition.Valid = definitionMap["valid"].(bool)
	definition.ServiceAccountId = definitionMap["service_account_id"].(string)
	if errors, err := helper.ParseFromMaps[Error](definitionMap["errors"].(*schema.Set).List()); err != nil {
		return err
	} else {
		definition.Errors = errors
	}

	return nil
}

func (routeInstance *RouteInstance) FromMap(routeInstanceMap map[string]any) error {
	routeInstance.Status = routeInstanceMap["status"].(string)
	routeInstance.LastStatusAt = routeInstanceMap["last_status_at"].(string)
	routeInstance.ContainerId = routeInstanceMap["container_id"].(string)
	routeInstance.LastMessageStartProcessAt = routeInstanceMap["last_message_start_process_at"].(string)
	routeInstance.LastMessageEndProcessAt = routeInstanceMap["last_message_end_process_at"].(string)
	routeInstance.NumberOfMessagesProcessed = routeInstanceMap["number_of_messages_processed"].(int)
	routeInstance.CamelRouteId = routeInstanceMap["camel_route_id"].(string)

	return nil
}

func (err *Error) FromMap(errorMap map[string]any) error {
	err.Code = errorMap["code"].(string)
	err.Message = errorMap["message"].(string)
	return nil
}
