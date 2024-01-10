# TencentAds\ImagesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImagesAdd**](ImagesApi.md#ImagesAdd) | **Get** /images/add | 添加图片文件
[**ImagesDelete**](ImagesApi.md#ImagesDelete) | **Get** /images/delete | 删除图片
[**ImagesGet**](ImagesApi.md#ImagesGet) | **Get** /images/get | 获取图片信息
[**ImagesUpdate**](ImagesApi.md#ImagesUpdate) | **Get** /images/update | 修改图片信息


# **ImagesAdd**
> ImagesAddResponse ImagesAdd(ctx, advertiserId, uploadType, imageSignature, optional)
添加图片文件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **advertiserId** | **int64**|  | 
  **uploadType** | **string**|  | 
  **imageSignature** | **string**|  | 
 **optional** | ***ImagesApiImagesAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImagesApiImagesAddOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **imageFile** | **optional.Interface of *os.File**|  | 
 **bytes** | **optional.String**|  | 
 **imageUsage** | **optional.String**|  | 
 **description** | **optional.String**|  | 
 **resizeWidth** | **optional.Int64**|  | 
 **resizeHeight** | **optional.Int64**|  | 
 **resizeFileSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ImagesAddResponse**](ImagesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImagesDelete**
> ImagesDeleteResponse ImagesDelete(ctx, advertiserId, imageId, optional)
删除图片

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **advertiserId** | **int64**|  | 
  **imageId** | **string**|  | 
 **optional** | ***ImagesApiImagesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImagesApiImagesDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ImagesDeleteResponse**](ImagesDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImagesGet**
> ImagesGetResponse ImagesGet(ctx, accountId, optional)
获取图片信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***ImagesApiImagesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImagesApiImagesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **labelId** | **optional.Int64**|  | 
 **businessScenario** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ImagesGetResponse**](ImagesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImagesUpdate**
> ImagesUpdateResponse ImagesUpdate(ctx, advertiserId, imageId, description, optional)
修改图片信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **advertiserId** | **int64**|  | 
  **imageId** | **string**|  | 
  **description** | **string**|  | 
 **optional** | ***ImagesApiImagesUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImagesApiImagesUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ImagesUpdateResponse**](ImagesUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
