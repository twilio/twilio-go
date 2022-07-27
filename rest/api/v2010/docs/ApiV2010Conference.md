# ApiV2010Conference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created this resource |
**DateCreated** | Pointer to **string** | The RFC 2822 date and time in GMT that this resource was created |
**DateUpdated** | Pointer to **string** | The RFC 2822 date and time in GMT that this resource was last updated |
**ApiVersion** | Pointer to **string** | The API version used to create this conference |
**FriendlyName** | Pointer to **string** | A string that you assigned to describe this conference room |
**Region** | Pointer to **string** | A string that represents the Twilio Region where the conference was mixed |
**Sid** | Pointer to **string** | The unique string that identifies this resource |
**Status** | Pointer to [**string**](ConferenceEnumStatus.md) |  |
**Uri** | Pointer to **string** | The URI of this resource, relative to `https://api.twilio.com` |
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their relative URIs |
**ReasonConferenceEnded** | Pointer to [**string**](ConferenceEnumReasonConferenceEnded.md) |  |
**CallSidEndingConference** | Pointer to **string** | The call SID that caused the conference to end |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


