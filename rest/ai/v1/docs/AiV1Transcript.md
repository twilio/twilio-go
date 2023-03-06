# AiV1Transcript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**CallSid** | Pointer to **string** | The unique SID identifier of the Call. |
**Confidence** | Pointer to **float32** | The overall confidence score for the transcript. |
**DataLogging** | Pointer to **bool** | Data logging allows Twilio to improve the quality of the speech recognition through using customer data to refine its speech recognition models. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this Transcript was created, given in ISO 8601 format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this Transcript was updated, given in ISO 8601 format. |
**Environment** | Pointer to **string** | The environment where the audio is coming from. Values can be 'telephony', 'meeting_room', or 'broadcast'. |
**LanguageLocale** | Pointer to **string** | The default language locale of the audio. |
**Links** | Pointer to **map[string]interface{}** | Contains a dictionary of URL links to nested resources of this Transcript. |
**ModelSid** | Pointer to **string** | The unique SID identifier of the Model. |
**Price** | Pointer to **float32** | The charge for the transcript in the currency associated with the account. This value is populated after the transcript is complete so it may not be available immediately. |
**PriceUnit** | Pointer to **string** | The currency in which price is measured, in ISO 4127 format (e.g. usd, eur, jpy). |
**RecordingSid** | Pointer to **string** | The unique SID identifier of the Recording. |
**ServiceSid** | Pointer to **string** | The unique SID identifier of the Service. |
**Sid** | Pointer to **string** | A 34 character string that uniquely identifies this Transcript. |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The date that this Transcript was started, given in ISO 8601 format. |
**Status** | Pointer to [**string**](TranscriptEnumStatus.md) |  |
**TranscriptStatus** | Pointer to [**string**](TranscriptEnumTranscriptStatus.md) |  |
**Url** | Pointer to **string** | The URL of this resource. |
**LupOutputs** | Pointer to **interface{}** | List of language understanding outputs. |
**LupAppliedModels** | Pointer to **interface{}** | List of language understanding models applied. |
**CallDirection** | Pointer to **string** | The direction of this Transcript's call. One of `inbound`, `outbound`, `internal` or `unknown`. |
**CallFrom** | Pointer to **string** | The From of this Transcript's call. A string represents the phone number. |
**CallTo** | Pointer to **string** | The To of this Transcript's call. A string represents the phone number. |
**CallDuration** | Pointer to **float32** | The duration of this Transcript's call. |
**CallStartTime** | Pointer to [**time.Time**](time.Time.md) | The date this transcript call was started, given in ISO 8601 format. |
**Participants** | Pointer to **interface{}** |  |
**HasSensitiveView** | Pointer to **bool** | If the transcript has been redacted, a sensitive alternative of the transcript will be available. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


