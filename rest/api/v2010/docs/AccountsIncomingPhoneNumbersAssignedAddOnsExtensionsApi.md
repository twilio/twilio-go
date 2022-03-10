# AccountsIncomingPhoneNumbersAssignedAddOnsExtensionsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchIncomingPhoneNumberAssignedAddOnExtension**](AccountsIncomingPhoneNumbersAssignedAddOnsExtensionsApi.md#FetchIncomingPhoneNumberAssignedAddOnExtension) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{AssignedAddOnSid}/Extensions/{Sid}.json | 
[**ListIncomingPhoneNumberAssignedAddOnExtension**](AccountsIncomingPhoneNumbersAssignedAddOnsExtensionsApi.md#ListIncomingPhoneNumberAssignedAddOnExtension) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{AssignedAddOnSid}/Extensions.json | 



## FetchIncomingPhoneNumberAssignedAddOnExtension

> ApiV2010IncomingPhoneNumberAssignedAddOnExtension FetchIncomingPhoneNumberAssignedAddOnExtension(ctx, ResourceSidAssignedAddOnSidSidoptional)



Fetch an instance of an Extension for the Assigned Add-on.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceSid** | **string** | The SID of the Phone Number to which the Add-on is assigned.
**AssignedAddOnSid** | **string** | The SID that uniquely identifies the assigned Add-on installation.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchIncomingPhoneNumberAssignedAddOnExtensionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resource to fetch.

### Return type

[**ApiV2010IncomingPhoneNumberAssignedAddOnExtension**](ApiV2010IncomingPhoneNumberAssignedAddOnExtension.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncomingPhoneNumberAssignedAddOnExtension

> []ApiV2010IncomingPhoneNumberAssignedAddOnExtension ListIncomingPhoneNumberAssignedAddOnExtension(ctx, ResourceSidAssignedAddOnSidoptional)



Retrieve a list of Extensions for the Assigned Add-on.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceSid** | **string** | The SID of the Phone Number to which the Add-on is assigned.
**AssignedAddOnSid** | **string** | The SID that uniquely identifies the assigned Add-on installation.

### Other Parameters

Other parameters are passed through a pointer to a ListIncomingPhoneNumberAssignedAddOnExtensionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010IncomingPhoneNumberAssignedAddOnExtension**](ApiV2010IncomingPhoneNumberAssignedAddOnExtension.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

