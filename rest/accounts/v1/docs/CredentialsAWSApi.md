# CredentialsAWSApi

All URIs are relative to *https://accounts.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredentialAws**](CredentialsAWSApi.md#CreateCredentialAws) | **Post** /v1/Credentials/AWS | 
[**DeleteCredentialAws**](CredentialsAWSApi.md#DeleteCredentialAws) | **Delete** /v1/Credentials/AWS/{Sid} | 
[**FetchCredentialAws**](CredentialsAWSApi.md#FetchCredentialAws) | **Get** /v1/Credentials/AWS/{Sid} | 
[**ListCredentialAws**](CredentialsAWSApi.md#ListCredentialAws) | **Get** /v1/Credentials/AWS | 
[**UpdateCredentialAws**](CredentialsAWSApi.md#UpdateCredentialAws) | **Post** /v1/Credentials/AWS/{Sid} | 



## CreateCredentialAws

> AccountsV1CredentialAws CreateCredentialAws(ctx, optional)



Create a new AWS Credential

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCredentialAwsParams struct


Name | Type | Description
------------- | ------------- | -------------
**Credentials** | **string** | A string that contains the AWS access credentials in the format `<AWS_ACCESS_KEY_ID>:<AWS_SECRET_ACCESS_KEY>`. For example, `AKIAIOSFODNN7EXAMPLE:wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY`
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**AccountSid** | **string** | The SID of the Subaccount that this Credential should be associated with. Must be a valid Subaccount of the account issuing the request.

### Return type

[**AccountsV1CredentialAws**](AccountsV1CredentialAws.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCredentialAws

> DeleteCredentialAws(ctx, Sid)



Delete a Credential from your account

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the AWS resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCredentialAwsParams struct


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


## FetchCredentialAws

> AccountsV1CredentialAws FetchCredentialAws(ctx, Sid)



Fetch the AWS credentials specified by the provided Credential Sid

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the AWS resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchCredentialAwsParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**AccountsV1CredentialAws**](AccountsV1CredentialAws.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredentialAws

> []AccountsV1CredentialAws ListCredentialAws(ctx, optional)



Retrieves a collection of AWS Credentials belonging to the account used to make the request

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCredentialAwsParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]AccountsV1CredentialAws**](AccountsV1CredentialAws.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredentialAws

> AccountsV1CredentialAws UpdateCredentialAws(ctx, Sidoptional)



Modify the properties of a given Account

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the AWS resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateCredentialAwsParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.

### Return type

[**AccountsV1CredentialAws**](AccountsV1CredentialAws.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

