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

func resourceSite() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSiteCreate,
		ReadContext:   resourceSiteRead,
		UpdateContext: resourceSiteUpdate,
		DeleteContext: resourceSiteDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getSiteSchema(),
	}
}

func resourceSiteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	//Convert resource data to API Object
	newSite := convertResourceDataToSiteCreateAPIObject(d)

	organizationId := d.Get("organization_id").(string)
	//Call the API
	createdSite, err := coxEdgeClient.CreateSite(newSite, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithDefaults(ctx, createdSite.TaskId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the Id
	d.SetId(taskResult.Data.Result.Id)

	return diags
}

func resourceSiteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the resource Id
	resourceId := d.Id()
	organizationId := d.Get("organization_id").(string)
	//Get the resource
	site, err := coxEdgeClient.GetSite(d.Get("environment_name").(string), resourceId, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Update state
	convertSiteAPIObjectToResourceData(d, site)

	return diags
}

func resourceSiteUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource Id
	resourceId := d.Id()
	organizationId := d.Get("organization_id").(string)
	value, hasValue := d.GetOk("operation")
	if hasValue {
		//Call the API
		_, err := coxEdgeClient.UpdateSite(resourceId, d.Get("environment_name").(string), value.(string), organizationId)
		if err != nil {
			return diag.FromErr(err)
		}
		//Set last_updated
		d.Set("last_updated", time.Now().Format(time.RFC850))
		return resourceSiteRead(ctx, d, m)
	}
	var diags diag.Diagnostics
	diag := diag.Diagnostic{
		Severity: diag.Error,
		Summary:  " Missing required argument for Update",
		Detail:   "The argument \"operation\" is required for update Site, but no definition was found",
	}
	diags = append(diags, diag)
	return diags
}

func resourceSiteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource Id
	resourceId := d.Id()
	organizationId := d.Get("organization_id").(string)
	//Delete the Site
	err := coxEdgeClient.DeleteSite(d.Get("environment_name").(string), resourceId, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func convertResourceDataToSiteCreateAPIObject(d *schema.ResourceData) apiclient.SiteCreateRequest {
	//Create update site struct
	updatedSite := apiclient.SiteCreateRequest{
		EnvironmentName: d.Get("environment_name").(string),
		Domain:          d.Get("domain").(string),
		Hostname:        d.Get("hostname").(string),
		Protocol:        d.Get("protocol").(string),
		//Optional
		AuthMethod: d.Get("auth_method").(string),
		Username:   d.Get("username").(string),
		Password:   d.Get("password").(string),
	}
	for _, svcName := range d.Get("services").([]interface{}) {
		updatedSite.Services = append(updatedSite.Services, svcName.(string))
	}

	return updatedSite
}

func convertSiteAPIObjectToResourceData(d *schema.ResourceData, site *apiclient.Site) {
	//Store the data
	d.Set("id", site.Id)
	d.Set("stack_id", site.StackId)
	d.Set("domain", site.Domain)
	d.Set("status", site.Status)
	d.Set("created_at", site.CreatedAt)
	d.Set("updated_at", site.UpdatedAt)
	d.Set("services", site.Services)
	d.Set("edge_address", site.EdgeAddress)
	d.Set("anycast_ip", site.AnycastIp)
	deliveryDomains := make([]map[string]string, len(site.DeliveryDomains), len(site.DeliveryDomains))
	for i, delDomain := range site.DeliveryDomains {
		item := make(map[string]string)
		item["domain"] = delDomain.Domain
		item["validated_at"] = delDomain.ValidatedAt
		deliveryDomains[i] = item
	}
	d.Set("delivery_domains", deliveryDomains)
}
