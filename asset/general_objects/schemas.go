package general_objects

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func PaginatedListSchema(content map[string]*schema.Schema) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"content": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: content,
			},
		},
		"total_elements": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"total_pages": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"number_of_elements": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"number": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"size": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"sort": SortSchema,
		"first": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"last": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"empty": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"pageable": PageableSchema,
	}
}

var SortSchema = &schema.Schema{
	Type:     schema.TypeList,
	Computed: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"direction": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"property": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ignore_case": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"null_handling": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ascending": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"descending": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	},
}

var PageableSchema = &schema.Schema{
	Type:     schema.TypeList,
	Computed: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"sort": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"direction": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"property": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ignore_case": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"null_handling": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ascending": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"descending": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"offset": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"page_number": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"page_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"paged": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"unpaged": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	},
}

var TagsSchema = &schema.Schema{
	Type:     schema.TypeSet,
	Optional: true,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	},
}
