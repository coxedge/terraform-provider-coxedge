package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("COXEDGE_KEY", nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"coxedge_organizations": dataSourceOrganization(),
			"coxedge_environments":  dataSourceEnvironment(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"coxedge_environment": resourceEnvironment(),
			"coxedge_workload":    resourceWorkload(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiKey := d.Get("key").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if apiKey != "" {
		c := apiclient.NewClient(apiKey)

		return c, diags
	}

	return nil, diag.Errorf("No key set for key")
}
