# \DefaultApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateByocTrunk**](DefaultApi.md#CreateByocTrunk) | **Post** /v1/ByocTrunks | 
[**CreateConnectionPolicy**](DefaultApi.md#CreateConnectionPolicy) | **Post** /v1/ConnectionPolicies | 
[**CreateConnectionPolicyTarget**](DefaultApi.md#CreateConnectionPolicyTarget) | **Post** /v1/ConnectionPolicies/{ConnectionPolicySid}/Targets | 
[**CreateDialingPermissionsCountryBulkUpdate**](DefaultApi.md#CreateDialingPermissionsCountryBulkUpdate) | **Post** /v1/DialingPermissions/BulkCountryUpdates | 
[**CreateIpRecord**](DefaultApi.md#CreateIpRecord) | **Post** /v1/IpRecords | 
[**CreateSourceIpMapping**](DefaultApi.md#CreateSourceIpMapping) | **Post** /v1/SourceIpMappings | 
[**DeleteByocTrunk**](DefaultApi.md#DeleteByocTrunk) | **Delete** /v1/ByocTrunks/{Sid} | 
[**DeleteConnectionPolicy**](DefaultApi.md#DeleteConnectionPolicy) | **Delete** /v1/ConnectionPolicies/{Sid} | 
[**DeleteConnectionPolicyTarget**](DefaultApi.md#DeleteConnectionPolicyTarget) | **Delete** /v1/ConnectionPolicies/{ConnectionPolicySid}/Targets/{Sid} | 
[**DeleteIpRecord**](DefaultApi.md#DeleteIpRecord) | **Delete** /v1/IpRecords/{Sid} | 
[**DeleteSourceIpMapping**](DefaultApi.md#DeleteSourceIpMapping) | **Delete** /v1/SourceIpMappings/{Sid} | 
[**FetchByocTrunk**](DefaultApi.md#FetchByocTrunk) | **Get** /v1/ByocTrunks/{Sid} | 
[**FetchConnectionPolicy**](DefaultApi.md#FetchConnectionPolicy) | **Get** /v1/ConnectionPolicies/{Sid} | 
[**FetchConnectionPolicyTarget**](DefaultApi.md#FetchConnectionPolicyTarget) | **Get** /v1/ConnectionPolicies/{ConnectionPolicySid}/Targets/{Sid} | 
[**FetchDialingPermissionsCountry**](DefaultApi.md#FetchDialingPermissionsCountry) | **Get** /v1/DialingPermissions/Countries/{IsoCode} | 
[**FetchDialingPermissionsSettings**](DefaultApi.md#FetchDialingPermissionsSettings) | **Get** /v1/Settings | 
[**FetchIpRecord**](DefaultApi.md#FetchIpRecord) | **Get** /v1/IpRecords/{Sid} | 
[**FetchSourceIpMapping**](DefaultApi.md#FetchSourceIpMapping) | **Get** /v1/SourceIpMappings/{Sid} | 
[**ListByocTrunk**](DefaultApi.md#ListByocTrunk) | **Get** /v1/ByocTrunks | 
[**ListConnectionPolicy**](DefaultApi.md#ListConnectionPolicy) | **Get** /v1/ConnectionPolicies | 
[**ListConnectionPolicyTarget**](DefaultApi.md#ListConnectionPolicyTarget) | **Get** /v1/ConnectionPolicies/{ConnectionPolicySid}/Targets | 
[**ListDialingPermissionsCountry**](DefaultApi.md#ListDialingPermissionsCountry) | **Get** /v1/DialingPermissions/Countries | 
[**ListDialingPermissionsHrsPrefixes**](DefaultApi.md#ListDialingPermissionsHrsPrefixes) | **Get** /v1/DialingPermissions/Countries/{IsoCode}/HighRiskSpecialPrefixes | 
[**ListIpRecord**](DefaultApi.md#ListIpRecord) | **Get** /v1/IpRecords | 
[**ListSourceIpMapping**](DefaultApi.md#ListSourceIpMapping) | **Get** /v1/SourceIpMappings | 
[**UpdateByocTrunk**](DefaultApi.md#UpdateByocTrunk) | **Post** /v1/ByocTrunks/{Sid} | 
[**UpdateConnectionPolicy**](DefaultApi.md#UpdateConnectionPolicy) | **Post** /v1/ConnectionPolicies/{Sid} | 
[**UpdateConnectionPolicyTarget**](DefaultApi.md#UpdateConnectionPolicyTarget) | **Post** /v1/ConnectionPolicies/{ConnectionPolicySid}/Targets/{Sid} | 
[**UpdateDialingPermissionsSettings**](DefaultApi.md#UpdateDialingPermissionsSettings) | **Post** /v1/Settings | 
[**UpdateIpRecord**](DefaultApi.md#UpdateIpRecord) | **Post** /v1/IpRecords/{Sid} | 
[**UpdateSourceIpMapping**](DefaultApi.md#UpdateSourceIpMapping) | **Post** /v1/SourceIpMappings/{Sid} | 



## CreateByocTrunk

> VoiceV1ByocTrunk CreateByocTrunk(ctx).CnamLookupEnabled(CnamLookupEnabled).ConnectionPolicySid(ConnectionPolicySid).FriendlyName(FriendlyName).FromDomainSid(FromDomainSid).StatusCallbackMethod(StatusCallbackMethod).StatusCallbackUrl(StatusCallbackUrl).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceUrl(VoiceUrl).Execute()



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
    CnamLookupEnabled := true // bool | Whether Caller ID Name (CNAM) lookup is enabled for the trunk. If enabled, all inbound calls to the BYOC Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information. (optional)
    ConnectionPolicySid := "ConnectionPolicySid_example" // string | The SID of the Connection Policy that Twilio will use when routing traffic to your communications infrastructure. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. (optional)
    FromDomainSid := "FromDomainSid_example" // string | The SID of the SIP Domain that should be used in the `From` header of originating calls sent to your SIP infrastructure. If your SIP infrastructure allows users to \\\"call back\\\" an incoming call, configure this with a [SIP Domain](https://www.twilio.com/docs/voice/api/sending-sip) to ensure proper routing. If not configured, the from domain will default to \\\"sip.twilio.com\\\". (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback_url`. Can be: `GET` or `POST`. (optional)
    StatusCallbackUrl := "StatusCallbackUrl_example" // string | The URL that we should call to pass status parameters (such as call ended) to your application. (optional)
    VoiceFallbackMethod := "VoiceFallbackMethod_example" // string | The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`. (optional)
    VoiceFallbackUrl := "VoiceFallbackUrl_example" // string | The URL that we should call when an error occurs while retrieving or executing the TwiML from `voice_url`. (optional)
    VoiceMethod := "VoiceMethod_example" // string | The HTTP method we should use to call `voice_url`. Can be: `GET` or `POST`. (optional)
    VoiceUrl := "VoiceUrl_example" // string | The URL we should call when the BYOC Trunk receives a call. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateByocTrunk(context.Background()).CnamLookupEnabled(CnamLookupEnabled).ConnectionPolicySid(ConnectionPolicySid).FriendlyName(FriendlyName).FromDomainSid(FromDomainSid).StatusCallbackMethod(StatusCallbackMethod).StatusCallbackUrl(StatusCallbackUrl).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceUrl(VoiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateByocTrunk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateByocTrunk`: VoiceV1ByocTrunk
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateByocTrunk`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateByocTrunkParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **CnamLookupEnabled** | **bool** | Whether Caller ID Name (CNAM) lookup is enabled for the trunk. If enabled, all inbound calls to the BYOC Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
 **ConnectionPolicySid** | **string** | The SID of the Connection Policy that Twilio will use when routing traffic to your communications infrastructure.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
 **FromDomainSid** | **string** | The SID of the SIP Domain that should be used in the &#x60;From&#x60; header of originating calls sent to your SIP infrastructure. If your SIP infrastructure allows users to \\\&quot;call back\\\&quot; an incoming call, configure this with a [SIP Domain](https://www.twilio.com/docs/voice/api/sending-sip) to ensure proper routing. If not configured, the from domain will default to \\\&quot;sip.twilio.com\\\&quot;.
 **StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **StatusCallbackUrl** | **string** | The URL that we should call to pass status parameters (such as call ended) to your application.
 **VoiceFallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs while retrieving or executing the TwiML from &#x60;voice_url&#x60;.
 **VoiceMethod** | **string** | The HTTP method we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **VoiceUrl** | **string** | The URL we should call when the BYOC Trunk receives a call.

### Return type

[**VoiceV1ByocTrunk**](VoiceV1ByocTrunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConnectionPolicy

> VoiceV1ConnectionPolicy CreateConnectionPolicy(ctx).FriendlyName(FriendlyName).Execute()



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
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateConnectionPolicy(context.Background()).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateConnectionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnectionPolicy`: VoiceV1ConnectionPolicy
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateConnectionPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateConnectionPolicyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.

### Return type

[**VoiceV1ConnectionPolicy**](VoiceV1ConnectionPolicy.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConnectionPolicyTarget

> VoiceV1ConnectionPolicyConnectionPolicyTarget CreateConnectionPolicyTarget(ctx, ConnectionPolicySid).Enabled(Enabled).FriendlyName(FriendlyName).Priority(Priority).Target(Target).Weight(Weight).Execute()



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
    ConnectionPolicySid := "ConnectionPolicySid_example" // string | The SID of the Connection Policy that owns the Target.
    Enabled := true // bool | Whether the Target is enabled. The default is `true`. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. (optional)
    Priority := int32(56) // int32 | The relative importance of the target. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important target. (optional)
    Target := "Target_example" // string | The SIP address you want Twilio to route your calls to. This must be a `sip:` schema. `sips` is NOT supported. (optional)
    Weight := int32(56) // int32 | The value that determines the relative share of the load the Target should receive compared to other Targets with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. Targets with higher values receive more load than those with lower ones with the same priority. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateConnectionPolicyTarget(context.Background(), ConnectionPolicySid).Enabled(Enabled).FriendlyName(FriendlyName).Priority(Priority).Target(Target).Weight(Weight).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateConnectionPolicyTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnectionPolicyTarget`: VoiceV1ConnectionPolicyConnectionPolicyTarget
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateConnectionPolicyTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConnectionPolicySid** | **string** | The SID of the Connection Policy that owns the Target.

### Other Parameters

Other parameters are passed through a pointer to a CreateConnectionPolicyTargetParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Enabled** | **bool** | Whether the Target is enabled. The default is &#x60;true&#x60;.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
 **Priority** | **int32** | The relative importance of the target. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important target.
 **Target** | **string** | The SIP address you want Twilio to route your calls to. This must be a &#x60;sip:&#x60; schema. &#x60;sips&#x60; is NOT supported.
 **Weight** | **int32** | The value that determines the relative share of the load the Target should receive compared to other Targets with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. Targets with higher values receive more load than those with lower ones with the same priority.

### Return type

[**VoiceV1ConnectionPolicyConnectionPolicyTarget**](VoiceV1ConnectionPolicyConnectionPolicyTarget.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDialingPermissionsCountryBulkUpdate

> VoiceV1DialingPermissionsDialingPermissionsCountryBulkUpdate CreateDialingPermissionsCountryBulkUpdate(ctx).UpdateRequest(UpdateRequest).Execute()





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
    UpdateRequest := "UpdateRequest_example" // string | URL encoded JSON array of update objects. example : `[ { \\\"iso_code\\\": \\\"GB\\\", \\\"low_risk_numbers_enabled\\\": \\\"true\\\", \\\"high_risk_special_numbers_enabled\\\":\\\"true\\\", \\\"high_risk_tollfraud_numbers_enabled\\\": \\\"false\\\" } ]` (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateDialingPermissionsCountryBulkUpdate(context.Background()).UpdateRequest(UpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateDialingPermissionsCountryBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDialingPermissionsCountryBulkUpdate`: VoiceV1DialingPermissionsDialingPermissionsCountryBulkUpdate
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateDialingPermissionsCountryBulkUpdate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateDialingPermissionsCountryBulkUpdateParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **UpdateRequest** | **string** | URL encoded JSON array of update objects. example : &#x60;[ { \\\&quot;iso_code\\\&quot;: \\\&quot;GB\\\&quot;, \\\&quot;low_risk_numbers_enabled\\\&quot;: \\\&quot;true\\\&quot;, \\\&quot;high_risk_special_numbers_enabled\\\&quot;:\\\&quot;true\\\&quot;, \\\&quot;high_risk_tollfraud_numbers_enabled\\\&quot;: \\\&quot;false\\\&quot; } ]&#x60;

### Return type

[**VoiceV1DialingPermissionsDialingPermissionsCountryBulkUpdate**](VoiceV1DialingPermissionsDialingPermissionsCountryBulkUpdate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIpRecord

> VoiceV1IpRecord CreateIpRecord(ctx).CidrPrefixLength(CidrPrefixLength).FriendlyName(FriendlyName).IpAddress(IpAddress).Execute()



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
    CidrPrefixLength := int32(56) // int32 | An integer representing the length of the [CIDR](https://tools.ietf.org/html/rfc4632) prefix to use with this IP address. By default the entire IP address is used, which for IPv4 is value 32. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. (optional)
    IpAddress := "IpAddress_example" // string | An IP address in dotted decimal notation, IPv4 only. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateIpRecord(context.Background()).CidrPrefixLength(CidrPrefixLength).FriendlyName(FriendlyName).IpAddress(IpAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIpRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIpRecord`: VoiceV1IpRecord
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIpRecord`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateIpRecordParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **CidrPrefixLength** | **int32** | An integer representing the length of the [CIDR](https://tools.ietf.org/html/rfc4632) prefix to use with this IP address. By default the entire IP address is used, which for IPv4 is value 32.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
 **IpAddress** | **string** | An IP address in dotted decimal notation, IPv4 only.

### Return type

[**VoiceV1IpRecord**](VoiceV1IpRecord.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSourceIpMapping

> VoiceV1SourceIpMapping CreateSourceIpMapping(ctx).IpRecordSid(IpRecordSid).SipDomainSid(SipDomainSid).Execute()



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
    IpRecordSid := "IpRecordSid_example" // string | The Twilio-provided string that uniquely identifies the IP Record resource to map from. (optional)
    SipDomainSid := "SipDomainSid_example" // string | The SID of the SIP Domain that the IP Record should be mapped to. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSourceIpMapping(context.Background()).IpRecordSid(IpRecordSid).SipDomainSid(SipDomainSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSourceIpMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSourceIpMapping`: VoiceV1SourceIpMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSourceIpMapping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSourceIpMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **IpRecordSid** | **string** | The Twilio-provided string that uniquely identifies the IP Record resource to map from.
 **SipDomainSid** | **string** | The SID of the SIP Domain that the IP Record should be mapped to.

### Return type

[**VoiceV1SourceIpMapping**](VoiceV1SourceIpMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteByocTrunk

> DeleteByocTrunk(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the BYOC Trunk resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteByocTrunk(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteByocTrunk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the BYOC Trunk resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteByocTrunkParams struct via the builder pattern


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


## DeleteConnectionPolicy

> DeleteConnectionPolicy(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The unique string that we created to identify the Connection Policy resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteConnectionPolicy(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteConnectionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Connection Policy resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConnectionPolicyParams struct via the builder pattern


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


## DeleteConnectionPolicyTarget

> DeleteConnectionPolicyTarget(ctx, ConnectionPolicySid, Sid).Execute()



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
    ConnectionPolicySid := "ConnectionPolicySid_example" // string | The SID of the Connection Policy that owns the Target.
    Sid := "Sid_example" // string | The unique string that we created to identify the Target resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteConnectionPolicyTarget(context.Background(), ConnectionPolicySid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteConnectionPolicyTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConnectionPolicySid** | **string** | The SID of the Connection Policy that owns the Target.
**Sid** | **string** | The unique string that we created to identify the Target resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConnectionPolicyTargetParams struct via the builder pattern


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


## DeleteIpRecord

> DeleteIpRecord(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the IP Record resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteIpRecord(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIpRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IP Record resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteIpRecordParams struct via the builder pattern


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


## DeleteSourceIpMapping

> DeleteSourceIpMapping(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the IP Record resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSourceIpMapping(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSourceIpMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IP Record resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSourceIpMappingParams struct via the builder pattern


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


## FetchByocTrunk

> VoiceV1ByocTrunk FetchByocTrunk(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the BYOC Trunk resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchByocTrunk(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchByocTrunk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchByocTrunk`: VoiceV1ByocTrunk
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchByocTrunk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the BYOC Trunk resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchByocTrunkParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**VoiceV1ByocTrunk**](VoiceV1ByocTrunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConnectionPolicy

> VoiceV1ConnectionPolicy FetchConnectionPolicy(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The unique string that we created to identify the Connection Policy resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchConnectionPolicy(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchConnectionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchConnectionPolicy`: VoiceV1ConnectionPolicy
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchConnectionPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Connection Policy resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchConnectionPolicyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**VoiceV1ConnectionPolicy**](VoiceV1ConnectionPolicy.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConnectionPolicyTarget

> VoiceV1ConnectionPolicyConnectionPolicyTarget FetchConnectionPolicyTarget(ctx, ConnectionPolicySid, Sid).Execute()



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
    ConnectionPolicySid := "ConnectionPolicySid_example" // string | The SID of the Connection Policy that owns the Target.
    Sid := "Sid_example" // string | The unique string that we created to identify the Target resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchConnectionPolicyTarget(context.Background(), ConnectionPolicySid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchConnectionPolicyTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchConnectionPolicyTarget`: VoiceV1ConnectionPolicyConnectionPolicyTarget
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchConnectionPolicyTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConnectionPolicySid** | **string** | The SID of the Connection Policy that owns the Target.
**Sid** | **string** | The unique string that we created to identify the Target resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchConnectionPolicyTargetParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**VoiceV1ConnectionPolicyConnectionPolicyTarget**](VoiceV1ConnectionPolicyConnectionPolicyTarget.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDialingPermissionsCountry

> VoiceV1DialingPermissionsDialingPermissionsCountryInstance FetchDialingPermissionsCountry(ctx, IsoCode).Execute()





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
    IsoCode := "IsoCode_example" // string | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the DialingPermissions Country resource to fetch

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchDialingPermissionsCountry(context.Background(), IsoCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchDialingPermissionsCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchDialingPermissionsCountry`: VoiceV1DialingPermissionsDialingPermissionsCountryInstance
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchDialingPermissionsCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IsoCode** | **string** | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the DialingPermissions Country resource to fetch

### Other Parameters

Other parameters are passed through a pointer to a FetchDialingPermissionsCountryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**VoiceV1DialingPermissionsDialingPermissionsCountryInstance**](VoiceV1DialingPermissionsDialingPermissionsCountryInstance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDialingPermissionsSettings

> VoiceV1DialingPermissionsDialingPermissionsSettings FetchDialingPermissionsSettings(ctx).Execute()





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
    resp, r, err := api_client.DefaultApi.FetchDialingPermissionsSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchDialingPermissionsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchDialingPermissionsSettings`: VoiceV1DialingPermissionsDialingPermissionsSettings
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchDialingPermissionsSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchDialingPermissionsSettingsParams struct via the builder pattern


### Return type

[**VoiceV1DialingPermissionsDialingPermissionsSettings**](VoiceV1DialingPermissionsDialingPermissionsSettings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIpRecord

> VoiceV1IpRecord FetchIpRecord(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the IP Record resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchIpRecord(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchIpRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIpRecord`: VoiceV1IpRecord
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchIpRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IP Record resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchIpRecordParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**VoiceV1IpRecord**](VoiceV1IpRecord.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSourceIpMapping

> VoiceV1SourceIpMapping FetchSourceIpMapping(ctx, Sid).Execute()



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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the IP Record resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSourceIpMapping(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSourceIpMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSourceIpMapping`: VoiceV1SourceIpMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSourceIpMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IP Record resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSourceIpMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**VoiceV1SourceIpMapping**](VoiceV1SourceIpMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListByocTrunk

> ListByocTrunkResponse ListByocTrunk(ctx).PageSize(PageSize).Execute()



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
    resp, r, err := api_client.DefaultApi.ListByocTrunk(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListByocTrunk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListByocTrunk`: ListByocTrunkResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListByocTrunk`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListByocTrunkParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListByocTrunkResponse**](ListByocTrunkResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectionPolicy

> ListConnectionPolicyResponse ListConnectionPolicy(ctx).PageSize(PageSize).Execute()



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
    resp, r, err := api_client.DefaultApi.ListConnectionPolicy(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConnectionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectionPolicy`: ListConnectionPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConnectionPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListConnectionPolicyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListConnectionPolicyResponse**](ListConnectionPolicyResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectionPolicyTarget

> ListConnectionPolicyTargetResponse ListConnectionPolicyTarget(ctx, ConnectionPolicySid).PageSize(PageSize).Execute()



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
    ConnectionPolicySid := "ConnectionPolicySid_example" // string | The SID of the Connection Policy from which to read the Targets.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListConnectionPolicyTarget(context.Background(), ConnectionPolicySid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListConnectionPolicyTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectionPolicyTarget`: ListConnectionPolicyTargetResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListConnectionPolicyTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConnectionPolicySid** | **string** | The SID of the Connection Policy from which to read the Targets.

### Other Parameters

Other parameters are passed through a pointer to a ListConnectionPolicyTargetParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListConnectionPolicyTargetResponse**](ListConnectionPolicyTargetResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDialingPermissionsCountry

> ListDialingPermissionsCountryResponse ListDialingPermissionsCountry(ctx).IsoCode(IsoCode).Continent(Continent).CountryCode(CountryCode).LowRiskNumbersEnabled(LowRiskNumbersEnabled).HighRiskSpecialNumbersEnabled(HighRiskSpecialNumbersEnabled).HighRiskTollfraudNumbersEnabled(HighRiskTollfraudNumbersEnabled).PageSize(PageSize).Execute()





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
    IsoCode := "IsoCode_example" // string | Filter to retrieve the country permissions by specifying the [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) (optional)
    Continent := "Continent_example" // string | Filter to retrieve the country permissions by specifying the continent (optional)
    CountryCode := "CountryCode_example" // string | Filter the results by specified [country codes](https://www.itu.int/itudoc/itu-t/ob-lists/icc/e164_763.html) (optional)
    LowRiskNumbersEnabled := true // bool | Filter to retrieve the country permissions with dialing to low-risk numbers enabled. Can be: `true` or `false`. (optional)
    HighRiskSpecialNumbersEnabled := true // bool | Filter to retrieve the country permissions with dialing to high-risk special service numbers enabled. Can be: `true` or `false` (optional)
    HighRiskTollfraudNumbersEnabled := true // bool | Filter to retrieve the country permissions with dialing to high-risk [toll fraud](https://www.twilio.com/learn/voice-and-video/toll-fraud) numbers enabled. Can be: `true` or `false`. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListDialingPermissionsCountry(context.Background()).IsoCode(IsoCode).Continent(Continent).CountryCode(CountryCode).LowRiskNumbersEnabled(LowRiskNumbersEnabled).HighRiskSpecialNumbersEnabled(HighRiskSpecialNumbersEnabled).HighRiskTollfraudNumbersEnabled(HighRiskTollfraudNumbersEnabled).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListDialingPermissionsCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDialingPermissionsCountry`: ListDialingPermissionsCountryResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListDialingPermissionsCountry`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListDialingPermissionsCountryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **IsoCode** | **string** | Filter to retrieve the country permissions by specifying the [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
 **Continent** | **string** | Filter to retrieve the country permissions by specifying the continent
 **CountryCode** | **string** | Filter the results by specified [country codes](https://www.itu.int/itudoc/itu-t/ob-lists/icc/e164_763.html)
 **LowRiskNumbersEnabled** | **bool** | Filter to retrieve the country permissions with dialing to low-risk numbers enabled. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **HighRiskSpecialNumbersEnabled** | **bool** | Filter to retrieve the country permissions with dialing to high-risk special service numbers enabled. Can be: &#x60;true&#x60; or &#x60;false&#x60;
 **HighRiskTollfraudNumbersEnabled** | **bool** | Filter to retrieve the country permissions with dialing to high-risk [toll fraud](https://www.twilio.com/learn/voice-and-video/toll-fraud) numbers enabled. Can be: &#x60;true&#x60; or &#x60;false&#x60;.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListDialingPermissionsCountryResponse**](ListDialingPermissionsCountryResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDialingPermissionsHrsPrefixes

> ListDialingPermissionsHrsPrefixesResponse ListDialingPermissionsHrsPrefixes(ctx, IsoCode).PageSize(PageSize).Execute()





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
    IsoCode := "IsoCode_example" // string | The [ISO 3166-1 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) to identify the country permissions from which high-risk special service number prefixes are fetched
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListDialingPermissionsHrsPrefixes(context.Background(), IsoCode).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListDialingPermissionsHrsPrefixes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDialingPermissionsHrsPrefixes`: ListDialingPermissionsHrsPrefixesResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListDialingPermissionsHrsPrefixes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IsoCode** | **string** | The [ISO 3166-1 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) to identify the country permissions from which high-risk special service number prefixes are fetched

### Other Parameters

Other parameters are passed through a pointer to a ListDialingPermissionsHrsPrefixesParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListDialingPermissionsHrsPrefixesResponse**](ListDialingPermissionsHrsPrefixesResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIpRecord

> ListIpRecordResponse ListIpRecord(ctx).PageSize(PageSize).Execute()



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
    resp, r, err := api_client.DefaultApi.ListIpRecord(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListIpRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIpRecord`: ListIpRecordResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListIpRecord`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListIpRecordParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListIpRecordResponse**](ListIpRecordResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSourceIpMapping

> ListSourceIpMappingResponse ListSourceIpMapping(ctx).PageSize(PageSize).Execute()



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
    resp, r, err := api_client.DefaultApi.ListSourceIpMapping(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSourceIpMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSourceIpMapping`: ListSourceIpMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSourceIpMapping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSourceIpMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSourceIpMappingResponse**](ListSourceIpMappingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateByocTrunk

> VoiceV1ByocTrunk UpdateByocTrunk(ctx, Sid).CnamLookupEnabled(CnamLookupEnabled).ConnectionPolicySid(ConnectionPolicySid).FriendlyName(FriendlyName).FromDomainSid(FromDomainSid).StatusCallbackMethod(StatusCallbackMethod).StatusCallbackUrl(StatusCallbackUrl).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceUrl(VoiceUrl).Execute()



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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the BYOC Trunk resource to update.
    CnamLookupEnabled := true // bool | Whether Caller ID Name (CNAM) lookup is enabled for the trunk. If enabled, all inbound calls to the BYOC Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information. (optional)
    ConnectionPolicySid := "ConnectionPolicySid_example" // string | The SID of the Connection Policy that Twilio will use when routing traffic to your communications infrastructure. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. (optional)
    FromDomainSid := "FromDomainSid_example" // string | The SID of the SIP Domain that should be used in the `From` header of originating calls sent to your SIP infrastructure. If your SIP infrastructure allows users to \\\"call back\\\" an incoming call, configure this with a [SIP Domain](https://www.twilio.com/docs/voice/api/sending-sip) to ensure proper routing. If not configured, the from domain will default to \\\"sip.twilio.com\\\". (optional)
    StatusCallbackMethod := "StatusCallbackMethod_example" // string | The HTTP method we should use to call `status_callback_url`. Can be: `GET` or `POST`. (optional)
    StatusCallbackUrl := "StatusCallbackUrl_example" // string | The URL that we should call to pass status parameters (such as call ended) to your application. (optional)
    VoiceFallbackMethod := "VoiceFallbackMethod_example" // string | The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`. (optional)
    VoiceFallbackUrl := "VoiceFallbackUrl_example" // string | The URL that we should call when an error occurs while retrieving or executing the TwiML requested by `voice_url`. (optional)
    VoiceMethod := "VoiceMethod_example" // string | The HTTP method we should use to call `voice_url` (optional)
    VoiceUrl := "VoiceUrl_example" // string | The URL we should call when the BYOC Trunk receives a call. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateByocTrunk(context.Background(), Sid).CnamLookupEnabled(CnamLookupEnabled).ConnectionPolicySid(ConnectionPolicySid).FriendlyName(FriendlyName).FromDomainSid(FromDomainSid).StatusCallbackMethod(StatusCallbackMethod).StatusCallbackUrl(StatusCallbackUrl).VoiceFallbackMethod(VoiceFallbackMethod).VoiceFallbackUrl(VoiceFallbackUrl).VoiceMethod(VoiceMethod).VoiceUrl(VoiceUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateByocTrunk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateByocTrunk`: VoiceV1ByocTrunk
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateByocTrunk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the BYOC Trunk resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateByocTrunkParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **CnamLookupEnabled** | **bool** | Whether Caller ID Name (CNAM) lookup is enabled for the trunk. If enabled, all inbound calls to the BYOC Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
 **ConnectionPolicySid** | **string** | The SID of the Connection Policy that Twilio will use when routing traffic to your communications infrastructure.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
 **FromDomainSid** | **string** | The SID of the SIP Domain that should be used in the &#x60;From&#x60; header of originating calls sent to your SIP infrastructure. If your SIP infrastructure allows users to \\\&quot;call back\\\&quot; an incoming call, configure this with a [SIP Domain](https://www.twilio.com/docs/voice/api/sending-sip) to ensure proper routing. If not configured, the from domain will default to \\\&quot;sip.twilio.com\\\&quot;.
 **StatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;status_callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **StatusCallbackUrl** | **string** | The URL that we should call to pass status parameters (such as call ended) to your application.
 **VoiceFallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
 **VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs while retrieving or executing the TwiML requested by &#x60;voice_url&#x60;.
 **VoiceMethod** | **string** | The HTTP method we should use to call &#x60;voice_url&#x60;
 **VoiceUrl** | **string** | The URL we should call when the BYOC Trunk receives a call.

### Return type

[**VoiceV1ByocTrunk**](VoiceV1ByocTrunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectionPolicy

> VoiceV1ConnectionPolicy UpdateConnectionPolicy(ctx, Sid).FriendlyName(FriendlyName).Execute()



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
    Sid := "Sid_example" // string | The unique string that we created to identify the Connection Policy resource to update.
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateConnectionPolicy(context.Background(), Sid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConnectionPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnectionPolicy`: VoiceV1ConnectionPolicy
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateConnectionPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Connection Policy resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConnectionPolicyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.

### Return type

[**VoiceV1ConnectionPolicy**](VoiceV1ConnectionPolicy.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectionPolicyTarget

> VoiceV1ConnectionPolicyConnectionPolicyTarget UpdateConnectionPolicyTarget(ctx, ConnectionPolicySid, Sid).Enabled(Enabled).FriendlyName(FriendlyName).Priority(Priority).Target(Target).Weight(Weight).Execute()



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
    ConnectionPolicySid := "ConnectionPolicySid_example" // string | The SID of the Connection Policy that owns the Target.
    Sid := "Sid_example" // string | The unique string that we created to identify the Target resource to update.
    Enabled := true // bool | Whether the Target is enabled. (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. (optional)
    Priority := int32(56) // int32 | The relative importance of the target. Can be an integer from 0 to 65535, inclusive. The lowest number represents the most important target. (optional)
    Target := "Target_example" // string | The SIP address you want Twilio to route your calls to. This must be a `sip:` schema. `sips` is NOT supported. (optional)
    Weight := int32(56) // int32 | The value that determines the relative share of the load the Target should receive compared to other Targets with the same priority. Can be an integer from 1 to 65535, inclusive. Targets with higher values receive more load than those with lower ones with the same priority. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateConnectionPolicyTarget(context.Background(), ConnectionPolicySid, Sid).Enabled(Enabled).FriendlyName(FriendlyName).Priority(Priority).Target(Target).Weight(Weight).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateConnectionPolicyTarget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnectionPolicyTarget`: VoiceV1ConnectionPolicyConnectionPolicyTarget
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateConnectionPolicyTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConnectionPolicySid** | **string** | The SID of the Connection Policy that owns the Target.
**Sid** | **string** | The unique string that we created to identify the Target resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConnectionPolicyTargetParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


 **Enabled** | **bool** | Whether the Target is enabled.
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
 **Priority** | **int32** | The relative importance of the target. Can be an integer from 0 to 65535, inclusive. The lowest number represents the most important target.
 **Target** | **string** | The SIP address you want Twilio to route your calls to. This must be a &#x60;sip:&#x60; schema. &#x60;sips&#x60; is NOT supported.
 **Weight** | **int32** | The value that determines the relative share of the load the Target should receive compared to other Targets with the same priority. Can be an integer from 1 to 65535, inclusive. Targets with higher values receive more load than those with lower ones with the same priority.

### Return type

[**VoiceV1ConnectionPolicyConnectionPolicyTarget**](VoiceV1ConnectionPolicyConnectionPolicyTarget.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDialingPermissionsSettings

> VoiceV1DialingPermissionsDialingPermissionsSettings UpdateDialingPermissionsSettings(ctx).DialingPermissionsInheritance(DialingPermissionsInheritance).Execute()





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
    DialingPermissionsInheritance := true // bool | `true` for the sub-account to inherit voice dialing permissions from the Master Project; otherwise `false`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateDialingPermissionsSettings(context.Background()).DialingPermissionsInheritance(DialingPermissionsInheritance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateDialingPermissionsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDialingPermissionsSettings`: VoiceV1DialingPermissionsDialingPermissionsSettings
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateDialingPermissionsSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateDialingPermissionsSettingsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **DialingPermissionsInheritance** | **bool** | &#x60;true&#x60; for the sub-account to inherit voice dialing permissions from the Master Project; otherwise &#x60;false&#x60;.

### Return type

[**VoiceV1DialingPermissionsDialingPermissionsSettings**](VoiceV1DialingPermissionsDialingPermissionsSettings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIpRecord

> VoiceV1IpRecord UpdateIpRecord(ctx, Sid).FriendlyName(FriendlyName).Execute()



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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the IP Record resource to update.
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateIpRecord(context.Background(), Sid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateIpRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIpRecord`: VoiceV1IpRecord
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateIpRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IP Record resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateIpRecordParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.

### Return type

[**VoiceV1IpRecord**](VoiceV1IpRecord.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSourceIpMapping

> VoiceV1SourceIpMapping UpdateSourceIpMapping(ctx, Sid).SipDomainSid(SipDomainSid).Execute()



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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the IP Record resource to update.
    SipDomainSid := "SipDomainSid_example" // string | The SID of the SIP Domain that the IP Record should be mapped to. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSourceIpMapping(context.Background(), Sid).SipDomainSid(SipDomainSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSourceIpMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSourceIpMapping`: VoiceV1SourceIpMapping
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSourceIpMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the IP Record resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSourceIpMappingParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **SipDomainSid** | **string** | The SID of the SIP Domain that the IP Record should be mapped to.

### Return type

[**VoiceV1SourceIpMapping**](VoiceV1SourceIpMapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

