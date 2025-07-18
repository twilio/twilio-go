# ConversationsV1ConfigurationAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |
**AccountSid** | Pointer to **string** | The unique ID of the [Account](https://www.twilio.com/docs/iam/api/account) the address belongs to |
**Type** | Pointer to **string** | Type of Address, value can be `whatsapp` or `sms`. |
**Address** | Pointer to **string** | The unique address to be configured. The address can be a whatsapp address or phone number |
**FriendlyName** | Pointer to **string** | The human-readable name of this configuration, limited to 256 characters. Optional. |
**AutoCreation** | Pointer to **interface{}** | Auto Creation configuration for the address. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was created. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this resource was last updated. |
**Url** | Pointer to **string** | An absolute API resource URL for this address configuration. |
**AddressCountry** | Pointer to **string** | An ISO 3166-1 alpha-2n country code which the address belongs to. This is currently only applicable to short code addresses. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


