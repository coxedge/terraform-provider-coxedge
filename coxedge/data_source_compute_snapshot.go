package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceComputeSnapshots() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeSnapshotsRead,
		Schema:      getComputeSnapshotSetSchema(),
	}
}

func dataSourceComputeSnapshotsRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	snapshotId := data.Get("snapshot_id").(string)

	if snapshotId != "" {
		computeSnapshot, err := coxEdgeClient.GetComputeSnapshotById(environmentName, organizationId, snapshotId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("snapshots", flattenComputeSnapshotData(&[]apiclient.ComputeSnapshot{*computeSnapshot})); err != nil {
			return diag.FromErr(err)
		}
	} else {
		computeSnapshot, err := coxEdgeClient.GetComputeSnapshots(environmentName, organizationId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("snapshots", flattenComputeSnapshotData(&computeSnapshot)); err != nil {
			return diag.FromErr(err)
		}
	}
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenComputeSnapshotData(dataSlice *[]apiclient.ComputeSnapshot) []interface{} {
	if dataSlice != nil {
		flattened := make([]interface{}, len(*dataSlice))

		for i, instance := range *dataSlice {
			item := map[string]interface{}{
				"id":              instance.ID,
				"prefix_id":       instance.PrefixID,
				"date_created":    instance.DateCreated,
				"description":     instance.Description,
				"size":            instance.Size,
				"compressed_size": instance.CompressedSize,
				"status":          instance.Status,
				"os_id":           instance.OSID,
				"app_id":          instance.AppID,
			}

			flattened[i] = item
		}
		return flattened
	}
	return make([]interface{}, 0)
}
