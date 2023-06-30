package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

func dataSourceSubnets() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSubnetsRead,
		Schema:      getSubnetsSetSchema(),
	}
}

func dataSourceSubnetsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	environmentName := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)
	vpcId := d.Get("vpc_id").(string)

	org, err := coxEdgeClient.GetAllSubnets(vpcId, environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("subnets", flattenSubnetsData(&org)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenSubnetsData(subnets *[]apiclient.Subnets) []interface{} {
	if subnets != nil {
		subnetsList := make([]interface{}, len(*subnets), len(*subnets))

		for i, subnet := range *subnets {
			item := make(map[string]interface{})
			item["id"] = subnet.Id
			item["stack_id"] = subnet.StackId
			item["vpc_id"] = subnet.VpcId
			item["vpc_name"] = subnet.VpcName
			item["vpc_slug"] = subnet.VpcSlug
			item["name"] = subnet.Name
			item["slug"] = subnet.Slug
			item["cidr"] = subnet.Cidr
			item["status"] = subnet.Status

			subnetsList[i] = item
		}
		return subnetsList
	}

	return make([]interface{}, 0)
}
