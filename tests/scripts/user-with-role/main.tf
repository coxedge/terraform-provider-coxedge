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

resource "coxedge_user" "testuser" {
  user_name = "nisai"
  first_name = "Nishanth"
  last_name = "Bhosale"
  email = "sai-nishanth.bhosle@capgemini.com"
  organization_id = "f1e0a327-dbf6-46c8-967d-bcd381c7c531"
  roles {
    id = "7c5c2fb3-53ae-408a-a97b-717da8758348"
  }
}

output "added_user" {
  value = coxedge_user.testuser
}