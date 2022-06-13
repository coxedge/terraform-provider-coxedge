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

output "envs" {
  value = data.coxedge_environments.test
}

# Workloads
resource "coxedge_workload" "test" {
  name = "demo-container-2"
  environment_name = "demo_env"
  type = "CONTAINER"
  image = "ubuntu:latest"
  specs = "SP-2"
  deployment {
    name = "test"
    enable_autoscaling = false
    pops = ["BTR"]
    instances_per_pop = 3
  }
}