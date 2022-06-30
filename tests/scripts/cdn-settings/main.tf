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

resource "coxedge_cdn_settings" "testing" {
  site_id                           = "d348165c-3a8c-4904-8001-3e3d183a8643"
  environment_name                  = "test"
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