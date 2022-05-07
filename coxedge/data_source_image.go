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

func dataSourceImage() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceImageRead,
		Schema:      getImageSetSchema(),
	}
}

func dataSourceImageRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the request params from the data source block
	requestedId := d.Get("id").(string)
	requestedEnvironment := d.Get("environment").(string)

	if requestedId != "" {
		org, err := coxEdgeClient.GetImage(requestedEnvironment, requestedId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("images", flattenImageData(&[]apiclient.Image{*org})); err != nil {
			return diag.FromErr(err)
		}
	} else {
		orgs, err := coxEdgeClient.GetImages(requestedEnvironment)
		if err != nil {
			return diag.FromErr(err)
		}

		if err := d.Set("images", flattenImageData(&orgs)); err != nil {
			return diag.FromErr(err)
		}
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenImageData(apiImages *[]apiclient.Image) []interface{} {
	if apiImages != nil {
		images := make([]interface{}, len(*apiImages), len(*apiImages))

		for i, img := range *apiImages {
			item := make(map[string]interface{})

			item["id"] = img.Id
			item["stack_id"] = img.StackId
			item["family"] = img.Family
			item["tag"] = img.Tag
			item["slug"] = img.Slug
			item["status"] = img.Status
			item["created_at"] = img.CreatedAt
			item["description"] = img.Description
			item["reference"] = img.Reference

			images[i] = item
		}

		return images
	}

	return make([]interface{}, 0)
}
