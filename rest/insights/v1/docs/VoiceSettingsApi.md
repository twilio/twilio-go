# VoiceSettingsApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAccountSettings**](VoiceSettingsApi.md#FetchAccountSettings) | **Get** /v1/Voice/Settings | Get the Voice Insights Settings.
[**UpdateAccountSettings**](VoiceSettingsApi.md#UpdateAccountSettings) | **Post** /v1/Voice/Settings | Update a specific Voice Insights Setting.



## FetchAccountSettings

> InsightsV1AccountSettings FetchAccountSettings(ctx, optional)

Get the Voice Insights Settings.

Get the Voice Insights Settings.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchAccountSettingsParams struct


Name | Type | Description
------------- | ------------- | -------------
**SubaccountSid** | **string** | The unique SID identifier of the Subaccount.

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

Update a specific Voice Insights Setting.

Update a specific Voice Insights Setting.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateAccountSettingsParams struct


Name | Type | Description
------------- | ------------- | -------------
**AdvancedFeatures** | **bool** | A boolean flag to enable Advanced Features for Voice Insights.
**VoiceTrace** | **bool** | A boolean flag to enable Voice Trace.
**SubaccountSid** | **string** | The unique SID identifier of the Subaccount.

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

