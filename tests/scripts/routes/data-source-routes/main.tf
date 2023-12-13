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

data "coxedge_routes" "route" {
  environment_name = "<environment_name>"
  organization_id  = "<organization_id>"
  vpc_id           = "<vpc_id>"
}

output "output_routes" {
  value = data.coxedge_routes.route
}