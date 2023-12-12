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


data "coxedge_compute_workload_ipv6_reverse_dns" "ipv6_reverse_dns" {
  environment_name = "test"
  organization_id  = "b0d424e4-4f78-4cb3-8c7c-26781bea9f7e"
  workload_id      = "79bc4c82-f884-452a-b790-eb12c2b58ea5"
}

output "output_ipv6_reverse_dns" {
  value = data.coxedge_compute_workload_ipv6_reverse_dns.ipv6_reverse_dns
}
