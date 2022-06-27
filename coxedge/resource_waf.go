/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package coxedge

import (
	"context"
	"coxedge/terraform-provider/coxedge/apiclient"
	"coxedge/terraform-provider/coxedge/utils"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
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
	updatedWAFSettings := convertResourceDataToWAFSettingsCreateAPIObject(ctx, d)
	d.SetId(updatedWAFSettings.Id)

	//Run Update since you do not "create" these
	resourceWAFSettingsUpdate(ctx, d, m)
	//resourceWAFSettingsRead(ctx, d, m)
	return diags
}

func resourceWAFSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//check the id comes with id & environment_name, then split the value -> in case of importing the resource
	//format is <site_id>:<environment_name>
	if strings.Contains(d.Id(), ":") {
		keys := strings.Split(d.Id(), ":")
		d.SetId(keys[0])
		d.Set("environment_name", keys[1])
	}
	//Get the resource ID
	resourceId := d.Id()
	//Get the resource
	wafSettings, err := coxEdgeClient.GetWAFSettings(d.Get("environment_name").(string), resourceId)
	//wafSettings, err := coxEdgeClient.GetWAFSettings("test-codecraft", resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	convertWAFSettingsAPIObjectToResourceData(d, wafSettings)
	return diags
}

func resourceWAFSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)
	//Convert resource data to API object
	updatedWAFSettings := convertResourceDataToWAFSettingsCreateAPIObject(ctx, d)

	//Call the API
	taskResp, err := coxEdgeClient.UpdateWAFSettings(updatedWAFSettings.Id, updatedWAFSettings)
	if err != nil {
		return diag.FromErr(err)
	}

	//Await
	taskResult, err := coxEdgeClient.AwaitTaskResolveWithDefaults(ctx, taskResp.TaskId)
	if err != nil {
		return diag.FromErr(err)
	}
	fmt.Println(taskResult)

	//Set last_updated
	return resourceWAFSettingsRead(ctx, d, m)
}

func resourceWAFSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	return diag.Errorf("Cannot delete WAF")
}

func convertResourceDataToWAFSettingsCreateAPIObject(ctx context.Context, d *schema.ResourceData) apiclient.WAFSettings {
	//Create update cdnSettings struct
	updatedWAFSettings := apiclient.WAFSettings{
		EnvironmentName:       d.Get("environment_name").(string),
		Id:                    d.Get("site_id").(string),
		StackId:               d.Get("stack_id").(string),
		Domain:                d.Get("domain").(string),
		MonitoringModeEnabled: utils.BoolAddr(d.Get("monitoring_mode_enabled").(bool)),
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
			SQLInjection:                        utils.BoolAddr(entry["sql_injection"].(bool)),
			XSSAttack:                           utils.BoolAddr(entry["xss_attack"].(bool)),
			ShellShockAttack:                    utils.BoolAddr(entry["shell_shock_attack"].(bool)),
			RemoteFileInclusion:                 utils.BoolAddr(entry["remote_file_inclusion"].(bool)),
			ApacheStrutsExploit:                 utils.BoolAddr(entry["apache_struts_exploit"].(bool)),
			LocalFileInclusion:                  utils.BoolAddr(entry["local_file_inclusion"].(bool)),
			CommonWebApplicationVulnerabilities: utils.BoolAddr(entry["common_web_application_vulnerabilities"].(bool)),
			WebShellExecutionAttempt:            utils.BoolAddr(entry["webshell_execution_attempt"].(bool)),
			ProtocolAttack:                      utils.BoolAddr(entry["protocol_attack"].(bool)),
			Csrf:                                utils.BoolAddr(entry["csrf"].(bool)),
			OpenRedirect:                        utils.BoolAddr(entry["open_redirect"].(bool)),
			ShellInjection:                      utils.BoolAddr(entry["shell_injection"].(bool)),
			CodeInjection:                       utils.BoolAddr(entry["code_injection"].(bool)),
			SensitiveDataExposure:               utils.BoolAddr(entry["sensitive_data_exposure"].(bool)),
			XmlExternalEntity:                   utils.BoolAddr(entry["xml_external_entity"].(bool)),
			PersonalIdentifiableInfo:            utils.BoolAddr(entry["personal_identifiable_info"].(bool)),
			ServerSideTemplateInjection:         utils.BoolAddr(entry["serverside_template_injection"].(bool)),
		}
	}

	for _, entryRaw := range d.Get("general_policies").([]interface{}) {
		entry := entryRaw.(map[string]interface{})
		updatedWAFSettings.GeneralPolicies = apiclient.WAFGeneralPolicies{
			BlockInvalidUserAgents: utils.BoolAddr(entry["block_invalid_user_agents"].(bool)),
			BlockUnknownUserAgents: utils.BoolAddr(entry["block_unknown_user_agents"].(bool)),
			HttpMethodValidation:   utils.BoolAddr(entry["http_method_validation"].(bool)),
		}
	}

	for _, entryRaw := range d.Get("traffic_sources").([]interface{}) {
		entry := entryRaw.(map[string]interface{})
		updatedWAFSettings.TrafficSources = apiclient.WAFTrafficSources{
			ViaTorNodes:                    utils.BoolAddr(entry["via_tor_nodes"].(bool)),
			ViaProxyNetworks:               utils.BoolAddr(entry["via_proxy_networks"].(bool)),
			ViaHostingServices:             utils.BoolAddr(entry["via_hosting_services"].(bool)),
			ViaVpn:                         utils.BoolAddr(entry["via_vpn"].(bool)),
			ConvictedBotTraffic:            utils.BoolAddr(entry["convicted_bot_traffic"].(bool)),
			TrafficFromSuspiciousNatRanges: utils.BoolAddr(entry["traffic_from_suspicious_nat_ranges"].(bool)),
			ExternalReputationBlockList:    utils.BoolAddr(entry["external_reputation_block_list"].(bool)),
			TrafficViaCDN:                  utils.BoolAddr(entry["traffic_via_cdn"].(bool)),
		}
	}

	for _, entryRaw := range d.Get("anti_automation_bot_protection").([]interface{}) {
		entry := entryRaw.(map[string]interface{})
		updatedWAFSettings.AntiAutomationBotProtection = apiclient.WAFAntiAutomationBotProtection{
			ForceBrowserValidationOnTrafficAnomalies: utils.BoolAddr(entry["force_browser_validation_on_traffic_anomalies"].(bool)),
			ChallengeAutomatedClients:                utils.BoolAddr(entry["challenge_automated_clients"].(bool)),
			ChallengeHeadlessBrowsers:                utils.BoolAddr(entry["challenge_headless_browsers"].(bool)),
			AntiScraping:                             utils.BoolAddr(entry["anti_scraping"].(bool)),
		}
	}

	for _, entryRaw := range d.Get("behavioral_waf").([]interface{}) {
		entry := entryRaw.(map[string]interface{})
		updatedWAFSettings.BehavioralWaf = apiclient.WAFBehavioralWaf{
			SpamProtection:                        utils.BoolAddr(entry["spam_protection"].(bool)),
			BlockProbingAndForcedBrowsing:         utils.BoolAddr(entry["block_probing_and_forced_browsing"].(bool)),
			ObfuscatedAttacksAndZeroDayMitigation: utils.BoolAddr(entry["obfuscated_attacks_and_zeroday_mitigation"].(bool)),
			RepeatedViolations:                    utils.BoolAddr(entry["repeated_violations"].(bool)),
			BruteForceProtection:                  utils.BoolAddr(entry["bruteforce_protection"].(bool)),
		}
	}

	for _, entryRaw := range d.Get("cms_protection").([]interface{}) {
		entry := entryRaw.(map[string]interface{})
		updatedWAFSettings.CmsProtection = apiclient.WAFCmsProtection{
			WordpressWafRuleset: utils.BoolAddr(entry["wordpress_waf_ruleset"].(bool)),
			WhiteListWordpress:  utils.BoolAddr(entry["whitelist_wordpress"].(bool)),
			WhiteListModx:       utils.BoolAddr(entry["whitelist_modx"].(bool)),
			WhiteListDrupal:     utils.BoolAddr(entry["whitelist_drupal"].(bool)),
			WhiteListJoomla:     utils.BoolAddr(entry["whitelist_joomla"].(bool)),
			WhiteMagento:        utils.BoolAddr(entry["whitelist_magento"].(bool)),
			WhiteListOriginIP:   utils.BoolAddr(entry["whitelist_origin_ip"].(bool)),
			WhiteListUmbraco:    utils.BoolAddr(entry["whitelist_umbraco"].(bool)),
		}
	}

	for _, entryRaw := range d.Get("allow_known_bots").([]interface{}) {
		entry := entryRaw.(map[string]interface{})
		updatedWAFSettings.AllowKnownBots = apiclient.WAFAllowKnownBots{
			InternetArchiveBot: utils.BoolAddr(entry["internet_archive_bot"].(bool)),
		}
	}

	return updatedWAFSettings
}

func convertWAFSettingsAPIObjectToResourceData(d *schema.ResourceData, wafSettings *apiclient.WAFSettings) {
	//Store the data
	d.Set("site_id", wafSettings.Id)
	d.Set("stack_id", wafSettings.StackId)
	d.Set("domain", wafSettings.Domain)
	d.Set("api_urls", wafSettings.APIUrls)
	d.Set("monitoring_mode_enabled", *wafSettings.MonitoringModeEnabled)

	ddosSettings := make([]map[string]int, 1)
	ddosSettings[0] = make(map[string]int)
	ddosSettings[0]["global_threshold"] = wafSettings.DdosSettings.GlobalThreshold
	ddosSettings[0]["burst_threshold"] = wafSettings.DdosSettings.BurstThreshold
	ddosSettings[0]["subsecond_burst_threshold"] = wafSettings.DdosSettings.SubSecondBurstThreshold
	d.Set("ddos_settings", ddosSettings)

	owaspThreats := make([]map[string]bool, 1)
	owaspThreats[0] = make(map[string]bool)
	owaspThreats[0]["sql_injection"] = *wafSettings.OwaspThreats.SQLInjection
	owaspThreats[0]["xss_attack"] = *wafSettings.OwaspThreats.XSSAttack
	owaspThreats[0]["shell_shock_attack"] = *wafSettings.OwaspThreats.ShellShockAttack
	owaspThreats[0]["remote_file_inclusion"] = *wafSettings.OwaspThreats.RemoteFileInclusion
	owaspThreats[0]["apache_struts_exploit"] = *wafSettings.OwaspThreats.ApacheStrutsExploit
	owaspThreats[0]["local_file_inclusion"] = *wafSettings.OwaspThreats.LocalFileInclusion
	owaspThreats[0]["common_web_application_vulnerabilities"] = *wafSettings.OwaspThreats.CommonWebApplicationVulnerabilities
	owaspThreats[0]["webshell_execution_attempt"] = *wafSettings.OwaspThreats.WebShellExecutionAttempt
	owaspThreats[0]["protocol_attack"] = *wafSettings.OwaspThreats.ProtocolAttack
	owaspThreats[0]["csrf"] = *wafSettings.OwaspThreats.Csrf
	owaspThreats[0]["open_redirect"] = *wafSettings.OwaspThreats.OpenRedirect
	owaspThreats[0]["shell_injection"] = *wafSettings.OwaspThreats.ShellInjection
	owaspThreats[0]["code_injection"] = *wafSettings.OwaspThreats.CodeInjection
	owaspThreats[0]["sensitive_data_exposure"] = *wafSettings.OwaspThreats.SensitiveDataExposure
	owaspThreats[0]["xml_external_entity"] = *wafSettings.OwaspThreats.XmlExternalEntity
	owaspThreats[0]["personal_identifiable_info"] = *wafSettings.OwaspThreats.PersonalIdentifiableInfo
	owaspThreats[0]["serverside_template_injection"] = *wafSettings.OwaspThreats.ServerSideTemplateInjection
	d.Set("owasp_threats", owaspThreats)

	generalPolicies := make([]map[string]bool, 1)
	generalPolicies[0] = make(map[string]bool)
	generalPolicies[0]["block_invalid_user_agents"] = *wafSettings.GeneralPolicies.BlockInvalidUserAgents
	generalPolicies[0]["block_unknown_user_agents"] = *wafSettings.GeneralPolicies.BlockUnknownUserAgents
	generalPolicies[0]["http_method_validation"] = *wafSettings.GeneralPolicies.HttpMethodValidation
	d.Set("general_policies", generalPolicies)

	trafficSources := make([]map[string]bool, 1)
	trafficSources[0] = make(map[string]bool)
	trafficSources[0]["via_tor_nodes"] = *wafSettings.TrafficSources.ViaTorNodes
	trafficSources[0]["via_proxy_networks"] = *wafSettings.TrafficSources.ViaProxyNetworks
	trafficSources[0]["via_hosting_services"] = *wafSettings.TrafficSources.ViaHostingServices
	trafficSources[0]["via_vpn"] = *wafSettings.TrafficSources.ViaVpn
	trafficSources[0]["convicted_bot_traffic"] = *wafSettings.TrafficSources.ConvictedBotTraffic
	trafficSources[0]["traffic_from_suspicious_nat_ranges"] = *wafSettings.TrafficSources.TrafficFromSuspiciousNatRanges
	trafficSources[0]["external_reputation_block_list"] = *wafSettings.TrafficSources.ExternalReputationBlockList
	trafficSources[0]["traffic_via_cdn"] = *wafSettings.TrafficSources.TrafficViaCDN
	d.Set("traffic_sources", trafficSources)

	antiAutomationBotProtection := make([]map[string]bool, 1)
	antiAutomationBotProtection[0] = make(map[string]bool)
	antiAutomationBotProtection[0]["force_browser_validation_on_traffic_anomalies"] = *wafSettings.AntiAutomationBotProtection.ForceBrowserValidationOnTrafficAnomalies
	antiAutomationBotProtection[0]["challenge_automated_clients"] = *wafSettings.AntiAutomationBotProtection.ChallengeAutomatedClients
	antiAutomationBotProtection[0]["challenge_headless_browsers"] = *wafSettings.AntiAutomationBotProtection.ChallengeHeadlessBrowsers
	antiAutomationBotProtection[0]["anti_scraping"] = *wafSettings.AntiAutomationBotProtection.AntiScraping
	d.Set("anti_automation_bot_protection", antiAutomationBotProtection)

	behavioralWaf := make([]map[string]bool, 1)
	behavioralWaf[0] = make(map[string]bool)
	behavioralWaf[0]["spam_protection"] = *wafSettings.BehavioralWaf.SpamProtection
	behavioralWaf[0]["block_probing_and_forced_browsing"] = *wafSettings.BehavioralWaf.BlockProbingAndForcedBrowsing
	behavioralWaf[0]["obfuscated_attacks_and_zeroday_mitigation"] = *wafSettings.BehavioralWaf.ObfuscatedAttacksAndZeroDayMitigation
	behavioralWaf[0]["repeated_violations"] = *wafSettings.BehavioralWaf.RepeatedViolations
	behavioralWaf[0]["bruteforce_protection"] = *wafSettings.BehavioralWaf.BruteForceProtection
	d.Set("behavioral_waf", behavioralWaf)

	cmsProtection := make([]map[string]bool, 1)
	cmsProtection[0] = make(map[string]bool)
	cmsProtection[0]["wordpress_waf_ruleset"] = *wafSettings.CmsProtection.WordpressWafRuleset
	cmsProtection[0]["whitelist_wordpress"] = *wafSettings.CmsProtection.WhiteListWordpress
	cmsProtection[0]["whitelist_modx"] = *wafSettings.CmsProtection.WhiteListModx
	cmsProtection[0]["whitelist_drupal"] = *wafSettings.CmsProtection.WhiteListDrupal
	cmsProtection[0]["whitelist_joomla"] = *wafSettings.CmsProtection.WhiteListJoomla
	cmsProtection[0]["whitelist_magento"] = *wafSettings.CmsProtection.WhiteMagento
	cmsProtection[0]["whitelist_origin_ip"] = *wafSettings.CmsProtection.WhiteListOriginIP
	cmsProtection[0]["whitelist_umbraco"] = *wafSettings.CmsProtection.WhiteListUmbraco
	d.Set("cms_protection", cmsProtection)

	allowKnownBots := make([]map[string]bool, 1)
	allowKnownBots[0] = make(map[string]bool)
	allowKnownBots[0]["internet_archive_bot"] = *wafSettings.AllowKnownBots.InternetArchiveBot
	d.Set("allow_known_bots", allowKnownBots)
}
