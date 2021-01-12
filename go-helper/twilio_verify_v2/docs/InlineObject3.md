# InlineObject3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetailsFields** | **[]map[string]interface{}** | A list of objects that describe the Fields included in the Challenge. Each object contains the label and value of the field. Used when &#x60;factor_type&#x60; is &#x60;push&#x60;. | [optional] 
**DetailsMessage** | **string** | Shown to the user when the push notification arrives. Required when &#x60;factor_type&#x60; is &#x60;push&#x60; | [optional] 
**ExpirationDate** | [**time.Time**](time.Time.md) | The date-time when this Challenge expires, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. The default value is five (5) minutes after Challenge creation. The max value is sixty (60) minutes after creation. | [optional] 
**FactorSid** | **string** | The unique SID identifier of the Factor. | 
**HiddenDetails** | [**map[string]interface{}**](.md) | Details provided to give context about the Challenge. Not shown to the end user. It must be a stringified JSON with only strings values eg. &#x60;{\&quot;ip\&quot;: \&quot;172.168.1.234\&quot;}&#x60; | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


