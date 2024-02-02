package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceComputeStorages() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeStoragesRead,
		Schema:      getComputeStorageSetSchema(),
	}
}

func dataSourceComputeStoragesRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	storageId := data.Get("storage_id").(string)

	if storageId != "" {
		computeStorage, err := coxEdgeClient.GetComputeStorageById(environmentName, organizationId, storageId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("storages", flattenComputeStoragesData(&[]apiclient.ComputeStorage{*computeStorage})); err != nil {
			return diag.FromErr(err)
		}
	} else {
		computeStorages, err := coxEdgeClient.GetComputeStorages(environmentName, organizationId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("storages", flattenComputeStoragesData(&computeStorages)); err != nil {
			return diag.FromErr(err)
		}
	}
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenComputeStoragesData(computeWorkloads *[]apiclient.ComputeStorage) []interface{} {
	if computeWorkloads != nil {
		storages := make([]interface{}, len(*computeWorkloads))

		for i, storage := range *computeWorkloads {
			item := make(map[string]interface{})

			item["id"] = storage.ID
			item["date_created"] = storage.DateCreated
			item["cost"] = storage.Cost
			item["status"] = storage.Status
			item["size_gb"] = storage.SizeGB
			item["region"] = storage.Region
			item["attached_to_instance"] = storage.AttachedToInstance
			item["label"] = storage.Label
			item["mount_id"] = storage.MountID
			item["block_type"] = storage.BlockType
			item["description"] = storage.Description
			item["type"] = storage.Type
			item["location"] = storage.Location
			item["attached_to"] = storage.AttachedTo
			item["manage_label"] = storage.ManageLabel
			item["price"] = storage.Price
			item["size_in_gb"] = storage.SizeInGB
			item["edit_block_storage_label"] = storage.EditBlockStorageLabel
			item["none"] = storage.None
			item["detach"] = storage.Detach
			item["attach"] = storage.Attach

			storages[i] = item
		}
		return storages
	}
	return make([]interface{}, 0)
}
