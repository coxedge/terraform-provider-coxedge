package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

func dataSourceOrganization() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOrganizationRead,
		Schema:      getOrganizationSetSchema(),
	}
}

func dataSourceOrganizationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	requestedId := d.Get("id").(string)
	if requestedId != "" {
		org, err := coxEdgeClient.GetOrganization(requestedId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("organizations", flattenOrganizationData(&[]apiclient.Organization{*org})); err != nil {
			return diag.FromErr(err)
		}
	} else {
		orgs, err := coxEdgeClient.GetOrganizations()
		if err != nil {
			return diag.FromErr(err)
		}

		if err := d.Set("organizations", flattenOrganizationData(&orgs)); err != nil {
			return diag.FromErr(err)
		}
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenOrganizationData(organizations *[]apiclient.Organization) []interface{} {
	if organizations != nil {
		orgs := make([]interface{}, len(*organizations), len(*organizations))

		for i, org := range *organizations {
			item := make(map[string]interface{})

			item["id"] = org.Id
			item["name"] = org.Name
			item["entry_point"] = org.EntryPoint
			item["tags"] = org.Tags

			serviceConnectionList := make([]map[string]string, len(org.ServiceConnections))
			for i, serviceConnection := range org.ServiceConnections {
				serviceConnectionList[i] = make(map[string]string)
				serviceConnectionList[i]["id"] = serviceConnection.Id
				serviceConnectionList[i]["name"] = serviceConnection.Name
				serviceConnectionList[i]["service_code"] = serviceConnection.ServiceCode
			}
			item["service_connections"] = serviceConnectionList

			orgs[i] = item
		}

		return orgs
	}

	return make([]interface{}, 0)
}
