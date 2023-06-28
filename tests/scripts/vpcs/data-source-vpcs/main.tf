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

data "coxedge_vpcs" "vpcs" {
  environment_name     = "<environment_name>"
  organization_id = "<organization_id>"
}

output "output_vpc" {
  value = data.coxedge_vpcs.vpcs
}