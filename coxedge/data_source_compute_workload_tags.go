package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceComputeWorkloadTags() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeWorkloadTagsRead,
		Schema:      getComputeWorkloadTagSetSchema(),
	}
}

func dataSourceComputeWorkloadTagsRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	workloadId := data.Get("workload_id").(string)

	tags, err := coxEdgeClient.GetComputeWorkloadTagById(environmentName, organizationId, workloadId)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("tags", flattenComputeWorkloadTagsData(&tags)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenComputeWorkloadTagsData(tags *[]apiclient.ComputeWorkloadTag) []interface{} {
	if tags != nil {
		tagList := make([]interface{}, len(*tags), len(*tags))

		for i, tg := range *tags {
			item := make(map[string]interface{})
			item["id"] = tg.ID
			item["tag"] = tg.Tag
			tagList[i] = item
		}
		return tagList
	}
	return make([]interface{}, 0)
}
