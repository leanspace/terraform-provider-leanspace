---
page_title: "leanspace_nodes Data Source - terraform-provider-leanspace"
subcategory: ""
description: |-
  
---

# leanspace_nodes (Data Source)



## Example Usage

```terraform
data "leanspace_nodes" "all" {
  filters {
    parent_node_ids = []
    property_ids    = []
    metric_ids      = []
    types           = ["ASSET"]
    kinds           = ["SATELLITE"]
    tags            = []
    ids             = []
    query           = ""
    page            = 0
    size            = 10
    sort            = ["name,asc"]
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

```json json_schema
{
	"properties": {
		"content": {
			"items": {
				"properties": {
					"created_at": {
						"description": "When it was created",
						"readOnly": true,
						"title": "created_at",
						"type": "string"
					},
					"created_by": {
						"description": "Who created it",
						"readOnly": true,
						"title": "created_by",
						"type": "string"
					},
					"description": {
						"title": "description",
						"type": "string"
					},
					"elevation": {
						"title": "elevation",
						"type": "number"
					},
					"id": {
						"readOnly": true,
						"title": "id",
						"type": "string"
					},
					"international_designator": {
						"title": "international_designator",
						"type": "string"
					},
					"kind": {
						"enum": [
							"GENERIC",
							"SATELLITE",
							"GROUND_STATION"
						],
						"title": "kind",
						"type": "string"
					},
					"last_modified_at": {
						"description": "When it was last modified",
						"readOnly": true,
						"title": "last_modified_at",
						"type": "string"
					},
					"last_modified_by": {
						"description": "Who modified it the last",
						"readOnly": true,
						"title": "last_modified_by",
						"type": "string"
					},
					"latitude": {
						"title": "latitude",
						"type": "number"
					},
					"longitude": {
						"title": "longitude",
						"type": "number"
					},
					"name": {
						"title": "name",
						"type": "string"
					},
					"nodes": {
						"items": {
							"properties": {
								"created_at": {
									"description": "When it was created",
									"readOnly": true,
									"title": "created_at",
									"type": "string"
								},
								"created_by": {
									"description": "Who created it",
									"readOnly": true,
									"title": "created_by",
									"type": "string"
								},
								"description": {
									"title": "description",
									"type": "string"
								},
								"elevation": {
									"title": "elevation",
									"type": "number"
								},
								"id": {
									"readOnly": true,
									"title": "id",
									"type": "string"
								},
								"international_designator": {
									"title": "international_designator",
									"type": "string"
								},
								"kind": {
									"enum": [
										"GENERIC",
										"SATELLITE",
										"GROUND_STATION"
									],
									"title": "kind",
									"type": "string"
								},
								"last_modified_at": {
									"description": "When it was last modified",
									"readOnly": true,
									"title": "last_modified_at",
									"type": "string"
								},
								"last_modified_by": {
									"description": "Who modified it the last",
									"readOnly": true,
									"title": "last_modified_by",
									"type": "string"
								},
								"latitude": {
									"title": "latitude",
									"type": "number"
								},
								"longitude": {
									"title": "longitude",
									"type": "number"
								},
								"name": {
									"title": "name",
									"type": "string"
								},
								"norad_id": {
									"description": "It must be 5 digits",
									"title": "norad_id",
									"type": "string"
								},
								"parent_node_id": {
									"title": "parent_node_id",
									"type": "string"
								},
								"tags": {
									"items": {
										"properties": {
											"key": {
												"title": "key",
												"type": "string"
											},
											"value": {
												"title": "value",
												"type": "string"
											}
										},
										"required": [
											"key"
										],
										"type": "object"
									},
									"title": "tags",
									"type": "array",
									"uniqueItems": true
								},
								"tle": {
									"description": "TLE composed of its 2 lines",
									"items": {
										"title": "tle",
										"type": "string"
									},
									"maxItems": 2,
									"minItems": 2,
									"title": "tle",
									"type": "array"
								},
								"type": {
									"enum": [
										"ASSET",
										"GROUP",
										"COMPONENT"
									],
									"title": "type",
									"type": "string"
								}
							},
							"required": [
								"created_at",
								"created_by",
								"id",
								"last_modified_at",
								"last_modified_by",
								"name",
								"type"
							],
							"type": "object"
						},
						"readOnly": true,
						"title": "nodes",
						"type": "array",
						"uniqueItems": true
					},
					"norad_id": {
						"description": "It must be 5 digits",
						"title": "norad_id",
						"type": "string"
					},
					"parent_node_id": {
						"title": "parent_node_id",
						"type": "string"
					},
					"tags": {
						"items": {
							"properties": {
								"key": {
									"title": "key",
									"type": "string"
								},
								"value": {
									"title": "value",
									"type": "string"
								}
							},
							"required": [
								"key"
							],
							"type": "object"
						},
						"title": "tags",
						"type": "array",
						"uniqueItems": true
					},
					"tle": {
						"description": "TLE composed of its 2 lines",
						"items": {
							"title": "tle",
							"type": "string"
						},
						"maxItems": 2,
						"minItems": 2,
						"title": "tle",
						"type": "array"
					},
					"type": {
						"enum": [
							"ASSET",
							"GROUP",
							"COMPONENT"
						],
						"title": "type",
						"type": "string"
					}
				},
				"required": [
					"created_at",
					"created_by",
					"id",
					"last_modified_at",
					"last_modified_by",
					"name",
					"nodes",
					"type"
				],
				"type": "object"
			},
			"readOnly": true,
			"title": "content",
			"type": "array"
		},
		"empty": {
			"description": "True if the content is empty",
			"readOnly": true,
			"title": "empty",
			"type": "boolean"
		},
		"filters": {
			"items": {
				"properties": {
					"ids": {
						"items": {
							"title": "ids",
							"type": "string"
						},
						"title": "ids",
						"type": "array"
					},
					"kinds": {
						"items": {
							"enum": [
								"GENERIC",
								"SATELLITE",
								"GROUND_STATION"
							],
							"title": "kinds",
							"type": "string"
						},
						"title": "kinds",
						"type": "array"
					},
					"metric_ids": {
						"items": {
							"title": "metric_ids",
							"type": "string"
						},
						"title": "metric_ids",
						"type": "array"
					},
					"page": {
						"default": 0,
						"title": "page",
						"type": "integer"
					},
					"parent_node_ids": {
						"items": {
							"title": "parent_node_ids",
							"type": "string"
						},
						"title": "parent_node_ids",
						"type": "array"
					},
					"property_ids": {
						"items": {
							"title": "property_ids",
							"type": "string"
						},
						"title": "property_ids",
						"type": "array"
					},
					"query": {
						"title": "query",
						"type": "string"
					},
					"size": {
						"default": 100,
						"title": "size",
						"type": "integer"
					},
					"sort": {
						"items": {
							"title": "sort",
							"type": "string"
						},
						"title": "sort",
						"type": "array"
					},
					"tags": {
						"items": {
							"title": "tags",
							"type": "string"
						},
						"title": "tags",
						"type": "array"
					},
					"types": {
						"items": {
							"enum": [
								"ASSET",
								"GROUP",
								"COMPONENT"
							],
							"title": "types",
							"type": "string"
						},
						"title": "types",
						"type": "array"
					}
				},
				"required": [],
				"type": "object"
			},
			"maxItems": 1,
			"minItems": 1,
			"title": "filters",
			"type": "array"
		},
		"first": {
			"description": "True if this is the first page",
			"readOnly": true,
			"title": "first",
			"type": "boolean"
		},
		"last": {
			"description": "True if this is the last page",
			"readOnly": true,
			"title": "last",
			"type": "boolean"
		},
		"number": {
			"description": "Page number",
			"readOnly": true,
			"title": "number",
			"type": "integer"
		},
		"number_of_elements": {
			"description": "Number of elements fetched in this page",
			"readOnly": true,
			"title": "number_of_elements",
			"type": "integer"
		},
		"pageable": {
			"items": {
				"properties": {
					"offset": {
						"description": "Number of elements in previous pages",
						"readOnly": true,
						"title": "offset",
						"type": "integer"
					},
					"page_number": {
						"description": "Page number",
						"readOnly": true,
						"title": "page_number",
						"type": "integer"
					},
					"page_size": {
						"description": "Size of this page",
						"readOnly": true,
						"title": "page_size",
						"type": "integer"
					},
					"paged": {
						"description": "True if this query is paged",
						"readOnly": true,
						"title": "paged",
						"type": "boolean"
					},
					"sort": {
						"items": {
							"properties": {
								"ascending": {
									"description": "True if the direction of the sorting is ascending",
									"readOnly": true,
									"title": "ascending",
									"type": "boolean"
								},
								"descending": {
									"description": "True if the direction of the sorting is descending",
									"readOnly": true,
									"title": "descending",
									"type": "boolean"
								},
								"direction": {
									"description": "Direction of the sorting, either DESC or ASC",
									"readOnly": true,
									"title": "direction",
									"type": "string"
								},
								"ignore_case": {
									"description": "True if the search ignores case",
									"readOnly": true,
									"title": "ignore_case",
									"type": "boolean"
								},
								"null_handling": {
									"description": "How null values are handled",
									"readOnly": true,
									"title": "null_handling",
									"type": "string"
								},
								"property": {
									"description": "Property used to sort by",
									"readOnly": true,
									"title": "property",
									"type": "string"
								}
							},
							"required": [
								"ascending",
								"descending",
								"direction",
								"ignore_case",
								"null_handling",
								"property"
							],
							"type": "object"
						},
						"readOnly": true,
						"title": "sort",
						"type": "array"
					},
					"unpaged": {
						"description": "True if this query is unpaged",
						"readOnly": true,
						"title": "unpaged",
						"type": "boolean"
					}
				},
				"required": [
					"offset",
					"page_number",
					"page_size",
					"paged",
					"sort",
					"unpaged"
				],
				"type": "object"
			},
			"readOnly": true,
			"title": "pageable",
			"type": "array"
		},
		"size": {
			"description": "Size of this page",
			"readOnly": true,
			"title": "size",
			"type": "integer"
		},
		"sort": {
			"items": {
				"properties": {
					"ascending": {
						"description": "True if the direction of the sorting is ascending",
						"readOnly": true,
						"title": "ascending",
						"type": "boolean"
					},
					"descending": {
						"description": "True if the direction of the sorting is descending",
						"readOnly": true,
						"title": "descending",
						"type": "boolean"
					},
					"direction": {
						"description": "Direction of the sorting, either DESC or ASC",
						"readOnly": true,
						"title": "direction",
						"type": "string"
					},
					"ignore_case": {
						"description": "True if the search ignores case",
						"readOnly": true,
						"title": "ignore_case",
						"type": "boolean"
					},
					"null_handling": {
						"description": "How null values are handled",
						"readOnly": true,
						"title": "null_handling",
						"type": "string"
					},
					"property": {
						"description": "Property used to sort by",
						"readOnly": true,
						"title": "property",
						"type": "string"
					}
				},
				"required": [
					"ascending",
					"descending",
					"direction",
					"ignore_case",
					"null_handling",
					"property"
				],
				"type": "object"
			},
			"readOnly": true,
			"title": "sort",
			"type": "array"
		},
		"total_elements": {
			"description": "Number of elements in total",
			"readOnly": true,
			"title": "total_elements",
			"type": "integer"
		},
		"total_pages": {
			"description": "Number of pages in total",
			"readOnly": true,
			"title": "total_pages",
			"type": "integer"
		}
	},
	"required": [
		"content",
		"empty",
		"first",
		"last",
		"number",
		"number_of_elements",
		"pageable",
		"size",
		"sort",
		"total_elements",
		"total_pages"
	],
	"title": "leanspace_nodes",
	"type": "object"
}
```