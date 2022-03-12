# NetworkAccessProfilesNetworksApi

All URIs are relative to *https://supersim.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkAccessProfileNetwork**](NetworkAccessProfilesNetworksApi.md#CreateNetworkAccessProfileNetwork) | **Post** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks | 
[**DeleteNetworkAccessProfileNetwork**](NetworkAccessProfilesNetworksApi.md#DeleteNetworkAccessProfileNetwork) | **Delete** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks/{Sid} | 
[**FetchNetworkAccessProfileNetwork**](NetworkAccessProfilesNetworksApi.md#FetchNetworkAccessProfileNetwork) | **Get** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks/{Sid} | 
[**ListNetworkAccessProfileNetwork**](NetworkAccessProfilesNetworksApi.md#ListNetworkAccessProfileNetwork) | **Get** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks | 



## CreateNetworkAccessProfileNetwork

> SupersimV1NetworkAccessProfileNetwork CreateNetworkAccessProfileNetwork(ctx, NetworkAccessProfileSidoptional)



Add a Network resource to the Network Access Profile resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**NetworkAccessProfileSid** | **string** | The unique string that identifies the Network Access Profile resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateNetworkAccessProfileNetworkParams struct


Name | Type | Description
------------- | ------------- | -------------
**Network** | **string** | The SID of the Network resource to be added to the Network Access Profile resource.

### Return type

[**SupersimV1NetworkAccessProfileNetwork**](SupersimV1NetworkAccessProfileNetwork.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkAccessProfileNetwork

> DeleteNetworkAccessProfileNetwork(ctx, NetworkAccessProfileSidSid)



Remove a Network resource from the Network Access Profile resource's.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**NetworkAccessProfileSid** | **string** | The unique string that identifies the Network Access Profile resource.
**Sid** | **string** | The SID of the Network resource to be removed from the Network Access Profile resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteNetworkAccessProfileNetworkParams struct


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


## FetchNetworkAccessProfileNetwork

> SupersimV1NetworkAccessProfileNetwork FetchNetworkAccessProfileNetwork(ctx, NetworkAccessProfileSidSid)



Fetch a Network Access Profile resource's Network resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**NetworkAccessProfileSid** | **string** | The unique string that identifies the Network Access Profile resource.
**Sid** | **string** | The SID of the Network resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchNetworkAccessProfileNetworkParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SupersimV1NetworkAccessProfileNetwork**](SupersimV1NetworkAccessProfileNetwork.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkAccessProfileNetwork

> []SupersimV1NetworkAccessProfileNetwork ListNetworkAccessProfileNetwork(ctx, NetworkAccessProfileSidoptional)



Retrieve a list of Network Access Profile resource's Network resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**NetworkAccessProfileSid** | **string** | The unique string that identifies the Network Access Profile resource.

### Other Parameters

Other parameters are passed through a pointer to a ListNetworkAccessProfileNetworkParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SupersimV1NetworkAccessProfileNetwork**](SupersimV1NetworkAccessProfileNetwork.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

