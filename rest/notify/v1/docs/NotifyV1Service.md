# NotifyV1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**ApnCredentialSid** | Pointer to **string** | The SID of the Credential to use for APN Bindings |
**GcmCredentialSid** | Pointer to **string** | The SID of the Credential to use for GCM Bindings |
**FcmCredentialSid** | Pointer to **string** | The SID of the Credential to use for FCM Bindings |
**MessagingServiceSid** | Pointer to **string** | The SID of the Messaging Service to use for SMS Bindings |
**FacebookMessengerPageId** | Pointer to **string** | Deprecated |
**DefaultApnNotificationProtocolVersion** | Pointer to **string** | The protocol version to use for sending APNS notifications |
**DefaultGcmNotificationProtocolVersion** | Pointer to **string** | The protocol version to use for sending GCM notifications |
**DefaultFcmNotificationProtocolVersion** | Pointer to **string** | The protocol version to use for sending FCM notifications |
**LogEnabled** | Pointer to **bool** | Whether to log notifications |
**Url** | Pointer to **string** | The absolute URL of the Service resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of the resources related to the service |
**AlexaSkillId** | Pointer to **string** | Deprecated |
**DefaultAlexaNotificationProtocolVersion** | Pointer to **string** | Deprecated |
**DeliveryCallbackUrl** | Pointer to **string** | Webhook URL |
**DeliveryCallbackEnabled** | Pointer to **bool** | Enable delivery callbacks |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


