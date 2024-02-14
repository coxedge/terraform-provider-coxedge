terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source  = "coxedge.com/cox/coxedge"
    }
  }
}

provider "coxedge" {
  key = "[INSERT API KEY HERE]"
}

resource "coxedge_compute_vpc2" "vpc2" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  location_id      = "<location_id>"
  ip_range         = "<0-auto assign/1-configure>"
  prefix_length    = "0"
  ip_block         = ""
  description      = "<description>"
}