# ConferencesApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchConference**](ConferencesApi.md#FetchConference) | **Get** /v1/Conferences/{ConferenceSid} | 
[**ListConference**](ConferencesApi.md#ListConference) | **Get** /v1/Conferences | 



## FetchConference

> InsightsV1Conference FetchConference(ctx, ConferenceSid)



Fetch a specific Conference.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConferenceSid** | **string** | The unique SID identifier of the Conference.

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

> []InsightsV1Conference ListConference(ctx, optional)



Retrieve a list of Conferences.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListConferenceParams struct


Name | Type | Description
------------- | ------------- | -------------
**ConferenceSid** | **string** | The SID of the conference.
**FriendlyName** | **string** | Custom label for the conference resource, up to 64 characters.
**Status** | **string** | Conference status.
**CreatedAfter** | **string** | Conferences created after the provided timestamp specified in ISO 8601 format
**CreatedBefore** | **string** | Conferences created before the provided timestamp specified in ISO 8601 format.
**MixerRegion** | **string** | Twilio region where the conference media was mixed.
**Tags** | **string** | Tags applied by Twilio for common potential configuration, quality, or performance issues.
**Subaccount** | **string** | Account SID for the subaccount whose resources you wish to retrieve.
**DetectedIssues** | **string** | Potential configuration, behavior, or performance issues detected during the conference.
**EndReason** | **string** | Conference end reason; e.g. last participant left, modified by API, etc.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]InsightsV1Conference**](InsightsV1Conference.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

