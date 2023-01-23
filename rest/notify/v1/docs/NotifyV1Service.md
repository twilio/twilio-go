# NotifyV1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**FriendlyName** | **string** | The string that you assigned to describe the resource |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |[optional] 
**ApnCredentialSid** | **string** | The SID of the Credential to use for APN Bindings |[optional] 
**GcmCredentialSid** | **string** | The SID of the Credential to use for GCM Bindings |[optional] 
**FcmCredentialSid** | **string** | The SID of the Credential to use for FCM Bindings |[optional] 
**MessagingServiceSid** | **string** | The SID of the Messaging Service to use for SMS Bindings |[optional] 
**FacebookMessengerPageId** | **string** | Deprecated |[optional] 
**DefaultApnNotificationProtocolVersion** | **string** | The protocol version to use for sending APNS notifications |[optional] 
**DefaultGcmNotificationProtocolVersion** | **string** | The protocol version to use for sending GCM notifications |[optional] 
**DefaultFcmNotificationProtocolVersion** | **string** | The protocol version to use for sending FCM notifications |[optional] 
**LogEnabled** | **bool** | Whether to log notifications |[optional] 
**Url** | **string** | The absolute URL of the Service resource |[optional] 
**Links** | **map[string]interface{}** | The URLs of the resources related to the service |[optional] 
**AlexaSkillId** | **string** | Deprecated |[optional] 
**DefaultAlexaNotificationProtocolVersion** | **string** | Deprecated |[optional] 
**DeliveryCallbackUrl** | **string** | Webhook URL |[optional] 
**DeliveryCallbackEnabled** | **bool** | Enable delivery callbacks |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


