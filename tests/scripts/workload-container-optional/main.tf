terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source = "coxedge.com/cox/coxedge"
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
  name = "demo-container"
  environment_name = "demo_env"
  type = "CONTAINER"
  image = "ubuntu:latest"
  specs = "SP-2"
  commands = [
    "/bin/sh -c \"sleep 50\""
  ]
  persistent_storages = ""
  deployment {
    name = "test"
    enable_autoscaling = false
    pops = ["PVD"]
    instances_per_pop = 2
  }
}