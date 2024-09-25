# AssistantsV1ToolWithPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Tool resource. |[optional] 
**Description** | **string** | The description of the tool. |
**Enabled** | **bool** | True if the tool is enabled. |
**Id** | **string** | The tool ID. |
**Meta** | **map[string]interface{}** | The metadata related to method, url, input_schema to used with the Tool. |
**Name** | **string** | The name of the tool. |
**RequiresAuth** | **bool** | The authentication requirement for the tool. |
**Type** | **string** | The type of the tool. ('WEBHOOK') |
**Url** | **string** | The url of the tool resource. |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date and time in GMT when the Tool was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | [**time.Time**](time.Time.md) | The date and time in GMT when the Tool was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Policies** | [**[]AssistantsV1Policy**](AssistantsV1Policy.md) | The Policies associated with the tool. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


