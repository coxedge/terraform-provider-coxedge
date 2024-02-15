package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceComputeReservedIP() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeReservedIPRead,
		Schema:      getComputeReservedIPSetSchema(),
	}
}

func dataSourceComputeReservedIPRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	reservedIPId := data.Get("reserved_ip_id").(string)

	if reservedIPId != "" {
		computeReservedIP, err := coxEdgeClient.GetComputeReservedIPById(environmentName, organizationId, reservedIPId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("reserved_ips", flattenComputeReservedIPData(&[]apiclient.ComputeReservedIP{*computeReservedIP})); err != nil {
			return diag.FromErr(err)
		}
	} else {
		computeReservedIP, err := coxEdgeClient.GetComputeReservedIPs(environmentName, organizationId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("reserved_ips", flattenComputeReservedIPData(&computeReservedIP)); err != nil {
			return diag.FromErr(err)
		}
	}
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenComputeReservedIPData(dataSlice *[]apiclient.ComputeReservedIP) []interface{} {
	if dataSlice != nil {
		flattened := make([]interface{}, len(*dataSlice))

		for i, instance := range *dataSlice {
			item := map[string]interface{}{
				"id":                 instance.ID,
				"region":             instance.Region,
				"location":           instance.Location,
				"ip_type":            instance.IPType,
				"subnet":             instance.Subnet,
				"subnet_size":        instance.SubnetSize,
				"label":              instance.Label,
				"instance_id":        instance.InstanceID,
				"reserved_ip":         instance.ReservedIP,
				"is_workload_attached": instance.IsWorkloadAttached,
			}

			flattened[i] = item
		}
		return flattened
	}
	return make([]interface{}, 0)
}
