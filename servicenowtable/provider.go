package servicenowtable

import (
	"context"
	"os"

	"github.com/ckkannan/servicenowtable_client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &servicenowtableProvider{}
)

type servicenowtableProviderModel struct {
	SN_url  types.String `tfsdk:"sn_url"`
	SN_user types.String `tfsdk:"sn_user"`
	SN_pass types.String `tfsdk:"sn_pass"`
	SN_auth types.String `tfsdk:"sn_auth"`
}

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider {
	return &servicenowtableProvider{}
}

// servicenowtableProvider is the provider implementation.
type servicenowtableProvider struct{}

// Metadata returns the provider type name.
func (p *servicenowtableProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "servicenowtable"
}

// Schema defines the provider-level schema for configuration data.
func (p *servicenowtableProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"sn_url": schema.StringAttribute{
				Required: true,
			},
			"sn_user": schema.StringAttribute{
				Required: true,
			},
			"sn_pass": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"sn_auth": schema.StringAttribute{
				Optional:    true,
				Description: "For Basic auth set to Basic. As of now only basic auth is supported.",
			},
		},
	}
}

// Configure prepares a servicenowtable API client for data sources and resources.
func (p *servicenowtableProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config servicenowtableProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.SN_url.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("sn_url"),
			"Unknown servicenowtable API Host",
			"The provider cannot create the servicenowtable API client as there is an unknown configuration value for the servicenowtable API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SN_url environment variable.",
		)
	}

	if config.SN_user.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("sn_user"),
			"Unknown servicenowtable API Username",
			"The provider cannot create the servicenowtable API client as there is an unknown configuration value for the servicenowtable API username. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SN_user environment variable.",
		)
	}

	if config.SN_pass.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("sn_pass"),
			"Unknown servicenowtable API Password",
			"The provider cannot create the servicenowtable API client as there is an unknown configuration value for the servicenowtable API password. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SN_pass environment variable.",
		)
	}
	if config.SN_pass.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("sn_authtype"),
			"Unknown servicenowtable Auth method",
			"The provider cannot create the servicenowtable API client as there is an unknown configuration value for the servicenowtable Auth method. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SN_auth environment variable.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	sn_url := os.Getenv("SN_URL")
	sn_user := os.Getenv("SN_USER")
	sn_pass := os.Getenv("SN_PASS")
	sn_auth := os.Getenv("SN_AUTH")

	if !config.SN_url.IsNull() {
		sn_url = config.SN_url.ValueString()
	}

	if !config.SN_user.IsNull() {
		sn_user = config.SN_user.ValueString()
	}

	if !config.SN_pass.IsNull() {
		sn_pass = config.SN_pass.ValueString()
	}
	if !config.SN_auth.IsNull() {
		sn_auth = config.SN_auth.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if sn_url == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("sn_url"),
			"Missing servicenowtable API Host",
			"The provider cannot create the servicenowtable API client as there is a missing or empty value for the servicenowtable API host. "+
				"Set the host value in the configuration or use the SN_url environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if sn_user == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("sn_user"),
			"Missing servicenowtable API Username",
			"The provider cannot create the servicenowtable API client as there is a missing or empty value for the servicenowtable API username. "+
				"Set the username value in the configuration or use the SN_user environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if sn_pass == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("sn_pass"),
			"Missing servicenowtable API Password",
			"The provider cannot create the servicenowtable API client as there is a missing or empty value for the servicenowtable API password. "+
				"Set the password value in the configuration or use the SN_pass environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}
	if sn_auth == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("sn_auth"),
			"Missing servicenowtable Auth Type",
			"The provider cannot create the servicenowtable API client as there is a missing or empty value for the servicenowtable API authtype. "+
				"Set the password value in the configuration or use the SN_auth environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	sndev := servicenowtable_client.ServicenowtableProviderInput{Sn_url: sn_url, Sn_user: sn_user, Sn_pass: sn_pass, SSLIgnore: true, Version: "1.0"}
	// Create a new servicenowtable client using the configuration values
	client, err := servicenowtable_client.NewClient(sndev)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create servicenowtable API Client",
			"An unexpected error occurred when creating the servicenowtable API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"servicenowtable Client Error: "+err.Error(),
		)
		return
	}

	// Make the servicenowtable client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client
}

// DataSources defines the data sources implemented in the provider.
func (p *servicenowtableProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewServiceNowTableDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *servicenowtableProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
