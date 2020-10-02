# InlineObject5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binding** | **string** | A unique binding for this Factor that identifies it. E.g. the algorithm and public key for &#x60;push&#x60; factors. It must be a json string with the required properties for the given factor type. Required when creating a new Factor. This value is never returned because it can contain customer secrets. | 
**Config** | **string** | The config required for this Factor. It must be a json string with the required properties for the given factor type | 
**FactorType** | **string** | The Type of this Factor. Currently only &#x60;push&#x60; is supported | 
**FriendlyName** | **string** | The friendly name of this Factor | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


