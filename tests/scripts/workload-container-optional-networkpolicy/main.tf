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

resource "coxedge_workload" "test" {
  name             = "demo-container-1"
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  type             = "CONTAINER"
  image            = "bitnami/nginx"
  specs            = "SP-2"
  deployment {
    name               = "test"
    enable_autoscaling = false
    pops               = ["LAS"]
    instances_per_pop  = 1
  }
  ports {
    protocol         = "TCP"
    public_port      = "80"
    public_port_desc = "Description"
    public_port_src  = "0.0.0.0/0"
  }
}
