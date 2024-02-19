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

resource "coxedge_compute_reserved_ips" "reserved-ip" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  region           = "<region>"
  ip_type          = "<ip_type>"
  label            = "<label>"
}