# Go API client for muxvideosdk

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:
```
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:
```golang
import "./muxvideosdk"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.mux.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AssetsApi* | [**CreateAsset**](docs/AssetsApi.md#createasset) | **Post** /video/v1/assets | Create an asset
*AssetsApi* | [**CreateAssetPlaybackId**](docs/AssetsApi.md#createassetplaybackid) | **Post** /video/v1/assets/{ASSET_ID}/playback-ids | Create a playback ID
*AssetsApi* | [**DeleteAsset**](docs/AssetsApi.md#deleteasset) | **Delete** /video/v1/assets/{ASSET_ID} | Delete an asset
*AssetsApi* | [**DeleteAssetPlaybackId**](docs/AssetsApi.md#deleteassetplaybackid) | **Delete** /video/v1/assets/{ASSET_ID}/playback-ids/{PLAYBACK_ID} | Delete a playback ID
*AssetsApi* | [**GetAsset**](docs/AssetsApi.md#getasset) | **Get** /video/v1/assets/{ASSET_ID} | Retrieve an asset
*AssetsApi* | [**GetAssetInputInfo**](docs/AssetsApi.md#getassetinputinfo) | **Get** /video/v1/assets/{ASSET_ID}/input-info | Retrieve asset input info
*AssetsApi* | [**GetAssetPlaybackId**](docs/AssetsApi.md#getassetplaybackid) | **Get** /video/v1/assets/{ASSET_ID}/playback-ids/{PLAYBACK_ID} | Retrieve a playback ID
*AssetsApi* | [**ListAssets**](docs/AssetsApi.md#listassets) | **Get** /video/v1/assets | List assets
*AssetsApi* | [**UpdateAssetMp4Support**](docs/AssetsApi.md#updateassetmp4support) | **Put** /video/v1/assets/{ASSET_ID}/mp4-support | Update MP4 support
*DirectUploadsApi* | [**CancelDirectUpload**](docs/DirectUploadsApi.md#canceldirectupload) | **Put** /video/v1/uploads/{UPLOAD_ID}/cancel | Cancel a direct upload
*DirectUploadsApi* | [**CreateDirectUpload**](docs/DirectUploadsApi.md#createdirectupload) | **Post** /video/v1/uploads | Create a new direct upload URL
*DirectUploadsApi* | [**GetDirectUpload**](docs/DirectUploadsApi.md#getdirectupload) | **Get** /video/v1/uploads/{UPLOAD_ID} | Retrieve a single direct upload&#39;s info
*DirectUploadsApi* | [**ListDirectUploads**](docs/DirectUploadsApi.md#listdirectuploads) | **Get** /video/v1/uploads | List direct uploads
*ErrorsApi* | [**ListErrors**](docs/ErrorsApi.md#listerrors) | **Get** /data/v1/errors | List Errors
*ExportsApi* | [**ListExports**](docs/ExportsApi.md#listexports) | **Get** /data/v1/exports | List property video view export links
*FiltersApi* | [**ListFilterValues**](docs/FiltersApi.md#listfiltervalues) | **Get** /data/v1/filters/{FILTER_ID} | Lists values for a specific filter
*FiltersApi* | [**ListFilters**](docs/FiltersApi.md#listfilters) | **Get** /data/v1/filters | List Filters
*LiveStreamsApi* | [**CreateLiveStream**](docs/LiveStreamsApi.md#createlivestream) | **Post** /video/v1/live-streams | Create a live stream
*LiveStreamsApi* | [**CreateLiveStreamPlaybackId**](docs/LiveStreamsApi.md#createlivestreamplaybackid) | **Post** /video/v1/live-streams/{LIVE_STREAM_ID}/playback-ids | Create a live stream playback ID
*LiveStreamsApi* | [**DeleteLiveStream**](docs/LiveStreamsApi.md#deletelivestream) | **Delete** /video/v1/live-streams/{LIVE_STREAM_ID} | Delete a live stream
*LiveStreamsApi* | [**DeleteLiveStreamPlaybackId**](docs/LiveStreamsApi.md#deletelivestreamplaybackid) | **Delete** /video/v1/live-streams/{LIVE_STREAM_ID}/playback-ids/{PLAYBACK_ID} | Delete a live stream playback ID
*LiveStreamsApi* | [**GetLiveStream**](docs/LiveStreamsApi.md#getlivestream) | **Get** /video/v1/live-streams/{LIVE_STREAM_ID} | Retrieve a live stream
*LiveStreamsApi* | [**ListLiveStreams**](docs/LiveStreamsApi.md#listlivestreams) | **Get** /video/v1/live-streams | List live streams
*LiveStreamsApi* | [**ResetStreamKey**](docs/LiveStreamsApi.md#resetstreamkey) | **Post** /video/v1/live-streams/{LIVE_STREAM_ID}/reset-stream-key | Reset a live stream’s stream key
*LiveStreamsApi* | [**SignalLiveStreamComplete**](docs/LiveStreamsApi.md#signallivestreamcomplete) | **Put** /video/v1/live-streams/{LIVE_STREAM_ID}/complete | Signal a live stream is finished
*MetricsApi* | [**GetMetricTimeseriesData**](docs/MetricsApi.md#getmetrictimeseriesdata) | **Get** /data/v1/metrics/{METRIC_ID}/timeseries | Get metric timeseries data
*MetricsApi* | [**GetOverallValues**](docs/MetricsApi.md#getoverallvalues) | **Get** /data/v1/metrics/{METRIC_ID}/overall | Get Overall values
*MetricsApi* | [**ListAllMetricValues**](docs/MetricsApi.md#listallmetricvalues) | **Get** /data/v1/metrics/comparison | List all metric values
*MetricsApi* | [**ListBreakdownValues**](docs/MetricsApi.md#listbreakdownvalues) | **Get** /data/v1/metrics/{METRIC_ID}/breakdown | List breakdown values
*MetricsApi* | [**ListInsights**](docs/MetricsApi.md#listinsights) | **Get** /data/v1/metrics/{METRIC_ID}/insights | List Insights
*URLSigningKeysApi* | [**CreateUrlSigningKey**](docs/URLSigningKeysApi.md#createurlsigningkey) | **Post** /video/v1/signing-keys | Create a URL signing key
*URLSigningKeysApi* | [**DeleteUrlSigningKey**](docs/URLSigningKeysApi.md#deleteurlsigningkey) | **Delete** /video/v1/signing-keys/{SIGNING_KEY_ID} | Delete a URL signing key
*URLSigningKeysApi* | [**GetUrlSigningKey**](docs/URLSigningKeysApi.md#geturlsigningkey) | **Get** /video/v1/signing-keys/{SIGNING_KEY_ID} | Retrieve a URL signing key
*URLSigningKeysApi* | [**ListUrlSigningKeys**](docs/URLSigningKeysApi.md#listurlsigningkeys) | **Get** /video/v1/signing-keys | List URL signing keys
*VideoViewsApi* | [**GetVideoView**](docs/VideoViewsApi.md#getvideoview) | **Get** /data/v1/video-views/{VIDEO_VIEW_ID} | Get a Video View
*VideoViewsApi* | [**ListVideoViews**](docs/VideoViewsApi.md#listvideoviews) | **Get** /data/v1/video-views | List Video Views


## Documentation For Models

 - [AbridgedVideoView](docs/AbridgedVideoView.md)
 - [Asset](docs/Asset.md)
 - [AssetErrors](docs/AssetErrors.md)
 - [AssetMaster](docs/AssetMaster.md)
 - [AssetResponse](docs/AssetResponse.md)
 - [AssetStaticRenditions](docs/AssetStaticRenditions.md)
 - [AssetStaticRenditionsFiles](docs/AssetStaticRenditionsFiles.md)
 - [BreakdownValue](docs/BreakdownValue.md)
 - [CreateAssetRequest](docs/CreateAssetRequest.md)
 - [CreateLiveStreamRequest](docs/CreateLiveStreamRequest.md)
 - [CreatePlaybackIdRequest](docs/CreatePlaybackIdRequest.md)
 - [CreatePlaybackIdResponse](docs/CreatePlaybackIdResponse.md)
 - [CreateUploadRequest](docs/CreateUploadRequest.md)
 - [Error](docs/Error.md)
 - [FilterValue](docs/FilterValue.md)
 - [GetAssetInputInfoResponse](docs/GetAssetInputInfoResponse.md)
 - [GetAssetPlaybackIdResponse](docs/GetAssetPlaybackIdResponse.md)
 - [GetMetricTimeseriesDataResponse](docs/GetMetricTimeseriesDataResponse.md)
 - [GetOverallValuesResponse](docs/GetOverallValuesResponse.md)
 - [InputFile](docs/InputFile.md)
 - [InputInfo](docs/InputInfo.md)
 - [InputSettings](docs/InputSettings.md)
 - [InputSettingsOverlaySettings](docs/InputSettingsOverlaySettings.md)
 - [InputTrack](docs/InputTrack.md)
 - [Insight](docs/Insight.md)
 - [ListAllMetricValuesResponse](docs/ListAllMetricValuesResponse.md)
 - [ListAssetsResponse](docs/ListAssetsResponse.md)
 - [ListBreakdownValuesResponse](docs/ListBreakdownValuesResponse.md)
 - [ListErrorsResponse](docs/ListErrorsResponse.md)
 - [ListExportsResponse](docs/ListExportsResponse.md)
 - [ListFilterValuesResponse](docs/ListFilterValuesResponse.md)
 - [ListFiltersResponse](docs/ListFiltersResponse.md)
 - [ListFiltersResponseData](docs/ListFiltersResponseData.md)
 - [ListInsightsResponse](docs/ListInsightsResponse.md)
 - [ListLiveStreamsResponse](docs/ListLiveStreamsResponse.md)
 - [ListSigningKeysResponse](docs/ListSigningKeysResponse.md)
 - [ListUploadsResponse](docs/ListUploadsResponse.md)
 - [ListVideoViewsResponse](docs/ListVideoViewsResponse.md)
 - [LiveStream](docs/LiveStream.md)
 - [LiveStreamResponse](docs/LiveStreamResponse.md)
 - [Metric](docs/Metric.md)
 - [OverallValues](docs/OverallValues.md)
 - [PlaybackId](docs/PlaybackId.md)
 - [PlaybackPolicy](docs/PlaybackPolicy.md)
 - [Score](docs/Score.md)
 - [SignalLiveStreamCompleteResponse](docs/SignalLiveStreamCompleteResponse.md)
 - [SigningKey](docs/SigningKey.md)
 - [SigningKeyResponse](docs/SigningKeyResponse.md)
 - [Track](docs/Track.md)
 - [UpdateAssetMp4SupportRequest](docs/UpdateAssetMp4SupportRequest.md)
 - [Upload](docs/Upload.md)
 - [UploadError](docs/UploadError.md)
 - [UploadResponse](docs/UploadResponse.md)
 - [VideoView](docs/VideoView.md)
 - [VideoViewEvent](docs/VideoViewEvent.md)
 - [VideoViewResponse](docs/VideoViewResponse.md)


## Documentation For Authorization

## accessToken
- **Type**: HTTP basic authentication

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

## Author



