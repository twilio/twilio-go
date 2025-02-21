# HostedNumberOrdersBulkApi

All URIs are relative to *https://numbers.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchBulkHostedNumberOrder**](HostedNumberOrdersBulkApi.md#FetchBulkHostedNumberOrder) | **Get** /v2/HostedNumber/Orders/Bulk/{BulkHostingSid} | 



## FetchBulkHostedNumberOrder

> NumbersV2BulkHostedNumberOrder FetchBulkHostedNumberOrder(ctx, BulkHostingSidoptional)



Fetch a specific BulkHostedNumberOrder.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**BulkHostingSid** | **string** | A 34 character string that uniquely identifies this BulkHostedNumberOrder.

### Other Parameters

Other parameters are passed through a pointer to a FetchBulkHostedNumberOrderParams struct


Name | Type | Description
------------- | ------------- | -------------
**OrderStatus** | **string** | Order status can be used for filtering on Hosted Number Order status values. To see a complete list of order statuses, please check 'https://www.twilio.com/docs/phone-numbers/hosted-numbers/hosted-numbers-api/hosted-number-order-resource#status-values'.

### Return type

[**NumbersV2BulkHostedNumberOrder**](NumbersV2BulkHostedNumberOrder.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

