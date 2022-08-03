package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var resourceMap = make(map[string]*schema.Resource)
var dataSourceMap = make(map[string]*schema.Resource)

// Subscribes this data type, adding it to the valid terraform resources and data types.
func (dataType DataSourceType[T, PT]) Subscribe() {
	resourceMap[dataType.ResourceIdentifier] = dataType.toResource()
	dataSourceMap[dataType.ResourceIdentifier] = dataType.toDataSource()
}

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"env": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ENV", nil),
				Description: "Only set this value if you are using a specific environment given by leanspace",
			},
			"tenant": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TENANT", nil),
				Description: "The name given to your organization",
			},
			"client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_ID", nil),
				Description: "Client id of your Service Account",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_SECRET", nil),
				Description: "Client secret of your Service Account",
			},
		},
		ResourcesMap:         resourceMap,
		DataSourcesMap:       dataSourceMap,
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
	env := d.Get("env").(string)
	tenant := d.Get("tenant").(string)
	clientId := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (clientId != "") && (clientSecret != "") && (tenant != "") {
		c, err := NewClient(nil, &env, &tenant, &clientId, &clientSecret)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return c, diags
	}

	c, err := NewClient(nil, nil, nil, nil, nil)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
