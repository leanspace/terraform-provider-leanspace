package routes

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type Route struct {
	general_objects.AuditModel
	Name           string                     `json:"name"`
	Description    *string                    `json:"description,omitempty"`
	Tags           []general_objects.KeyValue `json:"tags,omitempty"`
	Definition     Definition                 `json:"definition"`
	RouteInstances []RouteInstance            `json:"routeInstances,omitempty"`
	ProcessorIds   []string                   `json:"processorIds,omitempty"`
}

type Definition struct {
	Configuration    string  `json:"configuration"`
	LogLevel         string  `json:"logLevel"`
	Valid            bool    `json:"valid,omitempty"`
	ServiceAccountId *string `json:"serviceAccountId,omitempty"`
	Errors           []Error `json:"errors,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type RouteInstance struct {
	Status                    string  `json:"status"`
	LastStatusAt              *string `json:"lastStatusAt,omitempty"`
	ContainerId               string  `json:"containerId"`
	LastMessageStartProcessAt *string `json:"lastMessageStartProcessAt,omitempty"`
	LastMessageEndProcessAt   *string `json:"lastMessageEndProcessAt,omitempty"`
	NumberOfMessagesProcessed int     `json:"numberOfMessagesProcessed,omitempty"`
	CamelRouteId              *string `json:"camelRouteId,omitempty"`
}
