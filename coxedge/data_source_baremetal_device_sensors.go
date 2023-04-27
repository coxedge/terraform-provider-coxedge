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
)

func dataSourceBareMetalDeviceSensors() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBareMetalDeviceSensorsRead,
		Schema:      getBareMetalDeviceSensorsSetSchema(),
	}
}

func dataSourceBareMetalDeviceSensorsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)

	requestedId := d.Get("id").(string)

	bareMetalDeviceSensors, err := coxEdgeClient.GetBareMetalDeviceSensorsById(environmentName, organizationId, requestedId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("baremetal_device_sensors", flattenBareMetalDeviceSensorsData(&bareMetalDeviceSensors)); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(requestedId)
	return diags
}

func flattenBareMetalDeviceSensorsData(bareMetalDevicesSensors *[]apiclient.BareMetalDeviceSensor) []interface{} {
	if bareMetalDevicesSensors != nil {
		devices := make([]interface{}, len(*bareMetalDevicesSensors), len(*bareMetalDevicesSensors))

		for i, sensor := range *bareMetalDevicesSensors {
			item := make(map[string]interface{})
			item["id"] = sensor.Id
			item["ipmi_field"] = sensor.IpmiField
			item["ipmi_value"] = sensor.IpmiValue
			devices[i] = item
		}
		return devices
	}
	return make([]interface{}, 0)
}
