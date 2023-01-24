# ConversationsV1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique ID of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this service. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |
**FriendlyName** | Pointer to **string** | The human-readable name of this service, limited to 256 characters. Optional. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was last updated. |
**Url** | Pointer to **string** | An absolute API resource URL for this service. |
**Links** | Pointer to **map[string]interface{}** | Contains absolute API resource URLs to access conversations, users, roles, bindings and configuration of this service. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


