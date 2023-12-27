package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceBareMetalLocationProductOS() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBareMetalLocationProductOSRead,
		Schema:      getBareMetalLocationProductOSSetSchema(),
	}
}

func dataSourceBareMetalLocationProductOSRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	vendorProductId := data.Get("vendor_product_id").(string)

	bareMetalOSs, err := coxEdgeClient.GetBareMetalLocationProductOSs(environmentName, organizationId, vendorProductId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("baremetal_os", flattenBareMetalOSData(&bareMetalOSs)); err != nil {
		return diag.FromErr(err)
	}

	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenBareMetalOSData(oss *[]apiclient.BareMetalLocationProductOS) []interface{} {
	if oss != nil {
		osList := make([]interface{}, len(*oss), len(*oss))

		for i, os := range *oss {
			item := make(map[string]interface{})
			item["id"] = os.Id
			item["name"] = os.Name
			osList[i] = item
		}
		return osList
	}
	return make([]interface{}, 0)
}
