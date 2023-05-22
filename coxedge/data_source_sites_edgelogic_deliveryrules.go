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

func dataSourceSitesEdgeLogicDeliveryRules() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDeliveryRulesRead,
		Schema:      getEdgeLogicDeliveryRulesSetSchema(),
	}
}

func dataSourceDeliveryRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the request params from the data source block
	requestedId := d.Get("id").(string)
	requestedEnvironmentName := d.Get("environment_name").(string)
	requestedOrganizationId := d.Get("organization_id").(string)

	edgeLogic, err := coxEdgeClient.GetDeliveryRules(requestedEnvironmentName, requestedOrganizationId, requestedId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("edge_logic_delivery_rules", flattenEdgeLogicDeliveryRulesData(edgeLogic)); err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(requestedId)

	return diags
}

func flattenEdgeLogicDeliveryRulesData(apiDeliveryRule []apiclient.DeliveryRule) []interface{} {
	if apiDeliveryRule != nil {
		deliveryRules := make([]interface{}, len(apiDeliveryRule), len(apiDeliveryRule))

		for i, deliveryRule := range apiDeliveryRule {
			item := make(map[string]interface{})

			item["id"] = deliveryRule.Id
			item["stack_id"] = deliveryRule.StackId
			item["scope_id"] = deliveryRule.ScopeId
			item["name"] = deliveryRule.Name
			item["slug"] = deliveryRule.Slug
			item["actions"] = deliveryRule.Actions

			conditions := make([]map[string]interface{}, len(deliveryRule.Conditions))
			for j, condition := range deliveryRule.Conditions {
				conditionItem := make(map[string]interface{})
				conditionItem["trigger"] = condition.Trigger
				conditionItem["operator"] = condition.Operator
				conditionItem["http_methods"] = condition.HTTPMethods
				conditionItem["target"] = condition.Target
				conditions[j] = conditionItem
			}
			item["conditions"] = conditions

			actions := make([]map[string]interface{}, len(deliveryRule.Actions))
			for k, action := range deliveryRule.Actions {
				actionItem := make(map[string]interface{})
				actionItem["action_type"] = action.ActionType
				actionItem["cache_ttl"] = action.CacheTtl
				actionItem["redirect_url"] = action.RedirectUrl
				actionItem["header_pattern"] = action.HeaderPattern
				actionItem["passphrase"] = action.Passphrase
				actionItem["passphrase_field"] = action.PassphraseField
				actionItem["md5_token_field"] = action.MD5TokenField
				actionItem["ttl_field"] = action.TTLField
				actionItem["ip_address_filter"] = action.IPAddressFilter
				actionItem["url_signature_path_length"] = action.URLSignaturePathLength

				responseHeaders := make([]map[string]interface{}, len(action.ResponseHeaders))
				for l, header := range action.ResponseHeaders {
					responseHeaderItem := make(map[string]interface{})
					responseHeaderItem["key"] = header.Key
					responseHeaderItem["value"] = header.Value
					responseHeaders[l] = responseHeaderItem
				}
				actionItem["response_headers"] = responseHeaders

				originHeaders := make([]map[string]interface{}, len(action.OriginHeaders))
				for l, header := range action.OriginHeaders {
					originHeaderItem := make(map[string]interface{})
					originHeaderItem["key"] = header.Key
					originHeaderItem["value"] = header.Value
					originHeaders[l] = originHeaderItem
				}
				actionItem["origin_headers"] = originHeaders

				cdnHeaders := make([]map[string]interface{}, len(action.CDNHeaders))
				for l, header := range action.CDNHeaders {
					cdnHeaderItem := make(map[string]interface{})
					cdnHeaderItem["key"] = header.Key
					cdnHeaderItem["value"] = header.Value
					cdnHeaders[l] = cdnHeaderItem
				}
				actionItem["cdn_headers"] = cdnHeaders

				actions[k] = actionItem
			}
			item["actions"] = actions

			deliveryRules[i] = item
		}
		return deliveryRules
	}

	return make([]interface{}, 0)
}
