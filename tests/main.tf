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

resource "coxedge_workload" "test" {
  name               = "terr-test2"
  organization_id    = "b0d424e4-4f78-4cb3-8c7c-26781bea9f7e"
  environment_name   = "test-backend"
  type               = "VM"
  image              = "stackpath-edge/ubuntu-2004-focal:v202102241556"
  first_boot_ssh_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDbuOix2IU7SvzJQP61aZ05b2KEYAmTssbC6mrRI/4UZ33BKCPMH039ipw33MyCM3xVHY4n6VshLzExAQLb9qJw7F2yA7shPk1sN9XI1gAestp1IAW/9iF3E8cQyZIV2JUm/ixnRFzlLDNr1gYn8W0XxbW1FJ17QLY4PigEx2WB2LJ28BBOBItY2w8ipWP0ZuDYWVoV8zbWmr/3cdLxp7sNcL6D2MD9t6iYnks2REeGCBtmPaRzuZVgd/g66HmR+614vFAFeT66XBe5HEO2ERSohXMnS8DNSTeRSeN6KcWtEdWMqC5RWkypb+/sNI4WH7SB5TPjrY3jUPSUfpmsZd+WZQM7hG9DTm5mSgMfbVJAF4hSJnN31b/FvuNecDEpaGoTTByS4cZKFNjOxnvZHqtWThN3Y371VsOXMwsSaO1I8v5ylj1YrK86xd4XE3XxFaIFMGfxrM7JcqbmTwhdReivu9+TYKKXvdLzbGIqtD+mjwFQakQj9mWKg+aELdlnrH8= shehzabahammad@DESKTOP-6IH172L"
  specs              = "SP-1"
  deployment {
    name               = "testvm"
    enable_autoscaling = false
    pops               = ["LAS"]
    instances_per_pop  = 1
  }
  probe_configuration = "LIVENESS"
  #  liveness_probe {
  #    initial_delay_seconds = 0
  #    timeout_seconds       = 1
  #    period_seconds        = 10
  #    success_threshold     = 1
  #    failure_threshold     = 3
  #    protocol              = "TCP_SOCKET"
  #    tcp_socket {
  #      port = 22
  #    }
  #  }
  liveness_probe {
    initial_delay_seconds = 0
    timeout_seconds       = 1
    period_seconds        = 10
    success_threshold     = 1
    failure_threshold     = 3
    protocol              = "HTTP_GET"
    http_get {
      path   = "/health"
      port   = 80
      scheme = "HTTPS"
      http_headers {
        header_name  = "authorization"
        header_value = "123456"
      }
    }
  }
}