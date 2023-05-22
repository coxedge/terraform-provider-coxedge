terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source  = "coxedge.com/cox/coxedge"
    }
  }
}

provider "coxedge" {
  key = "GM3COPLOU6nOI12/NZ7HNg=="
}

#data "coxedge_sites_predefined_edge_logic" "edge_logic" {
#  id = "afd792e6-2764-4802-85ce-9b0a0361a3e1"
#  environment_name = "test-backend"
#  organization_id = "b0d424e4-4f78-4cb3-8c7c-26781bea9f7e"
#}
#
#output "output_edge_logic" {
#  value = data.coxedge_sites_predefined_edge_logic.edge_logic
#}

#resource "coxedge_sites_predefined_edge_logic" "edge_logic" {
#  environment_name            = "test-backend"
#  organization_id             = "b0d424e4-4f78-4cb3-8c7c-26781bea9f7e"
#  site_id                     = "afd792e6-2764-4802-85ce-9b0a0361a3e1"
#  force_www_enabled           = false
#  pseudo_streaming_enabled    = false
#  robots_txt_enabled          = false
#  robots_txt_file             = ""
#  referrer_protection_enabled = true
#  referrer_list               = tolist(["listadd"])
#  allow_empty_referrer        = false
#
#}

#resource "coxedge_sites_edge_logic_delivery_rule" "delivery_rule" {
#  environment_name = "test-backend"
#  organization_id  = "b0d424e4-4f78-4cb3-8c7c-26781bea9f7e"
#  site_id          = "afd792e6-2764-4802-85ce-9b0a0361a3e1"
#  name             = "terr-rule1"
#  conditions {
#    trigger  = "URL"
#    operator = "MATCHES"
#    target   = "codecraft"
#  }
#
#  actions {
#    action_type               = "SIGN_URL"
#    passphrase                = "code"
#    passphrase_field          = "code"
#    md5_token_field           = "code"
#    ttl_field                 = "code"
#    ip_address_filter         = "100.0.0.7"
#    url_signature_path_length = "10"
#  }
#
#}

#data "coxedge_sites_edge_logic_delivery_rules" "delivery_rule" {
#  environment_name = "test-backend"
#  organization_id  = "b0d424e4-4f78-4cb3-8c7c-26781bea9f7e"
#  id          = "afd792e6-2764-4802-85ce-9b0a0361a3e1"
#}
#
#output "delivery_rules_output" {
#  value = data.coxedge_sites_edge_logic_delivery_rules.delivery_rule
#}

data "coxedge_sites_edge_logic_custom_rules" "custom_rules" {
  environment_name = "test-backend"
  organization_id  = "b0d424e4-4f78-4cb3-8c7c-26781bea9f7e"
  id               = "afd792e6-2764-4802-85ce-9b0a0361a3e1"
}

output "delivery_rules_output" {
  value = data.coxedge_sites_edge_logic_custom_rules.custom_rules
}