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

data "coxedge_workload_instances" "test" {
  organization_id  = "b0d424e4-4f78-4cb3-8c7c-26781bea9f7e"
  environment_name = "test-codecraft"
  id               = "3fd19d97-22bf-40a4-a615-6436f4714633"
}

output "testing" {
  value = data.coxedge_workload_instances.test
}