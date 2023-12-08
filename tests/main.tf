terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source  = "coxedge.com/cox/coxedge"
    }
  }
}

provider "coxedge" {
  key = "GM3COPLOU6nOI12/NZ7HNg=="
}

resource "coxedge_compute_workload" "workload" {
  environment_name          = "test"
  organization_id           = "b0d424e4-4f78-4cb3-8c7c-26781bea9f7e"
  is_ipv6                   = true
  no_public_ipv4            = true
#  is_virtual_private_clouds = false
#  is_vpc2                   = false
#  server_id                 = "voc"
#  server_type_id            = "voc-g"
#  image_id                  = "OS_2076"
  operating_system_id       = "2076"
#  plan_filter               = "voc-g"
#  continent                 = "All Locations"
  location_id               = "ams"
  plan_id                   = "voc-g-1c-4gb-30s-amd"
  hostname                  = "testterraform1"
  label                     = "testterraform1"
  first_boot_ssh_key        = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDbuOix2IU7SvzJQP61aZ05b2KEYAmTssbC6mrRI/4UZ33BKCPMH039ipw33MyCM3xVHY4n6VshLzExAQLb9qJw7F2yA7shPk1sN9XI1gAestp1IAW/9iF3E8cQyZIV2JUm/ixnRFzlLDNr1gYn8W0XxbW1FJ17QLY4PigEx2WB2LJ28BBOBItY2w8ipWP0ZuDYWVoV8zbWmr/3cdLxp7sNcL6D2MD9t6iYnks2REeGCBtmPaRzuZVgd/g66HmR+614vFAFeT66XBe5HEO2ERSohXMnS8DNSTeRSeN6KcWtEdWMqC5RWkypb+/sNI4WH7SB5TPjrY3jUPSUfpmsZd+WZQM7hG9DTm5mSgMfbVJAF4hSJnN31b/FvuNecDEpaGoTTByS4cZKFNjOxnvZHqtWThN3Y371VsOXMwsSaO1I8v5ylj1YrK86xd4XE3XxFaIFMGfxrM7JcqbmTwhdReivu9+TYKKXvdLzbGIqtD+mjwFQakQj9mWKg+aELdlnrH8= shehzabahammad@DESKTOP-6IH172L"
  ssh_key_name              = "test-ssh"
  firewall_id               = "f7a1c207-1666-48c2-9c6c-339021c0d440"
  user_data                 = "test data here"
}
