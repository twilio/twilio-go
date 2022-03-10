# ServicesSessionsInteractionsApi

All URIs are relative to *https://proxy.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteInteraction**](ServicesSessionsInteractionsApi.md#DeleteInteraction) | **Delete** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Interactions/{Sid} | 
[**FetchInteraction**](ServicesSessionsInteractionsApi.md#FetchInteraction) | **Get** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Interactions/{Sid} | 
[**ListInteraction**](ServicesSessionsInteractionsApi.md#ListInteraction) | **Get** /v1/Services/{ServiceSid}/Sessions/{SessionSid}/Interactions | 



## DeleteInteraction

> DeleteInteraction(ctx, ServiceSidSessionSidSid)



Delete a specific Interaction.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to delete.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to delete.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Interaction resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteInteractionParams struct


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


## FetchInteraction

> ProxyV1Interaction FetchInteraction(ctx, ServiceSidSessionSidSid)



Retrieve a list of Interactions for a given [Session](https://www.twilio.com/docs/proxy/api/session).

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) of the resource to fetch.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) of the resource to fetch.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Interaction resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchInteractionParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ProxyV1Interaction**](ProxyV1Interaction.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInteraction

> []ProxyV1Interaction ListInteraction(ctx, ServiceSidSessionSidoptional)



Retrieve a list of all Interactions for a Session. A maximum of 100 records will be returned per page.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the parent [Service](https://www.twilio.com/docs/proxy/api/service) to read the resources from.
**SessionSid** | **string** | The SID of the parent [Session](https://www.twilio.com/docs/proxy/api/session) to read the resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListInteractionParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ProxyV1Interaction**](ProxyV1Interaction.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

