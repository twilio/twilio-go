# AccountsAuthorizedConnectAppsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchAuthorizedConnectApp**](AccountsAuthorizedConnectAppsApi.md#FetchAuthorizedConnectApp) | **Get** /2010-04-01/Accounts/{AccountSid}/AuthorizedConnectApps/{ConnectAppSid}.json | 
[**ListAuthorizedConnectApp**](AccountsAuthorizedConnectAppsApi.md#ListAuthorizedConnectApp) | **Get** /2010-04-01/Accounts/{AccountSid}/AuthorizedConnectApps.json | 



## FetchAuthorizedConnectApp

> ApiV2010AuthorizedConnectApp FetchAuthorizedConnectApp(ctx, ConnectAppSidoptional)



Fetch an instance of an authorized-connect-app

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConnectAppSid** | **string** | The SID of the Connect App to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchAuthorizedConnectAppParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resource to fetch.

### Return type

[**ApiV2010AuthorizedConnectApp**](ApiV2010AuthorizedConnectApp.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthorizedConnectApp

> []ApiV2010AuthorizedConnectApp ListAuthorizedConnectApp(ctx, optional)



Retrieve a list of authorized-connect-apps belonging to the account used to make the request

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListAuthorizedConnectAppParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the AuthorizedConnectApp resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010AuthorizedConnectApp**](ApiV2010AuthorizedConnectApp.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

