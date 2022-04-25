package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
)

func resourceNetworkPolicyRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNetworkPolicyRuleCreate,
		ReadContext:   resourceNetworkPolicyRuleRead,
		UpdateContext: resourceNetworkPolicyRuleUpdate,
		DeleteContext: resourceNetworkPolicyRuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getNetworkPolicyRuleSchema(),
	}
}

func resourceNetworkPolicyRuleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Convert resource data to API Object
	newNetworkPolicyRule := convertResourceDataToNetworkPolicyRuleCreateAPIObject(d)

	//Call the API
	createdNetworkPolicyRule, err := coxEdgeClient.CreateNetworkPolicyRule(newNetworkPolicyRule)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the ID
	d.SetId(createdNetworkPolicyRule.Id)

	return diags
}

func resourceNetworkPolicyRuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the resource ID
	resourceId := d.Id()

	//Get the resource
	networkPolicyRule, err := coxEdgeClient.GetNetworkPolicyRule(d.Get("environment_name").(string), resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Update state
	convertNetworkPolicyRuleAPIObjectToResourceData(d, networkPolicyRule)

	return diags
}

func resourceNetworkPolicyRuleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource ID
	resourceId := d.Id()

	//Convert resource data to API object
	updatedNetworkPolicyRule := convertResourceDataToNetworkPolicyRuleCreateAPIObject(d)

	//Call the API
	updatedRule, err := coxEdgeClient.UpdateNetworkPolicyRule(resourceId, updatedNetworkPolicyRule)
	if err != nil {
		return diag.FromErr(err)
	}

	//Set last_updated
	d.Set("last_updated", time.Now().Format(time.RFC850))
	//Set new id
	d.SetId(updatedRule.Id)

	return resourceNetworkPolicyRuleRead(ctx, d, m)
}

func resourceNetworkPolicyRuleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource ID
	resourceId := d.Id()

	//Delete the NetworkPolicyRule
	err := coxEdgeClient.DeleteNetworkPolicyRule(d.Get("environment_name").(string), resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func convertResourceDataToNetworkPolicyRuleCreateAPIObject(d *schema.ResourceData) apiclient.NetworkPolicyRuleCreateRequest {
	//Create update networkPolicyRule struct
	updatedNetworkPolicyRule := apiclient.NetworkPolicyRuleCreateRequest{
		EnvironmentName: d.Get("environment_name").(string),
		WorkloadId:      d.Get("workload_id").(string),
		Description:     d.Get("description").(string),
		Protocol:        d.Get("protocol").(string),
		Type:            d.Get("type").(string),
		Action:          d.Get("action").(string),
		Source:          d.Get("source").(string),
		PortRange:       d.Get("port_range").(string),
	}

	return updatedNetworkPolicyRule
}

func convertNetworkPolicyRuleAPIObjectToResourceData(d *schema.ResourceData, networkPolicyRule *apiclient.NetworkPolicyRule) {
	//Store the data
	d.Set("id", networkPolicyRule.Id)
	d.Set("workload_id", networkPolicyRule.WorkloadId)
	d.Set("stack_id", networkPolicyRule.StackId)
	d.Set("description", networkPolicyRule.Description)
	d.Set("protocol", networkPolicyRule.Protocol)
	d.Set("type", networkPolicyRule.Type)
	d.Set("action", networkPolicyRule.Action)
	d.Set("source", networkPolicyRule.Source)
	d.Set("port_range", networkPolicyRule.PortRange)
}
