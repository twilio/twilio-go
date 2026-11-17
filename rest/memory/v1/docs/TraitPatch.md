# TraitPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the trait. |
**TraitGroup** | **string** | The trait group name to which this trait belongs. |
**Value** | Pointer to **interface{}** | The value of a trait. Can be a string, integer, boolean, or an array of these types (arrays cannot contain nested arrays). |
**Source** | **string** | Caller-set, descriptive, untrusted provenance metadata. Must follow the {company}/{product}/{version}/_* convention. \"twilio\" is a reserved company name, only usable by authenticated internal callers. |[optional] 
**OccurredAt** | [**time.Time**](time.Time.md) | When the value was actually observed or produced, as distinct from when the write was received. Optional; defaults to write time if omitted. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


