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

func dataSourceBareMetal() *schema.Resource {
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
		if err := d.Set("baremetal_devices", flattenBareMetalDevicesData(&[]apiclient.BareMetalDevice{*bareMetalDevice})); err != nil {
			return diag.FromErr(err)
		}
	} else {
		bareMetalDevices, err := coxEdgeClient.GetBareMetalDevices(environmentName, organizationId)
		if err != nil {
			return diag.FromErr(err)
		}

		if err := d.Set("baremetal_devices", flattenBareMetalDevicesData(&bareMetalDevices)); err != nil {
			return diag.FromErr(err)
		}
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags

}

func flattenBareMetalDevicesData(bareMetalDevices *[]apiclient.BareMetalDevice) []interface{} {
	if bareMetalDevices != nil {
		devices := make([]interface{}, len(*bareMetalDevices), len(*bareMetalDevices))

		for i, device := range *bareMetalDevices {
			item := make(map[string]interface{})

			item["id"] = device.Id
			item["service_plan"] = device.ServicePlan
			item["name"] = device.Name
			item["device_type"] = device.DeviceType
			item["primary_ip"] = device.PrimaryIp
			item["status"] = device.Status
			item["monitors_total"] = device.MonitorsTotal
			item["monitors_up"] = device.MonitorsUp
			item["ipmi_address"] = device.IpmiAddress
			item["power_status"] = device.PowerStatus
			item["tags"] = device.Tags

			loc := make([]interface{}, 1, 1)
			locItem := make(map[string]interface{})
			locItem["facility"] = device.Location.Facility
			locItem["facility_title"] = device.Location.FacilityTitle
			loc[0] = locItem

			item["location"] = loc
			devices[i] = item
		}

		return devices
	}

	return make([]interface{}, 0)
}
