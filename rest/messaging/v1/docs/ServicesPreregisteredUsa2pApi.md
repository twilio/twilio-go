# ServicesPreregisteredUsa2pApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalCampaign**](ServicesPreregisteredUsa2pApi.md#CreateExternalCampaign) | **Post** /v1/Services/PreregisteredUsa2p | 



## CreateExternalCampaign

> MessagingV1ExternalCampaign CreateExternalCampaign(ctx, optional)





### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateExternalCampaignParams struct


Name | Type | Description
------------- | ------------- | -------------
**CampaignId** | **string** | ID of the preregistered campaign.
**MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) that the resource is associated with.

### Return type

[**MessagingV1ExternalCampaign**](MessagingV1ExternalCampaign.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

