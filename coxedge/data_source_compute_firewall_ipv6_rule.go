package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceComputeFirewallIPv6Rules() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeFirewallIPv6RulesRead,
		Schema:      getComputeFirewallIPRuleSetSchema(),
	}
}

func dataSourceComputeFirewallIPv6RulesRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	firewallId := data.Get("firewall_id").(string)
	ipv6Id := data.Get("ip_rule_id").(string)

	if ipv6Id != "" {
		computeFirewall, err := coxEdgeClient.GetComputeFirewallsIPv6RuleById(environmentName, organizationId, firewallId, ipv6Id)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("ip_rules", flattenComputeFirewallIPv6RulesData(&[]apiclient.ComputeFirewallRule{*computeFirewall})); err != nil {
			return diag.FromErr(err)
		}
	} else {
		computeFirewalls, err := coxEdgeClient.GetComputeFirewallsIPv6Rules(environmentName, organizationId, firewallId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("ip_rules", flattenComputeFirewallIPv6RulesData(&computeFirewalls)); err != nil {
			return diag.FromErr(err)
		}
	}
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags

}

func flattenComputeFirewallIPv6RulesData(computeFirewalls *[]apiclient.ComputeFirewallRule) []interface{} {
	if computeFirewalls != nil {
		firewalls := make([]interface{}, len(*computeFirewalls))

		for i, firewall := range *computeFirewalls {
			item := make(map[string]interface{})

			item["id"] = firewall.ID
			item["type"] = firewall.Type
			item["ip_type"] = firewall.IPType
			item["action"] = firewall.Action
			item["protocol"] = firewall.Protocol
			item["port"] = firewall.Port
			item["subnet"] = firewall.Subnet
			item["subnet_size"] = firewall.SubnetSize
			item["source"] = firewall.Source
			item["notes"] = firewall.Notes

			firewalls[i] = item
		}
		return firewalls
	}
	return make([]interface{}, 0)
}
