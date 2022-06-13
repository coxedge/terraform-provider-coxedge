/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
)

func resourceWAFSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWAFSettingsCreate,
		ReadContext:   resourceWAFSettingsRead,
		UpdateContext: resourceWAFSettingsUpdate,
		DeleteContext: resourceWAFSettingsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getWAFSettingsSchema(),
	}
}

func resourceWAFSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Convert resource data to API object
	updatedWAFSettings := convertResourceDataToWAFSettingsCreateAPIObject(d)
	d.SetId(updatedWAFSettings.Id)

	//Run Update since you do not "create" these
	resourceWAFSettingsUpdate(ctx, d, m)

	return diags
}

func resourceWAFSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the resource ID
	resourceId := d.Get("site_id").(string)

	//Get the resource
	cdnSettings, err := coxEdgeClient.GetWAFSettings(d.Get("environment_name").(string), resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	convertWAFSettingsAPIObjectToResourceData(d, cdnSettings)

	return diags
}

func resourceWAFSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Convert resource data to API object
	updatedWAFSettings := convertResourceDataToWAFSettingsCreateAPIObject(d)

	//Call the API
	_, err := coxEdgeClient.UpdateWAFSettings(updatedWAFSettings.Id, updatedWAFSettings)
	fmt.Println("BEH")
	if err != nil {
		fmt.Println(err)
		return diag.FromErr(err)
	}

	//Set last_updated
	d.Set("last_updated", time.Now().Format(time.RFC850))

	return resourceWAFSettingsRead(ctx, d, m)
}

func resourceWAFSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func convertResourceDataToWAFSettingsCreateAPIObject(d *schema.ResourceData) apiclient.WAFSettings {
	//Create update cdnSettings struct
	updatedWAFSettings := apiclient.WAFSettings{
		EnvironmentName:             d.Get("environment_name").(string),
		Id:                          d.Get("site_id").(string),
		Domain:                      d.Get("domain").(string),
		MonitoringEnabled:           d.Get("monitoring_enabled").(bool),
		SpamAndAbuseForm:            d.Get("spam_and_abuse_form").(bool),
		Csrf:                        d.Get("csrf").(bool),
		DdosSettings:                apiclient.WAFDdosSettings{},
		OwaspThreats:                apiclient.WAFOwaspThreats{},
		UserAgents:                  apiclient.WAFUserAgents{},
		TrafficSources:              apiclient.WAFTrafficSources{},
		AntiAutomationBotProtection: apiclient.WAFAntiAutomationBotProtection{},
		BehavioralWaf:               apiclient.WAFBehavioralWaf{},
		CmsProtection:               apiclient.WAFCmsProtection{},
		AllowKnownBots:              apiclient.WAFAllowKnownBots{},
	}

	updatedWAFSettings.APIUrls = []string{}
	for _, val := range d.Get("api_urls").([]interface{}) {
		updatedWAFSettings.APIUrls = append(updatedWAFSettings.APIUrls, val.(string))
	}

	for _, entryRaw := range d.Get("ddos_settings").([]interface{}) {
		entry := entryRaw.(map[string]interface{})
		updatedWAFSettings.DdosSettings = apiclient.WAFDdosSettings{
			GlobalThreshold:         entry["global_threshold"].(int),
			BurstThreshold:          entry["burst_threshold"].(int),
			SubSecondBurstThreshold: entry["subsecond_burst_threshold"].(int),
		}
	}

	for _, entryRaw := range d.Get("owasp_threats").([]interface{}) {
		entry := entryRaw.(map[string]interface{})
		updatedWAFSettings.OwaspThreats = apiclient.WAFOwaspThreats{
			SQLInjection:                        entry["sql_injection"].(bool),
			XSSAttack:                           entry["xss_attack"].(bool),
			RemoteFileInclusion:                 entry["remote_file_inclusion"].(bool),
			WordpressWafRuleset:                 entry["wordpress_waf_ruleset"].(bool),
			ApacheStrutsExploit:                 entry["apache_struts_exploit"].(bool),
			LocalFileInclusion:                  entry["local_file_inclusion"].(bool),
			CommonWebApplicationVulnerabilities: entry["common_web_application_vulnerabilities"].(bool),
			WebShellExecutionAttempt:            entry["webshell_execution_attempt"].(bool),
			ResponseHeaderInjection:             entry["response_header_injections"].(bool),
			OpenRedirect:                        entry["open_redirect"].(bool),
			ShellInjection:                      entry["shell_injection"].(bool),
		}
	}

	for _, entryRaw := range d.Get("user_agents").([]interface{}) {
		entry := entryRaw.(map[string]interface{})
		updatedWAFSettings.UserAgents = apiclient.WAFUserAgents{
			BlockInvalidUserAgents: entry["block_invalid_user_agents"].(bool),
			BlockUnknownUserAgents: entry["block_unknown_user_agents"].(bool),
		}
	}

	for _, entryRaw := range d.Get("traffic_sources").([]interface{}) {
		entry := entryRaw.(map[string]interface{})
		updatedWAFSettings.TrafficSources = apiclient.WAFTrafficSources{
			ViaTorNodes:                      entry["via_tor_nodes"].(bool),
			ViaProxyNetworks:                 entry["via_proxy_networks"].(bool),
			ViaHostingServices:               entry["via_hosting_services"].(bool),
			ViaVpn:                           entry["via_vpn"].(bool),
			ConvictedBotTraffic:              entry["convicted_bot_traffic"].(bool),
			SuspiciousTrafficByLocalIPFormat: entry["suspicious_traffic_by_local_ip_format"].(bool),
		}
	}

	for _, entryRaw := range d.Get("anti_automation_bot_protection").([]interface{}) {
		entry := entryRaw.(map[string]interface{})
		updatedWAFSettings.AntiAutomationBotProtection = apiclient.WAFAntiAutomationBotProtection{
			ForceBrowserValidationOnTrafficAnomalies: entry["force_browser_validation_on_traffic_anomalies"].(bool),
			ChallengeAutomatedClients:                entry["challenge_automated_clients"].(bool),
			ChallengeHeadlessBrowsers:                entry["challenge_headless_browsers"].(bool),
			AntiScraping:                             entry["anti_scraping"].(bool),
		}
	}

	for _, entryRaw := range d.Get("behavioral_waf").([]interface{}) {
		entry := entryRaw.(map[string]interface{})
		updatedWAFSettings.BehavioralWaf = apiclient.WAFBehavioralWaf{
			SpamProtection:                        entry["spam_protection"].(bool),
			BlockProbingAndForcedBrowsing:         entry["block_probing_and_forced_browsing"].(bool),
			ObfuscatedAttacksAndZeroDayMitigation: entry["obfuscated_attacks_and_zeroday_mitigation"].(bool),
			RepeatedViolations:                    entry["repeated_violations"].(bool),
			BruteForceProtection:                  entry["bruteforce_protection"].(bool),
		}
	}

	for _, entryRaw := range d.Get("cms_protection").([]interface{}) {
		entry := entryRaw.(map[string]interface{})
		updatedWAFSettings.CmsProtection = apiclient.WAFCmsProtection{
			WhiteListWordpress: entry["whitelist_wordpress"].(bool),
			WhiteListModx:      entry["whitelist_modx"].(bool),
			WhiteListDrupal:    entry["whitelist_drupal"].(bool),
			WhiteListJoomla:    entry["whitelist_joomla"].(bool),
			WhiteMagento:       entry["whitelist_magneto"].(bool),
			WhiteListOriginIP:  entry["whitelist_origin_ip"].(bool),
			WhiteListUmbraco:   entry["whitelist_umbraco"].(bool),
		}
	}

	for _, entryRaw := range d.Get("allow_known_bots").([]interface{}) {
		entry := entryRaw.(map[string]interface{})
		updatedWAFSettings.AllowKnownBots = apiclient.WAFAllowKnownBots{
			InternetArchiveBot: entry["internet_archive_bot"].(bool),
		}
	}

	return updatedWAFSettings
}

func convertWAFSettingsAPIObjectToResourceData(d *schema.ResourceData, wafSettings *apiclient.WAFSettings) {
	//Store the data
	d.Set("environment_name", wafSettings.EnvironmentName)
	d.Set("site_id", wafSettings.Id)
	d.Set("stack_id", wafSettings.StackId)
	d.Set("domain", wafSettings.Domain)
	d.Set("api_urls", wafSettings.APIUrls)
	d.Set("monitoring_enabled", wafSettings.MonitoringEnabled)
	d.Set("spam_and_abuse_form", wafSettings.SpamAndAbuseForm)
	d.Set("csrf", wafSettings.Csrf)

	ddosSettings := make(map[string]int)
	ddosSettings["global_threshold"] = wafSettings.DdosSettings.GlobalThreshold
	ddosSettings["burst_threshold"] = wafSettings.DdosSettings.BurstThreshold
	ddosSettings["subsecond_burst_threshold"] = wafSettings.DdosSettings.SubSecondBurstThreshold
	d.Set("ddos_settings", ddosSettings)

	owaspThreats := make(map[string]bool)
	owaspThreats["sql_injection"] = wafSettings.OwaspThreats.SQLInjection
	owaspThreats["xss_attack"] = wafSettings.OwaspThreats.XSSAttack
	owaspThreats["remote_file_inclusion"] = wafSettings.OwaspThreats.RemoteFileInclusion
	owaspThreats["wordpress_waf_ruleset"] = wafSettings.OwaspThreats.WordpressWafRuleset
	owaspThreats["apache_struts_exploit"] = wafSettings.OwaspThreats.ApacheStrutsExploit
	owaspThreats["local_file_inclusion"] = wafSettings.OwaspThreats.LocalFileInclusion
	owaspThreats["common_web_application_vulnerabilities"] = wafSettings.OwaspThreats.CommonWebApplicationVulnerabilities
	owaspThreats["webshell_execution_attempt"] = wafSettings.OwaspThreats.WebShellExecutionAttempt
	owaspThreats["response_header_injection"] = wafSettings.OwaspThreats.ResponseHeaderInjection
	owaspThreats["open_redirect"] = wafSettings.OwaspThreats.OpenRedirect
	owaspThreats["shell_injection"] = wafSettings.OwaspThreats.ShellInjection
	d.Set("owasp_threats", owaspThreats)

	userAgents := make(map[string]bool)
	userAgents["block_invalid_user_agents"] = wafSettings.UserAgents.BlockInvalidUserAgents
	userAgents["block_unknown_user_agents"] = wafSettings.UserAgents.BlockUnknownUserAgents
	d.Set("user_agents", userAgents)

	trafficSources := make(map[string]bool)
	trafficSources["via_tor_nodes"] = wafSettings.TrafficSources.ViaTorNodes
	trafficSources["via_proxy_networks"] = wafSettings.TrafficSources.ViaProxyNetworks
	trafficSources["via_hosting_services"] = wafSettings.TrafficSources.ViaHostingServices
	trafficSources["via_vpn"] = wafSettings.TrafficSources.ViaVpn
	trafficSources["convicted_bot_traffic"] = wafSettings.TrafficSources.ConvictedBotTraffic
	trafficSources["suspicious_traffic_by_local_ip_format"] = wafSettings.TrafficSources.SuspiciousTrafficByLocalIPFormat
	d.Set("traffic_sources", trafficSources)

	antiAutomationBotProtection := make(map[string]bool)
	antiAutomationBotProtection["force_browser_validation_on_traffic_anomalies"] = wafSettings.AntiAutomationBotProtection.ForceBrowserValidationOnTrafficAnomalies
	antiAutomationBotProtection["challenge_automated_clients"] = wafSettings.AntiAutomationBotProtection.ChallengeAutomatedClients
	antiAutomationBotProtection["challenge_headless_browsers"] = wafSettings.AntiAutomationBotProtection.ChallengeHeadlessBrowsers
	antiAutomationBotProtection["anti_scraping"] = wafSettings.AntiAutomationBotProtection.AntiScraping
	d.Set("anti_automation_bot_protection", antiAutomationBotProtection)

	behavioralWaf := make(map[string]bool)
	behavioralWaf["spam_protection"] = wafSettings.BehavioralWaf.SpamProtection
	behavioralWaf["block_probing_and_forced_browsing"] = wafSettings.BehavioralWaf.BlockProbingAndForcedBrowsing
	behavioralWaf["obfuscated_attacks_and_zeroday_mitigation"] = wafSettings.BehavioralWaf.ObfuscatedAttacksAndZeroDayMitigation
	behavioralWaf["repeated_violations"] = wafSettings.BehavioralWaf.RepeatedViolations
	behavioralWaf["bruteforce_protection"] = wafSettings.BehavioralWaf.BruteForceProtection
	d.Set("behavioral_waf", behavioralWaf)

	cmsProtection := make(map[string]bool)
	cmsProtection["whitelist_wordpress"] = wafSettings.CmsProtection.WhiteListWordpress
	cmsProtection["whitelist_modx"] = wafSettings.CmsProtection.WhiteListModx
	cmsProtection["whitelist_drupal"] = wafSettings.CmsProtection.WhiteListDrupal
	cmsProtection["whitelist_joomla"] = wafSettings.CmsProtection.WhiteListJoomla
	cmsProtection["whitelist_magento"] = wafSettings.CmsProtection.WhiteMagento
	cmsProtection["whitelist_origin_ip"] = wafSettings.CmsProtection.WhiteListOriginIP
	cmsProtection["whitelist_umbraco"] = wafSettings.CmsProtection.WhiteListUmbraco
	d.Set("cms_protection", cmsProtection)

	allowKnownBots := make(map[string]bool)
	allowKnownBots["internet_archive_bot"] = wafSettings.AllowKnownBots.InternetArchiveBot
	d.Set("allow_known_bots", allowKnownBots)
}
