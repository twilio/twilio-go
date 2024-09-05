# IntelligenceV2Operator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account the Operator belongs to. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this Operator. |
**FriendlyName** | Pointer to **string** | A human-readable name of this resource, up to 64 characters. |
**Description** | Pointer to **string** | A human-readable description of this resource, longer than the friendly name. |
**Author** | Pointer to **string** | The creator of the Operator. Either Twilio or the creating Account. |
**OperatorType** | Pointer to **string** | Operator Type for this Operator. References an existing Operator Type resource. |
**Version** | **int** | Numeric Operator version. Incremented with each update on the resource, used to ensure integrity when updating the Operator. |[optional] [default to 0]
**Availability** | Pointer to [**string**](OperatorEnumAvailability.md) |  |
**Config** | Pointer to **interface{}** | Operator configuration, following the schema defined by the Operator Type. Only available on Custom Operators created by the Account. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this Operator was created, given in ISO 8601 format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this Operator was updated, given in ISO 8601 format. |
**Url** | Pointer to **string** | The URL of this resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


