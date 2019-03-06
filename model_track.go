/*
 * Mux Video
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package muxvideosdk

type Track struct {
	Id string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	Duration float64 `json:"duration,omitempty"`
	MaxWidth int64 `json:"max_width,omitempty"`
	MaxHeight int64 `json:"max_height,omitempty"`
	MaxFrameRate float64 `json:"max_frame_rate,omitempty"`
	MaxChannels int64 `json:"max_channels,omitempty"`
	MaxChannelLayout string `json:"max_channel_layout,omitempty"`
	TextTrackType string `json:"text_track_type,omitempty"`
	Language string `json:"language,omitempty"`
}
