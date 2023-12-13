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

data "coxedge_baremetal_ssh_keys" "sshkeys" {
  id               = "<ssh-key-id>"
  environment_name = "<environment_name>"
  organization_id  = "<organization_id>"
}

output "sshkeys_output" {
  value = data.coxedge_baremetal_ssh_keys.sshkeys
}