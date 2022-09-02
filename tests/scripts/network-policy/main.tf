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

resource "coxedge_network_policy_rule" "testing" {
  organization_id  = "<organization id>"
  environment_name = "test-codecraft"
  network_policy {
    workload_id = "<workload Id>"
    # add id for update
    #    id          = "5a1839d3-234c-4a1b-b5fa-82e5dfd2ec36/INBOUND/1080815577/0"
    description = "inbound-1-s"
    protocol    = "TCP"
    type        = "INBOUND"
    action      = "ALLOW"
    source      = "0.0.0.0/32"
    port_range  = "30000-33001"
  }
  network_policy {
    workload_id = "<workload Id>"
    # add id for update
    #id          = "5a1839d3-234c-4a1b-b5fa-82e5dfd2ec36/INBOUND/-2053391905/0"
    description = "inbound-2-update"
    protocol    = "TCP"
    type        = "INBOUND"
    action      = "ALLOW"
    source      = "0.0.0.0/2"
    port_range  = "30000-33001"
  }
}

output "policy" {
  value = coxedge_network_policy_rule.testing.network_policy
}