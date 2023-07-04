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
//terraform import coxedge_route.route <route_id>:<vpc_id>:<environment_name>:<organization_id>
resource "coxedge_route" "route" {
  organization_id   = "<organization_id>"
  environment_name  = "<environment_name>"
  vpc_id            = "<vpc_id>"
  name              = "<name>"
  destination_cidrs = tolist(["<destination_cidrs>"])
  next_hops         = tolist(["<next_hops>"])
  status            = "ACTIVE"
}