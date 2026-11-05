# VideoV1RoomSipVideoCalls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ttid** | Pointer to **string** | The unique string that identifies the SIP video call. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SIP video call. |
**RoomSid** | Pointer to **string** | The SID of the SIP video call's room. |
**CallSid** | Pointer to **string** | The SID of the SIP video call's associated call. |
**Identity** | Pointer to **string** | The SIP participant's identity in the Room. Defaults to the `ttid` if not supplied in `Configuration.Identity`. |
**Status** | Pointer to [**string**](RoomSipVideoCallsEnumStatus.md) |  |
**SipResponseCode** | Pointer to **int** | The final SIP response code. Null if no final response was received. |
**Configuration** | Pointer to **map[string]interface{}** | The configuration supplied when creating the SIP video call. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The time the SIP video call was answered in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format. |
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The time when the SIP video call ended in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#UTC) format. |
**Duration** | Pointer to **int** | The duration in seconds that the SIP video call was `in-progress`. Populated after the call ends. |
**Url** | Pointer to **string** | The absolute URL of the resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


