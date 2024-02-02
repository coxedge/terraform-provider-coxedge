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

data "coxedge_compute_storages" "storages" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
}

output "output_storage" {
  value = data.coxedge_compute_storages.storages
}