package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceComputeFirewallLinkedInstances() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeFirewallLinkedInstancesRead,
		Schema:      getComputeFirewallLinkedInstancesSetSchema(),
	}
}

func dataSourceComputeFirewallLinkedInstancesRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	firewallId := data.Get("firewall_id").(string)
	linkedInstanceId := data.Get("linked_instance_id").(string)

	if linkedInstanceId != "" {
		computeFirewall, err := coxEdgeClient.GetComputeFirewallLinkedInstanceById(environmentName, organizationId, firewallId, linkedInstanceId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("linked_instances", flattenComputeFirewallLinkedInstanceData(&[]apiclient.ComputeFirewallLinkedInstance{*computeFirewall})); err != nil {
			return diag.FromErr(err)
		}
	} else {
		computeFirewalls, err := coxEdgeClient.GetComputeFirewallLinkedInstances(environmentName, organizationId, firewallId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("linked_instances", flattenComputeFirewallLinkedInstanceData(&computeFirewalls)); err != nil {
			return diag.FromErr(err)
		}
	}
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenComputeFirewallLinkedInstanceData(computeFirewalls *[]apiclient.ComputeFirewallLinkedInstance) []interface{} {
	if computeFirewalls != nil {
		firewalls := make([]interface{}, len(*computeFirewalls))

		for i, instance := range *computeFirewalls {
			item := make(map[string]interface{})

			item["id"] = instance.ID
			item["hostname"] = instance.Hostname
			item["label"] = instance.Label
			item["status"] = instance.Status
			item["os"] = instance.OS
			item["ram"] = instance.RAM
			item["date_created"] = instance.DateCreated
			item["region"] = instance.Region
			item["disk"] = instance.Disk
			item["main_ip"] = instance.MainIP
			item["vcpu_count"] = instance.VCPUCount
			item["plan"] = instance.Plan
			item["allowed_bandwidth"] = instance.AllowedBandwidth
			item["netmask_v4"] = instance.NetmaskV4
			item["gateway_v4"] = instance.GatewayV4
			item["power_status"] = instance.PowerStatus
			item["server_status"] = instance.ServerStatus
			item["v6_network"] = instance.V6Network
			item["v6_main_ip"] = instance.V6MainIP
			item["v6_network_size"] = instance.V6NetworkSize
			item["internal_ip"] = instance.InternalIP
			item["kvm"] = instance.KVM
			item["os_id"] = instance.OSID
			item["app_id"] = instance.AppID
			item["image_id"] = instance.ImageID
			item["firewall_group_id"] = instance.FirewallGroupID
			item["features"] = instance.Features
			item["tags"] = instance.Tags

			firewalls[i] = item
		}
		return firewalls
	}
	return make([]interface{}, 0)
}
