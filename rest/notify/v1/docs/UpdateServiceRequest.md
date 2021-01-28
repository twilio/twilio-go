# UpdateServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlexaSkillId** | **string** | Deprecated. | [optional] 
**ApnCredentialSid** | **string** | The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for APN Bindings. | [optional] 
**DefaultAlexaNotificationProtocolVersion** | **string** | Deprecated. | [optional] 
**DefaultApnNotificationProtocolVersion** | **string** | The protocol version to use for sending APNS notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource. | [optional] 
**DefaultFcmNotificationProtocolVersion** | **string** | The protocol version to use for sending FCM notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource. | [optional] 
**DefaultGcmNotificationProtocolVersion** | **string** | The protocol version to use for sending GCM notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource. | [optional] 
**DeliveryCallbackEnabled** | **bool** | Callback configuration that enables delivery callbacks, default false | [optional] 
**DeliveryCallbackUrl** | **string** | URL to send delivery status callback. | [optional] 
**FacebookMessengerPageId** | **string** | Deprecated. | [optional] 
**FcmCredentialSid** | **string** | The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for FCM Bindings. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 64 characters long. | [optional] 
**GcmCredentialSid** | **string** | The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for GCM Bindings. | [optional] 
**LogEnabled** | **bool** | Whether to log notifications. Can be: &#x60;true&#x60; or &#x60;false&#x60; and the default is &#x60;true&#x60;. | [optional] 
**MessagingServiceSid** | **string** | The SID of the [Messaging Service](https://www.twilio.com/docs/sms/send-messages#messaging-services) to use for SMS Bindings. This parameter must be set in order to send SMS notifications. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


