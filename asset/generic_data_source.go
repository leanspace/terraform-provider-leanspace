package asset

import (
	"context"
	"strconv"
	"time"

	"github.com/mitchellh/mapstructure"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type DataSourceType[T any] struct {
	schema        map[string]*schema.Schema
	clientConvert func(*Client) GenericResourceType[T]
}

// A data source type for properties
var propertyDataSource = DataSourceType[Property[any]]{
	schema:        propertySchema,
	clientConvert: (*Client).forProperties,
}

// A data source type for nodes
var nodeDataSource = DataSourceType[Node]{
	schema:        nodeSchema,
	clientConvert: (*Client).forNodes,
}

// A data source type for command definitions
var commandDefinitionDataSource = DataSourceType[CommandDefinition]{
	schema:        commandDefinitionSchema,
	clientConvert: (*Client).forCommandDefinitions,
}

func (dataSourceType DataSourceType[T]) toDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceType.read,
		Schema: map[string]*schema.Schema{
			"content": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: dataSourceType.schema,
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
			"sort": sortSchema,
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
			"pageable": pageableSchema,
		},
	}
}

func (dataSourceType DataSourceType[T]) read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	values, err := dataSourceType.clientConvert(client).GetAll()
	if err != nil {
		return diag.FromErr(err)
	}
	dataSourceType.setData(values, d)
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func (dataSourceType DataSourceType[T]) setData(paginated_list *PaginatedList[T], d *schema.ResourceData) {

	data_as_map := make(map[string]any)
	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "terra",
		Result:  &data_as_map,
	})
	decoder.Decode(paginated_list)

	for key, value := range data_as_map {
		d.Set(key, value)
	}
}
