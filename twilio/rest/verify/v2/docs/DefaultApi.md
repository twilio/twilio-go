# DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessToken**](DefaultApi.md#CreateAccessToken) | **Post** /v2/Services/{ServiceSid}/AccessTokens | 
[**CreateBucket**](DefaultApi.md#CreateBucket) | **Post** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets | 
[**CreateChallenge**](DefaultApi.md#CreateChallenge) | **Post** /v2/Services/{ServiceSid}/Entities/{Identity}/Challenges | 
[**CreateEntity**](DefaultApi.md#CreateEntity) | **Post** /v2/Services/{ServiceSid}/Entities | 
[**CreateFactor**](DefaultApi.md#CreateFactor) | **Post** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors | 
[**CreateMessagingConfiguration**](DefaultApi.md#CreateMessagingConfiguration) | **Post** /v2/Services/{ServiceSid}/MessagingConfigurations | 
[**CreateRateLimit**](DefaultApi.md#CreateRateLimit) | **Post** /v2/Services/{ServiceSid}/RateLimits | 
[**CreateService**](DefaultApi.md#CreateService) | **Post** /v2/Services | 
[**CreateVerification**](DefaultApi.md#CreateVerification) | **Post** /v2/Services/{ServiceSid}/Verifications | 
[**CreateVerificationCheck**](DefaultApi.md#CreateVerificationCheck) | **Post** /v2/Services/{ServiceSid}/VerificationCheck | 
[**CreateWebhook**](DefaultApi.md#CreateWebhook) | **Post** /v2/Services/{ServiceSid}/Webhooks | 
[**DeleteBucket**](DefaultApi.md#DeleteBucket) | **Delete** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets/{Sid} | 
[**DeleteEntity**](DefaultApi.md#DeleteEntity) | **Delete** /v2/Services/{ServiceSid}/Entities/{Identity} | 
[**DeleteFactor**](DefaultApi.md#DeleteFactor) | **Delete** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors/{Sid} | 
[**DeleteMessagingConfiguration**](DefaultApi.md#DeleteMessagingConfiguration) | **Delete** /v2/Services/{ServiceSid}/MessagingConfigurations/{Country} | 
[**DeleteRateLimit**](DefaultApi.md#DeleteRateLimit) | **Delete** /v2/Services/{ServiceSid}/RateLimits/{Sid} | 
[**DeleteService**](DefaultApi.md#DeleteService) | **Delete** /v2/Services/{Sid} | 
[**DeleteWebhook**](DefaultApi.md#DeleteWebhook) | **Delete** /v2/Services/{ServiceSid}/Webhooks/{Sid} | 
[**FetchBucket**](DefaultApi.md#FetchBucket) | **Get** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets/{Sid} | 
[**FetchChallenge**](DefaultApi.md#FetchChallenge) | **Get** /v2/Services/{ServiceSid}/Entities/{Identity}/Challenges/{Sid} | 
[**FetchEntity**](DefaultApi.md#FetchEntity) | **Get** /v2/Services/{ServiceSid}/Entities/{Identity} | 
[**FetchFactor**](DefaultApi.md#FetchFactor) | **Get** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors/{Sid} | 
[**FetchForm**](DefaultApi.md#FetchForm) | **Get** /v2/Forms/{FormType} | 
[**FetchMessagingConfiguration**](DefaultApi.md#FetchMessagingConfiguration) | **Get** /v2/Services/{ServiceSid}/MessagingConfigurations/{Country} | 
[**FetchRateLimit**](DefaultApi.md#FetchRateLimit) | **Get** /v2/Services/{ServiceSid}/RateLimits/{Sid} | 
[**FetchService**](DefaultApi.md#FetchService) | **Get** /v2/Services/{Sid} | 
[**FetchVerification**](DefaultApi.md#FetchVerification) | **Get** /v2/Services/{ServiceSid}/Verifications/{Sid} | 
[**FetchVerificationAttempt**](DefaultApi.md#FetchVerificationAttempt) | **Get** /v2/Attempts/{Sid} | 
[**FetchWebhook**](DefaultApi.md#FetchWebhook) | **Get** /v2/Services/{ServiceSid}/Webhooks/{Sid} | 
[**ListBucket**](DefaultApi.md#ListBucket) | **Get** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets | 
[**ListChallenge**](DefaultApi.md#ListChallenge) | **Get** /v2/Services/{ServiceSid}/Entities/{Identity}/Challenges | 
[**ListEntity**](DefaultApi.md#ListEntity) | **Get** /v2/Services/{ServiceSid}/Entities | 
[**ListFactor**](DefaultApi.md#ListFactor) | **Get** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors | 
[**ListMessagingConfiguration**](DefaultApi.md#ListMessagingConfiguration) | **Get** /v2/Services/{ServiceSid}/MessagingConfigurations | 
[**ListRateLimit**](DefaultApi.md#ListRateLimit) | **Get** /v2/Services/{ServiceSid}/RateLimits | 
[**ListService**](DefaultApi.md#ListService) | **Get** /v2/Services | 
[**ListVerificationAttempt**](DefaultApi.md#ListVerificationAttempt) | **Get** /v2/Attempts | 
[**ListWebhook**](DefaultApi.md#ListWebhook) | **Get** /v2/Services/{ServiceSid}/Webhooks | 
[**UpdateBucket**](DefaultApi.md#UpdateBucket) | **Post** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets/{Sid} | 
[**UpdateChallenge**](DefaultApi.md#UpdateChallenge) | **Post** /v2/Services/{ServiceSid}/Entities/{Identity}/Challenges/{Sid} | 
[**UpdateFactor**](DefaultApi.md#UpdateFactor) | **Post** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors/{Sid} | 
[**UpdateMessagingConfiguration**](DefaultApi.md#UpdateMessagingConfiguration) | **Post** /v2/Services/{ServiceSid}/MessagingConfigurations/{Country} | 
[**UpdateRateLimit**](DefaultApi.md#UpdateRateLimit) | **Post** /v2/Services/{ServiceSid}/RateLimits/{Sid} | 
[**UpdateService**](DefaultApi.md#UpdateService) | **Post** /v2/Services/{Sid} | 
[**UpdateVerification**](DefaultApi.md#UpdateVerification) | **Post** /v2/Services/{ServiceSid}/Verifications/{Sid} | 
[**UpdateWebhook**](DefaultApi.md#UpdateWebhook) | **Post** /v2/Services/{ServiceSid}/Webhooks/{Sid} | 



## CreateAccessToken

> VerifyV2ServiceAccessToken CreateAccessToken(ctx, ServiceSid, optional)



Create a new enrollment Access Token for the Entity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Service. | 
 **optional** | ***CreateAccessTokenRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAccessTokenRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FactorType** | **String**| The Type of this Factor. Eg. &#x60;push&#x60; | 
**Identity** | **String**| The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, and generated by your external system, such as your user&#39;s UUID, GUID, or SID. | 

### Return type

[**VerifyV2ServiceAccessToken**](VerifyV2ServiceAccessToken.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBucket

> VerifyV2ServiceRateLimitBucket CreateBucket(ctx, ServiceSid, RateLimitSid, optional)



Create a new Bucket for a Rate Limit

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
**RateLimitSid** | **string**| The Twilio-provided string that uniquely identifies the Rate Limit resource. | 
 **optional** | ***CreateBucketRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateBucketRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Interval** | **Int32**| Number of seconds that the rate limit will be enforced over. | 
**Max** | **Int32**| Maximum number of requests permitted in during the interval. | 

### Return type

[**VerifyV2ServiceRateLimitBucket**](VerifyV2ServiceRateLimitBucket.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateChallenge

> VerifyV2ServiceEntityChallenge CreateChallenge(ctx, ServiceSid, Identity, optional)



Create a new Challenge for the Factor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Service. | 
**Identity** | **string**| The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, and generated by your external system, such as your user&#39;s UUID, GUID, or SID. | 
 **optional** | ***CreateChallengeRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateChallengeRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**DetailsFields** | [**[]map[string]interface{}**](map[string]interface{}.md)| A list of objects that describe the Fields included in the Challenge. Each object contains the label and value of the field. Used when &#x60;factor_type&#x60; is &#x60;push&#x60;. | 
**DetailsMessage** | **String**| Shown to the user when the push notification arrives. Required when &#x60;factor_type&#x60; is &#x60;push&#x60; | 
**ExpirationDate** | **Time**| The date-time when this Challenge expires, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. The default value is five (5) minutes after Challenge creation. The max value is sixty (60) minutes after creation. | 
**FactorSid** | **String**| The unique SID identifier of the Factor. | 
**HiddenDetails** | [**map[string]interface{}**](map[string]interface{}.md)| Details provided to give context about the Challenge. Not shown to the end user. It must be a stringified JSON with only strings values eg. &#x60;{\\\&quot;ip\\\&quot;: \\\&quot;172.168.1.234\\\&quot;}&#x60; | 

### Return type

[**VerifyV2ServiceEntityChallenge**](VerifyV2ServiceEntityChallenge.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEntity

> VerifyV2ServiceEntity CreateEntity(ctx, ServiceSid, optional)



Create a new Entity for the Service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Service. | 
 **optional** | ***CreateEntityRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateEntityRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Identity** | **String**| The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, length between 8 and 64 characters, and generated by your external system, such as your user&#39;s UUID, GUID, or SID. | 

### Return type

[**VerifyV2ServiceEntity**](VerifyV2ServiceEntity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFactor

> VerifyV2ServiceEntityFactor CreateFactor(ctx, ServiceSid, Identity, optional)



Create a new Factor for the Entity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Service. | 
**Identity** | **string**| Customer unique identity for the Entity owner of the Factor | 
 **optional** | ***CreateFactorRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFactorRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**BindingAlg** | **String**| The algorithm used when &#x60;factor_type&#x60; is &#x60;push&#x60;. Algorithm supported: &#x60;ES256&#x60; | 
**BindingPublicKey** | **String**| The Ecdsa public key in PKIX, ASN.1 DER format encoded in Base64 | 
**ConfigAppId** | **String**| The ID that uniquely identifies your app in the Google or Apple store, such as &#x60;com.example.myapp&#x60;. Required when &#x60;factor_type&#x60; is &#x60;push&#x60; | 
**ConfigNotificationPlatform** | **String**| The transport technology used to generate the Notification Token. Can be &#x60;apn&#x60; or &#x60;fcm&#x60;. Required when &#x60;factor_type&#x60; is &#x60;push&#x60; | 
**ConfigNotificationToken** | **String**| For APN, the device token. For FCM the registration token. It used to send the push notifications. Required when &#x60;factor_type&#x60; is &#x60;push&#x60; | 
**ConfigSdkVersion** | **String**| The Verify Push SDK version used to configure the factor | 
**FactorType** | **String**| The Type of this Factor. Currently only &#x60;push&#x60; is supported | 
**FriendlyName** | **String**| The friendly name of this Factor | 

### Return type

[**VerifyV2ServiceEntityFactor**](VerifyV2ServiceEntityFactor.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMessagingConfiguration

> VerifyV2ServiceMessagingConfiguration CreateMessagingConfiguration(ctx, ServiceSid, optional)



Create a new MessagingConfiguration for a service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with. | 
 **optional** | ***CreateMessagingConfigurationRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateMessagingConfigurationRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Country** | **String**| The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value &#x60;all&#x60;. | 
**MessagingServiceSid** | **String**| The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) to be used to send SMS to the country of this configuration. | 

### Return type

[**VerifyV2ServiceMessagingConfiguration**](VerifyV2ServiceMessagingConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRateLimit

> VerifyV2ServiceRateLimit CreateRateLimit(ctx, ServiceSid, optional)



Create a new Rate Limit for a Service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
 **optional** | ***CreateRateLimitRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateRateLimitRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Description** | **String**| Description of this Rate Limit | 
**UniqueName** | **String**| Provides a unique and addressable name to be assigned to this Rate Limit, assigned by the developer, to be optionally used in addition to SID. **This value should not contain PII.** | 

### Return type

[**VerifyV2ServiceRateLimit**](VerifyV2ServiceRateLimit.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> VerifyV2Service CreateService(ctx, optional)



Create a new Verification Service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateServiceRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateServiceRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CodeLength** | **Int32**| The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive. | 
**CustomCodeEnabled** | **Bool**| Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers. | 
**DoNotShareWarningEnabled** | **Bool**| Whether to add a security warning at the end of an SMS verification body. Disabled by default and applies only to SMS. Example SMS body: &#x60;Your AppName verification code is: 1234. Don’t share this code with anyone; our employees will never ask for the code&#x60; | 
**DtmfInputRequired** | **Bool**| Whether to ask the user to press a number before delivering the verify code in a phone call. | 
**FriendlyName** | **String**| A descriptive string that you create to describe the verification service. It can be up to 30 characters long. **This value should not contain PII.** | 
**LookupEnabled** | **Bool**| Whether to perform a lookup with each verification started and return info about the phone number. | 
**Psd2Enabled** | **Bool**| Whether to pass PSD2 transaction parameters when starting a verification. | 
**PushApnCredentialSid** | **String**| Optional configuration for the Push factors. Set the APN Credential for this service. This will allow to send push notifications to iOS devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource) | 
**PushFcmCredentialSid** | **String**| Optional configuration for the Push factors. Set the FCM Credential for this service. This will allow to send push notifications to Android devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource) | 
**PushIncludeDate** | **Bool**| Optional configuration for the Push factors. If true, include the date in the Challenge&#39;s reponse. Otherwise, the date is omitted from the response. See [Challenge](https://www.twilio.com/docs/verify/api/challenge) resource’s details parameter for more info. Default: true | 
**SkipSmsToLandlines** | **Bool**| Whether to skip sending SMS verifications to landlines. Requires &#x60;lookup_enabled&#x60;. | 
**TtsName** | **String**| The name of an alternative text-to-speech service to use in phone calls. Applies only to TTS languages. | 

### Return type

[**VerifyV2Service**](VerifyV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVerification

> VerifyV2ServiceVerification CreateVerification(ctx, ServiceSid, optional)



Create a new Verification using a Service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to create the resource under. | 
 **optional** | ***CreateVerificationRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateVerificationRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Amount** | **String**| The amount of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled. | 
**AppHash** | **String**| Your [App Hash](https://developers.google.com/identity/sms-retriever/verify#computing_your_apps_hash_string) to be appended at the end of your verification SMS body. Applies only to SMS. Example SMS body: &#x60;&lt;#&gt; Your AppName verification code is: 1234 He42w354ol9&#x60;. | 
**Channel** | **String**| The verification method to use. Can be: [&#x60;email&#x60;](https://www.twilio.com/docs/verify/email), &#x60;sms&#x60; or &#x60;call&#x60;. | 
**ChannelConfiguration** | [**map[string]interface{}**](map[string]interface{}.md)| [&#x60;email&#x60;](https://www.twilio.com/docs/verify/email) channel configuration in json format. Must include &#39;from&#39; and &#39;from_name&#39;. | 
**CustomCode** | **String**| A pre-generated code to use for verification. The code can be between 4 and 10 characters, inclusive. | 
**CustomFriendlyName** | **String**| A custom user defined friendly name that overwrites the existing one in the verification message | 
**CustomMessage** | **String**| The text of a custom message to use for the verification. | 
**Locale** | **String**| The locale to use for the verification SMS or call. Can be: &#x60;af&#x60;, &#x60;ar&#x60;, &#x60;ca&#x60;, &#x60;cs&#x60;, &#x60;da&#x60;, &#x60;de&#x60;, &#x60;el&#x60;, &#x60;en&#x60;, &#x60;en-GB&#x60;, &#x60;es&#x60;, &#x60;fi&#x60;, &#x60;fr&#x60;, &#x60;he&#x60;, &#x60;hi&#x60;, &#x60;hr&#x60;, &#x60;hu&#x60;, &#x60;id&#x60;, &#x60;it&#x60;, &#x60;ja&#x60;, &#x60;ko&#x60;, &#x60;ms&#x60;, &#x60;nb&#x60;, &#x60;nl&#x60;, &#x60;pl&#x60;, &#x60;pt&#x60;, &#x60;pr-BR&#x60;, &#x60;ro&#x60;, &#x60;ru&#x60;, &#x60;sv&#x60;, &#x60;th&#x60;, &#x60;tl&#x60;, &#x60;tr&#x60;, &#x60;vi&#x60;, &#x60;zh&#x60;, &#x60;zh-CN&#x60;, or &#x60;zh-HK.&#x60; | 
**Payee** | **String**| The payee of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled. | 
**RateLimits** | [**map[string]interface{}**](map[string]interface{}.md)| The custom key-value pairs of Programmable Rate Limits. Keys correspond to &#x60;unique_name&#x60; fields defined when [creating your Rate Limit](https://www.twilio.com/docs/verify/api/service-rate-limits). Associated value pairs represent values in the request that you are rate limiting on. You may include multiple Rate Limit values in each request. | 
**SendDigits** | **String**| The digits to send after a phone call is answered, for example, to dial an extension. For more information, see the Programmable Voice documentation of [sendDigits](https://www.twilio.com/docs/voice/twiml/number#attributes-sendDigits). | 
**To** | **String**| The phone number or [email](https://www.twilio.com/docs/verify/email) to verify. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164). | 

### Return type

[**VerifyV2ServiceVerification**](VerifyV2ServiceVerification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVerificationCheck

> VerifyV2ServiceVerificationCheck CreateVerificationCheck(ctx, ServiceSid, optional)



challenge a specific Verification Check.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to create the resource under. | 
 **optional** | ***CreateVerificationCheckRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateVerificationCheckRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Amount** | **String**| The amount of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled. | 
**Code** | **String**| The 4-10 character string being verified. | 
**Payee** | **String**| The payee of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled. | 
**To** | **String**| The phone number or [email](https://www.twilio.com/docs/verify/email) to verify. Either this parameter or the &#x60;verification_sid&#x60; must be specified. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164). | 
**VerificationSid** | **String**| A SID that uniquely identifies the Verification Check. Either this parameter or the &#x60;to&#x60; phone number/[email](https://www.twilio.com/docs/verify/email) must be specified. | 

### Return type

[**VerifyV2ServiceVerificationCheck**](VerifyV2ServiceVerificationCheck.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebhook

> VerifyV2ServiceWebhook CreateWebhook(ctx, ServiceSid, optional)



Create a new Webhook for the Service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Service. | 
 **optional** | ***CreateWebhookRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateWebhookRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**EventTypes** | [**[]string**](string.md)| The array of events that this Webhook is subscribed to. Possible event types: &#x60;*, factor.deleted, factor.created, factor.verified, challenge.approved, challenge.denied&#x60; | 
**FriendlyName** | **String**| The string that you assigned to describe the webhook. **This value should not contain PII.** | 
**Status** | **String**| The webhook status. Default value is &#x60;enabled&#x60;. One of: &#x60;enabled&#x60; or &#x60;disabled&#x60; | 
**WebhookUrl** | **String**| The URL associated with this Webhook. | 

### Return type

[**VerifyV2ServiceWebhook**](VerifyV2ServiceWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBucket

> DeleteBucket(ctx, ServiceSid, RateLimitSid, Sid)



Delete a specific Bucket.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
**RateLimitSid** | **string**| The Twilio-provided string that uniquely identifies the Rate Limit resource. | 
**Sid** | **string**| A 34 character string that uniquely identifies this Bucket. | 

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


## DeleteEntity

> DeleteEntity(ctx, ServiceSid, Identity)



Delete a specific Entity.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Service. | 
**Identity** | **string**| The unique external identifier for the Entity of the Service | 

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


## DeleteFactor

> DeleteFactor(ctx, ServiceSid, Identity, Sid)



Delete a specific Factor.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Service. | 
**Identity** | **string**| Customer unique identity for the Entity owner of the Factor | 
**Sid** | **string**| A 34 character string that uniquely identifies this Factor. | 

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


## DeleteMessagingConfiguration

> DeleteMessagingConfiguration(ctx, ServiceSid, Country)



Delete a specific MessagingConfiguration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with. | 
**Country** | **string**| The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value &#x60;all&#x60;. | 

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


## DeleteRateLimit

> DeleteRateLimit(ctx, ServiceSid, Sid)



Delete a specific Rate Limit.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Rate Limit resource to fetch. | 

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


## DeleteService

> DeleteService(ctx, Sid)



Delete a specific Verification Service Instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Verification Service resource to delete. | 

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


## DeleteWebhook

> DeleteWebhook(ctx, ServiceSid, Sid)



Delete a specific Webhook.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Service. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Webhook resource to delete. | 

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


## FetchBucket

> VerifyV2ServiceRateLimitBucket FetchBucket(ctx, ServiceSid, RateLimitSid, Sid)



Fetch a specific Bucket.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
**RateLimitSid** | **string**| The Twilio-provided string that uniquely identifies the Rate Limit resource. | 
**Sid** | **string**| A 34 character string that uniquely identifies this Bucket. | 

### Return type

[**VerifyV2ServiceRateLimitBucket**](VerifyV2ServiceRateLimitBucket.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchChallenge

> VerifyV2ServiceEntityChallenge FetchChallenge(ctx, ServiceSid, Identity, Sid)



Fetch a specific Challenge.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Service. | 
**Identity** | **string**| The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, and generated by your external system, such as your user&#39;s UUID, GUID, or SID. | 
**Sid** | **string**| A 34 character string that uniquely identifies this Challenge. | 

### Return type

[**VerifyV2ServiceEntityChallenge**](VerifyV2ServiceEntityChallenge.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEntity

> VerifyV2ServiceEntity FetchEntity(ctx, ServiceSid, Identity)



Fetch a specific Entity.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Service. | 
**Identity** | **string**| The unique external identifier for the Entity of the Service | 

### Return type

[**VerifyV2ServiceEntity**](VerifyV2ServiceEntity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFactor

> VerifyV2ServiceEntityFactor FetchFactor(ctx, ServiceSid, Identity, Sid)



Fetch a specific Factor.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Service. | 
**Identity** | **string**| Customer unique identity for the Entity owner of the Factor | 
**Sid** | **string**| A 34 character string that uniquely identifies this Factor. | 

### Return type

[**VerifyV2ServiceEntityFactor**](VerifyV2ServiceEntityFactor.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchForm

> VerifyV2Form FetchForm(ctx, FormType)



Fetch the forms for a specific Form Type.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FormType** | **string**| The Type of this Form. Currently only &#x60;form-push&#x60; is supported. | 

### Return type

[**VerifyV2Form**](VerifyV2Form.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchMessagingConfiguration

> VerifyV2ServiceMessagingConfiguration FetchMessagingConfiguration(ctx, ServiceSid, Country)



Fetch a specific MessagingConfiguration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with. | 
**Country** | **string**| The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value &#x60;all&#x60;. | 

### Return type

[**VerifyV2ServiceMessagingConfiguration**](VerifyV2ServiceMessagingConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRateLimit

> VerifyV2ServiceRateLimit FetchRateLimit(ctx, ServiceSid, Sid)



Fetch a specific Rate Limit.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Rate Limit resource to fetch. | 

### Return type

[**VerifyV2ServiceRateLimit**](VerifyV2ServiceRateLimit.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> VerifyV2Service FetchService(ctx, Sid)



Fetch specific Verification Service Instance.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Verification Service resource to fetch. | 

### Return type

[**VerifyV2Service**](VerifyV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVerification

> VerifyV2ServiceVerification FetchVerification(ctx, ServiceSid, Sid)



Fetch a specific Verification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to fetch the resource from. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Verification resource to fetch. | 

### Return type

[**VerifyV2ServiceVerification**](VerifyV2ServiceVerification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVerificationAttempt

> VerifyV2VerificationAttempt FetchVerificationAttempt(ctx, Sid)



Fetch a specific verification attempt.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique SID identifier of a Verification Attempt | 

### Return type

[**VerifyV2VerificationAttempt**](VerifyV2VerificationAttempt.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchWebhook

> VerifyV2ServiceWebhook FetchWebhook(ctx, ServiceSid, Sid)



Fetch a specific Webhook.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Service. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Webhook resource to fetch. | 

### Return type

[**VerifyV2ServiceWebhook**](VerifyV2ServiceWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBucket

> ListBucketResponse ListBucket(ctx, ServiceSid, RateLimitSid, optional)



Retrieve a list of all Buckets for a Rate Limit.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
**RateLimitSid** | **string**| The Twilio-provided string that uniquely identifies the Rate Limit resource. | 
 **optional** | ***ListBucketRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListBucketRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListBucketResponse**](ListBucketResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChallenge

> ListChallengeResponse ListChallenge(ctx, ServiceSid, Identity, optional)



Retrieve a list of all Challenges for a Factor.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Service. | 
**Identity** | **string**| Customer unique identity for the Entity owner of the Challenge | 
 **optional** | ***ListChallengeRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListChallengeRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FactorSid** | **String**| The unique SID identifier of the Factor. | 
**Status** | **String**| The Status of the Challenges to fetch. One of &#x60;pending&#x60;, &#x60;expired&#x60;, &#x60;approved&#x60; or &#x60;denied&#x60;. | 
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListChallengeResponse**](ListChallengeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntity

> ListEntityResponse ListEntity(ctx, ServiceSid, optional)



Retrieve a list of all Entities for a Service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Service. | 
 **optional** | ***ListEntityRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEntityRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListEntityResponse**](ListEntityResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFactor

> ListFactorResponse ListFactor(ctx, ServiceSid, Identity, optional)



Retrieve a list of all Factors for an Entity.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Service. | 
**Identity** | **string**| Customer unique identity for the Entity owner of the Factor | 
 **optional** | ***ListFactorRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFactorRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListFactorResponse**](ListFactorResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessagingConfiguration

> ListMessagingConfigurationResponse ListMessagingConfiguration(ctx, ServiceSid, optional)



Retrieve a list of all Messaging Configurations for a Service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with. | 
 **optional** | ***ListMessagingConfigurationRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMessagingConfigurationRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListMessagingConfigurationResponse**](ListMessagingConfigurationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRateLimit

> ListRateLimitResponse ListRateLimit(ctx, ServiceSid, optional)



Retrieve a list of all Rate Limits for a service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
 **optional** | ***ListRateLimitRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRateLimitRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListRateLimitResponse**](ListRateLimitResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> ListServiceResponse ListService(ctx, optional)



Retrieve a list of all Verification Services for an account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListServiceRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListServiceRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListServiceResponse**](ListServiceResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVerificationAttempt

> ListVerificationAttemptResponse ListVerificationAttempt(ctx, optional)



List all the verification attempts for a given Account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListVerificationAttemptRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListVerificationAttemptRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**DateCreatedAfter** | **Time**| Datetime filter used to query Verification Attempts created after this datetime. | 
**DateCreatedBefore** | **Time**| Datetime filter used to query Verification Attempts created before this datetime. | 
**ChannelDataTo** | **String**| Destination of a verification. Depending on the type of channel, it could be a phone number in E.164 format or an email address. | 
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListVerificationAttemptResponse**](ListVerificationAttemptResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhook

> ListWebhookResponse ListWebhook(ctx, ServiceSid, optional)



Retrieve a list of all Webhooks for a Service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Service. | 
 **optional** | ***ListWebhookRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListWebhookRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListWebhookResponse**](ListWebhookResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBucket

> VerifyV2ServiceRateLimitBucket UpdateBucket(ctx, ServiceSid, RateLimitSid, Sid, optional)



Update a specific Bucket.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
**RateLimitSid** | **string**| The Twilio-provided string that uniquely identifies the Rate Limit resource. | 
**Sid** | **string**| A 34 character string that uniquely identifies this Bucket. | 
 **optional** | ***UpdateBucketRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateBucketRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Interval** | **Int32**| Number of seconds that the rate limit will be enforced over. | 
**Max** | **Int32**| Maximum number of requests permitted in during the interval. | 

### Return type

[**VerifyV2ServiceRateLimitBucket**](VerifyV2ServiceRateLimitBucket.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChallenge

> VerifyV2ServiceEntityChallenge UpdateChallenge(ctx, ServiceSid, Identity, Sid, optional)



Verify a specific Challenge.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Service. | 
**Identity** | **string**| Customer unique identity for the Entity owner of the Challenge | 
**Sid** | **string**| A 34 character string that uniquely identifies this Challenge. | 
 **optional** | ***UpdateChallengeRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateChallengeRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AuthPayload** | **String**| The optional payload needed to verify the Challenge. E.g., a TOTP would use the numeric code. | 

### Return type

[**VerifyV2ServiceEntityChallenge**](VerifyV2ServiceEntityChallenge.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFactor

> VerifyV2ServiceEntityFactor UpdateFactor(ctx, ServiceSid, Identity, Sid, optional)



Update a specific Factor. This endpoint can be used to Verify a Factor if passed an `AuthPayload` param.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Service. | 
**Identity** | **string**| Customer unique identity for the Entity owner of the Factor | 
**Sid** | **string**| A 34 character string that uniquely identifies this Factor. | 
 **optional** | ***UpdateFactorRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateFactorRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AuthPayload** | **String**| The optional payload needed to verify the Factor for the first time. E.g. for a TOTP, the numeric code. | 
**ConfigNotificationToken** | **String**| For APN, the device token. For FCM the registration token. It used to send the push notifications. Required when &#x60;factor_type&#x60; is &#x60;push&#x60; | 
**ConfigSdkVersion** | **String**| The Verify Push SDK version used to configure the factor | 
**FriendlyName** | **String**| The new friendly name of this Factor | 

### Return type

[**VerifyV2ServiceEntityFactor**](VerifyV2ServiceEntityFactor.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessagingConfiguration

> VerifyV2ServiceMessagingConfiguration UpdateMessagingConfiguration(ctx, ServiceSid, Country, optional)



Update a specific MessagingConfiguration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with. | 
**Country** | **string**| The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value &#x60;all&#x60;. | 
 **optional** | ***UpdateMessagingConfigurationRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateMessagingConfigurationRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**MessagingServiceSid** | **String**| The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) to be used to send SMS to the country of this configuration. | 

### Return type

[**VerifyV2ServiceMessagingConfiguration**](VerifyV2ServiceMessagingConfiguration.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRateLimit

> VerifyV2ServiceRateLimit UpdateRateLimit(ctx, ServiceSid, Sid, optional)



Update a specific Rate Limit.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Rate Limit resource to fetch. | 
 **optional** | ***UpdateRateLimitRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateRateLimitRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Description** | **String**| Description of this Rate Limit | 

### Return type

[**VerifyV2ServiceRateLimit**](VerifyV2ServiceRateLimit.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> VerifyV2Service UpdateService(ctx, Sid, optional)



Update a specific Verification Service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Service resource to update. | 
 **optional** | ***UpdateServiceRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateServiceRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CodeLength** | **Int32**| The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive. | 
**CustomCodeEnabled** | **Bool**| Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers. | 
**DoNotShareWarningEnabled** | **Bool**| Whether to add a privacy warning at the end of an SMS. **Disabled by default and applies only for SMS.** | 
**DtmfInputRequired** | **Bool**| Whether to ask the user to press a number before delivering the verify code in a phone call. | 
**FriendlyName** | **String**| A descriptive string that you create to describe the verification service. It can be up to 30 characters long. **This value should not contain PII.** | 
**LookupEnabled** | **Bool**| Whether to perform a lookup with each verification started and return info about the phone number. | 
**Psd2Enabled** | **Bool**| Whether to pass PSD2 transaction parameters when starting a verification. | 
**PushApnCredentialSid** | **String**| Optional configuration for the Push factors. Set the APN Credential for this service. This will allow to send push notifications to iOS devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource) | 
**PushFcmCredentialSid** | **String**| Optional configuration for the Push factors. Set the FCM Credential for this service. This will allow to send push notifications to Android devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource) | 
**PushIncludeDate** | **Bool**| Optional configuration for the Push factors. If true, include the date in the Challenge&#39;s reponse. Otherwise, the date is omitted from the response. See [Challenge](https://www.twilio.com/docs/verify/api/challenge) resource’s details parameter for more info. Default: true | 
**SkipSmsToLandlines** | **Bool**| Whether to skip sending SMS verifications to landlines. Requires &#x60;lookup_enabled&#x60;. | 
**TtsName** | **String**| The name of an alternative text-to-speech service to use in phone calls. Applies only to TTS languages. | 

### Return type

[**VerifyV2Service**](VerifyV2Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVerification

> VerifyV2ServiceVerification UpdateVerification(ctx, ServiceSid, Sid, optional)



Update a Verification status

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to update the resource from. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Verification resource to update. | 
 **optional** | ***UpdateVerificationRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateVerificationRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Status** | **String**| The new status of the resource. Can be: &#x60;canceled&#x60; or &#x60;approved&#x60;. | 

### Return type

[**VerifyV2ServiceVerification**](VerifyV2ServiceVerification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> VerifyV2ServiceWebhook UpdateWebhook(ctx, ServiceSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string**| The unique SID identifier of the Service. | 
**Sid** | **string**| The Twilio-provided string that uniquely identifies the Webhook resource to update. | 
 **optional** | ***UpdateWebhookRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateWebhookRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**EventTypes** | [**[]string**](string.md)| The array of events that this Webhook is subscribed to. Possible event types: &#x60;*, factor.deleted, factor.created, factor.verified, challenge.approved, challenge.denied&#x60; | 
**FriendlyName** | **String**| The string that you assigned to describe the webhook. **This value should not contain PII.** | 
**Status** | **String**| The webhook status. Default value is &#x60;enabled&#x60;. One of: &#x60;enabled&#x60; or &#x60;disabled&#x60; | 
**WebhookUrl** | **String**| The URL associated with this Webhook. | 

### Return type

[**VerifyV2ServiceWebhook**](VerifyV2ServiceWebhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

