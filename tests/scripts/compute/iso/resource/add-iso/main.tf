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

resource "coxedge_compute_isos" "iso" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  url              = "<url>" #http://205.185.126.191/web/iso/Server2022x64.iso
}