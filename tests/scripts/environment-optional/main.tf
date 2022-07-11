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

output "env_id" {
 value = coxedge_environment.demo_env.id
}

resource "coxedge_environment" "demo_env" {
  name = "demo_env"
  description = "New Terra Test Env"
  membership = "ALL_ORG_USERS"
  service_connection_id = "<service-connection-id>"
  organization_id = "<organization-id>"
}