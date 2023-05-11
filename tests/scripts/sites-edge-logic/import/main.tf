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

// terraform import coxedge_sites_predefined_edge_logic.edge_logic <site_id>:<environment_name>:<organization_id>
resource "coxedge_sites_predefined_edge_logic" "edge_logic" {
}