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
)

func dataSourceBareMetalDeviceCharts() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBareMetalDeviceChartsRead,
		Schema:      getBareMetalDeviceChartsSetSchema(),
	}
}

func dataSourceBareMetalDeviceChartsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)

	requestedId := d.Get("id").(string)

	customChart := d.Get("custom").(bool)

	if customChart {
		if d.Get("start_date").(string) == "" || d.Get("end_date").(string) == "" {
			diag := diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Missing required argument",
				Detail:   "start_date field and end_date field are required for custom charts",
			}
			diags = append(diags, diag)
			return diags
		}
		customRequest := convertResourceDataToBareMetalDeviceCustomChartAPIObject(d)
		customChart, err := coxEdgeClient.PostBareMetalDeviceCustomChartsById(customRequest, environmentName, organizationId, requestedId)
		if err != nil {
			return diag.FromErr(err)
		}
		tflog.Info(ctx, "Initiated Update. Awaiting task result.")
		//Await
		_, err = coxEdgeClient.AwaitTaskResolveWithDefaults(ctx, customChart.TaskId)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	bareMetalDeviceCharts, err := coxEdgeClient.GetBareMetalDeviceChartsById(environmentName, organizationId, requestedId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("baremetal_device_charts", flattenBareMetalDeviceChartsData(&bareMetalDeviceCharts)); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(requestedId)
	return diags
}

func convertResourceDataToBareMetalDeviceCustomChartAPIObject(d *schema.ResourceData) apiclient.CustomChartRequest {
	customRequest := apiclient.CustomChartRequest{
		StartDate: d.Get("start_date").(string),
		EndDate:   d.Get("end_date").(string),
	}
	return customRequest
}

func flattenBareMetalDeviceChartsData(bareMetalDevicesCharts *[]apiclient.BareMetalDeviceChart) []interface{} {
	if bareMetalDevicesCharts != nil {
		devices := make([]interface{}, len(*bareMetalDevicesCharts), len(*bareMetalDevicesCharts))

		for i, charts := range *bareMetalDevicesCharts {
			item := make(map[string]interface{})
			item["id"] = charts.Id
			item["filter"] = charts.Filter
			item["graph_image"] = charts.GraphImage
			item["interfaces"] = charts.Interfaces
			item["switch_id"] = charts.SwitchId
			devices[i] = item
		}
		return devices
	}
	return make([]interface{}, 0)
}
