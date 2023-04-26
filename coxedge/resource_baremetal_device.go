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

func resourceBareMetalDevice() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBareMetalDeviceCreate,
		ReadContext:   resourceBareMetalDeviceRead,
		UpdateContext: resourceBareMetalDeviceUpdate,
		DeleteContext: resourceBareMetalDeviceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getBareMetalDeviceSchema(),
		Timeouts: &schema.ResourceTimeout{
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceBareMetalDeviceCreate(ctx context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	return nil
}

func resourceBareMetalDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	//check the id comes with id,environment_name & organization_id, then split the value -> in case of importing the resource
	//format is <device_id>:<environment_name>:<organization_id>
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
	bareMetalDevice, err := coxEdgeClient.GetBareMetalDeviceById(environmentName, organizationId, resourceId)
	if err != nil {
		return diag.FromErr(err)
	}
	convertDeviceAPIObjectToResourceData(d, bareMetalDevice)
	return diags
}

func resourceBareMetalDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the resource Id
	resourceId := d.Id()
	environmentName := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)

	editDevice := convertResourceDataToBareMetalDeviceEditAPIObject(d)
	//Edit the BareMetal device
	editedDevice, err := coxEdgeClient.EditBareMetalDeviceById(editDevice, resourceId, environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}
	timeout := d.Timeout(schema.TimeoutUpdate)
	tflog.Info(ctx, "Initiated Update. Awaiting task result.")

	//Await
	_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, editedDevice.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}
	powerStatus := d.Get("power_status").(string)
	if powerStatus != "" {
		var operation string
		if powerStatus == "ON" {
			operation = "device-on"
		} else {
			operation = "device-off"
		}
		powerStatusDevice, err := coxEdgeClient.EditBareMetalDevicePowerById(resourceId, operation, environmentName, organizationId)
		if err != nil {
			return diag.FromErr(err)
		}
		//Await
		_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, powerStatusDevice.TaskId, timeout)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceBareMetalDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	resourceId := d.Id()
	environmentName := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)

	//Delete the BareMetal device
	deletedDevice, err := coxEdgeClient.DeleteBareMetalDeviceById(resourceId, environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Delete. Awaiting task result.")

	timeout := d.Timeout(schema.TimeoutDelete)
	//Await
	_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, deletedDevice.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func convertDeviceAPIObjectToResourceData(d *schema.ResourceData, bareMetalDevice *apiclient.BareMetalDevice) {
	d.Set("id", bareMetalDevice.Id)
	d.Set("service_plan", bareMetalDevice.ServicePlan)
	d.Set("name", bareMetalDevice.Name)
	d.Set("hostname", bareMetalDevice.Hostname)
	d.Set("device_type", bareMetalDevice.DeviceType)
	d.Set("primary_ip", bareMetalDevice.PrimaryIp)
	d.Set("status", bareMetalDevice.Status)
	d.Set("monitors_total", bareMetalDevice.MonitorsTotal)
	d.Set("monitors_up", bareMetalDevice.MonitorsUp)
	d.Set("ipmi_address", bareMetalDevice.IpmiAddress)
	d.Set("power_status", bareMetalDevice.PowerStatus)
	d.Set("tags", bareMetalDevice.Tags)

	loc := make([]interface{}, 1, 1)
	locItem := make(map[string]interface{})
	locItem["facility"] = bareMetalDevice.Location.Facility
	locItem["facility_title"] = bareMetalDevice.Location.FacilityTitle
	loc[0] = locItem

	d.Set("location", loc)
}

func convertResourceDataToBareMetalDeviceEditAPIObject(d *schema.ResourceData) apiclient.EditBareMetalDeviceRequest {

	editDevice := apiclient.EditBareMetalDeviceRequest{
		Name:     d.Get("name").(string),
		Hostname: d.Get("hostname").(string),
	}

	tgs := d.Get("tags").([]interface{})
	tags := make([]string, len(tgs))
	for i, data := range tgs {
		tags[i] = data.(string)
	}
	editDevice.Tags = tags
	return editDevice
}
