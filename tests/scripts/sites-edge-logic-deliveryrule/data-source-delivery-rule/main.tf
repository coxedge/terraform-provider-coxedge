terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source  = "coxedge.com/cox/coxedge"
    }
  }
}

provider "coxedge" {
  key = ""
}

data "coxedge_sites_edge_logic_delivery_rules" "delivery_rule" {
  environment_name = "<environment_name>"
  organization_id  = "<organization_id>"
  id               = "<id>"
}

output "delivery_rules_output" {
  value = data.coxedge_sites_edge_logic_delivery_rules.delivery_rule
}