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
	"strings"
	"time"
)

func resourceSitesPredefinedEdgeLogic() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePredefinedEdgeLogicCreate,
		ReadContext:   resourcePredefinedEdgeLogicRead,
		UpdateContext: resourcePredefinedEdgeLogicUpdate,
		DeleteContext: resourcePredefinedEdgeLogicDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getPredefinedEdgeLogicResourceSchema(),
		Timeouts: &schema.ResourceTimeout{
			Update: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourcePredefinedEdgeLogicCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Errorf("Users cannot create predefined edge logic, but can only modify its fields after importing it. More details can be found in the documentation.")
}

func resourcePredefinedEdgeLogicRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//check the id comes with id,environment_name & organization_id, then split the value -> in case of importing the resource
	//format is <site_id>:<environment_name>:<organization_id>
	if strings.Contains(d.Id(), ":") {
		keys := strings.Split(d.Id(), ":")
		d.Set("site_id", keys[0])
		d.Set("environment_name", keys[1])
		d.Set("organization_id", keys[2])
	}
	//Get the resource Id
	resourceId := d.Get("site_id").(string)
	organizationId := d.Get("organization_id").(string)
	environmentName := d.Get("environment_name").(string)

	edgeLogic, err := coxEdgeClient.GetPredefinedEdgeLogics(environmentName, organizationId, resourceId)
	if err != nil {
		return diag.FromErr(err)
	}
	convertPredefinedEdgeLogicAPIObjectToResourceData(d, edgeLogic)

	return diags
}

func resourcePredefinedEdgeLogicUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	resourceId := d.Get("site_id").(string)
	environmentName := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)

	requestBody := convertResourceDataToPredefinedEdgeLogicAPIObject(d)

	edgeLogic, err := coxEdgeClient.UpdatePredefinedEdgeLogic(requestBody, environmentName, organizationId, resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Update. Awaiting task result.")

	timeout := d.Timeout(schema.TimeoutUpdate)
	//Await
	_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, edgeLogic.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	////Save the Id
	//d.SetId(taskResult.Data.Result.Id)

	return diags
}

func resourcePredefinedEdgeLogicDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Errorf("Users cannot delete predefined edge logic, but can only modify its fields after importing it. More details can be found in the documentation.")
}

func convertPredefinedEdgeLogicAPIObjectToResourceData(d *schema.ResourceData, edgeLogic *apiclient.EdgeLogic) {
	//Store the data
	d.Set("id", edgeLogic.Id)
	d.Set("stack_id", edgeLogic.StackId)
	d.Set("scope_id", edgeLogic.ScopeId)
	d.Set("force_www_enabled", edgeLogic.ForceWwwEnabled)
	d.Set("robots_txt_enabled", edgeLogic.RobotsTxtEnabled)
	d.Set("robots_txt_file", edgeLogic.RobotTxtFile)
	d.Set("pseudo_streaming_enabled", edgeLogic.PseudoStreamingEnabled)
	d.Set("referrer_protection_enabled", edgeLogic.ReferrerProtectionEnabled)
	d.Set("allow_empty_referrer", edgeLogic.AllowEmptyReferrer)
	d.Set("referrer_list", edgeLogic.ReferrerList)
}

func convertResourceDataToPredefinedEdgeLogicAPIObject(d *schema.ResourceData) apiclient.PredefinedEdgeLogicRequest {
	forceWwwEnabled := d.Get("force_www_enabled").(bool)
	robotsTxtEnabled := d.Get("robots_txt_enabled").(bool)
	pseudoStreamingEnabled := d.Get("pseudo_streaming_enabled").(bool)
	referrerProtectionEnabled := d.Get("referrer_protection_enabled").(bool)
	allowEmptyReferrer := d.Get("allow_empty_referrer").(bool)

	referrerList := d.Get("referrer_list").([]interface{})
	refList := make([]string, len(referrerList))
	for i, v := range referrerList {
		refList[i] = v.(string)
	}
	predefinedEdgeLogicObject := apiclient.PredefinedEdgeLogicRequest{
		ForceWwwEnabled:           &forceWwwEnabled,
		RobotsTxtEnabled:          &robotsTxtEnabled,
		RobotsTxtFile:             d.Get("robots_txt_file").(string),
		PseudoStreamingEnabled:    &pseudoStreamingEnabled,
		ReferrerProtectionEnabled: &referrerProtectionEnabled,
		AllowEmptyReferrer:        &allowEmptyReferrer,
		ReferrerList:              refList,
	}

	return predefinedEdgeLogicObject
}
