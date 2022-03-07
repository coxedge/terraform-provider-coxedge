terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source = "coxedge.com/cox/coxedge"
    }
  }
}

provider "coxedge" {
  key = "VXl8IGsMCF5cC5SW3EMgpw=="
}
/*
data "coxedge_organizations" "test" {
}

output "org_test" {
  value = data.coxedge_organizations.test
}

data "coxedge_organizations" "test_with_id" {
  id = "eadbaec8-3221-4f29-8a03-0f5c6976afb6"
}
output "org_test_with_id" {
  value = data.coxedge_organizations.test_with_id
}
# Envs
resource "coxedge_environment" "test_env" {
  count = 0
  name = "test_env_created_by_tf"
  description = "This test env was created by terraform."
  organization_id = data.coxedge_organizations.test_with_id.id
  service_connection_id = data.coxedge_organizations.test_with_id.organizations[0].service_connections[1].id
}*/

data "coxedge_environments" "test" {

}

output "envs" {
  value = data.coxedge_environments.test
}

# Workloads
resource "coxedge_workload" "test" {
  name = "test"
  environment_name = data.coxedge_environments.test.environments[0].name
  type = "CONTAINER"
  image = "ubuntu:latest"
  specs = "SP-1"
  deployment {
    name = "test"
    enable_autoscaling = false
    pops = ["MIA"]
    instances_per_pop = 1
  }
}