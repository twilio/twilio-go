# ServicesEntitiesApi

All URIs are relative to *https://verify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEntity**](ServicesEntitiesApi.md#CreateEntity) | **Post** /v2/Services/{ServiceSid}/Entities | Create a new Entity for the Service
[**DeleteEntity**](ServicesEntitiesApi.md#DeleteEntity) | **Delete** /v2/Services/{ServiceSid}/Entities/{Identity} | Delete a specific Entity.
[**FetchEntity**](ServicesEntitiesApi.md#FetchEntity) | **Get** /v2/Services/{ServiceSid}/Entities/{Identity} | Fetch a specific Entity.
[**ListEntity**](ServicesEntitiesApi.md#ListEntity) | **Get** /v2/Services/{ServiceSid}/Entities | Retrieve a list of all Entities for a Service.



## CreateEntity

> VerifyV2Entity CreateEntity(ctx, ServiceSidoptional)

Create a new Entity for the Service

Create a new Entity for the Service

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.

### Other Parameters

Other parameters are passed through a pointer to a CreateEntityParams struct


Name | Type | Description
------------- | ------------- | -------------
**Identity** | **string** | The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, length between 8 and 64 characters, and generated by your external system, such as your user's UUID, GUID, or SID. It can only contain dash (-) separated alphanumeric characters.

### Return type

[**VerifyV2Entity**](VerifyV2Entity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEntity

> DeleteEntity(ctx, ServiceSidIdentity)

Delete a specific Entity.

Delete a specific Entity.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Identity** | **string** | The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, length between 8 and 64 characters, and generated by your external system, such as your user's UUID, GUID, or SID. It can only contain dash (-) separated alphanumeric characters.

### Other Parameters

Other parameters are passed through a pointer to a DeleteEntityParams struct


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


## FetchEntity

> VerifyV2Entity FetchEntity(ctx, ServiceSidIdentity)

Fetch a specific Entity.

Fetch a specific Entity.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.
**Identity** | **string** | The unique external identifier for the Entity of the Service. This identifier should be immutable, not PII, length between 8 and 64 characters, and generated by your external system, such as your user's UUID, GUID, or SID. It can only contain dash (-) separated alphanumeric characters.

### Other Parameters

Other parameters are passed through a pointer to a FetchEntityParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VerifyV2Entity**](VerifyV2Entity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntity

> []VerifyV2Entity ListEntity(ctx, ServiceSidoptional)

Retrieve a list of all Entities for a Service.

Retrieve a list of all Entities for a Service.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The unique SID identifier of the Service.

### Other Parameters

Other parameters are passed through a pointer to a ListEntityParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VerifyV2Entity**](VerifyV2Entity.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

