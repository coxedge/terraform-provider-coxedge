package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"time"
)

func resourceComputeReservedIPAttachDetachInstance() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeReservedIPAttachDetachInstanceCreate,
		ReadContext:   resourceComputeReservedIPAttachDetachInstanceRead,
		UpdateContext: resourceComputeReservedIPAttachDetachInstanceUpdate,
		DeleteContext: resourceComputeReservedIPAttachDetachInstanceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getResourceComputeReservedIPAttachDetachSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceComputeReservedIPAttachDetachInstanceCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	reservedIPId := data.Get("reserved_ip_id").(string)
	action := data.Get("action").(string)

	if action == "attach" {
		attachRequest := apiclient.ComputeReservedIPAttachDetachRequest{
			WorkloadId: data.Get("workload_id").(string),
		}

		attachResponse, err := coxEdgeClient.PatchAttachComputeReservedIP(attachRequest, environmentName, organizationId, reservedIPId)
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
		detachResponse, err := coxEdgeClient.PatchDetachComputeReservedIP(environmentName, organizationId, reservedIPId)
		if err != nil {
			return diag.FromErr(err)
		}

		tflog.Info(ctx, "Initiated Create. Awaiting task result.")

		timeout := data.Timeout(schema.TimeoutCreate)

		//Await
		_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, detachResponse.TaskId, timeout)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	data.SetId(reservedIPId)
	return diags
}

func resourceComputeReservedIPAttachDetachInstanceRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceComputeReservedIPAttachDetachInstanceUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceComputeReservedIPAttachDetachInstanceCreate(ctx, data, i)
	return diags
}

func resourceComputeReservedIPAttachDetachInstanceDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
