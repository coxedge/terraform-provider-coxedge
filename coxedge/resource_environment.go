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

func resourceEnvironment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEnvironmentCreate,
		ReadContext:   resourceEnvironmentRead,
		UpdateContext: resourceEnvironmentUpdate,
		DeleteContext: resourceEnvironmentDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getEnvironmentSchema(),
	}
}

func resourceEnvironmentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Convert resource data to API Object
	newEnvironment := convertResourceDataToEnvironmentCreateAPIObject(d)

	//Call the API
	createdEnvironment, err := coxEdgeClient.CreateEnvironment(newEnvironment)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the ID
	d.SetId(createdEnvironment.Id)

	return diags
}

func resourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the resource ID
	resourceId := d.Id()

	//Get the resource
	environment, err := coxEdgeClient.GetEnvironment(resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	convertEnvironmentAPIObjectToResourceData(d, environment)

	//Update state
	resourceEnvironmentRead(ctx, d, m)

	return diags
}

func resourceEnvironmentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource ID
	resourceId := d.Id()

	//Convert resource data to API object
	updatedEnvironment := convertResourceDataToEnvironmentCreateAPIObject(d)

	//Call the API
	_, err := coxEdgeClient.UpdateEnvironment(resourceId, updatedEnvironment)
	if err != nil {
		return diag.FromErr(err)
	}

	//Set last_updated
	d.Set("last_updated", time.Now().Format(time.RFC850))

	return resourceEnvironmentRead(ctx, d, m)
}

func resourceEnvironmentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource ID
	resourceId := d.Id()

	//Delete the Environment
	err := coxEdgeClient.DeleteEnvironment(resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func convertResourceDataToEnvironmentCreateAPIObject(d *schema.ResourceData) apiclient.EnvironmentCreateRequest {
	//Create update environment struct
	updatedEnvironment := apiclient.EnvironmentCreateRequest{
		EnvironmentName: d.Get("name").(string),
		Description:     d.Get("description").(string),
		Organization:    apiclient.IdOnlyHelper{Id: d.Get("organization_id").(string)},
		ServiceConnection: apiclient.IdOnlyHelper{
			Id: d.Get("service_connection_id").(string),
		},
	}

	return updatedEnvironment
}

func convertEnvironmentAPIObjectToResourceData(d *schema.ResourceData, environment *apiclient.Environment) {
	//Store the data
	d.Set("id", environment.Id)
	d.Set("name", environment.Name)
	d.Set("description", environment.Description)
	d.Set("organization_id", environment.Organization.Id)
	d.Set("service_connection_id", environment.ServiceConnection.Id)
	d.Set("creation_date", environment.CreationDate)
}
