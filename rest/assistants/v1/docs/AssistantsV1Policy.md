# AssistantsV1Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Policy ID. |[optional] 
**Name** | **string** | The name of the policy. |[optional] 
**Description** | **string** | The description of the policy. |[optional] 
**AccountSid** | **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Policy resource. |[optional] 
**UserSid** | **string** | The SID of the User that created the Policy resource. |[optional] 
**Type** | **string** | The type of the policy. |
**PolicyDetails** | **map[string]interface{}** | The details of the policy based on the type. |
**DateCreated** | [**time.Time**](time.Time.md) | The date and time in GMT when the Policy was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date and time in GMT when the Policy was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


