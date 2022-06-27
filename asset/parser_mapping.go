package asset

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type GenericResourceType[T any] struct {
	Client     *Client
	Path       string
	CreatePath func(T) string
}

type DataSourceType[T any] struct {
	Name       string // Will be used in the terraform file!
	Path       string
	CreatePath func(T) string

	Schema     map[string]*schema.Schema
	RootSchema map[string]*schema.Schema

	GetID       func(*T) string
	MapToStruct func(map[string]any) (T, error)
	StructToMap func(T) map[string]any
}

var NodeDataType = DataSourceType[Node]{
	Name: "node",
	Path: "asset-repository/nodes",

	Schema:     nodeSchema,
	RootSchema: rootNodeSchema,

	GetID:       func(n *Node) string { return n.ID },
	MapToStruct: nodeInterfaceToStruct,
	StructToMap: nodeStructToInterfaceBase,
}

var PropertyDataType = DataSourceType[Property[any]]{
	Name: "property",
	Path: "asset-repository/properties",
	CreatePath: func(p Property[any]) string {
		return fmt.Sprintf("asset-repository/nodes/%s/properties", p.NodeId)
	},

	Schema:     propertySchema,
	RootSchema: propertySchema,

	GetID:       func(p *Property[any]) string { return p.ID },
	MapToStruct: getPropertyData,
	StructToMap: propertyStructToInterface,
}

var CommandDataType = DataSourceType[CommandDefinition]{
	Name: "command_definition",
	Path: "asset-repository/command-definitions",
	CreatePath: func(c CommandDefinition) string {
		return fmt.Sprintf("asset-repository/nodes/%s/command-definitions", c.NodeId)
	},

	Schema:     commandDefinitionSchema,
	RootSchema: commandDefinitionSchema,

	GetID:       func(c *CommandDefinition) string { return c.ID },
	MapToStruct: getCommandDefinitionData,
	StructToMap: commandDefinitionStructToInterface,
}

func (dataSource DataSourceType[T]) convert(client *Client) GenericResourceType[T] {
	return GenericResourceType[T]{
		Client:     client,
		Path:       dataSource.Path,
		CreatePath: dataSource.CreatePath,
	}
}
