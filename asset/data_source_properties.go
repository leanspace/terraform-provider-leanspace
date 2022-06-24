package asset

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProperties() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePropertiesRead,
		Schema: map[string]*schema.Schema{
			"properties": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: propertySchema,
				},
			},
			"total_elements": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_pages": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"number_of_elements": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"number": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"sort": sortSchema,
			"first": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"empty": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"pageable": pageableSchema,
		},
	}
}

func dataSourcePropertiesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	properties, err := client.forProperties().GetAll()
	if err != nil {
		return diag.FromErr(err)
	}
	setPropertiesData(properties, d)
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func setPropertiesData(properties *PaginatedList[Property[interface{}]], d *schema.ResourceData) {
	propertyList := make([]map[string]interface{}, len(properties.Content))

	for i, property := range properties.Content {
		propertyList[i] = propertyStructToInterface(property)
	}
	d.Set("properties", propertyList)

	d.Set("total_elements", properties.TotalElements)
	d.Set("total_pages", properties.TotalPages)
	d.Set("number_of_elements", properties.NumberOfElements)
	d.Set("number", properties.Number)
	d.Set("size", properties.Size)
	d.Set("first", properties.First)
	d.Set("last", properties.Last)
	d.Set("empty", properties.Empty)

	sort := sortStructToInterface(properties.Sort)
	d.Set("sort", sort)
	d.Set("pageable", pageableStructToInterface(properties.Pageable, sort))
}
