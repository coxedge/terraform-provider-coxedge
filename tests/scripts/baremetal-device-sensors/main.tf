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

data "coxedge_baremetal_device_sensors" "sensors" {
  id               = "<id>"
  environment_name = "<environment_name>"
  organization_id  = "<organization_id>"
}

output "sensors_output" {
  value = data.coxedge_baremetal_device_sensors.sensors
}