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

data "coxedge_compute_isos" "isos" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  iso_id           = "<iso_id>"
}

output "output_tag" {
  value = data.coxedge_compute_isos.isos
}