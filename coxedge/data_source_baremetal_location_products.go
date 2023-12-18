package coxedge

import (
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

func dataSourceBareMetalLocationProducts() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBareMetalLocationProductsRead,
		Schema:      getBareMetalLocationProductsSetSchema(),
	}
}

func dataSourceBareMetalLocationProductsRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	coxEdgeClient := i.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	environmentName := data.Get("environment_name").(string)
	organizationId := data.Get("organization_id").(string)
	id := data.Get("id").(string)
	code := data.Get("code").(string)

	bareMetalProducts, err := coxEdgeClient.GetBareMetalLocationProducts(environmentName, organizationId, id, code)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("baremetal_products", flattenBareMetalLocationProductsData(&bareMetalProducts)); err != nil {
		return diag.FromErr(err)
	}

	data.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}

func flattenBareMetalLocationProductsData(products *[]apiclient.BareMetalLocationProduct) []interface{} {
	if products != nil {
		flattenedProducts := make([]interface{}, len(*products), len(*products))

		for i, product := range *products {
			flattenedProduct := make(map[string]interface{})

			flattenedProduct["id"] = product.ID
			flattenedProduct["drive"] = product.Drive
			flattenedProduct["cpu"] = product.CPU
			flattenedProduct["sub_title"] = product.SubTitle
			flattenedProduct["memory"] = product.Memory
			flattenedProduct["bandwidth"] = product.Bandwidth
			flattenedProduct["monthly_price"] = product.MonthlyPrice
			flattenedProduct["monthly_premium"] = product.MonthlyPremium
			flattenedProduct["stock"] = product.Stock
			flattenedProduct["cpu_cores"] = product.CPUCores
			flattenedProduct["gpu"] = product.GPU
			flattenedProduct["hourly_price"] = product.HourlyPrice
			flattenedProduct["hourly_premium"] = product.HourlyPremium
			flattenedProduct["vendor_product_id"] = product.VendorProductID

			// Flatten ProductProcessorInfo
			flattenedProduct["cores"] = product.ProcessorInfo.Cores
			flattenedProduct["sockets"] = product.ProcessorInfo.Sockets
			flattenedProduct["threads"] = product.ProcessorInfo.Threads
			flattenedProduct["vcpus"] = product.ProcessorInfo.VCPUs

			flattenedProducts[i] = flattenedProduct
		}
		return flattenedProducts
	}
	return make([]interface{}, 0)
}
