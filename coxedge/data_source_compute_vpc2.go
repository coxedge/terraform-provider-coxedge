package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceComputeVPC2() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeVPC2Read,
		Schema:      getComputeVPC2SetSchema(),
	}
}

func dataSourceComputeVPC2Read(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	vpc2Id := data.Get("vpc2_id").(string)

	if vpc2Id != "" {
		computeVPC2, err := coxEdgeClient.GetComputeVPC2ById(environmentName, organizationId, vpc2Id)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("vpc2s", flattenComputeVPC2Data(&[]apiclient.ComputeVPC2{*computeVPC2})); err != nil {
			return diag.FromErr(err)
		}
	} else {
		computeVPC2, err := coxEdgeClient.GetComputeVPC2(environmentName, organizationId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("vpc2s", flattenComputeVPC2Data(&computeVPC2)); err != nil {
			return diag.FromErr(err)
		}
	}
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenComputeVPC2Data(dataSlice *[]apiclient.ComputeVPC2) []interface{} {
	if dataSlice != nil {
		flattened := make([]interface{}, len(*dataSlice))

		for i, instance := range *dataSlice {
			item := make(map[string]interface{})

			item["id"] = instance.ID
			item["date_created"] = instance.DateCreated
			item["region"] = instance.Region
			item["location"] = instance.Location
			item["description"] = instance.Description
			item["ip_block"] = instance.IPBlock
			item["prefix_length"] = instance.PrefixLength
			item["subnet"] = instance.Subnet

			flattened[i] = item
		}
		return flattened
	}
	return make([]interface{}, 0)
}
