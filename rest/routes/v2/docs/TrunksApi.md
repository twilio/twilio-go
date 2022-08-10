# TrunksApi

All URIs are relative to *https://routes.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchTrunks**](TrunksApi.md#FetchTrunks) | **Get** /v2/Trunks/{SipTrunkDomain} | 
[**UpdateTrunks**](TrunksApi.md#UpdateTrunks) | **Post** /v2/Trunks/{SipTrunkDomain} | 



## FetchTrunks

> RoutesV2Trunks FetchTrunks(ctx, SipTrunkDomain)



Fetch the Inbound Processing Region assigned to a SIP Trunk.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SipTrunkDomain** | **string** | The absolute URL of the SIP Trunk

### Other Parameters

Other parameters are passed through a pointer to a FetchTrunksParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**RoutesV2Trunks**](RoutesV2Trunks.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrunks

> RoutesV2Trunks UpdateTrunks(ctx, SipTrunkDomainoptional)



Assign an Inbound Processing Region to a SIP Trunk

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**SipTrunkDomain** | **string** | The absolute URL of the SIP Trunk

### Other Parameters

Other parameters are passed through a pointer to a UpdateTrunksParams struct


Name | Type | Description
------------- | ------------- | -------------
**VoiceRegion** | **string** | The Inbound Processing Region used for this SIP Trunk for voice
**FriendlyName** | **string** | A human readable description of this resource, up to 64 characters.

### Return type

[**RoutesV2Trunks**](RoutesV2Trunks.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

