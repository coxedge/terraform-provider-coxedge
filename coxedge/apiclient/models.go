/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package apiclient

import (
	"time"
)

type TaskStatusResponse struct {
	TaskId     string `json:"taskId,omitempty"`
	TaskStatus string `json:"taskStatus,omitempty"`
}

type TaskStatus struct {
	Data struct {
		TaskId     string `json:"id,omitempty"`
		TaskStatus string `json:"status,omitempty"`
		Result     struct {
			Id   string `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"result,omitempty"`
	}
}

type ServiceConnection struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	ServiceCode string `json:"serviceCode,omitempty"`
}

type Role struct {
	Id        string         `json:"id,omitempty"`
	Name      string         `json:"name,omitempty"`
	IsDefault bool           `json:"isDefault,omitempty"`
	Users     []IdOnlyHelper `json:"users,omitempty"`
}

type Environment struct {
	Id                string            `json:"id,omitempty"`
	Name              string            `json:"name,omitempty"`
	Description       string            `json:"description,omitempty"`
	Membership        string            `json:"membership,omitempty"`
	CreationDate      string            `json:"creationDate,omitempty"`
	Organization      Organization      `json:"organization,omitempty"`
	ServiceConnection ServiceConnection `json:"serviceConnection,omitempty"`
	Roles             []Role            `json:"roles,omitempty"`
}
type WrappedEnvironments struct {
	Data []Environment `json:"data"`
}
type WrappedEnvironment struct {
	Data Environment `json:"data"`
}

type ParentOrganization struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Organization struct {
	Id                   string              `json:"id,omitempty"`
	Name                 string              `json:"name,omitempty"`
	EntryPoint           string              `json:"entryPoint,omitempty"`
	BillableStartDate    time.Time           `json:"billableStartDate,omitempty"`
	BillingDay           int8                `json:"billingDay,omitempty"`
	IsBillable           bool                `json:"isBillable,omitempty"`
	BillingMode          string              `json:"billingMode,omitempty"`
	IsReseller           bool                `json:"isReseller,omitempty"`
	Tags                 []string            `json:"tags,omitempty"`
	Parent               ParentOrganization  `json:"parent,omitempty"`
	Environments         []Environment       `json:"environments,omitempty"`
	Roles                []Role              `json:"roles,omitempty"`
	ServiceConnections   []ServiceConnection `json:"serviceConnections,omitempty"`
	ResourceCommitments  interface{}         `json:"resourceCommitments,omitempty"`
	Users                []User              `json:"users,omitempty"`
	Notes                string              `json:"notes,omitempty"`
	IsDbAuthentication   bool                `json:"isDbAuthentication,omitempty"`
	IsLdapAuthentication bool                `json:"isLdapAuthentication,omitempty"`
	IsTrial              bool                `json:"isTrial,omitempty"`
	CustomDomain         interface{}         `json:"customDomain,omitempty"`
}

type WrappedOrganizations struct {
	Data []Organization `json:"data"`
}
type WrappedOrganization struct {
	Data Organization `json:"data"`
}

type User struct {
	Id           string             `json:"id,omitempty"`
	UserName     string             `json:"userName,omitempty"`
	FirstName    string             `json:"firstName,omitempty"`
	LastName     string             `json:"lastName,omitempty"`
	Email        string             `json:"email,omitempty"`
	CreationDate time.Time          `json:"creationDate,omitempty"`
	Status       string             `json:"status,omitempty"`
	Organization ParentOrganization `json:"organization,omitempty"`
	Roles        []Role             `json:"roles,omitempty"`
}
type WrappedUsers struct {
	Data []User `json:"data"`
}
type WrappedUser struct {
	Data User `json:"data"`
}

// Workloads
type Workload struct {
	AddImagePullCreationsOption bool                          `json:"addImagePullCredentialsOption,omitempty"`
	AnycastIpAddress            string                        `json:"anycastIpAddress,omitempty"`
	Commands                    []string                      `json:"commands,omitempty"`
	ContainerEmail              string                        `json:"containerEmail,omitempty"`
	ContainerServer             string                        `json:"containerServer,omitempty"`
	ContainerUsername           string                        `json:"containerUsername,omitempty"`
	CPU                         string                        `json:"cpu,omitempty"`
	Created                     string                        `json:"created,omitempty"`
	Deployments                 []WorkloadAutoscaleDeployment `json:"deployments,omitempty"`
	EnvironmentVariables        []WorkloadEnvironmentVariable `json:"environmentVariables,omitempty"`
	FirstBootSshKey             string                        `json:"firstBootSshKey,omitempty"`
	Id                          string                        `json:"id,omitempty"`
	Image                       string                        `json:"image,omitempty"`
	IsRemoteManagementEnabled   bool                          `json:"isRemoteManagementEnabled,omitempty"`
	Memory                      string                        `json:"memory,omitempty"`
	Name                        string                        `json:"name,omitempty"`
	Network                     string                        `json:"network,omitempty"`
	PersistentStorages          []WorkloadPersistentStorage   `json:"persistentStorages,omitempty"`
	Ports                       []WorkloadPort                `json:"ports,omitempty"`
	SecretEnvironmentVariables  []WorkloadEnvironmentVariable `json:"secretEnvironmentVariables,omitempty"`
	Slug                        string                        `json:"slug,omitempty"`
	Specs                       string                        `json:"specs,omitempty"`
	StackId                     string                        `json:"stackId,omitempty"`
	Status                      string                        `json:"status,omitempty"`
	Type                        string                        `json:"type,omitempty"`
	Version                     string                        `json:"version,omitempty"`
}

type WorkloadAutoscaleDeployment struct {
	Name               string   `json:"name,omitempty"`
	Pops               []string `json:"pops,omitempty"`
	EnableAutoScaling  bool     `json:"enableAutoScaling,omitempty"`
	InstancesPerPop    int      `json:"instancesPerPop,string,omitempty"`
	MaxInstancesPerPop int      `json:"maxInstancesPerPop,string,omitempty"`
	MinInstancesPerPop int      `json:"minInstancesPerPop,string,omitempty"`
	CPUUtilization     int      `json:"cpuUtilization,omitempty"`
}

type WorkloadEnvironmentVariable struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type WorkloadPort struct {
	Protocol       string `json:"protocol,omitempty"`
	PublicPort     string `json:"publicPort,omitempty"`
	PublicPortDesc string `json:"publicPortDesc,omitempty"`
	PublicPortSrc  string `json:"publicPortSrc,omitempty"`
}

type WorkloadPersistentStorage struct {
	Path string `json:"path,omitempty"`
	Size int    `json:"size,omitempty"`
}

type WrappedWorkloads struct {
	Data []Workload `json:"data"`
}
type WrappedWorkload struct {
	Data Workload `json:"data"`
}

//Network Policy Rules
type NetworkPolicyRule struct {
	Id              string `json:"id,omitempty"`
	StackId         string `json:"stackId,omitempty"`
	WorkloadId      string `json:"workloadId,omitempty"`
	NetworkPolicyId string `json:"networkPolicyId,omitempty"`
	Description     string `json:"description,omitempty"`
	Type            string `json:"type,omitempty"`
	Source          string `json:"source,omitempty"`
	Action          string `json:"action,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	PortRange       string `json:"portRange,omitempty"`
}

type WrappedNetworkPolicyRule struct {
	Data NetworkPolicyRule
}

type WrappedNetworkPolicyRules struct {
	Data []NetworkPolicyRule
}

type Image struct {
	IncludeOnlySystemImages bool   `json:"includeOnlySystemImages,omitempty"`
	Id                      string `json:"id,omitempty"`
	StackId                 string `json:"stackId,omitempty"`
	Family                  string `json:"family,omitempty"`
	Tag                     string `json:"tag,omitempty"`
	Slug                    string `json:"slug,omitempty"`
	Status                  string `json:"status,omitempty"`
	CreatedAt               string `json:"createdAt,omitempty"`
	Description             string `json:"description,omitempty"`
	Reference               string `json:"reference,omitempty"`
}

type WrappedImage struct {
	Data Image `json:"data"`
}

type WrappedImages struct {
	Data []Image `json:"data"`
}

//Site
type Site struct {
	Id              string               `json:"id,omitempty"`
	StackId         string               `json:"stackId,omitempty"`
	Domain          string               `json:"domain,omitempty"`
	Status          string               `json:"status,omitempty"`
	CreatedAt       string               `json:"createdAt,omitempty"`
	UpdatedAt       string               `json:"updatedAt,omitempty"`
	Services        []string             `json:"services,omitempty"`
	EdgeAddress     string               `json:"edgeAddress,omitempty"`
	AnycastIp       string               `json:"anycastIp,omitempty"`
	DeliveryDomains []SiteDeliveryDomain `json:"deliveryDomains,omitempty"`
}

type SiteDeliveryDomain struct {
	Domain      string `json:"domain,omitempty"`
	ValidatedAt string `json:"validatedAt,omitempty"`
}

type WrappedSite struct {
	Data Site `json:"data"`
}

type WrappedSites struct {
	Data []Site `json:"data"`
}

//Origin Settings
type OriginSettings struct {
	EnvironmentName          string               `json:"-"`
	Id                       string               `json:"id,omitempty"`
	StackId                  string               `json:"stackId,omitempty"`
	ScopeConfigurationId     string               `json:"scopeConfigurationId,omitempty"`
	Domain                   string               `json:"domain,omitempty"`
	WebSocketsEnabled        bool                 `json:"webSocketsEnabled,omitempty"`
	SSLValidationEnabled     bool                 `json:"sslValidationEnabled,omitempty"`
	PullProtocol             string               `json:"pullProtocol,omitempty"`
	HostHeader               string               `json:"hostHeader,omitempty"`
	Origin                   OriginSettingsOrigin `json:"origin,omitempty"`
	BackupOriginEnabled      bool                 `json:"backupOriginEnabled,omitempty"`
	BackupOriginExcludeCodes []string             `json:"backupOriginExcludeCodes,omitempty"`
	BackupOrigin             OriginSettingsOrigin `json:"backupOrigin,omitempty"`
}

type OriginSettingsOrigin struct {
	Id                    string `json:"id,omitempty"`
	Address               string `json:"address,omitempty"`
	AuthMethod            string `json:"authMethod,omitempty"`
	Username              string `json:"username,omitempty"`
	Password              string `json:"password,omitempty"`
	CommonCertificateName string `json:"commonCertificateName,omitempty"`
}

type WrappedOriginSettings struct {
	Data OriginSettings
}

type WrappedOriginSettingsSet struct {
	Data []OriginSettings
}

//Delivery Domain
type DeliveryDomain struct {
	Id        string `json:"id,omitempty"`
	Domain    string `json:"domain,omitempty"`
	SiteId    string `json:"siteId,omitempty"`
	StackId   string `json:"stackId,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type WrappedDeliveryDomain struct {
	Data DeliveryDomain `json:"data"`
}

type WrappedDeliveryDomains struct {
	Data []DeliveryDomain `json:"data"`
}

//CDNSettings
type CDNSettings struct {
	EnvironmentName               string   `json:"-"`
	SiteId                        string   `json:"-"`
	CacheExpirePolicy             string   `json:"cacheExpirePolicy,omitempty"`
	CacheTtl                      int      `json:"cacheTtl,omitempty"`
	QueryStringControl            string   `json:"queryStringControl,omitempty"`
	CustomCachedQueryStrings      []string `json:"customCachedQueryStrings,omitempty"`
	DynamicCachingByHeaderEnabled bool     `json:"dynamicCachingByHeaderEnabled,omitempty"`
	CustomCacheHeaders            []string `json:"customCacheHeaders,omitempty"`
	GzipCompressionEnabled        bool     `json:"gzipCompressionEnabled,omitempty"`
	GzipCompressionLevel          int      `json:"gzipCompressionLevel,omitempty"`
	ContentPersistenceEnabled     bool     `json:"contentPersistenceEnabled,omitempty"`
	MaximumStaleFileTtl           int      `json:"maximumStaleFileTtl,omitempty"`
	VaryHeaderEnabled             bool     `json:"varyHeaderEnabled,omitempty"`
	BrowserCacheTtl               int      `json:"browserCacheTtl,omitempty"`
	CorsHeaderEnabled             bool     `json:"corsHeaderEnabled,omitempty"`
	AllowedCorsOrigins            string   `json:"allowedCorsOrigins,omitempty"`
	OriginsToAllowCors            []string `json:"originsToAllowCors,omitempty"`
	Http2SupportEnabled           bool     `json:"http2SupportEnabled,omitempty"`
	LinkHeader                    string   `json:"linkHeader,omitempty"`
	CanonicalHeaderEnabled        bool     `json:"canonicalHeaderEnabled,omitempty"`
	CanonicalHeader               string   `json:"canonicalHeader,omitempty"`
	UrlCachingEnabled             bool     `json:"urlCachingEnabled,omitempty"`
	UrlCachingTtl                 int      `json:"urlCachingTtl,omitempty"`
}

type WrappedCDNSettings struct {
	Data CDNSettings `json:"data"`
}

type WrappedCDNSettingsSet struct {
	Data []CDNSettings `json:"data"`
}

type CDNPurgeOptions struct {
	PurgeType string `json:"purgeType"`
	Items     []struct {
		URL             string   `json:"url,omitempty"`
		Recursive       bool     `json:"recursive,omitempty"`
		InvalidateOnly  bool     `json:"invalidateOnly,omitempty"`
		PurgeAllDynamic bool     `json:"purgeAllDynamic,omitempty"`
		Headers         []string `json:"headers,omitempty"`
		PurgeSelector   struct {
			SelectorName           string `json:"selectorName,omitempty"`
			SelectorValue          string `json:"selectorValue,omitempty"`
			SelectorType           string `json:"selectorType,omitempty"`
			SelectorValueDelimiter string `json:"selectorValueDelimiter,omitempty"`
		} `json:"purgeSelector,omitempty"`
	} `json:"items,omitempty"`
}

//WAF
type WAFSettings struct {
	EnvironmentName             string                         `json:"-"`
	Id                          string                         `json:"id,omitempty"`
	StackId                     string                         `json:"stackId,omitempty"`
	Domain                      string                         `json:"domain,omitempty"`
	APIUrls                     []string                       `json:"apiUrls,omitempty"`
	DdosSettings                WAFDdosSettings                `json:"ddosSettings,omitempty"`
	MonitoringEnabled           bool                           `json:"monitoringEnabled,omitempty"`
	OwaspThreats                WAFOwaspThreats                `json:"owaspThreats,omitempty"`
	UserAgents                  WAFUserAgents                  `json:"userAgents,omitempty"`
	Csrf                        bool                           `json:"csrf,omitempty"`
	TrafficSources              WAFTrafficSources              `json:"trafficSources,omitempty"`
	AntiAutomationBotProtection WAFAntiAutomationBotProtection `json:"antiAutomationBotProtection,omitempty"`
	SpamAndAbuseForm            bool                           `json:"spamAndAbuseForm,omitempty"`
	BehavioralWaf               WAFBehavioralWaf               `json:"behavioralWaf,omitempty"`
	CmsProtection               WAFCmsProtection               `json:"cmsProtection,omitempty"`
	AllowKnownBots              WAFAllowKnownBots              `json:"allowKnownBots,omitempty"`
}
type WAFDdosSettings struct {
	GlobalThreshold         int `json:"globalThreshold,omitempty"`
	BurstThreshold          int `json:"burstThreshold,omitempty"`
	SubSecondBurstThreshold int `json:"subSecondBurstThreshold,omitempty"`
}
type WAFOwaspThreats struct {
	SQLInjection                        bool `json:"sqlInjection,omitempty"`
	XSSAttack                           bool `json:"xssAttack,omitempty"`
	RemoteFileInclusion                 bool `json:"remoteFileInclusion,omitempty"`
	WordpressWafRuleset                 bool `json:"wordpressWafRuleset,omitempty"`
	ApacheStrutsExploit                 bool `json:"apacheStrutsExploit,omitempty"`
	LocalFileInclusion                  bool `json:"localFileInclusion,omitempty"`
	CommonWebApplicationVulnerabilities bool `json:"commonWebApplicationVulnerabilities,omitempty"`
	WebShellExecutionAttempt            bool `json:"webShellExecutionAttempt,omitempty"`
	ResponseHeaderInjection             bool `json:"responseHeaderInjection,omitempty"`
	OpenRedirect                        bool `json:"openRedirect,omitempty"`
	ShellInjection                      bool `json:"shellInjection,omitempty"`
}
type WAFUserAgents struct {
	BlockInvalidUserAgents bool `json:"blockInvalidUserAgents,omitempty"`
	BlockUnknownUserAgents bool `json:"blockUnknownUserAgents,omitempty"`
}
type WAFTrafficSources struct {
	ViaTorNodes                      bool `json:"viaTorNodes,omitempty"`
	ViaProxyNetworks                 bool `json:"viaProxyNetworks,omitempty"`
	ViaHostingServices               bool `json:"viaHostingServices,omitempty"`
	ViaVpn                           bool `json:"viaVpn,omitempty"`
	ConvictedBotTraffic              bool `json:"convictedBotTraffic,omitempty"`
	SuspiciousTrafficByLocalIPFormat bool `json:"suspiciousTrafficByLocalIpFormat,omitempty"`
}
type WAFAntiAutomationBotProtection struct {
	ForceBrowserValidationOnTrafficAnomalies bool `json:"forceBrowserValidationOnTrafficAnomalies,omitempty"`
	ChallengeAutomatedClients                bool `json:"challengeAutomatedClients,omitempty"`
	ChallengeHeadlessBrowsers                bool `json:"challengeHeadlessBrowsers,omitempty"`
	AntiScraping                             bool `json:"antiScraping,omitempty"`
}
type WAFBehavioralWaf struct {
	SpamProtection                        bool `json:"spamProtection,omitempty"`
	BlockProbingAndForcedBrowsing         bool `json:"blockProbingAndForcedBrowsing,omitempty"`
	ObfuscatedAttacksAndZeroDayMitigation bool `json:"obfuscatedAttacksAndZeroDayMitigation,omitempty"`
	RepeatedViolations                    bool `json:"repeatedViolations,omitempty"`
	BruteForceProtection                  bool `json:"bruteForceProtection,omitempty"`
}
type WAFCmsProtection struct {
	WhiteListWordpress bool `json:"whiteListWordpress,omitempty"`
	WhiteListModx      bool `json:"whiteListModx,omitempty"`
	WhiteListDrupal    bool `json:"whiteListDrupal,omitempty"`
	WhiteListJoomla    bool `json:"whiteListJoomla,omitempty"`
	WhiteMagento       bool `json:"whiteMagento,omitempty"`
	WhiteListOriginIP  bool `json:"whiteListOriginIp,omitempty"`
	WhiteListUmbraco   bool `json:"whiteListUmbraco,omitempty"`
}
type WAFAllowKnownBots struct {
	InternetArchiveBot bool `json:"Internet Archive Bot,omitempty"`
}

type WrappedWAFSettings struct {
	Data WAFSettings `json:"data"`
}

type WrappedWAFSettingsSet struct {
	Data []WAFSettings `json:"data"`
}

//Firewall Rule
type FirewallRule struct {
	Action          string `json:"action,omitempty"`
	Enabled         bool   `json:"enabled,omitempty"`
	Id              string `json:"id,omitempty"`
	IpEnd           string `json:"ipEnd,omitempty"`
	IpStart         string `json:"ipStart,omitempty"`
	Name            string `json:"name,omitempty"`
	SiteId          string `json:"siteId,omitempty"`
	EnvironmentName string `json:"-"`
}

type WrappedFirewallRule struct {
	Data FirewallRule `json:"data"`
}

type WrappedFirewallRules struct {
	Data []FirewallRule `json:"data"`
}

//Script
type Script struct {
	Id        string   `json:"id,omitempty"`
	StackId   string   `json:"stackId,omitempty"`
	SiteId    string   `json:"siteId,omitempty"`
	Name      string   `json:"name,omitempty"`
	CreatedAt string   `json:"createdAt,omitempty"`
	UpdatedAt string   `json:"updatedAt,omitempty"`
	Code      string   `json:"code,omitempty"`
	Version   string   `json:"version,omitempty"`
	Routes    []string `json:"routes,omitempty"`
}

type WrappedScript struct {
	Data Script `json:"data,omitempty"`
}
type WrappedScripts struct {
	Data []Script `json:"data,omitempty"`
}

//Edge Logic
type EdgeLogic struct {
	AllowEmptyReferrer        bool     `json:"allowEmptyReferrer,omitempty"`
	ForceWwwEnabled           bool     `json:"forceWwwEnabled,omitempty"`
	Id                        string   `json:"id,omitempty"`
	PseudoStreamingEnabled    bool     `json:"pseudoStreamingEnabled,omitempty"`
	ReferrerList              []string `json:"referrerList,omitempty"`
	ReferrerProtectionEnabled bool     `json:"referrerProtectionEnabled,omitempty"`
	RobotTxtEnabled           bool     `json:"robotTxtEnabled,omitempty"`
	RobotTxtFile              string   `json:"robotTxtFile,omitempty"`
	ScopeId                   string   `json:"scopeId,omitempty"`
	StackId                   string   `json:"stackId,omitempty"`
}
type WrappedEdgeLogic struct {
	Data EdgeLogic `json:"data,omitempty"`
}
