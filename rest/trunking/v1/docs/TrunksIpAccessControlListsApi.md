# TrunksIpAccessControlListsApi

All URIs are relative to *https://trunking.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIpAccessControlList**](TrunksIpAccessControlListsApi.md#CreateIpAccessControlList) | **Post** /v1/Trunks/{TrunkSid}/IpAccessControlLists | 
[**DeleteIpAccessControlList**](TrunksIpAccessControlListsApi.md#DeleteIpAccessControlList) | **Delete** /v1/Trunks/{TrunkSid}/IpAccessControlLists/{Sid} | 
[**FetchIpAccessControlList**](TrunksIpAccessControlListsApi.md#FetchIpAccessControlList) | **Get** /v1/Trunks/{TrunkSid}/IpAccessControlLists/{Sid} | 
[**ListIpAccessControlList**](TrunksIpAccessControlListsApi.md#ListIpAccessControlList) | **Get** /v1/Trunks/{TrunkSid}/IpAccessControlLists | 



## CreateIpAccessControlList

> TrunkingV1IpAccessControlList CreateIpAccessControlList(ctx, TrunkSidoptional)



Associate an IP Access Control List with a Trunk

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk to associate the IP Access Control List with.

### Other Parameters

Other parameters are passed through a pointer to a CreateIpAccessControlListParams struct


Name | Type | Description
------------- | ------------- | -------------
**IpAccessControlListSid** | **string** | The SID of the [IP Access Control List](https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource) that you want to associate with the trunk.

### Return type

[**TrunkingV1IpAccessControlList**](TrunkingV1IpAccessControlList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIpAccessControlList

> DeleteIpAccessControlList(ctx, TrunkSidSid)



Remove an associated IP Access Control List from a Trunk

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to delete the IP Access Control List.
**Sid** | **string** | The unique string that we created to identify the IpAccessControlList resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteIpAccessControlListParams struct


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


## FetchIpAccessControlList

> TrunkingV1IpAccessControlList FetchIpAccessControlList(ctx, TrunkSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to fetch the IP Access Control List.
**Sid** | **string** | The unique string that we created to identify the IpAccessControlList resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchIpAccessControlListParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrunkingV1IpAccessControlList**](TrunkingV1IpAccessControlList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIpAccessControlList

> []TrunkingV1IpAccessControlList ListIpAccessControlList(ctx, TrunkSidoptional)



List all IP Access Control Lists for a Trunk

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to read the IP Access Control Lists.

### Other Parameters

Other parameters are passed through a pointer to a ListIpAccessControlListParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]TrunkingV1IpAccessControlList**](TrunkingV1IpAccessControlList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

