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

//terraform import coxedge_compute_workload_ipv6_reverse_dns.ipv6_reverse_dns <workload_id>:<environment_name>:<organization_id>
resource "coxedge_compute_workload_ipv6_reverse_dns" "ipv6_reverse_dns" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  workload_id      = "<workload_id>"
  ip               = "2001:19f0:5001:2ce1:5400:04ff:feae:50f6"
  reverse          = "tester1.foos.com"
}