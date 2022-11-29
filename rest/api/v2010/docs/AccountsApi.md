# AccountsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](AccountsApi.md#CreateAccount) | **Post** /2010-04-01/Accounts.json | 
[**FetchAccount**](AccountsApi.md#FetchAccount) | **Get** /2010-04-01/Accounts/{Sid}.json | 
[**ListAccount**](AccountsApi.md#ListAccount) | **Get** /2010-04-01/Accounts.json | 
[**UpdateAccount**](AccountsApi.md#UpdateAccount) | **Post** /2010-04-01/Accounts/{Sid}.json | 



## CreateAccount

> ApiV2010Account CreateAccount(ctx, optional)



Create a new Twilio Subaccount from the account making the request

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateAccountParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A human readable description of the account to create, defaults to `SubAccount Created at {YYYY-MM-DD HH:MM meridian}`

### Return type

[**ApiV2010Account**](ApiV2010Account.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAccount

> ApiV2010Account FetchAccount(ctx, Sid)



Fetch the account specified by the provided Account Sid

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Account Sid that uniquely identifies the account to fetch

### Other Parameters

Other parameters are passed through a pointer to a FetchAccountParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**ApiV2010Account**](ApiV2010Account.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccount

> []ApiV2010Account ListAccount(ctx, optional)



Retrieves a collection of Accounts belonging to the account used to make the request

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAccountParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | Only return the Account resources with friendly names that exactly match this name.
**Status** | **string** | Only return Account resources with the given status. Can be `closed`, `suspended` or `active`.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010Account**](ApiV2010Account.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> ApiV2010Account UpdateAccount(ctx, Sidoptional)



Modify the properties of a given Account

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Account Sid that uniquely identifies the account to update

### Other Parameters

Other parameters are passed through a pointer to a UpdateAccountParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | Update the human-readable description of this Account
**Status** | **string** | 

### Return type

[**ApiV2010Account**](ApiV2010Account.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

