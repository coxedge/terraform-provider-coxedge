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

func resourceBareMetalDevice() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBareMetalDeviceCreate,
		ReadContext:   resourceBareMetalDeviceRead,
		UpdateContext: resourceBareMetalDeviceUpdate,
		DeleteContext: resourceBareMetalDeviceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getBareMetalDeviceResourceSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceBareMetalDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	sshKeyId := d.Get("ssh_key_id").(string)
	hasSshKey := d.Get("has_ssh_data").(bool)
	if sshKeyId != "" {
		if hasSshKey {
			return diag.Errorf("has_ssh_data field should be set false if ssh_key_id field is added")
		}
	}
	hasUserData := d.Get("has_user_data").(bool)
	userData := d.Get("user_data").(string)

	if hasUserData {
		if userData == "" {
			return diag.Errorf("user_data field is required if has_user_data field is set to true")
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

func resourceBareMetalDeviceUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return diag.Errorf("no option to update BareMetal devices - remove terraform state files to create new one")
}

func resourceBareMetalDeviceRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return nil
}

func resourceBareMetalDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
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
	for _, entry := range serverList {
		convertedEntry := entry.(map[string]interface{})
		server := apiclient.Server{
			Hostname: convertedEntry["hostname"].(string),
		}
		bareMetalDevice.Server = append(bareMetalDevice.Server, server)
	}
	bareMetalDevice.Quantity = len(serverList)

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
