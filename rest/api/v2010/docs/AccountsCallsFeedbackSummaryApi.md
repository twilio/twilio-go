# AccountsCallsFeedbackSummaryApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCallFeedbackSummary**](AccountsCallsFeedbackSummaryApi.md#CreateCallFeedbackSummary) | **Post** /2010-04-01/Accounts/{AccountSid}/Calls/FeedbackSummary.json | 
[**DeleteCallFeedbackSummary**](AccountsCallsFeedbackSummaryApi.md#DeleteCallFeedbackSummary) | **Delete** /2010-04-01/Accounts/{AccountSid}/Calls/FeedbackSummary/{Sid}.json | 
[**FetchCallFeedbackSummary**](AccountsCallsFeedbackSummaryApi.md#FetchCallFeedbackSummary) | **Get** /2010-04-01/Accounts/{AccountSid}/Calls/FeedbackSummary/{Sid}.json | 



## CreateCallFeedbackSummary

> ApiV2010CallFeedbackSummary CreateCallFeedbackSummary(ctx, optional)



Create a FeedbackSummary resource for a call

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCallFeedbackSummaryParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
**EndDate** | **string** | Only include feedback given on or before this date. Format is &#x60;YYYY-MM-DD&#x60; and specified in UTC.
**IncludeSubaccounts** | **bool** | Whether to also include Feedback resources from all subaccounts. &#x60;true&#x60; includes feedback from all subaccounts and &#x60;false&#x60;, the default, includes feedback from only the specified account.
**StartDate** | **string** | Only include feedback given on or after this date. Format is &#x60;YYYY-MM-DD&#x60; and specified in UTC.
**StatusCallback** | **string** | The URL that we will request when the feedback summary is complete.
**StatusCallbackMethod** | **string** | The HTTP method (&#x60;GET&#x60; or &#x60;POST&#x60;) we use to make the request to the &#x60;StatusCallback&#x60; URL.

### Return type

[**ApiV2010CallFeedbackSummary**](ApiV2010CallFeedbackSummary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCallFeedbackSummary

> DeleteCallFeedbackSummary(ctx, Sidoptional)



Delete a FeedbackSummary resource from a call

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCallFeedbackSummaryParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.

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


## FetchCallFeedbackSummary

> ApiV2010CallFeedbackSummary FetchCallFeedbackSummary(ctx, Sidoptional)



Fetch a FeedbackSummary resource from a call

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | A 34 character string that uniquely identifies this resource.

### Other Parameters

Other parameters are passed through a pointer to a FetchCallFeedbackSummaryParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.

### Return type

[**ApiV2010CallFeedbackSummary**](ApiV2010CallFeedbackSummary.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

