# AccountsIncomingPhoneNumbersAssignedAddOnsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIncomingPhoneNumberAssignedAddOn**](AccountsIncomingPhoneNumbersAssignedAddOnsApi.md#CreateIncomingPhoneNumberAssignedAddOn) | **Post** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns.json | 
[**DeleteIncomingPhoneNumberAssignedAddOn**](AccountsIncomingPhoneNumbersAssignedAddOnsApi.md#DeleteIncomingPhoneNumberAssignedAddOn) | **Delete** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{Sid}.json | 
[**FetchIncomingPhoneNumberAssignedAddOn**](AccountsIncomingPhoneNumbersAssignedAddOnsApi.md#FetchIncomingPhoneNumberAssignedAddOn) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{Sid}.json | 
[**ListIncomingPhoneNumberAssignedAddOn**](AccountsIncomingPhoneNumbersAssignedAddOnsApi.md#ListIncomingPhoneNumberAssignedAddOn) | **Get** /2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns.json | 



## CreateIncomingPhoneNumberAssignedAddOn

> ApiV2010IncomingPhoneNumberAssignedAddOn CreateIncomingPhoneNumberAssignedAddOn(ctx, ResourceSidoptional)



Assign an Add-on installation to the Number specified.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceSid** | **string** | The SID of the Phone Number to assign the Add-on.

### Other Parameters

Other parameters are passed through a pointer to a CreateIncomingPhoneNumberAssignedAddOnParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**InstalledAddOnSid** | **string** | The SID that identifies the Add-on installation.

### Return type

[**ApiV2010IncomingPhoneNumberAssignedAddOn**](ApiV2010IncomingPhoneNumberAssignedAddOn.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIncomingPhoneNumberAssignedAddOn

> DeleteIncomingPhoneNumberAssignedAddOn(ctx, ResourceSidSidoptional)



Remove the assignment of an Add-on installation from the Number specified.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceSid** | **string** | The SID of the Phone Number to which the Add-on is assigned.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteIncomingPhoneNumberAssignedAddOnParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to delete.

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIncomingPhoneNumberAssignedAddOn

> ApiV2010IncomingPhoneNumberAssignedAddOn FetchIncomingPhoneNumberAssignedAddOn(ctx, ResourceSidSidoptional)



Fetch an instance of an Add-on installation currently assigned to this Number.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceSid** | **string** | The SID of the Phone Number to which the Add-on is assigned.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchIncomingPhoneNumberAssignedAddOnParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resource to fetch.

### Return type

[**ApiV2010IncomingPhoneNumberAssignedAddOn**](ApiV2010IncomingPhoneNumberAssignedAddOn.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIncomingPhoneNumberAssignedAddOn

> []ApiV2010IncomingPhoneNumberAssignedAddOn ListIncomingPhoneNumberAssignedAddOn(ctx, ResourceSidoptional)



Retrieve a list of Add-on installations currently assigned to this Number.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ResourceSid** | **string** | The SID of the Phone Number to which the Add-on is assigned.

### Other Parameters

Other parameters are passed through a pointer to a ListIncomingPhoneNumberAssignedAddOnParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010IncomingPhoneNumberAssignedAddOn**](ApiV2010IncomingPhoneNumberAssignedAddOn.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

