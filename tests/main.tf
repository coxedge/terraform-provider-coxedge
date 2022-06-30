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

resource "coxedge_origin_setting" "testing" {
  site_id          = "352cdc1e-c071-49ad-bddd-371094880507"
  environment_name = "test-codecraft"
  host_header      = "www.cc789.com"
  origin {
    address = "cc.cox11.com"
  }
}

#resource "coxedge_origin_setting" "testing" {
#}
#
#output "sample_order" {
#  value = coxedge_origin_setting.testing
#}