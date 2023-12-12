package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

func dataSourceComputeWorkloadFirewallGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeWorkloadFirewallGroupRead,
		Schema:      getComputeWorkloadFirewallGroupSetSchema(),
	}
}

func dataSourceComputeWorkloadFirewallGroupRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	workloadId := data.Get("workload_id").(string)

	firewallGroup, err := coxEdgeClient.GetComputeWorkloadFirewallGroupById(environmentName, organizationId, workloadId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("firewall_group", flattenComputeWorkloadFirewallGroupData(&[]apiclient.FirewallGroup{*firewallGroup})); err != nil {
		return diag.FromErr(err)
	}

	// always run
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenComputeWorkloadFirewallGroupData(firewallGroup *[]apiclient.FirewallGroup) []interface{} {
	if firewallGroup != nil {
		firewallGroupList := make([]interface{}, len(*firewallGroup), len(*firewallGroup))

		for i, networkConfig := range *firewallGroup {
			item := make(map[string]interface{})
			item["id"] = networkConfig.Id
			item["firewall_id"] = networkConfig.FirewallId
			item["name"] = networkConfig.Name
			firewallGroupList[i] = item
		}
		return firewallGroupList
	}
	return make([]interface{}, 0)
}
