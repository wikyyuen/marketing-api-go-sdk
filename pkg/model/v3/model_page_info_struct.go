/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 排序
type PageInfoStruct struct {
	Page  *int64 `json:"page,omitempty"`
	Rows  *int64 `json:"rows,omitempty"`
	Total *int64 `json:"total,omitempty"`
}