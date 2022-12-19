---
page_title: "Provider: Leanspace"
---

# Leanspace Provider

The leanspace provider provides utilities for working with the
various resources available on the platform. It provides all
resources that are considered "static", ie. that are unlikely to
change frequently. For instance, command definitions are supported,
but not command instances.

Use the navigation to the left to read about the available resources,
and their data source counterparts.

## Example Usage

```terraform
terraform {
  required_providers {
    leanspace = {
      source = "app.terraform.io/leanspace/leanspace"
    }
  }
}

provider "leanspace" {
  tenant        = "my-org"
  client_id     = "client-id"
  client_secret = "client-secret"
}

resource "leanspace_nodes" "my_node" {
  name        = "MySatellite"
  description = "Using terraform is so easy!"
  type        = "ASSET"
  kind        = "SATELLITE"
}

resource "leanspace_properties" "mass_property" {
  name        = "Mass"
  description = "The mass of this satellite"
  node_id     = leanspace_nodes.my_node.id
  type        = "NUMERIC"
  value       = 800
}
```

<!-- schema generated by tfplugindocs -->
## Schema

```json json_schema
{
	"properties": {
		"client_id": {
			"description": "Client id of your Service Account",
			"title": "client_id",
			"type": "string"
		},
		"client_secret": {
			"description": "Client secret of your Service Account",
			"title": "client_secret",
			"type": "string"
		},
		"env": {
			"description": "Only set this value if you are using a specific environment given by leanspace",
			"title": "env",
			"type": "string"
		},
		"tenant": {
			"description": "The name given to your organization",
			"title": "tenant",
			"type": "string"
		}
	},
	"required": [],
	"title": "terraform-provider-leanspace",
	"type": "object"
}
```

## Limitations

### Syncing with the console

The resource created through this provider will be created on your
tenant, and will be accessible through the console. This also means
that this provider is not immune to name collisions! If you attempt
creating a resource that has the same name as an existing resource
on your tenant, an error will be thrown (usually with code `409`).
If this happens, either rename or delete (this can't be undone!) 
one of the two resources.