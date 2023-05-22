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

// terraform import coxedge_sites_edge_logic_delivery_rule.delivery_rule <delivery_rule_id>:<site_id>:<environment_name>:<organization_id>
resource "coxedge_sites_edge_logic_delivery_rule" "delivery_rule" {
  environment_name = "<environment_name>"
  organization_id  = "<organization_id>"
  site_id          = "<site_id>"
}