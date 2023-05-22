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

resource "coxedge_sites_edge_logic_delivery_rule" "delivery_rule" {
  environment_name = "<environment_name>"
  organization_id  = "<organization_id>"
  site_id          = "<site_id>"
  name             = "<name>"
  conditions {
    trigger  = "<trigger>"
    operator = "<operator>"
    target   = "<target>"
  }

  actions {
    action_type               = "<action_type>"
    passphrase                = "<passphrase>"
    passphrase_field          = "<passphrase_field>"
    md5_token_field           = "<md5_token_field>"
    ttl_field                 = "<ttl_field>"
    ip_address_filter         = "<ip_address_filter"
    url_signature_path_length = "<url_signature_path_length>"
  }
}