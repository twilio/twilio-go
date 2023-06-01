# ServicesApi

All URIs are relative to *https://intelligence.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateService**](ServicesApi.md#CreateService) | **Post** /v2/Services | 
[**DeleteService**](ServicesApi.md#DeleteService) | **Delete** /v2/Services/{Sid} | 
[**FetchService**](ServicesApi.md#FetchService) | **Get** /v2/Services/{Sid} | 
[**ListService**](ServicesApi.md#ListService) | **Get** /v2/Services | 
[**UpdateService**](ServicesApi.md#UpdateService) | **Post** /v2/Services/{Sid} | 



## CreateService

> IntelligenceV2Service CreateService(ctx, optional)



Create a new Service for the given Account

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**UniqueName** | **string** | Provides a unique and addressable name to be assigned to this Service, assigned by the developer, to be optionally used in addition to SID.
**AutoTranscribe** | **bool** | Instructs the Speech Recognition service to automatically transcribe all recordings made on the account.
**DataLogging** | **bool** | Data logging allows Twilio to improve the quality of the speech recognition through using customer data to refine its speech recognition models.
**FriendlyName** | **string** | A human readable description of this resource, up to 64 characters.
**LanguageCode** | **string** | The default language code of the audio.
**AutoRedaction** | **bool** | Instructs the Speech Recognition service to automatically redact PII from all transcripts made on this service.
**MediaRedaction** | **bool** | Instructs the Speech Recognition service to automatically redact PII from all transcripts media made on this service. The auto_redaction flag must be enabled, results in error otherwise.
**WebhookUrl** | **string** | The URL Twilio will request when executing the Webhook.
**WebhookHttpMethod** | **string** | 

### Return type

[**IntelligenceV2Service**](IntelligenceV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteService

> DeleteService(ctx, Sid)



Delete a specific Service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Service.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceParams struct


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


## FetchService

> IntelligenceV2Service FetchService(ctx, Sid)



Fetch a specific Service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Service.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IntelligenceV2Service**](IntelligenceV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> []IntelligenceV2Service ListService(ctx, optional)



Retrieves a list of all Services for an account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]IntelligenceV2Service**](IntelligenceV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> IntelligenceV2Service UpdateService(ctx, Sidoptional)



Update a specific Service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this Service.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceParams struct


Name | Type | Description
------------- | ------------- | -------------
**IfMatch** | **string** | The If-Match HTTP request header
**AutoTranscribe** | **bool** | Instructs the Speech Recognition service to automatically transcribe all recordings made on the account.
**DataLogging** | **bool** | Data logging allows Twilio to improve the quality of the speech recognition through using customer data to refine its speech recognition models.
**FriendlyName** | **string** | A human readable description of this resource, up to 64 characters.
**LanguageCode** | **string** | The default language code of the audio.
**UniqueName** | **string** | Provides a unique and addressable name to be assigned to this Service, assigned by the developer, to be optionally used in addition to SID.
**AutoRedaction** | **bool** | Instructs the Speech Recognition service to automatically redact PII from all transcripts made on this service.
**MediaRedaction** | **bool** | Instructs the Speech Recognition service to automatically redact PII from all transcripts media made on this service. The auto_redaction flag must be enabled, results in error otherwise.
**WebhookUrl** | **string** | The URL Twilio will request when executing the Webhook.
**WebhookHttpMethod** | **string** | 

### Return type

[**IntelligenceV2Service**](IntelligenceV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

