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

data "coxedge_baremetal_device_charts" "charts" {
  id               = "<id>"
  environment_name = "<environment_name>"
  organization_id  = "<organization_id>"
  custom           = true
  start_date       = "<unix_epoch_time>"
  end_date         = "<unix_epoch_time>"
}

output "charts_output" {
  value = data.coxedge_baremetal_device_charts.charts
}