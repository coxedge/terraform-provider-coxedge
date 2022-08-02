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


resource "coxedge_delivery_domain" "testing" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  domain           = "<domain name>"
  site_id          = "<site-id>"
}
