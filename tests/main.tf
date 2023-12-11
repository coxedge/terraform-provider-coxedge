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


data "coxedge_compute_workload_ipv6" "ipv6" {
  environment_name = "test"
  organization_id  = "b0d424e4-4f78-4cb3-8c7c-26781bea9f7e"
  workload_id      = "90ab4e15-14c2-4643-8d10-1580964fd09c"
}

output "output_ipv6s" {
  value = data.coxedge_compute_workload_ipv6.ipv6
}
