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

data "coxedge_compute_firewall_ipv4_rule" "ipv4" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  firewall_id      = "<firewall_id>"
}

output "output_ipv4" {
  value = data.coxedge_compute_firewall_ipv4_rule.ipv4
}