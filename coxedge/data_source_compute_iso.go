package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceComputeISO() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceComputeISORead,
		Schema:      getComputeISOSetSchema(),
	}
}

func dataSourceComputeISORead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	isoId := data.Get("iso_id").(string)

	if isoId != "" {
		computeReservedIP, err := coxEdgeClient.GetComputeISOById(environmentName, organizationId, isoId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("isos", flattenComputeISOData(&[]apiclient.ComputeISO{*computeReservedIP})); err != nil {
			return diag.FromErr(err)
		}
	} else {
		computeReservedIP, err := coxEdgeClient.GetComputeISOs(environmentName, organizationId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := data.Set("isos", flattenComputeISOData(&computeReservedIP)); err != nil {
			return diag.FromErr(err)
		}
	}
	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenComputeISOData(dataSlice *[]apiclient.ComputeISO) []interface{} {
	if dataSlice != nil {
		flattened := make([]interface{}, len(*dataSlice))

		for i, instance := range *dataSlice {
			item := map[string]interface{}{
				"id":           instance.ID,
				"prefix_id":    instance.PrefixID,
				"date_created": instance.DateCreated,
				"filename":     instance.Filename,
				"size":         instance.Size,
				"md5sum":       instance.MD5Sum,
				"sha512sum":    instance.SHA512Sum,
				"status":       instance.Status,
			}

			flattened[i] = item
		}
		return flattened
	}
	return make([]interface{}, 0)
}
