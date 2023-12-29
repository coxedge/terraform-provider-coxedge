package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceComputeWorkloadOS() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeWorkloadOSRead,
		Schema:      getComputeWorkloadOSSetSchema(),
	}
}

func dataSourceComputeWorkloadOSRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	workloadId := data.Get("workload_id").(string)

	os, err := coxEdgeClient.GetComputeWorkloadOSById(environmentName, organizationId, workloadId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("os", flattenComputeWorkloadOSData(&[]apiclient.ComputeWorkloadOS{*os})); err != nil {
		return diag.FromErr(err)
	}

	// always run
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenComputeWorkloadOSData(operatingSystem *[]apiclient.ComputeWorkloadOS) []interface{} {
	if operatingSystem != nil {
		osList := make([]interface{}, len(*operatingSystem), len(*operatingSystem))

		for i, os := range *operatingSystem {
			item := map[string]interface{}{
				"id":       os.ID,
				"plan_id":  os.PlanId,
				"os_label": os.OsLabel,
				"os_id":    os.OsID,
			}
			osList[i] = item
		}
		return osList
	}
	return make([]interface{}, 0)
}
