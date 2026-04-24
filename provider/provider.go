package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	providerschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var resourceFactories []func() resource.Resource
var dataSourceFactories []func() datasource.DataSource

func RegisterResource(factory func() resource.Resource) {
	resourceFactories = append(resourceFactories, factory)
}

func RegisterDataSource(factory func() datasource.DataSource) {
	dataSourceFactories = append(dataSourceFactories, factory)
}

// Provider
type LeanspaceProvider struct {
	client *Client
}

func New() func() provider.Provider {
	return func() provider.Provider {
		return &LeanspaceProvider{}
	}
}

func (p *LeanspaceProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "leanspace"
}

func (p *LeanspaceProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = providerschema.Schema{
		Attributes: map[string]providerschema.Attribute{
			"host": providerschema.StringAttribute{
				Optional:    true,
				Description: "Only set this value if you are using a specific URL given by leanspace",
			},
			"env": providerschema.StringAttribute{
				Optional:    true,
				Description: "Only set this value if you are using a specific environment given by leanspace",
			},
			"tenant": providerschema.StringAttribute{
				Optional:    true,
				Description: "The name given to your organization",
			},
			"client_id": providerschema.StringAttribute{
				Optional:    true,
				Description: "Client id of your Service Account",
			},
			"client_secret": providerschema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Client secret of your Service Account",
			},
		},
	}
}

type providerModel struct {
	Host         types.String `tfsdk:"host"`
	Env          types.String `tfsdk:"env"`
	Tenant       types.String `tfsdk:"tenant"`
	ClientId     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
}

func (p *LeanspaceProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config providerModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Apply environment variable defaults where config values are null
	host := stringValueOrEnv(config.Host, "HOST", "")
	env := stringValueOrEnv(config.Env, "ENV", "prod")
	tenant := stringValueOrEnv(config.Tenant, "TENANT", "")
	clientId := stringValueOrEnv(config.ClientId, "CLIENT_ID", "")
	clientSecret := stringValueOrEnv(config.ClientSecret, "CLIENT_SECRET", "")

	if clientId == "" || clientSecret == "" || tenant == "" {
		resp.Diagnostics.AddError("Missing Configuration", "Please provide tenant, client_id and client_secret")
		return
	}

	c, err := NewClient(&host, &env, &tenant, &clientId, &clientSecret)
	if err != nil {
		resp.Diagnostics.AddError("Failed to create client", err.Error())
		return
	}

	p.client = c
	resp.DataSourceData = c
	resp.ResourceData = c
}

func (p *LeanspaceProvider) Resources(_ context.Context) []func() resource.Resource {
	return resourceFactories
}

func (p *LeanspaceProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return dataSourceFactories
}

func stringValueOrEnv(val types.String, envVar string, defaultVal string) string {
	if !val.IsNull() && !val.IsUnknown() {
		return val.ValueString()
	}
	if v := os.Getenv(envVar); v != "" {
		return v
	}
	return defaultVal
}
