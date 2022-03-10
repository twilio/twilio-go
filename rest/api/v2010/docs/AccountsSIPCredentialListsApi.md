# AccountsSIPCredentialListsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSipCredentialList**](AccountsSIPCredentialListsApi.md#CreateSipCredentialList) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists.json | 
[**DeleteSipCredentialList**](AccountsSIPCredentialListsApi.md#DeleteSipCredentialList) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{Sid}.json | 
[**FetchSipCredentialList**](AccountsSIPCredentialListsApi.md#FetchSipCredentialList) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{Sid}.json | 
[**ListSipCredentialList**](AccountsSIPCredentialListsApi.md#ListSipCredentialList) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists.json | 
[**UpdateSipCredentialList**](AccountsSIPCredentialListsApi.md#UpdateSipCredentialList) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{Sid}.json | 



## CreateSipCredentialList

> ApiV2010SipCredentialList CreateSipCredentialList(ctx, optional)



Create a Credential List

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipCredentialListParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**FriendlyName** | **string** | A human readable descriptive text that describes the CredentialList, up to 64 characters long.

### Return type

[**ApiV2010SipCredentialList**](ApiV2010SipCredentialList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipCredentialList

> DeleteSipCredentialList(ctx, Sidoptional)



Delete a Credential List

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The credential list Sid that uniquely identifies this resource

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipCredentialListParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the Account that is responsible for this resource.

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


## FetchSipCredentialList

> ApiV2010SipCredentialList FetchSipCredentialList(ctx, Sidoptional)



Get a Credential List

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The credential list Sid that uniquely identifies this resource

### Other Parameters

Other parameters are passed through a pointer to a FetchSipCredentialListParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the Account that is responsible for this resource.

### Return type

[**ApiV2010SipCredentialList**](ApiV2010SipCredentialList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipCredentialList

> []ApiV2010SipCredentialList ListSipCredentialList(ctx, optional)



Get All Credential Lists

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSipCredentialListParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010SipCredentialList**](ApiV2010SipCredentialList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipCredentialList

> ApiV2010SipCredentialList UpdateSipCredentialList(ctx, Sidoptional)



Update a Credential List

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The credential list Sid that uniquely identifies this resource

### Other Parameters

Other parameters are passed through a pointer to a UpdateSipCredentialListParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**FriendlyName** | **string** | A human readable descriptive text for a CredentialList, up to 64 characters long.

### Return type

[**ApiV2010SipCredentialList**](ApiV2010SipCredentialList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

