package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
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
	var diags diag.Diagnostics
	return diags
}

func resourceComputeStorageUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceComputeStorageDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
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
