# WirelessV1DataSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the DataSession resource. |
**SimSid** | Pointer to **string** | The SID of the [Sim resource](https://www.twilio.com/docs/wireless/api/sim-resource) that the Data Session is for. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the DataSession resource. |
**RadioLink** | Pointer to **string** | The generation of wireless technology that the device was using. |
**OperatorMcc** | Pointer to **string** | The 'mobile country code' is the unique ID of the home country where the Data Session took place. See: [MCC/MNC lookup](http://mcc-mnc.com/). |
**OperatorMnc** | Pointer to **string** | The 'mobile network code' is the unique ID specific to the mobile operator network where the Data Session took place. |
**OperatorCountry** | Pointer to **string** | The three letter country code representing where the device's Data Session took place. This is determined by looking up the `operator_mcc`. |
**OperatorName** | Pointer to **string** | The friendly name of the mobile operator network that the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource)-connected device is attached to. This is determined by looking up the `operator_mnc`. |
**CellId** | Pointer to **string** | The unique ID of the cellular tower that the device was attached to at the moment when the Data Session was last updated. |
**CellLocationEstimate** | Pointer to **interface{}** | An object that describes the estimated location in latitude and longitude where the device's Data Session took place. The location is derived from the `cell_id` when the Data Session was last updated. See [Cell Location Estimate Object](https://www.twilio.com/docs/wireless/api/datasession-resource#cell-location-estimate-object).  |
**PacketsUploaded** | Pointer to **int** | The number of packets uploaded by the device between the `start` time and when the Data Session was last updated. |
**PacketsDownloaded** | Pointer to **int** | The number of packets downloaded by the device between the `start` time and when the Data Session was last updated. |
**LastUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that the resource was last updated, given as GMT in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format. |
**Start** | Pointer to [**time.Time**](time.Time.md) | The date that the Data Session started, given as GMT in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format. |
**End** | Pointer to [**time.Time**](time.Time.md) | The date that the record ended, given as GMT in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format. |
**Imei** | Pointer to **string** | The 'international mobile equipment identity' is the unique ID of the device using the SIM to connect. An IMEI is a 15-digit string: 14 digits for the device identifier plus a check digit calculated using the Luhn formula. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


