# terraform

This repository enables the use of Terraform for the service Topology & Asset of Leanspace

## Requirements

- terraform (`choco install terraform` (windows) or `brew install terraform` (mac)): >=1.2.1
- go (for plugin development): >=1.18

## How to use

### Change plugin

Run `make install` or `make install-windows` if you are on windows to apply the changes.

### Run the plugin

Either run the examples (navigate to `examples`, if so you can modify the master `main.tf` to point to the correct resource) or create custom files.

Then run `terraform init && terraform apply --auto-approve`: this will create the required resources.
If you made some changes you can run it again to update the resources.

You can use `terraform init && terraform destroy --auto-approve`: this will delete the created resources.
You can also import existing resources (navigate to `examples/imports`)
You can use `terraform init && terraform import leanspace_nodes.sample_node 3563e0f6-03e3-416f-96f5-6c7102e37a11`: this will import the node with the id 3563e0f6-03e3-416f-96f5-6c7102e37a11 in the resource named sample_node

## Provider

The attributes are as follows:
- tenant: mandatory
- environment: optional
- client_id: mandatory, refers to the client id of a service account
- client_secret: mandatory, refers to the client secret of a service account

This service account needs to have enough permissions (CRUD).

## Resources

The resources reflects the object of Topology & Asset (https://api.leanspace.io/asset-repository/swagger-ui/index.html?configUrl=/asset-repository/v3/api-docs/swagger-config#/).

### Shared resources

tags:
- key
- value

### leanspace_nodes

One asset block containing:
- id (filled by the API)
- name
- description: optional
- created_at (filled by the API)
- created_by (filled by the API)
- last_modified_at (filled by the API)
- last_modified_by (filled by the API)
- last_modified_by (filled by the API)
- parent_node_id: optional, id of the parent node
- type: `GROUP||ASSET||COMPONENT`
- kind: optional if not an `ASSET`
- tags: optional, zero to multiple blocks
- nodes: optional, zero to multiple blocks
- norad_id: optional, only usefull for ASSET
- international_designator: optional, only usefull for ASSET
- tle: optional, only usefull for ASSET
    - list of exactly 2 strings

It is possible to create nodes within nodes but it's also possible to create them separately and set the `parent_node_id` on the child node to the id of the parent node

### leanspace_properties

One property block containing:
- id (filled by the API)
- name
- description: optional
- node_id: id of the node to attach to
- created_at (filled by the API)
- created_by (filled by the API)
- last_modified_at (filled by the API)
- last_modified_by (filled by the API)
- last_modified_by (filled by the API)
- type: `NUMERIC||TEXT||BOOLEAN||ENUM||TIMESTAMP||DATE||TIME||GEOPOINT`

For Numeric:
- min: optional, floating value
- max: optional, floating value
- scale: optional, integer value
- precision: optional, integer value
- unit_id: optional, refers to the id of a unit
- value: floating or integer value

For Text:
- minLength: optional, integer value
- maxLength: optional, integer value
- pattern: optional, string
- precision: optional, integer value
- value: string

For Boolean:
- value: boolean

For Enum:
- options: map of key/value pairs, the key is an integer, the value is a string
- value: integer, represents a key in the options

For Timestamp:
- before: optional, string date time
- after:optional, string date time
- value: string date time

For Date:
- before: optional, string date
- after:optional, string date
- value: string date

For Time:
- before: optional, string time
- after:optional, string time
- value: string time

For GeoPoint:
- fields: one block
    - elevation: one block
        - id (filled by the API)
        - name
        - description: optional
        - created_at (filled by the API)
        - created_by (filled by the API)
        - last_modified_at (filled by the API)
        - last_modified_by (filled by the API)
        - last_modified_by (filled by the API)
        - value: floating or integer value
        - type: string: `NUMERIC`
    - latitude: one block, same as elevation
    - longitude: one block, same as elevation

### leanspace_command_definitions

One command_definition block containing:
- id (filled by the API)
- node_id: id of the node to attach to
- name
- description: optional
- identifier: optional
- created_at (filled by the API)
- created_by (filled by the API)
- last_modified_at (filled by the API)
- last_modified_by (filled by the API)
- last_modified_by (filled by the API)
- metadata: one or multiple blocks
    - id (filled by the API)
    - name
    - description: optional
    - unit_id: optional
    - value: optional
    - required: optional
    - type: `NUMERIC||TEXT||BOOLEAN||TIMESTAMP||DATE||TIME`
- arguments: one or multiple blocks
    - id (filled by the API)
    - name
    - identifier
    - description: optional
    - value: optional
    - required: optional
    - type: `NUMERIC||TEXT||BOOLEAN||TIMESTAMP||DATE||TIME`

    For Numeric:
    - min: optional, floating value
    - max: optional, floating value
    - scale: optional, integer value
    - precision: optional, integer value
    - unit_id: optional, refers to the id of a unit
    - default_value: floating or integer value

    For Text:
    - minLength: optional, integer value
    - maxLength: optional, integer value
    - pattern: optional, string
    - precision: optional, integer value
    - default_value: string

    For Boolean:
    - default_value: boolean

    For Enum:
    - options: map of key/value pairs, the key is an integer, the value is a string
    - default_value: integer, represents a key in the options

    For Timestamp:
    - before: optional, string date time
    - after:optional, string date time
    - default_value: string date time

    For Date:
    - before: optional, string date
    - after:optional, string date
    - default_value: string date

    For Time:
    - before: optional, string time
    - after:optional, string time
    - default_value: string time

## Datasource

### Common pattern

- content: variable list of objects
- total_elements: integer
- total_pages: integer
- number_of_elements: integer
- number: integer
- size: integer
- sort
    - direction: string `ASC||DESC`
    - property: string
    - ignore_case: boolean
    - null_handling: string `NATIVE||NULLS_FIRST||NULLS_LAST`
    - ascending: boolean
    - descending: boolean
- first: boolean
- last: boolean
- empty: boolean
- pageable
    - sort: same as sort in the parent object
    - offset: integer
    - page_number: integer
    - page_size: integer
    - paged: boolean
    - unpaged: boolean

### leanspace_nodes

- content: list of one or multiple blocks of assets (snapshot representation of the resource `leanspace_nodes`)
    - id (filled by the API)
    - type: `GROUP||ASSET||COMPONENT`
    - kind: optional if not an `ASSET`
    - name
    - description: optional

### leanspace_properties

- content: list of one or multiple blocks of properties

### leanspace_command_definitions

- content: list of one or multiple blocks of command_definition

## Examples

You can find examples in the `/examples` folder

### Structure

There is the `main.tf` that defines which module it should other terraform file to call.

There's 3 folders for each resource:
- asset: it has 2 `leanspace_nodes` resources, the first one is a "normal" node and the second one has a node instead itself (thus creating 2 nodes)
- property: it has as many `leanspace_properties` resources as available types (8)
- command definition: it has as 1 `leanspace_command_definitions` resource which has all possible metadata types (6) and all possible argument types (7)

Finally there is the `imports` folder containing sample resources for each resource to test the import.