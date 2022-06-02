/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package coxedge

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func getOrganizationSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"organizations": &schema.Schema{
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: getOrganizationSchema(),
			},
		},
	}
}

func getOrganizationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"entry_point": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"tags": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"service_connections": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:     schema.TypeString,
						Required: true,
					},
					"name": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"service_code": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
	}
}

func getEnvironmentSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"environments": &schema.Schema{
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: getEnvironmentSchema(),
			},
		},
	}
}

func getEnvironmentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"membership": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"organization_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"service_connection_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"creation_date": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"roles": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"is_default": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  false,
					},
					"users": {
						Type:     schema.TypeList,
						Required: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
				},
			},
		},
	}
}

func getUserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"user_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"first_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"last_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"email": {
			Type:     schema.TypeString,
			Required: true,
		},
		"organization_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"roles": {
			Type: schema.TypeList,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"id": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
			Optional: true,
		},
		"last_updated": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func getWorkloadSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"environment_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"image": {
			Type:     schema.TypeString,
			Required: true,
		},
		"specs": {
			Type:     schema.TypeString,
			Required: true,
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"deployment": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"pops": {
						Type: schema.TypeList,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
						Required: true,
					},
					"enable_autoscaling": {
						Type:     schema.TypeBool,
						Default:  false,
						Optional: true,
					},
					"instances_per_pop": {
						Type:     schema.TypeInt,
						Optional: true,
						Default:  -1,
					},
					"max_instances_per_pop": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"min_instances_per_pop": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"cpu_utilization": {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
		"add_anycast_ip_address": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"anycast_ip_address": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"commands": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
		"container_email": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"container_username": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"container_password": {
			Type:      schema.TypeString,
			Sensitive: true,
			Optional:  true,
		},
		"container_server": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"environment_variables": {
			Type:     schema.TypeMap,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		"first_boot_ssh_key": {
			Type:     schema.TypeString,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		"ports": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"protocol": {
						Type:     schema.TypeString,
						Required: true,
					},
					"public_port": {
						Type:     schema.TypeString,
						Required: true,
					},
					"public_port_desc": {
						Type:     schema.TypeString,
						Required: true,
					},
					"public_port_src": {
						Type:     schema.TypeString,
						Required: true,
					},
				},
			},
		},
		"persistent_storage": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"path": {
						Type:     schema.TypeString,
						Required: true,
					},
					"size": {
						Type:     schema.TypeInt,
						Required: true,
					},
				},
			},
		},
		"secret_environment_variables": {
			Type:     schema.TypeMap,
			Elem:     &schema.Schema{Type: schema.TypeString},
			Optional: true,
		},
		"slug": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func getImageSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"stack_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"slug": {
			Type:     schema.TypeString,
			Required: true,
		},
		"family": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tag": {
			Type:     schema.TypeString,
			Required: true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Required: true,
		},
		"description": {
			Type:     schema.TypeString,
			Required: true,
		},
		"reference": {
			Type:     schema.TypeString,
			Required: true,
		},
		"status": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func getImageSetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
			Optional: true,
		},
		"environment": {
			Type:     schema.TypeString,
			Required: true,
		},
		"images": &schema.Schema{
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: getImageSchema(),
			},
		},
	}
}

func getNetworkPolicyRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"stack_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"environment_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"workload_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"network_policy_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"source": {
			Type:     schema.TypeString,
			Required: true,
		},
		"action": {
			Type:     schema.TypeString,
			Required: true,
		},
		"protocol": {
			Type:     schema.TypeString,
			Required: true,
		},
		"port_range": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func getSiteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"environment_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"services": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required: true,
		},
		"protocol": {
			Type:     schema.TypeString,
			Required: true,
		},
		"domain": {
			Type:     schema.TypeString,
			Required: true,
		},
		"hostname": {
			Type:     schema.TypeString,
			Required: true,
		},
		"auth_method": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"username": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"password": {
			Type:     schema.TypeString,
			Optional: true,
		},
		//Computed properties
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"stack_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"edge_address": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"anycast_ip": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"delivery_domains": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"domain": {
						Type:     schema.TypeString,
						Computed: true,
						ForceNew: true,
					},
					"validated_at": {
						Type:     schema.TypeString,
						Computed: true,
						ForceNew: true,
					},
				},
			},
		},
	}
}

func getOriginSettingOriginSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"address": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"common_certificate_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"auth_method": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"username": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"password": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func getOriginSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"stack_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"scope_configuration_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"environment_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"domain": {
			Type:     schema.TypeString,
			Required: true,
		},
		"websockets_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"ssl_validation_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"pull_protocol": {
			Type:     schema.TypeString,
			Required: true,
		},
		"host_header": {
			Type:     schema.TypeString,
			Required: true,
		},
		"origin": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: getOriginSettingOriginSchema(),
			},
			Required: true,
		},
		"backup_origin_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"backup_origin": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: getOriginSettingOriginSchema(),
			},
			Optional: true,
		},
		"backup_origin_exclude_codes": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
	}
}

func getDeliveryDomainSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"stack_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"environment_name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"domain": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
	}
}

func getCDNSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"site_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"environment_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"cache_expire_policy": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"cache_ttl": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"query_control_string": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"custom_cached_query_strings": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
		"dynamic_caching_by_header_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"custom_cached_headers": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
		"gzip_compression_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"gzip_compression_level": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"content_persistence_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"maximum_stale_file_ttl": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"vary_header_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"browser_cache_ttl": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"cors_header_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"allowed_cors_origins": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"origins_to_allow_cors": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
		"http2_support_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"http2_server_push_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"link_header": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"canonical_header_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"canonical_header": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"url_caching_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"url_caching_ttl": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

func getCDNPurgeResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"site_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"environment_name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"purge_type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "URL",
			ForceNew: true,
		},
		"items": {
			Type:     schema.TypeList,
			Optional: true,
			ForceNew: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"url": {
						Type:     schema.TypeString,
						Required: true,
					},
					"recursive": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"invalidate_only": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"purge_all_dynamic": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"headers": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"purge_selector": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"selector_name": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"selector_type": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"selector_value": {
									Type:     schema.TypeString,
									Optional: true,
								},
								"selector_value_delimiter": {
									Type:     schema.TypeString,
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func getWAFSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"environment_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"site_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"domain": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"api_urls": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
		"ddos_settings": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"global_threshold": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"burst_threshold": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"subsecond_burst_threshold": {
						Type:     schema.TypeInt,
						Optional: true,
					},
				},
			},
		},
		"monitoring_enabled": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"owasp_threats": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"sql_injection": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"xss_attack": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"remote_file_inclusion": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"wordpress_waf_ruleset": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"apache_struts_exploit": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"local_file_inclusion": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"common_web_application_vulnerabilities": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"webshell_execution_attempt": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"response_header_injections": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"open_redirect": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"shell_injection": {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
		"user_agents": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"block_invalid_user_agents": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"block_unknown_user_agents": {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
		"csrf": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"traffic_sources": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"via_tor_nodes": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"via_proxy_networks": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"via_hosting_services": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"via_vpn": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"convicted_bot_traffic": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"suspicious_traffic_by_local_ip_format": {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
		"anti_automation_bot_protection": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"force_browser_validation_on_traffic_anomalies": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"challenge_automated_clients": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"challenge_headless_browsers": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"anti_scraping": {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
		"spam_and_abuse_form": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"behavioral_waf": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"spam_protection": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"block_probing_and_forced_browsing": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"obfuscated_attacks_and_zeroday_mitigation": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"repeated_violations": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"bruteforce_protection": {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
		"cms_protection": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"whitelist_wordpress": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"whitelist_modx": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"whitelist_drupal": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"whitelist_joomla": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"whitelist_magneto": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"whitelist_origin_ip": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"whitelist_umbraco": {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
		"allow_known_bots": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"internet_archive_bot": {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
	}
}

func getFirewallRuleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"environment_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"site_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"action": {
			Type:     schema.TypeString,
			Required: true,
		},
		"ip_start": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"ip_end": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func getScriptSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"stack_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"site_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"updated_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"version": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"code": {
			Type:     schema.TypeString,
			Required: true,
		},
		"routes": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Required: true,
		},
	}
}
