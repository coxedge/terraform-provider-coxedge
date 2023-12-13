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

func resourceBareMetalDeviceIPMI() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBareMetalDeviceIPMICreate,
		ReadContext:   resourceBareMetalDeviceIPMIRead,
		UpdateContext: resourceBareMetalDeviceIPMIUpdate,
		DeleteContext: resourceBareMetalDeviceIPMIDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getBareMetalDeviceConnectIPMISchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceBareMetalDeviceIPMICreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ipmiObj := convertResourceDataToBareMetalDeviceIPMICreateAPIObject(d)
	environmentName := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)
	resourceId := d.Get("device_id").(string)

	if strings.HasPrefix(resourceId, "HV_") {
		//Call the API
		connectIPMI, err := coxEdgeClient.PostBareMetalDeviceConnectToIPMIById(ipmiObj, environmentName, organizationId, resourceId)
		if err != nil {
			return diag.FromErr(err)
		}

		tflog.Info(ctx, "Initiated Create. Awaiting task result.")

		timeout := d.Timeout(schema.TimeoutCreate)
		//Await
		_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, connectIPMI.TaskId, timeout)
		if err != nil {
			return diag.FromErr(err)
		}

		//Save the Id
		d.SetId(resourceId)
	} else {
		diag := diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "IPMI not available for METALSOFT",
			Detail:   "IPMI not available for METALSOFT",
		}
		diags = append(diags, diag)
		return diags
	}

	return diags
}

func resourceBareMetalDeviceIPMIRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return nil
}

func resourceBareMetalDeviceIPMIUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return diag.Errorf("Unfortunately, it is not possible to update BareMetal devices IPMI from this resource. Please refer to the documentation.")
}

func resourceBareMetalDeviceIPMIDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)
	resourceId := d.Get("device_id").(string)

	if strings.HasPrefix(resourceId, "HV_") {
		//Call the API
		connectIPMI, err := coxEdgeClient.PostBareMetalDeviceClearIPMIById(environmentName, organizationId, resourceId)
		if err != nil {
			return diag.FromErr(err)
		}

		tflog.Info(ctx, "Initiated Delete. Awaiting task result.")

		timeout := d.Timeout(schema.TimeoutDelete)
		//Await
		_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, connectIPMI.TaskId, timeout)
		if err != nil {
			return diag.FromErr(err)
		}

		d.SetId("")
	} else {
		diag := diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "IPMI not available for METALSOFT",
			Detail:   "IPMI not available for METALSOFT",
		}
		diags = append(diags, diag)
		return diags
	}
	return diags
}

func convertResourceDataToBareMetalDeviceIPMICreateAPIObject(d *schema.ResourceData) apiclient.ConnectIPMIRequest {
	ipmiRequest := apiclient.ConnectIPMIRequest{
		CustomIP: d.Get("custom_ip").(string),
	}
	return ipmiRequest
}
