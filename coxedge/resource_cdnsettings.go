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
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"strings"
)

func resourceCDNSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCDNSettingsCreate,
		ReadContext:   resourceCDNSettingsRead,
		UpdateContext: resourceCDNSettingsUpdate,
		DeleteContext: resourceCDNSettingsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: getCDNSettingsSchema(),
	}
}

func resourceCDNSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Convert to struct
	updatedCDNSettings := convertResourceDataToCDNSettingsCreateAPIObject(d)
	d.SetId(updatedCDNSettings.Id)
	resourceCDNSettingsUpdate(ctx, d, m)

	return diags
}

func resourceCDNSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		d.Set("organization_id", keys[2])
	}

	//Get the resource ID
	resourceId := d.Id()
	organizationId := d.Get("organization_id").(string)
	//Get the resource
	cdnSettings, err := coxEdgeClient.GetCDNSettings(d.Get("environment_name").(string), resourceId, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	convertCDNSettingsAPIObjectToResourceData(d, cdnSettings)

	return diags
}

func resourceCDNSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Convert resource data to API object
	updatedCDNSettings := convertResourceDataToCDNSettingsCreateAPIObject(d)
	organizationId := d.Get("organization_id").(string)
	//Call the API
	taskResp, err := coxEdgeClient.UpdateCDNSettings(updatedCDNSettings.Id, updatedCDNSettings, organizationId)
	if err != nil {
		return diag.FromErr(err)
	}

	//Await
	_, err = coxEdgeClient.AwaitTaskResolveWithDefaults(ctx, taskResp.TaskId)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceCDNSettingsRead(ctx, d, m)
}

func resourceCDNSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func convertResourceDataToCDNSettingsCreateAPIObject(d *schema.ResourceData) apiclient.CDNSettings {
	//Create update cdnSettings struct
	updatedCDNSettings := apiclient.CDNSettings{
		EnvironmentName: d.Get("environment_name").(string),
		Id:              d.Get("site_id").(string),
		SiteId:          d.Get("site_id").(string),
		BrowserCacheTtl: d.Get("browser_cache_ttl").(int),
	}

	queryStringControlPre, queryStringControlCurrent := d.GetChange("query_string_control")
	if queryStringControlCurrent != "" {
		updatedCDNSettings.QueryStringControl = queryStringControlCurrent.(string)
		if updatedCDNSettings.QueryStringControl == "CUSTOM" {
			for _, val := range d.Get("custom_cached_query_strings").([]interface{}) {
				updatedCDNSettings.CustomCachedQueryStrings = append(updatedCDNSettings.CustomCachedQueryStrings, val.(string))
			}
		}
	} else {
		if queryStringControlPre.(string) == "CUSTOM" {
			for _, val := range d.Get("custom_cached_query_strings").([]interface{}) {
				updatedCDNSettings.CustomCachedQueryStrings = append(updatedCDNSettings.CustomCachedQueryStrings, val.(string))
			}
		}
	}

	cacheExpirePolicyPre, cacheExpirePolicyCurrent := d.GetChange("cache_expire_policy")
	if cacheExpirePolicyCurrent.(string) != "" {
		updatedCDNSettings.CacheExpirePolicy = cacheExpirePolicyCurrent.(string)
		if updatedCDNSettings.CacheExpirePolicy == "SPECIFY_CDN_TTL" {
			updatedCDNSettings.CacheTtl = d.Get("cache_ttl").(int)
		}
	} else {
		if cacheExpirePolicyPre.(string) == "SPECIFY_CDN_TTL" {
			updatedCDNSettings.CacheTtl = d.Get("cache_ttl").(int)
		}
	}

	urlCachingEnabledValue := false
	urlCachingEnabledValuePre, urlCachingEnabledValueCurrent := d.GetChange("url_caching_enabled")
	if urlCachingEnabledValueCurrent.(string) != "" {
		urlCachingEnabledValue, _ = strconv.ParseBool(urlCachingEnabledValueCurrent.(string))
		updatedCDNSettings.UrlCachingEnabled = utils.BoolAddr(urlCachingEnabledValue)
	} else {
		urlCachingEnabledValue, _ = strconv.ParseBool(urlCachingEnabledValuePre.(string))
	}
	if urlCachingEnabledValue {
		updatedCDNSettings.UrlCachingTtl = d.Get("url_caching_ttl").(int)
	}

	canonicalHeaderEnabledValue := false
	canonicalHeaderEnabledValuePre, canonicalHeaderEnabledValueCurrent := d.GetChange("canonical_header_enabled")
	if canonicalHeaderEnabledValueCurrent.(string) != "" {
		canonicalHeaderEnabledValue, _ = strconv.ParseBool(canonicalHeaderEnabledValueCurrent.(string))
		updatedCDNSettings.CanonicalHeaderEnabled = utils.BoolAddr(canonicalHeaderEnabledValue)
	} else {
		canonicalHeaderEnabledValue, _ = strconv.ParseBool(canonicalHeaderEnabledValuePre.(string))
	}
	if canonicalHeaderEnabledValue {
		updatedCDNSettings.CanonicalHeader = d.Get("canonical_header").(string)
	}

	corsHeaderEnabledValue := false
	corsHeaderEnabledValuePre, corsHeaderEnabledValueCurrent := d.GetChange("cors_header_enabled")
	if corsHeaderEnabledValueCurrent.(string) != "" {
		corsHeaderEnabledValue, _ = strconv.ParseBool(corsHeaderEnabledValueCurrent.(string))
		updatedCDNSettings.CorsHeaderEnabled = utils.BoolAddr(corsHeaderEnabledValue)
	} else {
		corsHeaderEnabledValue, _ = strconv.ParseBool(corsHeaderEnabledValuePre.(string))
	}
	if corsHeaderEnabledValue {
		updatedCDNSettings.AllowedCorsOrigins = d.Get("allowed_cors_origins").(string)
		if updatedCDNSettings.AllowedCorsOrigins == "SPECIFY_ORIGINS" {
			for _, val := range d.Get("origins_to_allow_cors").([]interface{}) {
				updatedCDNSettings.OriginsToAllowCors = append(updatedCDNSettings.OriginsToAllowCors, val.(string))
			}
		}
	}

	varyHeaderEnabled := d.Get("vary_header_enabled").(string)
	if varyHeaderEnabled != "" {
		boolValue, _ := strconv.ParseBool(varyHeaderEnabled)
		updatedCDNSettings.VaryHeaderEnabled = utils.BoolAddr(boolValue)
	}

	contentPersistenceEnabledValue := false
	contentPersistenceEnabledValuePre, contentPersistenceEnabledValueCurrent := d.GetChange("content_persistence_enabled")
	if contentPersistenceEnabledValueCurrent.(string) != "" {
		contentPersistenceEnabledValue, _ = strconv.ParseBool(contentPersistenceEnabledValueCurrent.(string))
		updatedCDNSettings.ContentPersistenceEnabled = utils.BoolAddr(contentPersistenceEnabledValue)
	} else {
		contentPersistenceEnabledValue, _ = strconv.ParseBool(contentPersistenceEnabledValuePre.(string))
	}
	if contentPersistenceEnabledValue {
		updatedCDNSettings.MaximumStaleFileTtl = d.Get("maximum_stale_file_ttl").(int)
	}

	dynamicCachingByHeaderEnabledValue := false
	dynamicCachingByHeaderEnabledValuePrev, dynamicCachingByHeaderEnabledValueCurrent := d.GetChange("dynamic_caching_by_header_enabled")
	if dynamicCachingByHeaderEnabledValueCurrent.(string) != "" {
		dynamicCachingByHeaderEnabledValue, _ = strconv.ParseBool(dynamicCachingByHeaderEnabledValueCurrent.(string))
		updatedCDNSettings.DynamicCachingByHeaderEnabled = utils.BoolAddr(dynamicCachingByHeaderEnabledValue)
	} else {
		dynamicCachingByHeaderEnabledValue, _ = strconv.ParseBool(dynamicCachingByHeaderEnabledValuePrev.(string))
	}
	if dynamicCachingByHeaderEnabledValue {
		for _, val := range d.Get("custom_cached_headers").([]interface{}) {
			updatedCDNSettings.CustomCacheHeaders = append(updatedCDNSettings.CustomCacheHeaders, val.(string))
		}
	}

	gzipCompressionEnabledValue := false
	gzipCompressionEnabledValuePrev, gzipCompressionEnabledValueCurrent := d.GetChange("gzip_compression_enabled")
	if gzipCompressionEnabledValueCurrent.(string) != "" {
		gzipCompressionEnabledValue, _ = strconv.ParseBool(gzipCompressionEnabledValueCurrent.(string))
		updatedCDNSettings.GzipCompressionEnabled = utils.BoolAddr(gzipCompressionEnabledValue)
	} else {
		gzipCompressionEnabledValue, _ = strconv.ParseBool(gzipCompressionEnabledValuePrev.(string))
	}
	if gzipCompressionEnabledValue {
		updatedCDNSettings.GzipCompressionLevel = d.Get("gzip_compression_level").(int)
	}

	http2SupportEnabledValue := false
	http2SupportEnabledValuePrevious, http2SupportEnabledValueCurrent := d.GetChange("http2_support_enabled")
	if http2SupportEnabledValueCurrent.(string) != "" {
		http2SupportEnabledValue, _ = strconv.ParseBool(http2SupportEnabledValueCurrent.(string))
		updatedCDNSettings.Http2SupportEnabled = utils.BoolAddr(http2SupportEnabledValue)
	} else {
		http2SupportEnabledValue, _ = strconv.ParseBool(http2SupportEnabledValuePrevious.(string))
	}
	http2ServerPushEnabledValue := false
	if http2SupportEnabledValue {
		http2ServerPushEnabledPrev, http2ServerPushEnabledCurrent := d.GetChange("http2_server_push_enabled")
		if http2ServerPushEnabledCurrent.(string) != "" {
			http2ServerPushEnabledValue, _ = strconv.ParseBool(http2ServerPushEnabledCurrent.(string))
			updatedCDNSettings.Http2ServerPushEnabled = utils.BoolAddr(http2ServerPushEnabledValue)
		} else {
			http2ServerPushEnabledValue, _ = strconv.ParseBool(http2ServerPushEnabledPrev.(string))
		}
	}
	if http2ServerPushEnabledValue {
		updatedCDNSettings.LinkHeader = d.Get("link_header").(string)
	}

	return updatedCDNSettings
}

func convertCDNSettingsAPIObjectToResourceData(d *schema.ResourceData, cdnSettings *apiclient.CDNSettings) {
	//Store the data
	d.Set("site_id", cdnSettings.Id)
	d.Set("cache_expire_policy", cdnSettings.CacheExpirePolicy)
	d.Set("cache_ttl", cdnSettings.CacheTtl)
	d.Set("query_string_control", cdnSettings.QueryStringControl)
	d.Set("custom_cached_query_strings", cdnSettings.CustomCachedQueryStrings)
	if cdnSettings.DynamicCachingByHeaderEnabled != nil {
		d.Set("dynamic_caching_by_header_enabled", strconv.FormatBool(*cdnSettings.DynamicCachingByHeaderEnabled))
	}
	d.Set("custom_cached_headers", cdnSettings.CustomCacheHeaders)
	d.Set("gzip_compression_enabled", cdnSettings.GzipCompressionEnabled)
	d.Set("gzip_compression_level", cdnSettings.GzipCompressionLevel)
	if cdnSettings.ContentPersistenceEnabled != nil {
		d.Set("content_persistence_enabled", strconv.FormatBool(*cdnSettings.ContentPersistenceEnabled))
	}
	d.Set("maximum_stale_file_ttl", cdnSettings.MaximumStaleFileTtl)
	if cdnSettings.VaryHeaderEnabled != nil {
		d.Set("vary_header_enabled", strconv.FormatBool(*cdnSettings.VaryHeaderEnabled))
	}
	d.Set("browser_cache_ttl", cdnSettings.BrowserCacheTtl)
	if cdnSettings.CorsHeaderEnabled != nil {
		d.Set("cors_header_enabled", strconv.FormatBool(*cdnSettings.CorsHeaderEnabled))
	}
	d.Set("allowed_cors_origins", cdnSettings.AllowedCorsOrigins)
	d.Set("origins_to_allow_cors", cdnSettings.OriginsToAllowCors)
	if cdnSettings.Http2SupportEnabled != nil {
		d.Set("http2_support_enabled", strconv.FormatBool(*cdnSettings.Http2SupportEnabled))
	}
	if cdnSettings.Http2ServerPushEnabled != nil {
		d.Set("http2_server_push_enabled", strconv.FormatBool(*cdnSettings.Http2ServerPushEnabled))
	}
	d.Set("link_header", cdnSettings.LinkHeader)
	if cdnSettings.CanonicalHeaderEnabled != nil {
		d.Set("canonical_header_enabled", strconv.FormatBool(*cdnSettings.CanonicalHeaderEnabled))
	}
	d.Set("canonical_header", cdnSettings.CanonicalHeader)
	if cdnSettings.UrlCachingEnabled != nil {
		d.Set("url_caching_enabled", strconv.FormatBool(*cdnSettings.UrlCachingEnabled))
	}
	d.Set("url_caching_ttl", cdnSettings.UrlCachingTtl)
}
