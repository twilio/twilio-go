# \DefaultApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessToken**](DefaultApi.md#CreateAccessToken) | **Post** /v2/Services/{ServiceSid}/AccessTokens | 
[**CreateBucket**](DefaultApi.md#CreateBucket) | **Post** /v2/Services/{ServiceSid}/RateLimits/{RateLimitSid}/Buckets | 
[**CreateChallenge**](DefaultApi.md#CreateChallenge) | **Post** /v2/Services/{ServiceSid}/Entities/{Identity}/Challenges | 
[**CreateEntity**](DefaultApi.md#CreateEntity) | **Post** /v2/Services/{ServiceSid}/Entities | 
[**CreateMessagingConfiguration**](DefaultApi.md#CreateMessagingConfiguration) | **Post** /v2/Services/{ServiceSid}/MessagingConfigurations | 
[**CreateNewFactor**](DefaultApi.md#CreateNewFactor) | **Post** /v2/Services/{ServiceSid}/Entities/{Identity}/Factors | 
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

> VerifyV2ServiceAccessToken CreateAccessToken(ctx, ServiceSid).FactorType(FactorType).Identity(Identity).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The unique SID identifier of the Service.
    FactorType := "FactorType_example" // string | The Type of this Factor. Eg. `push` (optional)
    Identity := "Identity_example" // string | The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, and generated by your external system, such as your user's UUID, GUID, or SID. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateAccessToken(context.Background(), ServiceSid).FactorType(FactorType).Identity(Identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccessToken`: VerifyV2ServiceAccessToken
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.

### Other Parameters

Other parameters are passed through a pointer to a CreateAccessTokenParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FactorType** | **string** | The Type of this Factor. Eg. &#x60;push&#x60;
 **Identity** | **string** | The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, and generated by your external system, such as your user&#39;s UUID, GUID, or SID.

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

> VerifyV2ServiceRateLimitBucket CreateBucket(ctx, ServiceSid, RateLimitSid).Interval(Interval).Max(Max).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
    RateLimitSid := "RateLimitSid_example" // string | The Twilio-provided string that uniquely identifies the Rate Limit resource.
    Interval := int32(56) // int32 | Number of seconds that the rate limit will be enforced over. (optional)
    Max := int32(56) // int32 | Maximum number of requests permitted in during the interval. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateBucket(context.Background(), ServiceSid, RateLimitSid).Interval(Interval).Max(Max).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBucket`: VerifyV2ServiceRateLimitBucket
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
**RateLimitSid** | **string** | The Twilio-provided string that uniquely identifies the Rate Limit resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateBucketParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Interval** | **int32** | Number of seconds that the rate limit will be enforced over.
 **Max** | **int32** | Maximum number of requests permitted in during the interval.

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

> VerifyV2ServiceEntityChallenge CreateChallenge(ctx, ServiceSid, Identity).DetailsFields(DetailsFields).DetailsMessage(DetailsMessage).ExpirationDate(ExpirationDate).FactorSid(FactorSid).HiddenDetails(HiddenDetails).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The unique SID identifier of the Service.
    Identity := "Identity_example" // string | The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, and generated by your external system, such as your user's UUID, GUID, or SID. This value must be between 8 and 64 characters long.
    DetailsFields := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | A list of objects that describe the Fields included in the Challenge. Each object contains the label and value of the field. Used when `factor_type` is `push`. (optional)
    DetailsMessage := "DetailsMessage_example" // string | Shown to the user when the push notification arrives. Required when `factor_type` is `push` (optional)
    ExpirationDate := time.Now() // time.Time | The date-time when this Challenge expires, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. The default value is five (5) minutes after Challenge creation. The max value is sixty (60) minutes after creation. (optional)
    FactorSid := "FactorSid_example" // string | The unique SID identifier of the Factor. (optional)
    HiddenDetails := TODO // map[string]interface{} | Details provided to give context about the Challenge. Not shown to the end user. It must be a stringified JSON with only strings values eg. `{\\\"ip\\\": \\\"172.168.1.234\\\"}` (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateChallenge(context.Background(), ServiceSid, Identity).DetailsFields(DetailsFields).DetailsMessage(DetailsMessage).ExpirationDate(ExpirationDate).FactorSid(FactorSid).HiddenDetails(HiddenDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateChallenge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateChallenge`: VerifyV2ServiceEntityChallenge
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateChallenge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Identity** | **string** | The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, and generated by your external system, such as your user&#39;s UUID, GUID, or SID. This value must be between 8 and 64 characters long.

### Other Parameters

Other parameters are passed through a pointer to a CreateChallengeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **DetailsFields** | **[]map[string]interface{}** | A list of objects that describe the Fields included in the Challenge. Each object contains the label and value of the field. Used when &#x60;factor_type&#x60; is &#x60;push&#x60;.
 **DetailsMessage** | **string** | Shown to the user when the push notification arrives. Required when &#x60;factor_type&#x60; is &#x60;push&#x60;
 **ExpirationDate** | **time.Time** | The date-time when this Challenge expires, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. The default value is five (5) minutes after Challenge creation. The max value is sixty (60) minutes after creation.
 **FactorSid** | **string** | The unique SID identifier of the Factor.
 **HiddenDetails** | [**map[string]interface{}**](map[string]interface{}.md) | Details provided to give context about the Challenge. Not shown to the end user. It must be a stringified JSON with only strings values eg. &#x60;{\\\&quot;ip\\\&quot;: \\\&quot;172.168.1.234\\\&quot;}&#x60;

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

> VerifyV2ServiceEntity CreateEntity(ctx, ServiceSid).Identity(Identity).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The unique SID identifier of the Service.
    Identity := "Identity_example" // string | The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, length between 8 and 64 characters, and generated by your external system, such as your user's UUID, GUID, or SID. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateEntity(context.Background(), ServiceSid).Identity(Identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEntity`: VerifyV2ServiceEntity
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.

### Other Parameters

Other parameters are passed through a pointer to a CreateEntityParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Identity** | **string** | The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, length between 8 and 64 characters, and generated by your external system, such as your user&#39;s UUID, GUID, or SID.

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


## CreateMessagingConfiguration

> VerifyV2ServiceMessagingConfiguration CreateMessagingConfiguration(ctx, ServiceSid).Country(Country).MessagingServiceSid(MessagingServiceSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with.
    Country := "Country_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value `all`. (optional)
    MessagingServiceSid := "MessagingServiceSid_example" // string | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) to be used to send SMS to the country of this configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateMessagingConfiguration(context.Background(), ServiceSid).Country(Country).MessagingServiceSid(MessagingServiceSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMessagingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMessagingConfiguration`: VerifyV2ServiceMessagingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateMessagingConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with.

### Other Parameters

Other parameters are passed through a pointer to a CreateMessagingConfigurationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Country** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value &#x60;all&#x60;.
 **MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) to be used to send SMS to the country of this configuration.

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


## CreateNewFactor

> VerifyV2ServiceEntityNewFactor CreateNewFactor(ctx, ServiceSid, Identity).BindingAlg(BindingAlg).BindingPublicKey(BindingPublicKey).BindingSecret(BindingSecret).ConfigAlg(ConfigAlg).ConfigAppId(ConfigAppId).ConfigCodeLength(ConfigCodeLength).ConfigNotificationPlatform(ConfigNotificationPlatform).ConfigNotificationToken(ConfigNotificationToken).ConfigSdkVersion(ConfigSdkVersion).ConfigSkew(ConfigSkew).ConfigTimeStep(ConfigTimeStep).FactorType(FactorType).FriendlyName(FriendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The unique SID identifier of the Service.
    Identity := "Identity_example" // string | Customer unique identity for the Entity owner of the Factor. This value must be between 8 and 64 characters long.
    BindingAlg := "BindingAlg_example" // string | The algorithm used when `factor_type` is `push`. Algorithm supported: `ES256` (optional)
    BindingPublicKey := "BindingPublicKey_example" // string | The Ecdsa public key in PKIX, ASN.1 DER format encoded in Base64 (optional)
    BindingSecret := "BindingSecret_example" // string | The shared secret for TOTP factors encoded in Base32 (optional)
    ConfigAlg := "ConfigAlg_example" // string | The algorithm used to derive the TOTP codes. Can be `sha1`, `sha256` or `sha512`. Defaults to `sha1` (optional)
    ConfigAppId := "ConfigAppId_example" // string | The ID that uniquely identifies your app in the Google or Apple store, such as `com.example.myapp`. Required when `factor_type` is `push`. If specified, it can be up to 100 characters long. (optional)
    ConfigCodeLength := int32(56) // int32 | Number of digits for generated TOTP codes. Must be between 3 and 8, inclusive. Defaults to 6 (optional)
    ConfigNotificationPlatform := "ConfigNotificationPlatform_example" // string | The transport technology used to generate the Notification Token. Can be `apn` or `fcm`. Required when `factor_type` is `push` (optional)
    ConfigNotificationToken := "ConfigNotificationToken_example" // string | For APN, the device token. For FCM the registration token. It used to send the push notifications. Required when `factor_type` is `push`. If specified, this value must be between 32 and 255 characters long. (optional)
    ConfigSdkVersion := "ConfigSdkVersion_example" // string | The Verify Push SDK version used to configure the factor (optional)
    ConfigSkew := int32(56) // int32 | The number of time-steps, past and future, that are valid for validation of TOTP codes. Must be between 0 and 2, inclusive. Defaults to 1 (optional)
    ConfigTimeStep := int32(56) // int32 | Defines how often, in seconds, are TOTP codes generated. i.e, a new TOTP code is generated every time_step seconds. Must be between 20 and 60 seconds, inclusive. Defaults to 30 seconds (optional)
    FactorType := "FactorType_example" // string | The Type of this Factor. Currently only `push` is supported (optional)
    FriendlyName := "FriendlyName_example" // string | The friendly name of this Factor. It can be up to 64 characters. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateNewFactor(context.Background(), ServiceSid, Identity).BindingAlg(BindingAlg).BindingPublicKey(BindingPublicKey).BindingSecret(BindingSecret).ConfigAlg(ConfigAlg).ConfigAppId(ConfigAppId).ConfigCodeLength(ConfigCodeLength).ConfigNotificationPlatform(ConfigNotificationPlatform).ConfigNotificationToken(ConfigNotificationToken).ConfigSdkVersion(ConfigSdkVersion).ConfigSkew(ConfigSkew).ConfigTimeStep(ConfigTimeStep).FactorType(FactorType).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNewFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewFactor`: VerifyV2ServiceEntityNewFactor
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateNewFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Identity** | **string** | Customer unique identity for the Entity owner of the Factor. This value must be between 8 and 64 characters long.

### Other Parameters

Other parameters are passed through a pointer to a CreateNewFactorParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **BindingAlg** | **string** | The algorithm used when &#x60;factor_type&#x60; is &#x60;push&#x60;. Algorithm supported: &#x60;ES256&#x60;
 **BindingPublicKey** | **string** | The Ecdsa public key in PKIX, ASN.1 DER format encoded in Base64
 **BindingSecret** | **string** | The shared secret for TOTP factors encoded in Base32
 **ConfigAlg** | **string** | The algorithm used to derive the TOTP codes. Can be &#x60;sha1&#x60;, &#x60;sha256&#x60; or &#x60;sha512&#x60;. Defaults to &#x60;sha1&#x60;
 **ConfigAppId** | **string** | The ID that uniquely identifies your app in the Google or Apple store, such as &#x60;com.example.myapp&#x60;. Required when &#x60;factor_type&#x60; is &#x60;push&#x60;. If specified, it can be up to 100 characters long.
 **ConfigCodeLength** | **int32** | Number of digits for generated TOTP codes. Must be between 3 and 8, inclusive. Defaults to 6
 **ConfigNotificationPlatform** | **string** | The transport technology used to generate the Notification Token. Can be &#x60;apn&#x60; or &#x60;fcm&#x60;. Required when &#x60;factor_type&#x60; is &#x60;push&#x60;
 **ConfigNotificationToken** | **string** | For APN, the device token. For FCM the registration token. It used to send the push notifications. Required when &#x60;factor_type&#x60; is &#x60;push&#x60;. If specified, this value must be between 32 and 255 characters long.
 **ConfigSdkVersion** | **string** | The Verify Push SDK version used to configure the factor
 **ConfigSkew** | **int32** | The number of time-steps, past and future, that are valid for validation of TOTP codes. Must be between 0 and 2, inclusive. Defaults to 1
 **ConfigTimeStep** | **int32** | Defines how often, in seconds, are TOTP codes generated. i.e, a new TOTP code is generated every time_step seconds. Must be between 20 and 60 seconds, inclusive. Defaults to 30 seconds
 **FactorType** | **string** | The Type of this Factor. Currently only &#x60;push&#x60; is supported
 **FriendlyName** | **string** | The friendly name of this Factor. It can be up to 64 characters.

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


## CreateRateLimit

> VerifyV2ServiceRateLimit CreateRateLimit(ctx, ServiceSid).Description(Description).UniqueName(UniqueName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
    Description := "Description_example" // string | Description of this Rate Limit (optional)
    UniqueName := "UniqueName_example" // string | Provides a unique and addressable name to be assigned to this Rate Limit, assigned by the developer, to be optionally used in addition to SID. **This value should not contain PII.** (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateRateLimit(context.Background(), ServiceSid).Description(Description).UniqueName(UniqueName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateRateLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRateLimit`: VerifyV2ServiceRateLimit
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateRateLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.

### Other Parameters

Other parameters are passed through a pointer to a CreateRateLimitParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Description** | **string** | Description of this Rate Limit
 **UniqueName** | **string** | Provides a unique and addressable name to be assigned to this Rate Limit, assigned by the developer, to be optionally used in addition to SID. **This value should not contain PII.**

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

> VerifyV2Service CreateService(ctx).CodeLength(CodeLength).CustomCodeEnabled(CustomCodeEnabled).DoNotShareWarningEnabled(DoNotShareWarningEnabled).DtmfInputRequired(DtmfInputRequired).FriendlyName(FriendlyName).LookupEnabled(LookupEnabled).Psd2Enabled(Psd2Enabled).PushApnCredentialSid(PushApnCredentialSid).PushFcmCredentialSid(PushFcmCredentialSid).PushIncludeDate(PushIncludeDate).SkipSmsToLandlines(SkipSmsToLandlines).TtsName(TtsName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    CodeLength := int32(56) // int32 | The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive. (optional)
    CustomCodeEnabled := true // bool | Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers. (optional)
    DoNotShareWarningEnabled := true // bool | Whether to add a security warning at the end of an SMS verification body. Disabled by default and applies only to SMS. Example SMS body: `Your AppName verification code is: 1234. Don’t share this code with anyone; our employees will never ask for the code` (optional)
    DtmfInputRequired := true // bool | Whether to ask the user to press a number before delivering the verify code in a phone call. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the verification service. It can be up to 30 characters long. **This value should not contain PII.** (optional)
    LookupEnabled := true // bool | Whether to perform a lookup with each verification started and return info about the phone number. (optional)
    Psd2Enabled := true // bool | Whether to pass PSD2 transaction parameters when starting a verification. (optional)
    PushApnCredentialSid := "PushApnCredentialSid_example" // string | Optional configuration for the Push factors. Set the APN Credential for this service. This will allow to send push notifications to iOS devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource) (optional)
    PushFcmCredentialSid := "PushFcmCredentialSid_example" // string | Optional configuration for the Push factors. Set the FCM Credential for this service. This will allow to send push notifications to Android devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource) (optional)
    PushIncludeDate := true // bool | Optional configuration for the Push factors. If true, include the date in the Challenge's reponse. Otherwise, the date is omitted from the response. See [Challenge](https://www.twilio.com/docs/verify/api/challenge) resource’s details parameter for more info. Default: true (optional)
    SkipSmsToLandlines := true // bool | Whether to skip sending SMS verifications to landlines. Requires `lookup_enabled`. (optional)
    TtsName := "TtsName_example" // string | The name of an alternative text-to-speech service to use in phone calls. Applies only to TTS languages. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateService(context.Background()).CodeLength(CodeLength).CustomCodeEnabled(CustomCodeEnabled).DoNotShareWarningEnabled(DoNotShareWarningEnabled).DtmfInputRequired(DtmfInputRequired).FriendlyName(FriendlyName).LookupEnabled(LookupEnabled).Psd2Enabled(Psd2Enabled).PushApnCredentialSid(PushApnCredentialSid).PushFcmCredentialSid(PushFcmCredentialSid).PushIncludeDate(PushIncludeDate).SkipSmsToLandlines(SkipSmsToLandlines).TtsName(TtsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateService`: VerifyV2Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateService`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **CodeLength** | **int32** | The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive.
 **CustomCodeEnabled** | **bool** | Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers.
 **DoNotShareWarningEnabled** | **bool** | Whether to add a security warning at the end of an SMS verification body. Disabled by default and applies only to SMS. Example SMS body: &#x60;Your AppName verification code is: 1234. Don’t share this code with anyone; our employees will never ask for the code&#x60;
 **DtmfInputRequired** | **bool** | Whether to ask the user to press a number before delivering the verify code in a phone call.
 **FriendlyName** | **string** | A descriptive string that you create to describe the verification service. It can be up to 30 characters long. **This value should not contain PII.**
 **LookupEnabled** | **bool** | Whether to perform a lookup with each verification started and return info about the phone number.
 **Psd2Enabled** | **bool** | Whether to pass PSD2 transaction parameters when starting a verification.
 **PushApnCredentialSid** | **string** | Optional configuration for the Push factors. Set the APN Credential for this service. This will allow to send push notifications to iOS devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource)
 **PushFcmCredentialSid** | **string** | Optional configuration for the Push factors. Set the FCM Credential for this service. This will allow to send push notifications to Android devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource)
 **PushIncludeDate** | **bool** | Optional configuration for the Push factors. If true, include the date in the Challenge&#39;s reponse. Otherwise, the date is omitted from the response. See [Challenge](https://www.twilio.com/docs/verify/api/challenge) resource’s details parameter for more info. Default: true
 **SkipSmsToLandlines** | **bool** | Whether to skip sending SMS verifications to landlines. Requires &#x60;lookup_enabled&#x60;.
 **TtsName** | **string** | The name of an alternative text-to-speech service to use in phone calls. Applies only to TTS languages.

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

> VerifyV2ServiceVerification CreateVerification(ctx, ServiceSid).Amount(Amount).AppHash(AppHash).Channel(Channel).ChannelConfiguration(ChannelConfiguration).CustomCode(CustomCode).CustomFriendlyName(CustomFriendlyName).CustomMessage(CustomMessage).Locale(Locale).Payee(Payee).RateLimits(RateLimits).SendDigits(SendDigits).To(To).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to create the resource under.
    Amount := "Amount_example" // string | The amount of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled. (optional)
    AppHash := "AppHash_example" // string | Your [App Hash](https://developers.google.com/identity/sms-retriever/verify#computing_your_apps_hash_string) to be appended at the end of your verification SMS body. Applies only to SMS. Example SMS body: `<#> Your AppName verification code is: 1234 He42w354ol9`. (optional)
    Channel := "Channel_example" // string | The verification method to use. Can be: [`email`](https://www.twilio.com/docs/verify/email), `sms` or `call`. (optional)
    ChannelConfiguration := TODO // map[string]interface{} | [`email`](https://www.twilio.com/docs/verify/email) channel configuration in json format. Must include 'from' and 'from_name'. (optional)
    CustomCode := "CustomCode_example" // string | A pre-generated code to use for verification. The code can be between 4 and 10 characters, inclusive. (optional)
    CustomFriendlyName := "CustomFriendlyName_example" // string | A custom user defined friendly name that overwrites the existing one in the verification message (optional)
    CustomMessage := "CustomMessage_example" // string | The text of a custom message to use for the verification. (optional)
    Locale := "Locale_example" // string | The locale to use for the verification SMS or call. Can be: `af`, `ar`, `ca`, `cs`, `da`, `de`, `el`, `en`, `en-GB`, `es`, `fi`, `fr`, `he`, `hi`, `hr`, `hu`, `id`, `it`, `ja`, `ko`, `ms`, `nb`, `nl`, `pl`, `pt`, `pr-BR`, `ro`, `ru`, `sv`, `th`, `tl`, `tr`, `vi`, `zh`, `zh-CN`, or `zh-HK.` (optional)
    Payee := "Payee_example" // string | The payee of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled. (optional)
    RateLimits := TODO // map[string]interface{} | The custom key-value pairs of Programmable Rate Limits. Keys correspond to `unique_name` fields defined when [creating your Rate Limit](https://www.twilio.com/docs/verify/api/service-rate-limits). Associated value pairs represent values in the request that you are rate limiting on. You may include multiple Rate Limit values in each request. (optional)
    SendDigits := "SendDigits_example" // string | The digits to send after a phone call is answered, for example, to dial an extension. For more information, see the Programmable Voice documentation of [sendDigits](https://www.twilio.com/docs/voice/twiml/number#attributes-sendDigits). (optional)
    To := "To_example" // string | The phone number or [email](https://www.twilio.com/docs/verify/email) to verify. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164). (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateVerification(context.Background(), ServiceSid).Amount(Amount).AppHash(AppHash).Channel(Channel).ChannelConfiguration(ChannelConfiguration).CustomCode(CustomCode).CustomFriendlyName(CustomFriendlyName).CustomMessage(CustomMessage).Locale(Locale).Payee(Payee).RateLimits(RateLimits).SendDigits(SendDigits).To(To).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateVerification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVerification`: VerifyV2ServiceVerification
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to create the resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateVerificationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Amount** | **string** | The amount of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.
 **AppHash** | **string** | Your [App Hash](https://developers.google.com/identity/sms-retriever/verify#computing_your_apps_hash_string) to be appended at the end of your verification SMS body. Applies only to SMS. Example SMS body: &#x60;&lt;#&gt; Your AppName verification code is: 1234 He42w354ol9&#x60;.
 **Channel** | **string** | The verification method to use. Can be: [&#x60;email&#x60;](https://www.twilio.com/docs/verify/email), &#x60;sms&#x60; or &#x60;call&#x60;.
 **ChannelConfiguration** | [**map[string]interface{}**](map[string]interface{}.md) | [&#x60;email&#x60;](https://www.twilio.com/docs/verify/email) channel configuration in json format. Must include &#39;from&#39; and &#39;from_name&#39;.
 **CustomCode** | **string** | A pre-generated code to use for verification. The code can be between 4 and 10 characters, inclusive.
 **CustomFriendlyName** | **string** | A custom user defined friendly name that overwrites the existing one in the verification message
 **CustomMessage** | **string** | The text of a custom message to use for the verification.
 **Locale** | **string** | The locale to use for the verification SMS or call. Can be: &#x60;af&#x60;, &#x60;ar&#x60;, &#x60;ca&#x60;, &#x60;cs&#x60;, &#x60;da&#x60;, &#x60;de&#x60;, &#x60;el&#x60;, &#x60;en&#x60;, &#x60;en-GB&#x60;, &#x60;es&#x60;, &#x60;fi&#x60;, &#x60;fr&#x60;, &#x60;he&#x60;, &#x60;hi&#x60;, &#x60;hr&#x60;, &#x60;hu&#x60;, &#x60;id&#x60;, &#x60;it&#x60;, &#x60;ja&#x60;, &#x60;ko&#x60;, &#x60;ms&#x60;, &#x60;nb&#x60;, &#x60;nl&#x60;, &#x60;pl&#x60;, &#x60;pt&#x60;, &#x60;pr-BR&#x60;, &#x60;ro&#x60;, &#x60;ru&#x60;, &#x60;sv&#x60;, &#x60;th&#x60;, &#x60;tl&#x60;, &#x60;tr&#x60;, &#x60;vi&#x60;, &#x60;zh&#x60;, &#x60;zh-CN&#x60;, or &#x60;zh-HK.&#x60;
 **Payee** | **string** | The payee of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.
 **RateLimits** | [**map[string]interface{}**](map[string]interface{}.md) | The custom key-value pairs of Programmable Rate Limits. Keys correspond to &#x60;unique_name&#x60; fields defined when [creating your Rate Limit](https://www.twilio.com/docs/verify/api/service-rate-limits). Associated value pairs represent values in the request that you are rate limiting on. You may include multiple Rate Limit values in each request.
 **SendDigits** | **string** | The digits to send after a phone call is answered, for example, to dial an extension. For more information, see the Programmable Voice documentation of [sendDigits](https://www.twilio.com/docs/voice/twiml/number#attributes-sendDigits).
 **To** | **string** | The phone number or [email](https://www.twilio.com/docs/verify/email) to verify. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).

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

> VerifyV2ServiceVerificationCheck CreateVerificationCheck(ctx, ServiceSid).Amount(Amount).Code(Code).Payee(Payee).To(To).VerificationSid(VerificationSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to create the resource under.
    Amount := "Amount_example" // string | The amount of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled. (optional)
    Code := "Code_example" // string | The 4-10 character string being verified. (optional)
    Payee := "Payee_example" // string | The payee of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled. (optional)
    To := "To_example" // string | The phone number or [email](https://www.twilio.com/docs/verify/email) to verify. Either this parameter or the `verification_sid` must be specified. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164). (optional)
    VerificationSid := "VerificationSid_example" // string | A SID that uniquely identifies the Verification Check. Either this parameter or the `to` phone number/[email](https://www.twilio.com/docs/verify/email) must be specified. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateVerificationCheck(context.Background(), ServiceSid).Amount(Amount).Code(Code).Payee(Payee).To(To).VerificationSid(VerificationSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateVerificationCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVerificationCheck`: VerifyV2ServiceVerificationCheck
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateVerificationCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to create the resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateVerificationCheckParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Amount** | **string** | The amount of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.
 **Code** | **string** | The 4-10 character string being verified.
 **Payee** | **string** | The payee of the associated PSD2 compliant transaction. Requires the PSD2 Service flag enabled.
 **To** | **string** | The phone number or [email](https://www.twilio.com/docs/verify/email) to verify. Either this parameter or the &#x60;verification_sid&#x60; must be specified. Phone numbers must be in [E.164 format](https://www.twilio.com/docs/glossary/what-e164).
 **VerificationSid** | **string** | A SID that uniquely identifies the Verification Check. Either this parameter or the &#x60;to&#x60; phone number/[email](https://www.twilio.com/docs/verify/email) must be specified.

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

> VerifyV2ServiceWebhook CreateWebhook(ctx, ServiceSid).EventTypes(EventTypes).FriendlyName(FriendlyName).Status(Status).WebhookUrl(WebhookUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The unique SID identifier of the Service.
    EventTypes := []string{"Inner_example"} // []string | The array of events that this Webhook is subscribed to. Possible event types: `*, factor.deleted, factor.created, factor.verified, challenge.approved, challenge.denied` (optional)
    FriendlyName := "FriendlyName_example" // string | The string that you assigned to describe the webhook. **This value should not contain PII.** (optional)
    Status := "Status_example" // string | The webhook status. Default value is `enabled`. One of: `enabled` or `disabled` (optional)
    WebhookUrl := "WebhookUrl_example" // string | The URL associated with this Webhook. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateWebhook(context.Background(), ServiceSid).EventTypes(EventTypes).FriendlyName(FriendlyName).Status(Status).WebhookUrl(WebhookUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebhook`: VerifyV2ServiceWebhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.

### Other Parameters

Other parameters are passed through a pointer to a CreateWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **EventTypes** | **[]string** | The array of events that this Webhook is subscribed to. Possible event types: &#x60;*, factor.deleted, factor.created, factor.verified, challenge.approved, challenge.denied&#x60;
 **FriendlyName** | **string** | The string that you assigned to describe the webhook. **This value should not contain PII.**
 **Status** | **string** | The webhook status. Default value is &#x60;enabled&#x60;. One of: &#x60;enabled&#x60; or &#x60;disabled&#x60;
 **WebhookUrl** | **string** | The URL associated with this Webhook.

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

> DeleteBucket(ctx, ServiceSid, RateLimitSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
    RateLimitSid := "RateLimitSid_example" // string | The Twilio-provided string that uniquely identifies the Rate Limit resource.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this Bucket.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteBucket(context.Background(), ServiceSid, RateLimitSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
**RateLimitSid** | **string** | The Twilio-provided string that uniquely identifies the Rate Limit resource.
**Sid** | **string** | A 34 character string that uniquely identifies this Bucket.

### Other Parameters

Other parameters are passed through a pointer to a DeleteBucketParams struct via the builder pattern


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


## DeleteEntity

> DeleteEntity(ctx, ServiceSid, Identity).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The unique SID identifier of the Service.
    Identity := "Identity_example" // string | The unique external identifier for the Entity of the Service

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteEntity(context.Background(), ServiceSid, Identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Identity** | **string** | The unique external identifier for the Entity of the Service

### Other Parameters

Other parameters are passed through a pointer to a DeleteEntityParams struct via the builder pattern


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


## DeleteFactor

> DeleteFactor(ctx, ServiceSid, Identity, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The unique SID identifier of the Service.
    Identity := "Identity_example" // string | Customer unique identity for the Entity owner of the Factor. This value must be between 8 and 64 characters long.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this Factor.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteFactor(context.Background(), ServiceSid, Identity, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Identity** | **string** | Customer unique identity for the Entity owner of the Factor. This value must be between 8 and 64 characters long.
**Sid** | **string** | A 34 character string that uniquely identifies this Factor.

### Other Parameters

Other parameters are passed through a pointer to a DeleteFactorParams struct via the builder pattern


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


## DeleteMessagingConfiguration

> DeleteMessagingConfiguration(ctx, ServiceSid, Country).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with.
    Country := "Country_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value `all`.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteMessagingConfiguration(context.Background(), ServiceSid, Country).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteMessagingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with.
**Country** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value &#x60;all&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a DeleteMessagingConfigurationParams struct via the builder pattern


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


## DeleteRateLimit

> DeleteRateLimit(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Rate Limit resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteRateLimit(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRateLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Rate Limit resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a DeleteRateLimitParams struct via the builder pattern


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


## DeleteService

> DeleteService(ctx, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Verification Service resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteService(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Verification Service resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteServiceParams struct via the builder pattern


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


## DeleteWebhook

> DeleteWebhook(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The unique SID identifier of the Service.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Webhook resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteWebhook(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Webhook resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteWebhookParams struct via the builder pattern


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


## FetchBucket

> VerifyV2ServiceRateLimitBucket FetchBucket(ctx, ServiceSid, RateLimitSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
    RateLimitSid := "RateLimitSid_example" // string | The Twilio-provided string that uniquely identifies the Rate Limit resource.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this Bucket.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchBucket(context.Background(), ServiceSid, RateLimitSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchBucket`: VerifyV2ServiceRateLimitBucket
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
**RateLimitSid** | **string** | The Twilio-provided string that uniquely identifies the Rate Limit resource.
**Sid** | **string** | A 34 character string that uniquely identifies this Bucket.

### Other Parameters

Other parameters are passed through a pointer to a FetchBucketParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




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

> VerifyV2ServiceEntityChallenge FetchChallenge(ctx, ServiceSid, Identity, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The unique SID identifier of the Service.
    Identity := "Identity_example" // string | The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, and generated by your external system, such as your user's UUID, GUID, or SID. This value must be between 8 and 64 characters long.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this Challenge.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchChallenge(context.Background(), ServiceSid, Identity, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchChallenge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchChallenge`: VerifyV2ServiceEntityChallenge
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchChallenge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Identity** | **string** | The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, and generated by your external system, such as your user&#39;s UUID, GUID, or SID. This value must be between 8 and 64 characters long.
**Sid** | **string** | A 34 character string that uniquely identifies this Challenge.

### Other Parameters

Other parameters are passed through a pointer to a FetchChallengeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------




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

> VerifyV2ServiceEntity FetchEntity(ctx, ServiceSid, Identity).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The unique SID identifier of the Service.
    Identity := "Identity_example" // string | The unique external identifier for the Entity of the Service

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchEntity(context.Background(), ServiceSid, Identity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchEntity`: VerifyV2ServiceEntity
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Identity** | **string** | The unique external identifier for the Entity of the Service

### Other Parameters

Other parameters are passed through a pointer to a FetchEntityParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



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

> VerifyV2ServiceEntityFactor FetchFactor(ctx, ServiceSid, Identity, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The unique SID identifier of the Service.
    Identity := "Identity_example" // string | Customer unique identity for the Entity owner of the Factor. This value must be between 8 and 64 characters long.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this Factor.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchFactor(context.Background(), ServiceSid, Identity, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchFactor`: VerifyV2ServiceEntityFactor
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Identity** | **string** | Customer unique identity for the Entity owner of the Factor. This value must be between 8 and 64 characters long.
**Sid** | **string** | A 34 character string that uniquely identifies this Factor.

### Other Parameters

Other parameters are passed through a pointer to a FetchFactorParams struct via the builder pattern


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


## FetchForm

> VerifyV2Form FetchForm(ctx, FormType).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    FormType := "FormType_example" // string | The Type of this Form. Currently only `form-push` is supported.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchForm(context.Background(), FormType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchForm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchForm`: VerifyV2Form
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**FormType** | **string** | The Type of this Form. Currently only &#x60;form-push&#x60; is supported.

### Other Parameters

Other parameters are passed through a pointer to a FetchFormParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


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

> VerifyV2ServiceMessagingConfiguration FetchMessagingConfiguration(ctx, ServiceSid, Country).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with.
    Country := "Country_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value `all`.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchMessagingConfiguration(context.Background(), ServiceSid, Country).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchMessagingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchMessagingConfiguration`: VerifyV2ServiceMessagingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchMessagingConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with.
**Country** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value &#x60;all&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a FetchMessagingConfigurationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



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

> VerifyV2ServiceRateLimit FetchRateLimit(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Rate Limit resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchRateLimit(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRateLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRateLimit`: VerifyV2ServiceRateLimit
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRateLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Rate Limit resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchRateLimitParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



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

> VerifyV2Service FetchService(ctx, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Verification Service resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchService(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchService`: VerifyV2Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Verification Service resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


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

> VerifyV2ServiceVerification FetchVerification(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to fetch the resource from.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Verification resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchVerification(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchVerification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchVerification`: VerifyV2ServiceVerification
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to fetch the resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Verification resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchVerificationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



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

> VerifyV2VerificationAttempt FetchVerificationAttempt(ctx, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The unique SID identifier of a Verification Attempt

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchVerificationAttempt(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchVerificationAttempt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchVerificationAttempt`: VerifyV2VerificationAttempt
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchVerificationAttempt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique SID identifier of a Verification Attempt

### Other Parameters

Other parameters are passed through a pointer to a FetchVerificationAttemptParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


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

> VerifyV2ServiceWebhook FetchWebhook(ctx, ServiceSid, Sid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The unique SID identifier of the Service.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Webhook resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchWebhook(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchWebhook`: VerifyV2ServiceWebhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Webhook resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



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

> ListBucketResponse ListBucket(ctx, ServiceSid, RateLimitSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
    RateLimitSid := "RateLimitSid_example" // string | The Twilio-provided string that uniquely identifies the Rate Limit resource.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListBucket(context.Background(), ServiceSid, RateLimitSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBucket`: ListBucketResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
**RateLimitSid** | **string** | The Twilio-provided string that uniquely identifies the Rate Limit resource.

### Other Parameters

Other parameters are passed through a pointer to a ListBucketParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListChallengeResponse ListChallenge(ctx, ServiceSid, Identity).FactorSid(FactorSid).Status(Status).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The unique SID identifier of the Service.
    Identity := "Identity_example" // string | Customer unique identity for the Entity owner of the Challenge. This value must be between 8 and 64 characters long.
    FactorSid := "FactorSid_example" // string | The unique SID identifier of the Factor. (optional)
    Status := "Status_example" // string | The Status of the Challenges to fetch. One of `pending`, `expired`, `approved` or `denied`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListChallenge(context.Background(), ServiceSid, Identity).FactorSid(FactorSid).Status(Status).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListChallenge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListChallenge`: ListChallengeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListChallenge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Identity** | **string** | Customer unique identity for the Entity owner of the Challenge. This value must be between 8 and 64 characters long.

### Other Parameters

Other parameters are passed through a pointer to a ListChallengeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **FactorSid** | **string** | The unique SID identifier of the Factor.
 **Status** | **string** | The Status of the Challenges to fetch. One of &#x60;pending&#x60;, &#x60;expired&#x60;, &#x60;approved&#x60; or &#x60;denied&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListEntityResponse ListEntity(ctx, ServiceSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The unique SID identifier of the Service.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListEntity(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListEntity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEntity`: ListEntityResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListEntity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.

### Other Parameters

Other parameters are passed through a pointer to a ListEntityParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListFactorResponse ListFactor(ctx, ServiceSid, Identity).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The unique SID identifier of the Service.
    Identity := "Identity_example" // string | Customer unique identity for the Entity owner of the Factor. This value must be between 8 and 64 characters long.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListFactor(context.Background(), ServiceSid, Identity).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFactor`: ListFactorResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Identity** | **string** | Customer unique identity for the Entity owner of the Factor. This value must be between 8 and 64 characters long.

### Other Parameters

Other parameters are passed through a pointer to a ListFactorParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListMessagingConfigurationResponse ListMessagingConfiguration(ctx, ServiceSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListMessagingConfiguration(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMessagingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMessagingConfiguration`: ListMessagingConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMessagingConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with.

### Other Parameters

Other parameters are passed through a pointer to a ListMessagingConfigurationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListRateLimitResponse ListRateLimit(ctx, ServiceSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListRateLimit(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRateLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRateLimit`: ListRateLimitResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRateLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.

### Other Parameters

Other parameters are passed through a pointer to a ListRateLimitParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListServiceResponse ListService(ctx).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListService(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListService`: ListServiceResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListService`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListServiceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListVerificationAttemptResponse ListVerificationAttempt(ctx).DateCreatedAfter(DateCreatedAfter).DateCreatedBefore(DateCreatedBefore).ChannelDataTo(ChannelDataTo).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    DateCreatedAfter := time.Now() // time.Time | Datetime filter used to query Verification Attempts created after this datetime. (optional)
    DateCreatedBefore := time.Now() // time.Time | Datetime filter used to query Verification Attempts created before this datetime. (optional)
    ChannelDataTo := "ChannelDataTo_example" // string | Destination of a verification. Depending on the type of channel, it could be a phone number in E.164 format or an email address. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListVerificationAttempt(context.Background()).DateCreatedAfter(DateCreatedAfter).DateCreatedBefore(DateCreatedBefore).ChannelDataTo(ChannelDataTo).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListVerificationAttempt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVerificationAttempt`: ListVerificationAttemptResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListVerificationAttempt`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListVerificationAttemptParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **DateCreatedAfter** | **time.Time** | Datetime filter used to query Verification Attempts created after this datetime.
 **DateCreatedBefore** | **time.Time** | Datetime filter used to query Verification Attempts created before this datetime.
 **ChannelDataTo** | **string** | Destination of a verification. Depending on the type of channel, it could be a phone number in E.164 format or an email address.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> ListWebhookResponse ListWebhook(ctx, ServiceSid).PageSize(PageSize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The unique SID identifier of the Service.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListWebhook(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebhook`: ListWebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.

### Other Parameters

Other parameters are passed through a pointer to a ListWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

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

> VerifyV2ServiceRateLimitBucket UpdateBucket(ctx, ServiceSid, RateLimitSid, Sid).Interval(Interval).Max(Max).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
    RateLimitSid := "RateLimitSid_example" // string | The Twilio-provided string that uniquely identifies the Rate Limit resource.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this Bucket.
    Interval := int32(56) // int32 | Number of seconds that the rate limit will be enforced over. (optional)
    Max := int32(56) // int32 | Maximum number of requests permitted in during the interval. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateBucket(context.Background(), ServiceSid, RateLimitSid, Sid).Interval(Interval).Max(Max).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBucket`: VerifyV2ServiceRateLimitBucket
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
**RateLimitSid** | **string** | The Twilio-provided string that uniquely identifies the Rate Limit resource.
**Sid** | **string** | A 34 character string that uniquely identifies this Bucket.

### Other Parameters

Other parameters are passed through a pointer to a UpdateBucketParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **Interval** | **int32** | Number of seconds that the rate limit will be enforced over.
 **Max** | **int32** | Maximum number of requests permitted in during the interval.

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

> VerifyV2ServiceEntityChallenge UpdateChallenge(ctx, ServiceSid, Identity, Sid).AuthPayload(AuthPayload).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The unique SID identifier of the Service.
    Identity := "Identity_example" // string | Customer unique identity for the Entity owner of the Challenge. This value must be between 8 and 64 characters long.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this Challenge.
    AuthPayload := "AuthPayload_example" // string | The optional payload needed to verify the Challenge. E.g., a TOTP would use the numeric code. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateChallenge(context.Background(), ServiceSid, Identity, Sid).AuthPayload(AuthPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateChallenge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateChallenge`: VerifyV2ServiceEntityChallenge
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateChallenge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Identity** | **string** | Customer unique identity for the Entity owner of the Challenge. This value must be between 8 and 64 characters long.
**Sid** | **string** | A 34 character string that uniquely identifies this Challenge.

### Other Parameters

Other parameters are passed through a pointer to a UpdateChallengeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **AuthPayload** | **string** | The optional payload needed to verify the Challenge. E.g., a TOTP would use the numeric code.

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

> VerifyV2ServiceEntityFactor UpdateFactor(ctx, ServiceSid, Identity, Sid).AuthPayload(AuthPayload).ConfigAlg(ConfigAlg).ConfigCodeLength(ConfigCodeLength).ConfigNotificationToken(ConfigNotificationToken).ConfigSdkVersion(ConfigSdkVersion).ConfigSkew(ConfigSkew).ConfigTimeStep(ConfigTimeStep).FriendlyName(FriendlyName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The unique SID identifier of the Service.
    Identity := "Identity_example" // string | Customer unique identity for the Entity owner of the Factor. This value must be between 8 and 64 characters long.
    Sid := "Sid_example" // string | A 34 character string that uniquely identifies this Factor.
    AuthPayload := "AuthPayload_example" // string | The optional payload needed to verify the Factor for the first time. E.g. for a TOTP, the numeric code. (optional)
    ConfigAlg := "ConfigAlg_example" // string | The algorithm used to derive the TOTP codes. Can be `sha1`, `sha256` or `sha512` (optional)
    ConfigCodeLength := int32(56) // int32 | Number of digits for generated TOTP codes. Must be between 3 and 8, inclusive (optional)
    ConfigNotificationToken := "ConfigNotificationToken_example" // string | For APN, the device token. For FCM the registration token. It used to send the push notifications. Required when `factor_type` is `push`. If specified, this value must be between 32 and 255 characters long. (optional)
    ConfigSdkVersion := "ConfigSdkVersion_example" // string | The Verify Push SDK version used to configure the factor (optional)
    ConfigSkew := int32(56) // int32 | The number of time-steps, past and future, that are valid for validation of TOTP codes. Must be between 0 and 2, inclusive (optional)
    ConfigTimeStep := int32(56) // int32 | Defines how often, in seconds, are TOTP codes generated. i.e, a new TOTP code is generated every time_step seconds. Must be between 20 and 60 seconds, inclusive (optional)
    FriendlyName := "FriendlyName_example" // string | The new friendly name of this Factor. It can be up to 64 characters. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateFactor(context.Background(), ServiceSid, Identity, Sid).AuthPayload(AuthPayload).ConfigAlg(ConfigAlg).ConfigCodeLength(ConfigCodeLength).ConfigNotificationToken(ConfigNotificationToken).ConfigSdkVersion(ConfigSdkVersion).ConfigSkew(ConfigSkew).ConfigTimeStep(ConfigTimeStep).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateFactor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFactor`: VerifyV2ServiceEntityFactor
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Identity** | **string** | Customer unique identity for the Entity owner of the Factor. This value must be between 8 and 64 characters long.
**Sid** | **string** | A 34 character string that uniquely identifies this Factor.

### Other Parameters

Other parameters are passed through a pointer to a UpdateFactorParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



 **AuthPayload** | **string** | The optional payload needed to verify the Factor for the first time. E.g. for a TOTP, the numeric code.
 **ConfigAlg** | **string** | The algorithm used to derive the TOTP codes. Can be &#x60;sha1&#x60;, &#x60;sha256&#x60; or &#x60;sha512&#x60;
 **ConfigCodeLength** | **int32** | Number of digits for generated TOTP codes. Must be between 3 and 8, inclusive
 **ConfigNotificationToken** | **string** | For APN, the device token. For FCM the registration token. It used to send the push notifications. Required when &#x60;factor_type&#x60; is &#x60;push&#x60;. If specified, this value must be between 32 and 255 characters long.
 **ConfigSdkVersion** | **string** | The Verify Push SDK version used to configure the factor
 **ConfigSkew** | **int32** | The number of time-steps, past and future, that are valid for validation of TOTP codes. Must be between 0 and 2, inclusive
 **ConfigTimeStep** | **int32** | Defines how often, in seconds, are TOTP codes generated. i.e, a new TOTP code is generated every time_step seconds. Must be between 20 and 60 seconds, inclusive
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


## UpdateMessagingConfiguration

> VerifyV2ServiceMessagingConfiguration UpdateMessagingConfiguration(ctx, ServiceSid, Country).MessagingServiceSid(MessagingServiceSid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with.
    Country := "Country_example" // string | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value `all`.
    MessagingServiceSid := "MessagingServiceSid_example" // string | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) to be used to send SMS to the country of this configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateMessagingConfiguration(context.Background(), ServiceSid, Country).MessagingServiceSid(MessagingServiceSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateMessagingConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMessagingConfiguration`: VerifyV2ServiceMessagingConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateMessagingConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) that the resource is associated with.
**Country** | **string** | The [ISO-3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) country code of the country this configuration will be applied to. If this is a global configuration, Country will take the value &#x60;all&#x60;.

### Other Parameters

Other parameters are passed through a pointer to a UpdateMessagingConfigurationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/services/api) to be used to send SMS to the country of this configuration.

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

> VerifyV2ServiceRateLimit UpdateRateLimit(ctx, ServiceSid, Sid).Description(Description).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Rate Limit resource to fetch.
    Description := "Description_example" // string | Description of this Rate Limit (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateRateLimit(context.Background(), ServiceSid, Sid).Description(Description).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateRateLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRateLimit`: VerifyV2ServiceRateLimit
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateRateLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Rate Limit resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRateLimitParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Description** | **string** | Description of this Rate Limit

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

> VerifyV2Service UpdateService(ctx, Sid).CodeLength(CodeLength).CustomCodeEnabled(CustomCodeEnabled).DoNotShareWarningEnabled(DoNotShareWarningEnabled).DtmfInputRequired(DtmfInputRequired).FriendlyName(FriendlyName).LookupEnabled(LookupEnabled).Psd2Enabled(Psd2Enabled).PushApnCredentialSid(PushApnCredentialSid).PushFcmCredentialSid(PushFcmCredentialSid).PushIncludeDate(PushIncludeDate).SkipSmsToLandlines(SkipSmsToLandlines).TtsName(TtsName).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Service resource to update.
    CodeLength := int32(56) // int32 | The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive. (optional)
    CustomCodeEnabled := true // bool | Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers. (optional)
    DoNotShareWarningEnabled := true // bool | Whether to add a privacy warning at the end of an SMS. **Disabled by default and applies only for SMS.** (optional)
    DtmfInputRequired := true // bool | Whether to ask the user to press a number before delivering the verify code in a phone call. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the verification service. It can be up to 30 characters long. **This value should not contain PII.** (optional)
    LookupEnabled := true // bool | Whether to perform a lookup with each verification started and return info about the phone number. (optional)
    Psd2Enabled := true // bool | Whether to pass PSD2 transaction parameters when starting a verification. (optional)
    PushApnCredentialSid := "PushApnCredentialSid_example" // string | Optional configuration for the Push factors. Set the APN Credential for this service. This will allow to send push notifications to iOS devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource) (optional)
    PushFcmCredentialSid := "PushFcmCredentialSid_example" // string | Optional configuration for the Push factors. Set the FCM Credential for this service. This will allow to send push notifications to Android devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource) (optional)
    PushIncludeDate := true // bool | Optional configuration for the Push factors. If true, include the date in the Challenge's reponse. Otherwise, the date is omitted from the response. See [Challenge](https://www.twilio.com/docs/verify/api/challenge) resource’s details parameter for more info. Default: true (optional)
    SkipSmsToLandlines := true // bool | Whether to skip sending SMS verifications to landlines. Requires `lookup_enabled`. (optional)
    TtsName := "TtsName_example" // string | The name of an alternative text-to-speech service to use in phone calls. Applies only to TTS languages. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateService(context.Background(), Sid).CodeLength(CodeLength).CustomCodeEnabled(CustomCodeEnabled).DoNotShareWarningEnabled(DoNotShareWarningEnabled).DtmfInputRequired(DtmfInputRequired).FriendlyName(FriendlyName).LookupEnabled(LookupEnabled).Psd2Enabled(Psd2Enabled).PushApnCredentialSid(PushApnCredentialSid).PushFcmCredentialSid(PushFcmCredentialSid).PushIncludeDate(PushIncludeDate).SkipSmsToLandlines(SkipSmsToLandlines).TtsName(TtsName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateService`: VerifyV2Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Service resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **CodeLength** | **int32** | The length of the verification code to generate. Must be an integer value between 4 and 10, inclusive.
 **CustomCodeEnabled** | **bool** | Whether to allow sending verifications with a custom code instead of a randomly generated one. Not available for all customers.
 **DoNotShareWarningEnabled** | **bool** | Whether to add a privacy warning at the end of an SMS. **Disabled by default and applies only for SMS.**
 **DtmfInputRequired** | **bool** | Whether to ask the user to press a number before delivering the verify code in a phone call.
 **FriendlyName** | **string** | A descriptive string that you create to describe the verification service. It can be up to 30 characters long. **This value should not contain PII.**
 **LookupEnabled** | **bool** | Whether to perform a lookup with each verification started and return info about the phone number.
 **Psd2Enabled** | **bool** | Whether to pass PSD2 transaction parameters when starting a verification.
 **PushApnCredentialSid** | **string** | Optional configuration for the Push factors. Set the APN Credential for this service. This will allow to send push notifications to iOS devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource)
 **PushFcmCredentialSid** | **string** | Optional configuration for the Push factors. Set the FCM Credential for this service. This will allow to send push notifications to Android devices. See [Credential Resource](https://www.twilio.com/docs/notify/api/credential-resource)
 **PushIncludeDate** | **bool** | Optional configuration for the Push factors. If true, include the date in the Challenge&#39;s reponse. Otherwise, the date is omitted from the response. See [Challenge](https://www.twilio.com/docs/verify/api/challenge) resource’s details parameter for more info. Default: true
 **SkipSmsToLandlines** | **bool** | Whether to skip sending SMS verifications to landlines. Requires &#x60;lookup_enabled&#x60;.
 **TtsName** | **string** | The name of an alternative text-to-speech service to use in phone calls. Applies only to TTS languages.

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

> VerifyV2ServiceVerification UpdateVerification(ctx, ServiceSid, Sid).Status(Status).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to update the resource from.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Verification resource to update.
    Status := "Status_example" // string | The new status of the resource. Can be: `canceled` or `approved`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateVerification(context.Background(), ServiceSid, Sid).Status(Status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateVerification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVerification`: VerifyV2ServiceVerification
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the verification [Service](https://www.twilio.com/docs/verify/api/service) to update the resource from.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Verification resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateVerificationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Status** | **string** | The new status of the resource. Can be: &#x60;canceled&#x60; or &#x60;approved&#x60;.

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

> VerifyV2ServiceWebhook UpdateWebhook(ctx, ServiceSid, Sid).EventTypes(EventTypes).FriendlyName(FriendlyName).Status(Status).WebhookUrl(WebhookUrl).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ServiceSid := "ServiceSid_example" // string | The unique SID identifier of the Service.
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the Webhook resource to update.
    EventTypes := []string{"Inner_example"} // []string | The array of events that this Webhook is subscribed to. Possible event types: `*, factor.deleted, factor.created, factor.verified, challenge.approved, challenge.denied` (optional)
    FriendlyName := "FriendlyName_example" // string | The string that you assigned to describe the webhook. **This value should not contain PII.** (optional)
    Status := "Status_example" // string | The webhook status. Default value is `enabled`. One of: `enabled` or `disabled` (optional)
    WebhookUrl := "WebhookUrl_example" // string | The URL associated with this Webhook. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateWebhook(context.Background(), ServiceSid, Sid).EventTypes(EventTypes).FriendlyName(FriendlyName).Status(Status).WebhookUrl(WebhookUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhook`: VerifyV2ServiceWebhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Webhook resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateWebhookParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **EventTypes** | **[]string** | The array of events that this Webhook is subscribed to. Possible event types: &#x60;*, factor.deleted, factor.created, factor.verified, challenge.approved, challenge.denied&#x60;
 **FriendlyName** | **string** | The string that you assigned to describe the webhook. **This value should not contain PII.**
 **Status** | **string** | The webhook status. Default value is &#x60;enabled&#x60;. One of: &#x60;enabled&#x60; or &#x60;disabled&#x60;
 **WebhookUrl** | **string** | The URL associated with this Webhook.

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

