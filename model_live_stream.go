/*
 * Mux Video
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package muxvideosdk

type LiveStream struct {
	Id string `json:"id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	StreamKey string `json:"stream_key,omitempty"`
	ActiveAssetId string `json:"active_asset_id,omitempty"`
	RecentAssetIds []string `json:"recent_asset_ids,omitempty"`
	Status string `json:"status,omitempty"`
	PlaybackIds []PlaybackId `json:"playback_ids,omitempty"`
	NewAssetSettings Asset `json:"new_asset_settings,omitempty"`
	Passthrough string `json:"passthrough,omitempty"`
	ReconnectWindow float64 `json:"reconnect_window,omitempty"`
}
