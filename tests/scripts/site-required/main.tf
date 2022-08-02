terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source  = "coxedge.com/cox/coxedge"
    }
  }
}

provider "coxedge" {
}

resource "coxedge_site" "testing" {
  organization_id  = "<organization_id>"
  environment_name = "test"
  domain           = "www.rpgfan.com"
  hostname         = "199.250.204.212"
  services         = [
    "CDN",
    "SERVERLESS_EDGE_ENGINE",
    "WAF"
  ]
  protocol = "HTTPS"
  #  for enabling or disabling CDN, WAF or Serverless Scripting, we need to use "operation" argument
  #  "operation" argument values are - enable_cdn, disable_cdn, enable_waf, disable_waf, enable_scripts, disable_scripts
  #  operation = "enable_scripts"
}

output "site_id" {
  value = coxedge_site.testing.id
}