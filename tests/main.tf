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

data "coxedge_baremetal_device_ips" "ips" {
  id               = "14092"
  environment_name = "bm-env"
  organization_id  = "e5290682-44f4-481b-9327-34f677a1c46c"
}

output "ips_output" {
  value = data.coxedge_baremetal_device_ips.ips
}

#resource "coxedge_baremetal_device_ipmi" "ipmi" {
#  device_id        = "14092"
#  environment_name = "bm-env"
#  organization_id  = "e5290682-44f4-481b-9327-34f677a1c46c"
#  custom_ip        = "103.153.104.63"
#
#}
#
#data "coxedge_baremetal_device_sensors" "sensors" {
#  id               = "14092"
#  environment_name = "bm-env"
#  organization_id  = "e5290682-44f4-481b-9327-34f677a1c46c"
#}
#
#output "sensors_output" {
#  value = data.coxedge_baremetal_device_sensors.sensors
#}

#resource "coxedge_baremetal_device" "device" {
#  id               = 14882
#  environment_name = "bm-env"
#  organization_id  = "e5290682-44f4-481b-9327-34f677a1c46c"
##  hostname         = "test022.coxedge.com"
##  name             = "test022.coxedge.com"
##  tags             = tolist(["tag1"])
#  power_status     = "ON"
#}

#resource "coxedge_baremetal_devices" "device" {
#  environment_name  = "bm-env"
#  organization_id   = "e5290682-44f4-481b-9327-34f677a1c46c"
#  location_name     = "ATL2"
#  has_user_data     = true
#  has_ssh_data      = false
#  product_option_id = 144463
#  product_id        = "580"
#  os_name           = "Ubuntu 20.x"
#  server {
#    hostname = "testterr006.coxedge.com"
#  }
#  user_data  = "test user data field from terraform"
#  ssh_key_id = "923"
#}

#//Get all BareMetal devices
#data "coxedge_baremetal_devices" "baremetals" {
#  environment_name  = "bm-env"
#  organization_id   = "e5290682-44f4-481b-9327-34f677a1c46c"
#  id = 14882
#}
#
#output "output" {
#  value = data.coxedge_baremetal_devices.baremetals
#}