package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceBareMetalLocations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBareMetalLocationsRead,
		Schema:      getBareMetalLocationsSetSchema(),
	}
}

func dataSourceBareMetalLocationsRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)

	bareMetalDeviceDisks, err := coxEdgeClient.GetBareMetalLocations(environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("baremetal_locations", flattenBareMetalLocationsData(&bareMetalDeviceDisks)); err != nil {
		return diag.FromErr(err)
	}

	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenBareMetalLocationsData(locations *[]apiclient.BareMetalLocation) []interface{} {
	if locations != nil {
		locationList := make([]interface{}, len(*locations), len(*locations))

		for i, location := range *locations {
			item := make(map[string]interface{})
			item["id"] = location.ID
			item["location_id"] = location.LocationID
			item["code"] = location.Code
			item["name"] = location.Name
			item["vendor"] = location.Vendor
			locationList[i] = item
		}
		return locationList
	}
	return make([]interface{}, 0)
}
