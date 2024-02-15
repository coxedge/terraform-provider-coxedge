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
			"coxedge_organizations":              dataSourceOrganization(),
			"coxedge_organizations_billing_info": dataSourceOrganizationBillingInfo(),
			"coxedge_environments":               dataSourceEnvironment(),
			"coxedge_images":                     dataSourceImage(),
			//"coxedge_origin_settings":            dataSourceOriginSetting(),
			"coxedge_roles":                             dataSourceRoles(),
			"coxedge_workload_instances":                dataWorkloadInstances(),
			"coxedge_vpcs":                              dataSourceVPCs(),
			"coxedge_subnets":                           dataSourceSubnets(),
			"coxedge_routes":                            dataSourceRoutes(),
			"coxedge_baremetal_devices":                 dataSourceBareMetalDevice(),
			"coxedge_baremetal_device_charts":           dataSourceBareMetalDeviceCharts(),
			"coxedge_baremetal_device_sensors":          dataSourceBareMetalDeviceSensors(),
			"coxedge_baremetal_device_ips":              dataSourceBareMetalDeviceIPs(),
			"coxedge_baremetal_ssh_keys":                dataSourceBareMetalSSHKeys(),
			"coxedge_baremetal_device_disk":             dataSourceBareMetalDeviceDisks(),
			"coxedge_baremetal_locations":               dataSourceBareMetalLocations(),
			"coxedge_baremetal_location_products":       dataSourceBareMetalLocationProducts(),
			"coxedge_baremetal_location_product_os":     dataSourceBareMetalLocationProductOS(),
			"coxedge_compute_workloads":                 dataSourceComputeWorkloads(),
			"coxedge_compute_workload_ipv4":             dataSourceComputeWorkloadIPv4(),
			"coxedge_compute_workload_ipv6":             dataSourceComputeWorkloadIPv6(),
			"coxedge_compute_workload_ipv6_reverse_dns": dataSourceComputeWorkloadIPv6ReverseDNS(),
			"coxedge_compute_workload_firewall_group":   dataSourceComputeWorkloadFirewallGroup(),
			"coxedge_compute_workload_hostname":         dataSourceComputeWorkloadHostname(),
			"coxedge_compute_workload_plan":             dataSourceComputeWorkloadPlan(),
			"coxedge_compute_workload_os":               dataSourceComputeWorkloadOS(),
			"coxedge_compute_workload_user_data":        dataSourceComputeWorkloadUserData(),
			"coxedge_compute_workload_tags":             dataSourceComputeWorkloadTags(),
			"coxedge_compute_storages":                  dataSourceComputeStorages(),
			"coxedge_compute_firewalls":                 dataSourceComputeFirewalls(),
			"coxedge_compute_firewall_ipv4_rule":        dataSourceComputeFirewallIPv4Rules(),
			"coxedge_compute_firewall_ipv6_rule":        dataSourceComputeFirewallIPv6Rules(),
			"coxedge_compute_firewall_linked_instances": dataSourceComputeFirewallLinkedInstances(),
			"coxedge_compute_vpc2":                      dataSourceComputeVPC2(),
			"coxedge_compute_vpc":                       dataSourceComputeVPC(),
			"coxedge_compute_reserved_ips":              dataSourceComputeReservedIP(),
		},
		ResourcesMap: map[string]*schema.Resource{
			//"coxedge_cdn_purge":           resourceCDNPurgeResource(),
			//"coxedge_cdn_settings":        resourceCDNSettings(),
			//"coxedge_delivery_domain":     resourceDeliveryDomain(),
			"coxedge_environment": resourceEnvironment(),
			//"coxedge_firewall_rule":       resourceFirewallRule(),
			"coxedge_network_policy_rule": resourceNetworkPolicyRule(),
			//"coxedge_origin_setting":      resourceOriginSettings(),
			//"coxedge_script":              resourceScript(),
			//"coxedge_site":                resourceSite(),
			"coxedge_user": resourceUser(),
			//"coxedge_waf_settings":        resourceWAFSettings(),
			"coxedge_workload":                               resourceWorkload(),
			"coxedge_vpc":                                    resourceVPC(),
			"coxedge_subnet":                                 resourceSubnet(),
			"coxedge_route":                                  resourceRoute(),
			"coxedge_baremetal_devices":                      resourceBareMetalDevices(),
			"coxedge_baremetal_device":                       resourceBareMetalDevice(),
			"coxedge_baremetal_device_ipmi":                  resourceBareMetalDeviceIPMI(),
			"coxedge_baremetal_ssh_key":                      resourceBareMetaSSHKey(),
			"coxedge_compute_workload":                       resourceComputeWorkload(),
			"coxedge_compute_workload_ipv6_reverse_dns":      resourceComputeWorkloadIPv6ReverseDNS(),
			"coxedge_compute_workload_firewall_group":        resourceComputeWorkloadFirewallGroup(),
			"coxedge_compute_workload_hostname":              resourceComputeWorkloadHostname(),
			"coxedge_compute_workload_plan":                  resourceComputeWorkloadPlan(),
			"coxedge_compute_workload_os":                    resourceComputeWorkloadOS(),
			"coxedge_compute_workload_user_data":             resourceComputeWorkloadUserData(),
			"coxedge_compute_workload_tags":                  resourceComputeWorkloadTags(),
			"coxedge_compute_workload_operation":             resourceComputeWorkloadOperation(),
			"coxedge_compute_storage":                        resourceComputeStorage(),
			"coxedge_compute_storage_attach_detach_instance": resourceComputeStorageAttachDetachInstance(),
			"coxedge_compute_firewalls":                      resourceComputeFirewall(),
			"coxedge_compute_firewall_ipv4_rule":             resourceComputeFirewallIPv4Rule(),
			"coxedge_compute_firewall_ipv6_rule":             resourceComputeFirewallIPv6Rule(),
			"coxedge_compute_firewall_linked_instances":      resourceComputeFirewallLinkedInstance(),
			"coxedge_compute_vpc2":                           resourceComputeVPC2(),
			"coxedge_compute_vpc":                            resourceComputeVPC(),
			"coxedge_compute_reserved_ips":                   resourceComputeReservedIP(),
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
