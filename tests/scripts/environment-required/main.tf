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

data "coxedge_organizations" "orgs" {

}

resource "coxedge_environment" "demo_env" {
  name = "demo_env"
  description = "Terra Test Env"
  service_connection_id = "a572df45-56fa-4521-8a66-b63b5ab19c21"
  organization_id = "f1e0a327-dbf6-46c8-967d-bcd381c7c531"
}

output "env_id" {
 value = coxedge_environment.demo_env.id
}