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

resource "coxedge_compute_firewall_ipv6_rule" "ipv6" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  firewall_id      = "<firewall_id>"
  cidr             = "::/0"
  protocol         = "tcp"
  source_option    = "anywhere"
  port             = "8080"
  notes            = "sdadsadsdsadsfewfadsa"
}
