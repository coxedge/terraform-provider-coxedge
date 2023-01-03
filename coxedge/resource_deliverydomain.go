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

func resourceDeliveryDomain() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDeliveryDomainCreate,
		ReadContext:   resourceDeliveryDomainRead,
		UpdateContext: resourceDeliveryDomainUpdate,
		DeleteContext: resourceDeliveryDomainDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getDeliveryDomainSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceDeliveryDomainCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Convert resource data to API Object
	newDeliveryDomain := convertResourceDataToDeliveryDomainCreateAPIObject(d)

	resourceId := d.Get("site_id").(string)
	organizationId := d.Get("organization_id").(string)
	//Call the API
	createdDeliveryDomain, err := coxEdgeClient.CreateDeliveryDomain(resourceId, newDeliveryDomain, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	timeout := d.Timeout(schema.TimeoutCreate)
	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, createdDeliveryDomain.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the Id
	d.SetId(taskResult.Data.Result.Id)

	return diags
}

func resourceDeliveryDomainRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the resource Id
	resourceId := d.Id()
	organizationId := d.Get("organization_id").(string)

	//Get the resource
	deliveryDomain, err := coxEdgeClient.GetDeliveryDomain(d.Get("environment_name").(string), resourceId, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	convertDeliveryDomainAPIObjectToResourceData(d, deliveryDomain)

	return diags
}

func resourceDeliveryDomainUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Errorf("no option to update delivery domain - remove terraform state files to create new one")
}

func resourceDeliveryDomainDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource Id
	resourceId := d.Id()
	organizationId := d.Get("organization_id").(string)

	//Delete the DeliveryDomain
	err := coxEdgeClient.DeleteDeliveryDomain(d.Get("environment_name").(string), resourceId, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func convertResourceDataToDeliveryDomainCreateAPIObject(d *schema.ResourceData) apiclient.DeliveryDomainCreateRequest {
	//Create update deliveryDomain struct
	updatedDeliveryDomain := apiclient.DeliveryDomainCreateRequest{
		EnvironmentName: d.Get("environment_name").(string),
		Domain:          d.Get("domain").(string),
	}

	return updatedDeliveryDomain
}

func convertDeliveryDomainAPIObjectToResourceData(d *schema.ResourceData, deliveryDomain *apiclient.DeliveryDomain) {
	//Store the data
	d.Set("id", deliveryDomain.Id)
	d.Set("stack_id", deliveryDomain.StackId)
	d.Set("domain", deliveryDomain.Domain)
	d.Set("updated_at", deliveryDomain.UpdatedAt)
}
