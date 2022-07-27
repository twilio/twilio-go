# SupersimV1IpCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**SimSid** | Pointer to **string** | The SID of the Super SIM that this IP Command was sent to or from |
**SimIccid** | Pointer to **string** | The ICCID of the Super SIM that this IP Command was sent to or from |
**Status** | Pointer to [**string**](IpCommandEnumStatus.md) |  |
**Direction** | Pointer to [**string**](IpCommandEnumDirection.md) |  |
**DeviceIp** | Pointer to **string** | The IP address of the device that the IP Command was sent to or received from |
**DevicePort** | Pointer to **int** | The port that the IP Command either originated from or was sent to |
**PayloadType** | Pointer to [**string**](IpCommandEnumPayloadType.md) |  |
**Payload** | Pointer to **string** | The payload of the IP Command sent to or from the Super SIM |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Url** | Pointer to **string** | The absolute URL of the IP Command resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


