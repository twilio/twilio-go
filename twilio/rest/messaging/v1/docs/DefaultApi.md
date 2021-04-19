# \DefaultApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlphaSender**](DefaultApi.md#CreateAlphaSender) | **Post** /v1/Services/{ServiceSid}/AlphaSenders | 
[**CreateBrandRegistrations**](DefaultApi.md#CreateBrandRegistrations) | **Post** /v1/a2p/BrandRegistrations | 
[**CreateExternalCampaign**](DefaultApi.md#CreateExternalCampaign) | **Post** /v1/Services/PreregisteredUsa2p | 
[**CreatePhoneNumber**](DefaultApi.md#CreatePhoneNumber) | **Post** /v1/Services/{ServiceSid}/PhoneNumbers | 
[**CreateService**](DefaultApi.md#CreateService) | **Post** /v1/Services | 
[**CreateShortCode**](DefaultApi.md#CreateShortCode) | **Post** /v1/Services/{ServiceSid}/ShortCodes | 
[**CreateUsAppToPerson**](DefaultApi.md#CreateUsAppToPerson) | **Post** /v1/Services/{MessagingServiceSid}/Compliance/Usa2p | 
[**DeleteAlphaSender**](DefaultApi.md#DeleteAlphaSender) | **Delete** /v1/Services/{ServiceSid}/AlphaSenders/{Sid} | 
[**DeletePhoneNumber**](DefaultApi.md#DeletePhoneNumber) | **Delete** /v1/Services/{ServiceSid}/PhoneNumbers/{Sid} | 
[**DeleteService**](DefaultApi.md#DeleteService) | **Delete** /v1/Services/{Sid} | 
[**DeleteShortCode**](DefaultApi.md#DeleteShortCode) | **Delete** /v1/Services/{ServiceSid}/ShortCodes/{Sid} | 
[**DeleteUsAppToPerson**](DefaultApi.md#DeleteUsAppToPerson) | **Delete** /v1/Services/{MessagingServiceSid}/Compliance/Usa2p | 
[**FetchAlphaSender**](DefaultApi.md#FetchAlphaSender) | **Get** /v1/Services/{ServiceSid}/AlphaSenders/{Sid} | 
[**FetchBrandRegistrations**](DefaultApi.md#FetchBrandRegistrations) | **Get** /v1/a2p/BrandRegistrations/{Sid} | 
[**FetchDeactivation**](DefaultApi.md#FetchDeactivation) | **Get** /v1/Deactivations | 
[**FetchPhoneNumber**](DefaultApi.md#FetchPhoneNumber) | **Get** /v1/Services/{ServiceSid}/PhoneNumbers/{Sid} | 
[**FetchService**](DefaultApi.md#FetchService) | **Get** /v1/Services/{Sid} | 
[**FetchShortCode**](DefaultApi.md#FetchShortCode) | **Get** /v1/Services/{ServiceSid}/ShortCodes/{Sid} | 
[**FetchUsAppToPerson**](DefaultApi.md#FetchUsAppToPerson) | **Get** /v1/Services/{MessagingServiceSid}/Compliance/Usa2p | 
[**FetchUsAppToPersonUsecase**](DefaultApi.md#FetchUsAppToPersonUsecase) | **Get** /v1/Services/{MessagingServiceSid}/Compliance/Usa2p/Usecases | 
[**FetchUsecase**](DefaultApi.md#FetchUsecase) | **Get** /v1/Services/Usecases | 
[**ListAlphaSender**](DefaultApi.md#ListAlphaSender) | **Get** /v1/Services/{ServiceSid}/AlphaSenders | 
[**ListBrandRegistrations**](DefaultApi.md#ListBrandRegistrations) | **Get** /v1/a2p/BrandRegistrations | 
[**ListPhoneNumber**](DefaultApi.md#ListPhoneNumber) | **Get** /v1/Services/{ServiceSid}/PhoneNumbers | 
[**ListService**](DefaultApi.md#ListService) | **Get** /v1/Services | 
[**ListShortCode**](DefaultApi.md#ListShortCode) | **Get** /v1/Services/{ServiceSid}/ShortCodes | 
[**UpdateService**](DefaultApi.md#UpdateService) | **Post** /v1/Services/{Sid} | 



## CreateAlphaSender

> MessagingV1ServiceAlphaSender CreateAlphaSender(ctx, ServiceSid).AlphaSender(AlphaSender).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under.
    AlphaSender := "AlphaSender_example" // string | The Alphanumeric Sender ID string. Can be up to 11 characters long. Valid characters are A-Z, a-z, 0-9, space, and hyphen `-`. This value cannot contain only numbers. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateAlphaSender(context.Background(), ServiceSid).AlphaSender(AlphaSender).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAlphaSender``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlphaSender`: MessagingV1ServiceAlphaSender
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAlphaSender`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateAlphaSenderParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **AlphaSender** | **string** | The Alphanumeric Sender ID string. Can be up to 11 characters long. Valid characters are A-Z, a-z, 0-9, space, and hyphen &#x60;-&#x60;. This value cannot contain only numbers.

### Return type

[**MessagingV1ServiceAlphaSender**](MessagingV1ServiceAlphaSender.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBrandRegistrations

> MessagingV1BrandRegistrations CreateBrandRegistrations(ctx).A2pProfileBundleSid(A2pProfileBundleSid).CustomerProfileBundleSid(CustomerProfileBundleSid).Execute()



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
    A2pProfileBundleSid := "A2pProfileBundleSid_example" // string | A2P Messaging Profile Bundle Sid. (optional)
    CustomerProfileBundleSid := "CustomerProfileBundleSid_example" // string | Customer Profile Bundle Sid. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateBrandRegistrations(context.Background()).A2pProfileBundleSid(A2pProfileBundleSid).CustomerProfileBundleSid(CustomerProfileBundleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateBrandRegistrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBrandRegistrations`: MessagingV1BrandRegistrations
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateBrandRegistrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateBrandRegistrationsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **A2pProfileBundleSid** | **string** | A2P Messaging Profile Bundle Sid.
 **CustomerProfileBundleSid** | **string** | Customer Profile Bundle Sid.

### Return type

[**MessagingV1BrandRegistrations**](MessagingV1BrandRegistrations.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExternalCampaign

> MessagingV1ExternalCampaign CreateExternalCampaign(ctx).CampaignId(CampaignId).MessagingServiceSid(MessagingServiceSid).Execute()



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
    CampaignId := "CampaignId_example" // string | ID of the preregistered campaign. (optional)
    MessagingServiceSid := "MessagingServiceSid_example" // string | The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) that the resource is associated with. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateExternalCampaign(context.Background()).CampaignId(CampaignId).MessagingServiceSid(MessagingServiceSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateExternalCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExternalCampaign`: MessagingV1ExternalCampaign
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateExternalCampaign`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateExternalCampaignParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **CampaignId** | **string** | ID of the preregistered campaign.
 **MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) that the resource is associated with.

### Return type

[**MessagingV1ExternalCampaign**](MessagingV1ExternalCampaign.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePhoneNumber

> MessagingV1ServicePhoneNumber CreatePhoneNumber(ctx, ServiceSid).PhoneNumberSid(PhoneNumberSid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under.
    PhoneNumberSid := "PhoneNumberSid_example" // string | The SID of the Phone Number being added to the Service. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreatePhoneNumber(context.Background(), ServiceSid).PhoneNumberSid(PhoneNumberSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreatePhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePhoneNumber`: MessagingV1ServicePhoneNumber
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreatePhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreatePhoneNumberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PhoneNumberSid** | **string** | The SID of the Phone Number being added to the Service.

### Return type

[**MessagingV1ServicePhoneNumber**](MessagingV1ServicePhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> MessagingV1Service CreateService(ctx).AreaCodeGeomatch(AreaCodeGeomatch).FallbackMethod(FallbackMethod).FallbackToLongCode(FallbackToLongCode).FallbackUrl(FallbackUrl).FriendlyName(FriendlyName).InboundMethod(InboundMethod).InboundRequestUrl(InboundRequestUrl).MmsConverter(MmsConverter).ScanMessageContent(ScanMessageContent).SmartEncoding(SmartEncoding).StatusCallback(StatusCallback).StickySender(StickySender).SynchronousValidation(SynchronousValidation).ValidityPeriod(ValidityPeriod).Execute()



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
    AreaCodeGeomatch := true // bool | Whether to enable [Area Code Geomatch](https://www.twilio.com/docs/sms/services#area-code-geomatch) on the Service Instance. (optional)
    FallbackMethod := "FallbackMethod_example" // string | The HTTP method we should use to call `fallback_url`. Can be: `GET` or `POST`. (optional)
    FallbackToLongCode := true // bool | Whether to enable [Fallback to Long Code](https://www.twilio.com/docs/sms/services#fallback-to-long-code) for messages sent through the Service instance. (optional)
    FallbackUrl := "FallbackUrl_example" // string | The URL that we should call using `fallback_method` if an error occurs while retrieving or executing the TwiML from the Inbound Request URL. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)
    InboundMethod := "InboundMethod_example" // string | The HTTP method we should use to call `inbound_request_url`. Can be `GET` or `POST` and the default is `POST`. (optional)
    InboundRequestUrl := "InboundRequestUrl_example" // string | The URL we should call using `inbound_method` when a message is received by any phone number or short code in the Service. When this property is `null`, receiving inbound messages is disabled. (optional)
    MmsConverter := true // bool | Whether to enable the [MMS Converter](https://www.twilio.com/docs/sms/services#mms-converter) for messages sent through the Service instance. (optional)
    ScanMessageContent := "ScanMessageContent_example" // string | Reserved. (optional)
    SmartEncoding := true // bool | Whether to enable [Smart Encoding](https://www.twilio.com/docs/sms/services#smart-encoding) for messages sent through the Service instance. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call to [pass status updates](https://www.twilio.com/docs/sms/api/message-resource#message-status-values) about message delivery. (optional)
    StickySender := true // bool | Whether to enable [Sticky Sender](https://www.twilio.com/docs/sms/services#sticky-sender) on the Service instance. (optional)
    SynchronousValidation := true // bool | Reserved. (optional)
    ValidityPeriod := int32(56) // int32 | How long, in seconds, messages sent from the Service are valid. Can be an integer from `1` to `14,400`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateService(context.Background()).AreaCodeGeomatch(AreaCodeGeomatch).FallbackMethod(FallbackMethod).FallbackToLongCode(FallbackToLongCode).FallbackUrl(FallbackUrl).FriendlyName(FriendlyName).InboundMethod(InboundMethod).InboundRequestUrl(InboundRequestUrl).MmsConverter(MmsConverter).ScanMessageContent(ScanMessageContent).SmartEncoding(SmartEncoding).StatusCallback(StatusCallback).StickySender(StickySender).SynchronousValidation(SynchronousValidation).ValidityPeriod(ValidityPeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateService`: MessagingV1Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateService`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateServiceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **AreaCodeGeomatch** | **bool** | Whether to enable [Area Code Geomatch](https://www.twilio.com/docs/sms/services#area-code-geomatch) on the Service Instance.
 **FallbackMethod** | **string** | The HTTP method we should use to call &#x60;fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **FallbackToLongCode** | **bool** | Whether to enable [Fallback to Long Code](https://www.twilio.com/docs/sms/services#fallback-to-long-code) for messages sent through the Service instance.
 **FallbackUrl** | **string** | The URL that we should call using &#x60;fallback_method&#x60; if an error occurs while retrieving or executing the TwiML from the Inbound Request URL.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
 **InboundMethod** | **string** | The HTTP method we should use to call &#x60;inbound_request_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;.
 **InboundRequestUrl** | **string** | The URL we should call using &#x60;inbound_method&#x60; when a message is received by any phone number or short code in the Service. When this property is &#x60;null&#x60;, receiving inbound messages is disabled.
 **MmsConverter** | **bool** | Whether to enable the [MMS Converter](https://www.twilio.com/docs/sms/services#mms-converter) for messages sent through the Service instance.
 **ScanMessageContent** | **string** | Reserved.
 **SmartEncoding** | **bool** | Whether to enable [Smart Encoding](https://www.twilio.com/docs/sms/services#smart-encoding) for messages sent through the Service instance.
 **StatusCallback** | **string** | The URL we should call to [pass status updates](https://www.twilio.com/docs/sms/api/message-resource#message-status-values) about message delivery.
 **StickySender** | **bool** | Whether to enable [Sticky Sender](https://www.twilio.com/docs/sms/services#sticky-sender) on the Service instance.
 **SynchronousValidation** | **bool** | Reserved.
 **ValidityPeriod** | **int32** | How long, in seconds, messages sent from the Service are valid. Can be an integer from &#x60;1&#x60; to &#x60;14,400&#x60;.

### Return type

[**MessagingV1Service**](MessagingV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShortCode

> MessagingV1ServiceShortCode CreateShortCode(ctx, ServiceSid).ShortCodeSid(ShortCodeSid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under.
    ShortCodeSid := "ShortCodeSid_example" // string | The SID of the ShortCode resource being added to the Service. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateShortCode(context.Background(), ServiceSid).ShortCodeSid(ShortCodeSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateShortCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateShortCode`: MessagingV1ServiceShortCode
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateShortCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to create the resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateShortCodeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **ShortCodeSid** | **string** | The SID of the ShortCode resource being added to the Service.

### Return type

[**MessagingV1ServiceShortCode**](MessagingV1ServiceShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUsAppToPerson

> MessagingV1ServiceUsAppToPerson CreateUsAppToPerson(ctx, MessagingServiceSid).BrandRegistrationSid(BrandRegistrationSid).Description(Description).HasEmbeddedLinks(HasEmbeddedLinks).HasEmbeddedPhone(HasEmbeddedPhone).MessageSamples(MessageSamples).UsAppToPersonUsecase(UsAppToPersonUsecase).Execute()



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
    MessagingServiceSid := "MessagingServiceSid_example" // string | The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) to create the resources from.
    BrandRegistrationSid := "BrandRegistrationSid_example" // string | A2P Brand Registration SID (optional)
    Description := "Description_example" // string | A short description of what this SMS campaign does. (optional)
    HasEmbeddedLinks := true // bool | Indicates that this SMS campaign will send messages that contain links. (optional)
    HasEmbeddedPhone := true // bool | Indicates that this SMS campaign will send messages that contain phone numbers. (optional)
    MessageSamples := []string{"Inner_example"} // []string | Message samples, up to 5 sample messages, <=1024 chars each. (optional)
    UsAppToPersonUsecase := "UsAppToPersonUsecase_example" // string | A2P Campaign Use Case. Examples: [ 2FA, EMERGENCY, MARKETING..] (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateUsAppToPerson(context.Background(), MessagingServiceSid).BrandRegistrationSid(BrandRegistrationSid).Description(Description).HasEmbeddedLinks(HasEmbeddedLinks).HasEmbeddedPhone(HasEmbeddedPhone).MessageSamples(MessageSamples).UsAppToPersonUsecase(UsAppToPersonUsecase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUsAppToPerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUsAppToPerson`: MessagingV1ServiceUsAppToPerson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateUsAppToPerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) to create the resources from.

### Other Parameters

Other parameters are passed through a pointer to a CreateUsAppToPersonParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **BrandRegistrationSid** | **string** | A2P Brand Registration SID
 **Description** | **string** | A short description of what this SMS campaign does.
 **HasEmbeddedLinks** | **bool** | Indicates that this SMS campaign will send messages that contain links.
 **HasEmbeddedPhone** | **bool** | Indicates that this SMS campaign will send messages that contain phone numbers.
 **MessageSamples** | **[]string** | Message samples, up to 5 sample messages, &lt;&#x3D;1024 chars each.
 **UsAppToPersonUsecase** | **string** | A2P Campaign Use Case. Examples: [ 2FA, EMERGENCY, MARKETING..]

### Return type

[**MessagingV1ServiceUsAppToPerson**](MessagingV1ServiceUsAppToPerson.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlphaSender

> DeleteAlphaSender(ctx, ServiceSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the resource from.
    Sid := "Sid_example" // string | The SID of the AlphaSender resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteAlphaSender(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAlphaSender``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the resource from.
**Sid** | **string** | The SID of the AlphaSender resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteAlphaSenderParams struct via the builder pattern


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


## DeletePhoneNumber

> DeletePhoneNumber(ctx, ServiceSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the resource from.
    Sid := "Sid_example" // string | The SID of the PhoneNumber resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeletePhoneNumber(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeletePhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the resource from.
**Sid** | **string** | The SID of the PhoneNumber resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeletePhoneNumberParams struct via the builder pattern


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
    Sid := "Sid_example" // string | The SID of the Service resource to delete.

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
**Sid** | **string** | The SID of the Service resource to delete.

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


## DeleteShortCode

> DeleteShortCode(ctx, ServiceSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the resource from.
    Sid := "Sid_example" // string | The SID of the ShortCode resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteShortCode(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteShortCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to delete the resource from.
**Sid** | **string** | The SID of the ShortCode resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteShortCodeParams struct via the builder pattern


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


## DeleteUsAppToPerson

> DeleteUsAppToPerson(ctx, MessagingServiceSid).Execute()



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
    MessagingServiceSid := "MessagingServiceSid_example" // string | The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) to delete the resource from.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteUsAppToPerson(context.Background(), MessagingServiceSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteUsAppToPerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) to delete the resource from.

### Other Parameters

Other parameters are passed through a pointer to a DeleteUsAppToPersonParams struct via the builder pattern


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


## FetchAlphaSender

> MessagingV1ServiceAlphaSender FetchAlphaSender(ctx, ServiceSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the resource from.
    Sid := "Sid_example" // string | The SID of the AlphaSender resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchAlphaSender(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchAlphaSender``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAlphaSender`: MessagingV1ServiceAlphaSender
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchAlphaSender`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the resource from.
**Sid** | **string** | The SID of the AlphaSender resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchAlphaSenderParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**MessagingV1ServiceAlphaSender**](MessagingV1ServiceAlphaSender.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchBrandRegistrations

> MessagingV1BrandRegistrations FetchBrandRegistrations(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The SID of the Brand Registration resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchBrandRegistrations(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchBrandRegistrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchBrandRegistrations`: MessagingV1BrandRegistrations
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchBrandRegistrations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Brand Registration resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchBrandRegistrationsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**MessagingV1BrandRegistrations**](MessagingV1BrandRegistrations.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDeactivation

> FetchDeactivation(ctx).Date(Date).Execute()





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
    Date := time.Now() // string | The request will return a list of all United States Phone Numbers that were deactivated on the day specified by this parameter. This date should be specified in YYYY-MM-DD format. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchDeactivation(context.Background()).Date(Date).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchDeactivation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchDeactivationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **Date** | **string** | The request will return a list of all United States Phone Numbers that were deactivated on the day specified by this parameter. This date should be specified in YYYY-MM-DD format.

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPhoneNumber

> MessagingV1ServicePhoneNumber FetchPhoneNumber(ctx, ServiceSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the resource from.
    Sid := "Sid_example" // string | The SID of the PhoneNumber resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchPhoneNumber(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchPhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchPhoneNumber`: MessagingV1ServicePhoneNumber
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the resource from.
**Sid** | **string** | The SID of the PhoneNumber resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchPhoneNumberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**MessagingV1ServicePhoneNumber**](MessagingV1ServicePhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> MessagingV1Service FetchService(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The SID of the Service resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchService(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchService`: MessagingV1Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Service resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchServiceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**MessagingV1Service**](MessagingV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchShortCode

> MessagingV1ServiceShortCode FetchShortCode(ctx, ServiceSid, Sid).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the resource from.
    Sid := "Sid_example" // string | The SID of the ShortCode resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchShortCode(context.Background(), ServiceSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchShortCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchShortCode`: MessagingV1ServiceShortCode
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchShortCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to fetch the resource from.
**Sid** | **string** | The SID of the ShortCode resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchShortCodeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**MessagingV1ServiceShortCode**](MessagingV1ServiceShortCode.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUsAppToPerson

> MessagingV1ServiceUsAppToPerson FetchUsAppToPerson(ctx, MessagingServiceSid).Execute()



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
    MessagingServiceSid := "MessagingServiceSid_example" // string | The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) to fetch the resource from.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchUsAppToPerson(context.Background(), MessagingServiceSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchUsAppToPerson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchUsAppToPerson`: MessagingV1ServiceUsAppToPerson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchUsAppToPerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) to fetch the resource from.

### Other Parameters

Other parameters are passed through a pointer to a FetchUsAppToPersonParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**MessagingV1ServiceUsAppToPerson**](MessagingV1ServiceUsAppToPerson.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUsAppToPersonUsecase

> MessagingV1ServiceUsAppToPersonUsecase FetchUsAppToPersonUsecase(ctx, MessagingServiceSid).Execute()



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
    MessagingServiceSid := "MessagingServiceSid_example" // string | The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) to fetch the resource from.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchUsAppToPersonUsecase(context.Background(), MessagingServiceSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchUsAppToPersonUsecase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchUsAppToPersonUsecase`: MessagingV1ServiceUsAppToPersonUsecase
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchUsAppToPersonUsecase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) to fetch the resource from.

### Other Parameters

Other parameters are passed through a pointer to a FetchUsAppToPersonUsecaseParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**MessagingV1ServiceUsAppToPersonUsecase**](MessagingV1ServiceUsAppToPersonUsecase.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchUsecase

> MessagingV1Usecase FetchUsecase(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchUsecase(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchUsecase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchUsecase`: MessagingV1Usecase
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchUsecase`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchUsecaseParams struct via the builder pattern


### Return type

[**MessagingV1Usecase**](MessagingV1Usecase.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAlphaSender

> ListAlphaSenderResponse ListAlphaSender(ctx, ServiceSid).PageSize(PageSize).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the resources from.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAlphaSender(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAlphaSender``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAlphaSender`: ListAlphaSenderResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAlphaSender`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListAlphaSenderParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListAlphaSenderResponse**](ListAlphaSenderResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBrandRegistrations

> ListBrandRegistrationsResponse ListBrandRegistrations(ctx).PageSize(PageSize).Execute()



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
    resp, r, err := api_client.DefaultApi.ListBrandRegistrations(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListBrandRegistrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBrandRegistrations`: ListBrandRegistrationsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListBrandRegistrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListBrandRegistrationsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListBrandRegistrationsResponse**](ListBrandRegistrationsResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhoneNumber

> ListPhoneNumberResponse ListPhoneNumber(ctx, ServiceSid).PageSize(PageSize).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the resources from.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListPhoneNumber(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListPhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPhoneNumber`: ListPhoneNumberResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListPhoneNumberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListPhoneNumberResponse**](ListPhoneNumberResponse.md)

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


## ListShortCode

> ListShortCodeResponse ListShortCode(ctx, ServiceSid).PageSize(PageSize).Execute()



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
    ServiceSid := "ServiceSid_example" // string | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the resources from.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListShortCode(context.Background(), ServiceSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListShortCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListShortCode`: ListShortCodeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListShortCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) to read the resources from.

### Other Parameters

Other parameters are passed through a pointer to a ListShortCodeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListShortCodeResponse**](ListShortCodeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> MessagingV1Service UpdateService(ctx, Sid).AreaCodeGeomatch(AreaCodeGeomatch).FallbackMethod(FallbackMethod).FallbackToLongCode(FallbackToLongCode).FallbackUrl(FallbackUrl).FriendlyName(FriendlyName).InboundMethod(InboundMethod).InboundRequestUrl(InboundRequestUrl).MmsConverter(MmsConverter).ScanMessageContent(ScanMessageContent).SmartEncoding(SmartEncoding).StatusCallback(StatusCallback).StickySender(StickySender).SynchronousValidation(SynchronousValidation).ValidityPeriod(ValidityPeriod).Execute()



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
    Sid := "Sid_example" // string | The SID of the Service resource to update.
    AreaCodeGeomatch := true // bool | Whether to enable [Area Code Geomatch](https://www.twilio.com/docs/sms/services#area-code-geomatch) on the Service Instance. (optional)
    FallbackMethod := "FallbackMethod_example" // string | The HTTP method we should use to call `fallback_url`. Can be: `GET` or `POST`. (optional)
    FallbackToLongCode := true // bool | Whether to enable [Fallback to Long Code](https://www.twilio.com/docs/sms/services#fallback-to-long-code) for messages sent through the Service instance. (optional)
    FallbackUrl := "FallbackUrl_example" // string | The URL that we should call using `fallback_method` if an error occurs while retrieving or executing the TwiML from the Inbound Request URL. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)
    InboundMethod := "InboundMethod_example" // string | The HTTP method we should use to call `inbound_request_url`. Can be `GET` or `POST` and the default is `POST`. (optional)
    InboundRequestUrl := "InboundRequestUrl_example" // string | The URL we should call using `inbound_method` when a message is received by any phone number or short code in the Service. When this property is `null`, receiving inbound messages is disabled. (optional)
    MmsConverter := true // bool | Whether to enable the [MMS Converter](https://www.twilio.com/docs/sms/services#mms-converter) for messages sent through the Service instance. (optional)
    ScanMessageContent := "ScanMessageContent_example" // string | Reserved. (optional)
    SmartEncoding := true // bool | Whether to enable [Smart Encoding](https://www.twilio.com/docs/sms/services#smart-encoding) for messages sent through the Service instance. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we should call to [pass status updates](https://www.twilio.com/docs/sms/api/message-resource#message-status-values) about message delivery. (optional)
    StickySender := true // bool | Whether to enable [Sticky Sender](https://www.twilio.com/docs/sms/services#sticky-sender) on the Service instance. (optional)
    SynchronousValidation := true // bool | Reserved. (optional)
    ValidityPeriod := int32(56) // int32 | How long, in seconds, messages sent from the Service are valid. Can be an integer from `1` to `14,400`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateService(context.Background(), Sid).AreaCodeGeomatch(AreaCodeGeomatch).FallbackMethod(FallbackMethod).FallbackToLongCode(FallbackToLongCode).FallbackUrl(FallbackUrl).FriendlyName(FriendlyName).InboundMethod(InboundMethod).InboundRequestUrl(InboundRequestUrl).MmsConverter(MmsConverter).ScanMessageContent(ScanMessageContent).SmartEncoding(SmartEncoding).StatusCallback(StatusCallback).StickySender(StickySender).SynchronousValidation(SynchronousValidation).ValidityPeriod(ValidityPeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateService`: MessagingV1Service
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Service resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateServiceParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **AreaCodeGeomatch** | **bool** | Whether to enable [Area Code Geomatch](https://www.twilio.com/docs/sms/services#area-code-geomatch) on the Service Instance.
 **FallbackMethod** | **string** | The HTTP method we should use to call &#x60;fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **FallbackToLongCode** | **bool** | Whether to enable [Fallback to Long Code](https://www.twilio.com/docs/sms/services#fallback-to-long-code) for messages sent through the Service instance.
 **FallbackUrl** | **string** | The URL that we should call using &#x60;fallback_method&#x60; if an error occurs while retrieving or executing the TwiML from the Inbound Request URL.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
 **InboundMethod** | **string** | The HTTP method we should use to call &#x60;inbound_request_url&#x60;. Can be &#x60;GET&#x60; or &#x60;POST&#x60; and the default is &#x60;POST&#x60;.
 **InboundRequestUrl** | **string** | The URL we should call using &#x60;inbound_method&#x60; when a message is received by any phone number or short code in the Service. When this property is &#x60;null&#x60;, receiving inbound messages is disabled.
 **MmsConverter** | **bool** | Whether to enable the [MMS Converter](https://www.twilio.com/docs/sms/services#mms-converter) for messages sent through the Service instance.
 **ScanMessageContent** | **string** | Reserved.
 **SmartEncoding** | **bool** | Whether to enable [Smart Encoding](https://www.twilio.com/docs/sms/services#smart-encoding) for messages sent through the Service instance.
 **StatusCallback** | **string** | The URL we should call to [pass status updates](https://www.twilio.com/docs/sms/api/message-resource#message-status-values) about message delivery.
 **StickySender** | **bool** | Whether to enable [Sticky Sender](https://www.twilio.com/docs/sms/services#sticky-sender) on the Service instance.
 **SynchronousValidation** | **bool** | Reserved.
 **ValidityPeriod** | **int32** | How long, in seconds, messages sent from the Service are valid. Can be an integer from &#x60;1&#x60; to &#x60;14,400&#x60;.

### Return type

[**MessagingV1Service**](MessagingV1Service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

