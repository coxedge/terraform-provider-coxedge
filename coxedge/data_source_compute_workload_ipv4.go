package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

func dataSourceComputeWorkloadIPv4() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeWorkloadIPv4Read,
		Schema:      getComputeWorkloadIPv4SetSchema(),
	}
}

func dataSourceComputeWorkloadIPv4Read(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	workloadId := data.Get("workload_id").(string)

	ipv4s, err := coxEdgeClient.GetComputeWorkloadIPv4ById(environmentName, organizationId, workloadId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("ipv4s", flattenComputeWorkloadIPv4Data(&ipv4s)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenComputeWorkloadIPv4Data(ipv4s *[]apiclient.NetworkConfiguration) []interface{} {
	if ipv4s != nil {
		ipv4List := make([]interface{}, len(*ipv4s), len(*ipv4s))

		for i, networkConfig := range *ipv4s {
			item := make(map[string]interface{})

			item["id"] = networkConfig.IP
			item["netmask"] = networkConfig.Netmask
			item["gateway"] = networkConfig.Gateway
			item["type"] = networkConfig.Type
			item["reverse"] = networkConfig.Reverse

			ipv4List[i] = item
		}

		return ipv4List
	}
	return make([]interface{}, 0)
}
