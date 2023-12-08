package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

func dataSourceComputeWorkloads() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeWorkloadsRead,
		Schema:      getComputeWorkloadSetSchema(),
	}
}

func dataSourceComputeWorkloadsRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)

	computeWorkloads, err := coxEdgeClient.GetComputeWorkloads(environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("workloads", flattenComputeWorkloadsData(computeWorkloads)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenComputeWorkloadsData(computeWorkloads []apiclient.ComputeWorkload) []interface{} {
	if computeWorkloads != nil {
		workloads := make([]interface{}, len(computeWorkloads), len(computeWorkloads))

		for i, workload := range computeWorkloads {
			item := make(map[string]interface{})

			item["id"] = workload.Id
			item["hostname"] = workload.Hostname
			item["label"] = workload.Label
			item["status"] = workload.Status
			item["os"] = workload.OS
			item["ram"] = workload.RAM
			item["date_created"] = workload.DateCreated
			item["region"] = workload.Region
			item["disk"] = workload.Disk
			item["main_ip"] = workload.MainIP
			item["vcpu_count"] = workload.VCPUCount
			item["plan"] = workload.Plan
			item["allowed_bandwidth"] = workload.AllowedBandwidth
			item["netmask_v4"] = workload.NetmaskV4
			item["gateway_v4"] = workload.GatewayV4
			item["power_status"] = workload.PowerStatus
			item["server_status"] = workload.ServerStatus
			item["v6_network"] = workload.V6Network
			item["v6_main_ip"] = workload.V6MainIP
			item["v6_network_size"] = workload.V6NetworkSize
			item["internal_ip"] = workload.InternalIP
			item["kvm"] = workload.KVM
			item["os_id"] = workload.OSID
			item["app_id"] = workload.AppID
			item["image_id"] = workload.ImageID
			item["firewall_group_id"] = workload.FirewallGroupID
			item["features"] = workload.Features
			item["tags"] = workload.Tags

			workloads[i] = item
		}

		return workloads
	}
	return make([]interface{}, 0)
}
