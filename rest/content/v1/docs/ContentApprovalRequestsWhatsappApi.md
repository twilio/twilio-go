# ContentApprovalRequestsWhatsappApi

All URIs are relative to *https://content.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApprovalCreate**](ContentApprovalRequestsWhatsappApi.md#CreateApprovalCreate) | **Post** /v1/Content/{ContentSid}/ApprovalRequests/whatsapp | Create Content Approval Request
[**UpdateApprovalCreate**](ContentApprovalRequestsWhatsappApi.md#UpdateApprovalCreate) | **Patch** /v1/Content/{ContentSid}/ApprovalRequests/whatsapp | Update Content Approval Request



## CreateApprovalCreate

> ContentV1ApprovalCreate CreateApprovalCreate(ctx, ContentSidoptional)

Create Content Approval Request

Create a ContentApprovalRequest for a content item

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ContentSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateApprovalCreateParams struct


Name | Type | Description
------------- | ------------- | -------------
**ContentApprovalRequest** | [**ContentApprovalRequest**](ContentApprovalRequest.md) | 

### Return type

[**ContentV1ApprovalCreate**](ContentV1ApprovalCreate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApprovalCreate

> ContentV1ApprovalCreate UpdateApprovalCreate(ctx, ContentSidoptional)

Update Content Approval Request

Update the optimization spec (bid amount) of a ContentApprovalRequest for a content item

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ContentSid** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a PatchApprovalCreateParams struct


Name | Type | Description
------------- | ------------- | -------------
**ContentApprovalUpdateRequest** | [**ContentApprovalUpdateRequest**](ContentApprovalUpdateRequest.md) | 

### Return type

[**ContentV1ApprovalCreate**](ContentV1ApprovalCreate.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

