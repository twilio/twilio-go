# \DefaultApi

All URIs are relative to *https://pricing.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchVoiceCountry**](DefaultApi.md#FetchVoiceCountry) | **Get** /v2/Voice/Countries/{IsoCountry} | 
[**FetchVoiceNumber**](DefaultApi.md#FetchVoiceNumber) | **Get** /v2/Voice/Numbers/{DestinationNumber} | 
[**ListVoiceCountry**](DefaultApi.md#ListVoiceCountry) | **Get** /v2/Voice/Countries | 



## FetchVoiceCountry

> PricingV2VoiceVoiceCountryInstance FetchVoiceCountry(ctx, IsoCountry).Execute()





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
    IsoCountry := "IsoCountry_example" // string | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the origin-based voice pricing information to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchVoiceCountry(context.Background(), IsoCountry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchVoiceCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchVoiceCountry`: PricingV2VoiceVoiceCountryInstance
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchVoiceCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IsoCountry** | **string** | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the origin-based voice pricing information to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchVoiceCountryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**PricingV2VoiceVoiceCountryInstance**](PricingV2VoiceVoiceCountryInstance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVoiceNumber

> PricingV2VoiceVoiceNumber FetchVoiceNumber(ctx, DestinationNumber).OriginationNumber(OriginationNumber).Execute()





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
    DestinationNumber := "DestinationNumber_example" // string | The destination phone number, in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, for which to fetch the origin-based voice pricing information. E.164 format consists of a + followed by the country code and subscriber number.
    OriginationNumber := "OriginationNumber_example" // string | The origination phone number, in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, for which to fetch the origin-based voice pricing information. E.164 format consists of a + followed by the country code and subscriber number. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchVoiceNumber(context.Background(), DestinationNumber).OriginationNumber(OriginationNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchVoiceNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchVoiceNumber`: PricingV2VoiceVoiceNumber
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchVoiceNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**DestinationNumber** | **string** | The destination phone number, in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, for which to fetch the origin-based voice pricing information. E.164 format consists of a + followed by the country code and subscriber number.

### Other Parameters

Other parameters are passed through a pointer to a FetchVoiceNumberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **OriginationNumber** | **string** | The origination phone number, in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, for which to fetch the origin-based voice pricing information. E.164 format consists of a + followed by the country code and subscriber number.

### Return type

[**PricingV2VoiceVoiceNumber**](PricingV2VoiceVoiceNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVoiceCountry

> ListVoiceCountryResponse ListVoiceCountry(ctx).PageSize(PageSize).Execute()



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
    resp, r, err := api_client.DefaultApi.ListVoiceCountry(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListVoiceCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVoiceCountry`: ListVoiceCountryResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListVoiceCountry`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListVoiceCountryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListVoiceCountryResponse**](ListVoiceCountryResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

