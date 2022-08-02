terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source  = "coxedge.com/cox/coxedge"
    }
  }
}

provider "coxedge" {
}

data "coxedge_environments" "test" {

}

output "added_workload" {
  value = coxedge_workload.test.id
}

# Workloads
resource "coxedge_workload" "test" {
  name             = "demo-container"
  organization_id  = "<organization_id>"
  environment_name = "demo_env"
  type             = "CONTAINER"
  image            = "bitnami/nginx"
  specs            = "SP-2"
  persistent_storages {
    path = "/var/lib/data"
    size = 1000
  }
  deployment {
    name               = "test"
    enable_autoscaling = false
    pops               = ["LAS"]
    instances_per_pop  = 1
  }
}