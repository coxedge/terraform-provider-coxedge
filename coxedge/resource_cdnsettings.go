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
	d.SetId(updatedCDNSettings.SiteId)
	resourceCDNSettingsUpdate(ctx, d, m)

	return diags
}

func resourceCDNSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//Get the resource ID
	resourceId := d.Id()

	//Get the resource
	cdnSettings, err := coxEdgeClient.GetCDNSettings(d.Get("environment_name").(string), resourceId)
	if err != nil {
		return diag.FromErr(err)
	}

	convertCDNSettingsAPIObjectToResourceData(d, cdnSettings)

	//Update state
	resourceCDNSettingsRead(ctx, d, m)

	return diags
}

func resourceCDNSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//Get the API Client
	coxEdgeClient := m.(apiclient.Client)

	//Get the resource ID
	resourceId := d.Id()

	//Convert resource data to API object
	updatedCDNSettings := convertResourceDataToCDNSettingsCreateAPIObject(d)

	//Call the API
	_, err := coxEdgeClient.UpdateCDNSettings(resourceId, updatedCDNSettings)
	if err != nil {
		return diag.FromErr(err)
	}

	//Set last_updated
	d.Set("last_updated", time.Now().Format(time.RFC850))

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
		EnvironmentName:               d.Get("environment_name").(string),
		SiteId:                        d.Get("site_id").(string),
		CacheExpirePolicy:             d.Get("cache_expire_policy").(string),
		CacheTtl:                      d.Get("cache_ttl").(string),
		QueryStringControl:            d.Get("query_string_control").(string),
		CustomCachedQueryStrings:      d.Get("custom_cached_query_strings").([]string),
		DynamicCachingByHeaderEnabled: d.Get("dynamic_caching_by_header_enabled").(bool),
		CustomCacheHeaders:            d.Get("custom_cached_headers").([]string),
		GzipCompressionEnabled:        d.Get("gzip_compression_enabled").(bool),
		GzipCompressionLevel:          d.Get("gzip_compression_level").(int),
		ContentPersistenceEnabled:     d.Get("content_persistence_enabled").(bool),
		MaximumStaleFileTtl:           d.Get("maximum_stale_file_ttl").(int),
		VaryHeaderEnabled:             d.Get("vary_header_enabled").(bool),
		BrowserCacheTtl:               d.Get("browser_cache_ttl").(int),
		CorsHeaderEnabled:             d.Get("cors_header_enabled").(bool),
		AllowedCorsOrigins:            d.Get("allowed_cors_origins").(string),
		OriginsToAllowCors:            d.Get("origins_to_allow_cors").([]string),
		Http2SupportEnabled:           d.Get("http2_support_enabled").(bool),
		LinkHeader:                    d.Get("link_header").(string),
		CanonicalHeaderEnabled:        d.Get("canonical_header_enabled").(bool),
		CanonicalHeader:               d.Get("canonical_header").(string),
		UrlCachingEnabled:             d.Get("url_caching_enabled").(bool),
		UrlCachingTtl:                 d.Get("url_caching_ttl").(int),
	}

	return updatedCDNSettings
}

func convertCDNSettingsAPIObjectToResourceData(d *schema.ResourceData, cdnSettings *apiclient.CDNSettings) {
	//Store the data
	d.Set("environment_name", cdnSettings.EnvironmentName)
	d.Set("site_id", cdnSettings.SiteId)
	d.Set("cache_expire_policy", cdnSettings.CacheExpirePolicy)
	d.Set("cache_ttl", cdnSettings.CacheTtl)
	d.Set("query_string_control", cdnSettings.QueryStringControl)
	d.Set("custom_cached_query_strings", cdnSettings.CustomCachedQueryStrings)
	d.Set("dynamic_caching_by_header_enabled", cdnSettings.DynamicCachingByHeaderEnabled)
	d.Set("custom_cached_headers", cdnSettings.CustomCacheHeaders)
	d.Set("gzip_compression_enabled", cdnSettings.GzipCompressionEnabled)
	d.Set("gzip_compression_level", cdnSettings.GzipCompressionLevel)
	d.Set("content_persistence_enabled", cdnSettings.ContentPersistenceEnabled)
	d.Set("maximum_stale_file_ttl", cdnSettings.MaximumStaleFileTtl)
	d.Set("vary_header_enabled", cdnSettings.VaryHeaderEnabled)
	d.Set("browser_cache_ttl", cdnSettings.BrowserCacheTtl)
	d.Set("cors_header_enabled", cdnSettings.CorsHeaderEnabled)
	d.Set("allowed_cors_origins", cdnSettings.AllowedCorsOrigins)
	d.Set("origins_to_allow_cors", cdnSettings.OriginsToAllowCors)
	d.Set("http2_support_enabled", cdnSettings.Http2SupportEnabled)
	d.Set("link_header", cdnSettings.LinkHeader)
	d.Set("canonical_header_enabled", cdnSettings.CanonicalHeaderEnabled)
	d.Set("canonical_header", cdnSettings.CanonicalHeader)
	d.Set("url_caching_enabled", cdnSettings.UrlCachingEnabled)
	d.Set("url_caching_ttl", cdnSettings.UrlCachingTtl)
}
