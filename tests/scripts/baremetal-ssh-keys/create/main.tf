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

resource "coxedge_baremetal_ssh_key" "sshkey" {
  environment_name = "<environment_name>"
  organization_id  = "<organization_id>"
  name             = "<name>"
  public_key       = "<public_key>"
}