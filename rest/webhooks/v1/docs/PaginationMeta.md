# PaginationMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key of the list property contains the actual data items. This enables programmatic iteration over paginated results.  |
**PageSize** | **int** | The page size used for this request (the requested value, or the default/max-capped value if none was requested). The number of items actually returned in the list may be less than this on the last page.  |
**PreviousToken** | Pointer to **string** | Token to fetch the previous page of results. Only included if there is a previous page, otherwise omitted.  |
**NextToken** | Pointer to **string** | Token to fetch the next page of results. Only included if there is a next page, otherwise omitted.  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


