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
	"time"
)

func resourceBareMetalDevices() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBareMetalDevicesCreate,
		ReadContext:   resourceBareMetalDevicesRead,
		UpdateContext: resourceBareMetalDevicesUpdate,
		DeleteContext: resourceBareMetalDevicesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getBareMetalDeviceResourceSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceBareMetalDevicesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	sshKeyId := d.Get("ssh_key_id").(string)
	hasSshKey := d.Get("has_ssh_data").(bool)
	if sshKeyId != "" {
		if hasSshKey {
			diag := diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Set 'has_ssh_data' to false when adding 'ssh_key_id' field to avoid issues.",
				Detail:   "When the ssh_key_id field is added, the has_ssh_data field should be set to false. Please ensure that both fields are configured correctly to avoid unexpected behavior.",
			}
			diags = append(diags, diag)
			return diags
		}
	}
	hasUserData := d.Get("has_user_data").(bool)
	userData := d.Get("user_data").(string)

	if hasUserData {
		if userData == "" {
			diag := diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Ensure 'user_data' field is configured when 'has_user_data' is true to avoid issues.",
				Detail:   "The 'user_data' field is required when 'has_user_data' is set to true. Please ensure that the 'user_data' field is configured correctly to avoid unexpected behavior.",
			}
			diags = append(diags, diag)
			return diags
		}
	}

	createRequest := convertResourceDataToBareMetalDeviceCreateAPIObject(d)
	environmentName := d.Get("environment_name").(string)
	organizationId := d.Get("organization_id").(string)

	//Call the API
	createdDevice, err := coxEdgeClient.CreateBareMetalDevice(createRequest, environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Create. Awaiting task result.")

	timeout := d.Timeout(schema.TimeoutCreate)
	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, createdDevice.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the Id
	d.SetId(taskResult.Data.Result.Id)

	return diags
}

func resourceBareMetalDevicesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Errorf("Unfortunately, it is not possible to update BareMetal devices from this resource. For guidance on updating devices, please refer to the documentation.")
}

func resourceBareMetalDevicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceBareMetalDevicesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Errorf("Unfortunately, it is not possible to delete BareMetal devices from this resource. For guidance on deleting devices, please refer to the documentation.")
}

func convertResourceDataToBareMetalDeviceCreateAPIObject(d *schema.ResourceData) apiclient.CreateBareMetalDeviceRequest {
	//Create/update BareMetal device struct

	hasUserData := d.Get("has_user_data").(bool)
	hasSshData := d.Get("has_ssh_data").(bool)
	bareMetalDevice := apiclient.CreateBareMetalDeviceRequest{
		LocationName:    d.Get("location_name").(string),
		HasUserData:     &hasUserData,
		HasSshData:      &hasSshData,
		ProductOptionId: d.Get("product_option_id").(int),
		ProductId:       d.Get("product_id").(string),
		OsName:          d.Get("os_name").(string),
		SshKey:          d.Get("ssh_key").(string),
		SshKeyName:      d.Get("ssh_key_name").(string),
		SshKeyId:        d.Get("ssh_key_id").(string),
	}

	//Convert server
	serverList := d.Get("server").([]interface{})
	var name string
	for i, entry := range serverList {
		convertedEntry := entry.(map[string]interface{})
		server := apiclient.Server{
			Hostname: convertedEntry["hostname"].(string),
		}
		if i == 0 {
			name = convertedEntry["hostname"].(string)
		} else {
			name += ", " + convertedEntry["hostname"].(string)
		}
		bareMetalDevice.Server = append(bareMetalDevice.Server, server)
	}
	bareMetalDevice.Quantity = len(serverList)
	bareMetalDevice.Name = name

	if bareMetalDevice.HasUserData != nil && *bareMetalDevice.HasUserData {
		bareMetalDevice.UserData = d.Get("user_data").(string)
	}

	if bareMetalDevice.HasSshData != nil && *bareMetalDevice.HasSshData {
		bareMetalDevice.SshKey = d.Get("ssh_key").(string)
		bareMetalDevice.SshKeyName = d.Get("ssh_key_name").(string)
	}

	sshKeyId := d.Get("ssh_key_id").(string)
	if sshKeyId != "" {
		if bareMetalDevice.HasSshData != nil && !*bareMetalDevice.HasSshData {
			bareMetalDevice.SshKeyId = sshKeyId
		}
	}
	return bareMetalDevice
}
