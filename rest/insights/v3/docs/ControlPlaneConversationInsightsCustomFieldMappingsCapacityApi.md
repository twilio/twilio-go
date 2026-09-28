# ControlPlaneConversationInsightsCustomFieldMappingsCapacityApi

All URIs are relative to *https://insights.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchCustomFieldMappingCapacity**](ControlPlaneConversationInsightsCustomFieldMappingsCapacityApi.md#FetchCustomFieldMappingCapacity) | **Get** /v3/ControlPlane/ConversationInsights/CustomFieldMappings/Capacity | Get custom field mapping capacity for an account



## FetchCustomFieldMappingCapacity

> CustomFieldCapacity FetchCustomFieldMappingCapacity(ctx, )

Get custom field mapping capacity for an account

Returns the account's custom measure slot capacity — slots consumed and remaining. Only the `measure` category is supported for now. 

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchCustomFieldMappingCapacityParams struct


### Return type

[**CustomFieldCapacity**](CustomFieldCapacity.md)

### Authorization

[basic_apikey_or_accountsid](../README.md#basic_apikey_or_accountsid)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

