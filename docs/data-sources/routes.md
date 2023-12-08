---

# generated by https://github.com/hashicorp/terraform-plugin-docs

page_title: "coxedge_routes Data Source - terraform-provider-coxedge"
subcategory: ""
description: Cox Edge Computing uses the concept of routes in VPC.
  
---

# coxedge_routes (Data Source)

Cox Edge Computing uses the concept of routes in VPC.

Example Usage
---

```
data "coxedge_routes" "route" {
  environment_name = "<environment_name>"
  organization_id  = "<organization_id>"
  vpc_id           = "<vpc_id>"
}

output "output_routes" {
  value = data.coxedge_routes.route
}
```

<!-- schema generated by tfplugindocs -->

## Schema

### Required

- `environment_name` (String) - Name of the environment belonging to the organization.
- `organization_id` (String) - The id of the organization.
- `vpc_id` (String) - The id of the VPC to which routes belongs.

### Read-Only

- `id` (String) The ID of this resource.
- `routes` (List of Object) (see [below for nested schema](#nestedatt--routes)) - Number of routes in the VPC.

<a id="nestedatt--routes"></a>

### Nested Schema for `routes`

Read-Only:

- `id` (String) - ID of the VPC.
- `name` (String) - Name of the VPC.
- `stack_id` (String) - The stack ID to which this VPC belongs.
- `slug` (String) - slug of VPC.
- `destination_cidrs` (List of String) - destination_cidrs of the VPC.
- `next_hops` (List of String) - next_hops of the VPC.
- `vpc_id` (String) - vpc_id of subnet belongs to.
- `status` (String) - The VPC status.

