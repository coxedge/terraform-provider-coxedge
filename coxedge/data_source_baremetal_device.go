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
	"strconv"
	"time"
)

func dataSourceBareMetalDevice() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBareMetalRead,
		Schema:      getBareMetalDeviceSetSchema(),
	}
}

func dataSourceBareMetalRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)

	requestedId := d.Get("id").(string)
	if requestedId != "" {
		bareMetalDevice, err := coxEdgeClient.GetBareMetalDeviceById(environmentName, organizationId, requestedId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("baremetal_devices", flattenBareMetalDevicesData(&[]apiclient.BareMetalDevice{*bareMetalDevice}, true)); err != nil {
			return diag.FromErr(err)
		}
	} else {
		bareMetalDevices, err := coxEdgeClient.GetBareMetalDevices(environmentName, organizationId)
		if err != nil {
			return diag.FromErr(err)
		}

		if err := d.Set("baremetal_devices", flattenBareMetalDevicesData(&bareMetalDevices, false)); err != nil {
			return diag.FromErr(err)
		}
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags

}

func flattenBareMetalDevicesData(bareMetalDevices *[]apiclient.BareMetalDevice, isId bool) []interface{} {
	if bareMetalDevices != nil {
		devices := make([]interface{}, len(*bareMetalDevices), len(*bareMetalDevices))

		for i, device := range *bareMetalDevices {
			item := make(map[string]interface{})

			item["id"] = device.Id
			item["service_plan"] = device.ServicePlan
			item["name"] = device.Name
			item["hostname"] = device.Hostname
			item["device_type"] = device.DeviceType
			item["primary_ip"] = device.PrimaryIp
			item["status"] = device.Status
			item["monitors_total"] = device.MonitorsTotal
			item["monitors_up"] = device.MonitorsUp
			item["ipmi_address"] = device.IpmiAddress
			item["power_status"] = device.PowerStatus
			item["tags"] = device.Tags
			item["vendor"] = device.Vendor
			item["is_network_policy_available"] = device.IsNetworkPolicyAvailable
			item["change_id"] = device.ChangeId

			loc := make([]interface{}, 1, 1)
			locItem := make(map[string]interface{})
			locItem["facility"] = device.Location.Facility
			locItem["facility_title"] = device.Location.FacilityTitle
			loc[0] = locItem
			item["location"] = loc

			if isId {
				deviceIpDetail := make([]interface{}, 1, 1)
				ipDetail := make(map[string]interface{})
				ipDetail["primary_ip"] = device.DeviceDetail.DeviceIPDetail.PrimaryIP
				ipDetail["description"] = device.DeviceDetail.DeviceIPDetail.Description
				ipDetail["gateway_ip"] = device.DeviceDetail.DeviceIPDetail.GatewayIP
				ipDetail["subnet_mask"] = device.DeviceDetail.DeviceIPDetail.SubnetMask
				ipDetail["usable_ips"] = device.DeviceDetail.DeviceIPDetail.UsableIPs
				deviceIpDetail[0] = ipDetail

				deviceDetail := make([]interface{}, 1, 1)
				detail := make(map[string]interface{})
				detail["product_id"] = device.DeviceDetail.ProductID
				detail["service_plan"] = device.DeviceDetail.ServicePlan
				detail["processor"] = device.DeviceDetail.Processor
				detail["primary_hard_drive"] = device.DeviceDetail.PrimaryHardDrive
				detail["memory"] = device.DeviceDetail.Memory
				detail["operating_system"] = device.DeviceDetail.OperatingSystem
				detail["bandwidth"] = device.DeviceDetail.Bandwidth
				detail["internal_network"] = device.DeviceDetail.InternalNetwork
				detail["ddos"] = device.DeviceDetail.DDoS
				detail["raid_set_up"] = device.DeviceDetail.RaidSetUp
				detail["next_renew"] = device.DeviceDetail.NextRenew
				detail["device_ip_detail"] = deviceIpDetail
				deviceDetail[0] = detail
				item["device_detail"] = deviceDetail

				deviceInitialPassword := make([]interface{}, 1, 1)
				initialPassword := make(map[string]interface{})
				initialPassword["password_returns_until"] = device.DeviceInitialPassword.PasswordReturnsUntil
				initialPassword["password_expires"] = device.DeviceInitialPassword.PasswordExpires
				initialPassword["port"] = device.DeviceInitialPassword.Port
				initialPassword["user"] = device.DeviceInitialPassword.User
				deviceInitialPassword[0] = initialPassword
				item["device_initial_password"] = deviceInitialPassword

				deviceIPs := make([]interface{}, 1, 1)
				ips := make(map[string]interface{})
				ips["subnet"] = device.DeviceIPs.Subnet
				ips["netmask"] = device.DeviceIPs.Netmask
				ips["usable_ips"] = device.DeviceIPs.UsableIPs
				deviceIPs[0] = ips
				item["device_ips"] = deviceIPs
			}
			devices[i] = item
		}
		return devices
	}

	return make([]interface{}, 0)
}
