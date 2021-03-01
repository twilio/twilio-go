# DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateByocTrunk**](DefaultApi.md#CreateByocTrunk) | **Post** /v1/ByocTrunks | 
[**CreateConnectionPolicy**](DefaultApi.md#CreateConnectionPolicy) | **Post** /v1/ConnectionPolicies | 
[**CreateConnectionPolicyTarget**](DefaultApi.md#CreateConnectionPolicyTarget) | **Post** /v1/ConnectionPolicies/{ConnectionPolicySid}/Targets | 
[**CreateDialingPermissionsCountryBulkUpdate**](DefaultApi.md#CreateDialingPermissionsCountryBulkUpdate) | **Post** /v1/DialingPermissions/BulkCountryUpdates | 
[**CreateIpRecord**](DefaultApi.md#CreateIpRecord) | **Post** /v1/IpRecords | 
[**CreateSourceIpMapping**](DefaultApi.md#CreateSourceIpMapping) | **Post** /v1/SourceIpMappings | 
[**DeleteByocTrunk**](DefaultApi.md#DeleteByocTrunk) | **Delete** /v1/ByocTrunks/{Sid} | 
[**DeleteConnectionPolicy**](DefaultApi.md#DeleteConnectionPolicy) | **Delete** /v1/ConnectionPolicies/{Sid} | 
[**DeleteConnectionPolicyTarget**](DefaultApi.md#DeleteConnectionPolicyTarget) | **Delete** /v1/ConnectionPolicies/{ConnectionPolicySid}/Targets/{Sid} | 
[**DeleteIpRecord**](DefaultApi.md#DeleteIpRecord) | **Delete** /v1/IpRecords/{Sid} | 
[**DeleteSourceIpMapping**](DefaultApi.md#DeleteSourceIpMapping) | **Delete** /v1/SourceIpMappings/{Sid} | 
[**FetchByocTrunk**](DefaultApi.md#FetchByocTrunk) | **Get** /v1/ByocTrunks/{Sid} | 
[**FetchConnectionPolicy**](DefaultApi.md#FetchConnectionPolicy) | **Get** /v1/ConnectionPolicies/{Sid} | 
[**FetchConnectionPolicyTarget**](DefaultApi.md#FetchConnectionPolicyTarget) | **Get** /v1/ConnectionPolicies/{ConnectionPolicySid}/Targets/{Sid} | 
[**FetchDialingPermissionsCountry**](DefaultApi.md#FetchDialingPermissionsCountry) | **Get** /v1/DialingPermissions/Countries/{IsoCode} | 
[**FetchDialingPermissionsSettings**](DefaultApi.md#FetchDialingPermissionsSettings) | **Get** /v1/Settings | 
[**FetchIpRecord**](DefaultApi.md#FetchIpRecord) | **Get** /v1/IpRecords/{Sid} | 
[**FetchSourceIpMapping**](DefaultApi.md#FetchSourceIpMapping) | **Get** /v1/SourceIpMappings/{Sid} | 
[**ListByocTrunk**](DefaultApi.md#ListByocTrunk) | **Get** /v1/ByocTrunks | 
[**ListConnectionPolicy**](DefaultApi.md#ListConnectionPolicy) | **Get** /v1/ConnectionPolicies | 
[**ListConnectionPolicyTarget**](DefaultApi.md#ListConnectionPolicyTarget) | **Get** /v1/ConnectionPolicies/{ConnectionPolicySid}/Targets | 
[**ListDialingPermissionsCountry**](DefaultApi.md#ListDialingPermissionsCountry) | **Get** /v1/DialingPermissions/Countries | 
[**ListDialingPermissionsHrsPrefixes**](DefaultApi.md#ListDialingPermissionsHrsPrefixes) | **Get** /v1/DialingPermissions/Countries/{IsoCode}/HighRiskSpecialPrefixes | 
[**ListIpRecord**](DefaultApi.md#ListIpRecord) | **Get** /v1/IpRecords | 
[**ListSourceIpMapping**](DefaultApi.md#ListSourceIpMapping) | **Get** /v1/SourceIpMappings | 
[**UpdateByocTrunk**](DefaultApi.md#UpdateByocTrunk) | **Post** /v1/ByocTrunks/{Sid} | 
[**UpdateConnectionPolicy**](DefaultApi.md#UpdateConnectionPolicy) | **Post** /v1/ConnectionPolicies/{Sid} | 
[**UpdateConnectionPolicyTarget**](DefaultApi.md#UpdateConnectionPolicyTarget) | **Post** /v1/ConnectionPolicies/{ConnectionPolicySid}/Targets/{Sid} | 
[**UpdateDialingPermissionsSettings**](DefaultApi.md#UpdateDialingPermissionsSettings) | **Post** /v1/Settings | 
[**UpdateIpRecord**](DefaultApi.md#UpdateIpRecord) | **Post** /v1/IpRecords/{Sid} | 
[**UpdateSourceIpMapping**](DefaultApi.md#UpdateSourceIpMapping) | **Post** /v1/SourceIpMappings/{Sid} | 



## CreateByocTrunk

> VoiceV1ByocTrunk CreateByocTrunk(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateByocTrunkRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateByocTrunkRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CnamLookupEnabled** | **Bool**| Whether Caller ID Name (CNAM) lookup is enabled for the trunk. If enabled, all inbound calls to the BYOC Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information. | 
**ConnectionPolicySid** | **String**| The SID of the Connection Policy that Twilio will use when routing traffic to your communications infrastructure. | 
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 
**FromDomainSid** | **String**| The SID of the SIP Domain that should be used in the &#x60;From&#x60; header of originating calls sent to your SIP infrastructure. If your SIP infrastructure allows users to \\\&quot;call back\\\&quot; an incoming call, configure this with a [SIP Domain](https://www.twilio.com/docs/voice/api/sending-sip) to ensure proper routing. If not configured, the from domain will default to \\\&quot;sip.twilio.com\\\&quot;. | 
**StatusCallbackMethod** | **String**| The HTTP method we should use to call &#x60;status_callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
**StatusCallbackUrl** | **String**| The URL that we should call to pass status parameters (such as call ended) to your application. | 
**VoiceFallbackMethod** | **String**| The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
**VoiceFallbackUrl** | **String**| The URL that we should call when an error occurs while retrieving or executing the TwiML from &#x60;voice_url&#x60;. | 
**VoiceMethod** | **String**| The HTTP method we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
**VoiceUrl** | **String**| The URL we should call when the BYOC Trunk receives a call. | 

### Return type

[**VoiceV1ByocTrunk**](voice.v1.byoc_trunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConnectionPolicy

> VoiceV1ConnectionPolicy CreateConnectionPolicy(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateConnectionPolicyRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateConnectionPolicyRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 

### Return type

[**VoiceV1ConnectionPolicy**](voice.v1.connection_policy.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConnectionPolicyTarget

> VoiceV1ConnectionPolicyConnectionPolicyTarget CreateConnectionPolicyTarget(ctx, ConnectionPolicySid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConnectionPolicySid** | **string**| The SID of the Connection Policy that owns the Target. | 
 **optional** | ***CreateConnectionPolicyTargetRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateConnectionPolicyTargetRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Enabled** | **Bool**| Whether the Target is enabled. The default is &#x60;true&#x60;. | 
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 
**Priority** | **Int32**| The relative importance of the target. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important target. | 
**Target** | **String**| The SIP address you want Twilio to route your calls to. This must be a &#x60;sip:&#x60; schema. &#x60;sips&#x60; is NOT supported. | 
**Weight** | **Int32**| The value that determines the relative share of the load the Target should receive compared to other Targets with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. Targets with higher values receive more load than those with lower ones with the same priority. | 

### Return type

[**VoiceV1ConnectionPolicyConnectionPolicyTarget**](voice.v1.connection_policy.connection_policy_target.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDialingPermissionsCountryBulkUpdate

> VoiceV1DialingPermissionsDialingPermissionsCountryBulkUpdate CreateDialingPermissionsCountryBulkUpdate(ctx, optional)



Create a bulk update request to change voice dialing country permissions of one or more countries identified by the corresponding [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateDialingPermissionsCountryBulkUpdateRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDialingPermissionsCountryBulkUpdateRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**UpdateRequest** | **String**| URL encoded JSON array of update objects. example : &#x60;[ { \\\&quot;iso_code\\\&quot;: \\\&quot;GB\\\&quot;, \\\&quot;low_risk_numbers_enabled\\\&quot;: \\\&quot;true\\\&quot;, \\\&quot;high_risk_special_numbers_enabled\\\&quot;:\\\&quot;true\\\&quot;, \\\&quot;high_risk_tollfraud_numbers_enabled\\\&quot;: \\\&quot;false\\\&quot; } ]&#x60; | 

### Return type

[**VoiceV1DialingPermissionsDialingPermissionsCountryBulkUpdate**](voice.v1.dialing_permissions.dialing_permissions_country_bulk_update.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIpRecord

> VoiceV1IpRecord CreateIpRecord(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateIpRecordRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateIpRecordRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CidrPrefixLength** | **Int32**| An integer representing the length of the [CIDR](https://tools.ietf.org/html/rfc4632) prefix to use with this IP address. By default the entire IP address is used, which for IPv4 is value 32. | 
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 
**IpAddress** | **String**| An IP address in dotted decimal notation, IPv4 only. | 

### Return type

[**VoiceV1IpRecord**](voice.v1.ip_record.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSourceIpMapping

> VoiceV1SourceIpMapping CreateSourceIpMapping(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateSourceIpMappingRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSourceIpMappingRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**IpRecordSid** | **String**| The Twilio-provided string that uniquely identifies the IP Record resource to map from. | 
**SipDomainSid** | **String**| The SID of the SIP Domain that the IP Record should be mapped to. | 

### Return type

[**VoiceV1SourceIpMapping**](voice.v1.source_ip_mapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteByocTrunk

> DeleteByocTrunk(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the BYOC Trunk resource to delete. | 

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


## DeleteConnectionPolicy

> DeleteConnectionPolicy(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique string that we created to identify the Connection Policy resource to delete. | 

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


## DeleteConnectionPolicyTarget

> DeleteConnectionPolicyTarget(ctx, ConnectionPolicySid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConnectionPolicySid** | **string**| The SID of the Connection Policy that owns the Target. | 
**Sid** | **string**| The unique string that we created to identify the Target resource to delete. | 

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


## DeleteIpRecord

> DeleteIpRecord(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the IP Record resource to delete. | 

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


## DeleteSourceIpMapping

> DeleteSourceIpMapping(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the IP Record resource to delete. | 

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


## FetchByocTrunk

> VoiceV1ByocTrunk FetchByocTrunk(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the BYOC Trunk resource to fetch. | 

### Return type

[**VoiceV1ByocTrunk**](voice.v1.byoc_trunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConnectionPolicy

> VoiceV1ConnectionPolicy FetchConnectionPolicy(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique string that we created to identify the Connection Policy resource to fetch. | 

### Return type

[**VoiceV1ConnectionPolicy**](voice.v1.connection_policy.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchConnectionPolicyTarget

> VoiceV1ConnectionPolicyConnectionPolicyTarget FetchConnectionPolicyTarget(ctx, ConnectionPolicySid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConnectionPolicySid** | **string**| The SID of the Connection Policy that owns the Target. | 
**Sid** | **string**| The unique string that we created to identify the Target resource to fetch. | 

### Return type

[**VoiceV1ConnectionPolicyConnectionPolicyTarget**](voice.v1.connection_policy.connection_policy_target.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDialingPermissionsCountry

> VoiceV1DialingPermissionsDialingPermissionsCountryInstance FetchDialingPermissionsCountry(ctx, IsoCode)



Retrieve voice dialing country permissions identified by the given ISO country code

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IsoCode** | **string**| The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the DialingPermissions Country resource to fetch | 

### Return type

[**VoiceV1DialingPermissionsDialingPermissionsCountryInstance**](voice.v1.dialing_permissions.dialing_permissions_country-instance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchDialingPermissionsSettings

> VoiceV1DialingPermissionsDialingPermissionsSettings FetchDialingPermissionsSettings(ctx, )



Retrieve voice dialing permissions inheritance for the sub-account

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**VoiceV1DialingPermissionsDialingPermissionsSettings**](voice.v1.dialing_permissions.dialing_permissions_settings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIpRecord

> VoiceV1IpRecord FetchIpRecord(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the IP Record resource to fetch. | 

### Return type

[**VoiceV1IpRecord**](voice.v1.ip_record.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSourceIpMapping

> VoiceV1SourceIpMapping FetchSourceIpMapping(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the IP Record resource to fetch. | 

### Return type

[**VoiceV1SourceIpMapping**](voice.v1.source_ip_mapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListByocTrunk

> ListByocTrunkResponse ListByocTrunk(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListByocTrunkRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListByocTrunkRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListByocTrunkResponse**](ListByocTrunkResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectionPolicy

> ListConnectionPolicyResponse ListConnectionPolicy(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListConnectionPolicyRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListConnectionPolicyRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListConnectionPolicyResponse**](ListConnectionPolicyResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectionPolicyTarget

> ListConnectionPolicyTargetResponse ListConnectionPolicyTarget(ctx, ConnectionPolicySid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConnectionPolicySid** | **string**| The SID of the Connection Policy from which to read the Targets. | 
 **optional** | ***ListConnectionPolicyTargetRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListConnectionPolicyTargetRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListConnectionPolicyTargetResponse**](ListConnectionPolicyTargetResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDialingPermissionsCountry

> ListDialingPermissionsCountryResponse ListDialingPermissionsCountry(ctx, optional)



Retrieve all voice dialing country permissions for this account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListDialingPermissionsCountryRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDialingPermissionsCountryRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**IsoCode** | **String**| Filter to retrieve the country permissions by specifying the [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) | 
**Continent** | **String**| Filter to retrieve the country permissions by specifying the continent | 
**CountryCode** | **String**| Filter the results by specified [country codes](https://www.itu.int/itudoc/itu-t/ob-lists/icc/e164_763.html) | 
**LowRiskNumbersEnabled** | **Bool**| Filter to retrieve the country permissions with dialing to low-risk numbers enabled. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
**HighRiskSpecialNumbersEnabled** | **Bool**| Filter to retrieve the country permissions with dialing to high-risk special service numbers enabled. Can be: &#x60;true&#x60; or &#x60;false&#x60; | 
**HighRiskTollfraudNumbersEnabled** | **Bool**| Filter to retrieve the country permissions with dialing to high-risk [toll fraud](https://www.twilio.com/learn/voice-and-video/toll-fraud) numbers enabled. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListDialingPermissionsCountryResponse**](ListDialingPermissionsCountryResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDialingPermissionsHrsPrefixes

> ListDialingPermissionsHrsPrefixesResponse ListDialingPermissionsHrsPrefixes(ctx, IsoCode, optional)



Fetch the high-risk special services prefixes from the country resource corresponding to the [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**IsoCode** | **string**| The [ISO 3166-1 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) to identify the country permissions from which high-risk special service number prefixes are fetched | 
 **optional** | ***ListDialingPermissionsHrsPrefixesRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDialingPermissionsHrsPrefixesRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListDialingPermissionsHrsPrefixesResponse**](ListDialingPermissionsHrsPrefixesResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIpRecord

> ListIpRecordResponse ListIpRecord(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListIpRecordRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListIpRecordRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListIpRecordResponse**](ListIpRecordResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSourceIpMapping

> ListSourceIpMappingResponse ListSourceIpMapping(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSourceIpMappingRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSourceIpMappingRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSourceIpMappingResponse**](ListSourceIpMappingResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateByocTrunk

> VoiceV1ByocTrunk UpdateByocTrunk(ctx, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the BYOC Trunk resource to update. | 
 **optional** | ***UpdateByocTrunkRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateByocTrunkRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CnamLookupEnabled** | **Bool**| Whether Caller ID Name (CNAM) lookup is enabled for the trunk. If enabled, all inbound calls to the BYOC Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information. | 
**ConnectionPolicySid** | **String**| The SID of the Connection Policy that Twilio will use when routing traffic to your communications infrastructure. | 
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 
**FromDomainSid** | **String**| The SID of the SIP Domain that should be used in the &#x60;From&#x60; header of originating calls sent to your SIP infrastructure. If your SIP infrastructure allows users to \\\&quot;call back\\\&quot; an incoming call, configure this with a [SIP Domain](https://www.twilio.com/docs/voice/api/sending-sip) to ensure proper routing. If not configured, the from domain will default to \\\&quot;sip.twilio.com\\\&quot;. | 
**StatusCallbackMethod** | **String**| The HTTP method we should use to call &#x60;status_callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
**StatusCallbackUrl** | **String**| The URL that we should call to pass status parameters (such as call ended) to your application. | 
**VoiceFallbackMethod** | **String**| The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
**VoiceFallbackUrl** | **String**| The URL that we should call when an error occurs while retrieving or executing the TwiML requested by &#x60;voice_url&#x60;. | 
**VoiceMethod** | **String**| The HTTP method we should use to call &#x60;voice_url&#x60; | 
**VoiceUrl** | **String**| The URL we should call when the BYOC Trunk receives a call. | 

### Return type

[**VoiceV1ByocTrunk**](voice.v1.byoc_trunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectionPolicy

> VoiceV1ConnectionPolicy UpdateConnectionPolicy(ctx, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique string that we created to identify the Connection Policy resource to update. | 
 **optional** | ***UpdateConnectionPolicyRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateConnectionPolicyRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 

### Return type

[**VoiceV1ConnectionPolicy**](voice.v1.connection_policy.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectionPolicyTarget

> VoiceV1ConnectionPolicyConnectionPolicyTarget UpdateConnectionPolicyTarget(ctx, ConnectionPolicySid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ConnectionPolicySid** | **string**| The SID of the Connection Policy that owns the Target. | 
**Sid** | **string**| The unique string that we created to identify the Target resource to update. | 
 **optional** | ***UpdateConnectionPolicyTargetRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateConnectionPolicyTargetRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Enabled** | **Bool**| Whether the Target is enabled. | 
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 
**Priority** | **Int32**| The relative importance of the target. Can be an integer from 0 to 65535, inclusive. The lowest number represents the most important target. | 
**Target** | **String**| The SIP address you want Twilio to route your calls to. This must be a &#x60;sip:&#x60; schema. &#x60;sips&#x60; is NOT supported. | 
**Weight** | **Int32**| The value that determines the relative share of the load the Target should receive compared to other Targets with the same priority. Can be an integer from 1 to 65535, inclusive. Targets with higher values receive more load than those with lower ones with the same priority. | 

### Return type

[**VoiceV1ConnectionPolicyConnectionPolicyTarget**](voice.v1.connection_policy.connection_policy_target.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDialingPermissionsSettings

> VoiceV1DialingPermissionsDialingPermissionsSettings UpdateDialingPermissionsSettings(ctx, optional)



Update voice dialing permissions inheritance for the sub-account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateDialingPermissionsSettingsRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDialingPermissionsSettingsRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**DialingPermissionsInheritance** | **Bool**| &#x60;true&#x60; for the sub-account to inherit voice dialing permissions from the Master Project; otherwise &#x60;false&#x60;. | 

### Return type

[**VoiceV1DialingPermissionsDialingPermissionsSettings**](voice.v1.dialing_permissions.dialing_permissions_settings.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIpRecord

> VoiceV1IpRecord UpdateIpRecord(ctx, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the IP Record resource to update. | 
 **optional** | ***UpdateIpRecordRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateIpRecordRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 

### Return type

[**VoiceV1IpRecord**](voice.v1.ip_record.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSourceIpMapping

> VoiceV1SourceIpMapping UpdateSourceIpMapping(ctx, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The Twilio-provided string that uniquely identifies the IP Record resource to update. | 
 **optional** | ***UpdateSourceIpMappingRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSourceIpMappingRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**SipDomainSid** | **String**| The SID of the SIP Domain that the IP Record should be mapped to. | 

### Return type

[**VoiceV1SourceIpMapping**](voice.v1.source_ip_mapping.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

