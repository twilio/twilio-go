# NumbersV1CreateEmbeddedRegistrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Registration identifier (BU-prefixed). |
**RegulationId** | **string** | The regulation ID for this registration. |
**RegulationVersion** | **int** | The regulation version. |
**FriendlyName** | **string** | The friendly name provided in the request. |
**Status** | **string** | Registration status. Always DRAFT on creation. |
**StatusNotificationEmail** | Pointer to **string** | Email address for status notifications. |
**StatusCallbackUrl** | Pointer to **string** | Callback URL for status webhooks. |
**Comments** | Pointer to **string** | Additional comments. |
**EmbeddedSession** | [**NumbersV1EmbeddedSession**](NumbersV1EmbeddedSession.md) |  |
**Data** | **map[string]interface{}** | Registration data echoed from the request. |
**DateCreated** | [**time.Time**](time.Time.md) | Timestamp of creation. |
**DateUpdated** | [**time.Time**](time.Time.md) | Timestamp of last update. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


