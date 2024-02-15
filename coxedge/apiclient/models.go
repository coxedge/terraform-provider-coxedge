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

type OrganizationBillingInfo struct {
	Id                       string       `json:"id,omitempty"`
	Organization             IdOnlyHelper `json:"organization,omitempty"`
	BillingProvider          IdOnlyHelper `json:"billingProvider,omitempty"`
	CardType                 string       `json:"cardType,omitempty"`
	CardMaskedNumber         string       `json:"cardMaskedNumber,omitempty"`
	CardName                 string       `json:"cardName,omitempty"`
	CardExp                  string       `json:"cardExp,omitempty"`
	BillingAddressLineOne    string       `json:"billingAddressLineOne,omitempty"`
	BillingAddressLineTwo    string       `json:"billingAddressLineTwo,omitempty"`
	BillingAddressCity       string       `json:"billingAddressCity,omitempty"`
	BillingAddressProvince   string       `json:"billingAddressProvince,omitempty"`
	BillingAddressPostalCode string       `json:"billingAddressPostalCode,omitempty"`
	BillingAddressCountry    string       `json:"billingAddressCountry,omitempty"`
}

type WrappedOrganizationBillingInfo struct {
	Data OrganizationBillingInfo `json:"data"`
}

type Roles struct {
	Id           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	IsSystem     bool   `json:"isSystem,omitempty"`
	DefaultScope string `json:"defaultScope,omitempty"`
}

type WrappedRolesData struct {
	Data []Roles `json:"data"`
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
	UserData                    string                        `json:"userData,omitempty"`
	Id                          string                        `json:"id,omitempty"`
	Image                       string                        `json:"image,omitempty"`
	IsRemoteManagementEnabled   bool                          `json:"isRemoteManagementEnabled,omitempty"`
	Memory                      string                        `json:"memory,omitempty"`
	Name                        string                        `json:"name,omitempty"`
	NetworkNames                []string                      `json:"networkNames,omitempty"`
	PersistentStorages          []WorkloadPersistentStorage   `json:"persistentStorages,omitempty"`
	Ports                       []WorkloadPort                `json:"ports,omitempty"`
	SecretEnvironmentVariables  []WorkloadEnvironmentVariable `json:"secretEnvironmentVariables,omitempty"`
	Slug                        string                        `json:"slug,omitempty"`
	Specs                       string                        `json:"specs,omitempty"`
	StackId                     string                        `json:"stackId,omitempty"`
	Status                      string                        `json:"status,omitempty"`
	Type                        string                        `json:"type,omitempty"`
	Version                     string                        `json:"version,omitempty"`
	NetworkInterfaces           []NetworkInterface            `json:"networkInterfaces,omitempty"`
	ProbeConfiguration          string                        `json:"probeConfiguration,omitempty"`
	LivenessProbe               LivenessProbe                 `json:"livenessProbe,omitempty"`
	ReadinessProbe              ReadinessProbe                `json:"readinessProbe,omitempty"`
}

type NetworkInterface struct {
	VpcSlug    string `json:"vpcSlug"`
	IpFamilies string `json:"ipFamilies"`
	SubnetSlug string `json:"subnetSlug"`
	IsPublicIP bool   `json:"isPublicIP"`
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

type LivenessProbe struct {
	InitialDelaySeconds *int       `json:"initialDelaySeconds,omitempty"`
	TimeoutSeconds      *int       `json:"timeoutSeconds,omitempty"`
	PeriodSeconds       *int       `json:"periodSeconds,omitempty"`
	SuccessThreshold    *int       `json:"successThreshold,omitempty"`
	FailureThreshold    *int       `json:"failureThreshold,omitempty"`
	Protocol            string     `json:"protocol,omitempty"`
	TcpSocket           *TCPSocket `json:"tcpSocket,omitempty"'`
	HttpGet             *HTTPGet   `json:"httpGet,omitempty"'`
}

type ReadinessProbe struct {
	InitialDelaySeconds *int       `json:"initialDelaySeconds,omitempty"`
	TimeoutSeconds      *int       `json:"timeoutSeconds,omitempty"`
	PeriodSeconds       *int       `json:"periodSeconds,omitempty"`
	SuccessThreshold    *int       `json:"successThreshold,omitempty"`
	FailureThreshold    *int       `json:"failureThreshold,omitempty"`
	Protocol            string     `json:"protocol,omitempty"`
	TcpSocket           *TCPSocket `json:"tcpSocket,omitempty"'`
	HttpGet             *HTTPGet   `json:"httpGet,omitempty"'`
}

type TCPSocket struct {
	Port *int `json:"port,omitempty"`
}

type HTTPGet struct {
	HttpHeaders []HTTPHeaders `json:"httpHeaders,omitempty"`
	Scheme      string        `json:"scheme,omitempty"`
	Path        string        `json:"path,omitempty"`
	Port        *int          `json:"port,omitempty"`
}

type HTTPHeaders struct {
	HeaderName  string `json:"headerName,omitempty"`
	HeaderValue string `json:"headerValue,omitempty"`
}

type WrappedWorkloads struct {
	Data []Workload `json:"data"`
}
type WrappedWorkload struct {
	Data Workload `json:"data"`
}

type WorkloadInstance struct {
	StackId         string   `json:"stackId"`
	WorkloadId      string   `json:"workloadId"`
	Name            string   `json:"name"`
	IPAddress       []string `json:"ipAddress"`
	PublicIPAddress string   `json:"publicIpAddress"`
	Location        string   `json:"location"`
	CreatedDate     string   `json:"createdDate"`
	StartedDate     string   `json:"startedDate"`
	Id              string   `json:"id"`
	Status          string   `json:"status"`
}

//Network Policy Rules
type NetworkPolicyRule struct {
	Id              string   `json:"id,omitempty"`
	StackId         string   `json:"stackId,omitempty"`
	WorkloadId      string   `json:"workloadId,omitempty"`
	NetworkPolicyId string   `json:"networkPolicyId,omitempty"`
	Description     string   `json:"description,omitempty"`
	Type            string   `json:"type,omitempty"`
	SourceIps       []string `json:"sourceIps,omitempty"`
	Action          string   `json:"action,omitempty"`
	Protocol        string   `json:"protocol,omitempty"`
	Ports           []string `json:"ports,omitempty"`
}
type WrapperWorkloadInstances struct {
	Data []WorkloadInstance
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
	WebSocketsEnabled        *bool                `json:"webSocketsEnabled,omitempty"`
	SSLValidationEnabled     *bool                `json:"sslValidationEnabled,omitempty"`
	PullProtocol             string               `json:"pullProtocol,omitempty"`
	HostHeader               string               `json:"hostHeader,omitempty"`
	Origin                   OriginSettingsOrigin `json:"origin,omitempty"`
	BackupOriginEnabled      *bool                `json:"backupOriginEnabled,omitempty"`
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
	SiteId                        string   `json:"siteId,omitempty"`
	Id                            string   `json:"siteId,omitempty"`
	CacheExpirePolicy             string   `json:"cacheExpirePolicy,omitempty"`
	CacheTtl                      int      `json:"cacheTtl,omitempty"`
	QueryStringControl            string   `json:"queryStringControl,omitempty"`
	CustomCachedQueryStrings      []string `json:"customCachedQueryStrings,omitempty"`
	DynamicCachingByHeaderEnabled *bool    `json:"dynamicCachingByHeaderEnabled,omitempty"`
	CustomCacheHeaders            []string `json:"customCachedHeaders,omitempty"`
	GzipCompressionEnabled        *bool    `json:"gzipCompressionEnabled,omitempty"`
	GzipCompressionLevel          int      `json:"gzipCompressionLevel,omitempty"`
	ContentPersistenceEnabled     *bool    `json:"contentPersistenceEnabled,omitempty"`
	MaximumStaleFileTtl           int      `json:"maximumStaleFileTtl,omitempty"`
	VaryHeaderEnabled             *bool    `json:"varyHeaderEnabled,omitempty"`
	BrowserCacheTtl               int      `json:"browserCacheTtl,omitempty"`
	CorsHeaderEnabled             *bool    `json:"corsHeaderEnabled,omitempty"`
	AllowedCorsOrigins            string   `json:"allowedCorsOrigins,omitempty"`
	OriginsToAllowCors            []string `json:"originsToAllowCors,omitempty"`
	Http2SupportEnabled           *bool    `json:"http2SupportEnabled,omitempty"`
	Http2ServerPushEnabled        *bool    `json:"http2ServerPushEnabled,omitempty"`
	LinkHeader                    string   `json:"linkHeader,omitempty"`
	CanonicalHeaderEnabled        *bool    `json:"canonicalHeaderEnabled,omitempty"`
	CanonicalHeader               string   `json:"canonicalHeader,omitempty"`
	UrlCachingEnabled             *bool    `json:"urlCachingEnabled,omitempty"`
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
	MonitoringModeEnabled       *bool                          `json:"monitoringModeEnabled,omitempty"`
	OwaspThreats                WAFOwaspThreats                `json:"owaspThreats,omitempty"`
	GeneralPolicies             WAFGeneralPolicies             `json:"generalPolicies,omitempty"`
	TrafficSources              WAFTrafficSources              `json:"trafficSources,omitempty"`
	AntiAutomationBotProtection WAFAntiAutomationBotProtection `json:"antiAutomationBotProtection,omitempty"`
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
	SQLInjection                        *bool `json:"sqlInjection,omitempty"`
	XSSAttack                           *bool `json:"xssAttack,omitempty"`
	ShellShockAttack                    *bool `json:"shellShockAttack,omitempty"`
	RemoteFileInclusion                 *bool `json:"remoteFileInclusion,omitempty"`
	ApacheStrutsExploit                 *bool `json:"apacheStrutsExploit,omitempty"`
	LocalFileInclusion                  *bool `json:"localFileInclusion,omitempty"`
	CommonWebApplicationVulnerabilities *bool `json:"commonWebApplicationVulnerabilities,omitempty"`
	WebShellExecutionAttempt            *bool `json:"webShellExecutionAttempt,omitempty"`
	ProtocolAttack                      *bool `json:"protocolAttack,omitempty"`
	Csrf                                *bool `json:"csrf,omitempty"`
	OpenRedirect                        *bool `json:"openRedirect,omitempty"`
	ShellInjection                      *bool `json:"shellInjection,omitempty"`
	CodeInjection                       *bool `json:"codeInjection,omitempty"`
	SensitiveDataExposure               *bool `json:"sensitiveDataExposure,omitempty"`
	XmlExternalEntity                   *bool `json:"xmlExternalEntity,omitempty"`
	PersonalIdentifiableInfo            *bool `json:"personalIdentifiableInfo,omitempty"`
	ServerSideTemplateInjection         *bool `json:"serverSideTemplateInjection,omitempty"`
}
type WAFGeneralPolicies struct {
	BlockInvalidUserAgents *bool `json:"blockInvalidUserAgents,omitempty"`
	BlockUnknownUserAgents *bool `json:"blockUnknownUserAgents,omitempty"`
	HttpMethodValidation   *bool `json:"httpMethodValidation,omitempty"`
}
type WAFTrafficSources struct {
	ViaTorNodes                    *bool `json:"viaTorNodes,omitempty"`
	ViaProxyNetworks               *bool `json:"viaProxyNetworks,omitempty"`
	ViaHostingServices             *bool `json:"viaHostingServices,omitempty"`
	ViaVpn                         *bool `json:"viaVpn,omitempty"`
	ConvictedBotTraffic            *bool `json:"convictedBotTraffic,omitempty"`
	TrafficFromSuspiciousNatRanges *bool `json:"trafficFromSuspiciousNatRanges,omitempty"`
	ExternalReputationBlockList    *bool `json:"externalReputationBlockList,omitempty"`
	TrafficViaCDN                  *bool `json:"trafficViaCDN,omitempty"`
}
type WAFAntiAutomationBotProtection struct {
	ForceBrowserValidationOnTrafficAnomalies *bool `json:"forceBrowserValidationOnTrafficAnomalies,omitempty"`
	ChallengeAutomatedClients                *bool `json:"challengeAutomatedClients,omitempty"`
	ChallengeHeadlessBrowsers                *bool `json:"challengeHeadlessBrowsers,omitempty"`
	AntiScraping                             *bool `json:"antiScraping,omitempty"`
}
type WAFBehavioralWaf struct {
	SpamProtection                        *bool `json:"spamProtection,omitempty"`
	BlockProbingAndForcedBrowsing         *bool `json:"blockProbingAndForcedBrowsing,omitempty"`
	ObfuscatedAttacksAndZeroDayMitigation *bool `json:"obfuscatedAttacksAndZeroDayMitigation,omitempty"`
	RepeatedViolations                    *bool `json:"repeatedViolations,omitempty"`
	BruteForceProtection                  *bool `json:"bruteForceProtection,omitempty"`
}
type WAFCmsProtection struct {
	WordpressWafRuleset *bool `json:"wordpressWafRuleset,omitempty"`
	WhiteListWordpress  *bool `json:"whiteListWordpress,omitempty"`
	WhiteListModx       *bool `json:"whiteListModx,omitempty"`
	WhiteListDrupal     *bool `json:"whiteListDrupal,omitempty"`
	WhiteListJoomla     *bool `json:"whiteListJoomla,omitempty"`
	WhiteMagento        *bool `json:"whiteMagento,omitempty"`
	WhiteListOriginIP   *bool `json:"whiteListOriginIp,omitempty"`
	WhiteListUmbraco    *bool `json:"whiteListUmbraco,omitempty"`
}
type WAFAllowKnownBots struct {
	AcquiaUptime                      *bool `json:"Acquia Uptime,omitempty"`
	AddSearchBot                      *bool `json:"AddSearch Bot,omitempty"`
	AdestraBot                        *bool `json:"Adestra bot,omitempty"`
	AdjustServers                     *bool `json:"Adjust Servers,omitempty"`
	AhrefsBot                         *bool `json:"Ahrefs Bot,omitempty"`
	AlertaBot                         *bool `json:"Alerta Bot,omitempty"`
	AlexaIaArchiver                   *bool `json:"Alexa ia archiver,omitempty"`
	AlexaTechnologies                 *bool `json:"Alexa technologies,omitempty"`
	AmazonRoute53HealthCheckService   *bool `json:"Amazon Route53 Health Check Service,omitempty"`
	Applebot                          *bool `json:"Applebot,omitempty"`
	AppleNewsBot                      *bool `json:"AppleNewsBot,omitempty"`
	AskJeevesBot                      *bool `json:"Ask Jeeves bot,omitempty"`
	AudistoBot                        *bool `json:"Audisto Bot,omitempty"`
	BaiduSpiderBot                    *bool `json:"Baidu Spider bot,omitempty"`
	BaiduSpiderJapanBot               *bool `json:"Baidu Spider Japan bot,omitempty"`
	BinaryCanary                      *bool `json:"BinaryCanary,omitempty"`
	BitbucketWebhook                  *bool `json:"Bitbucket webhook,omitempty"`
	BlekkoScoutJetBot                 *bool `json:"Blekko ScoutJet bot,omitempty"`
	ChromeCompressionProxy            *bool `json:"Chrome Compression Proxy,omitempty"`
	Coccocbot                         *bool `json:"Coccocbot,omitempty"`
	CookieBot                         *bool `json:"CookieBot,omitempty"`
	Cybersource                       *bool `json:"Cybersource,omitempty"`
	DaumoaBot                         *bool `json:"Daumoa bot,omitempty"`
	DetectifyScanner                  *bool `json:"Detectify Scanner,omitempty"`
	DigiCertDCVBot                    *bool `json:"DigiCert DCV Bot,omitempty"`
	DotmicDotBotCommercial            *bool `json:"Dotmic DotBot - Commercial,omitempty"`
	DuckDuckGoBot                     *bool `json:"DuckDuckGo bot,omitempty"`
	FacebookExternalHitBot            *bool `json:"Facebook External Hit bot,omitempty"`
	FeederCo                          *bool `json:"Feeder.co,omitempty"`
	FeedPress                         *bool `json:"FeedPress,omitempty"`
	FeedWind                          *bool `json:"FeedWind,omitempty"`
	FreshpingMonitoring               *bool `json:"Freshping Monitoring,omitempty"`
	Geckoboard                        *bool `json:"Geckoboard,omitempty"`
	GhostInspector                    *bool `json:"GhostInspector,omitempty"`
	Gomez                             *bool `json:"Gomez,omitempty"`
	GooJapanBot                       *bool `json:"goo Japan bot,omitempty"`
	GoogleAdsBot                      *bool `json:"Google ads bot,omitempty"`
	GoogleBot                         *bool `json:"Google bot,omitempty"`
	GoogleCloudMonitoringBot          *bool `json:"Google Cloud Monitoring bot,omitempty"`
	GoogleFeedFetcherBot              *bool `json:"Google FeedFetcher bot,omitempty"`
	GoogleImageBot                    *bool `json:"Google Image bot,omitempty"`
	GoogleImageProxy                  *bool `json:"Google Image Proxy,omitempty"`
	GoogleMediapartnersBot            *bool `json:"Google Mediapartners bot,omitempty"`
	GoogleMobileAdsBot                *bool `json:"Google Mobile Ads Bot,omitempty"`
	GoogleNewsBot                     *bool `json:"Google News bot,omitempty"`
	GooglePageSpeedInsights           *bool `json:"Google Page Speed Insights,omitempty"`
	GoogleStructuredDataTestingTool   *bool `json:"Google Structured Data Testing Tool,omitempty"`
	GoogleVerificationBot             *bool `json:"Google verification bot,omitempty"`
	GoogleVideoBot                    *bool `json:"Google Video bot,omitempty"`
	GoogleWebLight                    *bool `json:"Google Web Light,omitempty"`
	GrapeshotBotCommercial            *bool `json:"Grapeshot bot - Commercial,omitempty"`
	GreeJapanBot                      *bool `json:"Gree Japan bot,omitempty"`
	HetrixTools                       *bool `json:"Hetrix Tools,omitempty"`
	HiPay                             *bool `json:"HiPay,omitempty"`
	HyperspinBot                      *bool `json:"Hyperspin Bot,omitempty"`
	IASCrawlerCommercial              *bool `json:"IAS crawler - Commercial,omitempty"`
	InternetArchiveBot                *bool `json:"Internet Archive Bot,omitempty"`
	JetpackBot                        *bool `json:"Jetpack bot,omitempty"`
	JikeSpiderBot                     *bool `json:"JikeSpider bot,omitempty"`
	JWordJapanBot                     *bool `json:"JWord Japan bot,omitempty"`
	KAKAOUserAgent                    *bool `json:"KAKAO UserAgent,omitempty"`
	KyotoTohokuCrawler                *bool `json:"Kyoto Tohoku Crawler,omitempty"`
	LandauMediaSpider                 *bool `json:"Landau Media Spider,omitempty"`
	LetsEncrypt                       *bool `json:"Lets Encrypt,omitempty"`
	LineJapanBot                      *bool `json:"Line Japan bot,omitempty"`
	LinkedInBot                       *bool `json:"LinkedIn bot,omitempty"`
	LivedoorJapanBot                  *bool `json:"Livedoor Japan bot,omitempty"`
	MailRuBot                         *bool `json:"Mail.ru Bot,omitempty"`
	ManageWP                          *bool `json:"ManageWP,omitempty"`
	MicrosoftBingBot                  *bool `json:"Microsoft Bing bot,omitempty"`
	MicrosoftBingPreviewBot           *bool `json:"Microsoft Bing Preview bot,omitempty"`
	MicrosoftMSNBot                   *bool `json:"Microsoft MSN bot,omitempty"`
	MicrosoftSkypeBot                 *bool `json:"Microsoft Skype bot,omitempty"`
	MixiJapanBot                      *bool `json:"Mixi Japan bot,omitempty"`
	MobageJapanBot                    *bool `json:"Mobage Japan bot,omitempty"`
	NaverYetiBot                      *bool `json:"Naver Yeti bot,omitempty"`
	NewRelicBot                       *bool `json:"New Relic bot,omitempty"`
	OCNJapanBot                       *bool `json:"OCN Japan bot,omitempty"`
	PanoptaBot                        *bool `json:"Panopta bot,omitempty"`
	ParseLyScraper                    *bool `json:"parse.ly scraper,omitempty"`
	PayPalIPN                         *bool `json:"PayPal IPN,omitempty"`
	PetalBot                          *bool `json:"Petal bot,omitempty"`
	Pingdom                           *bool `json:"Pingdom,omitempty"`
	PinterestBot                      *bool `json:"Pinterest Bot,omitempty"`
	QwantifyBot                       *bool `json:"Qwantify bot,omitempty"`
	RogerBot                          *bool `json:"Roger bot,omitempty"`
	SagePay                           *bool `json:"SagePay,omitempty"`
	SectigoBot                        *bool `json:"Sectigo Bot,omitempty"`
	SemrushBot                        *bool `json:"Semrush Bot,omitempty"`
	ServerDensityServiceMonitoringBot *bool `json:"Server Density Service Monitoring bot,omitempty"`
	SeznamBot                         *bool `json:"Seznam bot,omitempty"`
	ShareaholicBot                    *bool `json:"Shareaholic Bot,omitempty"`
	Site24X7Bot                       *bool `json:"Site24X7 Bot,omitempty"`
	SiteimproveBot                    *bool `json:"Siteimprove bot,omitempty"`
	SiteLockSpider                    *bool `json:"SiteLock Spider,omitempty"`
	SlackBot                          *bool `json:"Slack bot,omitempty"`
	SogouBot                          *bool `json:"Sogou bot,omitempty"`
	SosoSpiderBot                     *bool `json:"Soso Spider bot,omitempty"`
	Spatineo                          *bool `json:"Spatineo,omitempty"`
	SpringBot                         *bool `json:"Spring Bot,omitempty"`
	Stackify                          *bool `json:"Stackify,omitempty"`
	StatusCakeBot                     *bool `json:"StatusCake bot,omitempty"`
	Stripe                            *bool `json:"Stripe,omitempty"`
	SucuriUptimeMonitorBot            *bool `json:"Sucuri Uptime Monitor Bot,omitempty"`
	TelegramBot                       *bool `json:"Telegram Bot,omitempty"`
	TestomatoBot                      *bool `json:"Testomato Bot,omitempty"`
	TheFindCrawler                    *bool `json:"TheFind Crawler,omitempty"`
	TwitterBot                        *bool `json:"Twitter bot,omitempty"`
	UptimeRobot                       *bool `json:"Uptime Robot,omitempty"`
	VkontakteExternalHitBot           *bool `json:"Vkontakte External Hit bot,omitempty"`
	W3C                               *bool `json:"W3C,omitempty"`
	WordfenceCentral                  *bool `json:"Wordfence Central,omitempty"`
	Workato                           *bool `json:"Workato,omitempty"`
	XMLSitemaps                       *bool `json:"xml-sitemaps,omitempty"`
	YahooInktomiSlurpBot              *bool `json:"Yahoo Inktomi Slurp bot,omitempty"`
	YahooJapanBot                     *bool `json:"Yahoo Japan bot,omitempty"`
	YahooLinkPreview                  *bool `json:"Yahoo Link Preview,omitempty"`
	YahooSeekerBot                    *bool `json:"Yahoo Seeker bot,omitempty"`
	YahooSlurpBot                     *bool `json:"Yahoo Slurp bot,omitempty"`
	YandexBot                         *bool `json:"Yandex bot,omitempty"`
	YisouSpiderCommercial             *bool `json:"YisouSpider - Commercial,omitempty"`
	YodaoBot                          *bool `json:"Yodao bot,omitempty"`
	ZendeskBot                        *bool `json:"Zendesk Bot,omitempty"`
	ZohoBot                           *bool `json:"Zoho bot,omitempty"`
	ZumBot                            *bool `json:"Zum Bot,omitempty"`
}

type WrappedWAFSettings struct {
	Data WAFSettings `json:"data"`
}

type WrappedWAFSettingsSet struct {
	Data []WAFSettings `json:"data"`
}

//Firewall Rule
type FirewallRule struct {
	Action  string `json:"action,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
	Id      string `json:"id,omitempty"`
	IpEnd   string `json:"ipEnd,omitempty"`
	IpStart string `json:"ipStart,omitempty"`
	Name    string `json:"name,omitempty"`
	SiteId  string `json:"siteId,omitempty"`
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

type VPCSubnets struct {
	Id      string `json:"id,omitempty"`
	StackId string `json:"stackId,omitempty"`
	VpcId   string `json:"vpcId,omitempty"`
	Name    string `json:"name,omitempty"`
	Slug    string `json:"slug,omitempty"`
	Cidr    string `json:"cidr,omitempty"`
	Status  string `json:"status,omitempty"`
}

type VPCRoutes struct {
	Id               string   `json:"id,omitempty"`
	StackId          string   `json:"stackId,omitempty"`
	VpcId            string   `json:"vpcId,omitempty"`
	Name             string   `json:"name,omitempty"`
	Slug             string   `json:"slug,omitempty"`
	DestinationCidrs []string `json:"destinationCidrs,omitempty"`
	NextHops         []string `json:"nextHops,omitempty"`
	Status           string   `json:"status,omitempty"`
}

type VPCs struct {
	Id         string       `json:"id,omitempty"`
	Name       string       `json:"name,omitempty"`
	StackId    string       `json:"stackId,omitempty"`
	Slug       string       `json:"slug,omitempty"`
	Cidr       string       `json:"cidr,omitempty"`
	DefaultVpc bool         `json:"defaultVpc,omitempty"`
	Created    string       `json:"created,omitempty"`
	Subnets    []VPCSubnets `json:"subnets,omitempty"`
	Routes     []VPCRoutes  `json:"routes,omitempty"`
	Status     string       `json:"status,omitempty"`
}

type WrappedVPCs struct {
	Data VPCs `json:"data,omitempty"`
}

type WrappedVPCsData struct {
	Data []VPCs `json:"data,omitempty"`
}

type Subnets struct {
	Id      string `json:"id,omitempty"`
	StackId string `json:"stackId,omitempty"`
	VpcId   string `json:"vpcId,omitempty"`
	VpcName string `json:"vpcName,omitempty"`
	VpcSlug string `json:"vpcSlug,omitempty"`
	Name    string `json:"name,omitempty"`
	Slug    string `json:"slug,omitempty"`
	Cidr    string `json:"cidr,omitempty"`
	Status  string `json:"status,omitempty"`
}

type WrappedSubnet struct {
	Data Subnets `json:"data,omitempty"`
}

type WrappedSubnetsData struct {
	Data []Subnets `json:"data,omitempty"`
}

type Route struct {
	Id               string   `json:"id,omitempty"`
	StackId          string   `json:"stackId,omitempty"`
	VpcId            string   `json:"vpcId,omitempty"`
	VpcName          string   `json:"vpcName,omitempty"`
	Name             string   `json:"name,omitempty"`
	Slug             string   `json:"slug,omitempty"`
	DestinationCidrs []string `json:"destinationCidrs,omitempty"`
	NextHops         []string `json:"nextHops,omitempty"`
	Status           string   `json:"status,omitempty"`
}

type WrappedRoute struct {
	Data Route `json:"data,omitempty"`
}

type WrappedRoutesData struct {
	Data []Route `json:"data,omitempty"`
}

//BareMetal

type BareMetalDevice struct {
	Id                       string                `json:"id,omitempty"`
	ServicePlan              string                `json:"servicePlan,omitempty"`
	Name                     string                `json:"name,omitempty"`
	Hostname                 string                `json:"hostname,omitempty"`
	DeviceType               string                `json:"deviceType,omitempty"`
	PrimaryIp                string                `json:"primaryIp,omitempty"`
	Status                   string                `json:"status,omitempty"`
	MonitorsTotal            int                   `json:"monitorsTotal,omitempty"`
	MonitorsUp               int                   `json:"monitorsUp,omitempty"`
	IpmiAddress              string                `json:"ipmiAddress,omitempty"`
	PowerStatus              string                `json:"powerStatus,omitempty"`
	Tags                     []string              `json:"tags,omitempty"`
	Location                 Location              `json:"location,omitempty"`
	DeviceDetail             DeviceDetail          `json:"deviceDetail,omitempty"`
	DeviceInitialPassword    DeviceInitialPassword `json:"deviceInitialPassword,omitempty"`
	DeviceIPs                DeviceIPs             `json:"deviceIPs,omitempty"`
	IsNetworkPolicyAvailable bool                  `json:"isNetworkPolicyAvailable,omitempty"`
	ChangeId                 string                `json:"changeId,omitempty"`
}

type Location struct {
	Facility      string `json:"facility,omitempty"`
	FacilityTitle string `json:"facility_title,omitempty"`
}

type WrappedBareMetalDevices struct {
	Data []BareMetalDevice `json:"data"`
}

type WrappedBareMetalDevice struct {
	Data BareMetalDevice `json:"data"`
}

type BareMetalDeviceChart struct {
	Id         string `json:"id,omitempty"`
	Filter     string `json:"filter,omitempty"`
	GraphImage string `json:"graphImage,omitempty"`
	Interfaces string `json:"interfaces,omitempty"`
	SwitchId   string `json:"switchId,omitempty"`
}

type WrappedBareMetalDeviceCharts struct {
	Data []BareMetalDeviceChart `json:"data,omitempty"`
}

type BareMetalDeviceSensor struct {
	Id        string `json:"id,omitempty"`
	IpmiField string `json:"ipmiField,omitempty"`
	IpmiValue string `json:"ipmiValue,omitempty"`
}

type WrappedBareMetalDeviceSensors struct {
	Data []BareMetalDeviceSensor `json:"data,omitempty"`
}

type BareMetalDeviceIP struct {
	IPName string `json:"ipName,omitempty"`
	Value  string `json:"value,omitempty"`
}

type WrappedBareMetalDeviceIPs struct {
	Data []BareMetalDeviceIP `json:"data,omitempty"`
}

type BareMetalSSHKey struct {
	Id        string `json:"id,omitempty"`
	PublicKey string `json:"publicKey,omitempty"`
	Name      string `json:"name,omitempty"`
}

type WrappedBareMetalSSHKeys struct {
	Data []BareMetalSSHKey `json:"data,omitempty"`
}
type WrappedBareMetalSSHKey struct {
	Data BareMetalSSHKey `json:"data,omitempty"`
}

type DeviceDetail struct {
	ProductID        string         `json:"productId,omitempty"`
	ServicePlan      string         `json:"servicePlan,omitempty"`
	Processor        string         `json:"processor,omitempty"`
	PrimaryHardDrive string         `json:"primaryHardDrive,omitempty"`
	Memory           string         `json:"memory,omitempty"`
	OperatingSystem  string         `json:"operatingSystem,omitempty"`
	Bandwidth        string         `json:"bandwidth,omitempty"`
	InternalNetwork  string         `json:"internalNetwork,omitempty"`
	DDoS             string         `json:"ddos,omitempty"`
	RaidSetUp        string         `json:"raidSetUp,omitempty"`
	NextRenew        string         `json:"nextRenew,omitempty"`
	DeviceIPDetail   DeviceIPDetail `json:"deviceIPDetail,omitempty"`
}

type DeviceIPDetail struct {
	PrimaryIP   string   `json:"primaryIp,omitempty"`
	Description string   `json:"description,omitempty"`
	GatewayIP   string   `json:"gatewayIp,omitempty"`
	SubnetMask  string   `json:"subnetMask,omitempty"`
	UsableIPs   []string `json:"usableIps,omitempty"`
}

type DeviceInitialPassword struct {
	PasswordReturnsUntil int    `json:"passwordReturnsUntil,omitempty"`
	PasswordExpires      string `json:"passwordExpires,omitempty"`
	Port                 int    `json:"port,omitempty"`
	User                 string `json:"user,omitempty"`
}

type DeviceIPs struct {
	Subnet    string   `json:"subnet,omitempty"`
	Netmask   string   `json:"netmask,omitempty"`
	UsableIPs []string `json:"usableIps,omitempty"`
}

type BareMetalDeviceDisk struct {
	ServerDiskID           int    `json:"server_disk_id,omitempty"`
	ServerDiskModel        string `json:"server_disk_model,omitempty"`
	ServerDiskSizeGB       int    `json:"server_disk_size_gb,omitempty"`
	ServerID               int    `json:"server_id,omitempty"`
	ServerDiskSerial       string `json:"server_disk_serial,omitempty"`
	ServerDiskVendor       string `json:"server_disk_vendor,omitempty"`
	ServerDiskStatus       string `json:"server_disk_status,omitempty"`
	ServerDiskType         string `json:"server_disk_type,omitempty"`
	ServerRaidControllerID int    `json:"server_raid_controller_id,omitempty"`
	Type                   string `json:"type,omitempty"`
}

type WrappedBareMetalDeviceDisks struct {
	Data []BareMetalDeviceDisk `json:"data,omitempty"`
}

type BareMetalLocation struct {
	ID         string `json:"id"`
	LocationID string `json:"locationId"`
	Code       string `json:"code"`
	Name       string `json:"name"`
}

type WrappedBareMetalLocations struct {
	Data []BareMetalLocation `json:"data,omitempty"`
}

type ProductProcessorInfo struct {
	Cores   int `json:"cores"`
	Sockets int `json:"sockets"`
	Threads int `json:"threads"`
	VCPUs   int `json:"vcpus"`
}

type BareMetalLocationProduct struct {
	ID              string               `json:"id"`
	Drive           string               `json:"drive"`
	CPU             string               `json:"cpu"`
	SubTitle        string               `json:"subTitle"`
	Memory          string               `json:"memory"`
	Bandwidth       string               `json:"bandwidth"`
	MonthlyPrice    string               `json:"monthlyPrice"`
	MonthlyPremium  string               `json:"monthlyPremium"`
	Stock           string               `json:"stock"`
	ProcessorInfo   ProductProcessorInfo `json:"processorInfo"`
	CPUCores        string               `json:"cpuCores"`
	GPU             string               `json:"gpu"`
	HourlyPrice     string               `json:"hourlyPrice"`
	HourlyPremium   string               `json:"hourlyPremium"`
	VendorProductID string               `json:"vendorProductId"`
}

type WrappedBareMetalLocationProduct struct {
	Data []BareMetalLocationProduct `json:"data,omitempty"`
}

type BareMetalLocationProductOS struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type WrappedBareMetalLocationProductOS struct {
	Data []BareMetalLocationProductOS `json:"data,omitempty"`
}

//Edge Compute 2.0

type ComputeWorkload struct {
	Id               string   `json:"id"`
	Hostname         string   `json:"hostname"`
	Label            string   `json:"label"`
	Status           string   `json:"status"`
	OS               string   `json:"os"`
	RAM              string   `json:"ram"`
	DateCreated      string   `json:"date_created"`
	Region           string   `json:"region"`
	Disk             string   `json:"disk"`
	MainIP           string   `json:"main_ip"`
	VCPUCount        int      `json:"vcpu_count"`
	Plan             string   `json:"plan"`
	AllowedBandwidth int      `json:"allowed_bandwidth"`
	NetmaskV4        string   `json:"netmask_v4"`
	GatewayV4        string   `json:"gateway_v4"`
	PowerStatus      string   `json:"power_status"`
	ServerStatus     string   `json:"server_status"`
	V6Network        string   `json:"v6_network"`
	V6MainIP         string   `json:"v6_main_ip"`
	V6NetworkSize    int      `json:"v6_network_size"`
	InternalIP       string   `json:"internal_ip"`
	KVM              string   `json:"kvm"`
	OSID             int      `json:"os_id"`
	AppID            int      `json:"app_id"`
	ImageID          string   `json:"image_id"`
	FirewallGroupID  string   `json:"firewall_group_id"`
	Features         []string `json:"features"`
	Tags             []string `json:"tags"`
}

type WrapperComputeWorkloads struct {
	Data []ComputeWorkload
}

type WrapperComputeWorkload struct {
	Data ComputeWorkload
}

type IPv4Configuration struct {
	IP      string `json:"ip"`
	Netmask string `json:"netmask"`
	Gateway string `json:"gateway"`
	Type    string `json:"type"`
	Reverse string `json:"reverse"`
}

type WrapperIPv4Configuration struct {
	Data []IPv4Configuration
}

type IPv6Configuration struct {
	Id          string `json:"id"`
	IP          string `json:"ip"`
	Network     string `json:"network"`
	NetworkSize int    `json:"network_size"`
	Type        string `json:"type"`
}

type WrapperIPv6Configuration struct {
	Data []IPv6Configuration
}

type IPv6ReverseDNSConfiguration struct {
	Id      string `json:"id"`
	Ip      string `json:"ip"`
	Reverse string `json:"reverse"`
}

type WrapperIPv6ReverseDNSConfiguration struct {
	Data []IPv6ReverseDNSConfiguration
}

type FirewallGroup struct {
	Id         string `json:"id"`
	FirewallId string `json:"firewallId"`
	Name       string `json:"name"`
}

type WrapperFirewallGroup struct {
	Data FirewallGroup
}

type Hostname struct {
	Id       string `json:"id"`
	Hostname string `json:"hostname"`
}

type WrapperHostname struct {
	Data Hostname
}

type WorkloadPlan struct {
	ID        string `json:"id"`
	PlanID    string `json:"planId"`
	Region    string `json:"region"`
	Server    string `json:"server"`
	PlanLabel string `json:"planLabel"`
	VCPUCount int    `json:"vCPUCount"`
}

type WrapperWorkloadPlan struct {
	Data WorkloadPlan
}

type ComputeWorkloadOS struct {
	ID      string `json:"id"`
	OsID    int    `json:"osId"`
	OsLabel string `json:"osLabel"`
	PlanId  string `json:"planId"`
}

type WrapperComputeWorkloadOS struct {
	Data ComputeWorkloadOS
}

type ComputeWorkloadUserData struct {
	ID       string `json:"id"`
	UserData string `json:"userdata"`
}

type WrapperComputeWorkloadUserData struct {
	Data ComputeWorkloadUserData
}

type ComputeWorkloadTag struct {
	ID  string `json:"id"`
	Tag string `json:"tag"`
}

type WrapperComputeWorkloadTags struct {
	Data []ComputeWorkloadTag
}

type WrapperComputeWorkloadTag struct {
	Data ComputeWorkloadTag
}

type ComputeStorage struct {
	ID                    string `json:"id"`
	DateCreated           string `json:"date_created"`
	Cost                  string `json:"cost"`
	Status                string `json:"status"`
	SizeGB                string `json:"size_gb"`
	Region                string `json:"region"`
	AttachedToInstance    string `json:"attached_to_instance"`
	Label                 string `json:"label"`
	MountID               string `json:"mount_id"`
	BlockType             string `json:"block_type"`
	Description           string `json:"description"`
	Type                  string `json:"type"`
	Location              string `json:"location"`
	AttachedTo            string `json:"attached_to"`
	ManageLabel           string `json:"manage_label"`
	Price                 string `json:"price"`
	SizeInGB              string `json:"size_in_gb"`
	EditBlockStorageLabel string `json:"edit_block_storage_label"`
	None                  string `json:"none"`
	Detach                string `json:"detach"`
	Attach                string `json:"attach"`
}

type WrapperComputeStorage struct {
	Data ComputeStorage
}

type WrapperComputeStorages struct {
	Data []ComputeStorage
}

type ComputeFirewall struct {
	ID            string `json:"id"`
	Description   string `json:"description"`
	DateCreated   string `json:"date_created"`
	DateModified  string `json:"date_modified"`
	InstanceCount int    `json:"instance_count"`
	RuleCount     int    `json:"rule_count"`
	MaxRuleCount  int    `json:"max_rule_count"`
}

type WrapperComputeFirewall struct {
	Data ComputeFirewall
}

type WrapperComputeFirewalls struct {
	Data []ComputeFirewall
}

type ComputeFirewallRule struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	IPType     string `json:"ip_type"`
	Action     string `json:"action"`
	Protocol   string `json:"protocol"`
	Port       string `json:"port"`
	Subnet     string `json:"subnet"`
	SubnetSize string `json:"subnet_size"`
	Source     string `json:"source"`
	Notes      string `json:"notes"`
}

type WrapperComputeFirewallRule struct {
	Data ComputeFirewallRule
}

type WrapperComputeFirewallRules struct {
	Data []ComputeFirewallRule
}

type ComputeFirewallLinkedInstance struct {
	ID               string   `json:"id"`
	Hostname         string   `json:"hostname"`
	Label            string   `json:"label"`
	Status           string   `json:"status"`
	OS               string   `json:"os"`
	RAM              string   `json:"ram"`
	DateCreated      string   `json:"date_created"`
	Region           string   `json:"region"`
	Disk             string   `json:"disk"`
	MainIP           string   `json:"main_ip"`
	VCPUCount        int      `json:"vcpu_count"`
	Plan             string   `json:"plan"`
	AllowedBandwidth int      `json:"allowed_bandwidth"`
	NetmaskV4        string   `json:"netmask_v4"`
	GatewayV4        string   `json:"gateway_v4"`
	PowerStatus      string   `json:"power_status"`
	ServerStatus     string   `json:"server_status"`
	V6Network        string   `json:"v6_network"`
	V6MainIP         string   `json:"v6_main_ip"`
	V6NetworkSize    int      `json:"v6_network_size"`
	InternalIP       string   `json:"internal_ip"`
	KVM              string   `json:"kvm"`
	OSID             int      `json:"os_id"`
	AppID            int      `json:"app_id"`
	ImageID          string   `json:"image_id"`
	FirewallGroupID  string   `json:"firewall_group_id"`
	Features         []string `json:"features"`
	Tags             []string `json:"tags"`
}

type WrapperComputeFirewallLinkedInstance struct {
	Data ComputeFirewallLinkedInstance
}

type WrapperComputeFirewallLinkedInstances struct {
	Data []ComputeFirewallLinkedInstance
}

type ComputeVPC2 struct {
	ID           string `json:"id"`
	DateCreated  string `json:"date_created"`
	Region       string `json:"region"`
	Location     string `json:"location"`
	Description  string `json:"description"`
	IPBlock      string `json:"ip_block"`
	PrefixLength int    `json:"prefix_length"`
	Subnet       string `json:"subnet"`
}

type WrapperComputeVPC2 struct {
	Data ComputeVPC2
}

type WrapperComputeVPC2s struct {
	Data []ComputeVPC2
}

type ComputeVPC struct {
	ID           string `json:"id"`
	DateCreated  string `json:"date_created"`
	Region       string `json:"region"`
	Description  string `json:"description"`
	V4Subnet     string `json:"v4_subnet"`
	V4SubnetMask string `json:"v4_subnet_mask"`
	Subnet       string `json:"subnet"`
	Location     string `json:"location"`
}

type WrapperComputeVPC struct {
	Data ComputeVPC
}

type WrapperComputeVPCs struct {
	Data []ComputeVPC
}

type ComputeReservedIP struct {
	ID                 string `json:"id"`
	Region             string `json:"region"`
	Location           string `json:"location"`
	IPType             string `json:"ip_type"`
	Subnet             string `json:"subnet"`
	SubnetSize         int    `json:"subnet_size"`
	Label              string `json:"label"`
	InstanceID         string `json:"instance_id"`
	ReservedIP         string `json:"reservedIp"`
	IsWorkloadAttached bool   `json:"isWorkloadAttached"`
}

type WrapperComputeReservedIP struct {
	Data ComputeReservedIP
}

type WrapperComputeReservedIPs struct {
	Data []ComputeReservedIP
}
