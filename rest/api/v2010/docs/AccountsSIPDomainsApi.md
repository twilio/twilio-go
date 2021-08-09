# AccountsSIPDomainsApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSipDomain**](AccountsSIPDomainsApi.md#CreateSipDomain) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains.json | 
[**DeleteSipDomain**](AccountsSIPDomainsApi.md#DeleteSipDomain) | **Delete** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{Sid}.json | 
[**FetchSipDomain**](AccountsSIPDomainsApi.md#FetchSipDomain) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{Sid}.json | 
[**ListSipDomain**](AccountsSIPDomainsApi.md#ListSipDomain) | **Get** /2010-04-01/Accounts/{AccountSid}/SIP/Domains.json | 
[**UpdateSipDomain**](AccountsSIPDomainsApi.md#UpdateSipDomain) | **Post** /2010-04-01/Accounts/{AccountSid}/SIP/Domains/{Sid}.json | 



## CreateSipDomain

> ApiV2010SipDomain CreateSipDomain(ctx, optional)



Create a new Domain

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSipDomainParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
**ByocTrunkSid** | **string** | The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with.
**DomainName** | **string** | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \\\&quot;-\\\&quot; and must end with &#x60;sip.twilio.com&#x60;.
**EmergencyCallerSid** | **string** | Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call.
**EmergencyCallingEnabled** | **bool** | Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses.
**FriendlyName** | **string** | A descriptive string that you created to describe the resource. It can be up to 64 characters long.
**Secure** | **bool** | Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain.
**SipRegistration** | **bool** | Whether to allow SIP Endpoints to register with the domain to receive calls. Can be &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; allows SIP Endpoints to register with the domain to receive calls, &#x60;false&#x60; does not.
**VoiceFallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs while retrieving or executing the TwiML from &#x60;voice_url&#x60;.
**VoiceMethod** | **string** | The HTTP method we should use to call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**VoiceStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_status_callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**VoiceStatusCallbackUrl** | **string** | The URL that we should call to pass status parameters (such as call ended) to your application.
**VoiceUrl** | **string** | The URL we should when the domain receives a call.

### Return type

[**ApiV2010SipDomain**](ApiV2010SipDomain.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSipDomain

> DeleteSipDomain(ctx, Sidoptional)



Delete an instance of a Domain

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the SipDomain resource to delete.

### Other Parameters

Other parameters are passed through a pointer to a DeleteSipDomainParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resources to delete.

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


## FetchSipDomain

> ApiV2010SipDomain FetchSipDomain(ctx, Sidoptional)



Fetch an instance of a Domain

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the SipDomain resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSipDomainParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resource to fetch.

### Return type

[**ApiV2010SipDomain**](ApiV2010SipDomain.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSipDomain

> ListSipDomainResponse ListSipDomain(ctx, optional)



Retrieve a list of domains belonging to the account used to make the request

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSipDomainParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resources to read.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**ListSipDomainResponse**](ListSipDomainResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSipDomain

> ApiV2010SipDomain UpdateSipDomain(ctx, Sidoptional)



Update the attributes of a domain

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The Twilio-provided string that uniquely identifies the SipDomain resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSipDomainParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resource to update.
**ByocTrunkSid** | **string** | The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with.
**DomainName** | **string** | The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \\\&quot;-\\\&quot; and must end with &#x60;sip.twilio.com&#x60;.
**EmergencyCallerSid** | **string** | Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call.
**EmergencyCallingEnabled** | **bool** | Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses.
**FriendlyName** | **string** | A descriptive string that you created to describe the resource. It can be up to 64 characters long.
**Secure** | **bool** | Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain.
**SipRegistration** | **bool** | Whether to allow SIP Endpoints to register with the domain to receive calls. Can be &#x60;true&#x60; or &#x60;false&#x60;. &#x60;true&#x60; allows SIP Endpoints to register with the domain to receive calls, &#x60;false&#x60; does not.
**VoiceFallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**VoiceFallbackUrl** | **string** | The URL that we should call when an error occurs while retrieving or executing the TwiML requested by &#x60;voice_url&#x60;.
**VoiceMethod** | **string** | The HTTP method we should use to call &#x60;voice_url&#x60;
**VoiceStatusCallbackMethod** | **string** | The HTTP method we should use to call &#x60;voice_status_callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;.
**VoiceStatusCallbackUrl** | **string** | The URL that we should call to pass status parameters (such as call ended) to your application.
**VoiceUrl** | **string** | The URL we should call when the domain receives a call.

### Return type

[**ApiV2010SipDomain**](ApiV2010SipDomain.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

