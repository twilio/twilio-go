# ApiV2010RecordingTranscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ApiVersion** | **string** | The API version used to create the transcription |[optional] 
**DateCreated** | **string** | The RFC 2822 date and time in GMT that the resource was created |[optional] 
**DateUpdated** | **string** | The RFC 2822 date and time in GMT that the resource was last updated |[optional] 
**Duration** | **string** | The duration of the transcribed audio in seconds. |[optional] 
**Price** | **float32** | The charge for the transcription |[optional] 
**PriceUnit** | **string** | The currency in which price is measured |[optional] 
**RecordingSid** | **string** | The SID that identifies the transcription's recording |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**Status** | Pointer to [**string**](RecordingTranscriptionEnumStatus.md) |  |
**TranscriptionText** | **string** | The text content of the transcription. |[optional] 
**Type** | **string** | The transcription type |[optional] 
**Uri** | **string** | The URI of the resource, relative to `https://api.twilio.com` |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


