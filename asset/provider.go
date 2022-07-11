package asset

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var resourceMap = make(map[string]*schema.Resource)
var dataSourceMap = make(map[string]*schema.Resource)

// Subscribes this data type, adding it to the valid terraform resources and data types.
func (dataType DataSourceType[T, PT]) Subscribe() {
	resourceMap[dataType.ResourceIdentifier] = dataType.resourceGenericDefinition()
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
			},
			"tenant": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TENANT", nil),
			},
			"client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_ID", nil),
			},
			"client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_SECRET", nil),
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

func TestProvider() (*schema.Provider, error) {
	providerConfigure := func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		env := "develop"
		tenant := "training"
		clientId := "1gqfr50rsdovh7i1qr3imda3qd"
		clientSecret := "1g7ogiq4i6jj3ebcd46tsjfqi5btft194raf0ht2jtmo7sb98d13"

		c, err := NewClient(nil, &env, &tenant, &clientId, &clientSecret)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return c, nil
	}

	provider := Provider()
	provider.ConfigureContextFunc = providerConfigure

	return provider, nil
}
