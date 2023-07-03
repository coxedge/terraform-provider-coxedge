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

//for importing
//terraform import coxedge_subnet.subnet <subnet_id>:<vpc_id>:<environment_name>:<organization_id>
resource "coxedge_subnet" "subnet" {
  environment_name = "<environment_name>"
  organization_id  = "<organization_id>"
  name             = "<name>"
  slug             = "<slug>"
  vpc_id           = "<vpc_id>"
  cidr             = "<cidr>"
}