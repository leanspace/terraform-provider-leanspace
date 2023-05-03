package nodes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/leanspace/terraform-provider-leanspace/provider"
)

type apiShiftNodeInfo struct {
	TargetParentNodeId string `json:"targetParentNodeId"`
}

func (node *Node) toAPIFormat() ([]byte, error) {
	shiftNode := apiShiftNodeInfo{
		TargetParentNodeId: node.ParentNodeId,
	}
	return json.Marshal(shiftNode)
}

func nodeAction(action string, nodeId string, node *Node, client *provider.Client) error {
	data, err := node.toAPIFormat()
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%s/%s/%s/shift", client.HostURL, NodeDataType.Path, nodeId)
	req, err := http.NewRequest(action, path, strings.NewReader(string(data)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	_, err, _ = client.DoRequest(req, &(client).Token)
	if err != nil {
		return err
	}
	return nil
}

func shiftNode(nodeId string, node *Node, client *provider.Client) error {
	return nodeAction("PUT", nodeId, node, client)
}

func (node *Node) PostUpdateProcess(client *provider.Client, updated any) error {
	updatedNode := updated.(*Node)
	if err := shiftNode(updatedNode.ID, node, client); err != nil {
		return err
	}
	return nil
}
