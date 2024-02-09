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

resource "coxedge_compute_firewalls" "firewalls" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  description      = "<description>"
}