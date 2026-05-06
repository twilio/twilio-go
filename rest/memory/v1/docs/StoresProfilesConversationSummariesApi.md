# StoresProfilesConversationSummariesApi

All URIs are relative to *https://memory.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProfileConversationSummary**](StoresProfilesConversationSummariesApi.md#CreateProfileConversationSummary) | **Post** /v1/Stores/{storeId}/Profiles/{profileId}/ConversationSummaries | Create Conversation Summaries
[**DeleteProfileConversationSummary**](StoresProfilesConversationSummariesApi.md#DeleteProfileConversationSummary) | **Delete** /v1/Stores/{storeId}/Profiles/{profileId}/ConversationSummaries/{summaryId} | Delete Conversation Summary
[**FetchProfileConversationSummary**](StoresProfilesConversationSummariesApi.md#FetchProfileConversationSummary) | **Get** /v1/Stores/{storeId}/Profiles/{profileId}/ConversationSummaries/{summaryId} | Retrieve Conversation Summary
[**ListProfileConversationSummaries**](StoresProfilesConversationSummariesApi.md#ListProfileConversationSummaries) | **Get** /v1/Stores/{storeId}/Profiles/{profileId}/ConversationSummaries | List Conversation Summaries
[**PatchProfileConversationSummary**](StoresProfilesConversationSummariesApi.md#PatchProfileConversationSummary) | **Patch** /v1/Stores/{storeId}/Profiles/{profileId}/ConversationSummaries/{summaryId} | Update Conversation Summary



## CreateProfileConversationSummary

> SummariesCreatedResponse CreateProfileConversationSummary(ctx, StoreIdProfileIdoptional)

Create Conversation Summaries

Create one or more conversation summaries associated with the specified profile. Supports both single summary creation and batch creation of up to 10 summaries. Supports request compression for large batch operations and response compression for the response. All summaries will be automatically indexed for semantic search capabilities. The content summary can be up to 4KB in length. Each summary will be created with a unique ID in Twilio Type ID (TTID) format.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.

### Other Parameters

Other parameters are passed through a pointer to a CreateProfileConversationSummaryParams struct


Name | Type | Description
------------- | ------------- | -------------
**AcceptEncoding** | **string** | Compression algorithms supported by the client (e.g., gzip, deflate, br)
**ContentEncoding** | **string** | Compression algorithm used for the request body (e.g., gzip, deflate, br)
**CreateSummariesRequest** | [**CreateSummariesRequest**](CreateSummariesRequest.md) | 

### Return type

[**SummariesCreatedResponse**](SummariesCreatedResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProfileConversationSummary

> DeleteProfileConversationSummaryResponse DeleteProfileConversationSummary(ctx, StoreIdProfileIdSummaryId)

Delete Conversation Summary

Delete a specific conversation summary by its ID. This action is irreversible.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.
**SummaryId** | **string** | The summary ID.

### Other Parameters

Other parameters are passed through a pointer to a DeleteProfileConversationSummaryParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**DeleteProfileConversationSummaryResponse**](DeleteProfileConversationSummary202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchProfileConversationSummary

> SummaryInfo FetchProfileConversationSummary(ctx, StoreIdProfileIdSummaryId)

Retrieve Conversation Summary

Retrieve a specific conversation summary by its ID for the given profile.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.
**SummaryId** | **string** | The summary ID.

### Other Parameters

Other parameters are passed through a pointer to a FetchProfileConversationSummaryParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SummaryInfo**](SummaryInfo.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProfileConversationSummaries

> []SummaryInfo ListProfileConversationSummaries(ctx, StoreIdProfileIdoptional)

List Conversation Summaries

Retrieve a paginated list of conversation summaries for a specific profile.  Supports response compression for large datasets.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.

### Other Parameters

Other parameters are passed through a pointer to a ListProfileConversationSummariesParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | The maximum number of items to return per page, maximum of 1000.
**PageToken** | **string** | The token for the page of results to retrieve.
**AcceptEncoding** | **string** | Compression algorithms supported by the client (e.g., gzip, deflate, br)
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SummaryInfo**](SummaryInfo.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchProfileConversationSummary

> PatchProfileConversationSummaryResponse PatchProfileConversationSummary(ctx, StoreIdProfileIdSummaryIdoptional)

Update Conversation Summary

Partially update a specific conversation summary by its ID. Only provided fields will be updated. The updated timestamp will be automatically set. This allows for selective updates without needing to provide all fields.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | A unique Memory Store ID using Twilio Type ID (TTID) format
**ProfileId** | **string** | The unique identifier for the profile using Twilio Type ID (TTID) format.
**SummaryId** | **string** | The summary ID.

### Other Parameters

Other parameters are passed through a pointer to a PatchProfileConversationSummaryParams struct


Name | Type | Description
------------- | ------------- | -------------
**SummaryCorePatch** | [**SummaryCorePatch**](SummaryCorePatch.md) | 

### Return type

[**PatchProfileConversationSummaryResponse**](PatchProfileConversationSummary202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

