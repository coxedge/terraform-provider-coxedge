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

func resourceSitesEdgeLogicDeliveryRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEdgeLogicDeliveryRuleCreate,
		ReadContext:   resourceEdgeLogicDeliveryRuleRead,
		UpdateContext: resourceEdgeLogicDeliveryRuleUpdate,
		DeleteContext: resourceEdgeLogicDeliveryRuleDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getEdgeLogicDeliveryRuleResourceSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceEdgeLogicDeliveryRuleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	siteId := d.Get("site_id").(string)
	environmentName := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)
	request := convertResourceDataToEdgeLogicDeliveryRuleAPIObject(d)

	createDeliveryRule, err := coxEdgeClient.AddDeliveryRule(request, environmentName, organizationId, siteId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Create. Awaiting task result.")

	timeout := d.Timeout(schema.TimeoutCreate)
	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, createDeliveryRule.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the Id
	d.SetId(taskResult.Data.Result.Id)

	return diags
}

func resourceEdgeLogicDeliveryRuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//check the id comes with id,environment_name & organization_id, then split the value -> in case of importing the resource
	//format is <delivery_rule_id>:<site_id>:<environment_name>:<organization_id>
	if strings.Contains(d.Id(), ":") {
		keys := strings.Split(d.Id(), ":")
		d.Set("id", keys[0])
		d.Set("site_id", keys[1])
		d.Set("environment_name", keys[2])
		d.Set("organization_id", keys[3])
	}
	//Get the resource Id
	resourceId := d.Get("id").(string)
	siteId := d.Get("site_id").(string)
	organizationId := d.Get("organization_id").(string)
	environmentName := d.Get("environment_name").(string)

	deliveryRules, err := coxEdgeClient.GetDeliveryRules(environmentName, organizationId, resourceId, siteId)
	if err != nil {
		return diag.FromErr(err)
	}

	convertEdgeLogicDeliveryRuleAPIObjectToResourceData(d, deliveryRules)
	return diags
}

func resourceEdgeLogicDeliveryRuleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	resourceId := d.Get("id").(string)
	siteId := d.Get("site_id").(string)
	environmentName := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)
	request := convertResourceDataToEdgeLogicDeliveryRuleAPIObject(d)

	updateDeliveryRule, err := coxEdgeClient.UpdateDeliveryRule(request, environmentName, organizationId, resourceId, siteId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Update. Awaiting task result.")

	timeout := d.Timeout(schema.TimeoutUpdate)
	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, updateDeliveryRule.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(taskResult.Data.Result.Id)
	return diags
}

func resourceEdgeLogicDeliveryRuleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	resourceId := d.Get("id").(string)
	siteId := d.Get("site_id").(string)
	environmentName := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)

	deleteDeliveryRule, err := coxEdgeClient.DeleteDeliveryRule(environmentName, organizationId, resourceId, siteId)
	if err != nil {
		return diag.FromErr(err)
	}
	timeout := d.Timeout(schema.TimeoutDelete)
	//Await
	_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, deleteDeliveryRule.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}

func convertEdgeLogicDeliveryRuleAPIObjectToResourceData(d *schema.ResourceData, deliveryRule *apiclient.DeliveryRule) {
	//Store the data
	d.Set("id", deliveryRule.Id)
	d.Set("name", deliveryRule.Name)
	d.Set("slug", deliveryRule.Slug)
	d.Set("stack_id", deliveryRule.StackId)
	d.Set("scope_id", deliveryRule.ScopeId)
	d.Set("site_id", deliveryRule.SiteId)

	conditions := make([]map[string]interface{}, len(deliveryRule.Conditions), len(deliveryRule.Conditions))
	for i, cndtn := range deliveryRule.Conditions {
		item := make(map[string]interface{})
		item["trigger"] = cndtn.Trigger
		item["operator"] = cndtn.Operator
		item["http_methods"] = cndtn.HTTPMethods
		item["target"] = cndtn.Target
		conditions[i] = item
	}
	d.Set("conditions", conditions)

	actions := make([]map[string]interface{}, len(deliveryRule.Actions), len(deliveryRule.Actions))
	for i, actn := range deliveryRule.Actions {
		item := make(map[string]interface{})
		item["action_type"] = actn.ActionType
		item["cache_ttl"] = actn.CacheTtl
		item["redirect_url"] = actn.RedirectUrl
		item["header_pattern"] = actn.HeaderPattern
		item["passphrase"] = actn.Passphrase
		item["passphrase_field"] = actn.PassphraseField
		item["md5_token_field"] = actn.MD5TokenField
		item["ttl_field"] = actn.TTLField
		item["ip_address_filter"] = actn.IPAddressFilter
		item["url_signature_path_length"] = actn.URLSignaturePathLength

		responseHeaders := make([]map[string]interface{}, len(actn.ResponseHeaders), len(actn.ResponseHeaders))
		for j, header := range actn.ResponseHeaders {
			itemHeader := make(map[string]interface{})
			itemHeader["key"] = header.Key
			itemHeader["value"] = header.Value
			responseHeaders[j] = itemHeader
		}
		item["response_headers"] = responseHeaders

		originHeaders := make([]map[string]interface{}, len(actn.OriginHeaders), len(actn.OriginHeaders))
		for j, header := range actn.OriginHeaders {
			itemHeader := make(map[string]interface{})
			itemHeader["key"] = header.Key
			itemHeader["value"] = header.Value
			originHeaders[j] = itemHeader
		}
		item["origin_headers"] = originHeaders

		cdnHeaders := make([]map[string]interface{}, len(actn.CDNHeaders), len(actn.CDNHeaders))
		for j, header := range actn.CDNHeaders {
			itemHeader := make(map[string]interface{})
			itemHeader["key"] = header.Key
			itemHeader["value"] = header.Value
			cdnHeaders[j] = itemHeader
		}
		item["cdn_headers"] = cdnHeaders

		actions[i] = item
	}
	d.Set("actions", actions)
}

func convertResourceDataToEdgeLogicDeliveryRuleAPIObject(d *schema.ResourceData) apiclient.DeliveryRuleRequest {

	request := apiclient.DeliveryRuleRequest{
		Id:   d.Get("id").(string),
		Name: d.Get("name").(string),
	}

	for _, entry := range d.Get("conditions").([]interface{}) {
		conditionEntry := entry.(map[string]interface{})
		condition := apiclient.ConditionRequest{
			Trigger:  conditionEntry["trigger"].(string),
			Operator: conditionEntry["operator"].(string),
			Target:   conditionEntry["target"].(string),
		}
		condition.HTTPMethods = make([]string, len(conditionEntry["http_methods"].([]interface{})))
		for i, method := range conditionEntry["http_methods"].([]interface{}) {
			condition.HTTPMethods[i] = method.(string)
		}
		request.Condition = append(request.Condition, condition)
	}

	for _, entry := range d.Get("actions").([]interface{}) {
		actionEntry := entry.(map[string]interface{})
		action := apiclient.ActionRequest{
			ActionType:             actionEntry["action_type"].(string),
			CacheTtl:               actionEntry["cache_ttl"].(int),
			RedirectUrl:            actionEntry["redirect_url"].(string),
			HeaderPattern:          actionEntry["header_pattern"].(string),
			Passphrase:             actionEntry["passphrase"].(string),
			PassphraseField:        actionEntry["passphrase_field"].(string),
			MD5TokenField:          actionEntry["md5_token_field"].(string),
			TTLField:               actionEntry["ttl_field"].(string),
			IPAddressFilter:        actionEntry["ip_address_filter"].(string),
			URLSignaturePathLength: actionEntry["url_signature_path_length"].(string),
		}

		action.ResponseHeaders = make([]apiclient.HeaderRequest, len(actionEntry["response_headers"].([]interface{})))
		for i, resEntry := range actionEntry["response_headers"].([]interface{}) {
			respEntry := resEntry.(map[string]interface{})
			header := apiclient.HeaderRequest{
				Key:   respEntry["key"].(string),
				Value: respEntry["value"].(string),
			}
			action.ResponseHeaders[i] = header
		}

		action.OriginHeaders = make([]apiclient.HeaderRequest, len(actionEntry["origin_headers"].([]interface{})))
		for i, originEntry := range actionEntry["origin_headers"].([]interface{}) {
			orgEntry := originEntry.(map[string]interface{})
			header := apiclient.HeaderRequest{
				Key:   orgEntry["key"].(string),
				Value: orgEntry["value"].(string),
			}
			action.OriginHeaders[i] = header
		}

		action.CDNHeaders = make([]apiclient.HeaderRequest, len(actionEntry["cdn_headers"].([]interface{})))
		for i, cdnEntry := range actionEntry["cdn_headers"].([]interface{}) {
			cdEntry := cdnEntry.(map[string]interface{})
			header := apiclient.HeaderRequest{
				Key:   cdEntry["key"].(string),
				Value: cdEntry["value"].(string),
			}
			action.CDNHeaders[i] = header
		}

		request.Action = append(request.Action, action)
	}

	return request
}
