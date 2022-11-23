---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "coxedge_organizations_billing_info Data Source - terraform-provider-coxedge"
subcategory: ""
description: Retrieve the billing information for an organization.
  
---

# coxedge_organizations_billing_info (Data Source)
Retrieve the billing information for an organization.

Example Usage
---
```
data "coxedge_organizations_billing_info" "test" {
  id = "899c1ef6-6dc5-49f6-9663-3512e12d6ewe"
}

output "testing" {
  value = data.coxedge_organizations_billing_info.test
}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `organizations_billing_info` (List of Object) (see [below for nested schema](#nestedatt--organizations_billing_info)) - Retrieve the billing information for an organization.

<a id="nestedatt--organizations_billing_info"></a>
### Nested Schema for `organizations_billing_info`

Read-Only:

- `billing_address_city` (String) - The city of the billing address.
- `billing_address_line_one` (String) - The address line 1 of the billing address.
- `billing_address_line_two` (String) - The address line 2 of the billing address.
- `billing_address_postal_code` (String) - The postal/zip code of the billing address.
- `billing_address_postal_country` (String)- The postal/zip code of the billing address.
- `billing_address_province` (String) - The province or state code (2 letters) of the billing address.
- `billing_provider_id` (String) - The billing provider associated to the credit card.
- `card_exp` (String) - The credit card expiration ('mmyy' format)
- `card_masked_number` (String) - The credit card masked number.
- `card_name` (String) - The name on the credit card
- `card_type` (String) - The credit card type.
- `id` (String) - The ID of this resource.
- `organization_id` (String) - The ID of the organization
 
