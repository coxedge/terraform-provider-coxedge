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
	"strconv"
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

	owaspThreats := make([]map[string]interface{}, 1)
	owaspThreats[0] = make(map[string]interface{})
	owaspThreats[0]["sql_injection"] = wafSettings.OwaspThreats.SQLInjection
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

func convertResourceDataToWAFSettingsCreateAPIObject(ctx context.Context, d *schema.ResourceData) apiclient.WAFSettings {
	//Create update cdnSettings struct
	updatedWAFSettings := apiclient.WAFSettings{
		EnvironmentName: d.Get("environment_name").(string),
		Id:              d.Get("site_id").(string),
		StackId:         d.Get("stack_id").(string),
		Domain:          d.Get("domain").(string),
	}

	monitoringModeEnabled := d.Get("monitoring_mode_enabled").(string)
	if monitoringModeEnabled != "" {
		boolValue, _ := strconv.ParseBool(monitoringModeEnabled)
		updatedWAFSettings.MonitoringModeEnabled = utils.BoolAddr(boolValue)
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
		updatedWAFSettings.OwaspThreats = mapOswapThreats(entryRaw.(map[string]interface{}))
	}

	for _, entryRaw := range d.Get("general_policies").([]interface{}) {
		updatedWAFSettings.GeneralPolicies = mapGeneralPolicies(entryRaw.(map[string]interface{}))
	}

	for _, entryRaw := range d.Get("traffic_sources").([]interface{}) {
		updatedWAFSettings.TrafficSources = mapTrafficSources(entryRaw.(map[string]interface{}))
	}

	for _, entryRaw := range d.Get("anti_automation_bot_protection").([]interface{}) {
		updatedWAFSettings.AntiAutomationBotProtection = mapAntiAutomationBotProtection(entryRaw.(map[string]interface{}))
	}

	for _, entryRaw := range d.Get("behavioral_waf").([]interface{}) {
		updatedWAFSettings.BehavioralWaf = mapBehavioralWAF(entryRaw.(map[string]interface{}))
	}

	for _, entryRaw := range d.Get("cms_protection").([]interface{}) {
		updatedWAFSettings.CmsProtection = mapCmsProtection(entryRaw.(map[string]interface{}))
	}

	for _, entryRaw := range d.Get("allow_known_bots").([]interface{}) {
		updatedWAFSettings.AllowKnownBots = mapAllowKnownBots(entryRaw.(map[string]interface{}))
	}

	return updatedWAFSettings
}

func mapCmsProtection(entry map[string]interface{}) apiclient.WAFCmsProtection {
	cmsProtection := apiclient.WAFCmsProtection{}
	for key, value := range entry {
		if value != "" {
			switch key {
			case "wordpress_waf_ruleset":
				{
					boolValue, _ := strconv.ParseBool(entry["wordpress_waf_ruleset"].(string))
					cmsProtection.WordpressWafRuleset = utils.BoolAddr(boolValue)
				}
				break
			case "whitelist_wordpress":
				{
					boolValue, _ := strconv.ParseBool(entry["whitelist_wordpress"].(string))
					cmsProtection.WhiteListWordpress = utils.BoolAddr(boolValue)
				}
				break
			case "whitelist_modx":
				{
					boolValue, _ := strconv.ParseBool(entry["whitelist_modx"].(string))
					cmsProtection.WhiteListModx = utils.BoolAddr(boolValue)
				}
				break
			case "whitelist_drupal":
				{
					boolValue, _ := strconv.ParseBool(entry["whitelist_drupal"].(string))
					cmsProtection.WhiteListDrupal = utils.BoolAddr(boolValue)
				}
				break
			case "whitelist_joomla":
				{
					boolValue, _ := strconv.ParseBool(entry["whitelist_joomla"].(string))
					cmsProtection.WhiteListJoomla = utils.BoolAddr(boolValue)
				}
				break
			case "whitelist_magento":
				{
					boolValue, _ := strconv.ParseBool(entry["whitelist_magento"].(string))
					cmsProtection.WhiteMagento = utils.BoolAddr(boolValue)
				}
				break
			case "whitelist_origin_ip":
				{
					boolValue, _ := strconv.ParseBool(entry["whitelist_origin_ip"].(string))
					cmsProtection.WhiteListOriginIP = utils.BoolAddr(boolValue)
				}
				break
			case "whitelist_umbraco":
				{
					boolValue, _ := strconv.ParseBool(entry["whitelist_umbraco"].(string))
					cmsProtection.WhiteListUmbraco = utils.BoolAddr(boolValue)
				}
				break
			}
		}
	}
	return cmsProtection
}

func mapBehavioralWAF(entry map[string]interface{}) apiclient.WAFBehavioralWaf {
	behavioralWaf := apiclient.WAFBehavioralWaf{}
	for key, value := range entry {
		if value != "" {
			switch key {
			case "spam_protection":
				{
					boolValue, _ := strconv.ParseBool(entry["spam_protection"].(string))
					behavioralWaf.SpamProtection = utils.BoolAddr(boolValue)
				}
				break
			case "block_probing_and_forced_browsing":
				{
					boolValue, _ := strconv.ParseBool(entry["block_probing_and_forced_browsing"].(string))
					behavioralWaf.BlockProbingAndForcedBrowsing = utils.BoolAddr(boolValue)
				}
				break
			case "obfuscated_attacks_and_zeroday_mitigation":
				{
					boolValue, _ := strconv.ParseBool(entry["obfuscated_attacks_and_zeroday_mitigation"].(string))
					behavioralWaf.ObfuscatedAttacksAndZeroDayMitigation = utils.BoolAddr(boolValue)
				}
				break
			case "repeated_violations":
				{
					boolValue, _ := strconv.ParseBool(entry["repeated_violations"].(string))
					behavioralWaf.RepeatedViolations = utils.BoolAddr(boolValue)
				}
				break
			case "bruteforce_protection":
				{
					boolValue, _ := strconv.ParseBool(entry["bruteforce_protection"].(string))
					behavioralWaf.BruteForceProtection = utils.BoolAddr(boolValue)
				}
				break
			}
		}
	}
	return behavioralWaf
}

func mapAntiAutomationBotProtection(entry map[string]interface{}) apiclient.WAFAntiAutomationBotProtection {
	antiAutomationBotProtection := apiclient.WAFAntiAutomationBotProtection{}
	for key, value := range entry {
		if value != "" {
			switch key {
			case "force_browser_validation_on_traffic_anomalies":
				{
					boolValue, _ := strconv.ParseBool(entry["force_browser_validation_on_traffic_anomalies"].(string))
					antiAutomationBotProtection.ForceBrowserValidationOnTrafficAnomalies = utils.BoolAddr(boolValue)
				}
				break
			case "challenge_automated_clients":
				{
					boolValue, _ := strconv.ParseBool(entry["challenge_automated_clients"].(string))
					antiAutomationBotProtection.ChallengeAutomatedClients = utils.BoolAddr(boolValue)
				}
				break
			case "challenge_headless_browsers":
				{
					boolValue, _ := strconv.ParseBool(entry["challenge_headless_browsers"].(string))
					antiAutomationBotProtection.ChallengeHeadlessBrowsers = utils.BoolAddr(boolValue)
				}
				break
			case "anti_scraping":
				{
					boolValue, _ := strconv.ParseBool(entry["anti_scraping"].(string))
					antiAutomationBotProtection.AntiScraping = utils.BoolAddr(boolValue)
				}
				break
			}
		}
	}
	return antiAutomationBotProtection
}

func mapTrafficSources(entry map[string]interface{}) apiclient.WAFTrafficSources {
	trafficSources := apiclient.WAFTrafficSources{}
	for key, value := range entry {
		if value != "" {
			switch key {
			case "via_tor_nodes":
				{
					boolValue, _ := strconv.ParseBool(entry["via_tor_nodes"].(string))
					trafficSources.ViaTorNodes = utils.BoolAddr(boolValue)
				}
				break
			case "via_proxy_networks":
				{
					boolValue, _ := strconv.ParseBool(entry["via_proxy_networks"].(string))
					trafficSources.ViaProxyNetworks = utils.BoolAddr(boolValue)
				}
				break
			case "via_hosting_services":
				{
					boolValue, _ := strconv.ParseBool(entry["via_hosting_services"].(string))
					trafficSources.ViaHostingServices = utils.BoolAddr(boolValue)
				}
				break
			case "via_vpn":
				{
					boolValue, _ := strconv.ParseBool(entry["via_vpn"].(string))
					trafficSources.ViaVpn = utils.BoolAddr(boolValue)
				}
				break
			case "convicted_bot_traffic":
				{
					boolValue, _ := strconv.ParseBool(entry["convicted_bot_traffic"].(string))
					trafficSources.ConvictedBotTraffic = utils.BoolAddr(boolValue)
				}
				break
			case "traffic_from_suspicious_nat_ranges":
				{
					boolValue, _ := strconv.ParseBool(entry["traffic_from_suspicious_nat_ranges"].(string))
					trafficSources.TrafficFromSuspiciousNatRanges = utils.BoolAddr(boolValue)
				}
				break
			case "external_reputation_block_list":
				{
					boolValue, _ := strconv.ParseBool(entry["external_reputation_block_list"].(string))
					trafficSources.ExternalReputationBlockList = utils.BoolAddr(boolValue)
				}
				break
			case "traffic_via_cdn":
				{
					boolValue, _ := strconv.ParseBool(entry["traffic_via_cdn"].(string))
					trafficSources.TrafficViaCDN = utils.BoolAddr(boolValue)
				}
				break
			}
		}
	}
	return trafficSources
}

func mapOswapThreats(entry map[string]interface{}) apiclient.WAFOwaspThreats {
	oswapThreats := apiclient.WAFOwaspThreats{}
	for key, value := range entry {
		if value != "" {
			switch key {
			case "sql_injection":
				{
					boolValue, _ := strconv.ParseBool(entry["sql_injection"].(string))
					oswapThreats.SQLInjection = utils.BoolAddr(boolValue)
				}
				break
			case "xss_attack":
				{
					boolValue, _ := strconv.ParseBool(entry["xss_attack"].(string))
					oswapThreats.XmlExternalEntity = utils.BoolAddr(boolValue)
				}
				break
			case "shell_shock_attack":
				{
					boolValue, _ := strconv.ParseBool(entry["shell_shock_attack"].(string))
					oswapThreats.ShellShockAttack = utils.BoolAddr(boolValue)
				}
				break
			case "remote_file_inclusion":
				{
					boolValue, _ := strconv.ParseBool(entry["remote_file_inclusion"].(string))
					oswapThreats.RemoteFileInclusion = utils.BoolAddr(boolValue)
				}
				break
			case "apache_struts_exploit":
				{
					boolValue, _ := strconv.ParseBool(entry["apache_struts_exploit"].(string))
					oswapThreats.ApacheStrutsExploit = utils.BoolAddr(boolValue)
				}
				break
			case "local_file_inclusion":
				{
					boolValue, _ := strconv.ParseBool(entry["local_file_inclusion"].(string))
					oswapThreats.LocalFileInclusion = utils.BoolAddr(boolValue)
				}
				break
			case "common_web_application_vulnerabilities":
				{
					boolValue, _ := strconv.ParseBool(entry["common_web_application_vulnerabilities"].(string))
					oswapThreats.CommonWebApplicationVulnerabilities = utils.BoolAddr(boolValue)
				}
				break
			case "webshell_execution_attempt":
				{
					boolValue, _ := strconv.ParseBool(entry["webshell_execution_attempt"].(string))
					oswapThreats.WebShellExecutionAttempt = utils.BoolAddr(boolValue)
				}
				break
			case "protocol_attack":
				{
					boolValue, _ := strconv.ParseBool(entry["protocol_attack"].(string))
					oswapThreats.ProtocolAttack = utils.BoolAddr(boolValue)
				}
				break
			case "csrf":
				{
					boolValue, _ := strconv.ParseBool(entry["csrf"].(string))
					oswapThreats.ProtocolAttack = utils.BoolAddr(boolValue)
				}
				break
			case "open_redirect":
				{
					boolValue, _ := strconv.ParseBool(entry["open_redirect"].(string))
					oswapThreats.OpenRedirect = utils.BoolAddr(boolValue)
				}
				break
			case "shell_injection":
				{
					boolValue, _ := strconv.ParseBool(entry["shell_injection"].(string))
					oswapThreats.ShellInjection = utils.BoolAddr(boolValue)
				}
				break
			case "code_injection":
				{
					boolValue, _ := strconv.ParseBool(entry["code_injection"].(string))
					oswapThreats.CodeInjection = utils.BoolAddr(boolValue)
				}
				break
			case "sensitive_data_exposure":
				{
					boolValue, _ := strconv.ParseBool(entry["sensitive_data_exposure"].(string))
					oswapThreats.SensitiveDataExposure = utils.BoolAddr(boolValue)
				}
				break
			case "xml_external_entity":
				{
					boolValue, _ := strconv.ParseBool(entry["xml_external_entity"].(string))
					oswapThreats.XmlExternalEntity = utils.BoolAddr(boolValue)
				}
				break
			case "personal_identifiable_info":
				{
					boolValue, _ := strconv.ParseBool(entry["personal_identifiable_info"].(string))
					oswapThreats.PersonalIdentifiableInfo = utils.BoolAddr(boolValue)
				}
				break
			case "serverside_template_injection":
				{
					boolValue, _ := strconv.ParseBool(entry["serverside_template_injection"].(string))
					oswapThreats.ServerSideTemplateInjection = utils.BoolAddr(boolValue)
				}
				break
			}
		}
	}
	return oswapThreats
}

func mapGeneralPolicies(entry map[string]interface{}) apiclient.WAFGeneralPolicies {
	generalPolicy := apiclient.WAFGeneralPolicies{}
	for key, value := range entry {
		if value != "" {
			switch key {
			case "block_invalid_user_agents":
				{
					boolValue, _ := strconv.ParseBool(entry["block_invalid_user_agents"].(string))
					generalPolicy.BlockInvalidUserAgents = utils.BoolAddr(boolValue)
					//updatedWAFSettings.GeneralPolicies.BlockInvalidUserAgents = utils.BoolAddr(boolValue)
				}
				break
			case "block_unknown_user_agents":
				{
					boolValue, _ := strconv.ParseBool(entry["block_unknown_user_agents"].(string))
					generalPolicy.BlockUnknownUserAgents = utils.BoolAddr(boolValue)
				}
				break
			case "http_method_validation":
				{
					boolValue, _ := strconv.ParseBool(entry["http_method_validation"].(string))
					generalPolicy.HttpMethodValidation = utils.BoolAddr(boolValue)
				}
				break
			}
		}
	}
	return generalPolicy
}

func mapAllowKnownBots(entry map[string]interface{}) apiclient.WAFAllowKnownBots {
	allowKnownBots := apiclient.WAFAllowKnownBots{}
	for key, value := range entry {
		if value != "" {
			switch key {
			case "internet_archive_bot":
				{
					boolValue, _ := strconv.ParseBool(entry["internet_archive_bot"].(string))
					allowKnownBots.InternetArchiveBot = utils.BoolAddr(boolValue)
				}
				break
			}
		}
	}
	return allowKnownBots
}
