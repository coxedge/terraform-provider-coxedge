terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source  = "coxedge.com/cox/coxedge"
    }
  }
}

provider "coxedge" {
  key = "[INSERT API KEY HERE]"
}

data "coxedge_environments" "test" {

}

output "envs" {
  value = data.coxedge_environments.test
}

# VM Workloads
resource "coxedge_workload" "test" {
  name               = "<name>"
  organization_id    = "<organization_id>"
  environment_name   = "<environment_name>"
  type               = "VM"
  image              = "stackpath-edge/ubuntu-1804-bionic:v202104291427"
  first_boot_ssh_key = "<first_boot_ssh_key>"
  specs              = "SP-1"
  deployment {
    name               = "testvm"
    enable_autoscaling = false
    pops               = ["PVD"]
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
  network_interfaces {
    vpc_slug     = "default"
    ip_families  = "IPv4"
    subnet_slug  = ""
    is_public_ip = true
  }
}
