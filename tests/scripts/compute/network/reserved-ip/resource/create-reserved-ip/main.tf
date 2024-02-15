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

resource "coxedge_compute_reserved_ip_convert" "convert-reserved-ip" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  ip_type          = "<ip_type>"
  ip_address       = "<ip_address>"
}