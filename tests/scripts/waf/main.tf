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

resource "coxedge_waf_settings" "testing" {
  site_id                 = "fedb00cb-3b4c-4f1c-9cc9-f1a3712a29d8"
  domain                  = "rpgfan.com"
  environment_name        = "test-codecraft"
  monitoring_mode_enabled = false
  api_urls                = [
    "/test/this/ways123"
  ]
  allow_known_bots {
    internet_archive_bot = true
  }
  anti_automation_bot_protection {
    anti_scraping                                 = true
    challenge_automated_clients                   = false
    challenge_headless_browsers                   = false
    force_browser_validation_on_traffic_anomalies = true
  }
  behavioral_waf {
    block_probing_and_forced_browsing         = false
    bruteforce_protection                     = true
    obfuscated_attacks_and_zeroday_mitigation = true
    repeated_violations                       = true
    spam_protection                           = true
  }
  cms_protection {
    whitelist_drupal      = false
    whitelist_joomla      = false
    whitelist_magento     = false
    whitelist_modx        = false
    whitelist_origin_ip   = false
    whitelist_umbraco     = false
    whitelist_wordpress   = false
    wordpress_waf_ruleset = true
  }
  ddos_settings {
    burst_threshold           = 110
    global_threshold          = 500
    subsecond_burst_threshold = 50
  }
  general_policies {
    block_invalid_user_agents = true
    block_unknown_user_agents = false
    http_method_validation    = false
  }
  owasp_threats {
    apache_struts_exploit                  = false
    code_injection                         = false
    common_web_application_vulnerabilities = false
    csrf                                   = true
    local_file_inclusion                   = false
    open_redirect                          = false
    personal_identifiable_info             = false
    protocol_attack                        = true
    remote_file_inclusion                  = false
    sensitive_data_exposure                = true
    serverside_template_injection          = true
    shell_injection                        = false
    shell_shock_attack                     = true
    sql_injection                          = true
    webshell_execution_attempt             = false
    wordpress_waf_ruleset                  = false
    xml_external_entity                    = true
    xss_attack                             = true
  }
  traffic_sources {
    convicted_bot_traffic              = true
    external_reputation_block_list     = false
    traffic_from_suspicious_nat_ranges = false
    traffic_via_cdn                    = false
    via_hosting_services               = false
    via_proxy_networks                 = false
    via_tor_nodes                      = false
    via_vpn                            = true
  }
}