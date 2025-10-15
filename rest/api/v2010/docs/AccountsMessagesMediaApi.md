# AccountsMessagesMediaApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMedia**](AccountsMessagesMediaApi.md#DeleteMedia) | **Delete** /2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Media/{Sid}.json | Delete the Media resource.
[**FetchMedia**](AccountsMessagesMediaApi.md#FetchMedia) | **Get** /2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Media/{Sid}.json | Fetch a single Media resource associated with a specific Message resource
[**ListMedia**](AccountsMessagesMediaApi.md#ListMedia) | **Get** /2010-04-01/Accounts/{AccountSid}/Messages/{MessageSid}/Media.json | Read a list of Media resources associated with a specific Message resource



## DeleteMedia

> DeleteMedia(ctx, MessageSidSidoptional)

Delete the Media resource.

Delete the Media resource.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessageSid** | **string** | The SID of the Message resource that is associated with the Media resource.
**Sid** | **string** | The unique identifier of the to-be-deleted Media resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteMediaParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is associated with the Media resource.

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


## FetchMedia

> ApiV2010Media FetchMedia(ctx, MessageSidSidoptional)

Fetch a single Media resource associated with a specific Message resource

Fetch a single Media resource associated with a specific Message resource

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessageSid** | **string** | The SID of the Message resource that is associated with the Media resource.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the Media resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchMediaParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) associated with the Media resource.

### Return type

[**ApiV2010Media**](ApiV2010Media.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMedia

> []ApiV2010Media ListMedia(ctx, MessageSidoptional)

Read a list of Media resources associated with a specific Message resource

Read a list of Media resources associated with a specific Message resource

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessageSid** | **string** | The SID of the Message resource that is associated with the Media resources.

### Other Parameters

Other parameters are passed through a pointer to a ListMediaParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that is associated with the Media resources.
**DateCreated** | **time.Time** | Only include Media resources that were created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read Media that were created on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read Media that were created on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read Media that were created on or after midnight of this date.
**DateCreatedBefore** | **time.Time** | Only include Media resources that were created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read Media that were created on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read Media that were created on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read Media that were created on or after midnight of this date.
**DateCreatedAfter** | **time.Time** | Only include Media resources that were created on this date. Specify a date as `YYYY-MM-DD` in GMT, for example: `2009-07-06`, to read Media that were created on this date. You can also specify an inequality, such as `StartTime<=YYYY-MM-DD`, to read Media that were created on or before midnight of this date, and `StartTime>=YYYY-MM-DD` to read Media that were created on or after midnight of this date.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]ApiV2010Media**](ApiV2010Media.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

