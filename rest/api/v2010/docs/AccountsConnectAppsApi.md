# AccountsConnectAppsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConnectApp**](AccountsConnectAppsApi.md#DeleteConnectApp) | **Delete** /2010-04-01/Accounts/{AccountSid}/ConnectApps/{Sid}.json | 
[**FetchConnectApp**](AccountsConnectAppsApi.md#FetchConnectApp) | **Get** /2010-04-01/Accounts/{AccountSid}/ConnectApps/{Sid}.json | 
[**ListConnectApp**](AccountsConnectAppsApi.md#ListConnectApp) | **Get** /2010-04-01/Accounts/{AccountSid}/ConnectApps.json | 
[**UpdateConnectApp**](AccountsConnectAppsApi.md#UpdateConnectApp) | **Post** /2010-04-01/Accounts/{AccountSid}/ConnectApps/{Sid}.json | 



## DeleteConnectApp

> DeleteConnectApp(ctx, Sidoptional)



Delete an instance of a connect-app

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ConnectApp resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConnectAppParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resource to fetch.

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


## FetchConnectApp

> ApiV2010ConnectApp FetchConnectApp(ctx, Sidoptional)



Fetch an instance of a connect-app

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ConnectApp resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchConnectAppParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resource to fetch.

### Return type

[**ApiV2010ConnectApp**](ApiV2010ConnectApp.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectApp

> []ApiV2010ConnectApp ListConnectApp(ctx, optional)



Retrieve a list of connect-apps belonging to the account used to make the request

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListConnectAppParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010ConnectApp**](ApiV2010ConnectApp.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectApp

> ApiV2010ConnectApp UpdateConnectApp(ctx, Sidoptional)



Update a connect-app with the specified parameters

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the ConnectApp resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConnectAppParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resources to update.
**AuthorizeRedirectUrl** | **string** | The URL to redirect the user to after we authenticate the user and obtain authorization to access the Connect App.
**CompanyName** | **string** | The company name to set for the Connect App.
**DeauthorizeCallbackMethod** | **string** | The HTTP method to use when calling `deauthorize_callback_url`.
**DeauthorizeCallbackUrl** | **string** | The URL to call using the `deauthorize_callback_method` to de-authorize the Connect App.
**Description** | **string** | A description of the Connect App.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**HomepageUrl** | **string** | A public URL where users can obtain more information about this Connect App.
**Permissions** | [**[]ConnectAppEnumPermission**](ConnectAppEnumPermission.md) | A comma-separated list of the permissions you will request from the users of this ConnectApp.  Can include: `get-all` and `post-all`.

### Return type

[**ApiV2010ConnectApp**](ApiV2010ConnectApp.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

