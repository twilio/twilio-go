# WirelessV1Command

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Command** | Pointer to **string** | The message being sent to or from the SIM |
**CommandMode** | Pointer to **string** | The mode used to send the SMS message |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated format |
**DeliveryReceiptRequested** | Pointer to **bool** | Whether to request a delivery receipt |
**Direction** | Pointer to **string** | The direction of the Command |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**SimSid** | Pointer to **string** | The SID of the Sim resource that the Command was sent to or from |
**Status** | Pointer to **string** | The status of the Command |
**Transport** | Pointer to **string** | The type of transport used |
**Url** | Pointer to **string** | The absolute URL of the resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


