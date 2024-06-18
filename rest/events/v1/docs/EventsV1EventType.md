# EventsV1EventType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A string that uniquely identifies this Event Type. |
**SchemaId** | Pointer to **string** | A string that uniquely identifies the Schema this Event Type adheres to. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this Event Type was created, given in ISO 8601 format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this Event Type was updated, given in ISO 8601 format. |
**Description** | Pointer to **string** | A human readable description for this Event Type. |
**Status** | Pointer to **string** | A string that describes how this Event Type can be used. For example: `available`, `deprecated`, `restricted`, `discontinued`. When the status is `available`, the Event Type can be used normally. |
**DocumentationUrl** | Pointer to **string** | The URL to the documentation or to the most relevant Twilio Changelog entry of this Event Type. |
**Url** | Pointer to **string** | The URL of this resource. |
**Links** | Pointer to **map[string]interface{}** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


