/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// OptimizationGoal : 广告优化目标类型
type OptimizationGoal string

// List of OptimizationGoal
const (
	OptimizationGoal_NONE                                   OptimizationGoal = "OPTIMIZATIONGOAL_NONE"
	OptimizationGoal_BRAND_CONVERSION                       OptimizationGoal = "OPTIMIZATIONGOAL_BRAND_CONVERSION"
	OptimizationGoal_FOLLOW                                 OptimizationGoal = "OPTIMIZATIONGOAL_FOLLOW"
	OptimizationGoal_CLICK                                  OptimizationGoal = "OPTIMIZATIONGOAL_CLICK"
	OptimizationGoal_IMPRESSION                             OptimizationGoal = "OPTIMIZATIONGOAL_IMPRESSION"
	OptimizationGoal_APP_DOWNLOAD                           OptimizationGoal = "OPTIMIZATIONGOAL_APP_DOWNLOAD"
	OptimizationGoal_APP_ACTIVATE                           OptimizationGoal = "OPTIMIZATIONGOAL_APP_ACTIVATE"
	OptimizationGoal_APP_REGISTER                           OptimizationGoal = "OPTIMIZATIONGOAL_APP_REGISTER"
	OptimizationGoal_ONE_DAY_RETENTION                      OptimizationGoal = "OPTIMIZATIONGOAL_ONE_DAY_RETENTION"
	OptimizationGoal_APP_PURCHASE                           OptimizationGoal = "OPTIMIZATIONGOAL_APP_PURCHASE"
	OptimizationGoal_ECOMMERCE_ORDER                        OptimizationGoal = "OPTIMIZATIONGOAL_ECOMMERCE_ORDER"
	OptimizationGoal_ECOMMERCE_CHECKOUT                     OptimizationGoal = "OPTIMIZATIONGOAL_ECOMMERCE_CHECKOUT"
	OptimizationGoal_LEADS                                  OptimizationGoal = "OPTIMIZATIONGOAL_LEADS"
	OptimizationGoal_ECOMMERCE_CART                         OptimizationGoal = "OPTIMIZATIONGOAL_ECOMMERCE_CART"
	OptimizationGoal_PROMOTION_CLICK_KEY_PAGE               OptimizationGoal = "OPTIMIZATIONGOAL_PROMOTION_CLICK_KEY_PAGE"
	OptimizationGoal_VIEW_COMMODITY_PAGE                    OptimizationGoal = "OPTIMIZATIONGOAL_VIEW_COMMODITY_PAGE"
	OptimizationGoal_ONLINE_CONSULTATION                    OptimizationGoal = "OPTIMIZATIONGOAL_ONLINE_CONSULTATION"
	OptimizationGoal_TELEPHONE_CONSULTATION                 OptimizationGoal = "OPTIMIZATIONGOAL_TELEPHONE_CONSULTATION"
	OptimizationGoal_PAGE_RESERVATION                       OptimizationGoal = "OPTIMIZATIONGOAL_PAGE_RESERVATION"
	OptimizationGoal_DELIVERY                               OptimizationGoal = "OPTIMIZATIONGOAL_DELIVERY"
	OptimizationGoal_MESSAGE_AFTER_FOLLOW                   OptimizationGoal = "OPTIMIZATIONGOAL_MESSAGE_AFTER_FOLLOW"
	OptimizationGoal_CLICK_MENU_AFTER_FOLLOW                OptimizationGoal = "OPTIMIZATIONGOAL_CLICK_MENU_AFTER_FOLLOW"
	OptimizationGoal_PAGE_EFFECTIVE_ONLINE_CONSULT          OptimizationGoal = "OPTIMIZATIONGOAL_PAGE_EFFECTIVE_ONLINE_CONSULT"
	OptimizationGoal_CLICK_KEY_PAGE                         OptimizationGoal = "OPTIMIZATIONGOAL_CLICK_KEY_PAGE"
	OptimizationGoal_MOBILE_APP_START                       OptimizationGoal = "OPTIMIZATIONGOAL_MOBILE_APP_START"
	OptimizationGoal_PAGE_DELIVER                           OptimizationGoal = "OPTIMIZATIONGOAL_PAGE_DELIVER"
	OptimizationGoal_PAGE_MAKE_PHONE_CALL                   OptimizationGoal = "OPTIMIZATIONGOAL_PAGE_MAKE_PHONE_CALL"
	OptimizationGoal_PAGE_ONLINE_CONSULT                    OptimizationGoal = "OPTIMIZATIONGOAL_PAGE_ONLINE_CONSULT"
	OptimizationGoal_MOBILE_APP_CHECKOUT                    OptimizationGoal = "OPTIMIZATIONGOAL_MOBILE_APP_CHECKOUT"
	OptimizationGoal_APP_INSTALL                            OptimizationGoal = "OPTIMIZATIONGOAL_APP_INSTALL"
	OptimizationGoal_PAGE_EFFECTIVE_PHONE_CALL              OptimizationGoal = "OPTIMIZATIONGOAL_PAGE_EFFECTIVE_PHONE_CALL"
	OptimizationGoal_CONFIRM_EFFECTIVE_LEADS_CONSULT        OptimizationGoal = "OPTIMIZATIONGOAL_CONFIRM_EFFECTIVE_LEADS_CONSULT"
	OptimizationGoal_CONFIRM_EFFECTIVE_LEADS_PHONE          OptimizationGoal = "OPTIMIZATIONGOAL_CONFIRM_EFFECTIVE_LEADS_PHONE"
	OptimizationGoal_LEADS_COLLECT                          OptimizationGoal = "OPTIMIZATIONGOAL_LEADS_COLLECT"
	OptimizationGoal_FIRST_PURCHASE                         OptimizationGoal = "OPTIMIZATIONGOAL_FIRST_PURCHASE"
	OptimizationGoal_APPLY                                  OptimizationGoal = "OPTIMIZATIONGOAL_APPLY"
	OptimizationGoal_PRE_CREDIT                             OptimizationGoal = "OPTIMIZATIONGOAL_PRE_CREDIT"
	OptimizationGoal_CREDIT                                 OptimizationGoal = "OPTIMIZATIONGOAL_CREDIT"
	OptimizationGoal_WITHDRAW_DEPOSITS                      OptimizationGoal = "OPTIMIZATIONGOAL_WITHDRAW_DEPOSITS"
	OptimizationGoal_PROMOTION_VIEW_KEY_PAGE                OptimizationGoal = "OPTIMIZATIONGOAL_PROMOTION_VIEW_KEY_PAGE"
	OptimizationGoal_MOBILE_APP_CREATE_ROLE                 OptimizationGoal = "OPTIMIZATIONGOAL_MOBILE_APP_CREATE_ROLE"
	OptimizationGoal_CANVAS_CLICK                           OptimizationGoal = "OPTIMIZATIONGOAL_CANVAS_CLICK"
	OptimizationGoal_PROMOTION_CLAIM_OFFER                  OptimizationGoal = "OPTIMIZATIONGOAL_PROMOTION_CLAIM_OFFER"
	OptimizationGoal_ECOMMERCE_ADD_TO_WISHLIST              OptimizationGoal = "OPTIMIZATIONGOAL_ECOMMERCE_ADD_TO_WISHLIST"
	OptimizationGoal_CONFIRM_EFFECTIVE_LEADS_RESERVATION    OptimizationGoal = "OPTIMIZATIONGOAL_CONFIRM_EFFECTIVE_LEADS_RESERVATION"
	OptimizationGoal_PAGE_RECEIPT                           OptimizationGoal = "OPTIMIZATIONGOAL_PAGE_RECEIPT"
	OptimizationGoal_PAGE_SCAN_CODE                         OptimizationGoal = "OPTIMIZATIONGOAL_PAGE_SCAN_CODE"
	OptimizationGoal_SELECT_COURSE                          OptimizationGoal = "OPTIMIZATIONGOAL_SELECT_COURSE"
	OptimizationGoal_CONFIRM_POTENTIAL_CUSTOMER_PHONE       OptimizationGoal = "OPTIMIZATIONGOAL_CONFIRM_POTENTIAL_CUSTOMER_PHONE"
	OptimizationGoal_MOBILE_APP_AD_INCOME                   OptimizationGoal = "OPTIMIZATIONGOAL_MOBILE_APP_AD_INCOME"
	OptimizationGoal_MOBILE_APP_ACCREDIT                    OptimizationGoal = "OPTIMIZATIONGOAL_MOBILE_APP_ACCREDIT"
	OptimizationGoal_PURCHASE_MEMBER_CARD                   OptimizationGoal = "OPTIMIZATIONGOAL_PURCHASE_MEMBER_CARD"
	OptimizationGoal_PAGE_CONFIRM_EFFECTIVE_LEADS           OptimizationGoal = "OPTIMIZATIONGOAL_PAGE_CONFIRM_EFFECTIVE_LEADS"
	OptimizationGoal_ADD_DESKTOP                            OptimizationGoal = "OPTIMIZATIONGOAL_ADD_DESKTOP"
	OptimizationGoal_RESERVATION                            OptimizationGoal = "OPTIMIZATIONGOAL_RESERVATION"
	OptimizationGoal_FIRST_ECOMMERCE_ORDER                  OptimizationGoal = "OPTIMIZATIONGOAL_FIRST_ECOMMERCE_ORDER"
	OptimizationGoal_FIRST_TWENTY_FOUR_HOUR_ECOMMERCE_ORDER OptimizationGoal = "OPTIMIZATIONGOAL_FIRST_TWENTY_FOUR_HOUR_ECOMMERCE_ORDER"
	OptimizationGoal_ECOMMERCE_SCANCODE_WX                  OptimizationGoal = "OPTIMIZATIONGOAL_ECOMMERCE_SCANCODE_WX"
	OptimizationGoal_CLASS_PARTICIPATED                     OptimizationGoal = "OPTIMIZATIONGOAL_CLASS_PARTICIPATED"
	OptimizationGoal_INSURANCE_PURCHASE                     OptimizationGoal = "OPTIMIZATIONGOAL_INSURANCE_PURCHASE"
	OptimizationGoal_MOBILE_APP_SEVEN_DAYS_RETENTION        OptimizationGoal = "OPTIMIZATIONGOAL_MOBILE_APP_SEVEN_DAYS_RETENTION"
	OptimizationGoal_LIKE                                   OptimizationGoal = "OPTIMIZATIONGOAL_LIKE"
	OptimizationGoal_EXTERNAL_LINK_CLICK                    OptimizationGoal = "OPTIMIZATIONGOAL_EXTERNAL_LINK_CLICK"
	OptimizationGoal_BUY_COUPONS                            OptimizationGoal = "OPTIMIZATIONGOAL_BUY_COUPONS"
	OptimizationGoal_LEAVE_INFORMATION                      OptimizationGoal = "OPTIMIZATIONGOAL_LEAVE_INFORMATION"
	OptimizationGoal_CORE_ACTION                            OptimizationGoal = "OPTIMIZATIONGOAL_CORE_ACTION"
	OptimizationGoal_ONE_DAY_RETENTION_RATIO                OptimizationGoal = "OPTIMIZATIONGOAL_ONE_DAY_RETENTION_RATIO"
	OptimizationGoal_PROMOTION_READ_ARTICLE                 OptimizationGoal = "OPTIMIZATIONGOAL_PROMOTION_READ_ARTICLE"
	OptimizationGoal_RESERVATION_CHECK                      OptimizationGoal = "OPTIMIZATIONGOAL_RESERVATION_CHECK"
	OptimizationGoal_OPEN_ACCOUNT                           OptimizationGoal = "OPTIMIZATIONGOAL_OPEN_ACCOUNT"
	OptimizationGoal_SEVEN_DAY_ECOMMERCE_ORDER              OptimizationGoal = "OPTIMIZATIONGOAL_SEVEN_DAY_ECOMMERCE_ORDER"
	OptimizationGoal_ADD_WECHAT                             OptimizationGoal = "OPTIMIZATIONGOAL_ADD_WECHAT"
)
