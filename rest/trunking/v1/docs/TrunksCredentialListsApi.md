# TrunksCredentialListsApi

All URIs are relative to *https://trunking.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredentialList**](TrunksCredentialListsApi.md#CreateCredentialList) | **Post** /v1/Trunks/{TrunkSid}/CredentialLists | 
[**DeleteCredentialList**](TrunksCredentialListsApi.md#DeleteCredentialList) | **Delete** /v1/Trunks/{TrunkSid}/CredentialLists/{Sid} | 
[**FetchCredentialList**](TrunksCredentialListsApi.md#FetchCredentialList) | **Get** /v1/Trunks/{TrunkSid}/CredentialLists/{Sid} | 
[**ListCredentialList**](TrunksCredentialListsApi.md#ListCredentialList) | **Get** /v1/Trunks/{TrunkSid}/CredentialLists | 



## CreateCredentialList

> TrunkingV1CredentialList CreateCredentialList(ctx, TrunkSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk to associate the credential list with.

### Other Parameters

Other parameters are passed through a pointer to a CreateCredentialListParams struct


Name | Type | Description
------------- | ------------- | -------------
**CredentialListSid** | **string** | The SID of the [Credential List](https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource) that you want to associate with the trunk. Once associated, we will authenticate access to the trunk against this list.

### Return type

[**TrunkingV1CredentialList**](TrunkingV1CredentialList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCredentialList

> DeleteCredentialList(ctx, TrunkSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to delete the credential list.
**Sid** | **string** | The unique string that we created to identify the CredentialList resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCredentialListParams struct


Name | Type | Description
------------- | ------------- | -------------

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


## FetchCredentialList

> TrunkingV1CredentialList FetchCredentialList(ctx, TrunkSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to fetch the credential list.
**Sid** | **string** | The unique string that we created to identify the CredentialList resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchCredentialListParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrunkingV1CredentialList**](TrunkingV1CredentialList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredentialList

> []TrunkingV1CredentialList ListCredentialList(ctx, TrunkSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to read the credential lists.

### Other Parameters

Other parameters are passed through a pointer to a ListCredentialListParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TrunkingV1CredentialList**](TrunkingV1CredentialList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

