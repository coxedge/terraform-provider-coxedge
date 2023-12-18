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

func resourceBareMetalDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Errorf("Creating devices from this resource is not possible. However, you can import devices using the following command: \"terraform import coxedge_baremetal_device.device <device_id>:<environment_name>:<organization_id>\". Please refer to the documentation for instructions on device creation.")
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

	var vendor string
	if strings.HasPrefix(resourceId, "HV_") {
		vendor = "HIVELOCITY"
	} else {
		vendor = "METALSOFT"
	}
	editDevice := convertResourceDataToBareMetalDeviceEditAPIObject(d, vendor)
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
	oldValue, newValue := d.GetChange("power_status")
	if newValue != oldValue {
		powerStatus := d.Get("power_status").(string)
		if powerStatus != "" {
			var operation string
			if powerStatus == "ON" {
				operation = "device-on"
			} else if powerStatus == "OFF" {
				operation = "device-off"
			} else if powerStatus == "RESTART" {
				operation = "device-restart"
			} else if powerStatus == "soft-device-off" {
				operation = "soft-device-off"
			}
			powerStatusDevice, err := coxEdgeClient.EditBareMetalDevicePowerById(resourceId, operation, environmentName, organizationId, vendor)
			if err != nil {
				return diag.FromErr(err)
			}
			//Await
			_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, powerStatusDevice.TaskId, timeout)
			if err != nil {
				return diag.FromErr(err)
			}
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

	deviceIpDetail := make([]interface{}, 1, 1)
	ipDetail := make(map[string]interface{})
	ipDetail["primary_ip"] = bareMetalDevice.DeviceDetail.DeviceIPDetail.PrimaryIP
	ipDetail["description"] = bareMetalDevice.DeviceDetail.DeviceIPDetail.Description
	ipDetail["gateway_ip"] = bareMetalDevice.DeviceDetail.DeviceIPDetail.GatewayIP
	ipDetail["subnet_mask"] = bareMetalDevice.DeviceDetail.DeviceIPDetail.SubnetMask
	ipDetail["usable_ips"] = bareMetalDevice.DeviceDetail.DeviceIPDetail.UsableIPs
	deviceIpDetail[0] = ipDetail

	deviceDetail := make([]interface{}, 1, 1)
	detail := make(map[string]interface{})
	detail["product_id"] = bareMetalDevice.DeviceDetail.ProductID
	detail["service_plan"] = bareMetalDevice.DeviceDetail.ServicePlan
	detail["processor"] = bareMetalDevice.DeviceDetail.Processor
	detail["primary_hard_drive"] = bareMetalDevice.DeviceDetail.PrimaryHardDrive
	detail["memory"] = bareMetalDevice.DeviceDetail.Memory
	detail["operating_system"] = bareMetalDevice.DeviceDetail.OperatingSystem
	detail["bandwidth"] = bareMetalDevice.DeviceDetail.Bandwidth
	detail["internal_network"] = bareMetalDevice.DeviceDetail.InternalNetwork
	detail["ddos"] = bareMetalDevice.DeviceDetail.DDoS
	detail["raid_set_up"] = bareMetalDevice.DeviceDetail.RaidSetUp
	detail["next_renew"] = bareMetalDevice.DeviceDetail.NextRenew
	detail["device_ip_detail"] = deviceIpDetail
	deviceDetail[0] = detail
	d.Set("device_detail", deviceDetail)

	deviceInitialPassword := make([]interface{}, 1, 1)
	initialPassword := make(map[string]interface{})
	initialPassword["password_returns_until"] = bareMetalDevice.DeviceInitialPassword.PasswordReturnsUntil
	initialPassword["password_expires"] = bareMetalDevice.DeviceInitialPassword.PasswordExpires
	initialPassword["port"] = bareMetalDevice.DeviceInitialPassword.Port
	initialPassword["user"] = bareMetalDevice.DeviceInitialPassword.User
	deviceInitialPassword[0] = initialPassword
	d.Set("device_initial_password", deviceInitialPassword)

	deviceIPs := make([]interface{}, 1, 1)
	ips := make(map[string]interface{})
	ips["subnet"] = bareMetalDevice.DeviceIPs.Subnet
	ips["netmask"] = bareMetalDevice.DeviceIPs.Netmask
	ips["usable_ips"] = bareMetalDevice.DeviceIPs.UsableIPs
	deviceIPs[0] = ips
	d.Set("device_ips", deviceIPs)
}

func convertResourceDataToBareMetalDeviceEditAPIObject(d *schema.ResourceData, vendor string) apiclient.EditBareMetalDeviceRequest {

	editDevice := apiclient.EditBareMetalDeviceRequest{
		Name: d.Get("name").(string),
	}

	if vendor == "HIVELOCITY" {
		editDevice.Hostname = d.Get("hostname").(string)
	}

	tgs := d.Get("tags").([]interface{})
	tags := make([]string, len(tgs))
	for i, data := range tgs {
		tags[i] = data.(string)
	}
	editDevice.Tags = tags
	return editDevice
}
