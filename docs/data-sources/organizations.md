---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "coxedge_organizations Data Source - terraform-provider-coxedge"
subcategory: ""
description: Organizations are the largest logical grouping of users, environments and resources available in Cox Edge. Each organization is isolated from other organizations. It is protected by its own customizable system roles. Additionally, provisioned resource usage is metered at the organization level facilitating cost tracking.
  
---

# coxedge_organizations (Data Source)
Organizations are the largest logical grouping of users, environments and resources available in Cox Edge. Each organization is isolated from other organizations. It is protected by its own customizable system roles. Additionally, provisioned resource usage is metered at the organization level facilitating cost tracking.

Example Usage
---
```
data "coxedge_organizations" "orgs" {
  id   = "organization-id"
}

output "testing" {
  value = data.coxedge_organizations.test
}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) - The ID of this resource.
- `organizations` (List of Object) (see [below for nested schema](#nestedatt--organizations)) - The organization of the environment.

<a id="nestedatt--organizations"></a>
### Nested Schema for `organizations`

Read-Only:

- `entry_point` (String) - The entry point of the organization is the subdomain of the organization in the Cox Edge URL : [entryPoint].Cox Edge.
- `id` (String) - The id of the organization.
- `name` (String) - The name of the organization.
- `service_connections` (List of Object) (see [below for nested schema](#nestedobjatt--organizations--service_connections)) - The services for which the organization is allowed to provision resources.
  includes: id,serviceCode
- `tags` (List of String) - Tags associated to the organization.

<a id="nestedobjatt--organizations--service_connections"></a>
### Nested Schema for `organizations.service_connections`

Read-Only:

- `id` (String) - Organization id.
- `service_code` (String) - - `name` (String) - The name of the environment.
- `users` (List of String) - The users that are members of the environment.
- `name` (String) - The name of the environment.
- `users` (List of String) - The users that are members of the environment.
