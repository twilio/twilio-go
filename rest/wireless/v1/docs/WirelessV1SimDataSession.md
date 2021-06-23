# WirelessV1SimDataSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**CellId** | Pointer to **string** | The unique ID of the cellular tower that the device was attached to at the moment when the Data Session was last updated |
**CellLocationEstimate** | Pointer to **map[string]interface{}** | An object with the estimated location where the device's Data Session took place |
**End** | Pointer to [**time.Time**](time.Time.md) | The date that the record ended, given as GMT in ISO 8601 format |
**Imei** | Pointer to **string** | The unique ID of the device using the SIM to connect |
**LastUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that the resource was last updated, given as GMT in ISO 8601 format |
**OperatorCountry** | Pointer to **string** | The three letter country code representing where the device's Data Session took place |
**OperatorMcc** | Pointer to **string** | The 'mobile country code' is the unique ID of the home country where the Data Session took place |
**OperatorMnc** | Pointer to **string** | The 'mobile network code' is the unique ID specific to the mobile operator network where the Data Session took place |
**OperatorName** | Pointer to **string** | The friendly name of the mobile operator network that the SIM-connected device is attached to |
**PacketsDownloaded** | Pointer to **int** | The number of packets downloaded by the device between the start time and when the Data Session was last updated |
**PacketsUploaded** | Pointer to **int** | The number of packets uploaded by the device between the start time and when the Data Session was last updated |
**RadioLink** | Pointer to **string** | The generation of wireless technology that the device was using |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**SimSid** | Pointer to **string** | The SID of the Sim resource that the Data Session is for |
**Start** | Pointer to [**time.Time**](time.Time.md) | The date that the Data Session started, given as GMT in ISO 8601 format |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


