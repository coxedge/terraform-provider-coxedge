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

resource "coxedge_workload" "test" {
  name               = "terr-test1"
  organization_id    = "96a582f8-2301-46f3-8200-3fb6afb64e69"
  environment_name   = "anothercleanenv"
  type               = "VM"
  image              = "stackpath-edge/ubuntu-2004-focal:v202102241556"
  first_boot_ssh_key = "<first_boot_ssh_key>"
  specs              = "SP-1"
  deployment {
    name               = "testvm"
    enable_autoscaling = false
    pops               = ["LAS"]
    instances_per_pop  = 1
  }
  probe_configuration = "LIVENESS_AND_READINESS"
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
  readiness_probe {
    initial_delay_seconds = 0
    timeout_seconds       = 1
    period_seconds        = 10
    success_threshold     = 1
    failure_threshold     = 3
    protocol              = "HTTP_GET"
    http_get {
      path   = "/ping"
      port   = 80
      scheme = "HTTP"
      http_headers {
        header_name  = "authorization"
        header_value = "123456"
      }
    }
  }
}