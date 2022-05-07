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

func dataSourceEnvironment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEnvironmentRead,
		Schema:      getEnvironmentSetSchema(),
	}
}

func dataSourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	requestedId := d.Get("id").(string)
	if requestedId != "" {
		org, err := coxEdgeClient.GetEnvironment(requestedId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("environments", flattenEnvironmentData(&[]apiclient.Environment{*org})); err != nil {
			return diag.FromErr(err)
		}
	} else {
		orgs, err := coxEdgeClient.GetEnvironments()
		if err != nil {
			return diag.FromErr(err)
		}

		if err := d.Set("environments", flattenEnvironmentData(&orgs)); err != nil {
			return diag.FromErr(err)
		}
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenEnvironmentData(environments *[]apiclient.Environment) []interface{} {
	if environments != nil {
		tfObjs := make([]interface{}, len(*environments), len(*environments))

		for i, apiItem := range *environments {
			item := make(map[string]interface{})

			item["id"] = apiItem.Id
			item["name"] = apiItem.Name
			item["description"] = apiItem.Description
			item["creation_date"] = apiItem.CreationDate

			tfObjs[i] = item
		}

		return tfObjs
	}

	return make([]interface{}, 0)
}
