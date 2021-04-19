# \DefaultApi

All URIs are relative to *https://trunking.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredentialList**](DefaultApi.md#CreateCredentialList) | **Post** /v1/Trunks/{TrunkSid}/CredentialLists | 
[**CreateIpAccessControlList**](DefaultApi.md#CreateIpAccessControlList) | **Post** /v1/Trunks/{TrunkSid}/IpAccessControlLists | 
[**CreateOriginationUrl**](DefaultApi.md#CreateOriginationUrl) | **Post** /v1/Trunks/{TrunkSid}/OriginationUrls | 
[**CreatePhoneNumber**](DefaultApi.md#CreatePhoneNumber) | **Post** /v1/Trunks/{TrunkSid}/PhoneNumbers | 
[**CreateTrunk**](DefaultApi.md#CreateTrunk) | **Post** /v1/Trunks | 
[**DeleteCredentialList**](DefaultApi.md#DeleteCredentialList) | **Delete** /v1/Trunks/{TrunkSid}/CredentialLists/{Sid} | 
[**DeleteIpAccessControlList**](DefaultApi.md#DeleteIpAccessControlList) | **Delete** /v1/Trunks/{TrunkSid}/IpAccessControlLists/{Sid} | 
[**DeleteOriginationUrl**](DefaultApi.md#DeleteOriginationUrl) | **Delete** /v1/Trunks/{TrunkSid}/OriginationUrls/{Sid} | 
[**DeletePhoneNumber**](DefaultApi.md#DeletePhoneNumber) | **Delete** /v1/Trunks/{TrunkSid}/PhoneNumbers/{Sid} | 
[**DeleteTrunk**](DefaultApi.md#DeleteTrunk) | **Delete** /v1/Trunks/{Sid} | 
[**FetchCredentialList**](DefaultApi.md#FetchCredentialList) | **Get** /v1/Trunks/{TrunkSid}/CredentialLists/{Sid} | 
[**FetchIpAccessControlList**](DefaultApi.md#FetchIpAccessControlList) | **Get** /v1/Trunks/{TrunkSid}/IpAccessControlLists/{Sid} | 
[**FetchOriginationUrl**](DefaultApi.md#FetchOriginationUrl) | **Get** /v1/Trunks/{TrunkSid}/OriginationUrls/{Sid} | 
[**FetchPhoneNumber**](DefaultApi.md#FetchPhoneNumber) | **Get** /v1/Trunks/{TrunkSid}/PhoneNumbers/{Sid} | 
[**FetchRecording**](DefaultApi.md#FetchRecording) | **Get** /v1/Trunks/{TrunkSid}/Recording | 
[**FetchTrunk**](DefaultApi.md#FetchTrunk) | **Get** /v1/Trunks/{Sid} | 
[**ListCredentialList**](DefaultApi.md#ListCredentialList) | **Get** /v1/Trunks/{TrunkSid}/CredentialLists | 
[**ListIpAccessControlList**](DefaultApi.md#ListIpAccessControlList) | **Get** /v1/Trunks/{TrunkSid}/IpAccessControlLists | 
[**ListOriginationUrl**](DefaultApi.md#ListOriginationUrl) | **Get** /v1/Trunks/{TrunkSid}/OriginationUrls | 
[**ListPhoneNumber**](DefaultApi.md#ListPhoneNumber) | **Get** /v1/Trunks/{TrunkSid}/PhoneNumbers | 
[**ListTrunk**](DefaultApi.md#ListTrunk) | **Get** /v1/Trunks | 
[**UpdateOriginationUrl**](DefaultApi.md#UpdateOriginationUrl) | **Post** /v1/Trunks/{TrunkSid}/OriginationUrls/{Sid} | 
[**UpdateRecording**](DefaultApi.md#UpdateRecording) | **Post** /v1/Trunks/{TrunkSid}/Recording | 
[**UpdateTrunk**](DefaultApi.md#UpdateTrunk) | **Post** /v1/Trunks/{Sid} | 



## CreateCredentialList

> TrunkingV1TrunkCredentialList CreateCredentialList(ctx, TrunkSid).CredentialListSid(CredentialListSid).Execute()



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
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk to associate the credential list with.
    CredentialListSid := "CredentialListSid_example" // string | The SID of the [Credential List](https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource) that you want to associate with the trunk. Once associated, we will authenticate access to the trunk against this list. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateCredentialList(context.Background(), TrunkSid).CredentialListSid(CredentialListSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCredentialList`: TrunkingV1TrunkCredentialList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk to associate the credential list with.

### Other Parameters

Other parameters are passed through a pointer to a CreateCredentialListParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **CredentialListSid** | **string** | The SID of the [Credential List](https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource) that you want to associate with the trunk. Once associated, we will authenticate access to the trunk against this list.

### Return type

[**TrunkingV1TrunkCredentialList**](TrunkingV1TrunkCredentialList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIpAccessControlList

> TrunkingV1TrunkIpAccessControlList CreateIpAccessControlList(ctx, TrunkSid).IpAccessControlListSid(IpAccessControlListSid).Execute()





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
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk to associate the IP Access Control List with.
    IpAccessControlListSid := "IpAccessControlListSid_example" // string | The SID of the [IP Access Control List](https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource) that you want to associate with the trunk. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateIpAccessControlList(context.Background(), TrunkSid).IpAccessControlListSid(IpAccessControlListSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIpAccessControlList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIpAccessControlList`: TrunkingV1TrunkIpAccessControlList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIpAccessControlList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk to associate the IP Access Control List with.

### Other Parameters

Other parameters are passed through a pointer to a CreateIpAccessControlListParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **IpAccessControlListSid** | **string** | The SID of the [IP Access Control List](https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource) that you want to associate with the trunk.

### Return type

[**TrunkingV1TrunkIpAccessControlList**](TrunkingV1TrunkIpAccessControlList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOriginationUrl

> TrunkingV1TrunkOriginationUrl CreateOriginationUrl(ctx, TrunkSid).Enabled(Enabled).FriendlyName(FriendlyName).Priority(Priority).SipUrl(SipUrl).Weight(Weight).Execute()



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
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk to associate the resource with.
    Enabled := true // bool | Whether the URL is enabled. The default is `true`. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)
    Priority := int32(56) // int32 | The relative importance of the URI. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important URI. (optional)
    SipUrl := "SipUrl_example" // string | The SIP address you want Twilio to route your Origination calls to. This must be a `sip:` schema. (optional)
    Weight := int32(56) // int32 | The value that determines the relative share of the load the URI should receive compared to other URIs with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. URLs with higher values receive more load than those with lower ones with the same priority. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateOriginationUrl(context.Background(), TrunkSid).Enabled(Enabled).FriendlyName(FriendlyName).Priority(Priority).SipUrl(SipUrl).Weight(Weight).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateOriginationUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOriginationUrl`: TrunkingV1TrunkOriginationUrl
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateOriginationUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk to associate the resource with.

### Other Parameters

Other parameters are passed through a pointer to a CreateOriginationUrlParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Enabled** | **bool** | Whether the URL is enabled. The default is &#x60;true&#x60;.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
 **Priority** | **int32** | The relative importance of the URI. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important URI.
 **SipUrl** | **string** | The SIP address you want Twilio to route your Origination calls to. This must be a &#x60;sip:&#x60; schema.
 **Weight** | **int32** | The value that determines the relative share of the load the URI should receive compared to other URIs with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. URLs with higher values receive more load than those with lower ones with the same priority.

### Return type

[**TrunkingV1TrunkOriginationUrl**](TrunkingV1TrunkOriginationUrl.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePhoneNumber

> TrunkingV1TrunkPhoneNumber CreatePhoneNumber(ctx, TrunkSid).PhoneNumberSid(PhoneNumberSid).Execute()



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
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk to associate the phone number with.
    PhoneNumberSid := "PhoneNumberSid_example" // string | The SID of the [Incoming Phone Number](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) that you want to associate with the trunk. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreatePhoneNumber(context.Background(), TrunkSid).PhoneNumberSid(PhoneNumberSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreatePhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePhoneNumber`: TrunkingV1TrunkPhoneNumber
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreatePhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk to associate the phone number with.

### Other Parameters

Other parameters are passed through a pointer to a CreatePhoneNumberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PhoneNumberSid** | **string** | The SID of the [Incoming Phone Number](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) that you want to associate with the trunk.

### Return type

[**TrunkingV1TrunkPhoneNumber**](TrunkingV1TrunkPhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTrunk

> TrunkingV1Trunk CreateTrunk(ctx).CnamLookupEnabled(CnamLookupEnabled).DisasterRecoveryMethod(DisasterRecoveryMethod).DisasterRecoveryUrl(DisasterRecoveryUrl).DomainName(DomainName).FriendlyName(FriendlyName).Secure(Secure).TransferMode(TransferMode).Execute()



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
    CnamLookupEnabled := true // bool | Whether Caller ID Name (CNAM) lookup should be enabled for the trunk. If enabled, all inbound calls to the SIP Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information. (optional)
    DisasterRecoveryMethod := "DisasterRecoveryMethod_example" // string | The HTTP method we should use to call the `disaster_recovery_url`. Can be: `GET` or `POST`. (optional)
    DisasterRecoveryUrl := "DisasterRecoveryUrl_example" // string | The URL we should call using the `disaster_recovery_method` if an error occurs while sending SIP traffic towards the configured Origination URL. We retrieve TwiML from the URL and execute the instructions like any other normal TwiML call. See [Disaster Recovery](https://www.twilio.com/docs/sip-trunking#disaster-recovery) for more information. (optional)
    DomainName := "DomainName_example" // string | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and `-` and must end with `pstn.twilio.com`. See [Termination Settings](https://www.twilio.com/docs/sip-trunking#termination) for more information. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)
    Secure := true // bool | Whether Secure Trunking is enabled for the trunk. If enabled, all calls going through the trunk will be secure using SRTP for media and TLS for signaling. If disabled, then RTP will be used for media. See [Secure Trunking](https://www.twilio.com/docs/sip-trunking#securetrunking) for more information. (optional)
    TransferMode := "TransferMode_example" // string | The call transfer settings for the trunk. Can be: `enable-all`, `sip-only` and `disable-all`. See [Transfer](https://www.twilio.com/docs/sip-trunking/call-transfer) for more information. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateTrunk(context.Background()).CnamLookupEnabled(CnamLookupEnabled).DisasterRecoveryMethod(DisasterRecoveryMethod).DisasterRecoveryUrl(DisasterRecoveryUrl).DomainName(DomainName).FriendlyName(FriendlyName).Secure(Secure).TransferMode(TransferMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTrunk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTrunk`: TrunkingV1Trunk
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTrunk`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateTrunkParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **CnamLookupEnabled** | **bool** | Whether Caller ID Name (CNAM) lookup should be enabled for the trunk. If enabled, all inbound calls to the SIP Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
 **DisasterRecoveryMethod** | **string** | The HTTP method we should use to call the &#x60;disaster_recovery_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **DisasterRecoveryUrl** | **string** | The URL we should call using the &#x60;disaster_recovery_method&#x60; if an error occurs while sending SIP traffic towards the configured Origination URL. We retrieve TwiML from the URL and execute the instructions like any other normal TwiML call. See [Disaster Recovery](https://www.twilio.com/docs/sip-trunking#disaster-recovery) for more information.
 **DomainName** | **string** | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and &#x60;-&#x60; and must end with &#x60;pstn.twilio.com&#x60;. See [Termination Settings](https://www.twilio.com/docs/sip-trunking#termination) for more information.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
 **Secure** | **bool** | Whether Secure Trunking is enabled for the trunk. If enabled, all calls going through the trunk will be secure using SRTP for media and TLS for signaling. If disabled, then RTP will be used for media. See [Secure Trunking](https://www.twilio.com/docs/sip-trunking#securetrunking) for more information.
 **TransferMode** | **string** | The call transfer settings for the trunk. Can be: &#x60;enable-all&#x60;, &#x60;sip-only&#x60; and &#x60;disable-all&#x60;. See [Transfer](https://www.twilio.com/docs/sip-trunking/call-transfer) for more information.

### Return type

[**TrunkingV1Trunk**](TrunkingV1Trunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCredentialList

> DeleteCredentialList(ctx, TrunkSid, Sid).Execute()



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
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk from which to delete the credential list.
    Sid := "Sid_example" // string | The unique string that we created to identify the CredentialList resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteCredentialList(context.Background(), TrunkSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to delete the credential list.
**Sid** | **string** | The unique string that we created to identify the CredentialList resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCredentialListParams struct via the builder pattern


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


## DeleteIpAccessControlList

> DeleteIpAccessControlList(ctx, TrunkSid, Sid).Execute()





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
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk from which to delete the IP Access Control List.
    Sid := "Sid_example" // string | The unique string that we created to identify the IpAccessControlList resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteIpAccessControlList(context.Background(), TrunkSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIpAccessControlList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to delete the IP Access Control List.
**Sid** | **string** | The unique string that we created to identify the IpAccessControlList resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteIpAccessControlListParams struct via the builder pattern


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


## DeleteOriginationUrl

> DeleteOriginationUrl(ctx, TrunkSid, Sid).Execute()



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
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk from which to delete the OriginationUrl.
    Sid := "Sid_example" // string | The unique string that we created to identify the OriginationUrl resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteOriginationUrl(context.Background(), TrunkSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteOriginationUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to delete the OriginationUrl.
**Sid** | **string** | The unique string that we created to identify the OriginationUrl resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteOriginationUrlParams struct via the builder pattern


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

> DeletePhoneNumber(ctx, TrunkSid, Sid).Execute()



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
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk from which to delete the PhoneNumber resource.
    Sid := "Sid_example" // string | The unique string that we created to identify the PhoneNumber resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeletePhoneNumber(context.Background(), TrunkSid, Sid).Execute()
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
**TrunkSid** | **string** | The SID of the Trunk from which to delete the PhoneNumber resource.
**Sid** | **string** | The unique string that we created to identify the PhoneNumber resource to delete.

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


## DeleteTrunk

> DeleteTrunk(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The unique string that we created to identify the Trunk resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteTrunk(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTrunk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Trunk resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteTrunkParams struct via the builder pattern


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


## FetchCredentialList

> TrunkingV1TrunkCredentialList FetchCredentialList(ctx, TrunkSid, Sid).Execute()



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
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk from which to fetch the credential list.
    Sid := "Sid_example" // string | The unique string that we created to identify the CredentialList resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchCredentialList(context.Background(), TrunkSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCredentialList`: TrunkingV1TrunkCredentialList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to fetch the credential list.
**Sid** | **string** | The unique string that we created to identify the CredentialList resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchCredentialListParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**TrunkingV1TrunkCredentialList**](TrunkingV1TrunkCredentialList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIpAccessControlList

> TrunkingV1TrunkIpAccessControlList FetchIpAccessControlList(ctx, TrunkSid, Sid).Execute()



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
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk from which to fetch the IP Access Control List.
    Sid := "Sid_example" // string | The unique string that we created to identify the IpAccessControlList resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchIpAccessControlList(context.Background(), TrunkSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchIpAccessControlList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIpAccessControlList`: TrunkingV1TrunkIpAccessControlList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchIpAccessControlList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to fetch the IP Access Control List.
**Sid** | **string** | The unique string that we created to identify the IpAccessControlList resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchIpAccessControlListParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**TrunkingV1TrunkIpAccessControlList**](TrunkingV1TrunkIpAccessControlList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchOriginationUrl

> TrunkingV1TrunkOriginationUrl FetchOriginationUrl(ctx, TrunkSid, Sid).Execute()



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
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk from which to fetch the OriginationUrl.
    Sid := "Sid_example" // string | The unique string that we created to identify the OriginationUrl resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchOriginationUrl(context.Background(), TrunkSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchOriginationUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchOriginationUrl`: TrunkingV1TrunkOriginationUrl
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchOriginationUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to fetch the OriginationUrl.
**Sid** | **string** | The unique string that we created to identify the OriginationUrl resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchOriginationUrlParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**TrunkingV1TrunkOriginationUrl**](TrunkingV1TrunkOriginationUrl.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPhoneNumber

> TrunkingV1TrunkPhoneNumber FetchPhoneNumber(ctx, TrunkSid, Sid).Execute()



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
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk from which to fetch the PhoneNumber resource.
    Sid := "Sid_example" // string | The unique string that we created to identify the PhoneNumber resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchPhoneNumber(context.Background(), TrunkSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchPhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchPhoneNumber`: TrunkingV1TrunkPhoneNumber
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to fetch the PhoneNumber resource.
**Sid** | **string** | The unique string that we created to identify the PhoneNumber resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchPhoneNumberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**TrunkingV1TrunkPhoneNumber**](TrunkingV1TrunkPhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecording

> TrunkingV1TrunkRecording FetchRecording(ctx, TrunkSid).Execute()



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
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk from which to fetch the recording settings.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchRecording(context.Background(), TrunkSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRecording`: TrunkingV1TrunkRecording
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to fetch the recording settings.

### Other Parameters

Other parameters are passed through a pointer to a FetchRecordingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**TrunkingV1TrunkRecording**](TrunkingV1TrunkRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTrunk

> TrunkingV1Trunk FetchTrunk(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The unique string that we created to identify the Trunk resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchTrunk(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchTrunk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchTrunk`: TrunkingV1Trunk
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchTrunk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Trunk resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchTrunkParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**TrunkingV1Trunk**](TrunkingV1Trunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredentialList

> ListCredentialListResponse ListCredentialList(ctx, TrunkSid).PageSize(PageSize).Execute()



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
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk from which to read the credential lists.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListCredentialList(context.Background(), TrunkSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCredentialList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCredentialList`: ListCredentialListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCredentialList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to read the credential lists.

### Other Parameters

Other parameters are passed through a pointer to a ListCredentialListParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListCredentialListResponse**](ListCredentialListResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIpAccessControlList

> ListIpAccessControlListResponse ListIpAccessControlList(ctx, TrunkSid).PageSize(PageSize).Execute()





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
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk from which to read the IP Access Control Lists.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListIpAccessControlList(context.Background(), TrunkSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListIpAccessControlList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIpAccessControlList`: ListIpAccessControlListResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListIpAccessControlList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to read the IP Access Control Lists.

### Other Parameters

Other parameters are passed through a pointer to a ListIpAccessControlListParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListIpAccessControlListResponse**](ListIpAccessControlListResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOriginationUrl

> ListOriginationUrlResponse ListOriginationUrl(ctx, TrunkSid).PageSize(PageSize).Execute()



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
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk from which to read the OriginationUrl.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListOriginationUrl(context.Background(), TrunkSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListOriginationUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOriginationUrl`: ListOriginationUrlResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListOriginationUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to read the OriginationUrl.

### Other Parameters

Other parameters are passed through a pointer to a ListOriginationUrlParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListOriginationUrlResponse**](ListOriginationUrlResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhoneNumber

> ListPhoneNumberResponse ListPhoneNumber(ctx, TrunkSid).PageSize(PageSize).Execute()



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
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk from which to read the PhoneNumber resources.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListPhoneNumber(context.Background(), TrunkSid).PageSize(PageSize).Execute()
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
**TrunkSid** | **string** | The SID of the Trunk from which to read the PhoneNumber resources.

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


## ListTrunk

> ListTrunkResponse ListTrunk(ctx).PageSize(PageSize).Execute()



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
    resp, r, err := api_client.DefaultApi.ListTrunk(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListTrunk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTrunk`: ListTrunkResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListTrunk`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListTrunkParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListTrunkResponse**](ListTrunkResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOriginationUrl

> TrunkingV1TrunkOriginationUrl UpdateOriginationUrl(ctx, TrunkSid, Sid).Enabled(Enabled).FriendlyName(FriendlyName).Priority(Priority).SipUrl(SipUrl).Weight(Weight).Execute()



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
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk from which to update the OriginationUrl.
    Sid := "Sid_example" // string | The unique string that we created to identify the OriginationUrl resource to update.
    Enabled := true // bool | Whether the URL is enabled. The default is `true`. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)
    Priority := int32(56) // int32 | The relative importance of the URI. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important URI. (optional)
    SipUrl := "SipUrl_example" // string | The SIP address you want Twilio to route your Origination calls to. This must be a `sip:` schema. `sips` is NOT supported. (optional)
    Weight := int32(56) // int32 | The value that determines the relative share of the load the URI should receive compared to other URIs with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. URLs with higher values receive more load than those with lower ones with the same priority. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateOriginationUrl(context.Background(), TrunkSid, Sid).Enabled(Enabled).FriendlyName(FriendlyName).Priority(Priority).SipUrl(SipUrl).Weight(Weight).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateOriginationUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOriginationUrl`: TrunkingV1TrunkOriginationUrl
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateOriginationUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to update the OriginationUrl.
**Sid** | **string** | The unique string that we created to identify the OriginationUrl resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateOriginationUrlParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Enabled** | **bool** | Whether the URL is enabled. The default is &#x60;true&#x60;.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
 **Priority** | **int32** | The relative importance of the URI. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important URI.
 **SipUrl** | **string** | The SIP address you want Twilio to route your Origination calls to. This must be a &#x60;sip:&#x60; schema. &#x60;sips&#x60; is NOT supported.
 **Weight** | **int32** | The value that determines the relative share of the load the URI should receive compared to other URIs with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. URLs with higher values receive more load than those with lower ones with the same priority.

### Return type

[**TrunkingV1TrunkOriginationUrl**](TrunkingV1TrunkOriginationUrl.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRecording

> TrunkingV1TrunkRecording UpdateRecording(ctx, TrunkSid).Mode(Mode).Trim(Trim).Execute()



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
    TrunkSid := "TrunkSid_example" // string | The SID of the Trunk that will have its recording settings updated.
    Mode := "Mode_example" // string | The recording mode for the trunk. Can be do-not-record (default), record-from-ringing, record-from-answer, record-from-ringing-dual, or record-from-answer-dual. (optional)
    Trim := "Trim_example" // string | The recording trim setting for the trunk. Can be do-not-trim (default) or trim-silence. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateRecording(context.Background(), TrunkSid).Mode(Mode).Trim(Trim).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateRecording``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRecording`: TrunkingV1TrunkRecording
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk that will have its recording settings updated.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRecordingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Mode** | **string** | The recording mode for the trunk. Can be do-not-record (default), record-from-ringing, record-from-answer, record-from-ringing-dual, or record-from-answer-dual.
 **Trim** | **string** | The recording trim setting for the trunk. Can be do-not-trim (default) or trim-silence.

### Return type

[**TrunkingV1TrunkRecording**](TrunkingV1TrunkRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrunk

> TrunkingV1Trunk UpdateTrunk(ctx, Sid).CnamLookupEnabled(CnamLookupEnabled).DisasterRecoveryMethod(DisasterRecoveryMethod).DisasterRecoveryUrl(DisasterRecoveryUrl).DomainName(DomainName).FriendlyName(FriendlyName).Secure(Secure).TransferMode(TransferMode).Execute()



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
    Sid := "Sid_example" // string | The unique string that we created to identify the OriginationUrl resource to update.
    CnamLookupEnabled := true // bool | Whether Caller ID Name (CNAM) lookup should be enabled for the trunk. If enabled, all inbound calls to the SIP Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information. (optional)
    DisasterRecoveryMethod := "DisasterRecoveryMethod_example" // string | The HTTP method we should use to call the `disaster_recovery_url`. Can be: `GET` or `POST`. (optional)
    DisasterRecoveryUrl := "DisasterRecoveryUrl_example" // string | The URL we should call using the `disaster_recovery_method` if an error occurs while sending SIP traffic towards the configured Origination URL. We retrieve TwiML from the URL and execute the instructions like any other normal TwiML call. See [Disaster Recovery](https://www.twilio.com/docs/sip-trunking#disaster-recovery) for more information. (optional)
    DomainName := "DomainName_example" // string | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and `-` and must end with `pstn.twilio.com`. See [Termination Settings](https://www.twilio.com/docs/sip-trunking#termination) for more information. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)
    Secure := true // bool | Whether Secure Trunking is enabled for the trunk. If enabled, all calls going through the trunk will be secure using SRTP for media and TLS for signaling. If disabled, then RTP will be used for media. See [Secure Trunking](https://www.twilio.com/docs/sip-trunking#securetrunking) for more information. (optional)
    TransferMode := "TransferMode_example" // string | The call transfer settings for the trunk. Can be: `enable-all`, `sip-only` and `disable-all`. See [Transfer](https://www.twilio.com/docs/sip-trunking/call-transfer) for more information. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateTrunk(context.Background(), Sid).CnamLookupEnabled(CnamLookupEnabled).DisasterRecoveryMethod(DisasterRecoveryMethod).DisasterRecoveryUrl(DisasterRecoveryUrl).DomainName(DomainName).FriendlyName(FriendlyName).Secure(Secure).TransferMode(TransferMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTrunk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTrunk`: TrunkingV1Trunk
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTrunk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the OriginationUrl resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateTrunkParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **CnamLookupEnabled** | **bool** | Whether Caller ID Name (CNAM) lookup should be enabled for the trunk. If enabled, all inbound calls to the SIP Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
 **DisasterRecoveryMethod** | **string** | The HTTP method we should use to call the &#x60;disaster_recovery_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **DisasterRecoveryUrl** | **string** | The URL we should call using the &#x60;disaster_recovery_method&#x60; if an error occurs while sending SIP traffic towards the configured Origination URL. We retrieve TwiML from the URL and execute the instructions like any other normal TwiML call. See [Disaster Recovery](https://www.twilio.com/docs/sip-trunking#disaster-recovery) for more information.
 **DomainName** | **string** | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and &#x60;-&#x60; and must end with &#x60;pstn.twilio.com&#x60;. See [Termination Settings](https://www.twilio.com/docs/sip-trunking#termination) for more information.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
 **Secure** | **bool** | Whether Secure Trunking is enabled for the trunk. If enabled, all calls going through the trunk will be secure using SRTP for media and TLS for signaling. If disabled, then RTP will be used for media. See [Secure Trunking](https://www.twilio.com/docs/sip-trunking#securetrunking) for more information.
 **TransferMode** | **string** | The call transfer settings for the trunk. Can be: &#x60;enable-all&#x60;, &#x60;sip-only&#x60; and &#x60;disable-all&#x60;. See [Transfer](https://www.twilio.com/docs/sip-trunking/call-transfer) for more information.

### Return type

[**TrunkingV1Trunk**](TrunkingV1Trunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

