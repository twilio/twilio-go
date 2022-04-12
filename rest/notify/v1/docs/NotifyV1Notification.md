# NotifyV1Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Action** | Pointer to **string** | The actions to display for the notification |
**Alexa** | Pointer to **interface{}** | Deprecated |
**Apn** | Pointer to **interface{}** | The APNS-specific payload that overrides corresponding attributes in a generic payload for APNS Bindings |
**Body** | Pointer to **string** | The notification body text |
**Data** | Pointer to **interface{}** | The custom key-value pairs of the notification's payload |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**FacebookMessenger** | Pointer to **interface{}** | Deprecated |
**Fcm** | Pointer to **interface{}** | The FCM-specific payload that overrides corresponding attributes in generic payload for FCM Bindings |
**Gcm** | Pointer to **interface{}** | The GCM-specific payload that overrides corresponding attributes in generic payload for GCM Bindings |
**Identities** | Pointer to **[]string** | The list of identity values of the Users to notify |
**Priority** | Pointer to **string** | The priority of the notification |
**Segments** | Pointer to **[]string** | The list of Segments to notify |
**ServiceSid** | Pointer to **string** | The SID of the Service that the resource is associated with |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Sms** | Pointer to **interface{}** | The SMS-specific payload that overrides corresponding attributes in generic payload for SMS Bindings |
**Sound** | Pointer to **string** | The name of the sound to be played for the notification |
**Tags** | Pointer to **[]string** | The tags that select the Bindings to notify |
**Title** | Pointer to **string** | The notification title |
**Ttl** | Pointer to **int** | How long, in seconds, the notification is valid |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


