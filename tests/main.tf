terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source  = "coxedge.com/cox/coxedge"
    }
  }
}

provider "coxedge" {
  key = "cJqctxd5QjVlwNXm/hOfcg=="
}

#resource "coxedge_baremetal_devices" "device" {
#  environment_name = "resellerenv"
#  organization_id  = "b0d424e4-4f78-4cb3-8c7c-26781bea9f7e"
#  location_code    = "cox-dvtc"
#  product_id       = "3"
#  os_id            = "14"
#  server_label     = "testterraform"
#  tags             = tolist(["tag added"])
#}

resource "coxedge_baremetal_device" "device" {
  environment_name = "resellerenv"
  organization_id  = "b0d424e4-4f78-4cb3-8c7c-26781bea9f7e"
  id               = "MS_372"
  tags             = tolist(["tag added", "tag 3"])
  power_status     = "ON"
}

#data "coxedge_baremetal_device_sensors" "sensors" {
#  environment_name = "resellerenv"
#  organization_id  = "b0d424e4-4f78-4cb3-8c7c-26781bea9f7e"
#  id               = "MS_370"
#}
#
#output "output" {
#  value = data.coxedge_baremetal_device_sensors.sensors
#}
#data "coxedge_baremetal_ssh_keys" "ssh" {
#  environment_name = "sanityhiv"
#  organization_id  = "7e80611c-29c5-4bab-8e6d-1c4fc5b2c035"
##  id               = "HV_14000"
##  id               = "MS_249"
#  id               = "48"
#}
#
#output "out_vpc" {
#  value = data.coxedge_baremetal_ssh_keys.ssh
#}

#resource "coxedge_baremetal_devices" "device" {
#  environment_name = "sanityhiv"
#  organization_id  = "7e80611c-29c5-4bab-8e6d-1c4fc5b2c035"
#  location_name    = "cox-dvtc"
#  #  has_user_data     = true
#  #  has_ssh_data      = true
#  #  product_option_id = 143004
#  product_id       = "1"
#  #  os_name           = "Ubuntu 20.x"
#  #  server {
#  #    hostname = "testterraform001.coxedge.com"
#  #  }
#  #  user_data    = "user data sample text"
#  #  ssh_key      = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDbuOix2IU7SvzJQP61aZ05b2KEYAmTssbC6mrRI/4UZ33BKCPMH039ipw33MyCM3xVHY4n6VshLzExAQLb9qJw7F2yA7shPk1sN9XI1gAestp1IAW/9iF3E8cQyZIV2JUm/ixnRFzlLDNr1gYn8W0XxbW1FJ17QLY4PigEx2WB2LJ28BBOBItY2w8ipWP0ZuDYWVoV8zbWmr/3cdLxp7sNcL6D2MD9t6iYnks2REeGCBtmPaRzuZVgd/g66HmR+614vFAFeT66XBe5HEO2ERSohXMnS8DNSTeRSeN6KcWtEdWMqC5RWkypb+/sNI4WH7SB5TPjrY3jUPSUfpmsZd+WZQM7hG9DTm5mSgMfbVJAF4hSJnN31b/FvuNecDEpaGoTTByS4cZKFNjOxnvZHqtWThN3Y371VsOXMwsSaO1I8v5ylj1YrK86xd4XE3XxFaIFMGfxrM7JcqbmTwhdReivu9+TYKKXvdLzbGIqtD+mjwFQakQj9mWKg+aELdlnrH8= shehzabahammad@DESKTOP-6IH172L"
#  #  ssh_key_name = "shehzab-machine"
#  vendor           = "METALSOFT"
#  os_id            = "4"
#  server_label     = "testterraformmm3"
#  tags             = tolist(["tag added"])
#}

#resource "coxedge_baremetal_device" "device" {
#  power_status = "OFF"
##  tags         = tolist(["hello2"])
##  name         = "terraformtest"
#}

#data "coxedge_baremetal_location_product_os" "os" {
#  environment_name = "sanityhiv"
#  organization_id  = "7e80611c-29c5-4bab-8e6d-1c4fc5b2c035"
#    vendor_product_id = "MS_2"
##  id               = "MS_4"
##  code             = "cox-dvtc"
#}
#
#output "output_disk" {
#  value = data.coxedge_baremetal_location_product_os.os
#}

#data "coxedge_baremetal_devices" "test" {
#    environment_name  = "sanityhiv"
#    organization_id   = "7e80611c-29c5-4bab-8e6d-1c4fc5b2c035"
##  id               = "<device_id>"
#}
#
#output "testing" {
#  value = data.coxedge_baremetal_devices.test
#}

#resource "coxedge_baremetal_devices" "device" {
#  environment_name = "sanityhiv"
#  organization_id  = "7e80611c-29c5-4bab-8e6d-1c4fc5b2c035"
#  location_code    = "cox-dvtc"
#  product_id       = "1"
#  os_id            = "4"
#  server_label     = "testterraform"
#  tags             = tolist(["tag added"])
#}

#resource "coxedge_baremetal_devices" "device" {
#  environment_name = "sanityhiv"
#  organization_id  = "7e80611c-29c5-4bab-8e6d-1c4fc5b2c035"
#  location_code     = "LAX2"
#  has_user_data     = true
#  has_ssh_data      = true
#  product_option_id = 144178
#  product_id        = "504"
#  os_name           = "Ubuntu 18.x"
#  server {
#    hostname = "testterraform001.coxedge.com"
#  }
#  user_data    = "user data sample text"
#  ssh_key      = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDbuOix2IU7SvzJQP61aZ05b2KEYAmTssbC6mrRI/4UZ33BKCPMH039ipw33MyCM3xVHY4n6VshLzExAQLb9qJw7F2yA7shPk1sN9XI1gAestp1IAW/9iF3E8cQyZIV2JUm/ixnRFzlLDNr1gYn8W0XxbW1FJ17QLY4PigEx2WB2LJ28BBOBItY2w8ipWP0ZuDYWVoV8zbWmr/3cdLxp7sNcL6D2MD9t6iYnks2REeGCBtmPaRzuZVgd/g66HmR+614vFAFeT66XBe5HEO2ERSohXMnS8DNSTeRSeN6KcWtEdWMqC5RWkypb+/sNI4WH7SB5TPjrY3jUPSUfpmsZd+WZQM7hG9DTm5mSgMfbVJAF4hSJnN31b/FvuNecDEpaGoTTByS4cZKFNjOxnvZHqtWThN3Y371VsOXMwsSaO1I8v5ylj1YrK86xd4XE3XxFaIFMGfxrM7JcqbmTwhdReivu9+TYKKXvdLzbGIqtD+mjwFQakQj9mWKg+aELdlnrH8= shehzabahammad@DESKTOP-6IH172L"
#  ssh_key_name = "shehzab-machine"
#}
