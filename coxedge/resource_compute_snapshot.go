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

func resourceComputeSnapshots() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeSnapshotsCreate,
		ReadContext:   resourceComputeSnapshotsRead,
		UpdateContext: resourceComputeSnapshotsUpdate,
		DeleteContext: resourceComputeSnapshotsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getResourceComputeSnapshotSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceComputeSnapshotsCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	request := apiclient.ComputeSnapshotRequest{
		InstanceId:  data.Get("instance_id").(string),
		Description: data.Get("description").(string),
	}

	existingList, err := coxEdgeClient.GetComputeSnapshots(environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}
	existingIDs := make(map[string]bool)
	for _, item := range existingList {
		existingIDs[item.ID] = true
	}

	//Call the API
	snapshotResponse, err := coxEdgeClient.CreateComputeSnapshot(request, environmentName, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Create. Awaiting task result.")

	timeout := data.Timeout(schema.TimeoutCreate)
	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, snapshotResponse.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	if taskResult.Data.TaskStatus == "SUCCESS" {
		afterList, err := coxEdgeClient.GetComputeSnapshots(environmentName, organizationId)
		if err != nil {
			return diag.FromErr(err)
		}
		var missingItem *apiclient.ComputeSnapshot
		for _, item := range afterList {
			if !existingIDs[item.ID] {
				missingItem = &item
				data.SetId(missingItem.ID)
				return diags
			}
		}
	}
	//Save the Id
	data.SetId(taskResult.Data.TaskId)
	return diags
}

func resourceComputeSnapshotsRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//check the id comes with id & environment_name, then split the value -> in case of importing the resource
	//format is <snapshot_id>:<environment_name>:<organization_id>
	if strings.Contains(data.Id(), ":") {
		keys := strings.Split(data.Id(), ":")
		data.SetId(keys[0])
		data.Set("environment_name", keys[2])
		data.Set("organization_id", keys[3])
	}

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	snapshotId := data.Id()

	computeIso, err := coxEdgeClient.GetComputeSnapshotById(environmentName, organizationId, snapshotId)
	if err != nil {
		return diag.FromErr(err)
	}
	convertSnapshotToResourceData(data, computeIso)
	data.SetId(snapshotId)
	return diags
}

func resourceComputeSnapshotsUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceComputeSnapshotsDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	//Get the resource Id
	resourceId := data.Id()
	organizationId := data.Get("organization_id").(string)
	environmentName := data.Get("environment_name").(string)

	//Delete the Storage
	err := coxEdgeClient.DeleteComputeSnapshotById(environmentName, organizationId, resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	// From Docs: d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	data.SetId("")

	return diags
}

func convertSnapshotToResourceData(d *schema.ResourceData, snapshot *apiclient.ComputeSnapshot) {
	d.Set("id", snapshot.ID)
	d.Set("prefix_id", snapshot.PrefixID)
	d.Set("date_created", snapshot.DateCreated)
	d.Set("description", snapshot.Description)
	d.Set("size", snapshot.Size)
	d.Set("compressed_size", snapshot.CompressedSize)
	d.Set("os_id", snapshot.OSID)
	d.Set("app_id", snapshot.AppID)
	d.Set("status", snapshot.Status)
}