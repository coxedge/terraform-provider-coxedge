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

##create device
#resource "coxedge_baremetal_devices" "device" {
#  environment_name  = "bm-env"
#  organization_id   = "e5290682-44f4-481b-9327-34f677a1c46c"
#  location_name     = "LAX2"
#  has_user_data     = true
#  has_ssh_data      = true
#  product_option_id = 143004
#  product_id        = "500"
#  os_name           = "Ubuntu 20.x"
#  server {
#    hostname = "test001.coxedge.com"
#  }
#  user_data    = "user data sample text"
#  ssh_key      = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDbuOix2IU7SvzJQP61aZ05b2KEYAmTssbC6mrRI/4UZ33BKCPMH039ipw33MyCM3xVHY4n6VshLzExAQLb9qJw7F2yA7shPk1sN9XI1gAestp1IAW/9iF3E8cQyZIV2JUm/ixnRFzlLDNr1gYn8W0XxbW1FJ17QLY4PigEx2WB2LJ28BBOBItY2w8ipWP0ZuDYWVoV8zbWmr/3cdLxp7sNcL6D2MD9t6iYnks2REeGCBtmPaRzuZVgd/g66HmR+614vFAFeT66XBe5HEO2ERSohXMnS8DNSTeRSeN6KcWtEdWMqC5RWkypb+/sNI4WH7SB5TPjrY3jUPSUfpmsZd+WZQM7hG9DTm5mSgMfbVJAF4hSJnN31b/FvuNecDEpaGoTTByS4cZKFNjOxnvZHqtWThN3Y371VsOXMwsSaO1I8v5ylj1YrK86xd4XE3XxFaIFMGfxrM7JcqbmTwhdReivu9+TYKKXvdLzbGIqtD+mjwFQakQj9mWKg+aELdlnrH8= shehzabahammad@DESKTOP-6IH172L"
#  ssh_key_name = "shehzab-machine"
#}
#
##import, delete device
##terraform import coxedge_baremetal_device.device <device_id>:<environment_name>:<organization_id>
resource "coxedge_baremetal_device" "device" {
}

output "device_details" {
  value = coxedge_baremetal_device.device
}
#
##update device
#resource "coxedge_baremetal_device" "device" {
#  id               = "14092"
#  environment_name = "bm-env"
#  organization_id  = "e5290682-44f4-481b-9327-34f677a1c46c"
#  hostname         = "test002.coxedge.com"
#  name             = "test002.coxedge.com"
#  tags             = tolist(["tag1"])
#  power_status     = "OFF"
#}
#
###charts
#data "coxedge_baremetal_device_charts" "charts" {
#  id               = "14092"
#  environment_name = "bm-env"
#  organization_id  = "e5290682-44f4-481b-9327-34f677a1c46c"
#  custom           = true
#  start_date       = "1668184639"
#  end_date         = "1668407548"
#}
#
#output "charts_output" {
#  value = data.coxedge_baremetal_device_charts.charts
#}
#
###sensors
#data "coxedge_baremetal_device_sensors" "sensors" {
#  id               = "14092"
#  environment_name = "bm-env"
#  organization_id  = "e5290682-44f4-481b-9327-34f677a1c46c"
#}
#
#output "sensors_output" {
#  value = data.coxedge_baremetal_device_sensors.sensors
#}
#
#ipmi
#resource "coxedge_baremetal_device_ipmi" "ipmi" {
#  device_id        = "14092"
#  environment_name = "bm-env"
#  organization_id  = "e5290682-44f4-481b-9327-34f677a1c46c"
#  custom_ip        = "103.147.208.242"
#}
#
##ips
#data "coxedge_baremetal_device_ips" "ips" {
#  id               = "14092"
#  environment_name = "bm-env"
#  organization_id  = "e5290682-44f4-481b-9327-34f677a1c46c"
#}
#
#output "ips_output" {
#  value = data.coxedge_baremetal_device_ips.ips
#}
#
##create sshkey
#resource "coxedge_baremetal_ssh_key" "sshkey" {
#  environment_name = "bm-env"
#  organization_id  = "e5290682-44f4-481b-9327-34f677a1c46c"
#  name             = "terraform-script"
#  public_key       = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDbuOix2IU7SvzJQP61aZ05b2KEYAmTssbC6mrRI/4UZ33BKCPMH039ipw33MyCM3xVHY4n6VshLzExAQLb9qJw7F2yA7shPk1sN9XI1gAestp1IAW/9iF3E8cQyZIV2JUm/ixnRFzlLDNr1gYn8W0XxbW1FJ17QLY4PigEx2WB2LJ28BBOBItY2w8ipWP0ZuDYWVoV8zbWmr/3cdLxp7sNcL6D2MD9t6iYnks2REeGCBtmPaRzuZVgd/g66HmR+614vFAFeT66XBe5HEO2ERSohXMnS8DNSTeRSeN6KcWtEdWMqC5RWkypb+/sNI4WH7SB5TPjrY3jUPSUfpmsZd+WZQM7hG9DTm5mSgMfbVJAF4hSJnN31b/FvuNecDEpaGoTTByS4cZKFNjOxnvZHqtWThN3Y371VsOXMwsSaO1I8v5ylj1YrK86xd4XE3XxFaIFMGfxrM7JcqbmTwhdReivu9+TYKKXvdLzbGIqtD+mjwFQakQj9mWKg+aELdlnrH8= shehzabahammad@DESKTOP-6IH172L"
#}
#
##sshkey
#data "coxedge_baremetal_ssh_keys" "sshkeys" {
#  environment_name = "bm-env"
#  organization_id  = "e5290682-44f4-481b-9327-34f677a1c46c"
#}
#
#output "sshkeys_output" {
#  value = data.coxedge_baremetal_ssh_keys.sshkeys
#}
