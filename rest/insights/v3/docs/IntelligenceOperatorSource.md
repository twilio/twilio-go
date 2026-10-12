# IntelligenceOperatorSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperatorId** | **string** | The ID of the Conversational Intelligence operator whose output supplies this field's value. Must match `^intelligence_operator_[0-7][0-9a-z]{25}$`. Cannot be changed after registration.  |
**OperatorName** | **string** | Optional, display-only name of the operator. Stored as supplied and not validated against the operator.  |[optional] 
**Field** | **string** | Path to the property in the operator's output schema that supplies this field's value. Rules: - one or two segments separated by `.` — a top-level property   (`leadScore`) or one level of nesting (`scores.leadScore`); - each segment contains only letters, digits and `_`; array   indexing (such as `items[0]`) is not supported; - the path must exist in the output schema of the operator's   latest active version and resolve to a property of type `number`   or `integer`.  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


