# NetworkAccessProfilesApi

All URIs are relative to *https://supersim.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkAccessProfile**](NetworkAccessProfilesApi.md#CreateNetworkAccessProfile) | **Post** /v1/NetworkAccessProfiles | 
[**FetchNetworkAccessProfile**](NetworkAccessProfilesApi.md#FetchNetworkAccessProfile) | **Get** /v1/NetworkAccessProfiles/{Sid} | 
[**ListNetworkAccessProfile**](NetworkAccessProfilesApi.md#ListNetworkAccessProfile) | **Get** /v1/NetworkAccessProfiles | 
[**UpdateNetworkAccessProfile**](NetworkAccessProfilesApi.md#UpdateNetworkAccessProfile) | **Post** /v1/NetworkAccessProfiles/{Sid} | 



## CreateNetworkAccessProfile

> SupersimV1NetworkAccessProfile CreateNetworkAccessProfile(ctx, optional)



Create a new Network Access Profile

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateNetworkAccessProfileParams struct


Name | Type | Description
------------- | ------------- | -------------
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource.
**Networks** | **[]string** | List of Network SIDs that this Network Access Profile will allow connections to.

### Return type

[**SupersimV1NetworkAccessProfile**](SupersimV1NetworkAccessProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNetworkAccessProfile

> SupersimV1NetworkAccessProfile FetchNetworkAccessProfile(ctx, Sid)



Fetch a Network Access Profile instance from your account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Network Access Profile resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchNetworkAccessProfileParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SupersimV1NetworkAccessProfile**](SupersimV1NetworkAccessProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkAccessProfile

> []SupersimV1NetworkAccessProfile ListNetworkAccessProfile(ctx, optional)



Retrieve a list of Network Access Profiles from your account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListNetworkAccessProfileParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SupersimV1NetworkAccessProfile**](SupersimV1NetworkAccessProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkAccessProfile

> SupersimV1NetworkAccessProfile UpdateNetworkAccessProfile(ctx, Sidoptional)



Updates the given properties of a Network Access Profile in your account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Network Access Profile to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateNetworkAccessProfileParams struct


Name | Type | Description
------------- | ------------- | -------------
**UniqueName** | **string** | The new unique name of the Network Access Profile.

### Return type

[**SupersimV1NetworkAccessProfile**](SupersimV1NetworkAccessProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

