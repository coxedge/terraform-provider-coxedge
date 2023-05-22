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

func dataSourceSitesEdgeLogicCustomRules() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCustomRulesRead,
		Schema:      getEdgeLogicCustomRulesSetSchema(),
	}
}

func dataSourceCustomRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the request params from the data source block
	requestedId := d.Get("id").(string)
	requestedEnvironmentName := d.Get("environment_name").(string)
	requestedOrganizationId := d.Get("organization_id").(string)

	edgeLogic, err := coxEdgeClient.GetCustomRules(requestedEnvironmentName, requestedOrganizationId, requestedId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("edge_logic_custom_rules", flattenEdgeLogicCustomRulesData(edgeLogic)); err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(requestedId)

	return diags
}

func flattenEdgeLogicCustomRulesData(apiCustomRule []apiclient.CustomRule) []interface{} {
	if apiCustomRule != nil {
		customRules := make([]interface{}, len(apiCustomRule), len(apiCustomRule))

		for i, customRule := range apiCustomRule {
			item := make(map[string]interface{})

			item["id"] = customRule.Id
			item["stack_id"] = customRule.StackId
			item["name"] = customRule.Name
			item["site_id"] = customRule.SiteId
			item["notes"] = customRule.Notes
			item["type"] = customRule.Type
			item["enabled"] = customRule.Enabled
			item["action"] = customRule.Action
			item["action_duration"] = customRule.ActionDuration
			item["status_code"] = customRule.StatusCode
			item["nb_request"] = customRule.NBRequest
			item["duration"] = customRule.Duration
			item["path_reg_exp"] = customRule.PathRegExp
			item["http_methods"] = customRule.HttpMethods
			item["ip_addresses"] = customRule.IPAddresses

			conditions := make([]map[string]interface{}, len(customRule.Conditions))
			for j, condition := range customRule.Conditions {
				conditionItem := make(map[string]interface{})
				conditionItem["type"] = condition.Type
				conditionItem["operation"] = condition.Operation
				conditionItem["value_list"] = condition.ValueList
				conditionItem["value"] = condition.Value
				conditionItem["end_value"] = condition.EndValue
				conditionItem["header"] = condition.Header
				conditions[j] = conditionItem
			}
			item["conditions"] = conditions

			customRules[i] = item
		}
		return customRules
	}

	return make([]interface{}, 0)
}
