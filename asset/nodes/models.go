package nodes

import "terraform-provider-asset/asset/general_objects"

type Node struct {
	ID                      string                `json:"id" terra:"id"`
	Name                    string                `json:"name" terra:"name"`
	Description             string                `json:"description,omitempty" terra:"description"`
	CreatedAt               string                `json:"createdAt" terra:"created_at"`
	CreatedBy               string                `json:"createdBy" terra:"created_by"`
	LastModifiedAt          string                `json:"lastModifiedAt" terra:"last_modified_at"`
	LastModifiedBy          string                `json:"lastModifiedBy" terra:"last_modified_by"`
	ParentNodeId            string                `json:"parentNodeId,omitempty" terra:"parent_node_id"`
	Tags                    []general_objects.Tag `json:"tags,omitempty" terra:"tags"`
	Nodes                   []Node                `json:"nodes,omitempty" terra:"nodes"`
	Type                    string                `json:"type" terra:"type"`
	Kind                    string                `json:"kind,omitempty" terra:"kind"`
	NoradId                 string                `json:"noradId,omitempty" terra:"norad_id"`
	InternationalDesignator string                `json:"internationalDesignator,omitempty" terra:"international_designator"`
	Tle                     []string              `json:"tle,omitempty" terra:"tle"`
	Latitude                float64               `json:"latitude,omitempty" terra:"latitude,omitempty"`
	Longitude               float64               `json:"longitude,omitempty" terra:"longitude,omitempty"`
	Elevation               float64               `json:"elevation,omitempty" terra:"elevation,omitempty"`
}

func (node *Node) GetID() string { return node.ID }
