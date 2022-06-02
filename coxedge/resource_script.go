/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
)

func resourceScript() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceScriptCreate,
		ReadContext:   resourceScriptRead,
		UpdateContext: resourceScriptUpdate,
		DeleteContext: resourceScriptDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getScriptSchema(),
	}
}

func resourceScriptCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	///Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Convert resource data to API Object
	newScript := convertResourceDataToScriptCreateAPIObject(d)

	//Call the API
	createdScript, err := coxEdgeClient.CreateScript(newScript)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Create. Awaiting task result.")

	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithDefaults(ctx, createdScript.TaskId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the ID
	d.SetId(taskResult.Data.Result.Id)

	return diags
}

func resourceScriptRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the resource ID
	resourceId := d.Id()

	//Get the resource
	script, err := coxEdgeClient.GetScript(resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	convertScriptAPIObjectToResourceData(d, script)

	//Update state
	resourceScriptRead(ctx, d, m)

	return diags
}

func resourceScriptUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource ID
	resourceId := d.Id()

	//Convert resource data to API object
	updatedScript := convertResourceDataToScriptCreateAPIObject(d)

	//Call the API
	updateScriptResponse, err := coxEdgeClient.UpdateScript(resourceId, updatedScript)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = coxEdgeClient.AwaitTaskResolveWithDefaults(ctx, updateScriptResponse.TaskId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Set last_updated
	d.Set("last_updated", time.Now().Format(time.RFC850))

	return resourceScriptRead(ctx, d, m)
}

func resourceScriptDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource ID
	resourceId := d.Id()

	//Delete the Script
	err := coxEdgeClient.DeleteScript(resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func convertResourceDataToScriptCreateAPIObject(d *schema.ResourceData) apiclient.ScriptCreateRequest {
	//Create update script struct
	updatedScript := apiclient.ScriptCreateRequest{
		SiteId: d.Get("site_id").(string),
		Name:   d.Get("name").(string),
		Code:   d.Get("code").(string),
	}

	//Convert Backup Origin Codes
	updatedScript.Routes = []string{}
	for _, route := range d.Get("routes").([]interface{}) {
		updatedScript.Routes = append(updatedScript.Routes, route.(string))
	}

	return updatedScript
}

func convertScriptAPIObjectToResourceData(d *schema.ResourceData, script *apiclient.Script) {
	//Store the data
	d.Set("id", script.Id)
	d.Set("stack_id", script.StackId)
	d.Set("site_id", script.SiteId)
	d.Set("name", script.Name)
	d.Set("created_at", script.CreatedAt)
	d.Set("updated_at", script.UpdatedAt)
	d.Set("version", script.Version)
	d.Set("code", script.Code)
	d.Set("routes", script.Routes)
}
