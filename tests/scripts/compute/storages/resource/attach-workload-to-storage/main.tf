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

resource "coxedge_compute_storage_attach_detach_instance" "attach" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  storage_id       = "<storage_id>"
  live             = false
  instance_id      = "<workload_id>"
  action           = "attach"
}
