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

data "coxedge_compute_workload_ipv6_reverse_dns" "ipv6_reverse_dns" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  workload_id      = "<workload_id>"
}

output "output_ipv6_reverse_dns" {
  value = data.coxedge_compute_workload_ipv6_reverse_dns.ipv6_reverse_dns
}