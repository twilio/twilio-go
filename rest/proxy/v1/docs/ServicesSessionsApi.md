# ServicesSessionsApi

All URIs are relative to *https://proxy.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSession**](ServicesSessionsApi.md#CreateSession) | **Post** /v1/Services/{ServiceSid}/Sessions | 
[**DeleteSession**](ServicesSessionsApi.md#DeleteSession) | **Delete** /v1/Services/{ServiceSid}/Sessions/{Sid} | 
[**FetchSession**](ServicesSessionsApi.md#FetchSession) | **Get** /v1/Services/{ServiceSid}/Sessions/{Sid} | 
[**ListSession**](ServicesSessionsApi.md#ListSession) | **Get** /v1/Services/{ServiceSid}/Sessions | 
[**UpdateSession**](ServicesSessionsApi.md#UpdateSession) | **Post** /v1/Services/{ServiceSid}/Sessions/{Sid} | 



## CreateSession

> ProxyV1Session CreateSession(ctx, ServiceSidoptional)



Create a new Session

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateSessionParams struct


Name | Type | Description
------------- | ------------- | -------------
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.**
**DateExpiry** | **time.Time** | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the `ttl` value.
**Ttl** | **int** | The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session's last Interaction.
**Mode** | **string** | 
**Status** | **string** | 
**Participants** | **[]interface{}** | The Participant objects to include in the new session.

### Return type

[**ProxyV1Session**](ProxyV1Session.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSession

> DeleteSession(ctx, ServiceSidSid)



Delete a specific Session.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Session resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSessionParams struct


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


## FetchSession

> ProxyV1Session FetchSession(ctx, ServiceSidSid)



Fetch a specific Session.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Session resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSessionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ProxyV1Session**](ProxyV1Session.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSession

> []ProxyV1Session ListSession(ctx, ServiceSidoptional)



Retrieve a list of all Sessions for the Service. A maximum of 100 records will be returned per page.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to read.

### Other Parameters

Other parameters are passed through a pointer to a ListSessionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ProxyV1Session**](ProxyV1Session.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSession

> ProxyV1Session UpdateSession(ctx, ServiceSidSidoptional)



Update a specific Session.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to update.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Session resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSessionParams struct


Name | Type | Description
------------- | ------------- | -------------
**DateExpiry** | **time.Time** | The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the `ttl` value.
**Ttl** | **int** | The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session's last Interaction.
**Status** | **string** | 

### Return type

[**ProxyV1Session**](ProxyV1Session.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

