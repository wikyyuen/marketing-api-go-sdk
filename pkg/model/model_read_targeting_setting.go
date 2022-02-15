/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 定向详细设置
type ReadTargetingSetting struct {
	Age                              *[]AgeStruct                       `json:"age,omitempty"`
	Gender                           *[]string                          `json:"gender,omitempty"`
	Education                        *[]string                          `json:"education,omitempty"`
	MaritalStatus                    *[]string                          `json:"marital_status,omitempty"`
	WorkingStatus                    *[]string                          `json:"working_status,omitempty"`
	GeoLocation                      *GeoLocations                      `json:"geo_location,omitempty"`
	UserOs                           *[]string                          `json:"user_os,omitempty"`
	NewDevice                        *[]string                          `json:"new_device,omitempty"`
	DevicePrice                      *[]string                          `json:"device_price,omitempty"`
	DeviceBrandModel                 *DeviceBrandModel                  `json:"device_brand_model,omitempty"`
	NetworkType                      *[]string                          `json:"network_type,omitempty"`
	NetworkOperator                  *[]string                          `json:"network_operator,omitempty"`
	NetworkScene                     *[]string                          `json:"network_scene,omitempty"`
	ConsumptionStatus                *[]string                          `json:"consumption_status,omitempty"`
	GameConsumptionLevel             *[]string                          `json:"game_consumption_level,omitempty"`
	ResidentialCommunityPrice        *[]ResidentialCommunityPriceStruct `json:"residential_community_price,omitempty"`
	FinancialSituation               *[]string                          `json:"financial_situation,omitempty"`
	ConsumptionType                  *[]string                          `json:"consumption_type,omitempty"`
	WechatAdBehavior                 *WechatAdBehavior                  `json:"wechat_ad_behavior,omitempty"`
	CustomAudience                   *[]int64                           `json:"custom_audience,omitempty"`
	ExcludedCustomAudience           *[]int64                           `json:"excluded_custom_audience,omitempty"`
	BehaviorOrInterest               *BehaviorOrInterest                `json:"behavior_or_interest,omitempty"`
	WechatOfficialAccountCategory    *[]int64                           `json:"wechat_official_account_category,omitempty"`
	MobileUnionCategory              *[]int64                           `json:"mobile_union_category,omitempty"`
	BusinessInterest                 *[]int64                           `json:"business_interest,omitempty"`
	Keyword                          *Keyword                           `json:"keyword,omitempty"`
	AppBehavior                      *AppBehavior                       `json:"app_behavior,omitempty"`
	PaidUser                         *[]string                          `json:"paid_user,omitempty"`
	DeprecatedCustomAudience         *[]int64                           `json:"deprecated_custom_audience,omitempty"`
	DeprecatedExcludedCustomAudience *[]int64                           `json:"deprecated_excluded_custom_audience,omitempty"`
	DeprecatedRegion                 *[]int64                           `json:"deprecated_region,omitempty"`
	ExcludedConvertedAudience        *ExcludedConvertedAudience         `json:"excluded_converted_audience,omitempty"`
}
