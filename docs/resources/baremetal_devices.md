---

# generated by https://github.com/hashicorp/terraform-plugin-docs

page_title: "coxedge_baremetal_devices Resource - terraform-provider-coxedge"
subcategory: ""
description: Allows you to deploy, and manage your Bare Metal machines and network resources.
  
---

# coxedge_baremetal_devices (Resource)

Allows you to deploy your Bare Metal machines and network resources.

Example Usage
---

```
//for Baremetal
resource "coxedge_baremetal_devices" "device" {
  environment_name  = "<environment_name>"
  organization_id   = "<organization_id>"
  location_code     = "ATL2"
  has_user_data     = false
  has_ssh_data      = false
  product_option_id = 144463
  product_id        = "580"
  os_name           = "Ubuntu 20.x"
  server {
    hostname = "example.coxedge.com"
  }
  timeouts {
    create = "20m"
  }
}

//for Edge Fabric
resource "coxedge_baremetal_devices" "device" {
  environment_name  = "<environment_name>"
  organization_id   = "<organization_id>"
  location_code    = "cox-dvtc"
  product_id       = "1"
  os_id            = "4"
  server_label     = "example"
  tags             = tolist(["tag added"])
}
```

<!-- schema generated by tfplugindocs -->

## Argument Reference

The following arguments are supported:

### Required

- `environment_name` (String) - The name of the environment.
- `organization_id` (String) - The id of the organization.
- `location_code` (String) - A facility code. For example NYC1.
- `has_user_data` (Boolean) - True if we're passing user data(Depends on location_code).
- `has_ssh_data` (Boolean) - True if we're passing SSH key(Depends on location_code).
- `product_option_id` (Number) - The unique ID of the desired product option(Depends on location_code).
- `product_id` (String) - The unique ID of the desired product to provision.
- `os_id` (String) - The unique ID of the operating system (Depends on location_code).
- `server_label` (String) - The name of server (Depends on location_code).
- `os_name` (String) - The name of the Operating System to provision on this device. Must match name of an operating
  system product option (Depends on location_code).
- `server` (List of Objects) (see [below for nested schema](#nestedblock--server)) - List of servers (Depends on location_code).

### Optional

- `user_data` (String) - Value of user inputs
- `ssh_key` (String) - SSH key value.
- `ssh_key_name` (String) - Name for newly adding SSH key.
- `ssh_key_id` (String) - The unique ID of the SSH key.
- `tags` (List of String) - List of tags (Depends on location_code).
- `timeouts` (Block List, Min: 1) - Can pass custom timeout while create . Example: create = "20m"  - such as "60m" for
  60 minutes, "10s" for ten seconds, or "2h" for two hours.

<a id="nestedblock--server"></a>

### Nested Schema for `server`

Required:

- `hostname` (String) - A FQDN for the device. For example: example.coxedge.com
