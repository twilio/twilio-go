# WirelessV1Command

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**SimSid** | **string** | The SID of the Sim resource that the Command was sent to or from |[optional] 
**Command** | **string** | The message being sent to or from the SIM |[optional] 
**CommandMode** | Pointer to [**string**](CommandEnumCommandMode.md) |  |
**Transport** | Pointer to [**string**](CommandEnumTransport.md) |  |
**DeliveryReceiptRequested** | **bool** | Whether to request a delivery receipt |[optional] 
**Status** | Pointer to [**string**](CommandEnumStatus.md) |  |
**Direction** | Pointer to [**string**](CommandEnumDirection.md) |  |
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated format |[optional] 
**Url** | **string** | The absolute URL of the resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


