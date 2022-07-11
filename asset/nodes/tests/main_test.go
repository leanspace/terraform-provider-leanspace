package nodes_tests

import (
	"terraform-provider-asset/asset"
	"terraform-provider-asset/asset/nodes"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAccNodes_basic(t *testing.T) {
	nodeName := acctest.RandString(12)

	nodes.NodeDataType.Subscribe()

	resource.Test(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"leanspace": asset.TestProvider,
		},
		Steps: []resource.TestStep{
			{
				Config: testAccNodes_basicGenerateConfig(nodeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("leanspace_nodes.basic", "node.0.name", nodeName),
					resource.TestCheckResourceAttr("leanspace_nodes.basic", "node.0.description", "description of "+nodeName),
					resource.TestCheckResourceAttr("leanspace_nodes.basic", "node.0.type", "GROUP"),
					resource.TestCheckResourceAttr("leanspace_nodes.basic", "node.0.tags.#", "1"),
					resource.TestCheckResourceAttr("leanspace_nodes.basic", "node.0.tags.0.key", "test"),
					resource.TestCheckResourceAttr("leanspace_nodes.basic", "node.0.tags.0.value", nodeName),
				),
			},
		},
	})
}
