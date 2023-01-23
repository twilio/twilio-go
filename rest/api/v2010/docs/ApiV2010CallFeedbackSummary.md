# ApiV2010CallFeedbackSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The unique sid that identifies this account |[optional] 
**CallCount** | **int** | The total number of calls |[optional] 
**CallFeedbackCount** | **int** | The total number of calls with a feedback entry |[optional] 
**DateCreated** | **string** | The date this resource was created |[optional] 
**DateUpdated** | **string** | The date this resource was last updated |[optional] 
**EndDate** | **string** | The latest feedback entry date in the summary |[optional] 
**IncludeSubaccounts** | **bool** | Whether the feedback summary includes subaccounts |[optional] 
**Issues** | **[]interface{}** | Issues experienced during the call |[optional] 
**QualityScoreAverage** | **float32** | The average QualityScore of the feedback entries |[optional] 
**QualityScoreMedian** | **float32** | The median QualityScore of the feedback entries |[optional] 
**QualityScoreStandardDeviation** | **float32** | The standard deviation of the quality scores |[optional] 
**Sid** | **string** | A string that uniquely identifies this feedback entry |[optional] 
**StartDate** | **string** | The earliest feedback entry date in the summary |[optional] 
**Status** | Pointer to [**string**](CallFeedbackSummaryEnumStatus.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


