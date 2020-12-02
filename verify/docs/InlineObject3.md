# InlineObject3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | **string** | Details provided to give context about the Challenge. Shown to the end user. It must be a stringified JSON with the following structure: {\&quot;message\&quot;: \&quot;string\&quot;, \&quot;fields\&quot;: [ { \&quot;label\&quot;: \&quot;string\&quot;, \&quot;value\&quot;: \&quot;string\&quot;}]}. &#x60;message&#x60; is required. If you send the &#x60;fields&#x60; property, each field has to include &#x60;label&#x60; and &#x60;value&#x60; properties. If you had set &#x60;include_date&#x3D;true&#x60; in the &#x60;push&#x60; configuration of the [service](https://www.twilio.com/docs/verify/api/service), the response will also include the challenge&#39;s date created value as an additional field called &#x60;date&#x60; | [optional] 
**ExpirationDate** | [**time.Time**](time.Time.md) | The future date in which this Challenge will expire, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. | [optional] 
**FactorSid** | **string** | The unique SID identifier of the Factor. | 
**HiddenDetails** | **string** | Details provided to give context about the Challenge. Not shown to the end user. It must be a stringified JSON with only strings values eg. &#x60;{\&quot;ip\&quot;: \&quot;172.168.1.234\&quot;}&#x60; | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


