// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type CreateLiveStreamRequest struct {
	PlaybackPolicy   []PlaybackPolicy   `json:"playback_policy,omitempty"`
	NewAssetSettings CreateAssetRequest `json:"new_asset_settings,omitempty"`
	// When live streaming software disconnects from Mux, either intentionally or due to a drop in the network, the Reconnect Window is the time in seconds that Mux should wait for the streaming software to reconnect before considering the live stream finished and completing the recorded asset. Defaults to 60 seconds on the API if not specified.
	ReconnectWindow float32 `json:"reconnect_window,omitempty"`
	Passthrough     string  `json:"passthrough,omitempty"`
	// Latency is the time from when the streamer does something in real life to when you see it happen in the player. Set this if you want lower latency for your live stream. Note: Reconnect windows are incompatible with Reduced Latency and will always be set to zero (0) seconds. Read more here: https://mux.com/blog/reduced-latency-for-mux-live-streaming-now-available/
	ReducedLatency bool `json:"reduced_latency,omitempty"`
}
