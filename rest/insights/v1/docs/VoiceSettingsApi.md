# VoiceSettingsApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAccountSettings**](VoiceSettingsApi.md#FetchAccountSettings) | **Get** /v1/Voice/Settings | 
[**UpdateAccountSettings**](VoiceSettingsApi.md#UpdateAccountSettings) | **Post** /v1/Voice/Settings | 



## FetchAccountSettings

> InsightsV1AccountSettings FetchAccountSettings(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchAccountSettingsParams struct


Name | Type | Description
------------- | ------------- | -------------
**SubaccountSid** | **string** | 

### Return type

[**InsightsV1AccountSettings**](InsightsV1AccountSettings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountSettings

> InsightsV1AccountSettings UpdateAccountSettings(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateAccountSettingsParams struct


Name | Type | Description
------------- | ------------- | -------------
**AdvancedFeatures** | **bool** | 
**VoiceTrace** | **bool** | 
**SubaccountSid** | **string** | 

### Return type

[**InsightsV1AccountSettings**](InsightsV1AccountSettings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

