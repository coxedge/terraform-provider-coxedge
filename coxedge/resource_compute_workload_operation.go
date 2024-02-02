package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"time"
)

func resourceComputeWorkloadOperation() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeWorkloadOperationCreate,
		ReadContext:   resourceComputeWorkloadOperationRead,
		UpdateContext: resourceComputeWorkloadOperationUpdate,
		DeleteContext: resourceComputeWorkloadOperationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getResourceComputeWorkloadPowerSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceComputeWorkloadOperationCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceComputeWorkloadOperationUpdate(ctx, data, i)
	return diags
}

func resourceComputeWorkloadOperationRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceComputeWorkloadOperationUpdate(ctx, data, i)
	return diags
}

func resourceComputeWorkloadOperationUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)
	var diags diag.Diagnostics
	resourceId := data.Get("workload_id").(string)
	organizationId := data.Get("organization_id").(string)
	environmentName := data.Get("environment_name").(string)
	operation := data.Get("operation").(string)

	operationResponse, err := coxEdgeClient.UpdateComputeWorkloadOperation(environmentName, organizationId, resourceId, operation)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Update. Awaiting task result.")

	timeout := data.Timeout(schema.TimeoutCreate)
	//Await
	_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, operationResponse.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(resourceId)
	return diags
}

func resourceComputeWorkloadOperationDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	return diag.Errorf("Unfortunately, this operation is not available")
}
