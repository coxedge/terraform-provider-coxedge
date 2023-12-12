package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceComputeWorkloadHostname() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeWorkloadHostnameRead,
		Schema:      getComputeWorkloadHostnameSetSchema(),
	}
}

func dataSourceComputeWorkloadHostnameRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	workloadId := data.Get("workload_id").(string)

	hostname, err := coxEdgeClient.GetComputeWorkloadHostnameById(environmentName, organizationId, workloadId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("hostnames", flattenComputeWorkloadHostnameData(&[]apiclient.Hostname{*hostname})); err != nil {
		return diag.FromErr(err)
	}

	// always run
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenComputeWorkloadHostnameData(hostname *[]apiclient.Hostname) []interface{} {
	if hostname != nil {
		hostnameList := make([]interface{}, len(*hostname), len(*hostname))

		for i, host := range *hostname {
			item := make(map[string]interface{})
			item["id"] = host.Id
			item["hostname"] = host.Hostname
			hostnameList[i] = item
		}
		return hostnameList
	}
	return make([]interface{}, 0)
}
