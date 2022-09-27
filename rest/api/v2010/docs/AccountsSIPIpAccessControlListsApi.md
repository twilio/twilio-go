# AccountsSIPIpAccessControlListsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSipIpAccessControlList**](AccountsSIPIpAccessControlListsApi.md#CreateSipIpAccessControlList) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists.json | 
[**DeleteSipIpAccessControlList**](AccountsSIPIpAccessControlListsApi.md#DeleteSipIpAccessControlList) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{Sid}.json | 
[**FetchSipIpAccessControlList**](AccountsSIPIpAccessControlListsApi.md#FetchSipIpAccessControlList) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{Sid}.json | 
[**ListSipIpAccessControlList**](AccountsSIPIpAccessControlListsApi.md#ListSipIpAccessControlList) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists.json | 
[**UpdateSipIpAccessControlList**](AccountsSIPIpAccessControlListsApi.md#UpdateSipIpAccessControlList) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{Sid}.json | 



## CreateSipIpAccessControlList

> ApiV2010SipIpAccessControlList CreateSipIpAccessControlList(ctx, optional)



Create a new IpAccessControlList resource

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipIpAccessControlListParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**FriendlyName** | **string** | A human readable descriptive text that describes the IpAccessControlList, up to 255 characters long.

### Return type

[**ApiV2010SipIpAccessControlList**](ApiV2010SipIpAccessControlList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipIpAccessControlList

> DeleteSipIpAccessControlList(ctx, Sidoptional)



Delete an IpAccessControlList from the requested account

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies the resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipIpAccessControlListParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.

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


## FetchSipIpAccessControlList

> ApiV2010SipIpAccessControlList FetchSipIpAccessControlList(ctx, Sidoptional)



Fetch a specific instance of an IpAccessControlList

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSipIpAccessControlListParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.

### Return type

[**ApiV2010SipIpAccessControlList**](ApiV2010SipIpAccessControlList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipIpAccessControlList

> []ApiV2010SipIpAccessControlList ListSipIpAccessControlList(ctx, optional)



Retrieve a list of IpAccessControlLists that belong to the account used to make the request

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSipIpAccessControlListParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010SipIpAccessControlList**](ApiV2010SipIpAccessControlList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipIpAccessControlList

> ApiV2010SipIpAccessControlList UpdateSipIpAccessControlList(ctx, Sidoptional)



Rename an IpAccessControlList

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies the resource to udpate.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSipIpAccessControlListParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**FriendlyName** | **string** | A human readable descriptive text, up to 255 characters long.

### Return type

[**ApiV2010SipIpAccessControlList**](ApiV2010SipIpAccessControlList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

