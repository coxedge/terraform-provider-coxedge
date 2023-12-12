package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceComputeWorkloadIPv6ReverseDNS() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeWorkloadIPv6ReverseDNSRead,
		Schema:      getComputeWorkloadIPv6ReverseDNSSetSchema(),
	}
}

func dataSourceComputeWorkloadIPv6ReverseDNSRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	workloadId := data.Get("workload_id").(string)

	ipv6ReverseDNSData, err := coxEdgeClient.GetComputeWorkloadIPv6ReverseDNSById(environmentName, organizationId, workloadId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("ipv6_reverse_dns", flattenComputeWorkloadIPv6ReverseDNSData(&ipv6ReverseDNSData)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenComputeWorkloadIPv6ReverseDNSData(ipv6s *[]apiclient.IPv6ReverseDNSConfiguration) []interface{} {
	if ipv6s != nil {
		ipv6ReversDNSList := make([]interface{}, len(*ipv6s), len(*ipv6s))

		for i, networkConfig := range *ipv6s {
			item := make(map[string]interface{})

			item["id"] = networkConfig.Id
			item["ip"] = networkConfig.Ip
			item["reverse"] = networkConfig.Reverse

			ipv6ReversDNSList[i] = item
		}
		return ipv6ReversDNSList
	}
	return make([]interface{}, 0)
}
