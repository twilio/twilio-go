# ServicesUsersSegmentMembershipsApi

All URIs are relative to *https://notify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSegmentMembership**](ServicesUsersSegmentMembershipsApi.md#CreateSegmentMembership) | **Post** /v1/Services/{ServiceSid}/Users/{Identity}/SegmentMemberships | 
[**DeleteSegmentMembership**](ServicesUsersSegmentMembershipsApi.md#DeleteSegmentMembership) | **Delete** /v1/Services/{ServiceSid}/Users/{Identity}/SegmentMemberships/{Segment} | 
[**FetchSegmentMembership**](ServicesUsersSegmentMembershipsApi.md#FetchSegmentMembership) | **Get** /v1/Services/{ServiceSid}/Users/{Identity}/SegmentMemberships/{Segment} | 



## CreateSegmentMembership

> NotifyV1SegmentMembership CreateSegmentMembership(ctx, ServiceSidIdentityoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Identity** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a CreateSegmentMembershipParams struct


Name | Type | Description
------------- | ------------- | -------------
**Segment** | **string** | 

### Return type

[**NotifyV1SegmentMembership**](NotifyV1SegmentMembership.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSegmentMembership

> DeleteSegmentMembership(ctx, ServiceSidIdentitySegment)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Identity** | **string** | 
**Segment** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteSegmentMembershipParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSegmentMembership

> NotifyV1SegmentMembership FetchSegmentMembership(ctx, ServiceSidIdentitySegment)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | 
**Identity** | **string** | 
**Segment** | **string** | 

### Other Parameters

Other parameters are passed through a pointer to a FetchSegmentMembershipParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**NotifyV1SegmentMembership**](NotifyV1SegmentMembership.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

