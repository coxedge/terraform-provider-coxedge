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

func resourceComputeFirewallIPv6Rule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeFirewallIPv6RuleCreate,
		ReadContext:   resourceComputeFirewallIPv6RuleRead,
		UpdateContext: resourceComputeFirewallIPv6RuleUpdate,
		DeleteContext: resourceComputeFirewallIPv6RuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getResourceComputeFirewallIPRuleSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceComputeFirewallIPv6RuleCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	firewallId := data.Get("firewall_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ipv6request := convertResourceDataToComputeFirewallIPv6RuleCreateAPIObject(data)

	existingList, err := coxEdgeClient.GetComputeFirewallsIPv6Rules(environmentName, organizationId, firewallId)
	if err != nil {
		return diag.FromErr(err)
	}
	existingIDs := make(map[string]bool)
	for _, item := range existingList {
		existingIDs[item.ID] = true
	}

	//Call the API
	firewallResponse, err := coxEdgeClient.CreateComputeFirewallIPv6Rule(ipv6request, environmentName, organizationId, firewallId)
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
		afterList, err := coxEdgeClient.GetComputeFirewallsIPv6Rules(environmentName, organizationId, firewallId)
		if err != nil {
			return diag.FromErr(err)
		}
		var missingItem *apiclient.ComputeFirewallRule
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

func resourceComputeFirewallIPv6RuleRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//check the id comes with id & environment_name, then split the value -> in case of importing the resource
	//format is <ipv4ruleId>:<firewallId>:<environment_name>:<organization_id>
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
	ipv6RuleId := data.Id()

	computeFirewallIPv6, err := coxEdgeClient.GetComputeFirewallsIPv6RuleById(environmentName, organizationId, firewallId, ipv6RuleId)
	if err != nil {
		return diag.FromErr(err)
	}
	convertFirewallIPv6RuleToResourceData(data, computeFirewallIPv6)
	data.SetId(ipv6RuleId)
	return diags
}

func resourceComputeFirewallIPv6RuleUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceComputeFirewallIPv6RuleDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
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
	err := coxEdgeClient.DeleteComputeFirewallIPv6RuleById(environmentName, organizationId, firewallId, resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")

	return diags
}

func convertResourceDataToComputeFirewallIPv6RuleCreateAPIObject(data *schema.ResourceData) apiclient.ComputeFirewallRuleRequest {
	firewallRequest := apiclient.ComputeFirewallRuleRequest{
		CIDR:         data.Get("cidr").(string),
		Protocol:     data.Get("protocol").(string),
		SourceOption: data.Get("source_option").(string),
		Port:         data.Get("port").(string),
		Notes:        data.Get("notes").(string),
	}
	return firewallRequest
}

func convertFirewallIPv6RuleToResourceData(d *schema.ResourceData, rule *apiclient.ComputeFirewallRule) {
	d.Set("id", rule.ID)
	d.Set("type", rule.Type)
	d.Set("ip_type", rule.IPType)
	d.Set("action", rule.Action)
	d.Set("protocol", rule.Protocol)
	d.Set("port", rule.Port)
	d.Set("subnet", rule.Subnet)
	d.Set("subnet_size", rule.SubnetSize)
	d.Set("source", rule.Source)
	d.Set("notes", rule.Notes)
}