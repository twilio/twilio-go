# DialingPermissionsBulkCountryUpdatesApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDialingPermissionsCountryBulkUpdate**](DialingPermissionsBulkCountryUpdatesApi.md#CreateDialingPermissionsCountryBulkUpdate) | **Post** /v1/DialingPermissions/BulkCountryUpdates | 



## CreateDialingPermissionsCountryBulkUpdate

> VoiceV1DialingPermissionsCountryBulkUpdate CreateDialingPermissionsCountryBulkUpdate(ctx, optional)



Create a bulk update request to change voice dialing country permissions of one or more countries identified by the corresponding [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateDialingPermissionsCountryBulkUpdateParams struct


Name | Type | Description
------------- | ------------- | -------------
**UpdateRequest** | **string** | URL encoded JSON array of update objects. example : `[ { \\\"iso_code\\\": \\\"GB\\\", \\\"low_risk_numbers_enabled\\\": \\\"true\\\", \\\"high_risk_special_numbers_enabled\\\":\\\"true\\\", \\\"high_risk_tollfraud_numbers_enabled\\\": \\\"false\\\" } ]`

### Return type

[**VoiceV1DialingPermissionsCountryBulkUpdate**](VoiceV1DialingPermissionsCountryBulkUpdate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

