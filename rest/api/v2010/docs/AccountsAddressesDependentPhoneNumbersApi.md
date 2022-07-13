# AccountsAddressesDependentPhoneNumbersApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDependentPhoneNumber**](AccountsAddressesDependentPhoneNumbersApi.md#ListDependentPhoneNumber) | **Get** /2010-04-01/Accounts/{AccountSid}/Addresses/{AddressSid}/DependentPhoneNumbers.json | 



## ListDependentPhoneNumber

> []ApiV2010DependentPhoneNumber ListDependentPhoneNumber(ctx, AddressSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**AddressSid** | **string** | The SID of the Address resource associated with the phone number.

### Other Parameters

Other parameters are passed through a pointer to a ListDependentPhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the DependentPhoneNumber resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010DependentPhoneNumber**](ApiV2010DependentPhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

