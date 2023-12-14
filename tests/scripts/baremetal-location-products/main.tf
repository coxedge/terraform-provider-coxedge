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

data "coxedge_baremetal_location_products" "products" {
  environment_name = "<environment_name>"
  organization_id  = "<organization_id>"
  id               = "<location_id>"
  code             = "<location_code>"
}

output "output_disk" {
  value = data.coxedge_baremetal_location_products.products
}