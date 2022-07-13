# AccountsAddressesApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAddress**](AccountsAddressesApi.md#CreateAddress) | **Post** /2010-04-01/Accounts/{AccountSid}/Addresses.json | 
[**DeleteAddress**](AccountsAddressesApi.md#DeleteAddress) | **Delete** /2010-04-01/Accounts/{AccountSid}/Addresses/{Sid}.json | 
[**FetchAddress**](AccountsAddressesApi.md#FetchAddress) | **Get** /2010-04-01/Accounts/{AccountSid}/Addresses/{Sid}.json | 
[**ListAddress**](AccountsAddressesApi.md#ListAddress) | **Get** /2010-04-01/Accounts/{AccountSid}/Addresses.json | 
[**UpdateAddress**](AccountsAddressesApi.md#UpdateAddress) | **Post** /2010-04-01/Accounts/{AccountSid}/Addresses/{Sid}.json | 



## CreateAddress

> ApiV2010Address CreateAddress(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateAddressParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will be responsible for the new Address resource.
**AutoCorrectAddress** | **bool** | Whether we should automatically correct the address. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. If empty or &#x60;true&#x60;, we will correct the address you provide if necessary. If &#x60;false&#x60;, we won&#39;t alter the address you provide.
**City** | **string** | The city of the new address.
**CustomerName** | **string** | The name to associate with the new address.
**EmergencyEnabled** | **bool** | Whether to enable emergency calling on the new address. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
**FriendlyName** | **string** | A descriptive string that you create to describe the new address. It can be up to 64 characters long.
**IsoCountry** | **string** | The ISO country code of the new address.
**PostalCode** | **string** | The postal code of the new address.
**Region** | **string** | The state or region of the new address.
**Street** | **string** | The number and street address of the new address.

### Return type

[**ApiV2010Address**](ApiV2010Address.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAddress

> DeleteAddress(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Address resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteAddressParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to delete.

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


## FetchAddress

> ApiV2010Address FetchAddress(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Address resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchAddressParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to fetch.

### Return type

[**ApiV2010Address**](ApiV2010Address.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAddress

> []ApiV2010Address ListAddress(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAddressParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to read.
**CustomerName** | **string** | The &#x60;customer_name&#x60; of the Address resources to read.
**FriendlyName** | **string** | The string that identifies the Address resources to read.
**IsoCountry** | **string** | The ISO country code of the Address resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010Address**](ApiV2010Address.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAddress

> ApiV2010Address UpdateAddress(ctx, Sidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Address resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateAddressParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is responsible for the Address resource to update.
**AutoCorrectAddress** | **bool** | Whether we should automatically correct the address. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. If empty or &#x60;true&#x60;, we will correct the address you provide if necessary. If &#x60;false&#x60;, we won&#39;t alter the address you provide.
**City** | **string** | The city of the address.
**CustomerName** | **string** | The name to associate with the address.
**EmergencyEnabled** | **bool** | Whether to enable emergency calling on the address. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
**FriendlyName** | **string** | A descriptive string that you create to describe the address. It can be up to 64 characters long.
**PostalCode** | **string** | The postal code of the address.
**Region** | **string** | The state or region of the address.
**Street** | **string** | The number and street address of the address.

### Return type

[**ApiV2010Address**](ApiV2010Address.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

