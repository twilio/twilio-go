# ServicesNotificationsApi

All URIs are relative to *https://notify.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotification**](ServicesNotificationsApi.md#CreateNotification) | **Post** /v1/Services/{ServiceSid}/Notifications | 



## CreateNotification

> NotifyV1ServiceNotification CreateNotification(ctx, ServiceSidoptional)



### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ServiceSid** | **string** | The SID of the [Service](https://www.twilio.com/docs/notify/api/service-resource) to create the resource under.

### Other Parameters

Other parameters are passed through a pointer to a CreateNotificationParams struct


Name | Type | Description
------------- | ------------- | -------------
**Action** | **string** | The actions to display for the notification. For APNS, translates to the &#x60;aps.category&#x60; value. For GCM, translates to the &#x60;data.twi_action&#x60; value. For SMS, this parameter is not supported and is omitted from deliveries to those channels.
**Alexa** | [**map[string]interface{}**](map[string]interface{}.md) | Deprecated.
**Apn** | [**map[string]interface{}**](map[string]interface{}.md) | The APNS-specific payload that overrides corresponding attributes in the generic payload for APNS Bindings. This property maps to the APNS &#x60;Payload&#x60; item, therefore the &#x60;aps&#x60; key must be used to change standard attributes. Adds custom key-value pairs to the root of the dictionary. See the [APNS documentation](https://developer.apple.com/library/content/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/CommunicatingwithAPNs.html) for more details. We reserve keys that start with &#x60;twi_&#x60; for future use. Custom keys that start with &#x60;twi_&#x60; are not allowed.
**Body** | **string** | The notification text. For FCM and GCM, translates to &#x60;data.twi_body&#x60;. For APNS, translates to &#x60;aps.alert.body&#x60;. For SMS, translates to &#x60;body&#x60;. SMS requires either this &#x60;body&#x60; value, or &#x60;media_urls&#x60; attribute defined in the &#x60;sms&#x60; parameter of the notification.
**Data** | [**map[string]interface{}**](map[string]interface{}.md) | The custom key-value pairs of the notification&#39;s payload. For FCM and GCM, this value translates to &#x60;data&#x60; in the FCM and GCM payloads. FCM and GCM [reserve certain keys](https://firebase.google.com/docs/cloud-messaging/http-server-ref) that cannot be used in those channels. For APNS, attributes of &#x60;data&#x60; are inserted into the APNS payload as custom properties outside of the &#x60;aps&#x60; dictionary. In all channels, we reserve keys that start with &#x60;twi_&#x60; for future use. Custom keys that start with &#x60;twi_&#x60; are not allowed and are rejected as 400 Bad request with no delivery attempted. For SMS, this parameter is not supported and is omitted from deliveries to those channels.
**DeliveryCallbackUrl** | **string** | URL to send webhooks.
**FacebookMessenger** | [**map[string]interface{}**](map[string]interface{}.md) | Deprecated.
**Fcm** | [**map[string]interface{}**](map[string]interface{}.md) | The FCM-specific payload that overrides corresponding attributes in the generic payload for FCM Bindings. This property maps to the root JSON dictionary. See the [FCM documentation](https://firebase.google.com/docs/cloud-messaging/http-server-ref#downstream) for more details. Target parameters &#x60;to&#x60;, &#x60;registration_ids&#x60;, &#x60;condition&#x60;, and &#x60;notification_key&#x60; are not allowed in this parameter. We reserve keys that start with &#x60;twi_&#x60; for future use. Custom keys that start with &#x60;twi_&#x60; are not allowed. FCM also [reserves certain keys](https://firebase.google.com/docs/cloud-messaging/http-server-ref), which cannot be used in that channel.
**Gcm** | [**map[string]interface{}**](map[string]interface{}.md) | The GCM-specific payload that overrides corresponding attributes in the generic payload for GCM Bindings.  This property maps to the root JSON dictionary. See the [GCM documentation](https://firebase.google.com/docs/cloud-messaging/http-server-ref) for more details. Target parameters &#x60;to&#x60;, &#x60;registration_ids&#x60;, and &#x60;notification_key&#x60; are not allowed. We reserve keys that start with &#x60;twi_&#x60; for future use. Custom keys that start with &#x60;twi_&#x60; are not allowed. GCM also [reserves certain keys](https://firebase.google.com/docs/cloud-messaging/http-server-ref).
**Identity** | **[]string** | The &#x60;identity&#x60; value that uniquely identifies the new resource&#39;s [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/notify/api/service-resource). Delivery will be attempted only to Bindings with an Identity in this list. No more than 20 items are allowed in this list.
**Priority** | **string** | The priority of the notification. Can be: &#x60;low&#x60; or &#x60;high&#x60; and the default is &#x60;high&#x60;. A value of &#x60;low&#x60; optimizes the client app&#39;s battery consumption; however, notifications may be delivered with unspecified delay. For FCM and GCM, &#x60;low&#x60; priority is the same as &#x60;Normal&#x60; priority. For APNS &#x60;low&#x60; priority is the same as &#x60;5&#x60;. A value of &#x60;high&#x60; sends the notification immediately, and can wake up a sleeping device. For FCM and GCM, &#x60;high&#x60; is the same as &#x60;High&#x60; priority. For APNS, &#x60;high&#x60; is a priority &#x60;10&#x60;. SMS does not support this property.
**Segment** | **[]string** | The Segment resource is deprecated. Use the &#x60;tag&#x60; parameter, instead.
**Sms** | [**map[string]interface{}**](map[string]interface{}.md) | The SMS-specific payload that overrides corresponding attributes in the generic payload for SMS Bindings.  Each attribute in this value maps to the corresponding &#x60;form&#x60; parameter of the Twilio [Message](https://www.twilio.com/docs/sms/send-messages) resource.  These parameters of the Message resource are supported in snake case format: &#x60;body&#x60;, &#x60;media_urls&#x60;, &#x60;status_callback&#x60;, and &#x60;max_price&#x60;.  The &#x60;status_callback&#x60; parameter overrides the corresponding parameter in the messaging service, if configured. The &#x60;media_urls&#x60; property expects a JSON array.
**Sound** | **string** | The name of the sound to be played for the notification. For FCM and GCM, this Translates to &#x60;data.twi_sound&#x60;.  For APNS, this translates to &#x60;aps.sound&#x60;.  SMS does not support this property.
**Tag** | **[]string** | A tag that selects the Bindings to notify. Repeat this parameter to specify more than one tag, up to a total of 5 tags. The implicit tag &#x60;all&#x60; is available to notify all Bindings in a Service instance. Similarly, the implicit tags &#x60;apn&#x60;, &#x60;fcm&#x60;, &#x60;gcm&#x60;, &#x60;sms&#x60; and &#x60;facebook-messenger&#x60; are available to notify all Bindings in a specific channel.
**Title** | **string** | The notification title. For FCM and GCM, this translates to the &#x60;data.twi_title&#x60; value. For APNS, this translates to the &#x60;aps.alert.title&#x60; value. SMS does not support this property. This field is not visible on iOS phones and tablets but appears on Apple Watch and Android devices.
**ToBinding** | **[]string** | The destination address specified as a JSON string.  Multiple &#x60;to_binding&#x60; parameters can be included but the total size of the request entity should not exceed 1MB. This is typically sufficient for 10,000 phone numbers.
**Ttl** | **int** | How long, in seconds, the notification is valid. Can be an integer between 0 and 2,419,200, which is 4 weeks, the default and the maximum supported time to live (TTL). Delivery should be attempted if the device is offline until the TTL elapses. Zero means that the notification delivery is attempted immediately, only once, and is not stored for future delivery. SMS does not support this property.

### Return type

[**NotifyV1ServiceNotification**](NotifyV1ServiceNotification.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

