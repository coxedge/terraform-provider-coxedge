terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source  = "coxedge.com/cox/coxedge"
    }
  }
}

provider "coxedge" {
  key = "GM3COPLOU6nOI12/NZ7HNg=="
}

data "coxedge_compute_workloads" "workloads" {
  environment_name = "test"
  organization_id = "b0d424e4-4f78-4cb3-8c7c-26781bea9f7e"
}

output "output_workloads" {
  value = data.coxedge_compute_workloads.workloads
}