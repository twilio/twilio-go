# ServicesEntitiesFactorsApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewFactor**](ServicesEntitiesFactorsApi.md#CreateNewFactor) | **Post** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors | 
[**DeleteFactor**](ServicesEntitiesFactorsApi.md#DeleteFactor) | **Delete** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors/{Sid} | 
[**FetchFactor**](ServicesEntitiesFactorsApi.md#FetchFactor) | **Get** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors/{Sid} | 
[**ListFactor**](ServicesEntitiesFactorsApi.md#ListFactor) | **Get** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors | 
[**UpdateFactor**](ServicesEntitiesFactorsApi.md#UpdateFactor) | **Post** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors/{Sid} | 



## CreateNewFactor

> VerifyV2ServiceEntityNewFactor CreateNewFactor(ctx, ServiceSidIdentityoptional)



Create a new Factor for the Entity

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Identity** | **string** | Customer unique identity for the Entity owner of the Factor. This value must be between 8 and 64 characters long.

### Other Parameters

Other parameters are passed through a pointer to a CreateNewFactorParams struct


Name | Type | Description
------------- | ------------- | -------------
**BindingAlg** | **string** | The algorithm used when &#x60;factor_type&#x60; is &#x60;push&#x60;. Algorithm supported: &#x60;ES256&#x60;
**BindingPublicKey** | **string** | The Ecdsa public key in PKIX, ASN.1 DER format encoded in Base64.  Required when &#x60;factor_type&#x60; is &#x60;push&#x60;
**BindingSecret** | **string** | The shared secret for TOTP factors encoded in Base32. This can be provided when creating the Factor, otherwise it will be generated.  Used when &#x60;factor_type&#x60; is &#x60;totp&#x60;
**ConfigAlg** | **string** | The algorithm used to derive the TOTP codes. Can be &#x60;sha1&#x60;, &#x60;sha256&#x60; or &#x60;sha512&#x60;. Defaults to &#x60;sha1&#x60;.  Used when &#x60;factor_type&#x60; is &#x60;totp&#x60;
**ConfigAppId** | **string** | The ID that uniquely identifies your app in the Google or Apple store, such as &#x60;com.example.myapp&#x60;. It can be up to 100 characters long.  Required when &#x60;factor_type&#x60; is &#x60;push&#x60;.
**ConfigCodeLength** | **int** | Number of digits for generated TOTP codes. Must be between 3 and 8, inclusive. The default value is defined at the service level in the property &#x60;totp.code_length&#x60;. If not configured defaults to 6.  Used when &#x60;factor_type&#x60; is &#x60;totp&#x60;
**ConfigNotificationPlatform** | **string** | The transport technology used to generate the Notification Token. Can be &#x60;apn&#x60; or &#x60;fcm&#x60;.  Required when &#x60;factor_type&#x60; is &#x60;push&#x60;.
**ConfigNotificationToken** | **string** | For APN, the device token. For FCM the registration token. It used to send the push notifications. Must be between 32 and 255 characters long.  Required when &#x60;factor_type&#x60; is &#x60;push&#x60;.
**ConfigSdkVersion** | **string** | The Verify Push SDK version used to configure the factor  Required when &#x60;factor_type&#x60; is &#x60;push&#x60;
**ConfigSkew** | **int** | The number of time-steps, past and future, that are valid for validation of TOTP codes. Must be between 0 and 2, inclusive. The default value is defined at the service level in the property &#x60;totp.skew&#x60;. If not configured defaults to 1.  Used when &#x60;factor_type&#x60; is &#x60;totp&#x60;
**ConfigTimeStep** | **int** | Defines how often, in seconds, are TOTP codes generated. i.e, a new TOTP code is generated every time_step seconds. Must be between 20 and 60 seconds, inclusive. The default value is defined at the service level in the property &#x60;totp.time_step&#x60;. Defaults to 30 seconds if not configured.  Used when &#x60;factor_type&#x60; is &#x60;totp&#x60;
**FactorType** | **string** | The Type of this Factor. Currently &#x60;push&#x60; and &#x60;totp&#x60; are supported. For &#x60;totp&#x60; to work, you need to contact [Twilio sales](https://www.twilio.com/help/sales) first to have the Verify TOTP feature enabled for your Twilio account.
**FriendlyName** | **string** | The friendly name of this Factor. This can be any string up to 64 characters, meant for humans to distinguish between Factors. For &#x60;factor_type&#x60; &#x60;push&#x60;, this could be a device name. For &#x60;factor_type&#x60; &#x60;totp&#x60;, this value is used as the “account name” in constructing the &#x60;binding.uri&#x60; property. At the same time, we recommend avoiding providing PII.

### Return type

[**VerifyV2ServiceEntityNewFactor**](VerifyV2ServiceEntityNewFactor.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFactor

> DeleteFactor(ctx, ServiceSidIdentitySid)



Delete a specific Factor.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Identity** | **string** | Customer unique identity for the Entity owner of the Factor. This value must be between 8 and 64 characters long.
**Sid** | **string** | A 34 character string that uniquely identifies this Factor.

### Other Parameters

Other parameters are passed through a pointer to a DeleteFactorParams struct


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


## FetchFactor

> VerifyV2ServiceEntityFactor FetchFactor(ctx, ServiceSidIdentitySid)



Fetch a specific Factor.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Identity** | **string** | Customer unique identity for the Entity owner of the Factor. This value must be between 8 and 64 characters long.
**Sid** | **string** | A 34 character string that uniquely identifies this Factor.

### Other Parameters

Other parameters are passed through a pointer to a FetchFactorParams struct


Name | Type | Description
------------- | ------------- | -------------

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


## ListFactor

> ListFactorResponse ListFactor(ctx, ServiceSidIdentityoptional)



Retrieve a list of all Factors for an Entity.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Identity** | **string** | Customer unique identity for the Entity owner of the Factor. This value must be between 8 and 64 characters long.

### Other Parameters

Other parameters are passed through a pointer to a ListFactorParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

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


## UpdateFactor

> VerifyV2ServiceEntityFactor UpdateFactor(ctx, ServiceSidIdentitySidoptional)



Update a specific Factor. This endpoint can be used to Verify a Factor if passed an `AuthPayload` param.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Identity** | **string** | Customer unique identity for the Entity owner of the Factor. This value must be between 8 and 64 characters long.
**Sid** | **string** | A 34 character string that uniquely identifies this Factor.

### Other Parameters

Other parameters are passed through a pointer to a UpdateFactorParams struct


Name | Type | Description
------------- | ------------- | -------------
**AuthPayload** | **string** | The optional payload needed to verify the Factor for the first time. E.g. for a TOTP, the numeric code.
**ConfigAlg** | **string** | The algorithm used to derive the TOTP codes. Can be &#x60;sha1&#x60;, &#x60;sha256&#x60; or &#x60;sha512&#x60;
**ConfigCodeLength** | **int** | Number of digits for generated TOTP codes. Must be between 3 and 8, inclusive
**ConfigNotificationToken** | **string** | For APN, the device token. For FCM the registration token. It used to send the push notifications. Required when &#x60;factor_type&#x60; is &#x60;push&#x60;. If specified, this value must be between 32 and 255 characters long.
**ConfigSdkVersion** | **string** | The Verify Push SDK version used to configure the factor
**ConfigSkew** | **int** | The number of time-steps, past and future, that are valid for validation of TOTP codes. Must be between 0 and 2, inclusive
**ConfigTimeStep** | **int** | Defines how often, in seconds, are TOTP codes generated. i.e, a new TOTP code is generated every time_step seconds. Must be between 20 and 60 seconds, inclusive
**FriendlyName** | **string** | The new friendly name of this Factor. It can be up to 64 characters.

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

