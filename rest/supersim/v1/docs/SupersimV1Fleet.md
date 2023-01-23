# SupersimV1Fleet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the Fleet resource |[optional] 
**DataEnabled** | **bool** | Defines whether SIMs in the Fleet are capable of using data connectivity |[optional] 
**DataLimit** | **int** | The total data usage (download and upload combined) in Megabytes that each Super SIM assigned to the Fleet can consume |[optional] 
**DataMetering** | Pointer to [**string**](FleetEnumDataMetering.md) |  |
**SmsCommandsEnabled** | **bool** | Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands |[optional] 
**SmsCommandsUrl** | **string** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the SMS Commands number |[optional] 
**SmsCommandsMethod** | **string** | A string representing the HTTP method to use when making a request to `sms_commands_url` |[optional] 
**NetworkAccessProfileSid** | **string** | The SID of the Network Access Profile of the Fleet |[optional] 
**IpCommandsUrl** | **string** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an IP Command from your device |[optional] 
**IpCommandsMethod** | **string** | A string representing the HTTP method to use when making a request to `ip_commands_url` |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


