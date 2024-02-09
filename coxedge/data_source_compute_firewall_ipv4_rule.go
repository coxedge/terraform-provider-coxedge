package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceComputeFirewallIPv4Rules() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeFirewallIPv4RulesRead,
		Schema:      getComputeFirewallIPv4RuleSetSchema(),
	}
}

func dataSourceComputeFirewallIPv4RulesRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	firewallId := data.Get("firewall_id").(string)
	ipv4Id := data.Get("ipv4_rule_id").(string)

	if ipv4Id != "" {
		computeFirewall, err := coxEdgeClient.GetComputeFirewallsIPv4RuleById(environmentName, organizationId, firewallId, ipv4Id)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("ipv4_rules", flattenComputeFirewallIPv4RulesData(&[]apiclient.ComputeFirewallRule{*computeFirewall})); err != nil {
			return diag.FromErr(err)
		}
	} else {
		computeFirewalls, err := coxEdgeClient.GetComputeFirewallsIPv4Rules(environmentName, organizationId, firewallId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("ipv4_rules", flattenComputeFirewallIPv4RulesData(&computeFirewalls)); err != nil {
			return diag.FromErr(err)
		}
	}
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags

}

func flattenComputeFirewallIPv4RulesData(computeFirewalls *[]apiclient.ComputeFirewallRule) []interface{} {
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
