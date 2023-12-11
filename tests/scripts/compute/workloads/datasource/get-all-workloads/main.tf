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

data "coxedge_compute_workloads" "workloads" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
}

output "output_workloads" {
  value = data.coxedge_compute_workloads.workloads
}