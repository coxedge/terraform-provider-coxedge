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

func dataSourceBareMetalSSHKeys() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBareMetalSSHKeysRead,
		Schema:      getBareMetalSSHKeysSetSchema(),
	}
}

func dataSourceBareMetalSSHKeysRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)
	resourceId := d.Get("id").(string)

	if resourceId != "" {
		bareMetalSSHKeys, err := coxEdgeClient.GetBareMetalSSHKeyById(environmentName, organizationId, resourceId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("baremetal_ssh_keys", flattenBareMetalSSHKeysData(&[]apiclient.BareMetalSSHKey{*bareMetalSSHKeys})); err != nil {
			return diag.FromErr(err)
		}
	} else {
		bareMetalSSHKeys, err := coxEdgeClient.GetBareMetalSSHKeys(environmentName, organizationId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("baremetal_ssh_keys", flattenBareMetalSSHKeysData(&bareMetalSSHKeys)); err != nil {
			return diag.FromErr(err)
		}
	}
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenBareMetalSSHKeysData(bareMetalSSHKeys *[]apiclient.BareMetalSSHKey) []interface{} {
	if bareMetalSSHKeys != nil {
		devices := make([]interface{}, len(*bareMetalSSHKeys), len(*bareMetalSSHKeys))

		for i, sshKey := range *bareMetalSSHKeys {
			item := make(map[string]interface{})
			item["id"] = sshKey.Id
			item["public_key"] = sshKey.PublicKey
			item["name"] = sshKey.Name
			devices[i] = item
		}
		return devices
	}
	return make([]interface{}, 0)
}
