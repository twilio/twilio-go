# \PreviewApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2FormsFormTypeGet**](PreviewApi.md#V2FormsFormTypeGet) | **Get** /v2/Forms/{FormType} | 
[**V2ServicesServiceSidAccessTokensPost**](PreviewApi.md#V2ServicesServiceSidAccessTokensPost) | **Post** /v2/Services/{ServiceSid}/AccessTokens | 
[**V2ServicesServiceSidEntitiesGet**](PreviewApi.md#V2ServicesServiceSidEntitiesGet) | **Get** /v2/Services/{ServiceSid}/Entities | 
[**V2ServicesServiceSidEntitiesIdentityChallengesGet**](PreviewApi.md#V2ServicesServiceSidEntitiesIdentityChallengesGet) | **Get** /v2/Services/{ServiceSid}/Entities/{Identity}/Challenges | 
[**V2ServicesServiceSidEntitiesIdentityChallengesPost**](PreviewApi.md#V2ServicesServiceSidEntitiesIdentityChallengesPost) | **Post** /v2/Services/{ServiceSid}/Entities/{Identity}/Challenges | 
[**V2ServicesServiceSidEntitiesIdentityChallengesSidGet**](PreviewApi.md#V2ServicesServiceSidEntitiesIdentityChallengesSidGet) | **Get** /v2/Services/{ServiceSid}/Entities/{Identity}/Challenges/{Sid} | 
[**V2ServicesServiceSidEntitiesIdentityChallengesSidPost**](PreviewApi.md#V2ServicesServiceSidEntitiesIdentityChallengesSidPost) | **Post** /v2/Services/{ServiceSid}/Entities/{Identity}/Challenges/{Sid} | 
[**V2ServicesServiceSidEntitiesIdentityDelete**](PreviewApi.md#V2ServicesServiceSidEntitiesIdentityDelete) | **Delete** /v2/Services/{ServiceSid}/Entities/{Identity} | 
[**V2ServicesServiceSidEntitiesIdentityFactorsGet**](PreviewApi.md#V2ServicesServiceSidEntitiesIdentityFactorsGet) | **Get** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors | 
[**V2ServicesServiceSidEntitiesIdentityFactorsPost**](PreviewApi.md#V2ServicesServiceSidEntitiesIdentityFactorsPost) | **Post** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors | 
[**V2ServicesServiceSidEntitiesIdentityFactorsSidDelete**](PreviewApi.md#V2ServicesServiceSidEntitiesIdentityFactorsSidDelete) | **Delete** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors/{Sid} | 
[**V2ServicesServiceSidEntitiesIdentityFactorsSidGet**](PreviewApi.md#V2ServicesServiceSidEntitiesIdentityFactorsSidGet) | **Get** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors/{Sid} | 
[**V2ServicesServiceSidEntitiesIdentityFactorsSidPost**](PreviewApi.md#V2ServicesServiceSidEntitiesIdentityFactorsSidPost) | **Post** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors/{Sid} | 
[**V2ServicesServiceSidEntitiesIdentityGet**](PreviewApi.md#V2ServicesServiceSidEntitiesIdentityGet) | **Get** /v2/Services/{ServiceSid}/Entities/{Identity} | 
[**V2ServicesServiceSidEntitiesPost**](PreviewApi.md#V2ServicesServiceSidEntitiesPost) | **Post** /v2/Services/{ServiceSid}/Entities | 
[**V2ServicesServiceSidWebhooksGet**](PreviewApi.md#V2ServicesServiceSidWebhooksGet) | **Get** /v2/Services/{ServiceSid}/Webhooks | 
[**V2ServicesServiceSidWebhooksPost**](PreviewApi.md#V2ServicesServiceSidWebhooksPost) | **Post** /v2/Services/{ServiceSid}/Webhooks | 
[**V2ServicesServiceSidWebhooksSidDelete**](PreviewApi.md#V2ServicesServiceSidWebhooksSidDelete) | **Delete** /v2/Services/{ServiceSid}/Webhooks/{Sid} | 
[**V2ServicesServiceSidWebhooksSidGet**](PreviewApi.md#V2ServicesServiceSidWebhooksSidGet) | **Get** /v2/Services/{ServiceSid}/Webhooks/{Sid} | 
[**V2ServicesServiceSidWebhooksSidPost**](PreviewApi.md#V2ServicesServiceSidWebhooksSidPost) | **Post** /v2/Services/{ServiceSid}/Webhooks/{Sid} | 



## V2FormsFormTypeGet

> VerifyV2Form V2FormsFormTypeGet(ctx, formType)



Fetch the forms for a specific Form Type.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formType** | **string**| The Type of this Form. Currently only &#x60;form-push&#x60; is supported. | 

### Return type

[**VerifyV2Form**](verify.v2.form.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidAccessTokensPost

> VerifyV2ServiceAccessToken V2ServicesServiceSidAccessTokensPost(ctx, serviceSid, optional)



Create a new enrollment Access Token for the Entity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**|  | 
 **optional** | ***V2ServicesServiceSidAccessTokensPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidAccessTokensPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **factorType** | **optional.**| The Type of this Factor. Eg. &#x60;push&#x60; | 
 **identity** | **optional.**| The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, and generated by your external system, such as your user&#39;s UUID, GUID, or SID. | 

### Return type

[**VerifyV2ServiceAccessToken**](verify.v2.service.access_token.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidEntitiesGet

> InlineResponse2001 V2ServicesServiceSidEntitiesGet(ctx, serviceSid, optional)



Retrieve a list of all Entities for a Service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Service. | 
 **optional** | ***V2ServicesServiceSidEntitiesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidEntitiesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **twilioSandboxMode** | **optional.**| The Twilio-Sandbox-Mode HTTP request header | 
 **pageSize** | **optional.**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidEntitiesIdentityChallengesGet

> InlineResponse2002 V2ServicesServiceSidEntitiesIdentityChallengesGet(ctx, serviceSid, identity, optional)



Retrieve a list of all Challenges for a Factor.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Service. | 
**identity** | **string**| Customer unique identity for the Entity owner of the Challenge | 
 **optional** | ***V2ServicesServiceSidEntitiesIdentityChallengesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidEntitiesIdentityChallengesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **twilioSandboxMode** | **optional.**| The Twilio-Sandbox-Mode HTTP request header | 
 **factorSid** | **optional.**| The unique SID identifier of the Factor. | 
 **status** | **optional.**| The Status of the Challenges to fetch. One of &#x60;pending&#x60;, &#x60;expired&#x60;, &#x60;approved&#x60; or &#x60;denied&#x60;. | 
 **pageSize** | **optional.**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidEntitiesIdentityChallengesPost

> VerifyV2ServiceEntityChallenge V2ServicesServiceSidEntitiesIdentityChallengesPost(ctx, serviceSid, identity, optional)



Create a new Challenge for the Factor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Service. | 
**identity** | **string**| The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, and generated by your external system, such as your user&#39;s UUID, GUID, or SID. | 
 **optional** | ***V2ServicesServiceSidEntitiesIdentityChallengesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidEntitiesIdentityChallengesPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **twilioSandboxMode** | **optional.**| The Twilio-Sandbox-Mode HTTP request header | 
 **details** | **optional.**| Details provided to give context about the Challenge. Shown to the end user. It must be a stringified JSON with the following structure: {\\\&quot;message\\\&quot;: \\\&quot;string\\\&quot;, \\\&quot;fields\\\&quot;: [ { \\\&quot;label\\\&quot;: \\\&quot;string\\\&quot;, \\\&quot;value\\\&quot;: \\\&quot;string\\\&quot;}]}. &#x60;message&#x60; is required. If you send the &#x60;fields&#x60; property, each field has to include &#x60;label&#x60; and &#x60;value&#x60; properties. If you had set &#x60;include_date&#x3D;true&#x60; in the &#x60;push&#x60; configuration of the [service](https://www.twilio.com/docs/verify/api/service), the response will also include the challenge&#39;s date created value as an additional field called &#x60;date&#x60; | 
 **expirationDate** | **optional.**| The future date in which this Challenge will expire, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. | 
 **factorSid** | **optional.**| The unique SID identifier of the Factor. | 
 **hiddenDetails** | **optional.**| Details provided to give context about the Challenge. Not shown to the end user. It must be a stringified JSON with only strings values eg. &#x60;{\\\&quot;ip\\\&quot;: \\\&quot;172.168.1.234\\\&quot;}&#x60; | 

### Return type

[**VerifyV2ServiceEntityChallenge**](verify.v2.service.entity.challenge.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidEntitiesIdentityChallengesSidGet

> VerifyV2ServiceEntityChallenge V2ServicesServiceSidEntitiesIdentityChallengesSidGet(ctx, serviceSid, identity, sid, optional)



Fetch a specific Challenge.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Service. | 
**identity** | **string**| The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, and generated by your external system, such as your user&#39;s UUID, GUID, or SID. | 
**sid** | **string**| A 34 character string that uniquely identifies this Challenge. | 
 **optional** | ***V2ServicesServiceSidEntitiesIdentityChallengesSidGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidEntitiesIdentityChallengesSidGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **twilioSandboxMode** | **optional.**| The Twilio-Sandbox-Mode HTTP request header | 

### Return type

[**VerifyV2ServiceEntityChallenge**](verify.v2.service.entity.challenge.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidEntitiesIdentityChallengesSidPost

> VerifyV2ServiceEntityChallenge V2ServicesServiceSidEntitiesIdentityChallengesSidPost(ctx, serviceSid, identity, sid, optional)



Verify a specific Challenge.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Service. | 
**identity** | **string**| Customer unique identity for the Entity owner of the Challenge | 
**sid** | **string**| A 34 character string that uniquely identifies this Challenge. | 
 **optional** | ***V2ServicesServiceSidEntitiesIdentityChallengesSidPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidEntitiesIdentityChallengesSidPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **twilioSandboxMode** | **optional.**| The Twilio-Sandbox-Mode HTTP request header | 
 **authPayload** | **optional.**| The optional payload needed to verify the Challenge. E.g., a TOTP would use the numeric code. | 

### Return type

[**VerifyV2ServiceEntityChallenge**](verify.v2.service.entity.challenge.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidEntitiesIdentityDelete

> V2ServicesServiceSidEntitiesIdentityDelete(ctx, serviceSid, identity, optional)



Delete a specific Entity.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Service. | 
**identity** | **string**| The unique external identifier for the Entity of the Service | 
 **optional** | ***V2ServicesServiceSidEntitiesIdentityDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidEntitiesIdentityDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **twilioSandboxMode** | **optional.**| The Twilio-Sandbox-Mode HTTP request header | 

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


## V2ServicesServiceSidEntitiesIdentityFactorsGet

> InlineResponse2003 V2ServicesServiceSidEntitiesIdentityFactorsGet(ctx, serviceSid, identity, optional)



Retrieve a list of all Factors for an Entity.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Service. | 
**identity** | **string**| Customer unique identity for the Entity owner of the Factor | 
 **optional** | ***V2ServicesServiceSidEntitiesIdentityFactorsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidEntitiesIdentityFactorsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **twilioSandboxMode** | **optional.**| The Twilio-Sandbox-Mode HTTP request header | 
 **pageSize** | **optional.**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidEntitiesIdentityFactorsPost

> VerifyV2ServiceEntityFactor V2ServicesServiceSidEntitiesIdentityFactorsPost(ctx, serviceSid, identity, optional)



Create a new Factor for the Entity

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Service. | 
**identity** | **string**| Customer unique identity for the Entity owner of the Factor | 
 **optional** | ***V2ServicesServiceSidEntitiesIdentityFactorsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidEntitiesIdentityFactorsPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **twilioSandboxMode** | **optional.**| The Twilio-Sandbox-Mode HTTP request header | 
 **binding** | **optional.**| A unique binding for this Factor that identifies it. E.g. the algorithm and public key for &#x60;push&#x60; factors. It must be a json string with the required properties for the given factor type. Required when creating a new Factor. This value is never returned because it can contain customer secrets. | 
 **config** | **optional.**| The config required for this Factor. It must be a json string with the required properties for the given factor type | 
 **factorType** | **optional.**| The Type of this Factor. Currently only &#x60;push&#x60; is supported | 
 **friendlyName** | **optional.**| The friendly name of this Factor | 

### Return type

[**VerifyV2ServiceEntityFactor**](verify.v2.service.entity.factor.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidEntitiesIdentityFactorsSidDelete

> V2ServicesServiceSidEntitiesIdentityFactorsSidDelete(ctx, serviceSid, identity, sid, optional)



Delete a specific Factor.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Service. | 
**identity** | **string**| Customer unique identity for the Entity owner of the Factor | 
**sid** | **string**| A 34 character string that uniquely identifies this Factor. | 
 **optional** | ***V2ServicesServiceSidEntitiesIdentityFactorsSidDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidEntitiesIdentityFactorsSidDeleteOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **twilioSandboxMode** | **optional.**| The Twilio-Sandbox-Mode HTTP request header | 

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


## V2ServicesServiceSidEntitiesIdentityFactorsSidGet

> VerifyV2ServiceEntityFactor V2ServicesServiceSidEntitiesIdentityFactorsSidGet(ctx, serviceSid, identity, sid, optional)



Fetch a specific Factor.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Service. | 
**identity** | **string**| Customer unique identity for the Entity owner of the Factor | 
**sid** | **string**| A 34 character string that uniquely identifies this Factor. | 
 **optional** | ***V2ServicesServiceSidEntitiesIdentityFactorsSidGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidEntitiesIdentityFactorsSidGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **twilioSandboxMode** | **optional.**| The Twilio-Sandbox-Mode HTTP request header | 

### Return type

[**VerifyV2ServiceEntityFactor**](verify.v2.service.entity.factor.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidEntitiesIdentityFactorsSidPost

> VerifyV2ServiceEntityFactor V2ServicesServiceSidEntitiesIdentityFactorsSidPost(ctx, serviceSid, identity, sid, optional)



Update a specific Factor. This endpoint can be used to Verify a Factor if passed an `AuthPayload` param.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Service. | 
**identity** | **string**| Customer unique identity for the Entity owner of the Factor | 
**sid** | **string**| A 34 character string that uniquely identifies this Factor. | 
 **optional** | ***V2ServicesServiceSidEntitiesIdentityFactorsSidPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidEntitiesIdentityFactorsSidPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **twilioSandboxMode** | **optional.**| The Twilio-Sandbox-Mode HTTP request header | 
 **authPayload** | **optional.**| The optional payload needed to verify the Factor for the first time. E.g. for a TOTP, the numeric code. | 
 **config** | **optional.**| The new config for this Factor. It must be a json string with the required properties for the given factor type | 
 **friendlyName** | **optional.**| The new friendly name of this Factor | 

### Return type

[**VerifyV2ServiceEntityFactor**](verify.v2.service.entity.factor.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidEntitiesIdentityGet

> VerifyV2ServiceEntity V2ServicesServiceSidEntitiesIdentityGet(ctx, serviceSid, identity, optional)



Fetch a specific Entity.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Service. | 
**identity** | **string**| The unique external identifier for the Entity of the Service | 
 **optional** | ***V2ServicesServiceSidEntitiesIdentityGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidEntitiesIdentityGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **twilioSandboxMode** | **optional.**| The Twilio-Sandbox-Mode HTTP request header | 

### Return type

[**VerifyV2ServiceEntity**](verify.v2.service.entity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidEntitiesPost

> VerifyV2ServiceEntity V2ServicesServiceSidEntitiesPost(ctx, serviceSid, optional)



Create a new Entity for the Service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Service. | 
 **optional** | ***V2ServicesServiceSidEntitiesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidEntitiesPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **twilioSandboxMode** | **optional.**| The Twilio-Sandbox-Mode HTTP request header | 
 **identity** | **optional.**| The unique external identifier for the Entity of the Service | 

### Return type

[**VerifyV2ServiceEntity**](verify.v2.service.entity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidWebhooksGet

> InlineResponse2007 V2ServicesServiceSidWebhooksGet(ctx, serviceSid, optional)



Retrieve a list of all Webhooks for a Service.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Service. | 
 **optional** | ***V2ServicesServiceSidWebhooksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidWebhooksGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidWebhooksPost

> VerifyV2ServiceWebhook V2ServicesServiceSidWebhooksPost(ctx, serviceSid, optional)



Create a new Webhook for the Service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Service. | 
 **optional** | ***V2ServicesServiceSidWebhooksPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidWebhooksPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventTypes** | [**optional.Interface of []string**](string.md)| The array of events that this Webhook is subscribed to. Possible event types: &#x60;*, factor.deleted, factor.created, factor.verified, challenge.approved, challenge.denied&#x60; | 
 **friendlyName** | **optional.**| The string that you assigned to describe the webhook. **This value should not contain PII.** | 
 **status** | **optional.**| The webhook status. Default value is &#x60;enabled&#x60;. One of: &#x60;enabled&#x60; or &#x60;disabled&#x60; | 
 **webhookUrl** | **optional.**| The URL associated with this Webhook. | 

### Return type

[**VerifyV2ServiceWebhook**](verify.v2.service.webhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidWebhooksSidDelete

> V2ServicesServiceSidWebhooksSidDelete(ctx, serviceSid, sid)



Delete a specific Webhook.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Service. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the Webhook resource to delete. | 

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


## V2ServicesServiceSidWebhooksSidGet

> VerifyV2ServiceWebhook V2ServicesServiceSidWebhooksSidGet(ctx, serviceSid, sid)



Fetch a specific Webhook.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Service. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the Webhook resource to fetch. | 

### Return type

[**VerifyV2ServiceWebhook**](verify.v2.service.webhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V2ServicesServiceSidWebhooksSidPost

> VerifyV2ServiceWebhook V2ServicesServiceSidWebhooksSidPost(ctx, serviceSid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The unique SID identifier of the Service. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the Webhook resource to update. | 
 **optional** | ***V2ServicesServiceSidWebhooksSidPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V2ServicesServiceSidWebhooksSidPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **eventTypes** | [**optional.Interface of []string**](string.md)| The array of events that this Webhook is subscribed to. Possible event types: &#x60;*, factor.deleted, factor.created, factor.verified, challenge.approved, challenge.denied&#x60; | 
 **friendlyName** | **optional.**| The string that you assigned to describe the webhook. **This value should not contain PII.** | 
 **status** | **optional.**| The webhook status. Default value is &#x60;enabled&#x60;. One of: &#x60;enabled&#x60; or &#x60;disabled&#x60; | 
 **webhookUrl** | **optional.**| The URL associated with this Webhook. | 

### Return type

[**VerifyV2ServiceWebhook**](verify.v2.service.webhook.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

