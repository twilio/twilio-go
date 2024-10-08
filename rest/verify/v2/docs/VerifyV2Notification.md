# VerifyV2Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this Notification. |
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**ServiceSid** | Pointer to **string** | The unique SID identifier of the Service. |
**EntitySid** | Pointer to **string** | The unique SID identifier of the Entity. |
**Identity** | Pointer to **string** | Customer unique identity for the Entity owner of the Challenge. This identifier should be immutable, not PII, length between 8 and 64 characters, and generated by your external system, such as your user's UUID, GUID, or SID. It can only contain dash (-) separated alphanumeric characters. |
**ChallengeSid** | Pointer to **string** | The unique SID identifier of the Challenge. |
**Priority** | Pointer to **string** | The priority of the notification. For `push` Challenges it's always `high` which sends the notification immediately, and can wake up a sleeping device. |
**Ttl** | **int** | How long, in seconds, the notification is valid. Max: 5 minutes |[optional] [default to 0]
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this Notification was created, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


