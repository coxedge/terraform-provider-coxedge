package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

func dataSourceVPCs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVPCsRead,
		Schema:      getVPCsSetSchema(),
	}
}

func dataSourceVPCsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	environmentName := d.Get("environment").(string)

	org, err := coxEdgeClient.GetAllVPCs(environmentName)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("vpcs", flattenVPCsData(&org)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenVPCsData(vpcs *[]apiclient.VPCs) []interface{} {
	if vpcs != nil {
		vpcsList := make([]interface{}, len(*vpcs), len(*vpcs))

		for i, org := range *vpcs {
			item := make(map[string]interface{})
			item["id"] = org.Id
			item["name"] = org.Name
			item["stack_id"] = org.StackId
			item["slug"] = org.Slug
			item["cidr"] = org.Cidr
			item["default_vpc"] = org.DefaultVpc
			item["created"] = org.Created
			item["status"] = org.Status

			subnetsList := make([]map[string]string, len(org.Subnets))
			for i, subnet := range org.Subnets {
				subnetsList[i] = make(map[string]string)
				subnetsList[i]["id"] = subnet.Id
				subnetsList[i]["stack_id"] = subnet.StackId
				subnetsList[i]["vpc_id"] = subnet.VpcId
				subnetsList[i]["name"] = subnet.Name
				subnetsList[i]["slug"] = subnet.Slug
				subnetsList[i]["cidr"] = subnet.Cidr
				subnetsList[i]["status"] = subnet.Status
			}
			item["subnets"] = subnetsList

			routesList := make([]map[string]interface{}, len(org.Routes))
			for i, route := range org.Routes {
				routesList[i] = make(map[string]interface{})
				routesList[i]["id"] = route.Id
				routesList[i]["stack_id"] = route.StackId
				routesList[i]["vpc_id"] = route.VpcId
				routesList[i]["name"] = route.Name
				routesList[i]["slug"] = route.Slug
				routesList[i]["destination_cidrs"] = route.DestinationCidrs
				routesList[i]["next_hops"] = route.NextHops
				routesList[i]["status"] = route.Status
			}
			item["routes"] = routesList

			vpcsList[i] = item
		}
		return vpcsList
	}

	return make([]interface{}, 0)
}
