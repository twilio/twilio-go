# SupersimV1IpCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**SimSid** | **string** | The SID of the Super SIM that this IP Command was sent to or from |[optional] 
**SimIccid** | **string** | The ICCID of the Super SIM that this IP Command was sent to or from |[optional] 
**Status** | Pointer to [**string**](IpCommandEnumStatus.md) |  |
**Direction** | Pointer to [**string**](IpCommandEnumDirection.md) |  |
**DeviceIp** | **string** | The IP address of the device that the IP Command was sent to or received from |[optional] 
**DevicePort** | **int** | The port that the IP Command either originated from or was sent to |[optional] 
**PayloadType** | Pointer to [**string**](IpCommandEnumPayloadType.md) |  |
**Payload** | **string** | The payload of the IP Command sent to or from the Super SIM |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the IP Command resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


