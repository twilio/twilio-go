# IntelligenceV2OperatorType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique name that references an Operator's Operator Type. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this Operator Type. |
**FriendlyName** | Pointer to **string** | A human-readable name of this resource, up to 64 characters. |
**Description** | Pointer to **string** | A human-readable description of this resource, longer than the friendly name. |
**DocsLink** | Pointer to **string** | Additional documentation for the Operator Type. |
**OutputType** | Pointer to [**string**](OperatorTypeEnumOutputType.md) |  |
**SupportedLanguages** | Pointer to **[]string** | List of languages this Operator Type supports. |
**Provider** | Pointer to [**string**](OperatorTypeEnumProvider.md) |  |
**Availability** | Pointer to [**string**](OperatorTypeEnumAvailability.md) |  |
**Configurable** | Pointer to **bool** | Operators can be created from configurable Operator Types. |
**ConfigSchema** | Pointer to **interface{}** | JSON Schema for configuring an Operator with this Operator Type. Following https://json-schema.org/ |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this Operator Type was created, given in ISO 8601 format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this Operator Type was updated, given in ISO 8601 format. |
**Url** | Pointer to **string** | The URL of this resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


