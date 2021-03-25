# SupersimV1Fleet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**CommandsEnabled** | Pointer to **bool** | Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands |
**CommandsMethod** | Pointer to **string** | A string representing the HTTP method to use when making a request to `commands_url` |
**CommandsUrl** | Pointer to **string** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the Commands number |
**DataEnabled** | Pointer to **bool** | Defines whether SIMs in the Fleet are capable of using data connectivity |
**DataLimit** | Pointer to **int32** | The total data usage (download and upload combined) in Megabytes that each Sim resource assigned to the Fleet resource can consume |
**DataMetering** | Pointer to **string** | The model by which a SIM is metered and billed |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**NetworkAccessProfileSid** | Pointer to **string** | The SID of the Network Access Profile of the Fleet |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**SmsCommandsEnabled** | Pointer to **bool** | Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands |
**SmsCommandsMethod** | Pointer to **string** | A string representing the HTTP method to use when making a request to `sms_commands_url` |
**SmsCommandsUrl** | Pointer to **string** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the SMS Commands number |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**Url** | Pointer to **string** | The absolute URL of the Fleet resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


