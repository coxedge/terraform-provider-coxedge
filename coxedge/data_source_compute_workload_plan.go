package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceComputeWorkloadPlan() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeWorkloadPlanRead,
		Schema:      getComputeWorkloadPlanSetSchema(),
	}
}

func dataSourceComputeWorkloadPlanRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	workloadId := data.Get("workload_id").(string)

	plan, err := coxEdgeClient.GetComputeWorkloadPlanById(environmentName, organizationId, workloadId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("plan", flattenComputeWorkloadPlanData(&[]apiclient.WorkloadPlan{*plan})); err != nil {
		return diag.FromErr(err)
	}

	// always run
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenComputeWorkloadPlanData(plan *[]apiclient.WorkloadPlan) []interface{} {
	if plan != nil {
		planList := make([]interface{}, len(*plan), len(*plan))

		for i, host := range *plan {
			item := map[string]interface{}{
				"id":         host.ID,
				"plan_id":    host.PlanID,
				"region":     host.Region,
				"server":     host.Server,
				"plan_label": host.PlanLabel,
				"vcpu_count": host.VCPUCount,
			}
			planList[i] = item
		}
		return planList
	}
	return make([]interface{}, 0)
}
