# \DefaultApi

All URIs are relative to *https://accounts.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredentialAws**](DefaultApi.md#CreateCredentialAws) | **Post** /v1/Credentials/AWS | 
[**CreateCredentialPublicKey**](DefaultApi.md#CreateCredentialPublicKey) | **Post** /v1/Credentials/PublicKeys | 
[**CreateSecondaryAuthToken**](DefaultApi.md#CreateSecondaryAuthToken) | **Post** /v1/AuthTokens/Secondary | 
[**DeleteCredentialAws**](DefaultApi.md#DeleteCredentialAws) | **Delete** /v1/Credentials/AWS/{Sid} | 
[**DeleteCredentialPublicKey**](DefaultApi.md#DeleteCredentialPublicKey) | **Delete** /v1/Credentials/PublicKeys/{Sid} | 
[**DeleteSecondaryAuthToken**](DefaultApi.md#DeleteSecondaryAuthToken) | **Delete** /v1/AuthTokens/Secondary | 
[**FetchCredentialAws**](DefaultApi.md#FetchCredentialAws) | **Get** /v1/Credentials/AWS/{Sid} | 
[**FetchCredentialPublicKey**](DefaultApi.md#FetchCredentialPublicKey) | **Get** /v1/Credentials/PublicKeys/{Sid} | 
[**ListCredentialAws**](DefaultApi.md#ListCredentialAws) | **Get** /v1/Credentials/AWS | 
[**ListCredentialPublicKey**](DefaultApi.md#ListCredentialPublicKey) | **Get** /v1/Credentials/PublicKeys | 
[**UpdateAuthTokenPromotion**](DefaultApi.md#UpdateAuthTokenPromotion) | **Post** /v1/AuthTokens/Promote | 
[**UpdateCredentialAws**](DefaultApi.md#UpdateCredentialAws) | **Post** /v1/Credentials/AWS/{Sid} | 
[**UpdateCredentialPublicKey**](DefaultApi.md#UpdateCredentialPublicKey) | **Post** /v1/Credentials/PublicKeys/{Sid} | 



## CreateCredentialAws

> AccountsV1CredentialCredentialAws CreateCredentialAws(ctx).AccountSid(AccountSid).Credentials(Credentials).FriendlyName(FriendlyName).Execute()





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
    AccountSid := "AccountSid_example" // string | The SID of the Subaccount that this Credential should be associated with. Must be a valid Subaccount of the account issuing the request. (optional)
    Credentials := "Credentials_example" // string | A string that contains the AWS access credentials in the format `<AWS_ACCESS_KEY_ID>:<AWS_SECRET_ACCESS_KEY>`. For example, `AKIAIOSFODNN7EXAMPLE:wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY` (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateCredentialAws(context.Background()).AccountSid(AccountSid).Credentials(Credentials).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCredentialAws``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCredentialAws`: AccountsV1CredentialCredentialAws
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCredentialAws`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCredentialAwsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **AccountSid** | **string** | The SID of the Subaccount that this Credential should be associated with. Must be a valid Subaccount of the account issuing the request.
 **Credentials** | **string** | A string that contains the AWS access credentials in the format &#x60;&lt;AWS_ACCESS_KEY_ID&gt;:&lt;AWS_SECRET_ACCESS_KEY&gt;&#x60;. For example, &#x60;AKIAIOSFODNN7EXAMPLE:wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY&#x60;
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.

### Return type

[**AccountsV1CredentialCredentialAws**](AccountsV1CredentialCredentialAws.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCredentialPublicKey

> AccountsV1CredentialCredentialPublicKey CreateCredentialPublicKey(ctx).AccountSid(AccountSid).FriendlyName(FriendlyName).PublicKey(PublicKey).Execute()





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
    AccountSid := "AccountSid_example" // string | The SID of the Subaccount that this Credential should be associated with. Must be a valid Subaccount of the account issuing the request (optional)
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)
    PublicKey := "PublicKey_example" // string | A URL encoded representation of the public key. For example, `-----BEGIN PUBLIC KEY-----MIIBIjANB.pa9xQIDAQAB-----END PUBLIC KEY-----` (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateCredentialPublicKey(context.Background()).AccountSid(AccountSid).FriendlyName(FriendlyName).PublicKey(PublicKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCredentialPublicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCredentialPublicKey`: AccountsV1CredentialCredentialPublicKey
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCredentialPublicKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCredentialPublicKeyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **AccountSid** | **string** | The SID of the Subaccount that this Credential should be associated with. Must be a valid Subaccount of the account issuing the request
 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
 **PublicKey** | **string** | A URL encoded representation of the public key. For example, &#x60;-----BEGIN PUBLIC KEY-----MIIBIjANB.pa9xQIDAQAB-----END PUBLIC KEY-----&#x60;

### Return type

[**AccountsV1CredentialCredentialPublicKey**](AccountsV1CredentialCredentialPublicKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecondaryAuthToken

> AccountsV1SecondaryAuthToken CreateSecondaryAuthToken(ctx).Execute()





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
    resp, r, err := api_client.DefaultApi.CreateSecondaryAuthToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSecondaryAuthToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSecondaryAuthToken`: AccountsV1SecondaryAuthToken
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSecondaryAuthToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSecondaryAuthTokenParams struct via the builder pattern


### Return type

[**AccountsV1SecondaryAuthToken**](AccountsV1SecondaryAuthToken.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCredentialAws

> DeleteCredentialAws(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the AWS resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteCredentialAws(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCredentialAws``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the AWS resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCredentialAwsParams struct via the builder pattern


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


## DeleteCredentialPublicKey

> DeleteCredentialPublicKey(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the PublicKey resource to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteCredentialPublicKey(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCredentialPublicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the PublicKey resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCredentialPublicKeyParams struct via the builder pattern


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


## DeleteSecondaryAuthToken

> DeleteSecondaryAuthToken(ctx).Execute()





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
    resp, r, err := api_client.DefaultApi.DeleteSecondaryAuthToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSecondaryAuthToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSecondaryAuthTokenParams struct via the builder pattern


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


## FetchCredentialAws

> AccountsV1CredentialCredentialAws FetchCredentialAws(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the AWS resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchCredentialAws(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCredentialAws``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCredentialAws`: AccountsV1CredentialCredentialAws
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCredentialAws`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the AWS resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchCredentialAwsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**AccountsV1CredentialCredentialAws**](AccountsV1CredentialCredentialAws.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCredentialPublicKey

> AccountsV1CredentialCredentialPublicKey FetchCredentialPublicKey(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the PublicKey resource to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchCredentialPublicKey(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchCredentialPublicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchCredentialPublicKey`: AccountsV1CredentialCredentialPublicKey
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchCredentialPublicKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the PublicKey resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchCredentialPublicKeyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**AccountsV1CredentialCredentialPublicKey**](AccountsV1CredentialCredentialPublicKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredentialAws

> ListCredentialAwsResponse ListCredentialAws(ctx).PageSize(PageSize).Execute()





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
    resp, r, err := api_client.DefaultApi.ListCredentialAws(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCredentialAws``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCredentialAws`: ListCredentialAwsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCredentialAws`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCredentialAwsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListCredentialAwsResponse**](ListCredentialAwsResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredentialPublicKey

> ListCredentialPublicKeyResponse ListCredentialPublicKey(ctx).PageSize(PageSize).Execute()





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
    resp, r, err := api_client.DefaultApi.ListCredentialPublicKey(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCredentialPublicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCredentialPublicKey`: ListCredentialPublicKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCredentialPublicKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCredentialPublicKeyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListCredentialPublicKeyResponse**](ListCredentialPublicKeyResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthTokenPromotion

> AccountsV1AuthTokenPromotion UpdateAuthTokenPromotion(ctx).Execute()





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
    resp, r, err := api_client.DefaultApi.UpdateAuthTokenPromotion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAuthTokenPromotion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthTokenPromotion`: AccountsV1AuthTokenPromotion
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAuthTokenPromotion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateAuthTokenPromotionParams struct via the builder pattern


### Return type

[**AccountsV1AuthTokenPromotion**](AccountsV1AuthTokenPromotion.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredentialAws

> AccountsV1CredentialCredentialAws UpdateCredentialAws(ctx, Sid).FriendlyName(FriendlyName).Execute()





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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the AWS resource to update.
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateCredentialAws(context.Background(), Sid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateCredentialAws``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCredentialAws`: AccountsV1CredentialCredentialAws
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateCredentialAws`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the AWS resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateCredentialAwsParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.

### Return type

[**AccountsV1CredentialCredentialAws**](AccountsV1CredentialCredentialAws.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredentialPublicKey

> AccountsV1CredentialCredentialPublicKey UpdateCredentialPublicKey(ctx, Sid).FriendlyName(FriendlyName).Execute()





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
    Sid := "Sid_example" // string | The Twilio-provided string that uniquely identifies the PublicKey resource to update.
    FriendlyName := "FriendlyName_example" // string | A descriptive string that you create to describe the resource. It can be up to 64 characters long. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateCredentialPublicKey(context.Background(), Sid).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateCredentialPublicKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCredentialPublicKey`: AccountsV1CredentialCredentialPublicKey
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateCredentialPublicKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the PublicKey resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateCredentialPublicKeyParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.

### Return type

[**AccountsV1CredentialCredentialPublicKey**](AccountsV1CredentialCredentialPublicKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

