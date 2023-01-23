# SupersimV1SmsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**SimSid** | **string** | The SID of the SIM that this SMS Command was sent to or from |[optional] 
**Payload** | **string** | The message body of the SMS Command sent to or from the SIM |[optional] 
**Status** | Pointer to [**string**](SmsCommandEnumStatus.md) |  |
**Direction** | Pointer to [**string**](SmsCommandEnumDirection.md) |  |
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the SMS Command resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


