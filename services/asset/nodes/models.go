package nodes

import "github.com/leanspace/terraform-provider-leanspace/helper/general_objects"

type Node struct {
	ID                      string                `json:"id"`
	Name                    string                `json:"name"`
	Description             string                `json:"description,omitempty"`
	CreatedAt               string                `json:"createdAt"`
	CreatedBy               string                `json:"createdBy"`
	LastModifiedAt          string                `json:"lastModifiedAt"`
	LastModifiedBy          string                `json:"lastModifiedBy"`
	ParentNodeId            string                `json:"parentNodeId,omitempty"`
	Tags                    []general_objects.Tag `json:"tags,omitempty"`
	Nodes                   []Node                `json:"nodes,omitempty"`
	Type                    string                `json:"type"`
	Kind                    string                `json:"kind,omitempty"`
	NoradId                 string                `json:"noradId,omitempty"`
	InternationalDesignator string                `json:"internationalDesignator,omitempty"`
	Tle                     []string              `json:"tle,omitempty"`
	Latitude                *float64              `json:"latitude,omitempty"`
	Longitude               *float64              `json:"longitude,omitempty"`
	Elevation               *float64              `json:"elevation,omitempty"`
	NumberOfChildren        int                   `json:"numberOfChildren"`
}

func (node *Node) GetID() string { return node.ID }
