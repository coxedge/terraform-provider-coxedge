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

data "coxedge_compute_snapshots" "snapshots" {
  organization_id  = "<organization_id>"
  environment_name = "<environment name>"
  snapshot_id      = "<snapshot_id>"
}

output "output_tag" {
  value = data.coxedge_compute_snapshots.snapshots
}