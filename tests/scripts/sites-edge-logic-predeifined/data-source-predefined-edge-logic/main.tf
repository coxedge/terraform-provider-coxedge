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

data "coxedge_sites_predefined_edge_logic" "edge_logic" {
  id = "<site_id>"
  environment_name = "<environment_name>"
  organization_id = "<organization_id>"
}

output "output_edge_logic" {
  value = data.coxedge_sites_predefined_edge_logic.edge_logic
}