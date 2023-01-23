# SupersimV1UsageRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that incurred the usage. |[optional] 
**SimSid** | **string** | SID of a Sim resource to which the UsageRecord belongs. |[optional] 
**NetworkSid** | **string** | SID of the Network resource on which the usage occurred. |[optional] 
**FleetSid** | **string** | SID of the Fleet resource on which the usage occurred. |[optional] 
**IsoCountry** | **string** | Alpha-2 ISO Country Code of the country the usage occurred in. |[optional] 
**Period** | Pointer to **interface{}** | The time period for which the usage is reported. |
**DataUpload** | **int64** | Total data uploaded in bytes, aggregated by the query parameters. |[optional] 
**DataDownload** | **int64** | Total data downloaded in bytes, aggregated by the query parameters. |[optional] 
**DataTotal** | **int64** | Total of data_upload and data_download. |[optional] 
**DataTotalBilled** | **float32** | Total amount in the `billed_unit` that was charged for the data uploaded or downloaded. |[optional] 
**BilledUnit** | **string** | The currency in which the billed amounts are measured, specified in the 3 letter ISO 4127 format (e.g. `USD`, `EUR`, `JPY`). |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


