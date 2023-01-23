# ApiV2010ConferenceRecording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ApiVersion** | **string** | The API version used to create the recording |[optional] 
**CallSid** | **string** | The SID of the Call the resource is associated with |[optional] 
**ConferenceSid** | **string** | The Conference SID that identifies the conference associated with the recording |[optional] 
**DateCreated** | **string** | The RFC 2822 date and time in GMT that the resource was created |[optional] 
**DateUpdated** | **string** | The RFC 2822 date and time in GMT that the resource was last updated |[optional] 
**StartTime** | **string** | The start time of the recording, given in RFC 2822 format |[optional] 
**Duration** | **string** | The length of the recording in seconds |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**Price** | **string** | The one-time cost of creating the recording. |[optional] 
**PriceUnit** | **string** | The currency used in the price property. |[optional] 
**Status** | Pointer to [**string**](ConferenceRecordingEnumStatus.md) |  |
**Channels** | **int** | The number of channels in the final recording file as an integer |[optional] 
**Source** | Pointer to [**string**](ConferenceRecordingEnumSource.md) |  |
**ErrorCode** | Pointer to **int** | More information about why the recording is missing, if status is `absent`. |
**EncryptionDetails** | Pointer to **interface{}** | How to decrypt the recording. |
**Uri** | **string** | The URI of the resource, relative to `https://api.twilio.com` |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


