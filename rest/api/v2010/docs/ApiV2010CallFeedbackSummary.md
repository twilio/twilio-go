# ApiV2010CallFeedbackSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique sid that identifies this account |
**CallCount** | Pointer to **int** | The total number of calls |
**CallFeedbackCount** | Pointer to **int** | The total number of calls with a feedback entry |
**DateCreated** | Pointer to **string** | The date this resource was created |
**DateUpdated** | Pointer to **string** | The date this resource was last updated |
**EndDate** | Pointer to **string** | The latest feedback entry date in the summary |
**IncludeSubaccounts** | Pointer to **bool** | Whether the feedback summary includes subaccounts |
**Issues** | Pointer to **[]interface{}** | Issues experienced during the call |
**QualityScoreAverage** | Pointer to **float32** | The average QualityScore of the feedback entries |
**QualityScoreMedian** | Pointer to **float32** | The median QualityScore of the feedback entries |
**QualityScoreStandardDeviation** | Pointer to **float32** | The standard deviation of the quality scores |
**Sid** | Pointer to **string** | A string that uniquely identifies this feedback entry |
**StartDate** | Pointer to **string** | The earliest feedback entry date in the summary |
**Status** | Pointer to **string** | The status of the feedback summary |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


