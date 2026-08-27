# RuleExecutionsApi

All URIs are relative to *https://intelligence.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRuleExecution**](RuleExecutionsApi.md#CreateRuleExecution) | **Post** /v3/RuleExecutions | Manually queue a rule execution



## CreateRuleExecution

> CreateRuleExecution(ctx, optional)

Manually queue a rule execution

Resolves the given configuration, rule, and conversation, derives the memoryStoreId from the conversation's configuration. Then executes the rule on the conversation. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateRuleExecutionParams struct


Name | Type | Description
------------- | ------------- | -------------
**CreateRuleExecutionRequest** | [**CreateRuleExecutionRequest**](CreateRuleExecutionRequest.md) | 

### Return type

 (empty response body)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

