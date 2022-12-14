# ServicesComplianceUsa2pApi

All URIs are relative to *https://messaging.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsAppToPerson**](ServicesComplianceUsa2pApi.md#CreateUsAppToPerson) | **Post** /v1/Services/{MessagingServiceSid}/Compliance/Usa2p | 
[**DeleteUsAppToPerson**](ServicesComplianceUsa2pApi.md#DeleteUsAppToPerson) | **Delete** /v1/Services/{MessagingServiceSid}/Compliance/Usa2p/{Sid} | 
[**FetchUsAppToPerson**](ServicesComplianceUsa2pApi.md#FetchUsAppToPerson) | **Get** /v1/Services/{MessagingServiceSid}/Compliance/Usa2p/{Sid} | 
[**ListUsAppToPerson**](ServicesComplianceUsa2pApi.md#ListUsAppToPerson) | **Get** /v1/Services/{MessagingServiceSid}/Compliance/Usa2p | 



## CreateUsAppToPerson

> MessagingV1UsAppToPerson CreateUsAppToPerson(ctx, MessagingServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) to create the resources from.

### Other Parameters

Other parameters are passed through a pointer to a CreateUsAppToPersonParams struct


Name | Type | Description
------------- | ------------- | -------------
**BrandRegistrationSid** | **string** | A2P Brand Registration SID
**Description** | **string** | A short description of what this SMS campaign does. Min length: 40 characters. Max length: 4096 characters.
**MessageFlow** | **string** | Required for all Campaigns. Details around how a consumer opts-in to their campaign, therefore giving consent to receive their messages. If multiple opt-in methods can be used for the same campaign, they must all be listed. 40 character minimum. 2048 character maximum.
**MessageSamples** | **[]string** | Message samples, at least 1 and up to 5 sample messages (at least 2 for sole proprietor), >=20 chars, <=1024 chars each.
**UsAppToPersonUsecase** | **string** | A2P Campaign Use Case. Examples: [ 2FA, EMERGENCY, MARKETING..]
**HasEmbeddedLinks** | **bool** | Indicates that this SMS campaign will send messages that contain links.
**HasEmbeddedPhone** | **bool** | Indicates that this SMS campaign will send messages that contain phone numbers.
**OptInMessage** | **string** | If end users can text in a keyword to start receiving messages from this campaign, the auto-reply messages sent to the end users must be provided. The opt-in response should include the Brand name, confirmation of opt-in enrollment to a recurring message campaign, how to get help, and clear description of how to opt-out. This field is required if end users can text in a keyword to start receiving messages from this campaign. 20 character minimum. 320 character maximum.
**OptOutMessage** | **string** | Upon receiving the opt-out keywords from the end users, Twilio customers are expected to send back an auto-generated response, which must provide acknowledgment of the opt-out request and confirmation that no further messages will be sent. It is also recommended that these opt-out messages include the brand name. This field is required if managing opt out keywords yourself (i.e. not using Twilio's Default or Advanced Opt Out features). 20 character minimum. 320 character maximum.
**HelpMessage** | **string** | When customers receive the help keywords from their end users, Twilio customers are expected to send back an auto-generated response; this may include the brand name and additional support contact information. This field is required if managing help keywords yourself (i.e. not using Twilio's Default or Advanced Opt Out features). 20 character minimum. 320 character maximum.
**OptInKeywords** | **[]string** | If end users can text in a keyword to start receiving messages from this campaign, those keywords must be provided. This field is required if end users can text in a keyword to start receiving messages from this campaign. Values must be alphanumeric. 255 character maximum.
**OptOutKeywords** | **[]string** | End users should be able to text in a keyword to stop receiving messages from this campaign. Those keywords must be provided. This field is required if managing opt out keywords yourself (i.e. not using Twilio's Default or Advanced Opt Out features). Values must be alphanumeric. 255 character maximum.
**HelpKeywords** | **[]string** | End users should be able to text in a keyword to receive help. Those keywords must be provided as part of the campaign registration request. This field is required if managing help keywords yourself (i.e. not using Twilio's Default or Advanced Opt Out features). Values must be alphanumeric. 255 character maximum.

### Return type

[**MessagingV1UsAppToPerson**](MessagingV1UsAppToPerson.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUsAppToPerson

> DeleteUsAppToPerson(ctx, MessagingServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) to delete the resource from.
**Sid** | **string** | The SID of the US A2P Compliance resource to delete `QE2c6890da8086d771620e9b13fadeba0b`.

### Other Parameters

Other parameters are passed through a pointer to a DeleteUsAppToPersonParams struct


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


## FetchUsAppToPerson

> MessagingV1UsAppToPerson FetchUsAppToPerson(ctx, MessagingServiceSidSid)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) to fetch the resource from.
**Sid** | **string** | The SID of the US A2P Compliance resource to fetch `QE2c6890da8086d771620e9b13fadeba0b`.

### Other Parameters

Other parameters are passed through a pointer to a FetchUsAppToPersonParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**MessagingV1UsAppToPerson**](MessagingV1UsAppToPerson.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsAppToPerson

> []MessagingV1UsAppToPerson ListUsAppToPerson(ctx, MessagingServiceSidoptional)





### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/messaging/services/api) to fetch the resource from.

### Other Parameters

Other parameters are passed through a pointer to a ListUsAppToPersonParams struct


Name | Type | Description
------------- | ------------- | -------------
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]MessagingV1UsAppToPerson**](MessagingV1UsAppToPerson.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

