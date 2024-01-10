/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tencentad/marketing-api-go-sdk/pkg/ads/v3"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api/v3"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config/v3"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors/v3"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model/v3"
)

type GameFeatureGetExample struct {
	TAds                    *ads.SDKClient
	AccessToken             string
	AccountId               int64
	MarketingTargetType     string
	MarketingTargetDetailId string
	GameFeatureGetOpts      *api.GameFeatureGetOpts
}

func (e *GameFeatureGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = 789
	e.MarketingTargetType = "marketingTargetType_example"
	e.MarketingTargetDetailId = "marketingTargetDetailId_example"
	e.GameFeatureGetOpts = &api.GameFeatureGetOpts{}
}

func (e *GameFeatureGetExample) RunExample() (model.GameFeatureGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.GameFeature().Get(ctx, e.AccountId, e.MarketingTargetType, e.MarketingTargetDetailId, e.GameFeatureGetOpts)
}

func main() {
	e := &GameFeatureGetExample{}
	e.Init()
	response, headers, err := e.RunExample()
	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Response error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("Response data:", response)
	fmt.Println("Headers:", headers)
}