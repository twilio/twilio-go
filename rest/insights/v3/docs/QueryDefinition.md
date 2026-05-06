# QueryDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Measures** | **[]string** | Array of measures to retrieve, representing quantitative values or metrics to be calculated |[optional] 
**Dimensions** | **[]string** | Array of dimensions to retrieve, representing categorical attributes for grouping and organizing data |[optional] 
**Filters** | [**[]QueryDefinitionFilters**](QueryDefinitionFilters.md) | Nested filter conditions. Always use `op` and `expressions`. |[optional] 
**OrderBy** | [**[]QueryDefinitionOrderBy**](QueryDefinitionOrderBy.md) | Specifications for sorting the query results by specific fields in ascending or descending order |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


