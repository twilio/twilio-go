# ApiV2010ConferenceRecording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Conference Recording resource. |
**ApiVersion** | Pointer to **string** | The API version used to create the recording. |
**CallSid** | Pointer to **string** | The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the Conference Recording resource is associated with. |
**ConferenceSid** | Pointer to **string** | The Conference SID that identifies the conference associated with the recording. |
**DateCreated** | Pointer to **string** | The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to **string** | The date and time in GMT that the resource was last updated, specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**StartTime** | Pointer to **string** | The start time of the recording in GMT and in [RFC 2822](https://www.php.net/manual/en/class.datetime.php#datetime.constants.rfc2822) format. |
**Duration** | Pointer to **string** | The length of the recording in seconds. |
**Sid** | Pointer to **string** | The unique string that that we created to identify the Conference Recording resource. |
**Price** | Pointer to **string** | The one-time cost of creating the recording in the `price_unit` currency. |
**PriceUnit** | Pointer to **string** | The currency used in the `price` property. Example: `USD`. |
**Status** | Pointer to [**string**](ConferenceRecordingEnumStatus.md) |  |
**Channels** | Pointer to **int** | The number of channels in the final recording file.  Can be: `1`, or `2`. Separating a two leg call into two separate channels of the recording file is supported in [Dial](https://www.twilio.com/docs/voice/twiml/dial#attributes-record) and [Outbound Rest API](https://www.twilio.com/docs/voice/make-calls) record options. |
**Source** | Pointer to [**string**](ConferenceRecordingEnumSource.md) |  |
**ErrorCode** | Pointer to **int** | The error code that describes why the recording is `absent`. The error code is described in our [Error Dictionary](https://www.twilio.com/docs/api/errors). This value is null if the recording `status` is not `absent`. |
**EncryptionDetails** | Pointer to **interface{}** | How to decrypt the recording if it was encrypted using [Call Recording Encryption](https://www.twilio.com/docs/voice/tutorials/voice-recording-encryption) feature. |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com`. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


