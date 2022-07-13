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

resource "coxedge_script" "testing" {
  site_id          = "<site-id>"
  environment_name = "<environment_name>"
  name             = "script-test"
  routes           = ["v1/api"]
  code             = "sample script test"
}
