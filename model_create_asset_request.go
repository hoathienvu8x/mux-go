// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxvideosdk

type CreateAssetRequest struct {
	Input []InputSettings `json:"input,omitempty"`
	PlaybackPolicy []PlaybackPolicy `json:"playback_policy,omitempty"`
	Demo bool `json:"demo,omitempty"`
	PerTitleEncode bool `json:"per_title_encode,omitempty"`
	Passthrough string `json:"passthrough,omitempty"`
	Mp4Support string `json:"mp4_support,omitempty"`
}