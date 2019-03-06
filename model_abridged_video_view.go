/*
 * Mux Video
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package muxvideosdk

type AbridgedVideoView struct {
	Id string `json:"id,omitempty"`
	ViewerOsFamily string `json:"viewer_os_family,omitempty"`
	ViewerApplicationName string `json:"viewer_application_name,omitempty"`
	VideoTitle string `json:"video_title,omitempty"`
	TotalRowCount int32 `json:"total_row_count,omitempty"`
	PlayerErrorMessage string `json:"player_error_message,omitempty"`
	PlayerErrorCode string `json:"player_error_code,omitempty"`
	ErrorTypeId int32 `json:"error_type_id,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	ViewStart string `json:"view_start,omitempty"`
	ViewEnd string `json:"view_end,omitempty"`
}
