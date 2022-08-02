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

resource "coxedge_origin_setting" "testing" {
  organization_id             = "<organization_id>"
  site_id                     = "<site_id>"
  environment_name            = "<environment_name>"
  domain                      = "bluegreen.com"
  websockets_enabled          = "false"
  ssl_validation_enabled      = "false"
  pull_protocol               = "MATCH"
  host_header                 = "Host: marvel.com"
  backup_origin_enabled       = "true"
  backup_origin_exclude_codes = ["415"]

  origin {
    #id = ""
    address     = "www.test.com:80"
    auth_method = "BASIC"
    username    = "terraform-user"

    common_certificate_name = "commanName"
  }
}