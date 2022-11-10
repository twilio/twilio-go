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
**Description** | **string** | A short description of what this SMS campaign does.
**MessageSamples** | **[]string** | Message samples, at least 2 and up to 5 sample messages, &lt;&#x3D;1024 chars each.
**UsAppToPersonUsecase** | **string** | A2P Campaign Use Case. Examples: [ 2FA, EMERGENCY, MARKETING..]
**HasEmbeddedLinks** | **bool** | Indicates that this SMS campaign will send messages that contain links.
**HasEmbeddedPhone** | **bool** | Indicates that this SMS campaign will send messages that contain phone numbers.
**MessageFlow** | **string** | Description of how end users opt-in to the SMS campaign, therefore giving consent to receive messages.
**OptInMessage** | **string** | The message that will be sent to the user when they opt in to the SMS campaign.
**OptOutMessage** | **string** | The message that will be sent to the user when they opt out of the SMS campaign.
**HelpMessage** | **string** | The message that will be sent to the user when they request help for the SMS campaign.
**OptInKeywords** | **[]string** | The keywords that will be used to opt in to the SMS campaign.
**OptOutKeywords** | **[]string** | The keywords that will be used to opt out of the SMS campaign.
**HelpKeywords** | **[]string** | The keywords that will be used to request help for the SMS campaign.

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
**Sid** | **string** | The SID of the US A2P Compliance resource to delete &#x60;QE2c6890da8086d771620e9b13fadeba0b&#x60;.

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
**Sid** | **string** | The SID of the US A2P Compliance resource to fetch &#x60;QE2c6890da8086d771620e9b13fadeba0b&#x60;.

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

