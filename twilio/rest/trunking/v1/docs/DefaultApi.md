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

> TrunkingV1TrunkCredentialList CreateCredentialList(ctx, TrunkSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk to associate the credential list with.

### Other Parameters

Other parameters are passed through a pointer to a CreateCredentialListParams struct


Name | Type | Description
------------- | ------------- | -------------
**CredentialListSid** | **string** | The SID of the [Credential List](https://www.twilio.com/docs/voice/sip/api/sip-credentiallist-resource) that you want to associate with the trunk. Once associated, we will authenticate access to the trunk against this list.

### Return type

[**TrunkingV1TrunkCredentialList**](TrunkingV1TrunkCredentialList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIpAccessControlList

> TrunkingV1TrunkIpAccessControlList CreateIpAccessControlList(ctx, TrunkSidoptional)



Associate an IP Access Control List with a Trunk

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk to associate the IP Access Control List with.

### Other Parameters

Other parameters are passed through a pointer to a CreateIpAccessControlListParams struct


Name | Type | Description
------------- | ------------- | -------------
**IpAccessControlListSid** | **string** | The SID of the [IP Access Control List](https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource) that you want to associate with the trunk.

### Return type

[**TrunkingV1TrunkIpAccessControlList**](TrunkingV1TrunkIpAccessControlList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOriginationUrl

> TrunkingV1TrunkOriginationUrl CreateOriginationUrl(ctx, TrunkSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk to associate the resource with.

### Other Parameters

Other parameters are passed through a pointer to a CreateOriginationUrlParams struct


Name | Type | Description
------------- | ------------- | -------------
**Enabled** | **bool** | Whether the URL is enabled. The default is &#x60;true&#x60;.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**Priority** | **int32** | The relative importance of the URI. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important URI.
**SipUrl** | **string** | The SIP address you want Twilio to route your Origination calls to. This must be a &#x60;sip:&#x60; schema.
**Weight** | **int32** | The value that determines the relative share of the load the URI should receive compared to other URIs with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. URLs with higher values receive more load than those with lower ones with the same priority.

### Return type

[**TrunkingV1TrunkOriginationUrl**](TrunkingV1TrunkOriginationUrl.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePhoneNumber

> TrunkingV1TrunkPhoneNumber CreatePhoneNumber(ctx, TrunkSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk to associate the phone number with.

### Other Parameters

Other parameters are passed through a pointer to a CreatePhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------
**PhoneNumberSid** | **string** | The SID of the [Incoming Phone Number](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) that you want to associate with the trunk.

### Return type

[**TrunkingV1TrunkPhoneNumber**](TrunkingV1TrunkPhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTrunk

> TrunkingV1Trunk CreateTrunk(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateTrunkParams struct


Name | Type | Description
------------- | ------------- | -------------
**CnamLookupEnabled** | **bool** | Whether Caller ID Name (CNAM) lookup should be enabled for the trunk. If enabled, all inbound calls to the SIP Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
**DisasterRecoveryMethod** | **string** | The HTTP method we should use to call the &#x60;disaster_recovery_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**DisasterRecoveryUrl** | **string** | The URL we should call using the &#x60;disaster_recovery_method&#x60; if an error occurs while sending SIP traffic towards the configured Origination URL. We retrieve TwiML from the URL and execute the instructions like any other normal TwiML call. See [Disaster Recovery](https://www.twilio.com/docs/sip-trunking#disaster-recovery) for more information.
**DomainName** | **string** | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and &#x60;-&#x60; and must end with &#x60;pstn.twilio.com&#x60;. See [Termination Settings](https://www.twilio.com/docs/sip-trunking#termination) for more information.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**Secure** | **bool** | Whether Secure Trunking is enabled for the trunk. If enabled, all calls going through the trunk will be secure using SRTP for media and TLS for signaling. If disabled, then RTP will be used for media. See [Secure Trunking](https://www.twilio.com/docs/sip-trunking#securetrunking) for more information.
**TransferMode** | **string** | The call transfer settings for the trunk. Can be: &#x60;enable-all&#x60;, &#x60;sip-only&#x60; and &#x60;disable-all&#x60;. See [Transfer](https://www.twilio.com/docs/sip-trunking/call-transfer) for more information.

### Return type

[**TrunkingV1Trunk**](TrunkingV1Trunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCredentialList

> DeleteCredentialList(ctx, TrunkSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to delete the credential list.
**Sid** | **string** | The unique string that we created to identify the CredentialList resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteCredentialListParams struct


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


## DeleteIpAccessControlList

> DeleteIpAccessControlList(ctx, TrunkSidSid)



Remove an associated IP Access Control List from a Trunk

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to delete the IP Access Control List.
**Sid** | **string** | The unique string that we created to identify the IpAccessControlList resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteIpAccessControlListParams struct


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


## DeleteOriginationUrl

> DeleteOriginationUrl(ctx, TrunkSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to delete the OriginationUrl.
**Sid** | **string** | The unique string that we created to identify the OriginationUrl resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteOriginationUrlParams struct


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


## DeletePhoneNumber

> DeletePhoneNumber(ctx, TrunkSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to delete the PhoneNumber resource.
**Sid** | **string** | The unique string that we created to identify the PhoneNumber resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeletePhoneNumberParams struct


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


## DeleteTrunk

> DeleteTrunk(ctx, Sid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Trunk resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteTrunkParams struct


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


## FetchCredentialList

> TrunkingV1TrunkCredentialList FetchCredentialList(ctx, TrunkSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to fetch the credential list.
**Sid** | **string** | The unique string that we created to identify the CredentialList resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchCredentialListParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrunkingV1TrunkCredentialList**](TrunkingV1TrunkCredentialList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIpAccessControlList

> TrunkingV1TrunkIpAccessControlList FetchIpAccessControlList(ctx, TrunkSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to fetch the IP Access Control List.
**Sid** | **string** | The unique string that we created to identify the IpAccessControlList resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchIpAccessControlListParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrunkingV1TrunkIpAccessControlList**](TrunkingV1TrunkIpAccessControlList.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchOriginationUrl

> TrunkingV1TrunkOriginationUrl FetchOriginationUrl(ctx, TrunkSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to fetch the OriginationUrl.
**Sid** | **string** | The unique string that we created to identify the OriginationUrl resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchOriginationUrlParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrunkingV1TrunkOriginationUrl**](TrunkingV1TrunkOriginationUrl.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPhoneNumber

> TrunkingV1TrunkPhoneNumber FetchPhoneNumber(ctx, TrunkSidSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to fetch the PhoneNumber resource.
**Sid** | **string** | The unique string that we created to identify the PhoneNumber resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchPhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrunkingV1TrunkPhoneNumber**](TrunkingV1TrunkPhoneNumber.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRecording

> TrunkingV1TrunkRecording FetchRecording(ctx, TrunkSid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to fetch the recording settings.

### Other Parameters

Other parameters are passed through a pointer to a FetchRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrunkingV1TrunkRecording**](TrunkingV1TrunkRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchTrunk

> TrunkingV1Trunk FetchTrunk(ctx, Sid)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the Trunk resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchTrunkParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**TrunkingV1Trunk**](TrunkingV1Trunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredentialList

> ListCredentialListResponse ListCredentialList(ctx, TrunkSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to read the credential lists.

### Other Parameters

Other parameters are passed through a pointer to a ListCredentialListParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListCredentialListResponse**](ListCredentialListResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIpAccessControlList

> ListIpAccessControlListResponse ListIpAccessControlList(ctx, TrunkSidoptional)



List all IP Access Control Lists for a Trunk

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to read the IP Access Control Lists.

### Other Parameters

Other parameters are passed through a pointer to a ListIpAccessControlListParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListIpAccessControlListResponse**](ListIpAccessControlListResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOriginationUrl

> ListOriginationUrlResponse ListOriginationUrl(ctx, TrunkSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to read the OriginationUrl.

### Other Parameters

Other parameters are passed through a pointer to a ListOriginationUrlParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListOriginationUrlResponse**](ListOriginationUrlResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhoneNumber

> ListPhoneNumberResponse ListPhoneNumber(ctx, TrunkSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to read the PhoneNumber resources.

### Other Parameters

Other parameters are passed through a pointer to a ListPhoneNumberParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListPhoneNumberResponse**](ListPhoneNumberResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrunk

> ListTrunkResponse ListTrunk(ctx, optional)



### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListTrunkParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000.

### Return type

[**ListTrunkResponse**](ListTrunkResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOriginationUrl

> TrunkingV1TrunkOriginationUrl UpdateOriginationUrl(ctx, TrunkSidSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk from which to update the OriginationUrl.
**Sid** | **string** | The unique string that we created to identify the OriginationUrl resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateOriginationUrlParams struct


Name | Type | Description
------------- | ------------- | -------------
**Enabled** | **bool** | Whether the URL is enabled. The default is &#x60;true&#x60;.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**Priority** | **int32** | The relative importance of the URI. Can be an integer from 0 to 65535, inclusive, and the default is 10. The lowest number represents the most important URI.
**SipUrl** | **string** | The SIP address you want Twilio to route your Origination calls to. This must be a &#x60;sip:&#x60; schema. &#x60;sips&#x60; is NOT supported.
**Weight** | **int32** | The value that determines the relative share of the load the URI should receive compared to other URIs with the same priority. Can be an integer from 1 to 65535, inclusive, and the default is 10. URLs with higher values receive more load than those with lower ones with the same priority.

### Return type

[**TrunkingV1TrunkOriginationUrl**](TrunkingV1TrunkOriginationUrl.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRecording

> TrunkingV1TrunkRecording UpdateRecording(ctx, TrunkSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**TrunkSid** | **string** | The SID of the Trunk that will have its recording settings updated.

### Other Parameters

Other parameters are passed through a pointer to a UpdateRecordingParams struct


Name | Type | Description
------------- | ------------- | -------------
**Mode** | **string** | The recording mode for the trunk. Can be do-not-record (default), record-from-ringing, record-from-answer, record-from-ringing-dual, or record-from-answer-dual.
**Trim** | **string** | The recording trim setting for the trunk. Can be do-not-trim (default) or trim-silence.

### Return type

[**TrunkingV1TrunkRecording**](TrunkingV1TrunkRecording.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTrunk

> TrunkingV1Trunk UpdateTrunk(ctx, Sidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The unique string that we created to identify the OriginationUrl resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateTrunkParams struct


Name | Type | Description
------------- | ------------- | -------------
**CnamLookupEnabled** | **bool** | Whether Caller ID Name (CNAM) lookup should be enabled for the trunk. If enabled, all inbound calls to the SIP Trunk from the United States and Canada automatically perform a CNAM Lookup and display Caller ID data on your phone. See [CNAM Lookups](https://www.twilio.com/docs/sip-trunking#CNAM) for more information.
**DisasterRecoveryMethod** | **string** | The HTTP method we should use to call the &#x60;disaster_recovery_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**DisasterRecoveryUrl** | **string** | The URL we should call using the &#x60;disaster_recovery_method&#x60; if an error occurs while sending SIP traffic towards the configured Origination URL. We retrieve TwiML from the URL and execute the instructions like any other normal TwiML call. See [Disaster Recovery](https://www.twilio.com/docs/sip-trunking#disaster-recovery) for more information.
**DomainName** | **string** | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and &#x60;-&#x60; and must end with &#x60;pstn.twilio.com&#x60;. See [Termination Settings](https://www.twilio.com/docs/sip-trunking#termination) for more information.
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long.
**Secure** | **bool** | Whether Secure Trunking is enabled for the trunk. If enabled, all calls going through the trunk will be secure using SRTP for media and TLS for signaling. If disabled, then RTP will be used for media. See [Secure Trunking](https://www.twilio.com/docs/sip-trunking#securetrunking) for more information.
**TransferMode** | **string** | The call transfer settings for the trunk. Can be: &#x60;enable-all&#x60;, &#x60;sip-only&#x60; and &#x60;disable-all&#x60;. See [Transfer](https://www.twilio.com/docs/sip-trunking/call-transfer) for more information.

### Return type

[**TrunkingV1Trunk**](TrunkingV1Trunk.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

