# SettingsApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchDialingPermissionsSettings**](SettingsApi.md#FetchDialingPermissionsSettings) | **Get** /v1/Settings | 
[**UpdateDialingPermissionsSettings**](SettingsApi.md#UpdateDialingPermissionsSettings) | **Post** /v1/Settings | 



## FetchDialingPermissionsSettings

> VoiceV1DialingPermissionsSettings FetchDialingPermissionsSettings(ctx, )



Retrieve voice dialing permissions inheritance for the sub-account

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchDialingPermissionsSettingsParams struct


### Return type

[**VoiceV1DialingPermissionsSettings**](VoiceV1DialingPermissionsSettings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDialingPermissionsSettings

> VoiceV1DialingPermissionsSettings UpdateDialingPermissionsSettings(ctx, optional)



Update voice dialing permissions inheritance for the sub-account

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateDialingPermissionsSettingsParams struct


Name | Type | Description
------------- | ------------- | -------------
**DialingPermissionsInheritance** | **bool** | `true` for the sub-account to inherit voice dialing permissions from the Master Project; otherwise `false`.

### Return type

[**VoiceV1DialingPermissionsSettings**](VoiceV1DialingPermissionsSettings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

