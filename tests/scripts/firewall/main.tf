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

resource "coxedge_firewall_rule" "testing" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  site_id          = "<site-id>"
  action           = "ALLOW"
  ip_start         = "192.168.0.6"
  name             = "firewall.test.1"
  ip_end           = "192.168.0.7"
}
