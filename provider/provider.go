package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var resourceMap = make(map[string]provider.Resource)
var dataSourceMap = make(map[string]provider.DataSource)

// Subscribes this data type, adding it to the valid terraform resources and data types.
func (dataType DataSourceType[T, PT]) Subscribe() {
	resourceMap[dataType.ResourceIdentifier] = dataType.toResource()
	dataSourceMap[dataType.ResourceIdentifier] = dataType.toDataSource()
}

// Provider -
func Provider() provider.Provider {
	return &schema.Provider{
		Schema: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("HOST", nil),
				Description: "Only set this value if you are using a specific URL given by leanspace",
			},
			"region": schema.StringAttribute{
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("REGION", "eu-central-1"),
				Description: "Only set this value if you are using a specific region given by leanspace",
			},
			"env": schema.StringAttribute{
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ENV", "prod"),
				Description: "Only set this value if you are using a specific environment given by leanspace",
			},
			"tenant": schema.StringAttribute{
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TENANT", nil),
				Description: "The name given to your organization",
			},
			"client_id": schema.StringAttribute{
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_ID", nil),
				Description: "Client id of your Service Account",
			},
			"client_secret": schema.StringAttribute{
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

func providerConfigure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var host, env, tenant, clientId, clientSecret, region types.String

	diags := req.Config.Get(ctx, &host, &env, &tenant, &clientId, &clientSecret, &region)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !clientId.IsNull() && !clientSecret.IsNull() && !tenant.IsNull() {
		c, err := NewClient(&host.Value, &env.Value, &tenant.Value, &clientId.Value, &clientSecret.Value, &region.Value)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to create Leanspace client",
				err.Error(),
			)
			return
		}

		resp.DataSourceData = c
		resp.ResourceData = c
		return
	}

	c, err := NewClient(nil, nil, nil, nil, nil, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create Leanspace client",
			err.Error(),
		)
		return
	}

	resp.DataSourceData = c
	resp.ResourceData = c
}
