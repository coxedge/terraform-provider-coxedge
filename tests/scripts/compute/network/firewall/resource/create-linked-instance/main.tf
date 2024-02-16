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

resource "coxedge_compute_firewall_linked_instances" "linked_instance" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  firewall_id      = "<firewall_id>"
  workload_id      = "<workload_id>"
}