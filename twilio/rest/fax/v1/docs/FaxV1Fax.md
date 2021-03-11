# FaxV1Fax

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ApiVersion** | Pointer to **string** | The API version used to transmit the fax |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 formatted date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 formatted date and time in GMT when the resource was last updated |
**Direction** | Pointer to **string** | The direction of the fax |
**Duration** | Pointer to **int32** | The time it took to transmit the fax |
**From** | Pointer to **string** | The number the fax was sent from |
**Links** | Pointer to **map[string]interface{}** | The URLs of the fax's related resources |
**MediaSid** | Pointer to **string** | The SID of the FaxMedia resource that is associated with the Fax |
**MediaUrl** | Pointer to **string** | The Twilio-hosted URL that can be used to download fax media |
**NumPages** | Pointer to **int32** | The number of pages contained in the fax document |
**Price** | Pointer to **float32** | The fax transmission price |
**PriceUnit** | Pointer to **string** | The ISO 4217 currency used for billing |
**Quality** | Pointer to **string** | The quality of the fax |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Status** | Pointer to **string** | The status of the fax |
**To** | Pointer to **string** | The phone number that received the fax |
**Url** | Pointer to **string** | The absolute URL of the fax resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


