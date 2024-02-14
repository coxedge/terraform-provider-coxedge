package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceComputeVPC() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeVPCRead,
		Schema:      getComputeVPCSetSchema(),
	}
}

func dataSourceComputeVPCRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	vpcId := data.Get("vpc_id").(string)

	if vpcId != "" {
		computeVPC, err := coxEdgeClient.GetComputeVPCById(environmentName, organizationId, vpcId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("vpcs", flattenComputeVPCData(&[]apiclient.ComputeVPC{*computeVPC})); err != nil {
			return diag.FromErr(err)
		}
	} else {
		computeVPC, err := coxEdgeClient.GetComputeVPC(environmentName, organizationId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("vpcs", flattenComputeVPCData(&computeVPC)); err != nil {
			return diag.FromErr(err)
		}
	}
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenComputeVPCData(dataSlice *[]apiclient.ComputeVPC) []interface{} {
	if dataSlice != nil {
		flattened := make([]interface{}, len(*dataSlice))

		for i, instance := range *dataSlice {
			item := make(map[string]interface{})

			item["id"] = instance.ID
			item["date_created"] = instance.DateCreated
			item["region"] = instance.Region
			item["location"] = instance.Location
			item["description"] = instance.Description
			item["v4_subnet"] = instance.V4Subnet
			item["v4_subnet_mask"] = instance.V4SubnetMask
			item["subnet"] = instance.Subnet

			flattened[i] = item
		}
		return flattened
	}
	return make([]interface{}, 0)
}
