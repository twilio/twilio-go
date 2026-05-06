# VoiceV3CreateRequestWithMediaUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TranscriptionConfigurationId** | **string** | The ID of the transcription configuration to use |
**InputSource** | **string** | Discriminator indicating the input source type |[optional] 
**MediaUrl** | **string** | URL to the media file to transcribe |
**AudioStartedAt** | [**time.Time**](time.Time.md) | The start time of the audio recording |[optional] 
**Participants** | [**[]VoiceV3Participant**](VoiceV3Participant.md) | Participants in the conversation. If omitted or partially specified, defaults from the transcription configuration will be applied.  |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


