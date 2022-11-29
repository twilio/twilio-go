# AccountsSIPCredentialListsCredentialsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSipCredential**](AccountsSIPCredentialListsCredentialsApi.md#CreateSipCredential) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials.json | 
[**DeleteSipCredential**](AccountsSIPCredentialListsCredentialsApi.md#DeleteSipCredential) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials/{Sid}.json | 
[**FetchSipCredential**](AccountsSIPCredentialListsCredentialsApi.md#FetchSipCredential) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials/{Sid}.json | 
[**ListSipCredential**](AccountsSIPCredentialListsCredentialsApi.md#ListSipCredential) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials.json | 
[**UpdateSipCredential**](AccountsSIPCredentialListsCredentialsApi.md#UpdateSipCredential) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/CredentialLists/{CredentialListSid}/Credentials/{Sid}.json | 



## CreateSipCredential

> ApiV2010SipCredential CreateSipCredential(ctx, CredentialListSidoptional)



Create a new credential resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CredentialListSid** | **string** | The unique id that identifies the credential list to include the created credential.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipCredentialParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**Username** | **string** | The username that will be passed when authenticating SIP requests. The username should be sent in response to Twilio's challenge of the initial INVITE. It can be up to 32 characters long.
**Password** | **string** | The password that the username will use when authenticating SIP requests. The password must be a minimum of 12 characters, contain at least 1 digit, and have mixed case. (eg `IWasAtSignal2018`)

### Return type

[**ApiV2010SipCredential**](ApiV2010SipCredential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipCredential

> DeleteSipCredential(ctx, CredentialListSidSidoptional)



Delete a credential resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CredentialListSid** | **string** | The unique id that identifies the credential list that contains the desired credentials.
**Sid** | **string** | The unique id that identifies the resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipCredentialParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the Account that is responsible for this resource.

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


## FetchSipCredential

> ApiV2010SipCredential FetchSipCredential(ctx, CredentialListSidSidoptional)



Fetch a single credential.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CredentialListSid** | **string** | The unique id that identifies the credential list that contains the desired credential.
**Sid** | **string** | The unique id that identifies the resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSipCredentialParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the Account that is responsible for this resource.

### Return type

[**ApiV2010SipCredential**](ApiV2010SipCredential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipCredential

> []ApiV2010SipCredential ListSipCredential(ctx, CredentialListSidoptional)



Retrieve a list of credentials.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CredentialListSid** | **string** | The unique id that identifies the credential list that contains the desired credentials.

### Other Parameters

Other parameters are passed through a pointer to a ListSipCredentialParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010SipCredential**](ApiV2010SipCredential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipCredential

> ApiV2010SipCredential UpdateSipCredential(ctx, CredentialListSidSidoptional)



Update a credential resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**CredentialListSid** | **string** | The unique id that identifies the credential list that includes this credential.
**Sid** | **string** | The unique id that identifies the resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSipCredentialParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the Account that is responsible for this resource.
**Password** | **string** | The password that the username will use when authenticating SIP requests. The password must be a minimum of 12 characters, contain at least 1 digit, and have mixed case. (eg `IWasAtSignal2018`)

### Return type

[**ApiV2010SipCredential**](ApiV2010SipCredential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

