# VoiceV2Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationType** | **string** | The configuration type discriminator. Always \"Recording\" for this resource. |
**CompositionPolicy** | [**VoiceV2CompositionPolicy**](VoiceV2CompositionPolicy.md) |  |[optional] 
**CallRecordingStatusCallback** | [**VoiceV2StatusCallback**](VoiceV2StatusCallback.md) |  |[optional] 
**ConferenceRecordingStatusCallback** | [**VoiceV2StatusCallback**](VoiceV2StatusCallback.md) |  |[optional] 
**Features** | Pointer to [**[]VoiceV2Feature**](VoiceV2Feature.md) | The features to apply to this recording. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


