# CreateV3TranscriptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TranscriptionConfigurationId** | **string** | The ID of the transcription configuration to use |
**InputSource** | **string** | Discriminator indicating the input source type |[optional] 
**SourceId** | **string** | The SID or TTID of the source audio to transcribe (e.g. a Twilio Recording SID). When provided, audioStartedAt is inferred from the recording's start time and does not need to be supplied by the caller.  |
**Participants** | [**[]VoiceV3Participant**](VoiceV3Participant.md) | Participants in the conversation. If omitted or partially specified, defaults from the transcription configuration will be applied.  |[optional] 
**MediaUrl** | **string** | URL to the media file to transcribe |
**AudioStartedAt** | [**time.Time**](time.Time.md) | The start time of the audio recording |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


