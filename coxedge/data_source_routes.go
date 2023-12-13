package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceRoutes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRoutesRead,
		Schema:      getRoutesSetSchema(),
	}
}

func dataSourceRoutesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	environmentName := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)
	vpcId := d.Get("vpc_id").(string)

	org, err := coxEdgeClient.GetAllRoutes(vpcId, environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("routes", flattenRoutesData(&org)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenRoutesData(routes *[]apiclient.Route) []interface{} {
	if routes != nil {
		routesList := make([]interface{}, len(*routes), len(*routes))

		for i, route := range *routes {
			item := make(map[string]interface{})
			item["id"] = route.Id
			item["stack_id"] = route.StackId
			item["vpc_id"] = route.VpcId
			item["vpc_name"] = route.VpcName
			item["name"] = route.Name
			item["slug"] = route.Slug
			item["destination_cidrs"] = route.DestinationCidrs
			item["next_hops"] = route.NextHops
			item["status"] = route.Status

			routesList[i] = item
		}
		return routesList
	}

	return make([]interface{}, 0)
}
