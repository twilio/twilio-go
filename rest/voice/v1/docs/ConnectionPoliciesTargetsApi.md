# ConnectionPoliciesTargetsApi

All URIs are relative to *https://voice.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectionPolicyTarget**](ConnectionPoliciesTargetsApi.md#CreateConnectionPolicyTarget) | **Post** /v1/ConnectionPolicies/{ConnectionPolicySid}/Targets | 
[**DeleteConnectionPolicyTarget**](ConnectionPoliciesTargetsApi.md#DeleteConnectionPolicyTarget) | **Delete** /v1/ConnectionPolicies/{ConnectionPolicySid}/Targets/{Sid} | 
[**FetchConnectionPolicyTarget**](ConnectionPoliciesTargetsApi.md#FetchConnectionPolicyTarget) | **Get** /v1/ConnectionPolicies/{ConnectionPolicySid}/Targets/{Sid} | 
[**ListConnectionPolicyTarget**](ConnectionPoliciesTargetsApi.md#ListConnectionPolicyTarget) | **Get** /v1/ConnectionPolicies/{ConnectionPolicySid}/Targets | 
[**UpdateConnectionPolicyTarget**](ConnectionPoliciesTargetsApi.md#UpdateConnectionPolicyTarget) | **Post** /v1/ConnectionPolicies/{ConnectionPolicySid}/Targets/{Sid} | 



## CreateConnectionPolicyTarget

> VoiceV1ConnectionPolicyTarget CreateConnectionPolicyTarget(ctx, ConnectionPolicySidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConnectionPolicySid** | **string** | The SID of the Connection Policy that owns the Target.

### Other Parameters

Other parameters are passed through a pointer to a CreateConnectionPolicyTargetParams struct


Name | Type | Description
------------- | ------------- | -------------
**Target** | **string** | The SIP address you want Twilio to route your calls to. This must be a `sip:` schema. `sips` is NOT supported.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
**Priority** | **int** | The relative importance of the target. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important target.
**Weight** | **int** | The value that determines the relative share of the load the Target should receive compared to other Targets with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. Targets with higher values receive more load than those with lower ones with the same priority.
**Enabled** | **bool** | Whether the Target is enabled. The default is `true`.

### Return type

[**VoiceV1ConnectionPolicyTarget**](VoiceV1ConnectionPolicyTarget.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectionPolicyTarget

> DeleteConnectionPolicyTarget(ctx, ConnectionPolicySidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConnectionPolicySid** | **string** | The SID of the Connection Policy that owns the Target.
**Sid** | **string** | The unique string that we created to identify the Target resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteConnectionPolicyTargetParams struct


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


## FetchConnectionPolicyTarget

> VoiceV1ConnectionPolicyTarget FetchConnectionPolicyTarget(ctx, ConnectionPolicySidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConnectionPolicySid** | **string** | The SID of the Connection Policy that owns the Target.
**Sid** | **string** | The unique string that we created to identify the Target resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchConnectionPolicyTargetParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**VoiceV1ConnectionPolicyTarget**](VoiceV1ConnectionPolicyTarget.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectionPolicyTarget

> []VoiceV1ConnectionPolicyTarget ListConnectionPolicyTarget(ctx, ConnectionPolicySidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConnectionPolicySid** | **string** | The SID of the Connection Policy from which to read the Targets.

### Other Parameters

Other parameters are passed through a pointer to a ListConnectionPolicyTargetParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]VoiceV1ConnectionPolicyTarget**](VoiceV1ConnectionPolicyTarget.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectionPolicyTarget

> VoiceV1ConnectionPolicyTarget UpdateConnectionPolicyTarget(ctx, ConnectionPolicySidSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConnectionPolicySid** | **string** | The SID of the Connection Policy that owns the Target.
**Sid** | **string** | The unique string that we created to identify the Target resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateConnectionPolicyTargetParams struct


Name | Type | Description
------------- | ------------- | -------------
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
**Target** | **string** | The SIP address you want Twilio to route your calls to. This must be a `sip:` schema. `sips` is NOT supported.
**Priority** | **int** | The relative importance of the target. Can be an integer from 0 to 65535, inclusive. The lowest number represents the most important target.
**Weight** | **int** | The value that determines the relative share of the load the Target should receive compared to other Targets with the same priority. Can be an integer from 1 to 65535, inclusive. Targets with higher values receive more load than those with lower ones with the same priority.
**Enabled** | **bool** | Whether the Target is enabled.

### Return type

[**VoiceV1ConnectionPolicyTarget**](VoiceV1ConnectionPolicyTarget.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

