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
#terraform import coxedge_baremetal_device.device <device_id>:<environment_name>:<organization_id>
resource "coxedge_baremetal_device" "device" {
}

output "device_details" {
  value = coxedge_baremetal_device.device
}

//for hivelocity
resource "coxedge_baremetal_devices" "device" {
  environment_name  = "bm-env"
  organization_id   = "e5290682-44f4-481b-9327-34f677a1c46c"
  location_name     = "LAX2"
  has_user_data     = true
  has_ssh_data      = true
  product_option_id = 143004
  product_id        = "500"
  os_name           = "Ubuntu 20.x"
  server {
    hostname = "test001.coxedge.com"
  }
  user_data    = "user data sample text"
  ssh_key      = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDbuOix2IU7SvzJQP61aZ05b2KEYAmTssbC6mrRI/4UZ33BKCPMH039ipw33MyCM3xVHY4n6VshLzExAQLb9qJw7F2yA7shPk1sN9XI1gAestp1IAW/9iF3E8cQyZIV2JUm/ixnRFzlLDNr1gYn8W0XxbW1FJ17QLY4PigEx2WB2LJ28BBOBItY2w8ipWP0ZuDYWVoV8zbWmr/3cdLxp7sNcL6D2MD9t6iYnks2REeGCBtmPaRzuZVgd/g66HmR+614vFAFeT66XBe5HEO2ERSohXMnS8DNSTeRSeN6KcWtEdWMqC5RWkypb+/sNI4WH7SB5TPjrY3jUPSUfpmsZd+WZQM7hG9DTm5mSgMfbVJAF4hSJnN31b/FvuNecDEpaGoTTByS4cZKFNjOxnvZHqtWThN3Y371VsOXMwsSaO1I8v5ylj1YrK86xd4XE3XxFaIFMGfxrM7JcqbmTwhdReivu9+TYKKXvdLzbGIqtD+mjwFQakQj9mWKg+aELdlnrH8= shehzabahammad@DESKTOP-6IH172L"
  ssh_key_name = "shehzab-machine"
}

//for metalsoft
resource "coxedge_baremetal_devices" "device" {
  environment_name = "sanityhiv"
  organization_id  = "7e80611c-29c5-4bab-8e6d-1c4fc5b2c035"
  location_name    = "cox-dvtc"
  product_id       = "1"
  vendor           = "METALSOFT"
  os_id            = "4"
  server_label     = "testterraformmm2"
  tags             = tolist(["tag added"])
}