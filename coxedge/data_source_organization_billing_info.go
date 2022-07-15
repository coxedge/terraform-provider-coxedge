package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

func dataSourceOrganizationBillingInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOrganizationBillingInfoRead,
		Schema:      getOrganizationBillingInfoSetSchema(),
	}
}

func dataSourceOrganizationBillingInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	requestedId := d.Get("id").(string)
	if requestedId != "" {
		org, err := coxEdgeClient.GetOrganizationBillingInfo(requestedId)
		if err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("organizations_billing_info", flattenOrganizationBillingInfoData(&[]apiclient.OrganizationBillingInfo{*org})); err != nil {
			return diag.FromErr(err)
		}
	} else {
		return diags
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenOrganizationBillingInfoData(organizations *[]apiclient.OrganizationBillingInfo) []interface{} {
	if organizations != nil {
		orgs := make([]interface{}, len(*organizations), len(*organizations))

		for i, org := range *organizations {
			item := make(map[string]interface{})
			item["id"] = org.Id
			item["organization_id"] = org.Organization.Id
			item["billing_provider_id"] = org.BillingProvider.Id
			item["card_type"] = org.CardType
			item["card_masked_number"] = org.CardMaskedNumber
			item["card_name"] = org.CardName
			item["card_exp"] = org.CardExp
			item["billing_address_line_one"] = org.BillingAddressLineOne
			item["billing_address_line_two"] = org.BillingAddressLineTwo
			item["billing_address_city"] = org.BillingAddressCity
			item["billing_address_province"] = org.BillingAddressProvince
			item["billing_address_postal_code"] = org.BillingAddressPostalCode
			item["billing_address_postal_country"] = org.BillingAddressCountry
			orgs[i] = item
		}

		return orgs
	}

	return make([]interface{}, 0)
}
