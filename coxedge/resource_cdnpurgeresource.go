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

func resourceCDNPurgeResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCDNPurgeResourceCreate,
		ReadContext:   resourceCDNPurgeResourceRead,
		DeleteContext: resourceCDNPurgeResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getCDNPurgeResourceSchema(),
	}
}

func resourceCDNPurgeResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Convert resource data to API Object
	newCDNPurgeResource := convertResourceDataToCDNPurgeResourceCreateAPIObject(d)

	//Call the API
	createdCDNPurgeResource, err := coxEdgeClient.PurgeCDN(
		d.Get("environment_name").(string),
		d.Get("site_id").(string),
		newCDNPurgeResource,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the ID
	d.SetId(createdCDNPurgeResource.TaskId)

	return diags
}

func resourceCDNPurgeResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourceCDNPurgeResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//NOOP

	return diags
}

func convertResourceDataToCDNPurgeResourceCreateAPIObject(d *schema.ResourceData) apiclient.CDNPurgeOptions {
	//Create update script struct
	updatedScript := apiclient.CDNPurgeOptions{
		PurgeType: d.Get("purge_type").(string),
		Items: []struct {
			URL             string   `json:"url,omitempty"`
			Recursive       bool     `json:"recursive,omitempty"`
			InvalidateOnly  bool     `json:"invalidateOnly,omitempty"`
			PurgeAllDynamic bool     `json:"purgeAllDynamic,omitempty"`
			Headers         []string `json:"headers,omitempty"`
			PurgeSelector   struct {
				SelectorName           string `json:"selectorName,omitempty"`
				SelectorValue          string `json:"selectorValue,omitempty"`
				SelectorType           string `json:"selectorType,omitempty"`
				SelectorValueDelimiter string `json:"selectorValueDelimiter,omitempty"`
			} `json:"purgeSelector,omitempty"`
		}{},
	}

	items, hasItems := d.GetOk("items")
	if hasItems {
		for _, rawItem := range items.([]map[string]interface{}) {
			newItem := struct {
				URL             string   `json:"url,omitempty"`
				Recursive       bool     `json:"recursive,omitempty"`
				InvalidateOnly  bool     `json:"invalidateOnly,omitempty"`
				PurgeAllDynamic bool     `json:"purgeAllDynamic,omitempty"`
				Headers         []string `json:"headers,omitempty"`
				PurgeSelector   struct {
					SelectorName           string `json:"selectorName,omitempty"`
					SelectorValue          string `json:"selectorValue,omitempty"`
					SelectorType           string `json:"selectorType,omitempty"`
					SelectorValueDelimiter string `json:"selectorValueDelimiter,omitempty"`
				} `json:"purgeSelector,omitempty"`
			}{
				URL: rawItem["url"].(string),
			}

			//Optional fields
			recursive, hasRecursive := rawItem["recursive"]
			if hasRecursive {
				newItem.Recursive = recursive.(bool)
			}
			invalidateOnly, hasInvalidateOnly := rawItem["invalidate_only"]
			if hasInvalidateOnly {
				newItem.InvalidateOnly = invalidateOnly.(bool)
			}
			purgeAllDynamic, hasPurgeAllDynamic := rawItem["purge_all_dynamic"]
			if hasPurgeAllDynamic {
				newItem.PurgeAllDynamic = purgeAllDynamic.(bool)
			}
			headers, hasHeaders := rawItem["headers"]
			if hasHeaders {
				newItem.Headers = headers.([]string)
			}

			//Optional Nested Structure
			purgeSelector, hasPurgeSelector := rawItem["purge_selector"]
			if hasPurgeSelector {
				selectors := purgeSelector.([]map[string]string)
				name, hasName := selectors[0]["selector_name"]
				if hasName {
					newItem.PurgeSelector.SelectorName = name
				}
				selType, hasType := selectors[0]["selector_type"]
				if hasType {
					newItem.PurgeSelector.SelectorType = selType
				}
				selValue, hasValue := selectors[0]["selector_value"]
				if hasValue {
					newItem.PurgeSelector.SelectorValue = selValue
				}
				selValDel, hasDel := selectors[0]["selector_value_delimiter"]
				if hasDel {
					newItem.PurgeSelector.SelectorValueDelimiter = selValDel
				}
			}

			updatedScript.Items = append(updatedScript.Items, newItem)

		}
	}

	return updatedScript
}
