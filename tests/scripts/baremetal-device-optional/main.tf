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

resource "coxedge_baremetal_devices" "device" {
  environment_name  = "<environment_name>"
  organization_id   = "<organization_id>"
  location_name     = "ATL2"
  has_user_data     = true
  has_ssh_data      = true
  product_option_id = 144463
  product_id        = "580"
  os_name           = "Ubuntu 20.x"
  server {
    hostname = "test001.coxedge.com"
  }
  user_data    = "<user_data>"
  ssh_key      = "<ssh_key>"
  ssh_key_name = "<ssh_key_name>"
  ssh_key_id   = "<ssh_key_id>"
}