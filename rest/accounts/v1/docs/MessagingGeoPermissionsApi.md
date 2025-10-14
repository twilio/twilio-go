# MessagingGeoPermissionsApi

All URIs are relative to *https://accounts.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchMessagingGeopermissions**](MessagingGeoPermissionsApi.md#FetchMessagingGeopermissions) | **Get** /v1/Messaging/GeoPermissions | 
[**UpdateMessagingGeopermissions**](MessagingGeoPermissionsApi.md#UpdateMessagingGeopermissions) | **Patch** /v1/Messaging/GeoPermissions | 



## FetchMessagingGeopermissions

> AccountsV1MessagingGeopermissions FetchMessagingGeopermissions(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchMessagingGeopermissionsParams struct


Name | Type | Description
------------- | ------------- | -------------
**CountryCode** | **string** | The country code to filter the geo permissions. If provided, only the geo permission for the specified country will be returned.

### Return type

[**AccountsV1MessagingGeopermissions**](AccountsV1MessagingGeopermissions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessagingGeopermissions

> AccountsV1MessagingGeopermissions UpdateMessagingGeopermissions(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateMessagingGeopermissionsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Permissions** | **[]interface{}** | A list of objects where each object represents the Geo Permission to be updated. Each object contains the following fields: `country_code`, unique code for each country of Geo Permission; `type`, permission type of the Geo Permission i.e. country; `enabled`, configure true for enabling the Geo Permission, false for disabling the Geo Permission.

### Return type

[**AccountsV1MessagingGeopermissions**](AccountsV1MessagingGeopermissions.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

