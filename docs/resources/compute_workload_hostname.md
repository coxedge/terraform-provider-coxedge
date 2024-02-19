---

# generated by https://github.com/hashicorp/terraform-plugin-docs

page_title: "coxedge_compute_workload_hostname Resource - terraform-provider-coxedge"
subcategory: ""
description: Allows you to manage hostname
  
---

# coxedge_compute_workload_hostname (Resource)

Allows you to manage hostname

Example Usage
---

```
resource "coxedge_compute_workload_hostname" "hostname" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  workload_id      = "<workload_id>"
  hostname         = "<hostname>"
}
```

<!-- schema generated by tfplugindocs -->

## Argument Reference

The following arguments are supported:

### Required

- `environment_name` (String) - The name of the environment.
- `organization_id` (String) - The id of the organization.
- `workload_id` (String) - The id of the workload.
- `hostname` (String) - Hostname of workload