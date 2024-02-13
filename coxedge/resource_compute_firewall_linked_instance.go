package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strings"
	"time"
)

func resourceComputeFirewallLinkedInstance() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeFirewallLinkedInstanceCreate,
		ReadContext:   resourceComputeFirewallLinkedInstanceRead,
		UpdateContext: resourceComputeFirewallLinkedInstanceUpdate,
		DeleteContext: resourceComputeFirewallLinkedInstanceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getResourceComputeFirewallLinkedInstanceSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceComputeFirewallLinkedInstanceCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	firewallId := data.Get("firewall_id").(string)
	workloadId := data.Get("workload_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	linkedInstanceRequest := apiclient.ComputeFirewallLinkedInstanceRequest{
		Id:         firewallId,
		WorkloadId: workloadId,
	}

	existingList, err := coxEdgeClient.GetComputeFirewallLinkedInstances(environmentName, organizationId, firewallId)
	if err != nil {
		return diag.FromErr(err)
	}
	existingIDs := make(map[string]bool)
	for _, item := range existingList {
		existingIDs[item.ID] = true
	}

	//Call the API
	firewallResponse, err := coxEdgeClient.CreateComputeFirewallLinkedInstance(linkedInstanceRequest, environmentName, organizationId, firewallId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Create. Awaiting task result.")

	timeout := data.Timeout(schema.TimeoutCreate)
	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, firewallResponse.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	if taskResult.Data.TaskStatus == "SUCCESS" {
		afterList, err := coxEdgeClient.GetComputeFirewallLinkedInstances(environmentName, organizationId, firewallId)
		if err != nil {
			return diag.FromErr(err)
		}
		var missingItem *apiclient.ComputeFirewallLinkedInstance
		for _, item := range afterList {
			if !existingIDs[item.ID] {
				missingItem = &item
				data.SetId(missingItem.ID)
				return diags
			}
		}
	}
	//Save the Id
	data.SetId(taskResult.Data.TaskId)
	return diags
}

func resourceComputeFirewallLinkedInstanceRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//check the id comes with id & environment_name, then split the value -> in case of importing the resource
	//format is <linkedInstanceId>:<firewallId>:<environment_name>:<organization_id>
	if strings.Contains(data.Id(), ":") {
		keys := strings.Split(data.Id(), ":")
		data.SetId(keys[0])
		data.Set("firewall_id", keys[1])
		data.Set("environment_name", keys[2])
		data.Set("organization_id", keys[3])
	}

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	firewallId := data.Get("firewall_id").(string)
	linkedInstanceId := data.Id()

	computeFirewall, err := coxEdgeClient.GetComputeFirewallLinkedInstanceById(environmentName, organizationId, firewallId, linkedInstanceId)
	if err != nil {
		return diag.FromErr(err)
	}
	data.SetId(computeFirewall.ID)
	return diags
}

func resourceComputeFirewallLinkedInstanceUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceComputeFirewallLinkedInstanceDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	//Get the resource Id
	resourceId := data.Id()
	organizationId := data.Get("organization_id").(string)
	environmentName := data.Get("environment_name").(string)
	firewallId := data.Get("firewall_id").(string)

	//Delete the Storage
	err := coxEdgeClient.DeleteComputeFirewallLinkedInstanceById(environmentName, organizationId, firewallId, resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")

	return diags
}
