# SupersimV1Sim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that the Super SIM belongs to |[optional] 
**Iccid** | **string** | The ICCID associated with the SIM |[optional] 
**Status** | Pointer to [**string**](SimEnumStatus.md) |  |
**FleetSid** | **string** | The unique ID of the Fleet configured for this SIM |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the Sim Resource |[optional] 
**Links** | **map[string]interface{}** |  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


