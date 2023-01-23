# SupersimV1EsimProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account to which the eSIM Profile resource belongs |[optional] 
**Iccid** | **string** | The ICCID associated with the Sim resource |[optional] 
**SimSid** | **string** | The SID of the Sim resource that this eSIM Profile controls |[optional] 
**Status** | Pointer to [**string**](EsimProfileEnumStatus.md) |  |
**Eid** | **string** | Identifier of the eUICC that can claim the eSIM Profile |[optional] 
**SmdpPlusAddress** | **string** | Address of the SM-DP+ server from which the Profile will be downloaded |[optional] 
**ErrorCode** | **string** | Code indicating the failure if the download of the SIM Profile failed and the eSIM Profile is in `failed` state |[optional] 
**ErrorMessage** | **string** | Error message describing the failure if the download of the SIM Profile failed and the eSIM Profile is in `failed` state |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the eSIM Profile resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


