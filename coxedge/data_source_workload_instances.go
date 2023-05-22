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

func dataWorkloadInstances() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWorkloadInstancesRead,
		Schema:      getWorkloadInstanceSetSchema(),
	}
}

func dataSourceWorkloadInstancesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	requestedId := d.Get("id").(string)
	environmentName := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)

	workloadInstances, err := coxEdgeClient.GetWorkloadInstances(environmentName, organizationId, requestedId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("workload_instances", flattenWorkloadInstancesData(workloadInstances)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenWorkloadInstancesData(workloadInstances []apiclient.WorkloadInstance) []interface{} {
	if workloadInstances != nil {
		instances := make([]interface{}, len(workloadInstances), len(workloadInstances))

		for i, instance := range workloadInstances {
			item := make(map[string]interface{})

			item["stack_id"] = instance.StackId
			item["workload_id"] = instance.WorkloadId
			item["name"] = instance.Name
			item["ip_address"] = instance.IPAddress
			item["public_ip_address"] = instance.PublicIPAddress
			item["location"] = instance.Location
			item["created_date"] = instance.CreatedDate
			item["started_date"] = instance.StartedDate
			item["id"] = instance.Id
			item["status"] = instance.Status
			instances[i] = item
		}

		return instances
	}

	return make([]interface{}, 0)
}
