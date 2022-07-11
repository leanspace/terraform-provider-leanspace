package nodes_tests

import "fmt"

func testAccNodes_basicGenerateConfig(name string) string {
	return fmt.Sprintf(`

resource "leanspace_nodes" "basic" {
	node {
		name = "%[1]s"
		description = "description of %[1]s"
		type = "GROUP"
		tags {
			key = "test"
			value = "%[1]s"
		}
	}
}

	`, name)
}
