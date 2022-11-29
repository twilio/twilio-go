# ServicesMessagingConfigurationsApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMessagingConfiguration**](ServicesMessagingConfigurationsApi.md#CreateMessagingConfiguration) | **Post** /v2/Services/{ServiceSid}/MessagingConfigurations | 
[**DeleteMessagingConfiguration**](ServicesMessagingConfigurationsApi.md#DeleteMessagingConfiguration) | **Delete** /v2/Services/{ServiceSid}/MessagingConfigurations/{Country} | 
[**FetchMessagingConfiguration**](ServicesMessagingConfigurationsApi.md#FetchMessagingConfiguration) | **Get** /v2/Services/{ServiceSid}/MessagingConfigurations/{Country} | 
[**ListMessagingConfiguration**](ServicesMessagingConfigurationsApi.md#ListMessagingConfiguration) | **Get** /v2/Services/{ServiceSid}/MessagingConfigurations | 
[**UpdateMessagingConfiguration**](ServicesMessagingConfigurationsApi.md#UpdateMessagingConfiguration) | **Post** /v2/Services/{ServiceSid}/MessagingConfigurations/{Country} | 



## CreateMessagingConfiguration

> VerifyV2MessagingConfiguration CreateMessagingConfiguration(ctx, ServiceSidoptional)



Create a new MessagingConfiguration for a service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with.

### Other Parameters

Other parameters are passed through a pointer to a CreateMessagingConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**Country** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value `all`.
**MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) to be used to send SMS to the country of this configuration.

### Return type

[**VerifyV2MessagingConfiguration**](VerifyV2MessagingConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMessagingConfiguration

> DeleteMessagingConfiguration(ctx, ServiceSidCountry)



Delete a specific MessagingConfiguration.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with.
**Country** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value `all`.

### Other Parameters

Other parameters are passed through a pointer to a DeleteMessagingConfigurationParams struct


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


## FetchMessagingConfiguration

> VerifyV2MessagingConfiguration FetchMessagingConfiguration(ctx, ServiceSidCountry)



Fetch a specific MessagingConfiguration.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with.
**Country** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value `all`.

### Other Parameters

Other parameters are passed through a pointer to a FetchMessagingConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VerifyV2MessagingConfiguration**](VerifyV2MessagingConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessagingConfiguration

> []VerifyV2MessagingConfiguration ListMessagingConfiguration(ctx, ServiceSidoptional)



Retrieve a list of all Messaging Configurations for a Service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with.

### Other Parameters

Other parameters are passed through a pointer to a ListMessagingConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VerifyV2MessagingConfiguration**](VerifyV2MessagingConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessagingConfiguration

> VerifyV2MessagingConfiguration UpdateMessagingConfiguration(ctx, ServiceSidCountryoptional)



Update a specific MessagingConfiguration

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with.
**Country** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value `all`.

### Other Parameters

Other parameters are passed through a pointer to a UpdateMessagingConfigurationParams struct


Name | Type | Description
------------- | ------------- | -------------
**MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) to be used to send SMS to the country of this configuration.

### Return type

[**VerifyV2MessagingConfiguration**](VerifyV2MessagingConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

