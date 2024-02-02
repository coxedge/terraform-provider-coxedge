terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source  = "coxedge.com/cox/coxedge"
    }
  }
}

provider "coxedge" {
  key = "[INSERT API KEY HERE]"
}

resource "coxedge_compute_workload_user_data" "user_data" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  workload_id      = "<workload_id>"
  operation        = "<operation>" #workload-on/workload-off/reinstall-workload/restart-workload
}

