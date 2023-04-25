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

#do import existing resource, do run below script
#terraform import coxedge_baremetal_device.device <site_id>:<environment_name>:<organization_id>
resource "coxedge_baremetal_device" "device" {
}

output "device_details" {
  value = coxedge_baremetal_device.device
}