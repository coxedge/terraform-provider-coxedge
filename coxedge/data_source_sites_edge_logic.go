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

func dataSourceSitesEdgeLogic() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSitesEdgeLogicRead,
		Schema:      getSitesEdgeLogicSetSchema(),
	}
}

func dataSourceSitesEdgeLogicRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the request params from the data source block
	requestedId := d.Get("id").(string)
	requestedEnvironmentName := d.Get("environment_name").(string)
	requestedOrganizationId := d.Get("organization_id").(string)

	edgeLogic, err := coxEdgeClient.GetPredefinedEdgeLogics(requestedEnvironmentName, requestedOrganizationId, requestedId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("predefined_edge_logic", flattenPredefinedEdgeLogicData(&[]apiclient.EdgeLogic{*edgeLogic})); err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenPredefinedEdgeLogicData(apiEdgeLogic *[]apiclient.EdgeLogic) []interface{} {
	if apiEdgeLogic != nil {
		edgeLogics := make([]interface{}, len(*apiEdgeLogic), len(*apiEdgeLogic))

		for i, edgeLogic := range *apiEdgeLogic {
			item := make(map[string]interface{})

			item["id"] = edgeLogic.Id
			item["stack_id"] = edgeLogic.StackId
			item["scope_id"] = edgeLogic.ScopeId
			item["force_www_enabled"] = edgeLogic.ForceWwwEnabled
			item["robots_txt_enabled"] = edgeLogic.RobotsTxtEnabled
			item["robots_txt_file"] = edgeLogic.RobotTxtFile
			item["pseudo_streaming_enabled"] = edgeLogic.PseudoStreamingEnabled
			item["referrer_protection_enabled"] = edgeLogic.ReferrerProtectionEnabled
			item["allow_empty_referrer"] = edgeLogic.AllowEmptyReferrer
			item["referrer_list"] = edgeLogic.ReferrerList

			edgeLogics[i] = item
		}
		return edgeLogics
	}

	return make([]interface{}, 0)
}
