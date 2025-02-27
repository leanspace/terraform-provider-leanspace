# terraform-provider-leanspace

This repository enables the use of Terraform for the different services of Leanspace.

## Requirements

- terraform (`choco install terraform` (windows) or `brew install hashicorp/tap/terraform` (mac)): >=1.5.0
- go (for plugin development): >=1.22

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

### Debugging

- If using VSCode, you can create a file in `./.vscode/launch.json` with the following input

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Debug Terraform Provider",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}",
            "env": {},
            "args": [
                "-debug",
            ]
        }
    ]
}
```

- In VSCode, start "Debug Terraform Provider"
- It will output something like `TF_REATTACH_PROVIDERS='{"registry.terraform.io/my-org/my-provider":{"Protocol":"grpc","Pid":3382870,"Test":true,"Addr":{"Network":"unix","String":"/tmp/plugin713096927"}}}'` in the Debug Console Tab
- Export this variable by doing `export TF_REATTACH_PROVIDERS=[...]`
- No need to initialize with `terraform init`, you can directly do `terraform plan` or `terraform apply`
- Like any other tool, once you do a modification you will have to restart the debugger (which means re-exporting the environment variable)

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

### Testing

Automatically a check will be performed to verify that:

- examples are present for each resource and data source
- testing are present for each resource and data source
- the implementation works (based on the testing folder) and that no modification is performed when reapplying

Since automated tests are run the following convention is to be followed:

- start the name of the resources with `Terraform`
- the name of states shall be `TERRAFORM_STATE`
- specific resources (`integration-leafspace/ground-stations/links` and `integration-leafspace/satellites/links`) does not follow these since they don't have names

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
- region: optional (defaults to `eu-central-1`)
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

There is also an `imports/main.tf` file, to test importing resources for the Topology & Assets service.
