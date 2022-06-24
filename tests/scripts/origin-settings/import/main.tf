terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source  = "coxedge.com/cox/coxedge"
    }
  }
}

provider "coxedge" {
  key = "SyJwzYaeVylQmtOrPJkq2A=="
}

#do import existing resource, do run below script
#terraform import coxedge_origin_setting.testing <site_id>:<environment_name>
resource "coxedge_origin_setting" "testing" {
}

output "sample_order" {
  value = coxedge_origin_setting.testing
}