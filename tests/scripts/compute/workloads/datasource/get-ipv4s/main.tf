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

data "coxedge_compute_workload_ipv4" "ipv4" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  workload_id      = "<workload_id>"
}

output "output_ipv4s" {
  value = data.coxedge_compute_workload_ipv4.ipv4
}