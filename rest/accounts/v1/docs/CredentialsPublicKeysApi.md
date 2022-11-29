# CredentialsPublicKeysApi

All URIs are relative to *https://accounts.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredentialPublicKey**](CredentialsPublicKeysApi.md#CreateCredentialPublicKey) | **Post** /v1/Credentials/PublicKeys | 
[**DeleteCredentialPublicKey**](CredentialsPublicKeysApi.md#DeleteCredentialPublicKey) | **Delete** /v1/Credentials/PublicKeys/{Sid} | 
[**FetchCredentialPublicKey**](CredentialsPublicKeysApi.md#FetchCredentialPublicKey) | **Get** /v1/Credentials/PublicKeys/{Sid} | 
[**ListCredentialPublicKey**](CredentialsPublicKeysApi.md#ListCredentialPublicKey) | **Get** /v1/Credentials/PublicKeys | 
[**UpdateCredentialPublicKey**](CredentialsPublicKeysApi.md#UpdateCredentialPublicKey) | **Post** /v1/Credentials/PublicKeys/{Sid} | 



## CreateCredentialPublicKey

> AccountsV1CredentialPublicKey CreateCredentialPublicKey(ctx, optional)



Create a new Public Key Credential

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCredentialPublicKeyParams struct


Name | Type | Description
------------- | ------------- | -------------
**PublicKey** | **string** | A URL encoded representation of the public key. For example, `-----BEGIN PUBLIC KEY-----MIIBIjANB.pa9xQIDAQAB-----END PUBLIC KEY-----`
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**AccountSid** | **string** | The SID of the Subaccount that this Credential should be associated with. Must be a valid Subaccount of the account issuing the request

### Return type

[**AccountsV1CredentialPublicKey**](AccountsV1CredentialPublicKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCredentialPublicKey

> DeleteCredentialPublicKey(ctx, Sid)



Delete a Credential from your account

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the PublicKey resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCredentialPublicKeyParams struct


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


## FetchCredentialPublicKey

> AccountsV1CredentialPublicKey FetchCredentialPublicKey(ctx, Sid)



Fetch the public key specified by the provided Credential Sid

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the PublicKey resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchCredentialPublicKeyParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AccountsV1CredentialPublicKey**](AccountsV1CredentialPublicKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredentialPublicKey

> []AccountsV1CredentialPublicKey ListCredentialPublicKey(ctx, optional)



Retrieves a collection of Public Key Credentials belonging to the account used to make the request

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCredentialPublicKeyParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AccountsV1CredentialPublicKey**](AccountsV1CredentialPublicKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredentialPublicKey

> AccountsV1CredentialPublicKey UpdateCredentialPublicKey(ctx, Sidoptional)



Modify the properties of a given Account

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the PublicKey resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateCredentialPublicKeyParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.

### Return type

[**AccountsV1CredentialPublicKey**](AccountsV1CredentialPublicKey.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

