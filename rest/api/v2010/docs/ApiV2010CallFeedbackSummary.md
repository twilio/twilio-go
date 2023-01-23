# ApiV2010CallFeedbackSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. |
**CallCount** | Pointer to **int** | The total number of calls. |
**CallFeedbackCount** | Pointer to **int** | The total number of calls with a feedback entry. |
**DateCreated** | Pointer to **string** | The date that this resource was created, given in [RFC 2822](https://www.php.net/manual/en/class.datetime.php#datetime.constants.rfc2822) format. |
**DateUpdated** | Pointer to **string** | The date that this resource was last updated, given in [RFC 2822](https://www.php.net/manual/en/class.datetime.php#datetime.constants.rfc2822) format. |
**EndDate** | Pointer to **string** | The last date for which feedback entries are included in this Feedback Summary, formatted as `YYYY-MM-DD` and specified in UTC. |
**IncludeSubaccounts** | Pointer to **bool** | Whether the feedback summary includes subaccounts; `true` if it does, otherwise `false`. |
**Issues** | Pointer to **[]interface{}** | A list of issues experienced during the call. The issues can be: `imperfect-audio`, `dropped-call`, `incorrect-caller-id`, `post-dial-delay`, `digits-not-captured`, `audio-latency`, or `one-way-audio`. |
**QualityScoreAverage** | Pointer to **float32** | The average QualityScore of the feedback entries. |
**QualityScoreMedian** | Pointer to **float32** | The median QualityScore of the feedback entries. |
**QualityScoreStandardDeviation** | Pointer to **float32** | The standard deviation of the quality scores. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |
**StartDate** | Pointer to **string** | The first date for which feedback entries are included in this feedback summary, formatted as `YYYY-MM-DD` and specified in UTC. |
**Status** | Pointer to [**string**](CallFeedbackSummaryEnumStatus.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


