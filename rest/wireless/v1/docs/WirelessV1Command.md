# WirelessV1Command

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**SimSid** | Pointer to **string** | The SID of the Sim resource that the Command was sent to or from |
**Command** | Pointer to **string** | The message being sent to or from the SIM |
**CommandMode** | Pointer to [**string**](CommandEnumCommandMode.md) |  |
**Transport** | Pointer to [**string**](CommandEnumTransport.md) |  |
**DeliveryReceiptRequested** | Pointer to **bool** | Whether to request a delivery receipt |
**Status** | Pointer to [**string**](CommandEnumStatus.md) |  |
**Direction** | Pointer to [**string**](CommandEnumDirection.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated format |
**Url** | Pointer to **string** | The absolute URL of the resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


