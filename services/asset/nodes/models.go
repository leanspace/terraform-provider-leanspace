package nodes

import (
	"github.com/leanspace/terraform-provider-leanspace/helper/general_objects"
)

// Adding a properties attribute would require an optional "node_id" attribute in services/asset/properties/schemas.go whereas "node_id" is required during property creation.
// As a consequence, properties management would be too complicated.
// In the end, it was decided to not include a properties attribute in Node schema/struc to force the user to use the property resource.
type Node struct {
	ID                      string                     `json:"id"`
	Name                    string                     `json:"name"`
	Description             string                     `json:"description,omitempty"`
	CreatedAt               string                     `json:"createdAt"`
	CreatedBy               string                     `json:"createdBy"`
	LastModifiedAt          string                     `json:"lastModifiedAt"`
	LastModifiedBy          string                     `json:"lastModifiedBy"`
	ParentNodeId            string                     `json:"parentNodeId,omitempty"`
	Tags                    []general_objects.KeyValue `json:"tags,omitempty"`
	Nodes                   []Node                     `json:"nodes,omitempty"`
	Type                    string                     `json:"type"`
	Kind                    string                     `json:"kind,omitempty"`
	NoradId                 string                     `json:"noradId,omitempty"`
	InternationalDesignator string                     `json:"internationalDesignator,omitempty"`
	Tle                     []string                   `json:"tle,omitempty"`
	Latitude                *float64                   `json:"latitude,omitempty"`
	Longitude               *float64                   `json:"longitude,omitempty"`
	Elevation               *float64                   `json:"elevation,omitempty"`
	NumberOfChildren        int                        `json:"numberOfChildren"`
}

func (node *Node) GetID() string { return node.ID }
