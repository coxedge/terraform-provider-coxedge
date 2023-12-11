package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

func dataSourceComputeWorkloadIPv6() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeWorkloadIPv6Read,
		Schema:      getComputeWorkloadIPv6SetSchema(),
	}
}

func dataSourceComputeWorkloadIPv6Read(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	workloadId := data.Get("workload_id").(string)

	ipv6s, err := coxEdgeClient.GetComputeWorkloadIPv6ById(environmentName, organizationId, workloadId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("ipv6s", flattenComputeWorkloadIPv6Data(&ipv6s)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenComputeWorkloadIPv6Data(ipv6s *[]apiclient.IPv6Configuration) []interface{} {
	if ipv6s != nil {
		ipv6List := make([]interface{}, len(*ipv6s), len(*ipv6s))

		for i, networkConfig := range *ipv6s {
			item := make(map[string]interface{})

			item["id"] = networkConfig.IP
			item["ip"] = networkConfig.IP
			item["network"] = networkConfig.Network
			item["network_size"] = networkConfig.NetworkSize
			item["type"] = networkConfig.Type

			ipv6List[i] = item
		}

		return ipv6List
	}
	return make([]interface{}, 0)
}
