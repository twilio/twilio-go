# ChatV2Member

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ChannelSid** | **string** | The SID of the Channel for the member |[optional] 
**ServiceSid** | **string** | The SID of the Service that the resource is associated with |[optional] 
**Identity** | **string** | The string that identifies the resource's User |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**RoleSid** | **string** | The SID of the Role assigned to the member |[optional] 
**LastConsumedMessageIndex** | Pointer to **int** | The index of the last Message that the Member has read within the Channel |
**LastConsumptionTimestamp** | [**time.Time**](time.Time.md) | The ISO 8601 based timestamp string that represents the datetime of the last Message read event for the Member within the Channel |[optional] 
**Url** | **string** | The absolute URL of the Member resource |[optional] 
**Attributes** | **string** | The JSON string that stores application-specific data |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


