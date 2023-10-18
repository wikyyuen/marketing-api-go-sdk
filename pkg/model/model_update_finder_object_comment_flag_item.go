/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 评论精选
type UpdateFinderObjectCommentFlagItem struct {
	AdcreativeId *int64                     `json:"adcreative_id,omitempty"`
	OpType       SetObjectCommentFlagOpType `json:"op_type,omitempty"`
	AccountId    *int64                     `json:"account_id,omitempty"`
	CommentId    *string                    `json:"comment_id,omitempty"`
}