# AssistantsRestoreApi

All URIs are relative to *https://autopilot.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateRestoreAssistant**](AssistantsRestoreApi.md#UpdateRestoreAssistant) | **Post** /v1/Assistants/Restore | 



## UpdateRestoreAssistant

> AutopilotV1RestoreAssistant UpdateRestoreAssistant(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRestoreAssistantParams struct


Name | Type | Description
------------- | ------------- | -------------
**Assistant** | **string** | The Twilio-provided string that uniquely identifies the Assistant resource to restore.

### Return type

[**AutopilotV1RestoreAssistant**](AutopilotV1RestoreAssistant.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

