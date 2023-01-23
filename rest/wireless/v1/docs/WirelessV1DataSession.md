# WirelessV1DataSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**SimSid** | **string** | The SID of the Sim resource that the Data Session is for |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**RadioLink** | **string** | The generation of wireless technology that the device was using |[optional] 
**OperatorMcc** | **string** | The 'mobile country code' is the unique ID of the home country where the Data Session took place |[optional] 
**OperatorMnc** | **string** | The 'mobile network code' is the unique ID specific to the mobile operator network where the Data Session took place |[optional] 
**OperatorCountry** | **string** | The three letter country code representing where the device's Data Session took place |[optional] 
**OperatorName** | **string** | The friendly name of the mobile operator network that the SIM-connected device is attached to |[optional] 
**CellId** | **string** | The unique ID of the cellular tower that the device was attached to at the moment when the Data Session was last updated |[optional] 
**CellLocationEstimate** | Pointer to **interface{}** | An object with the estimated location where the device's Data Session took place |
**PacketsUploaded** | **int** | The number of packets uploaded by the device between the start time and when the Data Session was last updated |[optional] 
**PacketsDownloaded** | **int** | The number of packets downloaded by the device between the start time and when the Data Session was last updated |[optional] 
**LastUpdated** | [**time.Time**](time.Time.md) | The date that the resource was last updated, given as GMT in ISO 8601 format |[optional] 
**Start** | [**time.Time**](time.Time.md) | The date that the Data Session started, given as GMT in ISO 8601 format |[optional] 
**End** | [**time.Time**](time.Time.md) | The date that the record ended, given as GMT in ISO 8601 format |[optional] 
**Imei** | **string** | The unique ID of the device using the SIM to connect |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


