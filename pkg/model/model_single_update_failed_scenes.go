/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 应用场景信息
type SingleUpdateFailedScenes struct {
	Scene      DataNexusScene                  `json:"scene,omitempty"`
	AssetIds   *[]SingleUpdateFailedAssetArray `json:"asset_ids,omitempty"`
	FailReason *string                         `json:"fail_reason,omitempty"`
}