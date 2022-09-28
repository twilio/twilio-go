# ConversationsV1Credential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |
**AccountSid** | Pointer to **string** | The unique ID of the Account responsible for this credential. |
**FriendlyName** | Pointer to **string** | The human-readable name of this credential. |
**Type** | Pointer to [**string**](CredentialEnumPushType.md) |  |
**Sandbox** | Pointer to **string** | [APN only] Whether to send the credential to sandbox APNs. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was last updated. |
**Url** | Pointer to **string** | An absolute URL for this credential. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


