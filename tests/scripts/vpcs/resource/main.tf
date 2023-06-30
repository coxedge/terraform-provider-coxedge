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

resource "coxedge_vpc" "vpc" {
  cidr             = "10.0.0.0/7"
  default_vpc      = false
  environment_name = "<environment_name>"
  name             = "<name>"
  organization_id  = "<organization_id>"
  slug             = "<slug>"
  status           = "ACTIVE"
}
