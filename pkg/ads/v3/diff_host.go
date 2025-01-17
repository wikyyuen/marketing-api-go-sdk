/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ads

import (
	"net/url"
)

var SpecialUrlInterfaceList = map[string]url.URL{
	"/async_task_files/get":      url.URL{Host: "https://dl.e.qq.com/v1.1", Path: "/async_task_files/get"},
	"/async_report_files/get":    url.URL{Host: "https://dl.e.qq.com/v1.1", Path: "/async_report_files/get"},
	"/authorization/wechat_bind": url.URL{Host: "https://developers.e.qq.com", Path: "/authorization/wechat_bind"},
	"/oauth/token":               url.URL{Host: "https://api.e.qq.com", Path: "/oauth/token"},
	"/oauth/authorize":           url.URL{Host: "https://developers.e.qq.com", Path: "/oauth/authorize"},
}
