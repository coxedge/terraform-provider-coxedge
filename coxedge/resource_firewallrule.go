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

func resourceFirewallRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFirewallRuleCreate,
		ReadContext:   resourceFirewallRuleRead,
		UpdateContext: resourceFirewallRuleUpdate,
		DeleteContext: resourceFirewallRuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getFirewallRuleSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceFirewallRuleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Convert resource data to API Object
	newFirewallRule := convertResourceDataToFirewallRuleCreateAPIObject(d)

	organizationId := d.Get("organization_id").(string)
	//Call the API
	createdFirewallRule, err := coxEdgeClient.CreateFirewallRule(d.Get("environment_name").(string), newFirewallRule, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	timeout := d.Timeout(schema.TimeoutCreate)
	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, createdFirewallRule.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the Id
	d.SetId(taskResult.Data.Result.Id)

	return diags
}

func resourceFirewallRuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the resource Id
	resourceId := d.Id()
	organizationId := d.Get("organization_id").(string)
	//Get the resource
	firewallRule, err := coxEdgeClient.GetFirewallRule(d.Get("environment_name").(string), d.Get("site_id").(string), resourceId, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	convertFirewallRuleAPIObjectToResourceData(d, firewallRule)

	return diags
}

func resourceFirewallRuleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource Id
	resourceId := d.Id()

	//Convert resource data to API object
	updatedFirewallRule := convertResourceDataToFirewallRuleCreateAPIObject(d)
	organizationId := d.Get("organization_id").(string)
	//Call the API
	_, err := coxEdgeClient.UpdateFirewallRule(d.Get("environment_name").(string), resourceId, updatedFirewallRule, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Set last_updated
	d.Set("last_updated", time.Now().Format(time.RFC850))

	return resourceFirewallRuleRead(ctx, d, m)
}

func resourceFirewallRuleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource Id
	resourceId := d.Id()
	organizationId := d.Get("organization_id").(string)
	//Delete the FirewallRule
	err := coxEdgeClient.DeleteFirewallRule(d.Get("environment_name").(string), d.Get("site_id").(string), resourceId, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func convertResourceDataToFirewallRuleCreateAPIObject(d *schema.ResourceData) apiclient.FirewallRule {
	//Create update firewallRule struct
	updatedFirewallRule := apiclient.FirewallRule{
		Action:  d.Get("action").(string),
		Enabled: d.Get("enabled").(bool),
		Id:      d.Get("id").(string),
		IpEnd:   d.Get("ip_end").(string),
		IpStart: d.Get("ip_start").(string),
		Name:    d.Get("name").(string),
		SiteId:  d.Get("site_id").(string),
	}

	return updatedFirewallRule
}

func convertFirewallRuleAPIObjectToResourceData(d *schema.ResourceData, firewallRule *apiclient.FirewallRule) {
	//Store the data
	d.Set("id", firewallRule.Id)
	d.Set("site_id", firewallRule.SiteId)
	d.Set("action", firewallRule.Action)
	d.Set("ip_start", firewallRule.IpStart)
	d.Set("name", firewallRule.Name)
	d.Set("enabled", firewallRule.Enabled)
	d.Set("ip_end", firewallRule.IpEnd)
}
