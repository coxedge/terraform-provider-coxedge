package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"time"
)

func resourceComputeStorageAttachDetachInstance() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeStorageAttachInstanceCreate,
		ReadContext:   resourceComputeStorageAttachInstanceRead,
		UpdateContext: resourceComputeStorageAttachInstanceUpdate,
		DeleteContext: resourceComputeStorageAttachInstanceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getResourceComputeStorageAttachSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceComputeStorageAttachInstanceCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	storageId := data.Get("storage_id").(string)
	instanceId := data.Get("instance_id").(string)
	live := data.Get("live").(bool)
	action := data.Get("action").(string)

	if action == "attach" {
		attachRequest := apiclient.AttachComputeStorageInstanceRequest{
			Live:       live,
			InstanceId: instanceId,
		}

		attachResponse, err := coxEdgeClient.AttachComputeStorageInstance(attachRequest, environmentName, organizationId, storageId)
		if err != nil {
			return diag.FromErr(err)
		}

		tflog.Info(ctx, "Initiated Create. Awaiting task result.")

		timeout := data.Timeout(schema.TimeoutCreate)

		//Await
		_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, attachResponse.TaskId, timeout)
		if err != nil {
			return diag.FromErr(err)
		}
	} else {
		detachRequest := apiclient.DetachComputeStorageInstanceRequest{
			Live: live,
		}

		attachResponse, err := coxEdgeClient.DetachComputeStorageInstance(detachRequest, environmentName, organizationId, storageId)
		if err != nil {
			return diag.FromErr(err)
		}

		tflog.Info(ctx, "Initiated Create. Awaiting task result.")

		timeout := data.Timeout(schema.TimeoutCreate)

		//Await
		_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, attachResponse.TaskId, timeout)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	data.SetId(storageId)
	return diags
}

func resourceComputeStorageAttachInstanceRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceComputeStorageAttachInstanceUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceComputeStorageAttachInstanceCreate(ctx, data, i)
	return diags
}

func resourceComputeStorageAttachInstanceDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
