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

func resourceComputeWorkloadPlan() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputeWorkloadPlanCreate,
		ReadContext:   resourceComputeWorkloadPlanRead,
		UpdateContext: resourceComputeWorkloadPlanUpdate,
		DeleteContext: resourceComputeWorkloadPlanDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getResourceComputeWorkloadPlanSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
	}
}

func resourceComputeWorkloadPlanCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceComputeWorkloadPlanUpdate(ctx, data, i)
	return diags
}

func resourceComputeWorkloadPlanRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//check the id comes with id & environment_name, then split the value -> in case of importing the resource
	//format is <workload_id>:<environment_name>:<organization_id>
	if strings.Contains(data.Id(), ":") {
		keys := strings.Split(data.Id(), ":")
		data.SetId(keys[0])
		data.Set("workload_id", keys[0])
		data.Set("environment_name", keys[1])
		data.Set("organization_id", keys[2])
	}
	//Get the resource Id
	resourceId := data.Get("workload_id").(string)
	organizationId := data.Get("organization_id").(string)
	environmentName := data.Get("environment_name").(string)

	//Get the resource
	planResponse, err := coxEdgeClient.GetComputeWorkloadPlanById(environmentName, organizationId, resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Update state
	convertComputeWorkloadPlanAPIObjectToResourceData(data, planResponse)

	return diags
}

func resourceComputeWorkloadPlanUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := i.(apiclient.Client)

	var diags diag.Diagnostics
	resourceId := data.Get("workload_id").(string)
	organizationId := data.Get("organization_id").(string)
	environmentName := data.Get("environment_name").(string)
	planRequest := apiclient.ComputeWorkloadPlanRequest{
		SelectedPlanId: data.Get("selected_plan_id").(string),
	}
	//Call the API
	planResponse, err := coxEdgeClient.UpdateComputeWorkloadPlan(planRequest, environmentName, organizationId, resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Info(ctx, "Initiated Update. Awaiting task result.")

	timeout := data.Timeout(schema.TimeoutCreate)
	//Await
	_, err = coxEdgeClient.AwaitTaskResolveWithCustomTimeout(ctx, planResponse.TaskId, timeout)
	if err != nil {
		return diag.FromErr(err)
	}

	//Save the Id
	data.SetId(resourceId)
	return diags
}

func resourceComputeWorkloadPlanDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func convertComputeWorkloadPlanAPIObjectToResourceData(d *schema.ResourceData, plan *apiclient.WorkloadPlan) {
	d.Set("id", plan.ID)
	d.Set("plan_id", plan.PlanID)
	d.Set("region", plan.Region)
	d.Set("server", plan.Server)
	d.Set("plan_label", plan.PlanLabel)
	d.Set("vcpu_count", plan.VCPUCount)
}
