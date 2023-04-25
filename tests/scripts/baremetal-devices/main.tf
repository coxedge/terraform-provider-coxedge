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

//Get all BareMetal devices
data "coxedge_baremetal_devices" "baremetals" {
  environment_name = "<environment_name>"
  organization_id  = "<organization_id>"
}

//Get BareMetal device by Id
data "coxedge_baremetal_devices" "baremetals" {
  environment_name = "<environment_name>"
  organization_id  = "<organization_id>"
  id               = "<device_id>"
}

output "output" {
  value = data.coxedge_baremetal_devices.baremetals
}