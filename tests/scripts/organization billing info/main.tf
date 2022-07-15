terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source  = "coxedge.com/cox/coxedge"
    }
  }
}

# test account apikey and org id to get billing info in UAT
# usr: terraformtester002@getnada.com/ pwd: Codecraft@123
provider "coxedge" {
  key = "Xt93vyn12lJEgXJz4Q/kbw=="
}

data "coxedge_organizations_billing_info" "test" {
  id = "899c1ef6-6dc5-49f6-9663-3512e12d6e3d"
}

output "testing" {
  value = data.coxedge_organizations_billing_info.test
}