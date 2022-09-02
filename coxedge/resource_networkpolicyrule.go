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

	organizationId := d.Get("organization_id").(string)

	//Call the API
	createdNetworkPolicyRule, err := coxEdgeClient.CreateNetworkPolicyRule(newNetworkPolicyRule, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	networkPolicy := make([]map[string]interface{}, len(createdNetworkPolicyRule), len(createdNetworkPolicyRule))
	for i, policy := range createdNetworkPolicyRule {
		item := make(map[string]interface{})
		item["id"] = policy.Id
		item["workload_id"] = policy.WorkloadId
		item["description"] = policy.Description
		item["network_policy_id"] = policy.NetworkPolicyId
		item["type"] = policy.Type
		item["source"] = policy.Source
		item["action"] = policy.Action
		item["protocol"] = policy.Protocol
		item["port_range"] = policy.PortRange
		networkPolicy[i] = item

	}
	d.SetId(newNetworkPolicyRule.NetworkPolicy[0].WorkloadId)
	d.Set("network_policy", networkPolicy)

	return diags
}

func resourceNetworkPolicyRuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the resource Id
	resourceId := d.Id()
	organizationId := d.Get("organization_id").(string)

	//Get the resource
	//networkPolicyRule, err := coxEdgeClient.GetNetworkPolicyRule(d.Get("environment_name").(string), resourceId, organizationId)
	networkPolicyRule, err := coxEdgeClient.GetNetworkPolicyRuleWorkload(d.Get("environment_name").(string), resourceId, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Update state
	convertNetworkPolicyRuleWorkloadAPIObjectToResourceData(d, networkPolicyRule)

	return diags
}

func resourceNetworkPolicyRuleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource Id
	resourceId := d.Id()
	organizationId := d.Get("organization_id").(string)

	//Convert resource data to API object
	updatedNetworkPolicyRule := convertResourceDataToNetworkPolicyRuleCreateAPIObject(d)

	//Call the API
	updatedRule, err := coxEdgeClient.UpdateNetworkPolicyRule(resourceId, updatedNetworkPolicyRule, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Set last_updated
	d.Set("last_updated", time.Now().Format(time.RFC850))
	//Set new id
	d.SetId(updatedRule[0].WorkloadId)

	return resourceNetworkPolicyRuleRead(ctx, d, m)
}

func resourceNetworkPolicyRuleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource Id
	resourceId := d.Id()
	organizationId := d.Get("organization_id").(string)

	updatedNetworkPolicyRule := apiclient.NetworkPolicyRuleCreateRequest{
		EnvironmentName: d.Get("environment_name").(string),
	}
	for _, entry := range d.Get("network_policy").([]interface{}) {
		convertedEntry := entry.(map[string]interface{})
		networkObj := apiclient.NetworkPolicyList{
			EnvironmentName: d.Get("environment_name").(string),
			Id:              convertedEntry["id"].(string),
			WorkloadId:      convertedEntry["workload_id"].(string),
			Description:     convertedEntry["description"].(string),
			Protocol:        convertedEntry["protocol"].(string),
			Type:            convertedEntry["type"].(string),
			Action:          convertedEntry["action"].(string),
			Source:          convertedEntry["source"].(string),
			PortRange:       convertedEntry["port_range"].(string),
		}
		updatedNetworkPolicyRule.NetworkPolicy = append(updatedNetworkPolicyRule.NetworkPolicy, networkObj)
	}

	//Delete the NetworkPolicyRule
	err := coxEdgeClient.DeleteNetworkPolicyRule(d.Get("environment_name").(string), resourceId, organizationId, updatedNetworkPolicyRule)
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
	}
	for _, entry := range d.Get("network_policy").([]interface{}) {
		convertedEntry := entry.(map[string]interface{})
		networkObj := apiclient.NetworkPolicyList{
			EnvironmentName: d.Get("environment_name").(string),
			Id:              convertedEntry["id"].(string),
			WorkloadId:      convertedEntry["workload_id"].(string),
			Description:     convertedEntry["description"].(string),
			Protocol:        convertedEntry["protocol"].(string),
			Type:            convertedEntry["type"].(string),
			Action:          convertedEntry["action"].(string),
			Source:          convertedEntry["source"].(string),
			PortRange:       convertedEntry["port_range"].(string),
		}
		updatedNetworkPolicyRule.NetworkPolicy = append(updatedNetworkPolicyRule.NetworkPolicy, networkObj)
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

func convertNetworkPolicyRuleWorkloadAPIObjectToResourceData(d *schema.ResourceData, networkPolicyRule []apiclient.NetworkPolicyRule) {
	networkPolicy := make([]map[string]interface{}, len(networkPolicyRule), len(networkPolicyRule))
	for i, policy := range networkPolicyRule {
		item := make(map[string]interface{})
		item["id"] = policy.Id
		item["workload_id"] = policy.WorkloadId
		item["description"] = policy.Description
		item["network_policy_id"] = policy.NetworkPolicyId
		item["type"] = policy.Type
		item["source"] = policy.Source
		item["action"] = policy.Action
		item["protocol"] = policy.Protocol
		item["port_range"] = policy.PortRange
		networkPolicy[i] = item

	}
	d.Set("network_policy", networkPolicy)
}
