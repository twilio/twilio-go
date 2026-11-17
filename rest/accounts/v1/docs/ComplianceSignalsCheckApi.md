# ComplianceSignalsCheckApi

All URIs are relative to *https://accounts.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComplianceSignalsCheck**](ComplianceSignalsCheckApi.md#CreateComplianceSignalsCheck) | **Post** /v1/ComplianceSignals/Check | Evaluate compliance signals for a message



## CreateComplianceSignalsCheck

> AccountsV1ComplianceSignalsCheck CreateComplianceSignalsCheck(ctx, optional)

Evaluate compliance signals for a message

Evaluates compliance signals for a message to a phone number. Only the `QUIET_HOURS` signal is supported for now. The request and the response are JSON. The errors returned by the service are RFC 9457 problem details; errors raised before the request reaches the service, such as authentication failures, use the standard Twilio error format.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateComplianceSignalsCheckParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreateComplianceSignalsCheckRequest** | [**CreateComplianceSignalsCheckRequest**](CreateComplianceSignalsCheckRequest.md) | 

### Return type

[**AccountsV1ComplianceSignalsCheck**](AccountsV1ComplianceSignalsCheck.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

