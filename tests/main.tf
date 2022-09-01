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

resource "coxedge_network_policy_rule" "testing" {
  organization_id  = "b0d424e4-4f78-4cb3-8c7c-26781bea9f7e"
  environment_name = "test-codecraft"
  network_policy {
    workload_id = "3fd19d97-22bf-40a4-a615-6436f4714633"
    description = "inbound-1"
    protocol    = "TCP"
    type        = "INBOUND"
    action      = "ALLOW"
    source      = "0.0.0.0/32"
    port_range  = "30000-33001"
  }
  network_policy {
    workload_id = "3fd19d97-22bf-40a4-a615-6436f4714633"
    description = "inbound-2"
    protocol    = "TCP"
    type        = "INBOUND"
    action      = "ALLOW"
    source      = "0.0.0.0/2"
    port_range  = "30000-33001"
  }
  network_policy {
    workload_id = "3fd19d97-22bf-40a4-a615-6436f4714633"
    description = "outbound-1"
    protocol    = "TCP"
    type        = "OUTBOUND"
    action      = "ALLOW"
    source      = "0.0.0.0/0"
    port_range  = "80"
  }
  network_policy {
    workload_id = "3fd19d97-22bf-40a4-a615-6436f4714633"
    description = "outbound-2"
    protocol    = "TCP"
    type        = "OUTBOUND"
    action      = "ALLOW"
    source      = "0.0.0.0/32"
    port_range  = "80"
  }
}

output "policy_id" {
  value = coxedge_network_policy_rule.testing.id
}