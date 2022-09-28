# SupersimV1Fleet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Url** | Pointer to **string** | The absolute URL of the Fleet resource |
**DataEnabled** | Pointer to **bool** | Defines whether SIMs in the Fleet are capable of using data connectivity |
**DataLimit** | Pointer to **int** | The total data usage (download and upload combined) in Megabytes that each Super SIM assigned to the Fleet can consume |
**DataMetering** | Pointer to [**string**](FleetEnumDataMetering.md) |  |
**SmsCommandsEnabled** | Pointer to **bool** | Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands |
**SmsCommandsUrl** | Pointer to **string** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the SMS Commands number |
**SmsCommandsMethod** | Pointer to **string** | A string representing the HTTP method to use when making a request to `sms_commands_url` |
**NetworkAccessProfileSid** | Pointer to **string** | The SID of the Network Access Profile of the Fleet |
**IpCommandsUrl** | Pointer to **string** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an IP Command from your device |
**IpCommandsMethod** | Pointer to **string** | A string representing the HTTP method to use when making a request to `ip_commands_url` |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


