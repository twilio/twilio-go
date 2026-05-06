# StoresProfilesEventsApi

All URIs are relative to *https://memory.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProfileEvents**](StoresProfilesEventsApi.md#CreateProfileEvents) | **Post** /v1/Stores/{storeId}/Profiles/{profileId}/Events | Add events to a specific profile.



## CreateProfileEvents

> CreateProfileEventsResponse CreateProfileEvents(ctx, StoreIdProfileIdoptional)

Add events to a specific profile.

Multiple events can be added to a profile in a single request. The maximum number of events per request is 1000 and they will be processed asynchronously.  Any traits that are configured for identifier promotion will be set as an identifier value on the profile.  Types of events:  - `twilio-communication-event` - Twilio communication event, used for tracking multi-channel communications.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**StoreId** | **string** | The storeId identifier
**ProfileId** | **string** | The profileId identifier

### Other Parameters

Other parameters are passed through a pointer to a CreateProfileEventsParams struct


Name | Type | Description
------------- | ------------- | -------------
**ProfileEventRequest** | [**ProfileEventRequest**](ProfileEventRequest.md) | 

### Return type

[**CreateProfileEventsResponse**](CreateProfileEvents202Response.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

