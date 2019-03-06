/*
 * Mux Video
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package muxvideosdk

type AssetStaticRenditions struct {
	// * `ready`: All MP4s are downloadable * `preparing`: We are preparing the MP4s * `disabled`: MP4 support was not requested or has been removed * `errored`: There was a Mux internal error that prevented the MP4s from being created 
	Status string `json:"status,omitempty"`
	Files []AssetStaticRenditionsFiles `json:"files,omitempty"`
}
