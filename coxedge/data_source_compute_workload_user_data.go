package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceComputeWorkloadUserData() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeWorkloadUserDataRead,
		Schema:      getComputeWorkloadUserDataSetSchema(),
	}
}

func dataSourceComputeWorkloadUserDataRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	workloadId := data.Get("workload_id").(string)

	userdata, err := coxEdgeClient.GetComputeWorkloadUserDataById(environmentName, organizationId, workloadId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("user_data", flattenComputeWorkloadUserDataData(&[]apiclient.ComputeWorkloadUserData{*userdata})); err != nil {
		return diag.FromErr(err)
	}

	// always run
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenComputeWorkloadUserDataData(userData *[]apiclient.ComputeWorkloadUserData) []interface{} {
	if userData != nil {
		osList := make([]interface{}, len(*userData), len(*userData))

		for i, ud := range *userData {
			item := map[string]interface{}{
				"id":        ud.ID,
				"user_data": ud.UserData,
			}
			osList[i] = item
		}
		return osList
	}
	return make([]interface{}, 0)
}
