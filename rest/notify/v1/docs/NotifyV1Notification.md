# NotifyV1Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ServiceSid** | **string** | The SID of the Service that the resource is associated with |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |[optional] 
**Identities** | **[]string** | The list of identity values of the Users to notify |[optional] 
**Tags** | **[]string** | The tags that select the Bindings to notify |[optional] 
**Segments** | **[]string** | The list of Segments to notify |[optional] 
**Priority** | Pointer to [**string**](NotificationEnumPriority.md) |  |
**Ttl** | **int** | How long, in seconds, the notification is valid |[optional] 
**Title** | **string** | The notification title |[optional] 
**Body** | **string** | The notification body text |[optional] 
**Sound** | **string** | The name of the sound to be played for the notification |[optional] 
**Action** | **string** | The actions to display for the notification |[optional] 
**Data** | Pointer to **interface{}** | The custom key-value pairs of the notification's payload |
**Apn** | Pointer to **interface{}** | The APNS-specific payload that overrides corresponding attributes in a generic payload for APNS Bindings |
**Gcm** | Pointer to **interface{}** | The GCM-specific payload that overrides corresponding attributes in generic payload for GCM Bindings |
**Fcm** | Pointer to **interface{}** | The FCM-specific payload that overrides corresponding attributes in generic payload for FCM Bindings |
**Sms** | Pointer to **interface{}** | The SMS-specific payload that overrides corresponding attributes in generic payload for SMS Bindings |
**FacebookMessenger** | Pointer to **interface{}** | Deprecated |
**Alexa** | Pointer to **interface{}** | Deprecated |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


