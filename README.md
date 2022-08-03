# terraform-provider-leanspace

This repository enables the use of Terraform for the different services of Leanspace

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

You can also import existing resources (navigate to `examples/imports`):
You can use `terraform init && terraform import leanspace_nodes.sample_node 3563e0f6-03e3-416f-96f5-6c7102e37a11`: this will import the node with the id 3563e0f6-03e3-416f-96f5-6c7102e37a11 in the resource named sample_node

## Provider

The attributes are as follows:
- tenant: mandatory
- environment: optional
- client_id: mandatory, refers to the client id of a service account
- client_secret: mandatory, refers to the client secret of a service account

This service account needs to have enough permissions (CRUD).

## Resources

The resources reflect the objects of the different services, and for most of them mimics closely the APIs. 
More information, as well as the API documentation, can be found [here](https://docs.leanspace.io/docs/services/)

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
- parent_node_id: optional, id of the parent node
- type: `GROUP||ASSET||COMPONENT`
- kind: optional if not an `ASSET`
- tags: optional, zero to multiple blocks
- nodes: zero to multiple blocks (filled by the API)
- norad_id: optional, only useful for ASSET
- international_designator: optional, only useful for ASSET
- tle: optional, only useful for ASSET
    - list of exactly 2 strings
- latitude: required if kind = GROUND_STATION, float of the ground station's latitude
- longitude: required if kind = GROUND_STATION, float of the ground station's longitude
- elevation: required if kind = GROUND_STATION, float of the ground station's elevation

Nesting of nodes is not possible. Instead, set the `parent_node_id` field for the child node (see `examples/asset/nodes` for an example).

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
- metadata: one or multiple blocks
    - id (filled by the API)
    - name
    - description: optional
    - attributes: one block
        - unit_id: optional
        - value: optional
        - type: `NUMERIC||TEXT||BOOLEAN||TIMESTAMP||DATE||TIME`
- arguments: one or multiple blocks
    - id (filled by the API)
    - name
    - identifier
    - description: optional
    - attributes: one block
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
        - min_length: optional, integer value
        - max_length: optional, integer value
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
        - after: optional, string date time
        - default_value: string date time

        For Date:
        - before: optional, string date
        - after:optional, string date
        - default_value: string date

        For Time:
        - before: optional, string time
        - after:optional, string time
        - default_value: string time

### leanspace_command_queues

One command_queue block containing:
- id (filled by the API)
- asset_id: id of the node to attach to
- name
- ground_station_ids: optional, a list of strings containing the IDs of ground stations
- command_transformer_plugin_id: optional, UUID
- protocol_transformer_plugin_id: optional, UUID
- protocol_transformer_init_data: optional, string
- created_at (filled by the API)
- created_by (filled by the API)
- last_modified_at (filled by the API)
- last_modified_by (filled by the API)

### leanspace_metrics

One command_definition block containing:
- id (filled by the API)
- node_id: id of the node to attach to
- name
- description: optional
- created_at (filled by the API)
- created_by (filled by the API)
- last_modified_at (filled by the API)
- last_modified_by (filled by the API)
- attributes: one block
    - type: `NUMERIC||TEXT||BOOLEAN||TIMESTAMP||DATE||ENUM`

    For Numeric:
    - min: optional, floating value
    - max: optional, floating value
    - scale: optional, integer value
    - precision: optional, integer value
    - unit_id: optional, refers to the id of a unit

    For Text:
    - min_length: optional, integer value
    - max_length: optional, integer value
    - pattern: optional, string
    - precision: optional, integer value

    For Boolean:
    - No extra field

    For Enum:
    - options: map of key/value pairs, the key is an integer, the value is a string

    For Timestamp:
    - before: optional, string date time
    - after: optional, string date time

    For Date:
    - before: optional, string date
    - after:optional, string date

### leanspace_streams

One command_definition block containing:
- id (filled by the API)
- version (filled by the API)
- asset_id: id of the asset to attach to
- name
- description: optional
- created_at (filled by the API)
- created_by (filled by the API)
- last_modified_at (filled by the API)
- last_modified_by (filled by the API)
- configuration: one block
  - endianness: the default endianness of the stream (BE for big endian and LE for little endian)
  - structure: one block
    - elements: multiple blocks
      - name
      - path (filled by the API, derived from the terraform config)
      - order (filled by the API, derived from the terraform config)
      - type: `CONTAINER || FIELD || SWITCH`
      - valid (filled by the API)
      - errors (filled by the API)
      - for FIELD:
        - processor: optional
        - data_type: `INTEGER || UINTEGER || DECIMAL || TEXT || BOOLEAN`
        - length_in_bits
          Extra rules apply:
          - For data_type = `INTEGER || UINTEGER`, the max value is 32 bits
          - For data_type = `DECIMAL`, the value must be either 32 or 64 bits
        - endianness: field specific endiannes
      - for SWITCH:
        - expression: one block, required
          - switch_on: name of the field on which the switch applies
          - options: one block minimum
            - component: name of the container used for this case
            - value: one block, required
              - data_type: `INTEGER || UINTEGER || DECIMAL || TEXT || BOOLEAN`
              - data: the value to switch for
      - for SWITCH and CONTAINER:
        - elements: zero or more components
    - valid (filled by the API)
    - errors (filled by the API)
  - metadata: one block
    - packet_id: one block (filled by the API)
      - valid (filled by the API)
      - errors (filled by the API)
    - timestamp: optional, one block
      - expression: JS expression for the timecode
      - valid (filled by the API)
      - errors (filled by the API)
    - valid (filled by the API)
    - errors (filled by the API)
  - computations: one block
    - elements: 0 or more blocks
      - name
      - order (filled by the API, derived from the terraform config)
      - type: `COMPUTATION` (filled by the API)
      - data_type: `INTEGER || UINTEGER || DECIMAL || TEXT || BOOLEAN`
      - expression: JS expression to get the field
      - valid (filled by the API)
      - errors (filled by the API)
    - valid (filled by the API)
    - errors (filled by the API)
  - valid (filled by the API)
  - errors (filled by the API)
- mappings: multiple blocks
  - metric_id: ID of the metric to map to
  - name of the component to map from

### leanspace_widgets

One widget block containing:
- id (filled by the API)
- name
- description: optional
- type: `TABLE || LINE || BAR || AREA || VALUE`
- granularity: `second || minute || hour || day || week || month || raw`
- series: one or more blocks:
  - id: id of a metric if `datasource` = `metric` or the raw stream attribute if `datasource` = `raw_stream`
  - datasource: `metric || raw_stream`
  - aggregation: `avg || count || sum || min || max || none`
  - filters: 0 to 3 blocks:
    - filter_by: id of a metric if `datasource` = `metric` or the raw stream attribute if `datasource` = `raw_stream`
    - operator: `gt || lt || equals || notEquals`
    - value
- metrics: (filled by the API, derived from `series`) one or more blocks
  - id: id of the metric
  - aggregation: `avg || count || sum || min || max || none`
- metadata: optional
  - y_axis_label: optional string
  - y_axis_range_min: optional *list* of floats (one item inside)
  - y_axis_range_max: optional *list* of floats (one item inside)
- dashboards: (filled by the API) one or more blocks
  - id: id of dashboard
  - name: name of the dashboard
- tags: zero or more blocks
  - key
  - value
- created_at (filled by the API)
- created_by (filled by the API)
- last_modified_at (filled by the API)
- last_modified_by (filled by the API)

### leanspace_dashboards
One dashboard block containing:
- id (filled by the API)
- name
- description: optional
- node_ids: optional list of strings with node IDs
- widget_info: zero or more blocks
  - id: id of the widget 
  - type: `TABLE || LINE || BAR || AREA || VALUE` must match the widget
  - x: integer
  - y: integer
  - w: integer
  - h: integer
  - min_w: optional integer
  - min_h: optional integer
- widgets: wero or more blocks of type widget (see above) (filled by the API)
- tags: zero or more blocks
  - key
  - value
- created_at (filled by the API)
- created_by (filled by the API)
- last_modified_at (filled by the API)
- last_modified_by (filled by the API)

### leanspace_remote_agents
One remote agent block containing:
- id (filled by the API)
- name
- description: optional
- service_account_id: optional, the ID of the service account to link (can also be filled by the API)
- connectors: zero to many blocks:
  - id (filled by the API)
  - gateway_id
  - type `INBOUND || OUTBOUND`
  - socket: one block
    - type: `TCP || UDP`
    - host (only required if type = `OUTBOUND`)
    - port: integer
  - stream_id (only required if type = `INBOUND`)
  - destination (filled by the API) (only set if type = `INBOUND`):
    - type (filled by the API)
    - binding (filled by the API)
  - command_queue_id (only required if type = `OUTBOUND`)
  - source (filled by the API) (only set if type = `OUTBOUND`)
    - type (filled by the API)
    - binding (filled by the API)
  - created_at (filled by the API)
  - created_by (filled by the API)
  - last_modified_at (filled by the API)
  - last_modified_by (filled by the API)
- created_at (filled by the API)
- created_by (filled by the API)
- last_modified_at (filled by the API)
- last_modified_by (filled by the API)

### leanspace_access_policies
One access policy block containing:
- id (filled by the API)
- name
- description: optional
- read_only (filled by the API)
- statements: zero or more blocks
  - name
  - actions: a list of strings of permissions, that match the pattern `serviceName:permissionName` or `serviceName:*`
- created_at (filled by the API)
- created_by (filled by the API)
- last_modified_at (filled by the API)
- last_modified_by (filled by the API)

### leanspace_members
One member block containing:
- id (filled by the API)
- name
- email
- status (filled by the API)
- policy_ids: list of policy IDs (UUID strings) to attach
- created_at (filled by the API)
- created_by (filled by the API)
- last_modified_at (filled by the API)
- last_modified_by (filled by the API)

### leanspace_service_accounts
One service account block containing:
- id (filled by the API)
- name
- policy_ids: list of policy IDs (UUID strings) to attach
- created_at (filled by the API)
- created_by (filled by the API)
- last_modified_at (filled by the API)
- last_modified_by (filled by the API)

### leanspace_teams
One team block containing:
- id (filled by the API)
- name
- policy_ids: list of policy IDs (UUID strings) to attach
- members: list of member IDs (UUID strings) to add
- created_at (filled by the API)
- created_by (filled by the API)
- last_modified_at (filled by the API)
- last_modified_by (filled by the API)

### leanspace_activity_definitions

One command_definition block containing:
- id (filled by the API)
- node_id: id of the node to attach to
- name
- description: optional
- estimated_duration: optional, in seconds
- created_at (filled by the API)
- created_by (filled by the API)
- last_modified_at (filled by the API)
- last_modified_by (filled by the API)
- metadata: one or multiple blocks
    - id (filled by the API)
    - name
    - description: optional
    - attributes: one block
        - unit_id: optional
        - value: optional
        - type: `NUMERIC||TEXT||BOOLEAN||TIMESTAMP||DATE||TIME`
- arguments: one or multiple blocks
    - id (filled by the API)
    - name
    - description: optional
    - attributes: one block
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
        - min_length: optional, integer value
        - max_length: optional, integer value
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
        - after: optional, string date time
        - default_value: string date time

        For Date:
        - before: optional, string date
        - after:optional, string date
        - default_value: string date

        For Time:
        - before: optional, string time
        - after:optional, string time
        - default_value: string time
- command_mappings: one or multiple blocks
  - command_definition_id: id of a command definition for the mapping
  - position (filled by the API)
  - delay_in_milliseconds: required integer
  - argument_mappings: zero or more blocks
    - activity_definition_argument_name: name of an argument in the activity definition, required
    - command_definition_argument_name: name of an argument in the command definition, required
  - metadata_mappings: zero or more blocks
    - activity_definition_metadata_name: name of a metadata value in the activity definition, required
    - command_definition_argument_name: name of an argument in the command definition, required

### leanspace_plugins

One plugin block containing:
- id (filled by the API)
- type: `STRING_IDENTITY_PLUGIN_TYPE || JSON_IDENTITY_PLUGIN_TYPE || COMMANDS_COMMAND_TRANSFORMER_PLUGIN_TYPE || COMMANDS_PROTOCOL_TRANSFORMER_PLUGIN_TYPE || SIMULATIONS_ANALYTICAL_NOMINAL_PROPAGATION_PLUGIN_TYPE`
- implementation_class_name: required string (e.g. `org.myplugin.ClassName`)
- name
- description: optional
- source_code_file_download_authorized: optional bool, if the source can be downloaded (defaults to true)
- file_path: the absolute path to the file to upload (we recommend using the terraform function `abspath`)
- created_at (filled by the API)
- created_by (filled by the API)
- last_modified_at (filled by the API)
- last_modified_by (filled by the API)

### leanspace_action_templates

One action_template block containing:
- id (filled by the API)
- name
- type: can only be `WEBHOOK`, optional
- url: a required http / https url
- payload: the webhook payload (can be a `jsonencode` string)
- headers: an optional map of strings
- created_at (filled by the API)
- created_by (filled by the API)
- last_modified_at (filled by the API)
- last_modified_by (filled by the API)

### leanspce_monitors

One monitor block containing:
- id (filled by the API)
- name
- description (optional)
- status (filled by the API)
- polling_frequency_in_minutes: `1 || 60 || 1440`
- metric_id: UUID of the metric to monitor
- node_id (filled by the API)
- statistics: one block (filled by the API)
  - last_evaluation: one block
    - timestamp (datetime string)
    - value: float
    - status: string
- expression: one block
  - comparison_operator: one of `GREATER_THAN || LESSER_THAN || GREATER_THAN_OR_EQUAL_TO || LESSER_THAN_OR_EQUAL_TO || EQUAL_TO || NOT_EQUAL_TO`
  - comparison_value: float to compare
  - aggregation_function: one of `AVERAGE_VALUE || HIGHEST_VALUE || LOWEST_VALUE || SUM_VALUE || COUNT_VALUE`
  - tolerance: float of tolerance, only allowed if comparison_operator is `EQUAL_TO || NOT_EQUAL_TO`
- action_templates: set of action templates (see above) (filled by the API)
- action_template_ids: the set of desired action templates
- tags: zero or more tag blocks
- created_at (filled by the API)
- created_by (filled by the API)
- last_modified_at (filled by the API)
- last_modified_by (filled by the API)

## Datasource

### Common pattern

- content: variable list of the objects of the resource type (see above)
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
  - property: it has as many `leanspace_properties` resources as available types (8)
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
