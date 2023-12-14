package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
)

func dataSourceBareMetalDeviceDisks() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBareMetalDeviceDisksRead,
		Schema:      getBareMetalDeviceDiskSetSchema(),
	}
}

func dataSourceBareMetalDeviceDisksRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)

	requestedId := data.Get("id").(string)

	bareMetalDeviceDisks, err := coxEdgeClient.GetBareMetalDeviceDisksById(environmentName, organizationId, requestedId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("baremetal_device_disks", flattenBareMetalDeviceDisksData(&bareMetalDeviceDisks)); err != nil {
		return diag.FromErr(err)
	}

	data.SetId(requestedId)
	return diags
}

func flattenBareMetalDeviceDisksData(bareMetalDeviceDisk *[]apiclient.BareMetalDeviceDisk) []interface{} {
	if bareMetalDeviceDisk != nil {
		disks := make([]interface{}, len(*bareMetalDeviceDisk), len(*bareMetalDeviceDisk))

		for i, disk := range *bareMetalDeviceDisk {
			item := make(map[string]interface{})
			item["server_disk_id"] = disk.ServerDiskID
			item["server_disk_model"] = disk.ServerDiskModel
			item["server_disk_size_gb"] = disk.ServerDiskSizeGB
			item["server_id"] = disk.ServerID
			item["server_disk_serial"] = disk.ServerDiskSerial
			item["server_disk_vendor"] = disk.ServerDiskVendor
			item["server_disk_status"] = disk.ServerDiskStatus
			item["server_disk_type"] = disk.ServerDiskType
			item["server_raid_controller_id"] = disk.ServerRaidControllerID
			item["type"] = disk.Type
			disks[i] = item
		}
		return disks
	}
	return make([]interface{}, 0)
}
