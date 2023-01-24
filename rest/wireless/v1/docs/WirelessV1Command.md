# WirelessV1Command

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Command resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Command resource. |
**SimSid** | Pointer to **string** | The SID of the [Sim resource](https://www.twilio.com/docs/wireless/api/sim-resource) that the Command was sent to or from. |
**Command** | Pointer to **string** | The message being sent to or from the SIM. For text mode messages, this can be up to 160 characters. For binary mode messages, this is a series of up to 140 bytes of data encoded using base64. |
**CommandMode** | Pointer to [**string**](CommandEnumCommandMode.md) |  |
**Transport** | Pointer to [**string**](CommandEnumTransport.md) |  |
**DeliveryReceiptRequested** | Pointer to **bool** | Whether to request a delivery receipt. |
**Status** | Pointer to [**string**](CommandEnumStatus.md) |  |
**Direction** | Pointer to [**string**](CommandEnumDirection.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format. |
**Url** | Pointer to **string** | The absolute URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


