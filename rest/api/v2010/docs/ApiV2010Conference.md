# ApiV2010Conference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created this Conference resource. |
**DateCreated** | Pointer to **string** | The date and time in GMT that this resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to **string** | The date and time in GMT that this resource was last updated, specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**ApiVersion** | Pointer to **string** | The API version used to create this conference. |
**FriendlyName** | Pointer to **string** | A string that you assigned to describe this conference room. Maxiumum length is 128 characters. |
**Region** | Pointer to **string** | A string that represents the Twilio Region where the conference audio was mixed. May be `us1`, `ie1`,  `de1`, `sg1`, `br1`, `au1`, and `jp1`. Basic conference audio will always be mixed in `us1`. Global Conference audio will be mixed nearest to the majority of participants. |
**Sid** | Pointer to **string** | The unique string that that we created to identify this Conference resource. |
**Status** | Pointer to [**string**](ConferenceEnumStatus.md) |  |
**Uri** | Pointer to **string** | The URI of this resource, relative to `https://api.twilio.com`. |
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their URIs relative to `https://api.twilio.com`. |
**ReasonConferenceEnded** | Pointer to [**string**](ConferenceEnumReasonConferenceEnded.md) |  |
**CallSidEndingConference** | Pointer to **string** | The call SID that caused the conference to end. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


