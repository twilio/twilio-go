# ApiV2010CallFeedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource. |
**DateCreated** | Pointer to **string** | The date that this resource was created, given in [RFC 2822](https://www.php.net/manual/en/class.datetime.php#datetime.constants.rfc2822) format. |
**DateUpdated** | Pointer to **string** | The date that this resource was last updated, given in [RFC 2822](https://www.php.net/manual/en/class.datetime.php#datetime.constants.rfc2822) format. |
**Issues** | Pointer to [**[]string**](CallFeedbackEnumIssues.md) | A list of issues experienced during the call. The issues can be: `imperfect-audio`, `dropped-call`, `incorrect-caller-id`, `post-dial-delay`, `digits-not-captured`, `audio-latency`, `unsolicited-call`, or `one-way-audio`. |
**QualityScore** | Pointer to **int** | `1` to `5` quality score where `1` represents imperfect experience and `5` represents a perfect call. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


