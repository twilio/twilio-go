# \GAApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2ServicesGet**](GAApi.md#V2ServicesGet) | **Get** /v2/Services | 
[**V2ServicesPost**](GAApi.md#V2ServicesPost) | **Post** /v2/Services | 
[**V2ServicesServiceSidMessagingConfigurationsCountryDelete**](GAApi.md#V2ServicesServiceSidMessagingConfigurationsCountryDelete) | **Delete** /v2/Services/{ServiceSid}/MessagingConfigurations/{Country} | 
[**V2ServicesServiceSidMessagingConfigurationsCountryGet**](GAApi.md#V2ServicesServiceSidMessagingConfigurationsCountryGet) | **Get** /v2/Services/{ServiceSid}/MessagingConfigurations/{Country} | 
[**V2ServicesServiceSidMessagingConfigurationsCountryPost**](GAApi.md#V2ServicesServiceSidMessagingConfigurationsCountryPost) | **Post** /v2/Services/{ServiceSid}/MessagingConfigurations/{Country} | 
[**V2ServicesServiceSidMessagingConfigurationsGet**](GAApi.md#V2ServicesServiceSidMessagingConfigurationsGet) | **Get** /v2/Services/{ServiceSid}/MessagingConfigurations | 
[**V2ServicesServiceSidMessagingConfigurationsPost**](GAApi.md#V2ServicesServiceSidMessagingConfigurationsPost) | **Post** /v2/Services/{ServiceSid}/MessagingConfigurations | 
[**V2ServicesServiceSidRateLimitsGet**](GAApi.md#V2ServicesServiceSidRateLimitsGet) | **Get** /v2/Services/{ServiceSid}/RateLimits | 
[**V2ServicesServiceSidRateLimitsPost**](GAApi.md#V2ServicesServiceSidRateLimitsPost) | **Post** /v2/Services/{ServiceSid}/RateLimits | 
[**V2ServicesServiceSidRateLimitsRateLimitSidBucketsGet**](GAApi.md#V2ServicesServiceSidRateLimitsRateLimitSidBucketsGet) | **Get** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets | 
[**V2ServicesServiceSidRateLimitsRateLimitSidBucketsPost**](GAApi.md#V2ServicesServiceSidRateLimitsRateLimitSidBucketsPost) | **Post** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets | 
[**V2ServicesServiceSidRateLimitsRateLimitSidBucketsSidDelete**](GAApi.md#V2ServicesServiceSidRateLimitsRateLimitSidBucketsSidDelete) | **Delete** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets/{Sid} | 
[**V2ServicesServiceSidRateLimitsRateLimitSidBucketsSidGet**](GAApi.md#V2ServicesServiceSidRateLimitsRateLimitSidBucketsSidGet) | **Get** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets/{Sid} | 
[**V2ServicesServiceSidRateLimitsRateLimitSidBucketsSidPost**](GAApi.md#V2ServicesServiceSidRateLimitsRateLimitSidBucketsSidPost) | **Post** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets/{Sid} | 
[**V2ServicesServiceSidRateLimitsSidDelete**](GAApi.md#V2ServicesServiceSidRateLimitsSidDelete) | **Delete** /v2/Services/{ServiceSid}/RateLimits/{Sid} | 
[**V2ServicesServiceSidRateLimitsSidGet**](GAApi.md#V2ServicesServiceSidRateLimitsSidGet) | **Get** /v2/Services/{ServiceSid}/RateLimits/{Sid} | 
[**V2ServicesServiceSidRateLimitsSidPost**](GAApi.md#V2ServicesServiceSidRateLimitsSidPost) | **Post** /v2/Services/{ServiceSid}/RateLimits/{Sid} | 
[**V2ServicesServiceSidVerificationCheckPost**](GAApi.md#V2ServicesServiceSidVerificationCheckPost) | **Post** /v2/Services/{ServiceSid}/VerificationCheck | 
[**V2ServicesServiceSidVerificationsPost**](GAApi.md#V2ServicesServiceSidVerificationsPost) | **Post** /v2/Services/{ServiceSid}/Verifications | 
[**V2ServicesServiceSidVerificationsSidGet**](GAApi.md#V2ServicesServiceSidVerificationsSidGet) | **Get** /v2/Services/{ServiceSid}/Verifications/{Sid} | 
[**V2ServicesServiceSidVerificationsSidPost**](GAApi.md#V2ServicesServiceSidVerificationsSidPost) | **Post** /v2/Services/{ServiceSid}/Verifications/{Sid} | 
[**V2ServicesSidDelete**](GAApi.md#V2ServicesSidDelete) | **Delete** /v2/Services/{Sid} | 
[**V2ServicesSidGet**](GAApi.md#V2ServicesSidGet) | **Get** /v2/Services/{Sid} | 
[**V2ServicesSidPost**](GAApi.md#V2ServicesSidPost) | **Post** /v2/Services/{Sid} | 



## V2ServicesGet

> InlineResponse200 V2ServicesGet(ctx, optional)



Retrieve a list of all Verification Services for an account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V2ServicesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesPost

> VerifyV2Service V2ServicesPost(ctx, optional)



Create a new Verification Service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V2ServicesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codeLength** | **optional.**| The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive. | 
 **customCodeEnabled** | **optional.**| Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers. | 
 **doNotShareWarningEnabled** | **optional.**| Whether to add a security warning at the end of an SMS verification body. Disabled by default and applies only to SMS. Example SMS body: &#x60;Your AppName verification code is: 1234. Don’t share this code with anyone; our employees will never ask for the code&#x60; | 
 **dtmfInputRequired** | **optional.**| Whether to ask the user to press a number before delivering the verify code in a phone call. | 
 **friendlyName** | **optional.**| A descriptive string that you create to describe the verification service. It can be up to 64 characters long. **This value should not contain PII.** | 
 **lookupEnabled** | **optional.**| Whether to perform a lookup with each verification started and return info about the phone number. | 
 **psd2Enabled** | **optional.**| Whether to pass PSD2 transaction parameters when starting a verification. | 
 **push** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| Configurations for the Push factors (channel) created under this Service. If present, it must be a json string with the following format: {\\\&quot;notify_service_sid\\\&quot;: \\\&quot;ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\\\&quot;, \\\&quot;include_date\\\&quot;: true}. If &#x60;include_date&#x60; is set to &#x60;true&#x60;, which is the default, that means that the push challenge’s response will include the date created value. If &#x60;include_date&#x60; is set to &#x60;false&#x60;, then the date created value will not be included. See [Challenge](https://www.twilio.com/docs/verify/api/challenge) resource’s details parameter for more info | 
 **skipSmsToLandlines** | **optional.**| Whether to skip sending SMS verifications to landlines. Requires &#x60;lookup_enabled&#x60;. | 
 **ttsName** | **optional.**| The name of an alternative text-to-speech service to use in phone calls. Applies only to TTS languages. | 

### Return type

[**VerifyV2Service**](verify.v2.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidMessagingConfigurationsCountryDelete

> V2ServicesServiceSidMessagingConfigurationsCountryDelete(ctx, serviceSid, country)



Delete a specific MessagingConfiguration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with. | 
**country** | **string**| The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value &#x60;all&#x60;. | 

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


## V2ServicesServiceSidMessagingConfigurationsCountryGet

> VerifyV2ServiceMessagingConfiguration V2ServicesServiceSidMessagingConfigurationsCountryGet(ctx, serviceSid, country)



Fetch a specific MessagingConfiguration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with. | 
**country** | **string**| The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value &#x60;all&#x60;. | 

### Return type

[**VerifyV2ServiceMessagingConfiguration**](verify.v2.service.messaging_configuration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidMessagingConfigurationsCountryPost

> VerifyV2ServiceMessagingConfiguration V2ServicesServiceSidMessagingConfigurationsCountryPost(ctx, serviceSid, country, optional)



Update a specific MessagingConfiguration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with. | 
**country** | **string**| The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value &#x60;all&#x60;. | 
 **optional** | ***V2ServicesServiceSidMessagingConfigurationsCountryPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidMessagingConfigurationsCountryPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **messagingServiceSid** | **optional.**| The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) to be used to send SMS to the country of this configuration. | 

### Return type

[**VerifyV2ServiceMessagingConfiguration**](verify.v2.service.messaging_configuration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidMessagingConfigurationsGet

> InlineResponse2004 V2ServicesServiceSidMessagingConfigurationsGet(ctx, serviceSid, optional)



Retrieve a list of all Messaging Configurations for a Service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with. | 
 **optional** | ***V2ServicesServiceSidMessagingConfigurationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidMessagingConfigurationsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidMessagingConfigurationsPost

> VerifyV2ServiceMessagingConfiguration V2ServicesServiceSidMessagingConfigurationsPost(ctx, serviceSid, optional)



Create a new MessagingConfiguration for a service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with. | 
 **optional** | ***V2ServicesServiceSidMessagingConfigurationsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidMessagingConfigurationsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **country** | **optional.**| The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value &#x60;all&#x60;. | 
 **messagingServiceSid** | **optional.**| The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) to be used to send SMS to the country of this configuration. | 

### Return type

[**VerifyV2ServiceMessagingConfiguration**](verify.v2.service.messaging_configuration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidRateLimitsGet

> InlineResponse2005 V2ServicesServiceSidRateLimitsGet(ctx, serviceSid, optional)



Retrieve a list of all Rate Limits for a service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
 **optional** | ***V2ServicesServiceSidRateLimitsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidRateLimitsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidRateLimitsPost

> VerifyV2ServiceRateLimit V2ServicesServiceSidRateLimitsPost(ctx, serviceSid, optional)



Create a new Rate Limit for a Service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
 **optional** | ***V2ServicesServiceSidRateLimitsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidRateLimitsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **description** | **optional.**| Description of this Rate Limit | 
 **uniqueName** | **optional.**| Provides a unique and addressable name to be assigned to this Rate Limit, assigned by the developer, to be optionally used in addition to SID. **This value should not contain PII.** | 

### Return type

[**VerifyV2ServiceRateLimit**](verify.v2.service.rate_limit.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidRateLimitsRateLimitSidBucketsGet

> InlineResponse2006 V2ServicesServiceSidRateLimitsRateLimitSidBucketsGet(ctx, serviceSid, rateLimitSid, optional)



Retrieve a list of all Buckets for a Rate Limit.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
**rateLimitSid** | **string**| The Twilio-provided string that uniquely identifies the Rate Limit resource. | 
 **optional** | ***V2ServicesServiceSidRateLimitsRateLimitSidBucketsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidRateLimitsRateLimitSidBucketsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidRateLimitsRateLimitSidBucketsPost

> VerifyV2ServiceRateLimitBucket V2ServicesServiceSidRateLimitsRateLimitSidBucketsPost(ctx, serviceSid, rateLimitSid, optional)



Create a new Bucket for a Rate Limit

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
**rateLimitSid** | **string**| The Twilio-provided string that uniquely identifies the Rate Limit resource. | 
 **optional** | ***V2ServicesServiceSidRateLimitsRateLimitSidBucketsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidRateLimitsRateLimitSidBucketsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **interval** | **optional.**| Number of seconds that the rate limit will be enforced over. | 
 **max** | **optional.**| Maximum number of requests permitted in during the interval. | 

### Return type

[**VerifyV2ServiceRateLimitBucket**](verify.v2.service.rate_limit.bucket.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidRateLimitsRateLimitSidBucketsSidDelete

> V2ServicesServiceSidRateLimitsRateLimitSidBucketsSidDelete(ctx, serviceSid, rateLimitSid, sid)



Delete a specific Bucket.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
**rateLimitSid** | **string**| The Twilio-provided string that uniquely identifies the Rate Limit resource. | 
**sid** | **string**| A 34 character string that uniquely identifies this Bucket. | 

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


## V2ServicesServiceSidRateLimitsRateLimitSidBucketsSidGet

> VerifyV2ServiceRateLimitBucket V2ServicesServiceSidRateLimitsRateLimitSidBucketsSidGet(ctx, serviceSid, rateLimitSid, sid)



Fetch a specific Bucket.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
**rateLimitSid** | **string**| The Twilio-provided string that uniquely identifies the Rate Limit resource. | 
**sid** | **string**| A 34 character string that uniquely identifies this Bucket. | 

### Return type

[**VerifyV2ServiceRateLimitBucket**](verify.v2.service.rate_limit.bucket.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidRateLimitsRateLimitSidBucketsSidPost

> VerifyV2ServiceRateLimitBucket V2ServicesServiceSidRateLimitsRateLimitSidBucketsSidPost(ctx, serviceSid, rateLimitSid, sid, optional)



Update a specific Bucket.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
**rateLimitSid** | **string**| The Twilio-provided string that uniquely identifies the Rate Limit resource. | 
**sid** | **string**| A 34 character string that uniquely identifies this Bucket. | 
 **optional** | ***V2ServicesServiceSidRateLimitsRateLimitSidBucketsSidPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidRateLimitsRateLimitSidBucketsSidPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **interval** | **optional.**| Number of seconds that the rate limit will be enforced over. | 
 **max** | **optional.**| Maximum number of requests permitted in during the interval. | 

### Return type

[**VerifyV2ServiceRateLimitBucket**](verify.v2.service.rate_limit.bucket.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidRateLimitsSidDelete

> V2ServicesServiceSidRateLimitsSidDelete(ctx, serviceSid, sid)



Delete a specific Rate Limit.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the Rate Limit resource to fetch. | 

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


## V2ServicesServiceSidRateLimitsSidGet

> VerifyV2ServiceRateLimit V2ServicesServiceSidRateLimitsSidGet(ctx, serviceSid, sid)



Fetch a specific Rate Limit.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the Rate Limit resource to fetch. | 

### Return type

[**VerifyV2ServiceRateLimit**](verify.v2.service.rate_limit.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidRateLimitsSidPost

> VerifyV2ServiceRateLimit V2ServicesServiceSidRateLimitsSidPost(ctx, serviceSid, sid, optional)



Update a specific Rate Limit.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the Rate Limit resource to fetch. | 
 **optional** | ***V2ServicesServiceSidRateLimitsSidPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidRateLimitsSidPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **description** | **optional.**| Description of this Rate Limit | 

### Return type

[**VerifyV2ServiceRateLimit**](verify.v2.service.rate_limit.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidVerificationCheckPost

> VerifyV2ServiceVerificationCheck V2ServicesServiceSidVerificationCheckPost(ctx, serviceSid, optional)



challenge a specific Verification Check.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to create the resource under. | 
 **optional** | ***V2ServicesServiceSidVerificationCheckPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidVerificationCheckPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amount** | **optional.**| The amount of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled. | 
 **code** | **optional.**| The 4-10 character string being verified. | 
 **payee** | **optional.**| The payee of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled. | 
 **to** | **optional.**| The phone number or [email](https://www.twilio.com/docs/verify/email) to verify. Either this parameter or the &#x60;verification_sid&#x60; must be specified. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164). | 
 **verificationSid** | **optional.**| A SID that uniquely identifies the Verification Check. Either this parameter or the &#x60;to&#x60; phone number/[email](https://www.twilio.com/docs/verify/email) must be specified. | 

### Return type

[**VerifyV2ServiceVerificationCheck**](verify.v2.service.verification_check.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidVerificationsPost

> VerifyV2ServiceVerification V2ServicesServiceSidVerificationsPost(ctx, serviceSid, optional)



Create a new Verification using a Service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to create the resource under. | 
 **optional** | ***V2ServicesServiceSidVerificationsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidVerificationsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amount** | **optional.**| The amount of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled. | 
 **appHash** | **optional.**| Your [App Hash](https://developers.google.com/identity/sms-retriever/verify#computing_your_apps_hash_string) to be appended at the end of your verification SMS body. Applies only to SMS. Example SMS body: &#x60;&lt;#&gt; Your AppName verification code is: 1234 He42w354ol9&#x60;. | 
 **channel** | **optional.**| The verification method to use. Can be: [&#x60;email&#x60;](https://www.twilio.com/docs/verify/email), &#x60;sms&#x60; or &#x60;call&#x60;. | 
 **channelConfiguration** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| [&#x60;email&#x60;](https://www.twilio.com/docs/verify/email) channel configuration in json format. Must include &#39;from&#39; and &#39;from_name&#39;. | 
 **customCode** | **optional.**| A pre-generated code to use for verification. The code can be between 4 and 10 characters, inclusive. | 
 **customFriendlyName** | **optional.**| A custom user defined friendly name that overwrites the existing one in the verification message | 
 **customMessage** | **optional.**| The text of a custom message to use for the verification. | 
 **locale** | **optional.**| The locale to use for the verification SMS or call. Can be: &#x60;af&#x60;, &#x60;ar&#x60;, &#x60;ca&#x60;, &#x60;cs&#x60;, &#x60;da&#x60;, &#x60;de&#x60;, &#x60;el&#x60;, &#x60;en&#x60;, &#x60;en-GB&#x60;, &#x60;es&#x60;, &#x60;fi&#x60;, &#x60;fr&#x60;, &#x60;he&#x60;, &#x60;hi&#x60;, &#x60;hr&#x60;, &#x60;hu&#x60;, &#x60;id&#x60;, &#x60;it&#x60;, &#x60;ja&#x60;, &#x60;ko&#x60;, &#x60;ms&#x60;, &#x60;nb&#x60;, &#x60;nl&#x60;, &#x60;pl&#x60;, &#x60;pt&#x60;, &#x60;pr-BR&#x60;, &#x60;ro&#x60;, &#x60;ru&#x60;, &#x60;sv&#x60;, &#x60;th&#x60;, &#x60;tl&#x60;, &#x60;tr&#x60;, &#x60;vi&#x60;, &#x60;zh&#x60;, &#x60;zh-CN&#x60;, or &#x60;zh-HK.&#x60; | 
 **payee** | **optional.**| The payee of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled. | 
 **rateLimits** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The custom key-value pairs of Programmable Rate Limits. Keys correspond to &#x60;unique_name&#x60; fields defined when [creating your Rate Limit](https://www.twilio.com/docs/verify/api/service-rate-limits). Associated value pairs represent values in the request that you are rate limiting on. You may include multiple Rate Limit values in each request. | 
 **sendDigits** | **optional.**| The digits to send after a phone call is answered, for example, to dial an extension. For more information, see the Programmable Voice documentation of [sendDigits](https://www.twilio.com/docs/voice/twiml/number#attributes-sendDigits). | 
 **to** | **optional.**| The phone number or [email](https://www.twilio.com/docs/verify/email) to verify. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164). | 

### Return type

[**VerifyV2ServiceVerification**](verify.v2.service.verification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidVerificationsSidGet

> VerifyV2ServiceVerification V2ServicesServiceSidVerificationsSidGet(ctx, serviceSid, sid)



Fetch a specific Verification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to fetch the resource from. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the Verification resource to fetch. | 

### Return type

[**VerifyV2ServiceVerification**](verify.v2.service.verification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidVerificationsSidPost

> VerifyV2ServiceVerification V2ServicesServiceSidVerificationsSidPost(ctx, serviceSid, sid, optional)



Update a Verification status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to update the resource from. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the Verification resource to update. | 
 **optional** | ***V2ServicesServiceSidVerificationsSidPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidVerificationsSidPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **optional.**| The new status of the resource. Can be: &#x60;canceled&#x60; or &#x60;approved&#x60;. | 

### Return type

[**VerifyV2ServiceVerification**](verify.v2.service.verification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesSidDelete

> V2ServicesSidDelete(ctx, sid)



Delete a specific Verification Service Instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the Verification Service resource to delete. | 

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


## V2ServicesSidGet

> VerifyV2Service V2ServicesSidGet(ctx, sid)



Fetch specific Verification Service Instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the Verification Service resource to fetch. | 

### Return type

[**VerifyV2Service**](verify.v2.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesSidPost

> VerifyV2Service V2ServicesSidPost(ctx, sid, optional)



Update a specific Verification Service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the Service resource to update. | 
 **optional** | ***V2ServicesSidPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesSidPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **codeLength** | **optional.**| The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive. | 
 **customCodeEnabled** | **optional.**| Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers. | 
 **doNotShareWarningEnabled** | **optional.**| Whether to add a privacy warning at the end of an SMS. **Disabled by default and applies only for SMS.** | 
 **dtmfInputRequired** | **optional.**| Whether to ask the user to press a number before delivering the verify code in a phone call. | 
 **friendlyName** | **optional.**| A descriptive string that you create to describe the verification service. It can be up to 64 characters long. **This value should not contain PII.** | 
 **lookupEnabled** | **optional.**| Whether to perform a lookup with each verification started and return info about the phone number. | 
 **psd2Enabled** | **optional.**| Whether to pass PSD2 transaction parameters when starting a verification. | 
 **push** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| Configurations for the Push factors (channel) created under this Service. If present, it must be a json string with the following format: {\\\&quot;notify_service_sid\\\&quot;: \\\&quot;ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\\\&quot;, \\\&quot;include_date\\\&quot;: true}. If &#x60;include_date&#x60; is set to &#x60;true&#x60;, which is the default, that means that the push challenge’s response will include the date created value. If &#x60;include_date&#x60; is set to &#x60;false&#x60;, then the date created value will not be included. See [Challenge](https://www.twilio.com/docs/verify/api/challenge) resource’s details parameter for more info | 
 **skipSmsToLandlines** | **optional.**| Whether to skip sending SMS verifications to landlines. Requires &#x60;lookup_enabled&#x60;. | 
 **ttsName** | **optional.**| The name of an alternative text-to-speech service to use in phone calls. Applies only to TTS languages. | 

### Return type

[**VerifyV2Service**](verify.v2.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

