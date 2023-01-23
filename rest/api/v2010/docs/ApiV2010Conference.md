# ApiV2010Conference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that created this resource |[optional] 
**DateCreated** | **string** | The RFC 2822 date and time in GMT that this resource was created |[optional] 
**DateUpdated** | **string** | The RFC 2822 date and time in GMT that this resource was last updated |[optional] 
**ApiVersion** | **string** | The API version used to create this conference |[optional] 
**FriendlyName** | **string** | A string that you assigned to describe this conference room |[optional] 
**Region** | **string** | A string that represents the Twilio Region where the conference was mixed |[optional] 
**Sid** | **string** | The unique string that identifies this resource |[optional] 
**Status** | Pointer to [**string**](ConferenceEnumStatus.md) |  |
**Uri** | **string** | The URI of this resource, relative to `https://api.twilio.com` |[optional] 
**SubresourceUris** | **map[string]interface{}** | A list of related resources identified by their relative URIs |[optional] 
**ReasonConferenceEnded** | Pointer to [**string**](ConferenceEnumReasonConferenceEnded.md) |  |
**CallSidEndingConference** | **string** | The call SID that caused the conference to end |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


