# DataMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | Name of the data mapping. |
**Description** | **string** | A human readable description of this resource, up to 512 characters. |[optional] 
**IsEnabled** | **bool** | Flag indicating whether the data mapping is active. When true, data will be ingested and mapped according to the configuration. When false, the data mapping will be inactive and no data will be ingested into the Memory Store. |[optional] [default to true]
**MappingTo** | [**DataMappingToTraits**](DataMappingToTraits.md) |  |
**MappingFrom** | [**DataMappingFromDataSet**](DataMappingFromDataSet.md) |  |
**Id** | **string** | The unique identifier for the data mapping. |
**CreatedAt** | [**time.Time**](time.Time.md) | The ISO 8601 timestamp when the  data mapping was created. |
**UpdatedAt** | [**time.Time**](time.Time.md) | The ISO 8601 timestamp when the data mapping was last updated. |
**Version** | **int** | The current version number of the DataMapping. Incremented on each successful update. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


