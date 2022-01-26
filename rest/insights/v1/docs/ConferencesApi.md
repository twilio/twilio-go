# ConferencesApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchConference**](ConferencesApi.md#FetchConference) | **Get** /v1/Conferences/{ConferenceSid} | 
[**ListConference**](ConferencesApi.md#ListConference) | **Get** /v1/Conferences | 



## FetchConference

> InsightsV1Conference FetchConference(ctx, ConferenceSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConferenceSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchConferenceParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**InsightsV1Conference**](InsightsV1Conference.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConference

> ListConferenceResponse ListConference(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListConferenceParams struct


Name | Type | Description
------------- | ------------- | -------------
**ConferenceSid** | **string** | 
**FriendlyName** | **string** | 
**Status** | **string** | 
**CreatedAfter** | **string** | 
**CreatedBefore** | **string** | 
**MixerRegion** | **string** | 
**Tags** | **string** | 
**Subaccount** | **string** | 
**DetectedIssues** | **string** | 
**EndReason** | **string** | 
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**ListConferenceResponse**](ListConferenceResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

