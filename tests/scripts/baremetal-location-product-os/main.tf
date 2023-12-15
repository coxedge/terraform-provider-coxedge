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

data "coxedge_baremetal_location_product_os" "operating_systems" {
  environment_name  = "<environment_name>"
  organization_id   = "<organization_id>"
  vendor_product_id = "<vendor_product_id>"
}

output "output_os" {
  value = data.coxedge_baremetal_location_product_os.operating_systems
}