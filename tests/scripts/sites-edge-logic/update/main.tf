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

resource "coxedge_sites_predefined_edge_logic" "edge_logic" {
  environment_name            = "<environment_name>"
  organization_id             = "<organization_id>"
  site_id                     = "<site_id>"
  force_www_enabled           = false
  pseudo_streaming_enabled    = false
  robots_txt_enabled          = false
  robots_txt_file             = ""
  referrer_protection_enabled = true
  referrer_list               = tolist(["<referrer_list>"])
  allow_empty_referrer        = false

}