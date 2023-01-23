# SupersimV1SmsCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the SMS Command resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SMS Command resource. |
**SimSid** | Pointer to **string** | The SID of the [SIM](https://www.twilio.com/docs/iot/supersim/api/sim-resource) that this SMS Command was sent to or from. |
**Payload** | Pointer to **string** | The message body of the SMS Command sent to or from the SIM. For text mode messages, this can be up to 160 characters. |
**Status** | Pointer to [**string**](SmsCommandEnumStatus.md) |  |
**Direction** | Pointer to [**string**](SmsCommandEnumDirection.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the SMS Command resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


