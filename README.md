# terraform-provider-leanspace

This repository enables the use of Terraform for the different services of Leanspace. The provider is hosted on a private registry in Terraform cloud

## Requirements

- terraform (`choco install terraform` (windows) or `brew install terraform` (mac)): >=1.2.1
- go (for plugin development): >=1.18

## Supported platform

- Windows
  - ARM 32 & 64 bits
  - x32 and x64
- Linux
  - ARM 32 & 64 bits
  - x32 and x64
- Mac OS
  - ARM 64 bits
  - x64
- FreeBSD
  - ARM 32 & 64 bits
  - x32 and x64

These platforms are defined in `.goreleaser.yml`.

## How to use

### Make modification

- Run `make install` or `make install-windows` if you are on windows to apply the changes.
- In the terraform files (.tf), modify the source to "leanspace.io/io/leanspace", this way it will try to find the provider in your local machine first

### Add a new service or a new resource

- Create a new folder under `services`
- Create a subfolder inside for the service folder with the name of the resource
- Create
  - main.go
  - models.go
  - parsers.go
  - requests.go (optionally)
  - schemas.go
- Add for each resource, the corresponding line to the root `main.go` file

### Create a new version

Create a tag on a commit with the following format `v{marjor}.{minor}.{patch}`; i.e.: v0.4.0; this will create a version accordingly.
It will create a release in github and push this version to the private repository in Terraform cloud (see [Terraform](https://app.terraform.io/app/leanspace/registry/providers/private/leanspace/leanspace/latest/overview)).

### Run the plugin

Since we host the provider in a private registry, you first have to login `terraform login` and then put a token to have access.

Then either run the examples (navigate to `examples`, if so you can modify the master `main.tf` to point to the correct resource) or create custom files.

Then run `terraform init && terraform apply --auto-approve`: this will create the required resources.
If you made some changes you can run it again to update the resources.

You can use `terraform init && terraform destroy --auto-approve`: this will delete the created resources.

You can also import existing resources (navigate to `examples/imports`):
You can use `terraform init && terraform import leanspace_nodes.sample_node 3563e0f6-03e3-416f-96f5-6c7102e37a11`: this will import the node with the id 3563e0f6-03e3-416f-96f5-6c7102e37a11 in the resource named sample_node

## Provider

The attributes are as follows:
- tenant: mandatory
- environment: optional
- client_id: mandatory, refers to the client id of a service account
- client_secret: mandatory, refers to the client secret of a service account

This service account needs to have enough permissions (CRUD).

It is also possible to avoid passing this information in the provider by using environment variables as follows:
- TENANT
- ENV
- CLIENT_ID
- CLIENT_SECRET

## Resources

The resources reflect the objects of the different services, and for most of them mimics closely the APIs.
More information, as well as the API documentation, can be found [here](https://docs.leanspace.io/docs/services/)

## Documentation

The resources and datasource are explained in the [docs](https://github.com/leanspace/terraform-provider-leanspace/blob/main/docs/index.md) folder.

## Examples

You can find examples in the `/examples` folder.

### Structure

The `main.tf` file imports all other modules. All modules are then organised per service, per resource: `examples/{service}/{resource}/main.tf`

The available resources per service are:
- activities:
  - activity definitions: it has one `leanspace_activity_definitions` resource, whth all possible metadata types (6) and all possible argument types (7), as well as two mappings
- agents:
  - remote agents: it has one `leanspace_remote_agents` resource, with one inbound and one outbond connectors.
- asset:
  - nodes: it has 2 `leanspace_nodes` resources, one inside the other.
  - properties: it has as many `leanspace_properties` resources as available types (8)
  - units: it generates 7 `leanspace_units` that are variants of a custom unit.
- commands:
  - command definition: it has one `leanspace_command_definitions` resource which has all possible metadata types (6) and all possible argument types (7)
  - command queue: it has one `leanspace_command_queues` resource which links the satellite and ground station nodes.
- dashboard:
  - widgets: it has as many `leanspace_widgets` resources as available types (5)
  - dashboards: it has one `leanspace_dashboards` resource with three widgets and linked to a node
- metrics:
  - metrics: it has as many `leanspace_metrics` resources as available types (6)
- monitors:
  - action_templates: it has one simple `leanspace_action_template` with a body and headers
  - monitors: it has two `leanspace_monitors`, one with and one without a tolerance set.
- plugins:
  - plugins: it has one `leanspace_plugins` resource, with basic filler data.
- streams:
  - streams: it has one `leanspace_streams` resource, with all available element types (3), all possible field types (5), a computed field and a mapping.
- teams:
  - access policies: it has one `leanspace_access_policies` resource, with two statements, one containing a global (`*`) action and one with specific actions.
  - members: it has three `leanspace_members` resources, created recursively.
  - service accounts: it has three `leanspace_service_accounts` resources, created recursively.  
  - teams: it has one `leanspace_teams` resource, created with the given members and policies.

There is also an `imports/main.tf` file, to test importing resources for the Topology & Assets service.
