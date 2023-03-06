# RouteConfigurationsTrunksApi

All URIs are relative to *https://routes.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTrunk**](RouteConfigurationsTrunksApi.md#CreateTrunk) | **Post** /v1/RouteConfigurations/{RouteConfigurationSid}/Trunks | 
[**DeleteTrunk**](RouteConfigurationsTrunksApi.md#DeleteTrunk) | **Delete** /v1/RouteConfigurations/{RouteConfigurationSid}/Trunks | 
[**FetchTrunk**](RouteConfigurationsTrunksApi.md#FetchTrunk) | **Get** /v1/RouteConfigurations/{RouteConfigurationSid}/Trunks | 



## CreateTrunk

> RoutesV1Trunk CreateTrunk(ctx, RouteConfigurationSidoptional)



Link a Resource to an existing Route Configuration.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RouteConfigurationSid** | **string** | The unique string that we created to identify the Route Configuration resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateTrunkParams struct


Name | Type | Description
------------- | ------------- | -------------
**Trunk** | **string** | TBD

### Return type

[**RoutesV1Trunk**](RoutesV1Trunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrunk

> DeleteTrunk(ctx, RouteConfigurationSid)



Delete a specific Resource linked to a Route Configuration instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RouteConfigurationSid** | **string** | The unique string that we created to identify the Route Configuration resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteTrunkParams struct


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


## FetchTrunk

> RoutesV1Trunk FetchTrunk(ctx, RouteConfigurationSid)



Fetch a specific Resource linked to a Route Configuration instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**RouteConfigurationSid** | **string** | The unique string that we created to identify the Route Configuration resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchTrunkParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**RoutesV1Trunk**](RoutesV1Trunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

