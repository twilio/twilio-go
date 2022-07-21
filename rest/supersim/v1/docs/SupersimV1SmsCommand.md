# SupersimV1SmsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**SimSid** | Pointer to **string** | The SID of the SIM that this SMS Command was sent to or from |
**Payload** | Pointer to **string** | The message body of the SMS Command sent to or from the SIM |
**Status** | Pointer to [**string**](SmsCommandEnumStatus.md) |  |
**Direction** | Pointer to [**string**](SmsCommandEnumDirection.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Url** | Pointer to **string** | The absolute URL of the SMS Command resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


