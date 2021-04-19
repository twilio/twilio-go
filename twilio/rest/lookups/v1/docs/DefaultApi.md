# \DefaultApi

All URIs are relative to *https://lookups.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPhoneNumber**](DefaultApi.md#FetchPhoneNumber) | **Get** /v1/PhoneNumbers/{PhoneNumber} | 



## FetchPhoneNumber

> LookupsV1PhoneNumber FetchPhoneNumber(ctx, PhoneNumber).CountryCode(CountryCode).Type(Type).AddOns(AddOns).AddOnsData(AddOnsData).Execute()



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
    PhoneNumber := "PhoneNumber_example" // string | The phone number to lookup in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number.
    CountryCode := "CountryCode_example" // string | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the phone number to fetch. This is used to specify the country when the phone number is provided in a national format. (optional)
    Type := []string{"Inner_example"} // []string | The type of information to return. Can be: `carrier` or `caller-name`. The default is null.  Carrier information costs $0.005 per phone number looked up.  Caller Name information is currently available only in the US and costs $0.01 per phone number looked up.  To retrieve both types on information, specify this parameter twice; once with `carrier` and once with `caller-name` as the value. (optional)
    AddOns := []string{"Inner_example"} // []string | The `unique_name` of an Add-on you would like to invoke. Can be the `unique_name` of an Add-on that is installed on your account. You can specify multiple instances of this parameter to invoke multiple Add-ons. For more information about  Add-ons, see the [Add-ons documentation](https://www.twilio.com/docs/add-ons). (optional)
    AddOnsData := TODO // map[string]interface{} | Data specific to the add-on you would like to invoke. The content and format of this value depends on the add-on. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchPhoneNumber(context.Background(), PhoneNumber).CountryCode(CountryCode).Type(Type).AddOns(AddOns).AddOnsData(AddOnsData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchPhoneNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchPhoneNumber`: LookupsV1PhoneNumber
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchPhoneNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**PhoneNumber** | **string** | The phone number to lookup in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number.

### Other Parameters

Other parameters are passed through a pointer to a FetchPhoneNumberParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **CountryCode** | **string** | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the phone number to fetch. This is used to specify the country when the phone number is provided in a national format.
 **Type** | **[]string** | The type of information to return. Can be: &#x60;carrier&#x60; or &#x60;caller-name&#x60;. The default is null.  Carrier information costs $0.005 per phone number looked up.  Caller Name information is currently available only in the US and costs $0.01 per phone number looked up.  To retrieve both types on information, specify this parameter twice; once with &#x60;carrier&#x60; and once with &#x60;caller-name&#x60; as the value.
 **AddOns** | **[]string** | The &#x60;unique_name&#x60; of an Add-on you would like to invoke. Can be the &#x60;unique_name&#x60; of an Add-on that is installed on your account. You can specify multiple instances of this parameter to invoke multiple Add-ons. For more information about  Add-ons, see the [Add-ons documentation](https://www.twilio.com/docs/add-ons).
 **AddOnsData** | [**map[string]interface{}**](map[string]interface{}map[string]interface{}.md) | Data specific to the add-on you would like to invoke. The content and format of this value depends on the add-on.

### Return type

[**LookupsV1PhoneNumber**](LookupsV1PhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

