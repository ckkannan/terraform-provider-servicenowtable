package servicenowtable

import (
	"context"
	"fmt"

	// "fmt"

	"ckkannan/servicenowtable_client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var (
	_ datasource.DataSource              = &serviceNowTableDataSource{}
	_ datasource.DataSourceWithConfigure = &serviceNowTableDataSource{}
)

// Function defined as Data Source
func NewServiceNowTableDataSource() datasource.DataSource {
	return &serviceNowTableDataSource{}
}

// type servicenowTableOrgSourceModel struct {
// 	orgRows map[string]orgrowModel
// }

type servicenowTableOrgSourceModel struct {
	// Id      string                 `tfsdk:"id"`
	OrgRows map[string]orgrowModel `tfsdk:"datarows"`
}

type orgrowModel struct {
	Sys_id      types.String `tfsdk:"sys_id"`
	To_adgroup  types.String `tfsdk:"to_adgroup"`
	To_org_type types.String `tfsdk:"to_org_type"`
	To_org_name types.String `tfsdk:"to_org_name"`
}

type serviceNowTableDataSource struct {
	client *servicenowtable_client.Client
}

// Configure adds the provider configured client to the data source.
func (d *serviceNowTableDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*servicenowtable_client.Client)
}

func (d *serviceNowTableDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_queryorg"
}

func (d *serviceNowTableDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			// "id": schema.StringAttribute{
			// 	Computed: true,
			// },
			"datarows": schema.MapNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"to_adgroup": schema.StringAttribute{
							Computed: true,
						},
						"to_org_type": schema.StringAttribute{
							Computed: true,
						},
						"to_org_name": schema.StringAttribute{
							Computed: true,
						},
						"sys_id": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *serviceNowTableDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// var state servicenowTableOrgSourceModel
	// state := make(map[string]orgrowModel)
	var state servicenowTableOrgSourceModel
	// var state map[string]orgrowModel

	tflog.Info(ctx, "Configuring Tables client")
	d.client.Table = "x_22541_terraform_organization"
	d.client.Query = ""
	servicenowtablerows, err := d.client.GetOrgRows()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Organization Table", err.Error(),
		)
		return
	}

	// var f servicenowTableOrgSourceModel
	state.OrgRows = make(map[string]orgrowModel, len(servicenowtablerows))
	for k, r := range servicenowtablerows {
		rowState := orgrowModel{
			Sys_id:      types.StringValue(r.Sys_id),
			To_adgroup:  types.StringValue(r.To_adgroup),
			To_org_name: types.StringValue(r.To_org_name),
			To_org_type: types.StringValue(r.To_org_type),
		}
		fmt.Println("ROWSCK ", k, rowState.To_adgroup)

		// state.Id = k
		state.OrgRows[k] = rowState

	}
	// path.Root("datarows")
	tflog.Debug(ctx, "Creating HashiCups client")

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}
