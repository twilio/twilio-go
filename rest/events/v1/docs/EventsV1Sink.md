# EventsV1Sink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this Sink was created, given in ISO 8601 format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this Sink was updated, given in ISO 8601 format. |
**Description** | Pointer to **string** | A human readable description for the Sink |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this Sink. |
**SinkConfiguration** | Pointer to **interface{}** | The information required for Twilio to connect to the provided Sink encoded as JSON. |
**SinkType** | Pointer to [**string**](SinkEnumSinkType.md) |  |
**Status** | Pointer to [**string**](SinkEnumStatus.md) |  |
**Url** | Pointer to **string** | The URL of this resource. |
**Links** | Pointer to **map[string]interface{}** | Contains a dictionary of URL links to nested resources of this Sink. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


