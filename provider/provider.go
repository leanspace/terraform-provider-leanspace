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
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("HOST", nil),
				Description: "Only set this value if you are using a specific URL given by leanspace",
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("REGION", "eu-central-1"),
				Description: "Only set this value if you are using a specific region given by leanspace",
			},
			"env": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ENV", "prod"),
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
	host := d.Get("host").(string)
	env := d.Get("env").(string)
	tenant := d.Get("tenant").(string)
	clientId := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)
	region := d.Get("region").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (clientId != "") && (clientSecret != "") && (tenant != "") {
		c, err := NewClient(&host, &env, &tenant, &clientId, &clientSecret, &region)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return c, diags
	}

	c, err := NewClient(nil, nil, nil, nil, nil, nil)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return c, diags
}
