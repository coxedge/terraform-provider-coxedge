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

func dataSourceOriginSetting() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOriginSettingRead,
		Schema:      getOriginSettingSetSchema(),
	}
}

func dataSourceOriginSettingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	requestedId := d.Get("id").(string)
	requestedEnv := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)
	if requestedId != "" && requestedEnv != "" {
		org, err := coxEdgeClient.GetOriginSettings(requestedEnv, requestedId, organizationId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("origin_settings", flattenOriginSettingsData(&[]apiclient.OriginSettings{*org})); err != nil {
			return diag.FromErr(err)
		}
		// always run
		d.SetId(requestedId)
	} else {
		return diags
	}
	return diags
}

func flattenOriginSettingsData(origin *[]apiclient.OriginSettings) []interface{} {
	if origin != nil {
		orgs := make([]interface{}, len(*origin), len(*origin))

		for i, org := range *origin {
			item := make(map[string]interface{})

			item["id"] = org.Id
			item["stack_id"] = org.StackId
			item["scope_configuration_id"] = org.ScopeConfigurationId
			item["domain"] = org.Domain
			item["websockets_enabled"] = org.WebSocketsEnabled
			item["ssl_validation_enabled"] = org.SSLValidationEnabled
			item["pull_protocol"] = org.PullProtocol
			item["host_header"] = org.HostHeader
			item["backup_origin_enabled"] = org.BackupOriginEnabled
			item["backup_origin_exclude_codes"] = org.BackupOriginExcludeCodes

			orgBack := make([]map[string]string, 1)
			orgBack[0] = make(map[string]string)
			orgBack[0]["id"] = org.BackupOrigin.Id
			orgBack[0]["address"] = org.BackupOrigin.Address
			orgBack[0]["auth_method"] = org.BackupOrigin.AuthMethod
			item["backup_origin"] = orgBack

			orgi := make([]map[string]string, 1)
			orgi[0] = make(map[string]string)
			orgi[0]["id"] = org.Origin.Id
			orgi[0]["address"] = org.Origin.Address
			orgi[0]["auth_method"] = org.Origin.AuthMethod
			orgi[0]["username"] = org.Origin.Username
			orgi[0]["password"] = org.Origin.Password
			orgi[0]["common_certificate_name"] = org.Origin.CommonCertificateName
			item["origin"] = orgi

			orgs[i] = item
		}
		return orgs
	}

	return make([]interface{}, 0)
}
