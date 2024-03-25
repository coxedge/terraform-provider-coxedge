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

resource "coxedge_compute_reserved_ip_attach_detach_instance" "reserved_ip" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  reserved_ip_id   = "<reserved_ip_id>"
  action           = "detach"
}