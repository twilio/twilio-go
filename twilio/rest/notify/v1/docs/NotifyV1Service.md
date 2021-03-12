# NotifyV1Service

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**AlexaSkillId** | Pointer to **string** | Deprecated |
**ApnCredentialSid** | Pointer to **string** | The SID of the Credential to use for APN Bindings |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**DefaultAlexaNotificationProtocolVersion** | Pointer to **string** | Deprecated |
**DefaultApnNotificationProtocolVersion** | Pointer to **string** | The protocol version to use for sending APNS notifications |
**DefaultFcmNotificationProtocolVersion** | Pointer to **string** | The protocol version to use for sending FCM notifications |
**DefaultGcmNotificationProtocolVersion** | Pointer to **string** | The protocol version to use for sending GCM notifications |
**DeliveryCallbackEnabled** | Pointer to **bool** | Enable delivery callbacks |
**DeliveryCallbackUrl** | Pointer to **string** | Webhook URL |
**FacebookMessengerPageId** | Pointer to **string** | Deprecated |
**FcmCredentialSid** | Pointer to **string** | The SID of the Credential to use for FCM Bindings |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource |
**GcmCredentialSid** | Pointer to **string** | The SID of the Credential to use for GCM Bindings |
**Links** | Pointer to **map[string]interface{}** | The URLs of the resources related to the service |
**LogEnabled** | Pointer to **bool** | Whether to log notifications |
**MessagingServiceSid** | Pointer to **string** | The SID of the Messaging Service to use for SMS Bindings |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Url** | Pointer to **string** | The absolute URL of the Service resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


