# ChatV1Member

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ChannelSid** | Pointer to **string** | The unique ID of the Channel for the member |
**ServiceSid** | Pointer to **string** | The SID of the Service that the resource is associated with |
**Identity** | Pointer to **string** | The string that identifies the resource's User |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT when the resource was last updated |
**RoleSid** | Pointer to **string** | The SID of the Role assigned to the member |
**LastConsumedMessageIndex** | Pointer to **int** | The index of the last Message that the Member has read within the Channel |
**LastConsumptionTimestamp** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 based timestamp string that represents the date-time of the last Message read event for the Member within the Channel |
**Url** | Pointer to **string** | The absolute URL of the Member resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


