---

# generated by https://github.com/hashicorp/terraform-plugin-docs

page_title: "coxedge_compute_workload Resource - terraform-provider-coxedge"
subcategory: ""
description: Allows you to manage your Workloads
  
---

# coxedge_compute_workload (Resource)

Allows you to manage your Workloads

Example Usage
---

```
//to import existing device which is created outside terraform script or from terraform script itself
//terraform import coxedge_compute_workload.workload <workload_id>:<environment_name>:<organization_id>
resource "coxedge_compute_workload" "workload" {
  organization_id           = "<organization_id>"
  environment_name          = "<environment name>"
  is_ipv6                   = true
  no_public_ipv4            = true
  is_virtual_private_clouds = false
  is_vpc2                   = false
  operating_system_id       = "2076"
  location_id               = "ams"
  plan_id                   = "voc-g-1c-4gb-30s-amd"
  hostname                  = "testterraform1"
  label                     = "testterraform1"
  first_boot_ssh_key        = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDbuOix2IU7SvzJQP61qwe5b2KEYAmTssbC6mrRI/4UZ33BKCPMH039ipw33MyCM3xVHY4n6VshLzExAQLb9qJw7F2yA7shPk1sN9XI1gAestp1IAW/9iF3E8cQyZIV2JUm/ixnRFzlLDNr1gYn8W0XxbW1FJ17QLY4PigEx2WB2LJ28BBOBItY2w8ipWP0ZuDYWVoV8zbWmr/3cdLxp7sNcL6D2MD9t6iYnks2REeGCBtmPaRzuZVgd/g66HmR+614vFAFeT66XBe5HEO2ERSohXMnS8DNSTeRSeN6KcWtEdWMqC5RWkypb+/sNI4WH7SB5TPjrY3jUPSUfpmsZd+WZQM7hG9DTm5mSgMfbVJAF4hSJnN31b/FvuNecDEpaGoTTByS4cZKFNjOxnvZHqtWThN3Y371VsOXMwsSaO1I8v5ylj1YrK86xd4XE3XxFaIFMGfxrM7JcqbmTwhdReivu9+TYKKXvdLzbGIqtD+mjwFQakQj9mWKg+aELdlnrH8= dumm@L"
  ssh_key_name              = "test-ssh"
  firewall_id               = "f7a1c207-1666-48c2-9c6c-339021c0d440"
  user_data                 = "test data here"
}
```

<!-- schema generated by tfplugindocs -->

## Argument Reference

The following arguments are supported:

### Required

- `environment_name` (String) - The name of the environment.
- `organization_id` (String) - The id of the organization.
- `is_ipv6` (Bool) - true/false
- `location_id` (String) - The Region id where the Workload is located.
- `plan_id` (String) - The Plan id to use when deploying this workload.
- `hostname` (String) - The hostname to use when deploying this workload.
- `label` (String) - A user-supplied label for this workload.

### Optional

- `workload_id` (String) - The id of the workload.
- `no_public_ipv4` (Bool) - Don't set up a public IPv4 address when is_ipv6 is enabled.
- `is_virtual_private_clouds` (Bool) - True/False
- `is_vpc2` (Bool) - True/False
- `operating_system_id` (String) - The Operating System id to use when deploying this workload.
- `first_boot_ssh_key` (String) - SSH key value
- `ssh_key_name` (String) - SSH Key name
- `firewall_id` (String) - The Firewall Group id to attach to this Workload.
- `user_data` (String) - Linux-only: The user scheme used for logging into this workload. By default, the "root" user is
  configured. Alternatively, a limited user with sudo permissions can be selected.