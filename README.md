# terraform-provider-leanspace

This repository enables the use of Terraform for the different services of Leanspace.

## Requirements

- terraform (`choco install terraform` (windows) or `brew install terraform` (mac)): >=1.3.0
- go (for plugin development): >=1.20

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

- To set a specific version for the built binary, make sure you setup the environment variable `VERSION`.
For example, `export VERSION=0.7.0` with linux or `set VERSION=0.7.0` with windows.

- Run `make install` or `make install-windows` if you are on windows to apply the changes.

⚠️ If you ever encounter the following error:

```shell
error obtaining VCS status: exit status 128
    Use -buildvcs=false to disable VCS stamping.
```

Just run the following command:

```shell
git config --global --add safe.directory [your dir here]
```

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

1. Go to the `Actions` tab of this repository.
2. In the left side plane, click on `Terraform Provider Release Deployer` workflow.
3. This will open the workflow window. Now click on the `Run workflow` dropdown.
4. You'll now be presented with 2 selection boxes.
5. The first box allows you to choose the `release version` you wish to deploy. Please make sure it is in the format `v{marjor}.{minor}.{patch}`; i.e.: v0.4.0
6. The second box allows you to define `release version type` like what kind of version you are deploying i.e. is it a patch release, a minor one or a major one.

It will create a release in github and push this version to the Terraform Registry (see [Terraform](https://registry.terraform.io/providers/leanspace/leanspace/latest)).

### Run the plugin

You can run the examples (navigate to `testing`, if so you can modify the master `main.tf` to point to the correct resource) or create custom files.

Then run `terraform init && terraform apply --auto-approve`: this will create the required resources.
If you made some changes you can run it again to update the resources.

You can use `terraform init && terraform destroy --auto-approve`: this will delete the created resources.

You can also import existing resources (navigate to `testing/imports`):
You can use `terraform init && terraform import leanspace_nodes.sample_node 3563e0f6-03e3-416f-96f5-6c7102e37a11`: this will import the node with the id 3563e0f6-03e3-416f-96f5-6c7102e37a11 in the resource named sample_node

## Provider

The attributes are as follows:
- tenant: mandatory
- env: optional
- host: optional
- region: optional
- client_id: mandatory, refers to the client id of a service account
- client_secret: mandatory, refers to the client secret of a service account

This service account needs to have enough permissions (CRUD).

It is also possible to avoid passing this information in the provider by using environment variables as follows:
- TENANT
- ENV
- REGION
- HOST
- CLIENT_ID
- CLIENT_SECRET

## Resources

The resources reflect the objects of the different services, and for most of them mimics closely the APIs.
More information, as well as the API documentation, can be found [here](https://docs.leanspace.io/docs/services/)

## Documentation

The resources and datasource are explained in the [docs](https://docs.leanspace.io/docs/tools/) folder.
To build the documentation locally, clone and build the [docs plugin](https://github.com/leanspace/terraform-plugin-docs), and run the executable in the directory of the provider.

## Examples

You can find examples in the `/examples` folder (visible to the users) and `/testing`.

### Structure

The `main.tf` file imports all other modules. All modules are then organised per service, per resource: `testing/{service}/{resource}/main.tf`

The available resources per service are:
- activities:
  - activity definitions: it has one `leanspace_activity_definitions` resource, whth all possible metadata types (6) and all possible argument types (7), as well as two mappings
  - activity_states: it has one `leanspace_activity_states` resource.
- agents:
  - remote agents: it has one `leanspace_remote_agents` resource, with one inbound and one outbond connectors.
- asset:
  - nodes: it has 2 `leanspace_nodes` resources, one inside the other.
  - properties: it has as many `leanspace_properties` resources as available types
  - units: it generates 7 `leanspace_units` that are variants of a custom unit.
- commands:
  - command definition: it has one `leanspace_command_definitions` resource which has all possible metadata types (6) and all possible argument types (7)
  - command queue: it has one `leanspace_command_queues` resource which links the satellite and ground station nodes.
  - release queue: it has one `leanspace_release_queues` resource which links the satellite and a transformation strategy
  - command sequence state: it has one `leanspace_command_sequence_states` resource.
- dashboard:
  - widgets: it has as many `leanspace_widgets` resources as available types (5)
  - dashboards: it has one `leanspace_dashboards` resource with three widgets and linked to a node
- metrics:
  - metrics: it has as many `leanspace_metrics` resources as available types (7)
- monitors:
  - action_templates: it has one simple `leanspace_action_template` with a body and headers
  - monitors: it has two `leanspace_monitors`, one with and one without a tolerance set.
- orbits:
  - orbit_resources: it has one simple `leanspace_orbit_resources` with a body and headers
- pass:
  - pass_states: it has one `leanspace_pass_states` resource.
- plans:
  - plan_states: it has one `leanspace_plan_states` resource.
- plugins:
  - plugins: it has one `leanspace_plugins` resource, with basic filler data.
- routes:
  - processors: it has one `leanspace_processors` resource, with basic filler data.
  - routes: it has one `leanspace_routes` resource, with basic filler data.
- streams:
  - streams: it has one `leanspace_streams` resource, with all available element types (3), all possible field types (5), a computed field and a mapping.
- teams:
  - access policies: it has one `leanspace_access_policies` resource, with two statements, one containing a global (`*`) action and one with specific actions.
  - members: it has three `leanspace_members` resources, created recursively.
  - service accounts: it has three `leanspace_service_accounts` resources, created recursively.
  - teams: it has one `leanspace_teams` resource, created with the given members and policies.
- routes:
  - routes: it has one `leanspace_routes` resource.
- processors:
  - processors: it has one `leanspace_processors` resource.

There is also an `imports/main.tf` file, to test importing resources for the Topology & Assets service.
