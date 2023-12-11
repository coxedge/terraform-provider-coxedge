package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"time"
)

func resourceComputeWorkload() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeWorkloadCreate,
		ReadContext:   resourceComputeWorkloadRead,
		UpdateContext: resourceComputeWorkloadUpdate,
		DeleteContext: resourceComputeWorkloadDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getResourceComputeWorkloadSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceComputeWorkloadRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceComputeWorkloadCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//convert resource data to API object
	workloadRequest := convertResourceDataToComputeWorkloadCreateAPIObject(data)

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)

	//Call the API
	createdWorkload, err := coxEdgeClient.CreateComputeWorkload(workloadRequest, environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Create. Awaiting task result.")

	timeout := data.Timeout(schema.TimeoutCreate)
	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, createdWorkload.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the Id
	data.SetId(taskResult.Data.TaskId)

	return diags
}

func resourceComputeWorkloadUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceComputeWorkloadDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func convertResourceDataToComputeWorkloadCreateAPIObject(data *schema.ResourceData) apiclient.ComputeWorkloadRequest {

	workloadRequest := apiclient.ComputeWorkloadRequest{
		IsIPv6:                 data.Get("is_ipv6").(bool),
		NoPublicIPv4:           data.Get("no_public_ipv4").(bool),
		IsVirtualPrivateClouds: data.Get("is_virtual_private_clouds").(bool),
		IsVPC2:                 data.Get("is_vpc2").(bool),
		OperatingSystemId:      data.Get("operating_system_id").(string),
		LocationId:             data.Get("location_id").(string),
		PlanId:                 data.Get("plan_id").(string),
		Hostname:               data.Get("hostname").(string),
		Label:                  data.Get("label").(string),
		Name:                   data.Get("label").(string),
		FirstBootSshKey:        data.Get("first_boot_ssh_key").(string),
		SshKeyName:             data.Get("ssh_key_name").(string),
		FirewallId:             data.Get("firewall_id").(string),
		UserData:               data.Get("user_data").(string),
	}

	return workloadRequest
}
