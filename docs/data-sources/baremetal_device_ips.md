---

# generated by https://github.com/hashicorp/terraform-plugin-docs

page_title: "coxedge_baremetal_device_ips Data Source - terraform-provider-coxedge"
subcategory: ""
description: Get device IPs in a given environment.
  
---

# coxedge_baremetal_device_ips (Data Source)

Get device IPs in a given environment.

Example Usage
---

```
data "coxedge_baremetal_device_ips" "ips" {
  id               = "<device_id>"
  environment_name = "<environment_name"
  organization_id  = "<organization_id>"
}

output "ips_output" {
  value = data.coxedge_baremetal_device_ips.ips
}
```

<!-- schema generated by tfplugindocs -->

## Schema

### Required

- `environment_name` (String) - Name of the environment belonging to the organization.
- `organization_id` (String) - The id of the organization.
- `id` (String) - The unique ID of device

### Read-Only

- `id` (String) - The ID of this resource.
- `baremetal_device_ips` (List of Object) (see [below for nested schema](#nestedatt--baremetal_device_ips)) - The IP
  details of the device.

<a id="nestedatt--baremetal_device_ips"></a>

### Nested Schema for `baremetal_device_ips`

### Read-Only:

- `ip_name` (String) - Name of IP
- `value` (String) - Value of IP