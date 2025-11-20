# ShortCodesApplicationsApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateShortCodeApplication**](ShortCodesApplicationsApi.md#CreateShortCodeApplication) | **Post** /v2/ShortCodes/Applications | Create a new short code application
[**FetchShortCodeApplication**](ShortCodesApplicationsApi.md#FetchShortCodeApplication) | **Get** /v2/ShortCodes/Applications/{sid} | Fetch a specific Short Code Application instance.
[**ListShortCodeApplications**](ShortCodesApplicationsApi.md#ListShortCodeApplications) | **Get** /v2/ShortCodes/Applications | List Short Code Applications



## CreateShortCodeApplication

> CreateShortCodeApplicationResponse CreateShortCodeApplication(ctx, optional)

Create a new short code application

Create a new short code application for an account

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateShortCodeApplicationParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreateShortCodeApplicationRequest** | [**CreateShortCodeApplicationRequest**](CreateShortCodeApplicationRequest.md) | 

### Return type

[**CreateShortCodeApplicationResponse**](CreateShortCodeApplicationResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchShortCodeApplication

> ShortCodeApplication FetchShortCodeApplication(ctx, Sid)

Fetch a specific Short Code Application instance.

Fetch a specific Short Code Application instance.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that identifies the Short Code Application resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchShortCodeApplicationParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ShortCodeApplication**](ShortCodeApplication.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListShortCodeApplications

> []ShortCodeApplication ListShortCodeApplications(ctx, optional)

List Short Code Applications

list of all short code applications for an account

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListShortCodeApplicationsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 50.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ShortCodeApplication**](ShortCodeApplication.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

