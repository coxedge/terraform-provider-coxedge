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
	updatedWAFSettings := convertResourceDataToWAFSettingsCreateAPIObject(d)

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
	d.Set("monitoring_mode_enabled", strconv.FormatBool(*wafSettings.MonitoringModeEnabled))

	ddosSettings := make([]map[string]int, 1)
	ddosSettings[0] = make(map[string]int)
	ddosSettings[0]["global_threshold"] = wafSettings.DdosSettings.GlobalThreshold
	ddosSettings[0]["burst_threshold"] = wafSettings.DdosSettings.BurstThreshold
	ddosSettings[0]["subsecond_burst_threshold"] = wafSettings.DdosSettings.SubSecondBurstThreshold
	d.Set("ddos_settings", ddosSettings)

	owaspThreats := make([]map[string]interface{}, 1)
	owaspThreats[0] = make(map[string]interface{})
	owaspThreats[0]["sql_injection"] = strconv.FormatBool(*wafSettings.OwaspThreats.SQLInjection)
	owaspThreats[0]["xss_attack"] = strconv.FormatBool(*wafSettings.OwaspThreats.XSSAttack)
	owaspThreats[0]["shell_shock_attack"] = strconv.FormatBool(*wafSettings.OwaspThreats.ShellShockAttack)
	owaspThreats[0]["remote_file_inclusion"] = strconv.FormatBool(*wafSettings.OwaspThreats.RemoteFileInclusion)
	owaspThreats[0]["apache_struts_exploit"] = strconv.FormatBool(*wafSettings.OwaspThreats.ApacheStrutsExploit)
	owaspThreats[0]["local_file_inclusion"] = strconv.FormatBool(*wafSettings.OwaspThreats.LocalFileInclusion)
	owaspThreats[0]["common_web_application_vulnerabilities"] = strconv.FormatBool(*wafSettings.OwaspThreats.CommonWebApplicationVulnerabilities)
	owaspThreats[0]["webshell_execution_attempt"] = strconv.FormatBool(*wafSettings.OwaspThreats.WebShellExecutionAttempt)
	owaspThreats[0]["protocol_attack"] = strconv.FormatBool(*wafSettings.OwaspThreats.ProtocolAttack)
	owaspThreats[0]["csrf"] = strconv.FormatBool(*wafSettings.OwaspThreats.Csrf)
	owaspThreats[0]["open_redirect"] = strconv.FormatBool(*wafSettings.OwaspThreats.OpenRedirect)
	owaspThreats[0]["shell_injection"] = strconv.FormatBool(*wafSettings.OwaspThreats.ShellInjection)
	owaspThreats[0]["code_injection"] = strconv.FormatBool(*wafSettings.OwaspThreats.CodeInjection)
	owaspThreats[0]["sensitive_data_exposure"] = strconv.FormatBool(*wafSettings.OwaspThreats.SensitiveDataExposure)
	owaspThreats[0]["xml_external_entity"] = strconv.FormatBool(*wafSettings.OwaspThreats.XmlExternalEntity)
	owaspThreats[0]["personal_identifiable_info"] = strconv.FormatBool(*wafSettings.OwaspThreats.PersonalIdentifiableInfo)
	owaspThreats[0]["serverside_template_injection"] = strconv.FormatBool(*wafSettings.OwaspThreats.ServerSideTemplateInjection)
	d.Set("owasp_threats", owaspThreats)

	generalPolicies := make([]map[string]interface{}, 1)
	generalPolicies[0] = make(map[string]interface{})
	generalPolicies[0]["block_invalid_user_agents"] = strconv.FormatBool(*wafSettings.GeneralPolicies.BlockInvalidUserAgents)
	generalPolicies[0]["block_unknown_user_agents"] = strconv.FormatBool(*wafSettings.GeneralPolicies.BlockUnknownUserAgents)
	generalPolicies[0]["http_method_validation"] = strconv.FormatBool(*wafSettings.GeneralPolicies.HttpMethodValidation)
	d.Set("general_policies", generalPolicies)

	trafficSources := make([]map[string]interface{}, 1)
	trafficSources[0] = make(map[string]interface{})
	trafficSources[0]["via_tor_nodes"] = strconv.FormatBool(*wafSettings.TrafficSources.ViaTorNodes)
	trafficSources[0]["via_proxy_networks"] = strconv.FormatBool(*wafSettings.TrafficSources.ViaProxyNetworks)
	trafficSources[0]["via_hosting_services"] = strconv.FormatBool(*wafSettings.TrafficSources.ViaHostingServices)
	trafficSources[0]["via_vpn"] = strconv.FormatBool(*wafSettings.TrafficSources.ViaVpn)
	trafficSources[0]["convicted_bot_traffic"] = strconv.FormatBool(*wafSettings.TrafficSources.ConvictedBotTraffic)
	trafficSources[0]["traffic_from_suspicious_nat_ranges"] = strconv.FormatBool(*wafSettings.TrafficSources.TrafficFromSuspiciousNatRanges)
	trafficSources[0]["external_reputation_block_list"] = strconv.FormatBool(*wafSettings.TrafficSources.ExternalReputationBlockList)
	trafficSources[0]["traffic_via_cdn"] = strconv.FormatBool(*wafSettings.TrafficSources.TrafficViaCDN)
	d.Set("traffic_sources", trafficSources)

	antiAutomationBotProtection := make([]map[string]interface{}, 1)
	antiAutomationBotProtection[0] = make(map[string]interface{})
	antiAutomationBotProtection[0]["force_browser_validation_on_traffic_anomalies"] = strconv.FormatBool(*wafSettings.AntiAutomationBotProtection.ForceBrowserValidationOnTrafficAnomalies)
	antiAutomationBotProtection[0]["challenge_automated_clients"] = strconv.FormatBool(*wafSettings.AntiAutomationBotProtection.ChallengeAutomatedClients)
	antiAutomationBotProtection[0]["challenge_headless_browsers"] = strconv.FormatBool(*wafSettings.AntiAutomationBotProtection.ChallengeHeadlessBrowsers)
	antiAutomationBotProtection[0]["anti_scraping"] = strconv.FormatBool(*wafSettings.AntiAutomationBotProtection.AntiScraping)
	d.Set("anti_automation_bot_protection", antiAutomationBotProtection)

	behavioralWaf := make([]map[string]interface{}, 1)
	behavioralWaf[0] = make(map[string]interface{})
	behavioralWaf[0]["spam_protection"] = strconv.FormatBool(*wafSettings.BehavioralWaf.SpamProtection)
	behavioralWaf[0]["block_probing_and_forced_browsing"] = strconv.FormatBool(*wafSettings.BehavioralWaf.BlockProbingAndForcedBrowsing)
	behavioralWaf[0]["obfuscated_attacks_and_zeroday_mitigation"] = strconv.FormatBool(*wafSettings.BehavioralWaf.ObfuscatedAttacksAndZeroDayMitigation)
	behavioralWaf[0]["repeated_violations"] = strconv.FormatBool(*wafSettings.BehavioralWaf.RepeatedViolations)
	behavioralWaf[0]["bruteforce_protection"] = strconv.FormatBool(*wafSettings.BehavioralWaf.BruteForceProtection)
	d.Set("behavioral_waf", behavioralWaf)

	cmsProtection := make([]map[string]interface{}, 1)
	cmsProtection[0] = make(map[string]interface{})
	cmsProtection[0]["wordpress_waf_ruleset"] = strconv.FormatBool(*wafSettings.CmsProtection.WordpressWafRuleset)
	cmsProtection[0]["whitelist_wordpress"] = strconv.FormatBool(*wafSettings.CmsProtection.WhiteListWordpress)
	cmsProtection[0]["whitelist_modx"] = strconv.FormatBool(*wafSettings.CmsProtection.WhiteListModx)
	cmsProtection[0]["whitelist_drupal"] = strconv.FormatBool(*wafSettings.CmsProtection.WhiteListDrupal)
	cmsProtection[0]["whitelist_joomla"] = strconv.FormatBool(*wafSettings.CmsProtection.WhiteListJoomla)
	cmsProtection[0]["whitelist_magento"] = strconv.FormatBool(*wafSettings.CmsProtection.WhiteMagento)
	cmsProtection[0]["whitelist_origin_ip"] = strconv.FormatBool(*wafSettings.CmsProtection.WhiteListOriginIP)
	cmsProtection[0]["whitelist_umbraco"] = strconv.FormatBool(*wafSettings.CmsProtection.WhiteListUmbraco)
	d.Set("cms_protection", cmsProtection)

	allowKnownBots := make([]map[string]interface{}, 1)
	allowKnownBots[0] = mapResourceAllowKnownBots(allowKnownBots, wafSettings)
	d.Set("allow_known_bots", allowKnownBots)
}

func mapResourceAllowKnownBots(allowKnownBots []map[string]interface{}, wafSettings *apiclient.WAFSettings) map[string]interface{} {
	allowKnownBots[0] = make(map[string]interface{})
	allowKnownBots[0]["acquia_uptime"] = strconv.FormatBool(*wafSettings.AllowKnownBots.AcquiaUptime)
	allowKnownBots[0]["add_search_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.AddSearchBot)
	allowKnownBots[0]["adestra_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.AdestraBot)
	allowKnownBots[0]["adjust_servers"] = strconv.FormatBool(*wafSettings.AllowKnownBots.AdjustServers)
	allowKnownBots[0]["ahrefs_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.AhrefsBot)
	allowKnownBots[0]["alerta_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.AlertaBot)
	allowKnownBots[0]["alexa_ia_archiver"] = strconv.FormatBool(*wafSettings.AllowKnownBots.AlexaIaArchiver)
	allowKnownBots[0]["alexa_technologies"] = strconv.FormatBool(*wafSettings.AllowKnownBots.AlexaTechnologies)
	allowKnownBots[0]["amazon_route_53_health_check_service"] = strconv.FormatBool(*wafSettings.AllowKnownBots.AmazonRoute53HealthCheckService)
	allowKnownBots[0]["applebot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.Applebot)
	allowKnownBots[0]["apple_news_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.AppleNewsBot)
	allowKnownBots[0]["ask_jeeves_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.AskJeevesBot)
	allowKnownBots[0]["audisto_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.AudistoBot)
	allowKnownBots[0]["baidu_spider_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.BaiduSpiderBot)
	allowKnownBots[0]["baidu_spider_japan_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.BaiduSpiderJapanBot)
	allowKnownBots[0]["binary_canary"] = strconv.FormatBool(*wafSettings.AllowKnownBots.BinaryCanary)
	allowKnownBots[0]["bitbucket_webhook"] = strconv.FormatBool(*wafSettings.AllowKnownBots.BitbucketWebhook)
	allowKnownBots[0]["blekko_scout_jet_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.BlekkoScoutJetBot)
	allowKnownBots[0]["chrome_compression_proxy"] = strconv.FormatBool(*wafSettings.AllowKnownBots.ChromeCompressionProxy)
	allowKnownBots[0]["coccocbot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.Coccocbot)
	allowKnownBots[0]["cookie_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.CookieBot)
	allowKnownBots[0]["cybersource"] = strconv.FormatBool(*wafSettings.AllowKnownBots.Cybersource)
	allowKnownBots[0]["daumoa_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.DaumoaBot)
	allowKnownBots[0]["detectify_scanner"] = strconv.FormatBool(*wafSettings.AllowKnownBots.DetectifyScanner)
	allowKnownBots[0]["digi_cert_dcv_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.DigiCertDCVBot)
	allowKnownBots[0]["dotmic_dot_bot_commercial"] = strconv.FormatBool(*wafSettings.AllowKnownBots.DotmicDotBotCommercial)
	allowKnownBots[0]["duck_duck_go_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.DuckDuckGoBot)
	allowKnownBots[0]["facebook_external_hit_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.FacebookExternalHitBot)
	allowKnownBots[0]["feeder_co"] = strconv.FormatBool(*wafSettings.AllowKnownBots.FeederCo)
	allowKnownBots[0]["feed_press"] = strconv.FormatBool(*wafSettings.AllowKnownBots.FeedPress)
	allowKnownBots[0]["feed_wind"] = strconv.FormatBool(*wafSettings.AllowKnownBots.FeedWind)
	allowKnownBots[0]["freshping_monitoring"] = strconv.FormatBool(*wafSettings.AllowKnownBots.FreshpingMonitoring)
	allowKnownBots[0]["geckoboard"] = strconv.FormatBool(*wafSettings.AllowKnownBots.Geckoboard)
	allowKnownBots[0]["ghost_inspector"] = strconv.FormatBool(*wafSettings.AllowKnownBots.GhostInspector)
	allowKnownBots[0]["gomez"] = strconv.FormatBool(*wafSettings.AllowKnownBots.Gomez)
	allowKnownBots[0]["goo_japan_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.GooJapanBot)
	allowKnownBots[0]["google_ads_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.GoogleAdsBot)
	allowKnownBots[0]["google_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.GoogleBot)
	allowKnownBots[0]["google_cloud_monitoring_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.GoogleCloudMonitoringBot)
	allowKnownBots[0]["google_feed_fetcher_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.GoogleFeedFetcherBot)
	allowKnownBots[0]["google_image_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.GoogleImageBot)
	allowKnownBots[0]["google_image_proxy"] = strconv.FormatBool(*wafSettings.AllowKnownBots.GoogleImageProxy)
	allowKnownBots[0]["google_mediapartners_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.GoogleMediapartnersBot)
	allowKnownBots[0]["google_mobile_ads_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.GoogleMobileAdsBot)
	allowKnownBots[0]["google_news_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.GoogleNewsBot)
	allowKnownBots[0]["google_page_speed_insights"] = strconv.FormatBool(*wafSettings.AllowKnownBots.GooglePageSpeedInsights)
	//todo need to remove this condition once UAT and Prod response matched
	if wafSettings.AllowKnownBots.GoogleStructuredDataTestingTool != nil {
		allowKnownBots[0]["google_structured_data_testing_tool"] = strconv.FormatBool(*wafSettings.AllowKnownBots.GoogleStructuredDataTestingTool)
	}
	allowKnownBots[0]["google_verification_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.GoogleVerificationBot)
	allowKnownBots[0]["google_video_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.GoogleVideoBot)
	allowKnownBots[0]["google_web_light"] = strconv.FormatBool(*wafSettings.AllowKnownBots.GoogleWebLight)
	allowKnownBots[0]["grapeshot_bot_commercial"] = strconv.FormatBool(*wafSettings.AllowKnownBots.GrapeshotBotCommercial)
	allowKnownBots[0]["gree_japan_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.GreeJapanBot)
	allowKnownBots[0]["hetrix_tools"] = strconv.FormatBool(*wafSettings.AllowKnownBots.HetrixTools)
	allowKnownBots[0]["hi_pay"] = strconv.FormatBool(*wafSettings.AllowKnownBots.HiPay)
	allowKnownBots[0]["hyperspin_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.HyperspinBot)
	allowKnownBots[0]["ias_crawler_commercial"] = strconv.FormatBool(*wafSettings.AllowKnownBots.IASCrawlerCommercial)
	allowKnownBots[0]["internet_archive_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.InternetArchiveBot)
	allowKnownBots[0]["jetpack_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.JetpackBot)
	allowKnownBots[0]["jike_spider_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.JikeSpiderBot)
	allowKnownBots[0]["j_word_japan_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.JWordJapanBot)
	allowKnownBots[0]["kakao_user_agent"] = strconv.FormatBool(*wafSettings.AllowKnownBots.KAKAOUserAgent)
	allowKnownBots[0]["kyoto_tohoku_crawler"] = strconv.FormatBool(*wafSettings.AllowKnownBots.KyotoTohokuCrawler)
	allowKnownBots[0]["landau_media_spider"] = strconv.FormatBool(*wafSettings.AllowKnownBots.LandauMediaSpider)
	allowKnownBots[0]["lets_encrypt"] = strconv.FormatBool(*wafSettings.AllowKnownBots.LetsEncrypt)
	allowKnownBots[0]["line_japan_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.LineJapanBot)
	allowKnownBots[0]["linked_in_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.LinkedInBot)
	allowKnownBots[0]["livedoor_japan_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.LivedoorJapanBot)
	allowKnownBots[0]["mail_ru_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.MailRuBot)
	allowKnownBots[0]["manage_wp"] = strconv.FormatBool(*wafSettings.AllowKnownBots.ManageWP)
	allowKnownBots[0]["microsoft_bing_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.MicrosoftBingBot)
	allowKnownBots[0]["microsoft_bing_preview_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.MicrosoftBingPreviewBot)
	allowKnownBots[0]["microsoft_msn_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.MicrosoftMSNBot)
	allowKnownBots[0]["microsoft_skype_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.MicrosoftSkypeBot)
	allowKnownBots[0]["mixi_japan_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.MixiJapanBot)
	allowKnownBots[0]["mobage_japan_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.MobageJapanBot)
	allowKnownBots[0]["naver_yeti_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.NaverYetiBot)
	allowKnownBots[0]["new_relic_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.NewRelicBot)
	allowKnownBots[0]["ocn_japan_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.OCNJapanBot)
	allowKnownBots[0]["panopta_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.PanoptaBot)
	allowKnownBots[0]["parse_ly_scraper"] = strconv.FormatBool(*wafSettings.AllowKnownBots.ParseLyScraper)
	allowKnownBots[0]["pay_pal_ipn"] = strconv.FormatBool(*wafSettings.AllowKnownBots.PayPalIPN)
	allowKnownBots[0]["petal_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.PetalBot)
	allowKnownBots[0]["pingdom"] = strconv.FormatBool(*wafSettings.AllowKnownBots.Pingdom)
	allowKnownBots[0]["pinterest_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.PinterestBot)
	allowKnownBots[0]["qwantify_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.QwantifyBot)
	allowKnownBots[0]["roger_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.RogerBot)
	allowKnownBots[0]["sage_pay"] = strconv.FormatBool(*wafSettings.AllowKnownBots.SagePay)
	allowKnownBots[0]["sectigo_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.SectigoBot)
	allowKnownBots[0]["semrush_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.SemrushBot)
	allowKnownBots[0]["server_density_service_monitoring_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.ServerDensityServiceMonitoringBot)
	allowKnownBots[0]["seznam_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.SeznamBot)
	allowKnownBots[0]["shareaholic_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.ShareaholicBot)
	allowKnownBots[0]["site_24_x_7_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.Site24X7Bot)
	allowKnownBots[0]["siteimprove_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.SiteimproveBot)
	allowKnownBots[0]["site_lock_spider"] = strconv.FormatBool(*wafSettings.AllowKnownBots.SiteLockSpider)
	allowKnownBots[0]["slack_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.SlackBot)
	allowKnownBots[0]["sogou_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.SogouBot)
	allowKnownBots[0]["soso_spider_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.SosoSpiderBot)
	allowKnownBots[0]["spatineo"] = strconv.FormatBool(*wafSettings.AllowKnownBots.Spatineo)
	allowKnownBots[0]["spring_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.SpringBot)
	allowKnownBots[0]["stackify"] = strconv.FormatBool(*wafSettings.AllowKnownBots.Stackify)
	allowKnownBots[0]["status_cake_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.StatusCakeBot)
	allowKnownBots[0]["stripe"] = strconv.FormatBool(*wafSettings.AllowKnownBots.Stripe)
	allowKnownBots[0]["sucuri_uptime_monitor_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.SucuriUptimeMonitorBot)
	allowKnownBots[0]["telegram_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.TelegramBot)
	allowKnownBots[0]["testomato_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.TestomatoBot)
	allowKnownBots[0]["the_find_crawler"] = strconv.FormatBool(*wafSettings.AllowKnownBots.TheFindCrawler)
	allowKnownBots[0]["twitter_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.TwitterBot)
	allowKnownBots[0]["uptime_robot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.UptimeRobot)
	allowKnownBots[0]["vkontakte_external_hit_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.VkontakteExternalHitBot)
	allowKnownBots[0]["w_3_c"] = strconv.FormatBool(*wafSettings.AllowKnownBots.W3C)
	allowKnownBots[0]["wordfence_central"] = strconv.FormatBool(*wafSettings.AllowKnownBots.WordfenceCentral)
	allowKnownBots[0]["workato"] = strconv.FormatBool(*wafSettings.AllowKnownBots.Workato)
	allowKnownBots[0]["xml_sitemaps"] = strconv.FormatBool(*wafSettings.AllowKnownBots.XMLSitemaps)
	allowKnownBots[0]["yahoo_inktomi_slurp_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.YahooInktomiSlurpBot)
	allowKnownBots[0]["yahoo_japan_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.YahooJapanBot)
	allowKnownBots[0]["yahoo_link_preview"] = strconv.FormatBool(*wafSettings.AllowKnownBots.YahooLinkPreview)
	allowKnownBots[0]["yahoo_seeker_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.YahooSeekerBot)
	allowKnownBots[0]["yahoo_slurp_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.YahooSlurpBot)
	allowKnownBots[0]["yandex_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.YandexBot)
	allowKnownBots[0]["yisou_spider_commercial"] = strconv.FormatBool(*wafSettings.AllowKnownBots.YisouSpiderCommercial)
	allowKnownBots[0]["yodao_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.YodaoBot)
	allowKnownBots[0]["zendesk_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.ZendeskBot)
	allowKnownBots[0]["zoho_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.ZohoBot)
	allowKnownBots[0]["zum_bot"] = strconv.FormatBool(*wafSettings.AllowKnownBots.ZumBot)
	return allowKnownBots[0]
}

func convertResourceDataToWAFSettingsCreateAPIObject(d *schema.ResourceData) apiclient.WAFSettings {
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
					boolValue, _ := strconv.ParseBool(value.(string))
					cmsProtection.WordpressWafRuleset = utils.BoolAddr(boolValue)
				}
				break
			case "whitelist_wordpress":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					cmsProtection.WhiteListWordpress = utils.BoolAddr(boolValue)
				}
				break
			case "whitelist_modx":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					cmsProtection.WhiteListModx = utils.BoolAddr(boolValue)
				}
				break
			case "whitelist_drupal":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					cmsProtection.WhiteListDrupal = utils.BoolAddr(boolValue)
				}
				break
			case "whitelist_joomla":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					cmsProtection.WhiteListJoomla = utils.BoolAddr(boolValue)
				}
				break
			case "whitelist_magento":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					cmsProtection.WhiteMagento = utils.BoolAddr(boolValue)
				}
				break
			case "whitelist_origin_ip":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					cmsProtection.WhiteListOriginIP = utils.BoolAddr(boolValue)
				}
				break
			case "whitelist_umbraco":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
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
					boolValue, _ := strconv.ParseBool(value.(string))
					behavioralWaf.SpamProtection = utils.BoolAddr(boolValue)
				}
				break
			case "block_probing_and_forced_browsing":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					behavioralWaf.BlockProbingAndForcedBrowsing = utils.BoolAddr(boolValue)
				}
				break
			case "obfuscated_attacks_and_zeroday_mitigation":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					behavioralWaf.ObfuscatedAttacksAndZeroDayMitigation = utils.BoolAddr(boolValue)
				}
				break
			case "repeated_violations":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					behavioralWaf.RepeatedViolations = utils.BoolAddr(boolValue)
				}
				break
			case "bruteforce_protection":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
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
					boolValue, _ := strconv.ParseBool(value.(string))
					antiAutomationBotProtection.ForceBrowserValidationOnTrafficAnomalies = utils.BoolAddr(boolValue)
				}
				break
			case "challenge_automated_clients":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					antiAutomationBotProtection.ChallengeAutomatedClients = utils.BoolAddr(boolValue)
				}
				break
			case "challenge_headless_browsers":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					antiAutomationBotProtection.ChallengeHeadlessBrowsers = utils.BoolAddr(boolValue)
				}
				break
			case "anti_scraping":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
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
					boolValue, _ := strconv.ParseBool(value.(string))
					trafficSources.ViaTorNodes = utils.BoolAddr(boolValue)
				}
				break
			case "via_proxy_networks":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					trafficSources.ViaProxyNetworks = utils.BoolAddr(boolValue)
				}
				break
			case "via_hosting_services":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					trafficSources.ViaHostingServices = utils.BoolAddr(boolValue)
				}
				break
			case "via_vpn":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					trafficSources.ViaVpn = utils.BoolAddr(boolValue)
				}
				break
			case "convicted_bot_traffic":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					trafficSources.ConvictedBotTraffic = utils.BoolAddr(boolValue)
				}
				break
			case "traffic_from_suspicious_nat_ranges":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					trafficSources.TrafficFromSuspiciousNatRanges = utils.BoolAddr(boolValue)
				}
				break
			case "external_reputation_block_list":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					trafficSources.ExternalReputationBlockList = utils.BoolAddr(boolValue)
				}
				break
			case "traffic_via_cdn":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
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
					boolValue, _ := strconv.ParseBool(value.(string))
					oswapThreats.SQLInjection = utils.BoolAddr(boolValue)
				}
				break
			case "xss_attack":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					oswapThreats.XmlExternalEntity = utils.BoolAddr(boolValue)
				}
				break
			case "shell_shock_attack":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					oswapThreats.ShellShockAttack = utils.BoolAddr(boolValue)
				}
				break
			case "remote_file_inclusion":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					oswapThreats.RemoteFileInclusion = utils.BoolAddr(boolValue)
				}
				break
			case "apache_struts_exploit":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					oswapThreats.ApacheStrutsExploit = utils.BoolAddr(boolValue)
				}
				break
			case "local_file_inclusion":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					oswapThreats.LocalFileInclusion = utils.BoolAddr(boolValue)
				}
				break
			case "common_web_application_vulnerabilities":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					oswapThreats.CommonWebApplicationVulnerabilities = utils.BoolAddr(boolValue)
				}
				break
			case "webshell_execution_attempt":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					oswapThreats.WebShellExecutionAttempt = utils.BoolAddr(boolValue)
				}
				break
			case "protocol_attack":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					oswapThreats.ProtocolAttack = utils.BoolAddr(boolValue)
				}
				break
			case "csrf":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					oswapThreats.ProtocolAttack = utils.BoolAddr(boolValue)
				}
				break
			case "open_redirect":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					oswapThreats.OpenRedirect = utils.BoolAddr(boolValue)
				}
				break
			case "shell_injection":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					oswapThreats.ShellInjection = utils.BoolAddr(boolValue)
				}
				break
			case "code_injection":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					oswapThreats.CodeInjection = utils.BoolAddr(boolValue)
				}
				break
			case "sensitive_data_exposure":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					oswapThreats.SensitiveDataExposure = utils.BoolAddr(boolValue)
				}
				break
			case "xml_external_entity":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					oswapThreats.XmlExternalEntity = utils.BoolAddr(boolValue)
				}
				break
			case "personal_identifiable_info":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					oswapThreats.PersonalIdentifiableInfo = utils.BoolAddr(boolValue)
				}
				break
			case "serverside_template_injection":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
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
					boolValue, _ := strconv.ParseBool(value.(string))
					generalPolicy.BlockInvalidUserAgents = utils.BoolAddr(boolValue)
				}
				break
			case "block_unknown_user_agents":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					generalPolicy.BlockUnknownUserAgents = utils.BoolAddr(boolValue)
				}
				break
			case "http_method_validation":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
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
			case "acquia_uptime":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.AcquiaUptime = utils.BoolAddr(boolValue)
				}
				break
			case "add_search_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.AddSearchBot = utils.BoolAddr(boolValue)
				}
				break
			case "adestra_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.AdestraBot = utils.BoolAddr(boolValue)
				}
				break
			case "adjust_servers":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.AdjustServers = utils.BoolAddr(boolValue)
				}
				break
			case "ahrefs_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.AhrefsBot = utils.BoolAddr(boolValue)
				}
				break
			case "alerta_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.AlertaBot = utils.BoolAddr(boolValue)
				}
				break
			case "alexa_ia_archiver":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.AlexaIaArchiver = utils.BoolAddr(boolValue)
				}
				break
			case "alexa_technologies":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.AlexaTechnologies = utils.BoolAddr(boolValue)
				}
				break
			case "amazon_route_53_health_check_service":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.AmazonRoute53HealthCheckService = utils.BoolAddr(boolValue)
				}
				break
			case "applebot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.Applebot = utils.BoolAddr(boolValue)
				}
				break
			case "apple_news_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.AppleNewsBot = utils.BoolAddr(boolValue)
				}
				break
			case "ask_jeeves_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.AskJeevesBot = utils.BoolAddr(boolValue)
				}
				break
			case "audisto_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.AudistoBot = utils.BoolAddr(boolValue)
				}
				break
			case "baidu_spider_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.BaiduSpiderBot = utils.BoolAddr(boolValue)
				}
				break
			case "baidu_spider_japan_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.BaiduSpiderJapanBot = utils.BoolAddr(boolValue)
				}
				break
			case "binary_canary":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.BinaryCanary = utils.BoolAddr(boolValue)
				}
				break
			case "bitbucket_webhook":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.BitbucketWebhook = utils.BoolAddr(boolValue)
				}
				break
			case "blekko_scout_jet_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.BlekkoScoutJetBot = utils.BoolAddr(boolValue)
				}
				break
			case "chrome_compression_proxy":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.ChromeCompressionProxy = utils.BoolAddr(boolValue)
				}
				break
			case "coccocbot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.Coccocbot = utils.BoolAddr(boolValue)
				}
				break
			case "cookie_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.CookieBot = utils.BoolAddr(boolValue)
				}
				break
			case "cybersource":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.Cybersource = utils.BoolAddr(boolValue)
				}
				break
			case "daumoa_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.DaumoaBot = utils.BoolAddr(boolValue)
				}
				break
			case "detectify_scanner":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.DetectifyScanner = utils.BoolAddr(boolValue)
				}
				break
			case "digi_cert_dcv_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.DigiCertDCVBot = utils.BoolAddr(boolValue)
				}
				break
			case "dotmic_dot_bot_commercial":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.DotmicDotBotCommercial = utils.BoolAddr(boolValue)
				}
				break
			case "duck_duck_go_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.DuckDuckGoBot = utils.BoolAddr(boolValue)
				}
				break
			case "facebook_external_hit_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.FacebookExternalHitBot = utils.BoolAddr(boolValue)
				}
				break
			case "feeder_co":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.FeederCo = utils.BoolAddr(boolValue)
				}
				break
			case "feed_press":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.FeedPress = utils.BoolAddr(boolValue)
				}
				break
			case "feed_wind":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.FeedWind = utils.BoolAddr(boolValue)
				}
				break
			case "freshping_monitoring":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.FreshpingMonitoring = utils.BoolAddr(boolValue)
				}
				break
			case "geckoboard":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.Geckoboard = utils.BoolAddr(boolValue)
				}
				break
			case "ghost_inspector":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.GhostInspector = utils.BoolAddr(boolValue)
				}
				break
			case "gomez":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.Gomez = utils.BoolAddr(boolValue)
				}
				break
			case "goo_japan_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.GooJapanBot = utils.BoolAddr(boolValue)
				}
				break
			case "google_ads_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.GoogleAdsBot = utils.BoolAddr(boolValue)
				}
				break
			case "google_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.GoogleBot = utils.BoolAddr(boolValue)
				}
				break
			case "google_cloud_monitoring_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.GoogleCloudMonitoringBot = utils.BoolAddr(boolValue)
				}
				break
			case "google_feed_fetcher_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.GoogleFeedFetcherBot = utils.BoolAddr(boolValue)
				}
				break
			case "google_image_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.GoogleImageBot = utils.BoolAddr(boolValue)
				}
				break
			case "google_image_proxy":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.GoogleImageProxy = utils.BoolAddr(boolValue)
				}
				break
			case "google_mediapartners_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.GoogleMediapartnersBot = utils.BoolAddr(boolValue)
				}
				break
			case "google_mobile_ads_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.GoogleMobileAdsBot = utils.BoolAddr(boolValue)
				}
				break
			case "google_news_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.GoogleNewsBot = utils.BoolAddr(boolValue)
				}
				break
			case "google_page_speed_insights":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.GooglePageSpeedInsights = utils.BoolAddr(boolValue)
				}
				break
			case "google_structured_data_testing_tool":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.GoogleStructuredDataTestingTool = utils.BoolAddr(boolValue)
				}
				break
			case "google_verification_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.GoogleVerificationBot = utils.BoolAddr(boolValue)
				}
				break
			case "google_video_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.GoogleVideoBot = utils.BoolAddr(boolValue)
				}
				break
			case "google_web_light":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.GoogleWebLight = utils.BoolAddr(boolValue)
				}
				break
			case "grapeshot_bot_commercial":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.GrapeshotBotCommercial = utils.BoolAddr(boolValue)
				}
				break
			case "gree_japan_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.GreeJapanBot = utils.BoolAddr(boolValue)
				}
				break
			case "hetrix_tools":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.HetrixTools = utils.BoolAddr(boolValue)
				}
				break
			case "hi_pay":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.HiPay = utils.BoolAddr(boolValue)
				}
				break
			case "hyperspin_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.HyperspinBot = utils.BoolAddr(boolValue)
				}
				break
			case "ias_crawler_commercial":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.IASCrawlerCommercial = utils.BoolAddr(boolValue)
				}
				break
			case "internet_archive_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.InternetArchiveBot = utils.BoolAddr(boolValue)
				}
				break
			case "jetpack_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.JetpackBot = utils.BoolAddr(boolValue)
				}
				break
			case "jike_spider_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.JikeSpiderBot = utils.BoolAddr(boolValue)
				}
				break
			case "j_word_japan_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.JWordJapanBot = utils.BoolAddr(boolValue)
				}
				break
			case "kakao_user_agent":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.KAKAOUserAgent = utils.BoolAddr(boolValue)
				}
				break
			case "kyoto_tohoku_crawler":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.KyotoTohokuCrawler = utils.BoolAddr(boolValue)
				}
				break
			case "landau_media_spider":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.LandauMediaSpider = utils.BoolAddr(boolValue)
				}
				break
			case "lets_encrypt":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.LetsEncrypt = utils.BoolAddr(boolValue)
				}
				break
			case "line_japan_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.LineJapanBot = utils.BoolAddr(boolValue)
				}
				break
			case "linked_in_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.LinkedInBot = utils.BoolAddr(boolValue)
				}
				break
			case "livedoor_japan_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.LivedoorJapanBot = utils.BoolAddr(boolValue)
				}
				break
			case "mail_ru_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.MailRuBot = utils.BoolAddr(boolValue)
				}
				break
			case "manage_wp":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.ManageWP = utils.BoolAddr(boolValue)
				}
				break
			case "microsoft_bing_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.MicrosoftBingBot = utils.BoolAddr(boolValue)
				}
				break
			case "microsoft_bing_preview_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.MicrosoftBingPreviewBot = utils.BoolAddr(boolValue)
				}
				break
			case "microsoft_msn_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.MicrosoftMSNBot = utils.BoolAddr(boolValue)
				}
				break
			case "microsoft_skype_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.MicrosoftSkypeBot = utils.BoolAddr(boolValue)
				}
				break
			case "mixi_japan_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.MixiJapanBot = utils.BoolAddr(boolValue)
				}
				break
			case "mobage_japan_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.MobageJapanBot = utils.BoolAddr(boolValue)
				}
				break
			case "naver_yeti_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.NaverYetiBot = utils.BoolAddr(boolValue)
				}
				break
			case "new_relic_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.NewRelicBot = utils.BoolAddr(boolValue)
				}
				break
			case "ocn_japan_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.OCNJapanBot = utils.BoolAddr(boolValue)
				}
				break
			case "panopta_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.PanoptaBot = utils.BoolAddr(boolValue)
				}
				break
			case "parse_ly_scraper":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.ParseLyScraper = utils.BoolAddr(boolValue)
				}
				break
			case "pay_pal_ipn":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.PayPalIPN = utils.BoolAddr(boolValue)
				}
				break
			case "petal_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.PetalBot = utils.BoolAddr(boolValue)
				}
				break
			case "pingdom":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.Pingdom = utils.BoolAddr(boolValue)
				}
				break
			case "pinterest_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.PinterestBot = utils.BoolAddr(boolValue)
				}
				break
			case "qwantify_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.QwantifyBot = utils.BoolAddr(boolValue)
				}
				break
			case "roger_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.RogerBot = utils.BoolAddr(boolValue)
				}
				break
			case "sage_pay":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.SagePay = utils.BoolAddr(boolValue)
				}
				break
			case "sectigo_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.SectigoBot = utils.BoolAddr(boolValue)
				}
				break
			case "semrush_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.SemrushBot = utils.BoolAddr(boolValue)
				}
				break
			case "server_density_service_monitoring_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.ServerDensityServiceMonitoringBot = utils.BoolAddr(boolValue)
				}
				break
			case "seznam_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.SeznamBot = utils.BoolAddr(boolValue)
				}
				break
			case "shareaholic_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.ShareaholicBot = utils.BoolAddr(boolValue)
				}
				break
			case "site_24_x_7_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.Site24X7Bot = utils.BoolAddr(boolValue)
				}
				break
			case "siteimprove_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.SiteimproveBot = utils.BoolAddr(boolValue)
				}
				break
			case "site_lock_spider":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.SiteLockSpider = utils.BoolAddr(boolValue)
				}
				break
			case "slack_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.SlackBot = utils.BoolAddr(boolValue)
				}
				break
			case "sogou_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.SogouBot = utils.BoolAddr(boolValue)
				}
				break
			case "soso_spider_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.SosoSpiderBot = utils.BoolAddr(boolValue)
				}
				break
			case "spatineo":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.Spatineo = utils.BoolAddr(boolValue)
				}
				break
			case "spring_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.SpringBot = utils.BoolAddr(boolValue)
				}
				break
			case "stackify":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.Stackify = utils.BoolAddr(boolValue)
				}
				break
			case "status_cake_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.StatusCakeBot = utils.BoolAddr(boolValue)
				}
				break
			case "stripe":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.Stripe = utils.BoolAddr(boolValue)
				}
				break
			case "sucuri_uptime_monitor_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.SucuriUptimeMonitorBot = utils.BoolAddr(boolValue)
				}
				break
			case "telegram_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.TelegramBot = utils.BoolAddr(boolValue)
				}
				break
			case "testomato_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.TestomatoBot = utils.BoolAddr(boolValue)
				}
				break
			case "the_find_crawler":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.TheFindCrawler = utils.BoolAddr(boolValue)
				}
				break
			case "twitter_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.TwitterBot = utils.BoolAddr(boolValue)
				}
				break
			case "uptime_robot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.UptimeRobot = utils.BoolAddr(boolValue)
				}
				break
			case "vkontakte_external_hit_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.VkontakteExternalHitBot = utils.BoolAddr(boolValue)
				}
				break
			case "w_3_c":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.W3C = utils.BoolAddr(boolValue)
				}
				break
			case "wordfence_central":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.WordfenceCentral = utils.BoolAddr(boolValue)
				}
				break
			case "workato":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.Workato = utils.BoolAddr(boolValue)
				}
				break
			case "xml_sitemaps":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.XMLSitemaps = utils.BoolAddr(boolValue)
				}
				break
			case "yahoo_inktomi_slurp_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.YahooInktomiSlurpBot = utils.BoolAddr(boolValue)
				}
				break
			case "yahoo_japan_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.YahooJapanBot = utils.BoolAddr(boolValue)
				}
				break
			case "yahoo_link_preview":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.YahooLinkPreview = utils.BoolAddr(boolValue)
				}
				break
			case "yahoo_seeker_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.YahooSeekerBot = utils.BoolAddr(boolValue)
				}
				break
			case "yahoo_slurp_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.YahooSlurpBot = utils.BoolAddr(boolValue)
				}
				break
			case "yandex_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.YandexBot = utils.BoolAddr(boolValue)
				}
				break
			case "yisou_spider_commercial":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.YisouSpiderCommercial = utils.BoolAddr(boolValue)
				}
				break
			case "yodao_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.YodaoBot = utils.BoolAddr(boolValue)
				}
				break
			case "zendesk_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.ZendeskBot = utils.BoolAddr(boolValue)
				}
				break
			case "zoho_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.ZohoBot = utils.BoolAddr(boolValue)
				}
				break
			case "zum_bot":
				{
					boolValue, _ := strconv.ParseBool(value.(string))
					allowKnownBots.ZumBot = utils.BoolAddr(boolValue)
				}
				break
			}
		}
	}
	return allowKnownBots
}
