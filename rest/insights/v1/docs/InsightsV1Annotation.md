# InsightsV1Annotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallSid** | Pointer to **string** | The unique SID identifier of the Call. |
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**AnsweredBy** | Pointer to [**string**](AnnotationEnumAnsweredBy.md) |  |
**ConnectivityIssue** | Pointer to [**string**](AnnotationEnumConnectivityIssue.md) |  |
**QualityIssues** | Pointer to **[]string** | Specifies if the call had any subjective quality issues. Possible values are one or more of `no_quality_issue`, `low_volume`, `choppy_robotic`, `echo`, `dtmf`, `latency`, `owa`, or `static_noise`. |
**Spam** | Pointer to **bool** | Specifies if the call was a spam call. Use this to provide feedback on whether calls placed from your account were marked as spam, or if inbound calls received by your account were unwanted spam. Is of type Boolean: true, false. Use true if the call was a spam call. |
**CallScore** | Pointer to **int** | Specifies the Call Score, if available. This is of type integer. Use a range of 1-5 to indicate the call experience score, with the following mapping as a reference for rating the call [5: Excellent, 4: Good, 3 : Fair, 2 : Poor, 1: Bad]. |
**Comment** | Pointer to **string** | Specifies any comments pertaining to the call. Twilio does not treat this field as PII, so no PII should be included in comments. |
**Incident** | Pointer to **string** | Incident or support ticket associated with this call. The `incident` property is of type string with a maximum character limit of 100. Twilio does not treat this field as PII, so no PII should be included in `incident`. |
**Url** | Pointer to **string** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


