# \SubaccountcleanlistApi

All URIs are relative to *https://api.sendpost.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCleaningRouterCleanBulkEmailList**](SubaccountcleanlistApi.md#ListCleaningRouterCleanBulkEmailList) | **Post** /subaccount/cleanlist/bulk | 
[**ListCleaningRouterCleanEmailist**](SubaccountcleanlistApi.md#ListCleaningRouterCleanEmailist) | **Post** /subaccount/cleanlist/ | 


# **ListCleaningRouterCleanBulkEmailList**
> ModelsCleanedList ListCleaningRouterCleanBulkEmailList(ctx, xSubAccountApiKey)


Send Email To Contacts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xSubAccountApiKey** | **string**| Sub-Account API Key | 

### Return type

[**ModelsCleanedList**](models.CleanedList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCleaningRouterCleanEmailist**
> ModelsCleanedList ListCleaningRouterCleanEmailist(ctx, xSubAccountApiKey)


Clean email list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xSubAccountApiKey** | **string**| Sub-Account API Key | 

### Return type

[**ModelsCleanedList**](models.CleanedList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

