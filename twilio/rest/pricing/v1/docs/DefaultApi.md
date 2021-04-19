# \DefaultApi

All URIs are relative to *https://pricing.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchMessagingCountry**](DefaultApi.md#FetchMessagingCountry) | **Get** /v1/Messaging/Countries/{IsoCountry} | 
[**FetchPhoneNumberCountry**](DefaultApi.md#FetchPhoneNumberCountry) | **Get** /v1/PhoneNumbers/Countries/{IsoCountry} | 
[**FetchVoiceCountry**](DefaultApi.md#FetchVoiceCountry) | **Get** /v1/Voice/Countries/{IsoCountry} | 
[**FetchVoiceNumber**](DefaultApi.md#FetchVoiceNumber) | **Get** /v1/Voice/Numbers/{Number} | 
[**ListMessagingCountry**](DefaultApi.md#ListMessagingCountry) | **Get** /v1/Messaging/Countries | 
[**ListPhoneNumberCountry**](DefaultApi.md#ListPhoneNumberCountry) | **Get** /v1/PhoneNumbers/Countries | 
[**ListVoiceCountry**](DefaultApi.md#ListVoiceCountry) | **Get** /v1/Voice/Countries | 



## FetchMessagingCountry

> PricingV1MessagingMessagingCountryInstance FetchMessagingCountry(ctx, IsoCountry).Execute()



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
    IsoCountry := "IsoCountry_example" // string | The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the pricing information to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchMessagingCountry(context.Background(), IsoCountry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchMessagingCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchMessagingCountry`: PricingV1MessagingMessagingCountryInstance
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchMessagingCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IsoCountry** | **string** | The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the pricing information to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchMessagingCountryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**PricingV1MessagingMessagingCountryInstance**](PricingV1MessagingMessagingCountryInstance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPhoneNumberCountry

> PricingV1PhoneNumberPhoneNumberCountryInstance FetchPhoneNumberCountry(ctx, IsoCountry).Execute()



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
    IsoCountry := "IsoCountry_example" // string | The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the pricing information to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchPhoneNumberCountry(context.Background(), IsoCountry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchPhoneNumberCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchPhoneNumberCountry`: PricingV1PhoneNumberPhoneNumberCountryInstance
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchPhoneNumberCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IsoCountry** | **string** | The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the pricing information to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchPhoneNumberCountryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**PricingV1PhoneNumberPhoneNumberCountryInstance**](PricingV1PhoneNumberPhoneNumberCountryInstance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVoiceCountry

> PricingV1VoiceVoiceCountryInstance FetchVoiceCountry(ctx, IsoCountry).Execute()



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
    IsoCountry := "IsoCountry_example" // string | The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the pricing information to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchVoiceCountry(context.Background(), IsoCountry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchVoiceCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchVoiceCountry`: PricingV1VoiceVoiceCountryInstance
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchVoiceCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IsoCountry** | **string** | The [ISO country code](http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the pricing information to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchVoiceCountryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**PricingV1VoiceVoiceCountryInstance**](PricingV1VoiceVoiceCountryInstance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVoiceNumber

> PricingV1VoiceVoiceNumber FetchVoiceNumber(ctx, Number).Execute()



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
    Number := "Number_example" // string | The phone number to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchVoiceNumber(context.Background(), Number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchVoiceNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchVoiceNumber`: PricingV1VoiceVoiceNumber
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchVoiceNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Number** | **string** | The phone number to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchVoiceNumberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**PricingV1VoiceVoiceNumber**](PricingV1VoiceVoiceNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMessagingCountry

> ListMessagingCountryResponse ListMessagingCountry(ctx).PageSize(PageSize).Execute()



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
    resp, r, err := api_client.DefaultApi.ListMessagingCountry(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMessagingCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMessagingCountry`: ListMessagingCountryResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMessagingCountry`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListMessagingCountryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListMessagingCountryResponse**](ListMessagingCountryResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhoneNumberCountry

> ListPhoneNumberCountryResponse ListPhoneNumberCountry(ctx).PageSize(PageSize).Execute()



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
    resp, r, err := api_client.DefaultApi.ListPhoneNumberCountry(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListPhoneNumberCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPhoneNumberCountry`: ListPhoneNumberCountryResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListPhoneNumberCountry`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListPhoneNumberCountryParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListPhoneNumberCountryResponse**](ListPhoneNumberCountryResponse.md)

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

