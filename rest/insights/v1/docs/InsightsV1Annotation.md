# InsightsV1Annotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallSid** | **string** | Call SID. |[optional] 
**AccountSid** | **string** | Account SID. |[optional] 
**AnsweredBy** | Pointer to [**string**](AnnotationEnumAnsweredBy.md) |  |
**ConnectivityIssue** | Pointer to [**string**](AnnotationEnumConnectivityIssue.md) |  |
**QualityIssues** | **[]string** | Indicates if the call had audio quality issues. |[optional] 
**Spam** | **bool** | Call spam indicator |[optional] 
**CallScore** | Pointer to **int** | Call Score |
**Comment** | **string** | User comments |[optional] 
**Incident** | **string** | Call tag for incidents or support ticket |[optional] 
**Url** | **string** | The URL of this resource. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


