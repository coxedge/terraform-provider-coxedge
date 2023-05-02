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

func resourceBareMetaSSHKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBareMetalSSHKeyCreate,
		ReadContext:   resourceBareMetalSSHKeyRead,
		UpdateContext: resourceBareMetalSSHKeyUpdate,
		DeleteContext: resourceBareMetalSSHKeyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getBareMetalSSHKeyResourceSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceBareMetalSSHKeyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if d.Get("name").(string) == "" {
		diag := diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Missing required argument",
			Detail:   "The argument 'name'' is required, but no definition was found.",
		}
		diags = append(diags, diag)
		return diags
	}
	if d.Get("public_key").(string) == "" {
		diag := diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Missing required argument",
			Detail:   "The argument 'public_key'' is required, but no definition was found.",
		}
		diags = append(diags, diag)
		return diags
	}

	createRequest := convertResourceDataToBareMetalSSHKeyCreateAPIObject(d)
	environmentName := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)
	//Call the API
	createdSSHKey, err := coxEdgeClient.CreateBareMetalSSHKey(createRequest, environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Create. Awaiting task result.")

	timeout := d.Timeout(schema.TimeoutCreate)
	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, createdSSHKey.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the Id
	d.SetId(taskResult.Data.Result.Id)

	return diags
}

func resourceBareMetalSSHKeyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	//check the id comes with id,environment_name & organization_id, then split the value -> in case of importing the resource
	//format is <ssh_key_d>:<environment_name>:<organization_id>
	if strings.Contains(d.Id(), ":") {
		keys := strings.Split(d.Id(), ":")
		d.SetId(keys[0])
		d.Set("environment_name", keys[1])
		d.Set("organization_id", keys[2])

	}
	//Get the resource Id
	resourceId := d.Id()
	environmentName := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)
	bareMetalSSHKey, err := coxEdgeClient.GetBareMetalSSHKeyById(environmentName, organizationId, resourceId)
	if err != nil {
		return diag.FromErr(err)
	}
	convertSSHKeyAPIObjectToResourceData(d, bareMetalSSHKey)
	return diags
}

func resourceBareMetalSSHKeyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Errorf("Unfortunately, it is not possible to update BareMetal SSHKey from this resource. Please refer to the documentation.")
}

func resourceBareMetalSSHKeyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	resourceId := d.Id()
	environmentName := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)
	//Call the API
	deletedSSHKey, err := coxEdgeClient.DeleteBareMetalSSHKeyById(environmentName, organizationId, resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Delete. Awaiting task result.")

	timeout := d.Timeout(schema.TimeoutDelete)
	//Await
	_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, deletedSSHKey.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags

}

func convertResourceDataToBareMetalSSHKeyCreateAPIObject(d *schema.ResourceData) apiclient.CreateBareMetalSSHKeyRequest {

	bareMetalSSHKey := apiclient.CreateBareMetalSSHKeyRequest{
		Name:      d.Get("name").(string),
		PublicKey: d.Get("public_key").(string),
	}
	return bareMetalSSHKey
}

func convertSSHKeyAPIObjectToResourceData(d *schema.ResourceData, bareMetalSSHKey *apiclient.BareMetalSSHKey) {
	d.Set("id", bareMetalSSHKey.Id)
	d.Set("name", bareMetalSSHKey.Name)
	d.Set("public_key", bareMetalSSHKey.PublicKey)
}
