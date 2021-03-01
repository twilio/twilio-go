# DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredentialList**](DefaultApi.md#CreateCredentialList) | **Post** /v1/Trunks/{TrunkSid}/CredentialLists | 
[**CreateIpAccessControlList**](DefaultApi.md#CreateIpAccessControlList) | **Post** /v1/Trunks/{TrunkSid}/IpAccessControlLists | 
[**CreateOriginationUrl**](DefaultApi.md#CreateOriginationUrl) | **Post** /v1/Trunks/{TrunkSid}/OriginationUrls | 
[**CreatePhoneNumber**](DefaultApi.md#CreatePhoneNumber) | **Post** /v1/Trunks/{TrunkSid}/PhoneNumbers | 
[**CreateTrunk**](DefaultApi.md#CreateTrunk) | **Post** /v1/Trunks | 
[**DeleteCredentialList**](DefaultApi.md#DeleteCredentialList) | **Delete** /v1/Trunks/{TrunkSid}/CredentialLists/{Sid} | 
[**DeleteIpAccessControlList**](DefaultApi.md#DeleteIpAccessControlList) | **Delete** /v1/Trunks/{TrunkSid}/IpAccessControlLists/{Sid} | 
[**DeleteOriginationUrl**](DefaultApi.md#DeleteOriginationUrl) | **Delete** /v1/Trunks/{TrunkSid}/OriginationUrls/{Sid} | 
[**DeletePhoneNumber**](DefaultApi.md#DeletePhoneNumber) | **Delete** /v1/Trunks/{TrunkSid}/PhoneNumbers/{Sid} | 
[**DeleteTrunk**](DefaultApi.md#DeleteTrunk) | **Delete** /v1/Trunks/{Sid} | 
[**FetchCredentialList**](DefaultApi.md#FetchCredentialList) | **Get** /v1/Trunks/{TrunkSid}/CredentialLists/{Sid} | 
[**FetchIpAccessControlList**](DefaultApi.md#FetchIpAccessControlList) | **Get** /v1/Trunks/{TrunkSid}/IpAccessControlLists/{Sid} | 
[**FetchOriginationUrl**](DefaultApi.md#FetchOriginationUrl) | **Get** /v1/Trunks/{TrunkSid}/OriginationUrls/{Sid} | 
[**FetchPhoneNumber**](DefaultApi.md#FetchPhoneNumber) | **Get** /v1/Trunks/{TrunkSid}/PhoneNumbers/{Sid} | 
[**FetchRecording**](DefaultApi.md#FetchRecording) | **Get** /v1/Trunks/{TrunkSid}/Recording | 
[**FetchTrunk**](DefaultApi.md#FetchTrunk) | **Get** /v1/Trunks/{Sid} | 
[**ListCredentialList**](DefaultApi.md#ListCredentialList) | **Get** /v1/Trunks/{TrunkSid}/CredentialLists | 
[**ListIpAccessControlList**](DefaultApi.md#ListIpAccessControlList) | **Get** /v1/Trunks/{TrunkSid}/IpAccessControlLists | 
[**ListOriginationUrl**](DefaultApi.md#ListOriginationUrl) | **Get** /v1/Trunks/{TrunkSid}/OriginationUrls | 
[**ListPhoneNumber**](DefaultApi.md#ListPhoneNumber) | **Get** /v1/Trunks/{TrunkSid}/PhoneNumbers | 
[**ListTrunk**](DefaultApi.md#ListTrunk) | **Get** /v1/Trunks | 
[**UpdateOriginationUrl**](DefaultApi.md#UpdateOriginationUrl) | **Post** /v1/Trunks/{TrunkSid}/OriginationUrls/{Sid} | 
[**UpdateRecording**](DefaultApi.md#UpdateRecording) | **Post** /v1/Trunks/{TrunkSid}/Recording | 
[**UpdateTrunk**](DefaultApi.md#UpdateTrunk) | **Post** /v1/Trunks/{Sid} | 



## CreateCredentialList

> TrunkingV1TrunkCredentialList CreateCredentialList(ctx, TrunkSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string**| The SID of the Trunk to associate the credential list with. | 
 **optional** | ***CreateCredentialListRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCredentialListRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CredentialListSid** | **String**| The SID of the [Credential List](https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource) that you want to associate with the trunk. Once associated, we will authenticate access to the trunk against this list. | 

### Return type

[**TrunkingV1TrunkCredentialList**](trunking.v1.trunk.credential_list.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIpAccessControlList

> TrunkingV1TrunkIpAccessControlList CreateIpAccessControlList(ctx, TrunkSid, optional)



Associate an IP Access Control List with a Trunk

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string**| The SID of the Trunk to associate the IP Access Control List with. | 
 **optional** | ***CreateIpAccessControlListRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateIpAccessControlListRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**IpAccessControlListSid** | **String**| The SID of the [IP Access Control List](https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource) that you want to associate with the trunk. | 

### Return type

[**TrunkingV1TrunkIpAccessControlList**](trunking.v1.trunk.ip_access_control_list.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOriginationUrl

> TrunkingV1TrunkOriginationUrl CreateOriginationUrl(ctx, TrunkSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string**| The SID of the Trunk to associate the resource with. | 
 **optional** | ***CreateOriginationUrlRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOriginationUrlRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Enabled** | **Bool**| Whether the URL is enabled. The default is &#x60;true&#x60;. | 
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
**Priority** | **Int32**| The relative importance of the URI. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important URI. | 
**SipUrl** | **String**| The SIP address you want Twilio to route your Origination calls to. This must be a &#x60;sip:&#x60; schema. | 
**Weight** | **Int32**| The value that determines the relative share of the load the URI should receive compared to other URIs with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. URLs with higher values receive more load than those with lower ones with the same priority. | 

### Return type

[**TrunkingV1TrunkOriginationUrl**](trunking.v1.trunk.origination_url.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePhoneNumber

> TrunkingV1TrunkPhoneNumber CreatePhoneNumber(ctx, TrunkSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string**| The SID of the Trunk to associate the phone number with. | 
 **optional** | ***CreatePhoneNumberRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePhoneNumberRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PhoneNumberSid** | **String**| The SID of the [Incoming Phone Number](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) that you want to associate with the trunk. | 

### Return type

[**TrunkingV1TrunkPhoneNumber**](trunking.v1.trunk.phone_number.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTrunk

> TrunkingV1Trunk CreateTrunk(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateTrunkRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateTrunkRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CnamLookupEnabled** | **Bool**| Whether Caller ID Name (CNAM) lookup should be enabled for the trunk. If enabled, all inbound calls to the SIP Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information. | 
**DisasterRecoveryMethod** | **String**| The HTTP method we should use to call the &#x60;disaster_recovery_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
**DisasterRecoveryUrl** | **String**| The URL we should call using the &#x60;disaster_recovery_method&#x60; if an error occurs while sending SIP traffic towards the configured Origination URL. We retrieve TwiML from the URL and execute the instructions like any other normal TwiML call. See [Disaster Recovery](https://www.twilio.com/docs/sip-trunking#disaster-recovery) for more information. | 
**DomainName** | **String**| The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and &#x60;-&#x60; and must end with &#x60;pstn.twilio.com&#x60;. See [Termination Settings](https://www.twilio.com/docs/sip-trunking#termination) for more information. | 
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
**Secure** | **Bool**| Whether Secure Trunking is enabled for the trunk. If enabled, all calls going through the trunk will be secure using SRTP for media and TLS for signaling. If disabled, then RTP will be used for media. See [Secure Trunking](https://www.twilio.com/docs/sip-trunking#securetrunking) for more information. | 
**TransferMode** | **String**| The call transfer settings for the trunk. Can be: &#x60;enable-all&#x60;, &#x60;sip-only&#x60; and &#x60;disable-all&#x60;. See [Transfer](https://www.twilio.com/docs/sip-trunking/call-transfer) for more information. | 

### Return type

[**TrunkingV1Trunk**](trunking.v1.trunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCredentialList

> DeleteCredentialList(ctx, TrunkSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string**| The SID of the Trunk from which to delete the credential list. | 
**Sid** | **string**| The unique string that we created to identify the CredentialList resource to delete. | 

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


## DeleteIpAccessControlList

> DeleteIpAccessControlList(ctx, TrunkSid, Sid)



Remove an associated IP Access Control List from a Trunk

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string**| The SID of the Trunk from which to delete the IP Access Control List. | 
**Sid** | **string**| The unique string that we created to identify the IpAccessControlList resource to delete. | 

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


## DeleteOriginationUrl

> DeleteOriginationUrl(ctx, TrunkSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string**| The SID of the Trunk from which to delete the OriginationUrl. | 
**Sid** | **string**| The unique string that we created to identify the OriginationUrl resource to delete. | 

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


## DeletePhoneNumber

> DeletePhoneNumber(ctx, TrunkSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string**| The SID of the Trunk from which to delete the PhoneNumber resource. | 
**Sid** | **string**| The unique string that we created to identify the PhoneNumber resource to delete. | 

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


## DeleteTrunk

> DeleteTrunk(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique string that we created to identify the Trunk resource to delete. | 

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


## FetchCredentialList

> TrunkingV1TrunkCredentialList FetchCredentialList(ctx, TrunkSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string**| The SID of the Trunk from which to fetch the credential list. | 
**Sid** | **string**| The unique string that we created to identify the CredentialList resource to fetch. | 

### Return type

[**TrunkingV1TrunkCredentialList**](trunking.v1.trunk.credential_list.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIpAccessControlList

> TrunkingV1TrunkIpAccessControlList FetchIpAccessControlList(ctx, TrunkSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string**| The SID of the Trunk from which to fetch the IP Access Control List. | 
**Sid** | **string**| The unique string that we created to identify the IpAccessControlList resource to fetch. | 

### Return type

[**TrunkingV1TrunkIpAccessControlList**](trunking.v1.trunk.ip_access_control_list.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchOriginationUrl

> TrunkingV1TrunkOriginationUrl FetchOriginationUrl(ctx, TrunkSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string**| The SID of the Trunk from which to fetch the OriginationUrl. | 
**Sid** | **string**| The unique string that we created to identify the OriginationUrl resource to fetch. | 

### Return type

[**TrunkingV1TrunkOriginationUrl**](trunking.v1.trunk.origination_url.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPhoneNumber

> TrunkingV1TrunkPhoneNumber FetchPhoneNumber(ctx, TrunkSid, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string**| The SID of the Trunk from which to fetch the PhoneNumber resource. | 
**Sid** | **string**| The unique string that we created to identify the PhoneNumber resource to fetch. | 

### Return type

[**TrunkingV1TrunkPhoneNumber**](trunking.v1.trunk.phone_number.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecording

> TrunkingV1TrunkRecording FetchRecording(ctx, TrunkSid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string**| The SID of the Trunk from which to fetch the recording settings. | 

### Return type

[**TrunkingV1TrunkRecording**](trunking.v1.trunk.recording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTrunk

> TrunkingV1Trunk FetchTrunk(ctx, Sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique string that we created to identify the Trunk resource to fetch. | 

### Return type

[**TrunkingV1Trunk**](trunking.v1.trunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredentialList

> ListCredentialListResponse ListCredentialList(ctx, TrunkSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string**| The SID of the Trunk from which to read the credential lists. | 
 **optional** | ***ListCredentialListRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCredentialListRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListCredentialListResponse**](ListCredentialListResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIpAccessControlList

> ListIpAccessControlListResponse ListIpAccessControlList(ctx, TrunkSid, optional)



List all IP Access Control Lists for a Trunk

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string**| The SID of the Trunk from which to read the IP Access Control Lists. | 
 **optional** | ***ListIpAccessControlListRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListIpAccessControlListRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListIpAccessControlListResponse**](ListIpAccessControlListResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOriginationUrl

> ListOriginationUrlResponse ListOriginationUrl(ctx, TrunkSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string**| The SID of the Trunk from which to read the OriginationUrl. | 
 **optional** | ***ListOriginationUrlRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOriginationUrlRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListOriginationUrlResponse**](ListOriginationUrlResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhoneNumber

> ListPhoneNumberResponse ListPhoneNumber(ctx, TrunkSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string**| The SID of the Trunk from which to read the PhoneNumber resources. | 
 **optional** | ***ListPhoneNumberRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPhoneNumberRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListPhoneNumberResponse**](ListPhoneNumberResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrunk

> ListTrunkResponse ListTrunk(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListTrunkRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListTrunkRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListTrunkResponse**](ListTrunkResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOriginationUrl

> TrunkingV1TrunkOriginationUrl UpdateOriginationUrl(ctx, TrunkSid, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string**| The SID of the Trunk from which to update the OriginationUrl. | 
**Sid** | **string**| The unique string that we created to identify the OriginationUrl resource to update. | 
 **optional** | ***UpdateOriginationUrlRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOriginationUrlRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Enabled** | **Bool**| Whether the URL is enabled. The default is &#x60;true&#x60;. | 
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
**Priority** | **Int32**| The relative importance of the URI. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important URI. | 
**SipUrl** | **String**| The SIP address you want Twilio to route your Origination calls to. This must be a &#x60;sip:&#x60; schema. &#x60;sips&#x60; is NOT supported. | 
**Weight** | **Int32**| The value that determines the relative share of the load the URI should receive compared to other URIs with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. URLs with higher values receive more load than those with lower ones with the same priority. | 

### Return type

[**TrunkingV1TrunkOriginationUrl**](trunking.v1.trunk.origination_url.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRecording

> TrunkingV1TrunkRecording UpdateRecording(ctx, TrunkSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string**| The SID of the Trunk that will have its recording settings updated. | 
 **optional** | ***UpdateRecordingRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateRecordingRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Mode** | **String**| The recording mode for the trunk. Can be do-not-record (default), record-from-ringing, record-from-answer, record-from-ringing-dual, or record-from-answer-dual. | 
**Trim** | **String**| The recording trim setting for the trunk. Can be do-not-trim (default) or trim-silence. | 

### Return type

[**TrunkingV1TrunkRecording**](trunking.v1.trunk.recording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrunk

> TrunkingV1Trunk UpdateTrunk(ctx, Sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The unique string that we created to identify the OriginationUrl resource to update. | 
 **optional** | ***UpdateTrunkRequest** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateTrunkRequest struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CnamLookupEnabled** | **Bool**| Whether Caller ID Name (CNAM) lookup should be enabled for the trunk. If enabled, all inbound calls to the SIP Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information. | 
**DisasterRecoveryMethod** | **String**| The HTTP method we should use to call the &#x60;disaster_recovery_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
**DisasterRecoveryUrl** | **String**| The URL we should call using the &#x60;disaster_recovery_method&#x60; if an error occurs while sending SIP traffic towards the configured Origination URL. We retrieve TwiML from the URL and execute the instructions like any other normal TwiML call. See [Disaster Recovery](https://www.twilio.com/docs/sip-trunking#disaster-recovery) for more information. | 
**DomainName** | **String**| The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and &#x60;-&#x60; and must end with &#x60;pstn.twilio.com&#x60;. See [Termination Settings](https://www.twilio.com/docs/sip-trunking#termination) for more information. | 
**FriendlyName** | **String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
**Secure** | **Bool**| Whether Secure Trunking is enabled for the trunk. If enabled, all calls going through the trunk will be secure using SRTP for media and TLS for signaling. If disabled, then RTP will be used for media. See [Secure Trunking](https://www.twilio.com/docs/sip-trunking#securetrunking) for more information. | 
**TransferMode** | **String**| The call transfer settings for the trunk. Can be: &#x60;enable-all&#x60;, &#x60;sip-only&#x60; and &#x60;disable-all&#x60;. See [Transfer](https://www.twilio.com/docs/sip-trunking/call-transfer) for more information. | 

### Return type

[**TrunkingV1Trunk**](trunking.v1.trunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

