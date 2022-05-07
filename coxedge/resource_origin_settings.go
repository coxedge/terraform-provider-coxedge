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
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Convert resource data to API Object
	newOriginSettings := convertResourceDataToOriginSettingsCreateAPIObject(d)

	//Call the API
	createdOriginSettings, err := coxEdgeClient.CreateOriginSettings(newOriginSettings)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the ID
	d.SetId(createdOriginSettings.Id)

	return diags
}

func resourceOriginSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the resource ID
	resourceId := d.Id()

	//Get the resource
	originSettings, err := coxEdgeClient.GetOriginSettings(d.Get("environment_name").(string), resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	convertOriginSettingsAPIObjectToResourceData(d, originSettings)

	//Update state
	resourceOriginSettingsRead(ctx, d, m)

	return diags
}

func resourceOriginSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource ID
	resourceId := d.Id()

	//Convert resource data to API object
	updatedOriginSettings := convertResourceDataToOriginSettingsCreateAPIObject(d)

	//Call the API
	_, err := coxEdgeClient.UpdateOriginSettings(resourceId, updatedOriginSettings)
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
		EnvironmentName:          d.Get("environment_name").(string),
		Id:                       d.Get("id").(string),
		StackId:                  d.Get("stack_id").(string),
		ScopeConfigurationId:     d.Get("scope_configuration_id").(string),
		Domain:                   d.Get("domain").(string),
		WebSocketsEnabled:        d.Get("websockets_enabled").(bool),
		SSLValidationEnabled:     d.Get("ssl_validation_enabled").(bool),
		PullProtocol:             d.Get("pull_protocol").(string),
		HostHeader:               d.Get("host_header").(string),
		BackupOriginEnabled:      d.Get("backup_origin_enabled").(bool),
		BackupOriginExcludeCodes: d.Get("backup_origin_exclude_codes").([]string),
	}

	//Convert origin
	for _, originSpec := range d.Get("origin").([]map[string]string) {
		origin := apiclient.OriginSettingsOrigin{
			Id:                    originSpec["id"],
			Address:               originSpec["address"],
			AuthMethod:            originSpec["auth_method"],
			Username:              originSpec["username"],
			Password:              originSpec["password"],
			CommonCertificateName: originSpec["common_certificate_name"],
		}
		updatedOriginSettings.Origin = origin
	}

	if updatedOriginSettings.BackupOriginEnabled {
		//Convert origin
		for _, originSpec := range d.Get("backup_origin").([]map[string]string) {
			origin := apiclient.OriginSettingsOrigin{
				Id:                    originSpec["id"],
				Address:               originSpec["address"],
				AuthMethod:            originSpec["auth_method"],
				Username:              originSpec["username"],
				Password:              originSpec["password"],
				CommonCertificateName: originSpec["common_certificate_name"],
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
	d.Set("environment_name", originSettings.EnvironmentName)
	d.Set("domain", originSettings.Domain)
	d.Set("websockets_enabled", originSettings.WebSocketsEnabled)
	d.Set("ssl_validation_enabled", originSettings.SSLValidationEnabled)
	d.Set("pull_protocol", originSettings.PullProtocol)
	d.Set("host_header", originSettings.HostHeader)
	d.Set("origin", originSettings.Origin)
	d.Set("backup_origin_enabled", originSettings.BackupOriginEnabled)
	d.Set("backup_origin_exclude_codes", originSettings.BackupOriginExcludeCodes)
}
