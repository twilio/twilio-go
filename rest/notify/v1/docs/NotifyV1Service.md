# NotifyV1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Service resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Service resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**ApnCredentialSid** | Pointer to **string** | The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for APN Bindings. |
**GcmCredentialSid** | Pointer to **string** | The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for GCM Bindings. |
**FcmCredentialSid** | Pointer to **string** | The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for FCM Bindings. |
**MessagingServiceSid** | Pointer to **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/send-messages#messaging-services) to use for SMS Bindings. In order to send SMS notifications this parameter has to be set. |
**FacebookMessengerPageId** | Pointer to **string** | Deprecated. |
**DefaultApnNotificationProtocolVersion** | Pointer to **string** | The protocol version to use for sending APNS notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource. |
**DefaultGcmNotificationProtocolVersion** | Pointer to **string** | The protocol version to use for sending GCM notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource. |
**DefaultFcmNotificationProtocolVersion** | Pointer to **string** | The protocol version to use for sending FCM notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource. |
**LogEnabled** | Pointer to **bool** | Whether to log notifications. Can be: `true` or `false` and the default is `true`. |
**Url** | Pointer to **string** | The absolute URL of the Service resource. |
**Links** | Pointer to **map[string]interface{}** | The URLs of the Binding, Notification, Segment, and User resources related to the service. |
**AlexaSkillId** | Pointer to **string** | Deprecated. |
**DefaultAlexaNotificationProtocolVersion** | Pointer to **string** | Deprecated. |
**DeliveryCallbackUrl** | Pointer to **string** | URL to send delivery status callback. |
**DeliveryCallbackEnabled** | Pointer to **bool** | Callback configuration that enables delivery callbacks, default false |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


