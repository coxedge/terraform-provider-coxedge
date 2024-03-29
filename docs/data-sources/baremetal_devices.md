---

# generated by https://github.com/hashicorp/terraform-plugin-docs

page_title: "coxedge_baremetal_devices Data Source - terraform-provider-coxedge"
subcategory: ""
description: Retrieve a list of all devices in a given environment.
  
---

# coxedge_baremetal_devices (Data Source)

Retrieve a list of all devices in a given environment.

Example Usage
---

```
data "coxedge_baremetal_devices" "test" {
  environment_name = "<environment_name>"
  organization_id  = "<organization_id>"
  id               = "<device_id>"
}

output "testing" {
  value = data.coxedge_baremetal_devices.test
}
```

<!-- schema generated by tfplugindocs -->

## Schema

### Required

- `environment_name` (String) - Name of the environment belonging to the organization.
- `organization_id` (String) - The id of the organization.

### Optional

- `id` (String) - The id of the device to retrieve

### Read-Only

- `id` (String) - The ID of this resource.
- `baremetal_devices` (List of Object) (see [below for nested schema](#nestedatt--baremetal_devices)) - The baremetal
  devices of the environment.

<a id="nestedatt--baremetal_devices"></a>

### Nested Schema for `baremetal_devices`

### Read-Only:

- `id` (String) - The unique ID of the device.
- `environment_name` (String) - The name of the environment.
- `organization_id` (String) - The id of the organization.
- `service_plan` (String) - The unique ID of the service associated with this device.
- `name` (String) - User given custom name.
- `hostname` (String) - A FQDN for the device. For example: example.hivelocity.net.
- `device_type` (String) - Generic description of device. Usually type and rack unit size.
- `primary_ip` (String) - The first assigned public IP for accessing this device.
- `status` (String) - active/inactive
- `monitors_total` (String) - Total # device monitors.
- `monitors_up` (String) - Number of passing device monitors.
- `ipmi_address` (String) - IP address for IPMI connection. Requires you to whitelist your current IP or be on IPMI VPN.
- `power_status` (String) - ON/OFF
- `tags` (List of String) - List of all user set device tags.
- `change_id` (String) - This property helps ensure that edit operations don’t overwrite other, more recent changes made
  to the same object. It gets updated automatically after each successful edit operation.(METALSOFT)
- `location` (Object) (see [below for nested schema](#nestedobjatt--location)) - Object of Location

<a id="nestedobjatt--baremetal_devices--location"></a>

### Nested Schema for `location`

Read-Only:

- `facility` (String) - A facility code. For example NYC1.
- `facility_title` (String) - A facility name.
