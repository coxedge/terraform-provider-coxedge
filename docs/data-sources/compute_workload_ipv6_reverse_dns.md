---

# generated by https://github.com/hashicorp/terraform-plugin-docs

page_title: "coxedge_compute_workload_ipv6_reverse_dns Data Source - terraform-provider-coxedge"
subcategory: ""
description: Retrieve a list of reverse dns of IPv6 from workload in a given environment.
  
---

# coxedge_compute_workload_ipv6_reverse_dns (Data Source)

Retrieve a list of reverse dns of IPv6 from workload in a given environment.

Example Usage
---

```
data "coxedge_compute_workload_ipv6_reverse_dns" "ipv6_reverse_dns" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  workload_id      = "<workload_id>"
}

output "output_ipv6_reverse_dns" {
  value = data.coxedge_compute_workload_ipv6_reverse_dns.ipv6_reverse_dns
}
```

<!-- schema generated by tfplugindocs -->

## Schema

### Required

- `environment_name` (String) - Name of the environment belonging to the organization.
- `organization_id` (String) - The id of the organization.
- `workload_id` (String) - The id of the workload.

### Read-Only

- `id` (String) - The ID of this resource.
- `ipv6_reverse_dns` (List of Object) (see [below for nested schema](#nestedatt--ipv6_reverse_dns)) - The ipv6_reverse_dns from workload of the
  environment.

<a id="nestedatt--ipv6_reverse_dns"></a>

### Nested Schema for `ipv6_reverse_dns`

### Read-Only:

- `id` (String) - The unique ID of the IPv6.
- `environment_name` (String) - The name of the environment.
- `organization_id` (String) - The id of the organization.
- `ip` (String) - The IPv6 address.
- `reverse` (String) - The IPv6 reverse entry.
