# EventsV1Schema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the schema. Each schema can have multiple versions, that share the same id. |
**Url** | Pointer to **string** | The URL of this resource. |
**Links** | Pointer to **map[string]interface{}** | Contains a dictionary of URL links to nested resources of this schema. |
**LatestVersionDateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that the latest schema version was created, given in ISO 8601 format. |
**LatestVersion** | Pointer to **int** | The latest version published of this schema. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


