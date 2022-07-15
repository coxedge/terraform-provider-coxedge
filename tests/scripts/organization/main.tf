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

data "coxedge_organizations" "test" {
  id = "<organization_id>"
}

output "testing" {
  value = data.coxedge_organizations.test
}