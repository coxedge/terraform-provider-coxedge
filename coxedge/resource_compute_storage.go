package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strings"
	"time"
)

func resourceComputeStorage() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeStorageCreate,
		ReadContext:   resourceComputeStorageRead,
		UpdateContext: resourceComputeStorageUpdate,
		DeleteContext: resourceComputeStorageDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getResourceComputeStorageSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceComputeStorageCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//convert resource data to API object
	storageRequest := convertResourceDataToComputeStorageCreateAPIObject(data)

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)

	//Call the API
	createdStorage, err := coxEdgeClient.CreateComputeStorage(storageRequest, environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Create. Awaiting task result.")

	timeout := data.Timeout(schema.TimeoutCreate)
	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, createdStorage.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the Id
	data.SetId(taskResult.Data.TaskId)
	return diags
}

func resourceComputeStorageRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//check the id comes with id & environment_name, then split the value -> in case of importing the resource
	//format is <storage_id>:<environment_name>:<organization_id>
	if strings.Contains(data.Id(), ":") {
		keys := strings.Split(data.Id(), ":")
		data.Set("storage_id", keys[0])
		data.Set("environment_name", keys[1])
		data.Set("organization_id", keys[2])
	}

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	storageId := data.Get("storage_id").(string)

	computeStorage, err := coxEdgeClient.GetComputeStorageById(environmentName, organizationId, storageId)
	if err != nil {
		return diag.FromErr(err)
	}
	convertStorageAPIDataToResourceData(data, computeStorage)
	data.SetId(storageId)
	return diags
}

func resourceComputeStorageUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	storageId := data.Get("storage_id").(string)

	tflog.Info(ctx, "Initiated Update. Awaiting task result.")

	if data.HasChange("size_gb") {
		sizeRequest := apiclient.UpdateComputeStorageSizeRequest{
			SizeGB: data.Get("size_gb").(string),
		}
		//Call the API
		updateStorageSize, err := coxEdgeClient.UpdateComputeStorageSize(sizeRequest, environmentName, organizationId, storageId)
		if err != nil {
			return diag.FromErr(err)
		}
		timeout := data.Timeout(schema.TimeoutUpdate)

		//Await
		_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, updateStorageSize.TaskId, timeout)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if data.HasChange("label") {
		labelRequest := apiclient.UpdateComputeStorageLabelRequest{
			Label: data.Get("label").(string),
		}
		//Call the API
		updateStorageLabel, err := coxEdgeClient.UpdateComputeStorageLabel(labelRequest, environmentName, organizationId, storageId)
		if err != nil {
			return diag.FromErr(err)
		}
		timeout := data.Timeout(schema.TimeoutUpdate)

		//Await
		_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, updateStorageLabel.TaskId, timeout)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	//Save the Id
	data.SetId(storageId)
	return diags
}

func resourceComputeStorageDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	//Get the resource Id
	resourceId := data.Id()
	organizationId := data.Get("organization_id").(string)
	environmentName := data.Get("environment_name").(string)

	//Delete the Workload
	err := coxEdgeClient.DeleteComputeStorageById(environmentName, organizationId, resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")

	return diags
}

func convertResourceDataToComputeStorageCreateAPIObject(data *schema.ResourceData) apiclient.ComputeStorageRequest {
	storageRequest := apiclient.ComputeStorageRequest{
		Region:    data.Get("region").(string),
		SizeGB:    data.Get("size_gb").(string),
		Label:     data.Get("label").(string),
		BlockType: "storage_opt",
	}
	return storageRequest
}

func convertStorageAPIDataToResourceData(d *schema.ResourceData, storage *apiclient.ComputeStorage) {
	d.Set("id", storage.ID)
	d.Set("date_created", storage.DateCreated)
	d.Set("cost", storage.Cost)
	d.Set("status", storage.Status)
	d.Set("size_gb", storage.SizeGB)
	d.Set("region", storage.Region)
	d.Set("attached_to_instance", storage.AttachedToInstance)
	d.Set("label", storage.Label)
	d.Set("mount_id", storage.MountID)
	d.Set("block_type", storage.BlockType)
	d.Set("description", storage.Description)
	d.Set("type", storage.Type)
	d.Set("location", storage.Location)
	d.Set("attached_to", storage.AttachedTo)
	d.Set("manage_label", storage.ManageLabel)
	d.Set("price", storage.Price)
	d.Set("size_in_gb", storage.SizeInGB)
	d.Set("edit_block_storage_label", storage.EditBlockStorageLabel)
	d.Set("none", storage.None)
}
