# SettingsUpdatesApi

All URIs are relative to *https://supersim.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSettingsUpdate**](SettingsUpdatesApi.md#ListSettingsUpdate) | **Get** /v1/SettingsUpdates | 



## ListSettingsUpdate

> []SupersimV1SettingsUpdate ListSettingsUpdate(ctx, optional)



Retrieve a list of Settings Updates.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSettingsUpdateParams struct


Name | Type | Description
------------- | ------------- | -------------
**Sim** | **string** | Filter the Settings Updates by a Super SIM's SID or UniqueName.
**Status** | **string** | Filter the Settings Updates by status. Can be `scheduled`, `in-progress`, `successful`, or `failed`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SupersimV1SettingsUpdate**](SupersimV1SettingsUpdate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

