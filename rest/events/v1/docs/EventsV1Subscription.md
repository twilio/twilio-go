# EventsV1Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this Subscription. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this Subscription was created, given in ISO 8601 format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this Subscription was updated, given in ISO 8601 format. |
**Description** | Pointer to **string** | A human readable description for the Subscription |
**SinkSid** | Pointer to **string** | The SID of the sink that events selected by this subscription should be sent to. Sink must be active for the subscription to be created. |
**Url** | Pointer to **string** | The URL of this resource. |
**Links** | Pointer to **map[string]interface{}** | Contains a dictionary of URL links to nested resources of this Subscription. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


