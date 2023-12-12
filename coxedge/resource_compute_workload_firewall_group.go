package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"time"
)

func resourceComputeWorkloadFirewallGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeWorkloadFirewallGroupCreate,
		ReadContext:   resourceComputeWorkloadFirewallGroupRead,
		UpdateContext: resourceComputeWorkloadFirewallGroupUpdate,
		DeleteContext: resourceComputeWorkloadFirewallGroupDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getResourceComputeWorkloadFirewallGroupSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceComputeWorkloadFirewallGroupCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceComputeWorkloadFirewallGroupUpdate(ctx, data, i)
	return diags
}

func resourceComputeWorkloadFirewallGroupRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceComputeWorkloadFirewallGroupUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//convert resource data to API object
	firewallGroupRequest := convertResourceDataToComputeWorkloadFirewallGroupCreateAPIObject(data)

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	workloadId := data.Get("workload_id").(string)

	//Call the API
	createdReverseDNS, err := coxEdgeClient.UpdateComputeWorkloadFirewallGroupById(firewallGroupRequest, environmentName, organizationId, workloadId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Update. Awaiting task result.")

	timeout := data.Timeout(schema.TimeoutCreate)
	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, createdReverseDNS.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the Id
	data.SetId(taskResult.Data.TaskId)
	return diags
}

func resourceComputeWorkloadFirewallGroupDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func convertResourceDataToComputeWorkloadFirewallGroupCreateAPIObject(data *schema.ResourceData) apiclient.ComputeWorkloadFirewallGroupRequest {
	firewallGroupRequest := apiclient.ComputeWorkloadFirewallGroupRequest{
		FirewallId: data.Get("firewall_id").(string),
	}
	return firewallGroupRequest
}
