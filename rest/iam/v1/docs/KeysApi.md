# KeysApi

All URIs are relative to *https://iam.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewKey**](KeysApi.md#CreateNewKey) | **Post** /v1/Keys | 
[**DeleteKey**](KeysApi.md#DeleteKey) | **Delete** /v1/Keys/{Sid} | 
[**FetchKey**](KeysApi.md#FetchKey) | **Get** /v1/Keys/{Sid} | 
[**ListGetKeys**](KeysApi.md#ListGetKeys) | **Get** /v1/Keys | 
[**UpdateKey**](KeysApi.md#UpdateKey) | **Post** /v1/Keys/{Sid} | 



## CreateNewKey

> IamV1NewKey CreateNewKey(ctx, optional)



Create a new Signing Key for the account making the request.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateNewKeyParams struct


Name | Type | Description
------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Payments resource.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**KeyType** | **string** | 
**Policy** | [**interface{}**](interface{}.md) | The \\\\`Policy\\\\` object is a collection that specifies the allowed Twilio permissions for the restricted key. For more information on the permissions available with restricted API keys, refer to the [Twilio documentation](https://www.twilio.com/docs/iam/api-keys/restricted-api-keys#permissions-available-with-restricted-api-keys).

### Return type

[**IamV1NewKey**](IamV1NewKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKey

> DeleteKey(ctx, Sid)



Delete a specific Key.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Key resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteKeyParams struct


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


## FetchKey

> IamV1Key FetchKey(ctx, Sid)



Fetch a specific Key.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Key resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchKeyParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**IamV1Key**](IamV1Key.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGetKeys

> []IamV1GetKeys ListGetKeys(ctx, optional)



Retrieve a list of all Keys for a account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListGetKeysParams struct


Name | Type | Description
------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Payments resource.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]IamV1GetKeys**](IamV1GetKeys.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKey

> IamV1Key UpdateKey(ctx, Sidoptional)



Update a specific Key.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Key resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateKeyParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**Policy** | [**interface{}**](interface{}.md) | The \\\\`Policy\\\\` object is a collection that specifies the allowed Twilio permissions for the restricted key. For more information on the permissions available with restricted API keys, refer to the [Twilio documentation](https://www.twilio.com/docs/iam/api-keys/restricted-api-keys#permissions-available-with-restricted-api-keys).

### Return type

[**IamV1Key**](IamV1Key.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

