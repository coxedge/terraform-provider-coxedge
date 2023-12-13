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

#terraform import coxedge_baremetal_ssh_key.sshkey <ssh_key_d>:<environment_name>:<organization_id>
resource "coxedge_baremetal_ssh_key" "sshkey" {
}