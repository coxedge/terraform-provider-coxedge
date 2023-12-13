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

resource "coxedge_baremetal_device" "device" {
  id               = 0
  environment_name = "<environment_name>"
  organization_id  = "<organization_id>"
  hostname         = "<hostname>"
  name             = "<name>"
  tags             = tolist(["<tags>"])
  power_status     = "<power_status>" // ON or OFF
}