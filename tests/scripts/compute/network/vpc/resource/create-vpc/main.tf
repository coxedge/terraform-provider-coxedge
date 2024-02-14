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

resource "coxedge_compute_vpc" "vpc" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  location_id      = "<location_id>"
  v4_subnet_mask   = "0"
  network_prefix   = 0
  ip_range         = "0"
  route_id         = "0"
  v4_subnet        = ""
  description      = "<description>"
  routes {
    destination    = ""
    network_prefix = ""
    target_address = ""
  }
}