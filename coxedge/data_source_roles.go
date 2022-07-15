package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

func dataSourceRoles() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRolesRead,
		Schema:      getRolesSetSchema(),
	}
}

func dataSourceRolesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	org, err := coxEdgeClient.GetRoles()
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("roles", flattenRolesData(&org)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenRolesData(roles *[]apiclient.Roles) []interface{} {
	if roles != nil {
		orgs := make([]interface{}, len(*roles), len(*roles))

		for i, org := range *roles {
			item := make(map[string]interface{})
			item["id"] = org.Id
			item["name"] = org.Name
			item["is_system"] = org.IsSystem
			item["default_scope"] = org.DefaultScope
			orgs[i] = item
		}
		return orgs
	}

	return make([]interface{}, 0)
}
