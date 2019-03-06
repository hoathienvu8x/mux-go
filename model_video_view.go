/*
 * Mux Video
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package muxvideosdk

type VideoView struct {
	ViewTotalUpscaling string `json:"view_total_upscaling,omitempty"`
	PrerollAdAssetHostname string `json:"preroll_ad_asset_hostname,omitempty"`
	PlayerSourceDomain string `json:"player_source_domain,omitempty"`
	Region string `json:"region,omitempty"`
	ViewerUserAgent string `json:"viewer_user_agent,omitempty"`
	PrerollRequested bool `json:"preroll_requested,omitempty"`
	PageType string `json:"page_type,omitempty"`
	StartupScore string `json:"startup_score,omitempty"`
	ViewSeekDuration string `json:"view_seek_duration,omitempty"`
	CountryName string `json:"country_name,omitempty"`
	PlayerSourceHeight int32 `json:"player_source_height,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	BufferingCount string `json:"buffering_count,omitempty"`
	VideoDuration string `json:"video_duration,omitempty"`
	PlayerSourceType string `json:"player_source_type,omitempty"`
	City string `json:"city,omitempty"`
	ViewId string `json:"view_id,omitempty"`
	PlatformDescription string `json:"platform_description,omitempty"`
	VideoStartupPrerollRequestTime string `json:"video_startup_preroll_request_time,omitempty"`
	ViewerDeviceName string `json:"viewer_device_name,omitempty"`
	VideoSeries string `json:"video_series,omitempty"`
	ViewerApplicationName string `json:"viewer_application_name,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	ViewTotalContentPlaybackTime string `json:"view_total_content_playback_time,omitempty"`
	Cdn string `json:"cdn,omitempty"`
	PlayerInstanceId string `json:"player_instance_id,omitempty"`
	VideoLanguage string `json:"video_language,omitempty"`
	PlayerSourceWidth int32 `json:"player_source_width,omitempty"`
	PlayerErrorMessage string `json:"player_error_message,omitempty"`
	PlayerMuxPluginVersion string `json:"player_mux_plugin_version,omitempty"`
	Watched bool `json:"watched,omitempty"`
	PlaybackScore string `json:"playback_score,omitempty"`
	PageUrl string `json:"page_url,omitempty"`
	Metro string `json:"metro,omitempty"`
	ViewMaxRequestLatency string `json:"view_max_request_latency,omitempty"`
	RequestsForFirstPreroll string `json:"requests_for_first_preroll,omitempty"`
	ViewTotalDownscaling string `json:"view_total_downscaling,omitempty"`
	Latitude string `json:"latitude,omitempty"`
	PlayerSourceHostName string `json:"player_source_host_name,omitempty"`
	InsertedAt string `json:"inserted_at,omitempty"`
	ViewEnd string `json:"view_end,omitempty"`
	MuxEmbedVersion string `json:"mux_embed_version,omitempty"`
	PlayerLanguage string `json:"player_language,omitempty"`
	PageLoadTime int32 `json:"page_load_time,omitempty"`
	ViewerDeviceCategory string `json:"viewer_device_category,omitempty"`
	VideoStartupPrerollLoadTime string `json:"video_startup_preroll_load_time,omitempty"`
	PlayerVersion string `json:"player_version,omitempty"`
	WatchTime int32 `json:"watch_time,omitempty"`
	PlayerSourceStreamType string `json:"player_source_stream_type,omitempty"`
	PrerollAdTagHostname string `json:"preroll_ad_tag_hostname,omitempty"`
	ViewerDeviceManufacturer string `json:"viewer_device_manufacturer,omitempty"`
	RebufferingScore string `json:"rebuffering_score,omitempty"`
	ExperimentName string `json:"experiment_name,omitempty"`
	ViewerOsVersion string `json:"viewer_os_version,omitempty"`
	PlayerPreload bool `json:"player_preload,omitempty"`
	BufferingDuration string `json:"buffering_duration,omitempty"`
	PlayerViewCount int32 `json:"player_view_count,omitempty"`
	PlayerSoftware string `json:"player_software,omitempty"`
	PlayerLoadTime string `json:"player_load_time,omitempty"`
	PlatformSummary string `json:"platform_summary,omitempty"`
	VideoEncodingVariant string `json:"video_encoding_variant,omitempty"`
	PlayerWidth int32 `json:"player_width,omitempty"`
	ViewSeekCount string `json:"view_seek_count,omitempty"`
	ViewerExperienceScore string `json:"viewer_experience_score,omitempty"`
	ViewErrorId int32 `json:"view_error_id,omitempty"`
	VideoVariantName string `json:"video_variant_name,omitempty"`
	PrerollPlayed bool `json:"preroll_played,omitempty"`
	ViewerApplicationEngine string `json:"viewer_application_engine,omitempty"`
	ViewerOsArchitecture string `json:"viewer_os_architecture,omitempty"`
	PlayerErrorCode string `json:"player_error_code,omitempty"`
	BufferingRate string `json:"buffering_rate,omitempty"`
	Events []VideoViewEvent `json:"events,omitempty"`
	PlayerName string `json:"player_name,omitempty"`
	ViewStart string `json:"view_start,omitempty"`
	ViewAverageRequestThroughput string `json:"view_average_request_throughput,omitempty"`
	VideoProducer string `json:"video_producer,omitempty"`
	ErrorTypeId int32 `json:"error_type_id,omitempty"`
	MuxViewerId string `json:"mux_viewer_id,omitempty"`
	VideoId string `json:"video_id,omitempty"`
	ContinentCode string `json:"continent_code,omitempty"`
	SessionId string `json:"session_id,omitempty"`
	ExitBeforeVideoStart bool `json:"exit_before_video_start,omitempty"`
	VideoContentType string `json:"video_content_type,omitempty"`
	ViewerOsFamily string `json:"viewer_os_family,omitempty"`
	PlayerPoster string `json:"player_poster,omitempty"`
	ViewAverageRequestLatency string `json:"view_average_request_latency,omitempty"`
	VideoVariantId string `json:"video_variant_id,omitempty"`
	PlayerSourceDuration int32 `json:"player_source_duration,omitempty"`
	PlayerSourceUrl string `json:"player_source_url,omitempty"`
	MuxApiVersion string `json:"mux_api_version,omitempty"`
	VideoTitle string `json:"video_title,omitempty"`
	Id string `json:"id,omitempty"`
	ShortTime string `json:"short_time,omitempty"`
	RebufferPercentage string `json:"rebuffer_percentage,omitempty"`
	TimeToFirstFrame string `json:"time_to_first_frame,omitempty"`
	ViewerUserId string `json:"viewer_user_id,omitempty"`
	VideoStreamType string `json:"video_stream_type,omitempty"`
	PlayerStartupTime int32 `json:"player_startup_time,omitempty"`
	ViewerApplicationVersion string `json:"viewer_application_version,omitempty"`
	ViewMaxDownscalePercentage string `json:"view_max_downscale_percentage,omitempty"`
	ViewMaxUpscalePercentage string `json:"view_max_upscale_percentage,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	UsedFullscreen bool `json:"used_fullscreen,omitempty"`
	Isp string `json:"isp,omitempty"`
	PropertyId int32 `json:"property_id,omitempty"`
	PlayerAutoplay bool `json:"player_autoplay,omitempty"`
	PlayerHeight int32 `json:"player_height,omitempty"`
	Asn int32 `json:"asn,omitempty"`
	QualityScore string `json:"quality_score,omitempty"`
	PlayerSoftwareVersion string `json:"player_software_version,omitempty"`
}
