---

# generated by https://github.com/hashicorp/terraform-plugin-docs

page_title: "coxedge_baremetal_device_charts Data Source - terraform-provider-coxedge"
subcategory: ""
description: Get device charts in a given environment.
  
---

# coxedge_baremetal_device_charts (Data Source)

Get device charts in a given environment.

Example Usage
---

```
data "coxedge_baremetal_device_charts" "charts" {
  id               = "<id>"
  environment_name = "<environment_name>"
  organization_id  = "<organization_id>"
  custom           = true
  start_date       = "<unix_epoch_time>"
  end_date         = "<unix_epoch_time>"
}

output "charts_output" {
  value = data.coxedge_baremetal_device_charts.charts
}
```

<!-- schema generated by tfplugindocs -->

## Schema

### Required

- `environment_name` (String) - Name of the environment belonging to the organization.
- `organization_id` (String) - The id of the organization.
- `id` (String) - The unique ID of device

### Optional

- `custom` (Boolean) - The id of the device to retrieve
- `start_date` (String) - Start Time of Custom Time Period. (Unix Epoch Time)
- `end_date` (String) - End Time of Custom Time Period (Unix Epoch Time)

### Read-Only

- `id` (String) - The ID of this resource.
- `baremetal_device_charts` (List of Object) (see [below for nested schema](#nestedatt--baremetal_device_charts)) - The
  chart details of the device.

<a id="nestedatt--baremetal_device_charts"></a>

### Nested Schema for `baremetal_device_charts`

### Read-Only:

- `id` (String) - The unique ID of the device.
- `filter` (String) - Return data in the given period. Day, week, month will return the previous day, week, month from now.
- `graph_image` (String) - A PNG image of bandwidth usage.
- `interfaces` (String) - The interface(s) displayed in the image.
- `switch_id` (String) - The unique ID of the switch where bandwidth data is measured.