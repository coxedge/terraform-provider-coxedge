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
  organization_id = "eadbaec8-3221-4f29-8a03-0f5c6976afb6"
}

output "added_user" {
  value = coxedge_user.testuser
}