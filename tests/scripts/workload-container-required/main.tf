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

output "envs" {
  value = data.coxedge_environments.test
}
//to import existing workload which is created outside terraform script and to update state file.
//terraform import coxedge_workload.test <workload_id>:<environment_name>:<organization_id>

//to sync with UI changes.
//terraform apply --refresh-only

# Workloads
resource "coxedge_workload" "test" {
  name             = "demo-container-2"
  organization_id  = "<organization_id>"
  environment_name = "demo_env"
  type             = "CONTAINER"
  image            = "bitnami/nginx"
  specs            = "SP-2"
  deployment {
    name               = "test"
    enable_autoscaling = false
    pops               = ["BTR"]
    instances_per_pop  = 1
  }
}