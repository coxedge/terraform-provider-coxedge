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

data "coxedge_compute_reserved_ips" "reserved_ip" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  reserved_ip_id   = "<reserved_ip_id>"
}

output "output_storage" {
  value = data.coxedge_compute_reserved_ips.reserved_ip
}