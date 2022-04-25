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

resource "coxedge_site" "testing" {
    environment_name = "test"
    domain = "www.rpgfan.com"
    hostname = "199.250.204.212"
    services = [
      "CDN",
      "SERVERLESS_EDGE_ENGINE",
      "WAF"
    ]
    protocol = "HTTPS"
}

output "site_id" {
  value = coxedge_site.testing.id
}