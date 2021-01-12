# \DefaultApi

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
 **optional** | ***CreateByocTrunkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateByocTrunkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cnamLookupEnabled** | **optional.Bool**| Whether Caller ID Name (CNAM) lookup is enabled for the trunk. If enabled, all inbound calls to the BYOC Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information. | 
 **connectionPolicySid** | **optional.String**| The SID of the Connection Policy that Twilio will use when routing traffic to your communications infrastructure. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 
 **fromDomainSid** | **optional.String**| The SID of the SIP Domain that should be used in the &#x60;From&#x60; header of originating calls sent to your SIP infrastructure. If your SIP infrastructure allows users to \\\&quot;call back\\\&quot; an incoming call, configure this with a [SIP Domain](https://www.twilio.com/docs/voice/api/sending-sip) to ensure proper routing. If not configured, the from domain will default to \\\&quot;sip.twilio.com\\\&quot;. | 
 **statusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;status_callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **statusCallbackUrl** | **optional.String**| The URL that we should call to pass status parameters (such as call ended) to your application. | 
 **voiceFallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **voiceFallbackUrl** | **optional.String**| The URL that we should call when an error occurs while retrieving or executing the TwiML from &#x60;voice_url&#x60;. | 
 **voiceMethod** | **optional.String**| The HTTP method we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **voiceUrl** | **optional.String**| The URL we should call when the BYOC Trunk receives a call. | 

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
 **optional** | ***CreateConnectionPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateConnectionPolicyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 

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

> VoiceV1ConnectionPolicyConnectionPolicyTarget CreateConnectionPolicyTarget(ctx, connectionPolicySid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionPolicySid** | **string**| The SID of the Connection Policy that owns the Target. | 
 **optional** | ***CreateConnectionPolicyTargetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateConnectionPolicyTargetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enabled** | **optional.Bool**| Whether the Target is enabled. The default is &#x60;true&#x60;. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 
 **priority** | **optional.Int32**| The relative importance of the target. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important target. | 
 **target** | **optional.String**| The SIP address you want Twilio to route your calls to. This must be a &#x60;sip:&#x60; schema. &#x60;sips&#x60; is NOT supported. | 
 **weight** | **optional.Int32**| The value that determines the relative share of the load the Target should receive compared to other Targets with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. Targets with higher values receive more load than those with lower ones with the same priority. | 

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
 **optional** | ***CreateDialingPermissionsCountryBulkUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDialingPermissionsCountryBulkUpdateOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRequest** | **optional.String**| URL encoded JSON array of update objects. example : &#x60;[ { \\\&quot;iso_code\\\&quot;: \\\&quot;GB\\\&quot;, \\\&quot;low_risk_numbers_enabled\\\&quot;: \\\&quot;true\\\&quot;, \\\&quot;high_risk_special_numbers_enabled\\\&quot;:\\\&quot;true\\\&quot;, \\\&quot;high_risk_tollfraud_numbers_enabled\\\&quot;: \\\&quot;false\\\&quot; } ]&#x60; | 

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
 **optional** | ***CreateIpRecordOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateIpRecordOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cidrPrefixLength** | **optional.Int32**| An integer representing the length of the [CIDR](https://tools.ietf.org/html/rfc4632) prefix to use with this IP address. By default the entire IP address is used, which for IPv4 is value 32. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 
 **ipAddress** | **optional.String**| An IP address in dotted decimal notation, IPv4 only. | 

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
 **optional** | ***CreateSourceIpMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSourceIpMappingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ipRecordSid** | **optional.String**| The Twilio-provided string that uniquely identifies the IP Record resource to map from. | 
 **sipDomainSid** | **optional.String**| The SID of the SIP Domain that the IP Record should be mapped to. | 

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

> DeleteByocTrunk(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the BYOC Trunk resource to delete. | 

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

> DeleteConnectionPolicy(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The unique string that we created to identify the Connection Policy resource to delete. | 

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

> DeleteConnectionPolicyTarget(ctx, connectionPolicySid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionPolicySid** | **string**| The SID of the Connection Policy that owns the Target. | 
**sid** | **string**| The unique string that we created to identify the Target resource to delete. | 

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

> DeleteIpRecord(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the IP Record resource to delete. | 

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

> DeleteSourceIpMapping(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the IP Record resource to delete. | 

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

> VoiceV1ByocTrunk FetchByocTrunk(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the BYOC Trunk resource to fetch. | 

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

> VoiceV1ConnectionPolicy FetchConnectionPolicy(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The unique string that we created to identify the Connection Policy resource to fetch. | 

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

> VoiceV1ConnectionPolicyConnectionPolicyTarget FetchConnectionPolicyTarget(ctx, connectionPolicySid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionPolicySid** | **string**| The SID of the Connection Policy that owns the Target. | 
**sid** | **string**| The unique string that we created to identify the Target resource to fetch. | 

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

> VoiceV1DialingPermissionsDialingPermissionsCountryInstance FetchDialingPermissionsCountry(ctx, isoCode)



Retrieve voice dialing country permissions identified by the given ISO country code

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**isoCode** | **string**| The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the DialingPermissions Country resource to fetch | 

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

> VoiceV1IpRecord FetchIpRecord(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the IP Record resource to fetch. | 

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

> VoiceV1SourceIpMapping FetchSourceIpMapping(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the IP Record resource to fetch. | 

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

> InlineResponse200 ListByocTrunk(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListByocTrunkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListByocTrunkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectionPolicy

> InlineResponse2001 ListConnectionPolicy(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListConnectionPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListConnectionPolicyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectionPolicyTarget

> InlineResponse2002 ListConnectionPolicyTarget(ctx, connectionPolicySid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionPolicySid** | **string**| The SID of the Connection Policy from which to read the Targets. | 
 **optional** | ***ListConnectionPolicyTargetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListConnectionPolicyTargetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDialingPermissionsCountry

> InlineResponse2003 ListDialingPermissionsCountry(ctx, optional)



Retrieve all voice dialing country permissions for this account

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListDialingPermissionsCountryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDialingPermissionsCountryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isoCode** | **optional.String**| Filter to retrieve the country permissions by specifying the [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) | 
 **continent** | **optional.String**| Filter to retrieve the country permissions by specifying the continent | 
 **countryCode** | **optional.String**| Filter the results by specified [country codes](https://www.itu.int/itudoc/itu-t/ob-lists/icc/e164_763.html) | 
 **lowRiskNumbersEnabled** | **optional.Bool**| Filter to retrieve the country permissions with dialing to low-risk numbers enabled. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **highRiskSpecialNumbersEnabled** | **optional.Bool**| Filter to retrieve the country permissions with dialing to high-risk special service numbers enabled. Can be: &#x60;true&#x60; or &#x60;false&#x60; | 
 **highRiskTollfraudNumbersEnabled** | **optional.Bool**| Filter to retrieve the country permissions with dialing to high-risk [toll fraud](https://www.twilio.com/learn/voice-and-video/toll-fraud) numbers enabled. Can be: &#x60;true&#x60; or &#x60;false&#x60;. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDialingPermissionsHrsPrefixes

> InlineResponse2004 ListDialingPermissionsHrsPrefixes(ctx, isoCode, optional)



Fetch the high-risk special services prefixes from the country resource corresponding to the [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**isoCode** | **string**| The [ISO 3166-1 country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) to identify the country permissions from which high-risk special service number prefixes are fetched | 
 **optional** | ***ListDialingPermissionsHrsPrefixesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDialingPermissionsHrsPrefixesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIpRecord

> InlineResponse2005 ListIpRecord(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListIpRecordOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListIpRecordOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSourceIpMapping

> InlineResponse2006 ListSourceIpMapping(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSourceIpMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSourceIpMappingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateByocTrunk

> VoiceV1ByocTrunk UpdateByocTrunk(ctx, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the BYOC Trunk resource to update. | 
 **optional** | ***UpdateByocTrunkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateByocTrunkOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cnamLookupEnabled** | **optional.Bool**| Whether Caller ID Name (CNAM) lookup is enabled for the trunk. If enabled, all inbound calls to the BYOC Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information. | 
 **connectionPolicySid** | **optional.String**| The SID of the Connection Policy that Twilio will use when routing traffic to your communications infrastructure. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 
 **fromDomainSid** | **optional.String**| The SID of the SIP Domain that should be used in the &#x60;From&#x60; header of originating calls sent to your SIP infrastructure. If your SIP infrastructure allows users to \\\&quot;call back\\\&quot; an incoming call, configure this with a [SIP Domain](https://www.twilio.com/docs/voice/api/sending-sip) to ensure proper routing. If not configured, the from domain will default to \\\&quot;sip.twilio.com\\\&quot;. | 
 **statusCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;status_callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **statusCallbackUrl** | **optional.String**| The URL that we should call to pass status parameters (such as call ended) to your application. | 
 **voiceFallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **voiceFallbackUrl** | **optional.String**| The URL that we should call when an error occurs while retrieving or executing the TwiML requested by &#x60;voice_url&#x60;. | 
 **voiceMethod** | **optional.String**| The HTTP method we should use to call &#x60;voice_url&#x60; | 
 **voiceUrl** | **optional.String**| The URL we should call when the BYOC Trunk receives a call. | 

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

> VoiceV1ConnectionPolicy UpdateConnectionPolicy(ctx, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The unique string that we created to identify the Connection Policy resource to update. | 
 **optional** | ***UpdateConnectionPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateConnectionPolicyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 

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

> VoiceV1ConnectionPolicyConnectionPolicyTarget UpdateConnectionPolicyTarget(ctx, connectionPolicySid, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionPolicySid** | **string**| The SID of the Connection Policy that owns the Target. | 
**sid** | **string**| The unique string that we created to identify the Target resource to update. | 
 **optional** | ***UpdateConnectionPolicyTargetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateConnectionPolicyTargetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enabled** | **optional.Bool**| Whether the Target is enabled. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 
 **priority** | **optional.Int32**| The relative importance of the target. Can be an integer from 0 to 65535, inclusive. The lowest number represents the most important target. | 
 **target** | **optional.String**| The SIP address you want Twilio to route your calls to. This must be a &#x60;sip:&#x60; schema. &#x60;sips&#x60; is NOT supported. | 
 **weight** | **optional.Int32**| The value that determines the relative share of the load the Target should receive compared to other Targets with the same priority. Can be an integer from 1 to 65535, inclusive. Targets with higher values receive more load than those with lower ones with the same priority. | 

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
 **optional** | ***UpdateDialingPermissionsSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDialingPermissionsSettingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dialingPermissionsInheritance** | **optional.Bool**| &#x60;true&#x60; for the sub-account to inherit voice dialing permissions from the Master Project; otherwise &#x60;false&#x60;. | 

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

> VoiceV1IpRecord UpdateIpRecord(ctx, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the IP Record resource to update. | 
 **optional** | ***UpdateIpRecordOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateIpRecordOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long. | 

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

> VoiceV1SourceIpMapping UpdateSourceIpMapping(ctx, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the IP Record resource to update. | 
 **optional** | ***UpdateSourceIpMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSourceIpMappingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sipDomainSid** | **optional.String**| The SID of the SIP Domain that the IP Record should be mapped to. | 

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

