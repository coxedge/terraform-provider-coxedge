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

func resourceComputeFirewall() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeFirewallCreate,
		ReadContext:   resourceComputeFirewallRead,
		UpdateContext: resourceComputeFirewallUpdate,
		DeleteContext: resourceComputeFirewallDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getResourceComputeFirewallSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceComputeFirewallCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//convert resource data to API object
	firewallRequest := convertResourceDataToComputeFirewallCreateAPIObject(data)
	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)

	existingList, err := coxEdgeClient.GetComputeFirewalls(environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}
	existingIDs := make(map[string]bool)
	for _, item := range existingList {
		existingIDs[item.ID] = true
	}
	//Call the API
	firewallResponse, err := coxEdgeClient.CreateComputeFirewall(firewallRequest, environmentName, organizationId)
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
		afterList, err := coxEdgeClient.GetComputeFirewalls(environmentName, organizationId)
		if err != nil {
			return diag.FromErr(err)
		}
		var missingItem *apiclient.ComputeFirewall
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

func resourceComputeFirewallRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//check the id comes with id & environment_name, then split the value -> in case of importing the resource
	//format is <firewallId>:<environment_name>:<organization_id>
	if strings.Contains(data.Id(), ":") {
		keys := strings.Split(data.Id(), ":")
		data.SetId(keys[0])
		data.Set("environment_name", keys[1])
		data.Set("organization_id", keys[2])
	}

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	firewallId := data.Id()

	computeFirewall, err := coxEdgeClient.GetComputeFirewallById(environmentName, organizationId, firewallId)
	if err != nil {
		return diag.FromErr(err)
	}
	convertComputeFirewallToResourceData(data, computeFirewall)
	data.SetId(firewallId)
	return diags
}

func resourceComputeFirewallUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	firewallId := data.Id()

	updateRequest := apiclient.UpdateComputeFirewallRequest{
		Id:          firewallId,
		Description: data.Get("description").(string),
	}
	//Call the API
	updateFirewall, err := coxEdgeClient.UpdateComputeFirewall(updateRequest, environmentName, organizationId, firewallId)
	if err != nil {
		return diag.FromErr(err)
	}
	timeout := data.Timeout(schema.TimeoutUpdate)
	tflog.Info(ctx, "Initiated Update. Awaiting task result.")
	//Await
	_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, updateFirewall.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(firewallId)
	return diags
}

func resourceComputeFirewallDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	//Get the resource Id
	resourceId := data.Id()
	organizationId := data.Get("organization_id").(string)
	environmentName := data.Get("environment_name").(string)

	//Delete the Storage
	err := coxEdgeClient.DeleteComputeFirewallById(environmentName, organizationId, resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")

	return diags
}

func convertResourceDataToComputeFirewallCreateAPIObject(data *schema.ResourceData) apiclient.ComputeFirewallRequest {
	firewallRequest := apiclient.ComputeFirewallRequest{
		Description: data.Get("description").(string),
	}
	return firewallRequest
}

func convertComputeFirewallToResourceData(d *schema.ResourceData, firewall *apiclient.ComputeFirewall) {
	d.Set("id", firewall.ID)
	d.Set("description", firewall.Description)
	d.Set("date_created", firewall.DateCreated)
	d.Set("date_modified", firewall.DateModified)
	d.Set("instance_count", firewall.InstanceCount)
	d.Set("rule_count", firewall.RuleCount)
	d.Set("max_rule_count", firewall.MaxRuleCount)
}
