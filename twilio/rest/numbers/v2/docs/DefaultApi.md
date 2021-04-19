# \DefaultApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBundle**](DefaultApi.md#CreateBundle) | **Post** /v2/RegulatoryCompliance/Bundles | 
[**CreateEndUser**](DefaultApi.md#CreateEndUser) | **Post** /v2/RegulatoryCompliance/EndUsers | 
[**CreateEvaluation**](DefaultApi.md#CreateEvaluation) | **Post** /v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations | 
[**CreateItemAssignment**](DefaultApi.md#CreateItemAssignment) | **Post** /v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments | 
[**CreateSupportingDocument**](DefaultApi.md#CreateSupportingDocument) | **Post** /v2/RegulatoryCompliance/SupportingDocuments | 
[**DeleteBundle**](DefaultApi.md#DeleteBundle) | **Delete** /v2/RegulatoryCompliance/Bundles/{Sid} | 
[**DeleteEndUser**](DefaultApi.md#DeleteEndUser) | **Delete** /v2/RegulatoryCompliance/EndUsers/{Sid} | 
[**DeleteItemAssignment**](DefaultApi.md#DeleteItemAssignment) | **Delete** /v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments/{Sid} | 
[**DeleteSupportingDocument**](DefaultApi.md#DeleteSupportingDocument) | **Delete** /v2/RegulatoryCompliance/SupportingDocuments/{Sid} | 
[**FetchBundle**](DefaultApi.md#FetchBundle) | **Get** /v2/RegulatoryCompliance/Bundles/{Sid} | 
[**FetchEndUser**](DefaultApi.md#FetchEndUser) | **Get** /v2/RegulatoryCompliance/EndUsers/{Sid} | 
[**FetchEndUserType**](DefaultApi.md#FetchEndUserType) | **Get** /v2/RegulatoryCompliance/EndUserTypes/{Sid} | 
[**FetchEvaluation**](DefaultApi.md#FetchEvaluation) | **Get** /v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations/{Sid} | 
[**FetchItemAssignment**](DefaultApi.md#FetchItemAssignment) | **Get** /v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments/{Sid} | 
[**FetchRegulation**](DefaultApi.md#FetchRegulation) | **Get** /v2/RegulatoryCompliance/Regulations/{Sid} | 
[**FetchSupportingDocument**](DefaultApi.md#FetchSupportingDocument) | **Get** /v2/RegulatoryCompliance/SupportingDocuments/{Sid} | 
[**FetchSupportingDocumentType**](DefaultApi.md#FetchSupportingDocumentType) | **Get** /v2/RegulatoryCompliance/SupportingDocumentTypes/{Sid} | 
[**ListBundle**](DefaultApi.md#ListBundle) | **Get** /v2/RegulatoryCompliance/Bundles | 
[**ListEndUser**](DefaultApi.md#ListEndUser) | **Get** /v2/RegulatoryCompliance/EndUsers | 
[**ListEndUserType**](DefaultApi.md#ListEndUserType) | **Get** /v2/RegulatoryCompliance/EndUserTypes | 
[**ListEvaluation**](DefaultApi.md#ListEvaluation) | **Get** /v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations | 
[**ListItemAssignment**](DefaultApi.md#ListItemAssignment) | **Get** /v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments | 
[**ListRegulation**](DefaultApi.md#ListRegulation) | **Get** /v2/RegulatoryCompliance/Regulations | 
[**ListSupportingDocument**](DefaultApi.md#ListSupportingDocument) | **Get** /v2/RegulatoryCompliance/SupportingDocuments | 
[**ListSupportingDocumentType**](DefaultApi.md#ListSupportingDocumentType) | **Get** /v2/RegulatoryCompliance/SupportingDocumentTypes | 
[**UpdateBundle**](DefaultApi.md#UpdateBundle) | **Post** /v2/RegulatoryCompliance/Bundles/{Sid} | 
[**UpdateEndUser**](DefaultApi.md#UpdateEndUser) | **Post** /v2/RegulatoryCompliance/EndUsers/{Sid} | 
[**UpdateSupportingDocument**](DefaultApi.md#UpdateSupportingDocument) | **Post** /v2/RegulatoryCompliance/SupportingDocuments/{Sid} | 



## CreateBundle

> NumbersV2RegulatoryComplianceBundle CreateBundle(ctx).Email(Email).EndUserType(EndUserType).FriendlyName(FriendlyName).IsoCountry(IsoCountry).NumberType(NumberType).RegulationSid(RegulationSid).StatusCallback(StatusCallback).Execute()





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
    Email := "Email_example" // string | The email address that will receive updates when the Bundle resource changes status. (optional)
    EndUserType := "EndUserType_example" // string | The type of End User of the Bundle resource. (optional)
    FriendlyName := "FriendlyName_example" // string | The string that you assigned to describe the resource. (optional)
    IsoCountry := "IsoCountry_example" // string | The ISO country code of the Bundle's phone number country ownership request. (optional)
    NumberType := "NumberType_example" // string | The type of phone number of the Bundle's ownership request. (optional)
    RegulationSid := "RegulationSid_example" // string | The unique string of a regulation that is associated to the Bundle resource. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we call to inform your application of status changes. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateBundle(context.Background()).Email(Email).EndUserType(EndUserType).FriendlyName(FriendlyName).IsoCountry(IsoCountry).NumberType(NumberType).RegulationSid(RegulationSid).StatusCallback(StatusCallback).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBundle`: NumbersV2RegulatoryComplianceBundle
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateBundle`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateBundleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **Email** | **string** | The email address that will receive updates when the Bundle resource changes status.
 **EndUserType** | **string** | The type of End User of the Bundle resource.
 **FriendlyName** | **string** | The string that you assigned to describe the resource.
 **IsoCountry** | **string** | The ISO country code of the Bundle&#39;s phone number country ownership request.
 **NumberType** | **string** | The type of phone number of the Bundle&#39;s ownership request.
 **RegulationSid** | **string** | The unique string of a regulation that is associated to the Bundle resource.
 **StatusCallback** | **string** | The URL we call to inform your application of status changes.

### Return type

[**NumbersV2RegulatoryComplianceBundle**](NumbersV2RegulatoryComplianceBundle.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEndUser

> NumbersV2RegulatoryComplianceEndUser CreateEndUser(ctx).Attributes(Attributes).FriendlyName(FriendlyName).Type(Type).Execute()





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
    Attributes := TODO // map[string]interface{} | The set of parameters that are the attributes of the End User resource which are derived End User Types. (optional)
    FriendlyName := "FriendlyName_example" // string | The string that you assigned to describe the resource. (optional)
    Type := "Type_example" // string | The type of end user of the Bundle resource - can be `individual` or `business`. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateEndUser(context.Background()).Attributes(Attributes).FriendlyName(FriendlyName).Type(Type).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateEndUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEndUser`: NumbersV2RegulatoryComplianceEndUser
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateEndUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateEndUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **Attributes** | [**map[string]interface{}**](map[string]interface{}.md) | The set of parameters that are the attributes of the End User resource which are derived End User Types.
 **FriendlyName** | **string** | The string that you assigned to describe the resource.
 **Type** | **string** | The type of end user of the Bundle resource - can be &#x60;individual&#x60; or &#x60;business&#x60;.

### Return type

[**NumbersV2RegulatoryComplianceEndUser**](NumbersV2RegulatoryComplianceEndUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEvaluation

> NumbersV2RegulatoryComplianceBundleEvaluation CreateEvaluation(ctx, BundleSid).Execute()





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
    BundleSid := "BundleSid_example" // string | The unique string that identifies the Bundle resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateEvaluation(context.Background(), BundleSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateEvaluation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEvaluation`: NumbersV2RegulatoryComplianceBundleEvaluation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateEvaluation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that identifies the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateEvaluationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**NumbersV2RegulatoryComplianceBundleEvaluation**](NumbersV2RegulatoryComplianceBundleEvaluation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateItemAssignment

> NumbersV2RegulatoryComplianceBundleItemAssignment CreateItemAssignment(ctx, BundleSid).ObjectSid(ObjectSid).Execute()





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
    BundleSid := "BundleSid_example" // string | The unique string that we created to identify the Bundle resource.
    ObjectSid := "ObjectSid_example" // string | The SID of an object bag that holds information of the different items. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateItemAssignment(context.Background(), BundleSid).ObjectSid(ObjectSid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateItemAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateItemAssignment`: NumbersV2RegulatoryComplianceBundleItemAssignment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateItemAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that we created to identify the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a CreateItemAssignmentParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **ObjectSid** | **string** | The SID of an object bag that holds information of the different items.

### Return type

[**NumbersV2RegulatoryComplianceBundleItemAssignment**](NumbersV2RegulatoryComplianceBundleItemAssignment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSupportingDocument

> NumbersV2RegulatoryComplianceSupportingDocument CreateSupportingDocument(ctx).Attributes(Attributes).FriendlyName(FriendlyName).Type(Type).Execute()





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
    Attributes := TODO // map[string]interface{} | The set of parameters that are the attributes of the Supporting Documents resource which are derived Supporting Document Types. (optional)
    FriendlyName := "FriendlyName_example" // string | The string that you assigned to describe the resource. (optional)
    Type := "Type_example" // string | The type of the Supporting Document. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateSupportingDocument(context.Background()).Attributes(Attributes).FriendlyName(FriendlyName).Type(Type).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSupportingDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSupportingDocument`: NumbersV2RegulatoryComplianceSupportingDocument
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSupportingDocument`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSupportingDocumentParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **Attributes** | [**map[string]interface{}**](map[string]interface{}.md) | The set of parameters that are the attributes of the Supporting Documents resource which are derived Supporting Document Types.
 **FriendlyName** | **string** | The string that you assigned to describe the resource.
 **Type** | **string** | The type of the Supporting Document.

### Return type

[**NumbersV2RegulatoryComplianceSupportingDocument**](NumbersV2RegulatoryComplianceSupportingDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBundle

> DeleteBundle(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The unique string that we created to identify the Bundle resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteBundle(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteBundleParams struct via the builder pattern


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


## DeleteEndUser

> DeleteEndUser(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The unique string created by Twilio to identify the End User resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteEndUser(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteEndUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string created by Twilio to identify the End User resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteEndUserParams struct via the builder pattern


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


## DeleteItemAssignment

> DeleteItemAssignment(ctx, BundleSid, Sid).Execute()





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
    BundleSid := "BundleSid_example" // string | The unique string that we created to identify the Bundle resource.
    Sid := "Sid_example" // string | The unique string that we created to identify the Identity resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteItemAssignment(context.Background(), BundleSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteItemAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that we created to identify the Bundle resource.
**Sid** | **string** | The unique string that we created to identify the Identity resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteItemAssignmentParams struct via the builder pattern


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


## DeleteSupportingDocument

> DeleteSupportingDocument(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The unique string created by Twilio to identify the Supporting Document resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteSupportingDocument(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSupportingDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string created by Twilio to identify the Supporting Document resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSupportingDocumentParams struct via the builder pattern


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


## FetchBundle

> NumbersV2RegulatoryComplianceBundle FetchBundle(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The unique string that we created to identify the Bundle resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchBundle(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchBundle`: NumbersV2RegulatoryComplianceBundle
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchBundleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**NumbersV2RegulatoryComplianceBundle**](NumbersV2RegulatoryComplianceBundle.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEndUser

> NumbersV2RegulatoryComplianceEndUser FetchEndUser(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The unique string created by Twilio to identify the End User resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchEndUser(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchEndUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchEndUser`: NumbersV2RegulatoryComplianceEndUser
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchEndUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string created by Twilio to identify the End User resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchEndUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**NumbersV2RegulatoryComplianceEndUser**](NumbersV2RegulatoryComplianceEndUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEndUserType

> NumbersV2RegulatoryComplianceEndUserType FetchEndUserType(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The unique string that identifies the End-User Type resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchEndUserType(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchEndUserType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchEndUserType`: NumbersV2RegulatoryComplianceEndUserType
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchEndUserType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that identifies the End-User Type resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchEndUserTypeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**NumbersV2RegulatoryComplianceEndUserType**](NumbersV2RegulatoryComplianceEndUserType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEvaluation

> NumbersV2RegulatoryComplianceBundleEvaluation FetchEvaluation(ctx, BundleSid, Sid).Execute()





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
    BundleSid := "BundleSid_example" // string | The unique string that we created to identify the Bundle resource.
    Sid := "Sid_example" // string | The unique string that identifies the Evaluation resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchEvaluation(context.Background(), BundleSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchEvaluation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchEvaluation`: NumbersV2RegulatoryComplianceBundleEvaluation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchEvaluation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that we created to identify the Bundle resource.
**Sid** | **string** | The unique string that identifies the Evaluation resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchEvaluationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**NumbersV2RegulatoryComplianceBundleEvaluation**](NumbersV2RegulatoryComplianceBundleEvaluation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchItemAssignment

> NumbersV2RegulatoryComplianceBundleItemAssignment FetchItemAssignment(ctx, BundleSid, Sid).Execute()





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
    BundleSid := "BundleSid_example" // string | The unique string that we created to identify the Bundle resource.
    Sid := "Sid_example" // string | The unique string that we created to identify the Identity resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchItemAssignment(context.Background(), BundleSid, Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchItemAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchItemAssignment`: NumbersV2RegulatoryComplianceBundleItemAssignment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchItemAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that we created to identify the Bundle resource.
**Sid** | **string** | The unique string that we created to identify the Identity resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchItemAssignmentParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------



### Return type

[**NumbersV2RegulatoryComplianceBundleItemAssignment**](NumbersV2RegulatoryComplianceBundleItemAssignment.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRegulation

> NumbersV2RegulatoryComplianceRegulation FetchRegulation(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The unique string that identifies the Regulation resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchRegulation(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchRegulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchRegulation`: NumbersV2RegulatoryComplianceRegulation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchRegulation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that identifies the Regulation resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchRegulationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**NumbersV2RegulatoryComplianceRegulation**](NumbersV2RegulatoryComplianceRegulation.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSupportingDocument

> NumbersV2RegulatoryComplianceSupportingDocument FetchSupportingDocument(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The unique string created by Twilio to identify the Supporting Document resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSupportingDocument(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSupportingDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSupportingDocument`: NumbersV2RegulatoryComplianceSupportingDocument
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSupportingDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string created by Twilio to identify the Supporting Document resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchSupportingDocumentParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**NumbersV2RegulatoryComplianceSupportingDocument**](NumbersV2RegulatoryComplianceSupportingDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSupportingDocumentType

> NumbersV2RegulatoryComplianceSupportingDocumentType FetchSupportingDocumentType(ctx, Sid).Execute()





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
    Sid := "Sid_example" // string | The unique string that identifies the Supporting Document Type resource.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.FetchSupportingDocumentType(context.Background(), Sid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.FetchSupportingDocumentType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchSupportingDocumentType`: NumbersV2RegulatoryComplianceSupportingDocumentType
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.FetchSupportingDocumentType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that identifies the Supporting Document Type resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchSupportingDocumentTypeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------


### Return type

[**NumbersV2RegulatoryComplianceSupportingDocumentType**](NumbersV2RegulatoryComplianceSupportingDocumentType.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBundle

> ListBundleResponse ListBundle(ctx).Status(Status).FriendlyName(FriendlyName).RegulationSid(RegulationSid).IsoCountry(IsoCountry).NumberType(NumberType).PageSize(PageSize).Execute()





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
    Status := "Status_example" // string | The verification status of the Bundle resource. (optional)
    FriendlyName := "FriendlyName_example" // string | The string that you assigned to describe the resource. (optional)
    RegulationSid := "RegulationSid_example" // string | The unique string of a regulation that is associated to the Bundle resource. (optional)
    IsoCountry := "IsoCountry_example" // string | The ISO country code of the Bundle's phone number country ownership request. (optional)
    NumberType := "NumberType_example" // string | The type of phone number of the Bundle's ownership request. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListBundle(context.Background()).Status(Status).FriendlyName(FriendlyName).RegulationSid(RegulationSid).IsoCountry(IsoCountry).NumberType(NumberType).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBundle`: ListBundleResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListBundle`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListBundleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **Status** | **string** | The verification status of the Bundle resource.
 **FriendlyName** | **string** | The string that you assigned to describe the resource.
 **RegulationSid** | **string** | The unique string of a regulation that is associated to the Bundle resource.
 **IsoCountry** | **string** | The ISO country code of the Bundle&#39;s phone number country ownership request.
 **NumberType** | **string** | The type of phone number of the Bundle&#39;s ownership request.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListBundleResponse**](ListBundleResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEndUser

> ListEndUserResponse ListEndUser(ctx).PageSize(PageSize).Execute()





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
    resp, r, err := api_client.DefaultApi.ListEndUser(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListEndUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEndUser`: ListEndUserResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListEndUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListEndUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListEndUserResponse**](ListEndUserResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEndUserType

> ListEndUserTypeResponse ListEndUserType(ctx).PageSize(PageSize).Execute()





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
    resp, r, err := api_client.DefaultApi.ListEndUserType(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListEndUserType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEndUserType`: ListEndUserTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListEndUserType`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListEndUserTypeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListEndUserTypeResponse**](ListEndUserTypeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvaluation

> ListEvaluationResponse ListEvaluation(ctx, BundleSid).PageSize(PageSize).Execute()





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
    BundleSid := "BundleSid_example" // string | The unique string that identifies the Bundle resource.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListEvaluation(context.Background(), BundleSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListEvaluation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEvaluation`: ListEvaluationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListEvaluation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that identifies the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a ListEvaluationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListEvaluationResponse**](ListEvaluationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListItemAssignment

> ListItemAssignmentResponse ListItemAssignment(ctx, BundleSid).PageSize(PageSize).Execute()





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
    BundleSid := "BundleSid_example" // string | The unique string that we created to identify the Bundle resource.
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListItemAssignment(context.Background(), BundleSid).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListItemAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListItemAssignment`: ListItemAssignmentResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListItemAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BundleSid** | **string** | The unique string that we created to identify the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a ListItemAssignmentParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListItemAssignmentResponse**](ListItemAssignmentResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRegulation

> ListRegulationResponse ListRegulation(ctx).EndUserType(EndUserType).IsoCountry(IsoCountry).NumberType(NumberType).PageSize(PageSize).Execute()





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
    EndUserType := "EndUserType_example" // string | The type of End User the regulation requires - can be `individual` or `business`. (optional)
    IsoCountry := "IsoCountry_example" // string | The ISO country code of the phone number's country. (optional)
    NumberType := "NumberType_example" // string | The type of phone number that the regulatory requiremnt is restricting. (optional)
    PageSize := int32(56) // int32 | How many resources to return in each list page. The default is 50, and the maximum is 1000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListRegulation(context.Background()).EndUserType(EndUserType).IsoCountry(IsoCountry).NumberType(NumberType).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRegulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRegulation`: ListRegulationResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRegulation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListRegulationParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **EndUserType** | **string** | The type of End User the regulation requires - can be &#x60;individual&#x60; or &#x60;business&#x60;.
 **IsoCountry** | **string** | The ISO country code of the phone number&#39;s country.
 **NumberType** | **string** | The type of phone number that the regulatory requiremnt is restricting.
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListRegulationResponse**](ListRegulationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportingDocument

> ListSupportingDocumentResponse ListSupportingDocument(ctx).PageSize(PageSize).Execute()





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
    resp, r, err := api_client.DefaultApi.ListSupportingDocument(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSupportingDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSupportingDocument`: ListSupportingDocumentResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSupportingDocument`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSupportingDocumentParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSupportingDocumentResponse**](ListSupportingDocumentResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSupportingDocumentType

> ListSupportingDocumentTypeResponse ListSupportingDocumentType(ctx).PageSize(PageSize).Execute()





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
    resp, r, err := api_client.DefaultApi.ListSupportingDocumentType(context.Background()).PageSize(PageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSupportingDocumentType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSupportingDocumentType`: ListSupportingDocumentTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSupportingDocumentType`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSupportingDocumentTypeParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------
 **PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListSupportingDocumentTypeResponse**](ListSupportingDocumentTypeResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBundle

> NumbersV2RegulatoryComplianceBundle UpdateBundle(ctx, Sid).Email(Email).FriendlyName(FriendlyName).Status(Status).StatusCallback(StatusCallback).Execute()





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
    Sid := "Sid_example" // string | The unique string that we created to identify the Bundle resource.
    Email := "Email_example" // string | The email address that will receive updates when the Bundle resource changes status. (optional)
    FriendlyName := "FriendlyName_example" // string | The string that you assigned to describe the resource. (optional)
    Status := "Status_example" // string | The verification status of the Bundle resource. (optional)
    StatusCallback := "StatusCallback_example" // string | The URL we call to inform your application of status changes. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateBundle(context.Background(), Sid).Email(Email).FriendlyName(FriendlyName).Status(Status).StatusCallback(StatusCallback).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateBundle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBundle`: NumbersV2RegulatoryComplianceBundle
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Bundle resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateBundleParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Email** | **string** | The email address that will receive updates when the Bundle resource changes status.
 **FriendlyName** | **string** | The string that you assigned to describe the resource.
 **Status** | **string** | The verification status of the Bundle resource.
 **StatusCallback** | **string** | The URL we call to inform your application of status changes.

### Return type

[**NumbersV2RegulatoryComplianceBundle**](NumbersV2RegulatoryComplianceBundle.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndUser

> NumbersV2RegulatoryComplianceEndUser UpdateEndUser(ctx, Sid).Attributes(Attributes).FriendlyName(FriendlyName).Execute()





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
    Sid := "Sid_example" // string | The unique string created by Twilio to identify the End User resource.
    Attributes := TODO // map[string]interface{} | The set of parameters that are the attributes of the End User resource which are derived End User Types. (optional)
    FriendlyName := "FriendlyName_example" // string | The string that you assigned to describe the resource. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateEndUser(context.Background(), Sid).Attributes(Attributes).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateEndUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEndUser`: NumbersV2RegulatoryComplianceEndUser
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateEndUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string created by Twilio to identify the End User resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateEndUserParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Attributes** | [**map[string]interface{}**](map[string]interface{}.md) | The set of parameters that are the attributes of the End User resource which are derived End User Types.
 **FriendlyName** | **string** | The string that you assigned to describe the resource.

### Return type

[**NumbersV2RegulatoryComplianceEndUser**](NumbersV2RegulatoryComplianceEndUser.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSupportingDocument

> NumbersV2RegulatoryComplianceSupportingDocument UpdateSupportingDocument(ctx, Sid).Attributes(Attributes).FriendlyName(FriendlyName).Execute()





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
    Sid := "Sid_example" // string | The unique string created by Twilio to identify the Supporting Document resource.
    Attributes := TODO // map[string]interface{} | The set of parameters that are the attributes of the Supporting Document resource which are derived Supporting Document Types. (optional)
    FriendlyName := "FriendlyName_example" // string | The string that you assigned to describe the resource. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateSupportingDocument(context.Background(), Sid).Attributes(Attributes).FriendlyName(FriendlyName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSupportingDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSupportingDocument`: NumbersV2RegulatoryComplianceSupportingDocument
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSupportingDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string created by Twilio to identify the Supporting Document resource.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSupportingDocumentParams struct via the builder pattern


Name | Type | Description
------------- | ------------- | -------------

 **Attributes** | [**map[string]interface{}**](map[string]interface{}.md) | The set of parameters that are the attributes of the Supporting Document resource which are derived Supporting Document Types.
 **FriendlyName** | **string** | The string that you assigned to describe the resource.

### Return type

[**NumbersV2RegulatoryComplianceSupportingDocument**](NumbersV2RegulatoryComplianceSupportingDocument.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

