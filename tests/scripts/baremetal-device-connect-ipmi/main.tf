terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source  = "coxedge.com/cox/coxedge"
    }
  }
}

provider "coxedge" {
  key = ""
}

resource "coxedge_baremetal_device_ipmi" "ipmi" {
  device_id        = "<device_id>"
  environment_name = "<environment_name>"
  organization_id  = "<organization_id>"
  custom_ip        = "<custom_ip>"

}