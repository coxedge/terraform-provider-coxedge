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

#to import existing resource, do run below script
#terraform import coxedge_cdn_settings.testing <site_id>:<environment_name>:<organization_id>

resource "coxedge_cdn_settings" "testing" {
  organization_id                   = "<organization_id>"
  site_id                           = "<site_id>"
  environment_name                  = "<environment_name>"
  cache_expire_policy               = "SPECIFY_CDN_TTL"
  cache_ttl                         = "60"
  custom_cached_query_strings       = ["customQuery"]
  dynamic_caching_by_header_enabled = "true"
  custom_cached_headers             = ["mydynamicheader"]
  gzip_compression_enabled          = "true"
  gzip_compression_level            = "2"
  content_persistence_enabled       = "true"
  maximum_stale_file_ttl            = "30"
  vary_header_enabled               = "true"
  browser_cache_ttl                 = "60"
  cors_header_enabled               = "true"
  allowed_cors_origins              = "SPECIFY_ORIGINS"
  origins_to_allow_cors             = ["www.sai.com"]
  http2_support_enabled             = "true"
  link_header                       = "my-link-header"

  canonical_header_enabled = "true"
  canonical_header         = "my-canonical-header"
  url_caching_enabled      = "true"
  url_caching_ttl          = "300"
}