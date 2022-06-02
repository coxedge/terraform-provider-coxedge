/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
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
	resourceId := d.Id()

	//Get the resource
	cdnSettings, err := coxEdgeClient.GetWAFSettings(d.Get("environment_name").(string), resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	convertWAFSettingsAPIObjectToResourceData(d, cdnSettings)

	//Update state
	resourceWAFSettingsRead(ctx, d, m)

	return diags
}

func resourceWAFSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource ID
	resourceId := d.Id()

	//Convert resource data to API object
	updatedWAFSettings := convertResourceDataToWAFSettingsCreateAPIObject(d)

	//Call the API
	_, err := coxEdgeClient.UpdateWAFSettings(resourceId, updatedWAFSettings)
	if err != nil {
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
		Id:                          d.Get("id").(string),
		StackId:                     d.Get("stack_id").(string),
		Domain:                      d.Get("domain").(string),
		APIUrls:                     d.Get("api_urls").([]string),
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

	for _, entry := range d.Get("ddos_settings").([]map[string]int) {
		updatedWAFSettings.DdosSettings = apiclient.WAFDdosSettings{
			GlobalThreshold:         entry["global_threshold"],
			BurstThreshold:          entry["burst_threshold"],
			SubSecondBurstThreshold: entry["subsecond_burst_threshold"],
		}
	}

	for _, entry := range d.Get("owasp_threats").([]map[string]bool) {
		updatedWAFSettings.OwaspThreats = apiclient.WAFOwaspThreats{
			SQLInjection:                        entry["sql_injection"],
			XSSAttack:                           entry["xss_attack"],
			RemoteFileInclusion:                 entry["remote_file_inclusion"],
			WordpressWafRuleset:                 entry["wordpress_waf_ruleset"],
			ApacheStrutsExploit:                 entry["apache_struts_exploit"],
			LocalFileInclusion:                  entry["local_file_inclusion"],
			CommonWebApplicationVulnerabilities: entry["common_web_application_vulnerabilities"],
			WebShellExecutionAttempt:            entry["webshell_execution_attempt"],
			ResponseHeaderInjection:             entry["response_header_injections"],
			OpenRedirect:                        entry["open_redirect"],
			ShellInjection:                      entry["shell_injection"],
		}
	}

	for _, entry := range d.Get("user_agents").([]map[string]bool) {
		updatedWAFSettings.UserAgents = apiclient.WAFUserAgents{
			BlockInvalidUserAgents: entry["block_invalid_user_agents"],
			BlockUnknownUserAgents: entry["block_unknown_user_agents"],
		}
	}

	for _, entry := range d.Get("traffic_sources").([]map[string]bool) {
		updatedWAFSettings.TrafficSources = apiclient.WAFTrafficSources{
			ViaTorNodes:                      entry["via_tor_nodes"],
			ViaProxyNetworks:                 entry["via_proxy_networks"],
			ViaHostingServices:               entry["via_hosting_services"],
			ViaVpn:                           entry["via_vpn"],
			ConvictedBotTraffic:              entry["convicted_bot_traffic"],
			SuspiciousTrafficByLocalIPFormat: entry["suspicious_traffic_by_local_ip_format"],
		}
	}

	for _, entry := range d.Get("anti_automation_bot_protection").([]map[string]bool) {
		updatedWAFSettings.AntiAutomationBotProtection = apiclient.WAFAntiAutomationBotProtection{
			ForceBrowserValidationOnTrafficAnomalies: entry["force_browser_validation_on_traffic_anomalies"],
			ChallengeAutomatedClients:                entry["challenge_automated_clients"],
			ChallengeHeadlessBrowsers:                entry["challenge_headless_browsers"],
			AntiScraping:                             entry["anti_scraping"],
		}
	}

	for _, entry := range d.Get("behavioral_waf").([]map[string]bool) {
		updatedWAFSettings.BehavioralWaf = apiclient.WAFBehavioralWaf{
			SpamProtection:                        entry["spam_protection"],
			BlockProbingAndForcedBrowsing:         entry["block_probing_and_forced_browsing"],
			ObfuscatedAttacksAndZeroDayMitigation: entry["obfuscated_attacks_and_zeroday_mitigation"],
			RepeatedViolations:                    entry["repeated_violations"],
			BruteForceProtection:                  entry["bruteforce_protection"],
		}
	}

	for _, entry := range d.Get("cms_protection").([]map[string]bool) {
		updatedWAFSettings.CmsProtection = apiclient.WAFCmsProtection{
			WhiteListWordpress: entry["whitelist_wordpress"],
			WhiteListModx:      entry["whitelist_modx"],
			WhiteListDrupal:    entry["whitelist_drupal"],
			WhiteListJoomla:    entry["whitelist_joomla"],
			WhiteMagento:       entry["whitelist_magento"],
			WhiteListOriginIP:  entry["whitelist_origin_ip"],
			WhiteListUmbraco:   entry["whitelist_umbraco"],
		}
	}

	for _, entry := range d.Get("all_known_bots").([]map[string]bool) {
		updatedWAFSettings.AllowKnownBots = apiclient.WAFAllowKnownBots{
			InternetArchiveBot: entry["internet_archive_bot"],
		}
	}

	return updatedWAFSettings
}

func convertWAFSettingsAPIObjectToResourceData(d *schema.ResourceData, wafSettings *apiclient.WAFSettings) {
	//Store the data
	d.Set("environment_name", wafSettings.EnvironmentName)
	d.Set("id", wafSettings.Id)
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
