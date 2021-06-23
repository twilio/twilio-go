# SupersimV1UsageRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that incurred the usage. |
**DataDownload** | Pointer to **int** | Total data downloaded in bytes, aggregated by the query parameters. |
**DataTotal** | Pointer to **int** | Total of data_upload and data_download. |
**DataUpload** | Pointer to **int** | Total data uploaded in bytes, aggregated by the query parameters. |
**FleetSid** | Pointer to **string** | SID of the Fleet resource on which the usage occurred. |
**IsoCountry** | Pointer to **string** | Alpha-2 ISO Country Code of the country the usage occurred in. |
**NetworkSid** | Pointer to **string** | SID of the Network resource on which the usage occurred. |
**Period** | Pointer to **map[string]interface{}** | The time period for which the usage is reported. |
**SimSid** | Pointer to **string** | SID of a Sim resource to which the UsageRecord belongs. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


