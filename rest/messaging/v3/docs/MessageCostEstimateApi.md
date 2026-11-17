# MessageCostEstimateApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchV3WhatsappMessageCostEstimate**](MessageCostEstimateApi.md#FetchV3WhatsappMessageCostEstimate) | **Get** /v3/MessageCostEstimate | Fetch a WhatsApp message cost estimate



## FetchV3WhatsappMessageCostEstimate

> MessageCostEstimateResponse FetchV3WhatsappMessageCostEstimate(ctx, optional)

Fetch a WhatsApp message cost estimate

Fetch the estimated cost and delivery ranges for WhatsApp Marketing Max Pricing messages sent to a single country, at multiple max price (bid) amounts. This applies to Marketing Max Pricing messages only. Estimates are generated from historical data over the requested lookback window and are informational only.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchV3WhatsappMessageCostEstimateParams struct


Name | Type | Description
------------- | ------------- | -------------
**Country** | **string** | The ISO-3166-1 alpha-2 country code to estimate for. Exactly one country is accepted.
**LookbackWindow** | **string** | The historical lookback window used to compute the estimate.

### Return type

[**MessageCostEstimateResponse**](MessageCostEstimateResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

