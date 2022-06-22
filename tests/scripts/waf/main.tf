terraform {
  required_providers {
    coxedge = {
      version = "0.1"
      source = "coxedge.com/cox/coxedge"
    }
  }
}

provider "coxedge" {
  key = "[INSERT API KEY HERE]"
}

resource "coxedge_waf_settings" "testing" {
  environment_name = "test"
  site_id = "fedb00cb-3b4c-4f1c-9cc9-f1a3712a29d8"
  domain = "rpgfan.com"
  api_urls = ["/test/this/way"]
  monitoring_enabled = "false"
  spam_and_abuse_form = "false"
  csrf = "true"

  ddos_settings {
    global_threshold = "5000"
    burst_threshold = "110"
    subsecond_burst_threshold = "50"
  }
  owasp_threats {
    sql_injection = "false"
    xss_attack = "true"
    remote_file_inclusion = "true"
    wordpress_waf_ruleset = "true"
    apache_struts_exploit = "true"
    local_file_inclusion = "false"
    common_web_application_vulnerabilities = "true"
    webshell_execution_attempt = "true"
    response_header_injections = "true"
    open_redirect = "false"
    shell_injection = "false"
  }
  user_agents {
    block_invalid_user_agents = "false"
    block_unknown_user_agents = "true"
  }
  traffic_sources {
    via_tor_nodes = "true"
    via_proxy_networks = "true"
    via_hosting_services = "true"
    via_vpn = "true"
    convicted_bot_traffic = "true"
    suspicious_traffic_by_local_ip_format = "true"
  }
  anti_automation_bot_protection {
    force_browser_validation_on_traffic_anomalies = "true"
    challenge_automated_clients = "false"
    challenge_headless_browsers = "false"
    anti_scraping = "false"
  }
  behavioral_waf {
    spam_protection = "true"
    block_probing_and_forced_browsing = "true"
    obfuscated_attacks_and_zeroday_mitigation = "true"
    repeated_violations = "true"
    bruteforce_protection = "true"
  }
  cms_protection {
    whitelist_wordpress = "false"
    whitelist_modx = "false"
    whitelist_drupal = "false"
    whitelist_joomla = "false"
    whitelist_magneto = "false"
    whitelist_origin_ip = "true"
    whitelist_umbraco = "false"
  }
  allow_known_bots {
    internet_archive_bot = "true"
  }
}