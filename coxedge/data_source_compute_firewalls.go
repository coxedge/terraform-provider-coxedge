package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceComputeFirewalls() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeFirewallsRead,
		Schema:      getComputeFirewallSetSchema(),
	}
}

func dataSourceComputeFirewallsRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	firewallId := data.Get("firewall_id").(string)

	if firewallId != "" {
		computeFirewall, err := coxEdgeClient.GetComputeFirewallById(environmentName, organizationId, firewallId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("firewalls", flattenComputeFirewallsData(&[]apiclient.ComputeFirewall{*computeFirewall})); err != nil {
			return diag.FromErr(err)
		}
	} else {
		computeFirewalls, err := coxEdgeClient.GetComputeFirewalls(environmentName, organizationId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("firewalls", flattenComputeFirewallsData(&computeFirewalls)); err != nil {
			return diag.FromErr(err)
		}
	}
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags

}

func flattenComputeFirewallsData(computeFirewalls *[]apiclient.ComputeFirewall) []interface{} {
	if computeFirewalls != nil {
		firewalls := make([]interface{}, len(*computeFirewalls))

		for i, firewall := range *computeFirewalls {
			item := make(map[string]interface{})

			item["id"] = firewall.ID
			item["description"] = firewall.Description
			item["date_created"] = firewall.DateCreated
			item["date_modified"] = firewall.DateModified
			item["instance_count"] = firewall.InstanceCount
			item["rule_count"] = firewall.RuleCount
			item["max_rule_count"] = firewall.MaxRuleCount

			firewalls[i] = item
		}
		return firewalls
	}
	return make([]interface{}, 0)
}
