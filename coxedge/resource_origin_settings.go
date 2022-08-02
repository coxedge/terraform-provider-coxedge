/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"coxedge/terraform-provider/coxedge/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"strings"
	"time"
)

func resourceOriginSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOriginSettingsCreate,
		ReadContext:   resourceOriginSettingsRead,
		UpdateContext: resourceOriginSettingsUpdate,
		DeleteContext: resourceOriginSettingsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getOriginSettingsSchema(),
	}
}

func resourceOriginSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Convert resource data to API object
	updatedOriginSettings := convertResourceDataToOriginSettingsCreateAPIObject(d)
	d.SetId(updatedOriginSettings.Id)

	//Run Update since you do not "create" these
	resourceOriginSettingsUpdate(ctx, d, m)

	return diags
}

func resourceOriginSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//check the id comes with id & environment_name, then split the value -> in case of importing the resource
	//format is <site_id>:<environment_name>:<organization_id>
	if strings.Contains(d.Id(), ":") {
		keys := strings.Split(d.Id(), ":")
		d.SetId(keys[0])
		d.Set("environment_name", keys[1])
		d.Set("organization_id", keys[2])
	}

	//Get the resource ID
	resourceId := d.Id()
	organizationId := d.Get("organization_id").(string)

	//Get the resource
	originSettings, err := coxEdgeClient.GetOriginSettings(d.Get("environment_name").(string), resourceId, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}
	convertOriginSettingsAPIObjectToResourceData(d, originSettings)

	return diags
}

func resourceOriginSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource ID
	resourceId := d.Id()

	//Convert resource data to API object
	updatedOriginSettings := convertResourceDataToOriginSettingsCreateAPIObject(d)
	organizationId := d.Get("organization_id").(string)
	//Call the API
	_, err := coxEdgeClient.UpdateOriginSettings(resourceId, updatedOriginSettings, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Set last_updated
	d.Set("last_updated", time.Now().Format(time.RFC850))

	return resourceOriginSettingsRead(ctx, d, m)
}

func resourceOriginSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource ID
	resourceId := d.Id()

	//Delete the OriginSettings
	err := coxEdgeClient.DeleteOriginSettings(d.Get("environment_name").(string), resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func convertResourceDataToOriginSettingsCreateAPIObject(d *schema.ResourceData) apiclient.OriginSettings {
	//Create update originSettings struct
	updatedOriginSettings := apiclient.OriginSettings{
		EnvironmentName:      d.Get("environment_name").(string),
		Id:                   d.Get("site_id").(string),
		StackId:              d.Get("stack_id").(string),
		ScopeConfigurationId: d.Get("scope_configuration_id").(string),
		Domain:               d.Get("domain").(string),
		PullProtocol:         d.Get("pull_protocol").(string),
		HostHeader:           d.Get("host_header").(string),
	}

	webSocketEnabled := d.Get("websockets_enabled").(string)
	if webSocketEnabled != "" {
		boolValue, _ := strconv.ParseBool(webSocketEnabled)
		updatedOriginSettings.WebSocketsEnabled = utils.BoolAddr(boolValue)
	}

	ssValidationEnabled := d.Get("ssl_validation_enabled").(string)
	if ssValidationEnabled != "" {
		boolValue, _ := strconv.ParseBool(ssValidationEnabled)
		updatedOriginSettings.SSLValidationEnabled = utils.BoolAddr(boolValue)
	}

	backupOriginEnabledValue := false
	backupOriginEnabled := d.Get("backup_origin_enabled").(string)
	if backupOriginEnabled != "" {
		boolValue, _ := strconv.ParseBool(backupOriginEnabled)
		backupOriginEnabledValue = boolValue
		updatedOriginSettings.BackupOriginEnabled = utils.BoolAddr(boolValue)
	}

	//Convert Backup Origin Codes
	updatedOriginSettings.BackupOriginExcludeCodes = []string{}
	for _, excludeCode := range d.Get("backup_origin_exclude_codes").([]interface{}) {
		updatedOriginSettings.BackupOriginExcludeCodes = append(updatedOriginSettings.BackupOriginExcludeCodes, excludeCode.(string))
	}

	//Convert origin
	for _, originSpecRaw := range d.Get("origin").([]interface{}) {
		originSpec := originSpecRaw.(map[string]interface{})
		origin := apiclient.OriginSettingsOrigin{
			Id:                    originSpec["id"].(string),
			Address:               originSpec["address"].(string),
			AuthMethod:            originSpec["auth_method"].(string),
			Username:              originSpec["username"].(string),
			Password:              originSpec["password"].(string),
			CommonCertificateName: originSpec["common_certificate_name"].(string),
		}
		updatedOriginSettings.Origin = origin
	}

	if backupOriginEnabledValue {
		//Convert origin
		for _, originSpecRaw := range d.Get("backup_origin").([]interface{}) {
			originSpec := originSpecRaw.(map[string]interface{})
			origin := apiclient.OriginSettingsOrigin{
				Id:                    originSpec["id"].(string),
				Address:               originSpec["address"].(string),
				CommonCertificateName: originSpec["common_certificate_name"].(string),
			}
			updatedOriginSettings.BackupOrigin = origin
		}
	}

	return updatedOriginSettings
}

func convertOriginSettingsAPIObjectToResourceData(d *schema.ResourceData, originSettings *apiclient.OriginSettings) {
	d.Set("id", originSettings.Id)
	d.Set("stack_id", originSettings.StackId)
	d.Set("scope_configuration_id", originSettings.ScopeConfigurationId)
	d.Set("domain", originSettings.Domain)
	d.Set("websockets_enabled", strconv.FormatBool(*originSettings.WebSocketsEnabled))
	d.Set("ssl_validation_enabled", strconv.FormatBool(*originSettings.SSLValidationEnabled))
	d.Set("pull_protocol", originSettings.PullProtocol)
	d.Set("host_header", originSettings.HostHeader)
	origin := make([]map[string]string, 1)
	origin[0] = make(map[string]string)
	origin[0]["id"] = originSettings.Origin.Id
	origin[0]["address"] = originSettings.Origin.Address
	origin[0]["auth_method"] = originSettings.Origin.AuthMethod
	origin[0]["username"] = originSettings.Origin.Username
	origin[0]["password"] = originSettings.Origin.Password
	origin[0]["common_certificate_name"] = originSettings.Origin.CommonCertificateName
	d.Set("origin", origin)
	d.Set("backup_origin_enabled", strconv.FormatBool(*originSettings.BackupOriginEnabled))
	d.Set("backup_origin_exclude_codes", originSettings.BackupOriginExcludeCodes)
}
