# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBinding**](DefaultApi.md#CreateBinding) | **Post** /v1/Services/{ServiceSid}/Bindings | 
[**CreateCredential**](DefaultApi.md#CreateCredential) | **Post** /v1/Credentials | 
[**CreateNotification**](DefaultApi.md#CreateNotification) | **Post** /v1/Services/{ServiceSid}/Notifications | 
[**CreateService**](DefaultApi.md#CreateService) | **Post** /v1/Services | 
[**DeleteBinding**](DefaultApi.md#DeleteBinding) | **Delete** /v1/Services/{ServiceSid}/Bindings/{Sid} | 
[**DeleteCredential**](DefaultApi.md#DeleteCredential) | **Delete** /v1/Credentials/{Sid} | 
[**DeleteService**](DefaultApi.md#DeleteService) | **Delete** /v1/Services/{Sid} | 
[**FetchBinding**](DefaultApi.md#FetchBinding) | **Get** /v1/Services/{ServiceSid}/Bindings/{Sid} | 
[**FetchCredential**](DefaultApi.md#FetchCredential) | **Get** /v1/Credentials/{Sid} | 
[**FetchService**](DefaultApi.md#FetchService) | **Get** /v1/Services/{Sid} | 
[**ListBinding**](DefaultApi.md#ListBinding) | **Get** /v1/Services/{ServiceSid}/Bindings | 
[**ListCredential**](DefaultApi.md#ListCredential) | **Get** /v1/Credentials | 
[**ListService**](DefaultApi.md#ListService) | **Get** /v1/Services | 
[**UpdateCredential**](DefaultApi.md#UpdateCredential) | **Post** /v1/Credentials/{Sid} | 
[**UpdateService**](DefaultApi.md#UpdateService) | **Post** /v1/Services/{Sid} | 



## CreateBinding

> NotifyV1ServiceBinding CreateBinding(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/notify/api/service-resource) to create the resource under. | 
 **optional** | ***CreateBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateBindingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **address** | **optional.String**| The channel-specific address. For APNS, the device token. For FCM and GCM, the registration token. For SMS, a phone number in E.164 format. For Facebook Messenger, the Messenger ID of the user or a phone number in E.164 format. | 
 **bindingType** | **optional.String**| The transport technology to use for the Binding. Can be: &#x60;apn&#x60;, &#x60;fcm&#x60;, &#x60;gcm&#x60;, &#x60;sms&#x60;, or &#x60;facebook-messenger&#x60;. | 
 **credentialSid** | **optional.String**| The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) resource to be used to send notifications to this Binding. If present, this overrides the Credential specified in the Service resource. Applies to only &#x60;apn&#x60;, &#x60;fcm&#x60;, and &#x60;gcm&#x60; type Bindings. | 
 **endpoint** | **optional.String**| Deprecated. | 
 **identity** | **optional.String**| The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/notify/api/service-resource). Up to 20 Bindings can be created for the same Identity in a given Service. | 
 **notificationProtocolVersion** | **optional.String**| The protocol version to use to send the notification. This defaults to the value of &#x60;default_xxxx_notification_protocol_version&#x60; for the protocol in the [Service](https://www.twilio.com/docs/notify/api/service-resource). The current version is &#x60;\\\&quot;3\\\&quot;&#x60; for &#x60;apn&#x60;, &#x60;fcm&#x60;, and &#x60;gcm&#x60; type Bindings. The parameter is not applicable to &#x60;sms&#x60; and &#x60;facebook-messenger&#x60; type Bindings as the data format is fixed. | 
 **tag** | [**optional.Interface of []string**](string.md)| A tag that can be used to select the Bindings to notify. Repeat this parameter to specify more than one tag, up to a total of 20 tags. | 

### Return type

[**NotifyV1ServiceBinding**](notify.v1.service.binding.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCredential

> NotifyV1Credential CreateCredential(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCredentialOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKey** | **optional.String**| [GCM only] The &#x60;Server key&#x60; of your project from Firebase console under Settings / Cloud messaging. | 
 **certificate** | **optional.String**| [APN only] The URL-encoded representation of the certificate. Strip everything outside of the headers, e.g. &#x60;-----BEGIN CERTIFICATE-----MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEFBQAwgZYxCzAJBgNV.....A&#x3D;&#x3D;-----END CERTIFICATE-----&#x60; | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
 **privateKey** | **optional.String**| [APN only] The URL-encoded representation of the private key. Strip everything outside of the headers, e.g. &#x60;-----BEGIN RSA PRIVATE KEY-----MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fGgvCI1l9s+cmBY3WIz+cUDqmxiieR\\\\n.-----END RSA PRIVATE KEY-----&#x60; | 
 **sandbox** | **optional.Bool**| [APN only] Whether to send the credential to sandbox APNs. Can be &#x60;true&#x60; to send to sandbox APNs or &#x60;false&#x60; to send to production. | 
 **secret** | **optional.String**| [FCM only] The &#x60;Server key&#x60; of your project from Firebase console under Settings / Cloud messaging. | 
 **type_** | **optional.String**| The Credential type. Can be: &#x60;gcm&#x60;, &#x60;fcm&#x60;, or &#x60;apn&#x60;. | 

### Return type

[**NotifyV1Credential**](notify.v1.credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNotification

> NotifyV1ServiceNotification CreateNotification(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/notify/api/service-resource) to create the resource under. | 
 **optional** | ***CreateNotificationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateNotificationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **optional.String**| The actions to display for the notification. For APNS, translates to the &#x60;aps.category&#x60; value. For GCM, translates to the &#x60;data.twi_action&#x60; value. For SMS, this parameter is not supported and is omitted from deliveries to those channels. | 
 **alexa** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| Deprecated. | 
 **apn** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The APNS-specific payload that overrides corresponding attributes in the generic payload for APNS Bindings. This property maps to the APNS &#x60;Payload&#x60; item, therefore the &#x60;aps&#x60; key must be used to change standard attributes. Adds custom key-value pairs to the root of the dictionary. See the [APNS documentation](https://developer.apple.com/library/content/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/CommunicatingwithAPNs.html) for more details. We reserve keys that start with &#x60;twi_&#x60; for future use. Custom keys that start with &#x60;twi_&#x60; are not allowed. | 
 **body** | **optional.String**| The notification text. For FCM and GCM, translates to &#x60;data.twi_body&#x60;. For APNS, translates to &#x60;aps.alert.body&#x60;. For SMS, translates to &#x60;body&#x60;. SMS requires either this &#x60;body&#x60; value, or &#x60;media_urls&#x60; attribute defined in the &#x60;sms&#x60; parameter of the notification. | 
 **data** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The custom key-value pairs of the notification&#39;s payload. For FCM and GCM, this value translates to &#x60;data&#x60; in the FCM and GCM payloads. FCM and GCM [reserve certain keys](https://firebase.google.com/docs/cloud-messaging/http-server-ref) that cannot be used in those channels. For APNS, attributes of &#x60;data&#x60; are inserted into the APNS payload as custom properties outside of the &#x60;aps&#x60; dictionary. In all channels, we reserve keys that start with &#x60;twi_&#x60; for future use. Custom keys that start with &#x60;twi_&#x60; are not allowed and are rejected as 400 Bad request with no delivery attempted. For SMS, this parameter is not supported and is omitted from deliveries to those channels. | 
 **deliveryCallbackUrl** | **optional.String**| URL to send webhooks. | 
 **facebookMessenger** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| Deprecated. | 
 **fcm** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The FCM-specific payload that overrides corresponding attributes in the generic payload for FCM Bindings. This property maps to the root JSON dictionary. See the [FCM documentation](https://firebase.google.com/docs/cloud-messaging/http-server-ref#downstream) for more details. Target parameters &#x60;to&#x60;, &#x60;registration_ids&#x60;, &#x60;condition&#x60;, and &#x60;notification_key&#x60; are not allowed in this parameter. We reserve keys that start with &#x60;twi_&#x60; for future use. Custom keys that start with &#x60;twi_&#x60; are not allowed. FCM also [reserves certain keys](https://firebase.google.com/docs/cloud-messaging/http-server-ref), which cannot be used in that channel. | 
 **gcm** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The GCM-specific payload that overrides corresponding attributes in the generic payload for GCM Bindings.  This property maps to the root JSON dictionary. See the [GCM documentation](https://firebase.google.com/docs/cloud-messaging/http-server-ref) for more details. Target parameters &#x60;to&#x60;, &#x60;registration_ids&#x60;, and &#x60;notification_key&#x60; are not allowed. We reserve keys that start with &#x60;twi_&#x60; for future use. Custom keys that start with &#x60;twi_&#x60; are not allowed. GCM also [reserves certain keys](https://firebase.google.com/docs/cloud-messaging/http-server-ref). | 
 **identity** | [**optional.Interface of []string**](string.md)| The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/notify/api/service-resource). Delivery will be attempted only to Bindings with an Identity in this list. No more than 20 items are allowed in this list. | 
 **priority** | **optional.String**| The priority of the notification. Can be: &#x60;low&#x60; or &#x60;high&#x60; and the default is &#x60;high&#x60;. A value of &#x60;low&#x60; optimizes the client app&#39;s battery consumption; however, notifications may be delivered with unspecified delay. For FCM and GCM, &#x60;low&#x60; priority is the same as &#x60;Normal&#x60; priority. For APNS &#x60;low&#x60; priority is the same as &#x60;5&#x60;. A value of &#x60;high&#x60; sends the notification immediately, and can wake up a sleeping device. For FCM and GCM, &#x60;high&#x60; is the same as &#x60;High&#x60; priority. For APNS, &#x60;high&#x60; is a priority &#x60;10&#x60;. SMS does not support this property. | 
 **segment** | [**optional.Interface of []string**](string.md)| The Segment resource is deprecated. Use the &#x60;tag&#x60; parameter, instead. | 
 **sms** | [**optional.Interface of map[string]interface{}**](map[string]interface{}.md)| The SMS-specific payload that overrides corresponding attributes in the generic payload for SMS Bindings.  Each attribute in this value maps to the corresponding &#x60;form&#x60; parameter of the Twilio [Message](https://www.twilio.com/docs/sms/send-messages) resource.  These parameters of the Message resource are supported in snake case format: &#x60;body&#x60;, &#x60;media_urls&#x60;, &#x60;status_callback&#x60;, and &#x60;max_price&#x60;.  The &#x60;status_callback&#x60; parameter overrides the corresponding parameter in the messaging service, if configured. The &#x60;media_urls&#x60; property expects a JSON array. | 
 **sound** | **optional.String**| The name of the sound to be played for the notification. For FCM and GCM, this Translates to &#x60;data.twi_sound&#x60;.  For APNS, this translates to &#x60;aps.sound&#x60;.  SMS does not support this property. | 
 **tag** | [**optional.Interface of []string**](string.md)| A tag that selects the Bindings to notify. Repeat this parameter to specify more than one tag, up to a total of 5 tags. The implicit tag &#x60;all&#x60; is available to notify all Bindings in a Service instance. Similarly, the implicit tags &#x60;apn&#x60;, &#x60;fcm&#x60;, &#x60;gcm&#x60;, &#x60;sms&#x60; and &#x60;facebook-messenger&#x60; are available to notify all Bindings in a specific channel. | 
 **title** | **optional.String**| The notification title. For FCM and GCM, this translates to the &#x60;data.twi_title&#x60; value. For APNS, this translates to the &#x60;aps.alert.title&#x60; value. SMS does not support this property. This field is not visible on iOS phones and tablets but appears on Apple Watch and Android devices. | 
 **toBinding** | [**optional.Interface of []string**](string.md)| The destination address specified as a JSON string.  Multiple &#x60;to_binding&#x60; parameters can be included but the total size of the request entity should not exceed 1MB. This is typically sufficient for 10,000 phone numbers. | 
 **ttl** | **optional.Int32**| How long, in seconds, the notification is valid. Can be an integer between 0 and 2,419,200, which is 4 weeks, the default and the maximum supported time to live (TTL). Delivery should be attempted if the device is offline until the TTL elapses. Zero means that the notification delivery is attempted immediately, only once, and is not stored for future delivery. SMS does not support this property. | 

### Return type

[**NotifyV1ServiceNotification**](notify.v1.service.notification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateService

> NotifyV1Service CreateService(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alexaSkillId** | **optional.String**| Deprecated. | 
 **apnCredentialSid** | **optional.String**| The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for APN Bindings. | 
 **defaultAlexaNotificationProtocolVersion** | **optional.String**| Deprecated. | 
 **defaultApnNotificationProtocolVersion** | **optional.String**| The protocol version to use for sending APNS notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource. | 
 **defaultFcmNotificationProtocolVersion** | **optional.String**| The protocol version to use for sending FCM notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource. | 
 **defaultGcmNotificationProtocolVersion** | **optional.String**| The protocol version to use for sending GCM notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource. | 
 **deliveryCallbackEnabled** | **optional.Bool**| Callback configuration that enables delivery callbacks, default false | 
 **deliveryCallbackUrl** | **optional.String**| URL to send delivery status callback. | 
 **facebookMessengerPageId** | **optional.String**| Deprecated. | 
 **fcmCredentialSid** | **optional.String**| The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for FCM Bindings. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
 **gcmCredentialSid** | **optional.String**| The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for GCM Bindings. | 
 **logEnabled** | **optional.Bool**| Whether to log notifications. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **messagingServiceSid** | **optional.String**| The SID of the [Messaging Service](https://www.twilio.com/docs/sms/send-messages#messaging-services) to use for SMS Bindings. This parameter must be set in order to send SMS notifications. | 

### Return type

[**NotifyV1Service**](notify.v1.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBinding

> DeleteBinding(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/notify/api/service-resource) to delete the resource from. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the Binding resource to delete. | 

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


## DeleteCredential

> DeleteCredential(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the Credential resource to delete. | 

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


## DeleteService

> DeleteService(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the Service resource to delete. | 

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


## FetchBinding

> NotifyV1ServiceBinding FetchBinding(ctx, serviceSid, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/notify/api/service-resource) to fetch the resource from. | 
**sid** | **string**| The Twilio-provided string that uniquely identifies the Binding resource to fetch. | 

### Return type

[**NotifyV1ServiceBinding**](notify.v1.service.binding.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCredential

> NotifyV1Credential FetchCredential(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the Credential resource to fetch. | 

### Return type

[**NotifyV1Credential**](notify.v1.credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchService

> NotifyV1Service FetchService(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the Service resource to fetch. | 

### Return type

[**NotifyV1Service**](notify.v1.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBinding

> NotifyV1ServiceBindingReadResponse ListBinding(ctx, serviceSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceSid** | **string**| The SID of the [Service](https://www.twilio.com/docs/notify/api/service-resource) to read the resource from. | 
 **optional** | ***ListBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListBindingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **optional.Time**| Only include usage that has occurred on or after this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. | 
 **endDate** | **optional.Time**| Only include usage that occurred on or before this date. Specify the date in GMT and format as &#x60;YYYY-MM-DD&#x60;. | 
 **identity** | [**optional.Interface of []string**](string.md)| The [User](https://www.twilio.com/docs/chat/rest/user-resource)&#39;s &#x60;identity&#x60; value of the resources to read. | 
 **tag** | [**optional.Interface of []string**](string.md)| Only list Bindings that have all of the specified Tags. The following implicit tags are available: &#x60;all&#x60;, &#x60;apn&#x60;, &#x60;fcm&#x60;, &#x60;gcm&#x60;, &#x60;sms&#x60;, &#x60;facebook-messenger&#x60;. Up to 5 tags are allowed. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**NotifyV1ServiceBindingReadResponse**](notify_v1_service_bindingReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredential

> NotifyV1CredentialReadResponse ListCredential(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCredentialOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**NotifyV1CredentialReadResponse**](notify_v1_credentialReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListService

> NotifyV1ServiceReadResponse ListService(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **friendlyName** | **optional.String**| The string that identifies the Service resources to read. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**NotifyV1ServiceReadResponse**](notify_v1_serviceReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredential

> NotifyV1Credential UpdateCredential(ctx, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the Credential resource to update. | 
 **optional** | ***UpdateCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCredentialOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKey** | **optional.String**| [GCM only] The &#x60;Server key&#x60; of your project from Firebase console under Settings / Cloud messaging. | 
 **certificate** | **optional.String**| [APN only] The URL-encoded representation of the certificate. Strip everything outside of the headers, e.g. &#x60;-----BEGIN CERTIFICATE-----MIIFnTCCBIWgAwIBAgIIAjy9H849+E8wDQYJKoZIhvcNAQEFBQAwgZYxCzAJBgNV.....A&#x3D;&#x3D;-----END CERTIFICATE-----&#x60; | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
 **privateKey** | **optional.String**| [APN only] The URL-encoded representation of the private key. Strip everything outside of the headers, e.g. &#x60;-----BEGIN RSA PRIVATE KEY-----MIIEpQIBAAKCAQEAuyf/lNrH9ck8DmNyo3fGgvCI1l9s+cmBY3WIz+cUDqmxiieR\\\\n.-----END RSA PRIVATE KEY-----&#x60; | 
 **sandbox** | **optional.Bool**| [APN only] Whether to send the credential to sandbox APNs. Can be &#x60;true&#x60; to send to sandbox APNs or &#x60;false&#x60; to send to production. | 
 **secret** | **optional.String**| [FCM only] The &#x60;Server key&#x60; of your project from Firebase console under Settings / Cloud messaging. | 

### Return type

[**NotifyV1Credential**](notify.v1.credential.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateService

> NotifyV1Service UpdateService(ctx, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The Twilio-provided string that uniquely identifies the Service resource to update. | 
 **optional** | ***UpdateServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateServiceOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alexaSkillId** | **optional.String**| Deprecated. | 
 **apnCredentialSid** | **optional.String**| The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for APN Bindings. | 
 **defaultAlexaNotificationProtocolVersion** | **optional.String**| Deprecated. | 
 **defaultApnNotificationProtocolVersion** | **optional.String**| The protocol version to use for sending APNS notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource. | 
 **defaultFcmNotificationProtocolVersion** | **optional.String**| The protocol version to use for sending FCM notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource. | 
 **defaultGcmNotificationProtocolVersion** | **optional.String**| The protocol version to use for sending GCM notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource. | 
 **deliveryCallbackEnabled** | **optional.Bool**| Callback configuration that enables delivery callbacks, default false | 
 **deliveryCallbackUrl** | **optional.String**| URL to send delivery status callback. | 
 **facebookMessengerPageId** | **optional.String**| Deprecated. | 
 **fcmCredentialSid** | **optional.String**| The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for FCM Bindings. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It can be up to 64 characters long. | 
 **gcmCredentialSid** | **optional.String**| The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for GCM Bindings. | 
 **logEnabled** | **optional.Bool**| Whether to log notifications. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | 
 **messagingServiceSid** | **optional.String**| The SID of the [Messaging Service](https://www.twilio.com/docs/sms/send-messages#messaging-services) to use for SMS Bindings. This parameter must be set in order to send SMS notifications. | 

### Return type

[**NotifyV1Service**](notify.v1.service.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

