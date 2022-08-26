package remote_agents

import (
	. "leanspace-terraform-provider/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var connectorValidators = Validators{
	Equivalence(
		Equals("type", "INBOUND"),
		IsSet("stream_id"),
	),
	Equivalence(
		Equals("type", "OUTBOUND"),
		IsSet("command_queue_id"),
	),
}

func (remoteAgent *RemoteAgent) Validate(data map[string]any) error {
	for _, connector := range data["connectors"].(*schema.Set).List() {
		if err := connectorValidators.Check(connector.(map[string]any)); err != nil {
			return err
		}
	}
	return nil
}
