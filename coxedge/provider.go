/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("COXEDGE_KEY", nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"coxedge_organizations":                   dataSourceOrganization(),
			"coxedge_organizations_billing_info":      dataSourceOrganizationBillingInfo(),
			"coxedge_environments":                    dataSourceEnvironment(),
			"coxedge_images":                          dataSourceImage(),
			"coxedge_origin_settings":                 dataSourceOriginSetting(),
			"coxedge_roles":                           dataSourceRoles(),
			"coxedge_workload_instances":              dataWorkloadInstances(),
			"coxedge_sites_predefined_edge_logic":     dataSourceSitesEdgeLogic(),
			"coxedge_sites_edge_logic_delivery_rules": dataSourceSitesEdgeLogicDeliveryRules(),
			"coxedge_sites_edge_logic_custom_rules":   dataSourceSitesEdgeLogicCustomRules(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"coxedge_cdn_purge":                      resourceCDNPurgeResource(),
			"coxedge_cdn_settings":                   resourceCDNSettings(),
			"coxedge_delivery_domain":                resourceDeliveryDomain(),
			"coxedge_environment":                    resourceEnvironment(),
			"coxedge_firewall_rule":                  resourceFirewallRule(),
			"coxedge_network_policy_rule":            resourceNetworkPolicyRule(),
			"coxedge_origin_setting":                 resourceOriginSettings(),
			"coxedge_script":                         resourceScript(),
			"coxedge_site":                           resourceSite(),
			"coxedge_user":                           resourceUser(),
			"coxedge_waf_settings":                   resourceWAFSettings(),
			"coxedge_workload":                       resourceWorkload(),
			"coxedge_sites_predefined_edge_logic":    resourceSitesPredefinedEdgeLogic(),
			"coxedge_sites_edge_logic_delivery_rule": resourceSitesEdgeLogicDeliveryRule(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiKey := d.Get("key").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if apiKey != "" {
		c := apiclient.NewClient(apiKey)

		return c, diags
	}

	return nil, diag.Errorf("No key set for key")
}
