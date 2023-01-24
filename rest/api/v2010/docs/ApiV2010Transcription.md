# ApiV2010Transcription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Transcription resource. |
**ApiVersion** | Pointer to **string** | The API version used to create the transcription. |
**DateCreated** | Pointer to **string** | The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to **string** | The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**Duration** | Pointer to **string** | The duration of the transcribed audio in seconds. |
**Price** | Pointer to **float32** | The charge for the transcript in the currency associated with the account. This value is populated after the transcript is complete so it may not be available immediately. |
**PriceUnit** | Pointer to **string** | The currency in which `price` is measured, in [ISO 4127](https://www.iso.org/iso/home/standards/currency_codes.htm) format (e.g. `usd`, `eur`, `jpy`). |
**RecordingSid** | Pointer to **string** | The SID of the [Recording](https://www.twilio.com/docs/voice/api/recording) from which the transcription was created. |
**Sid** | Pointer to **string** | The unique string that that we created to identify the Transcription resource. |
**Status** | Pointer to [**string**](TranscriptionEnumStatus.md) |  |
**TranscriptionText** | Pointer to **string** | The text content of the transcription. |
**Type** | Pointer to **string** | The transcription type. Can only be: `fast`. |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com`. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


