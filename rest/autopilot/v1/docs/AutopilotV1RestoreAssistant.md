# AutopilotV1RestoreAssistant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Assistant resource. |
**Sid** | Pointer to **string** | The unique string that we created to identify the Assistant resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource's `sid` in the URL to address the resource. |
**FriendlyName** | Pointer to **string** | The string that you assigned to describe the resource. It is not unique and can be up to 255 characters long. |
**NeedsModelBuild** | Pointer to **bool** | Whether model needs to be rebuilt. |
**LatestModelBuildSid** | Pointer to **string** | Reserved. |
**LogQueries** | Pointer to **bool** | Whether queries should be logged and kept after training. Can be: `true` or `false` and defaults to `true`. If `true`, queries are stored for 30 days, and then deleted. If `false`, no queries are stored. |
**DevelopmentStage** | Pointer to **string** | A string describing the state of the assistant. |
**CallbackUrl** | Pointer to **string** | Reserved. |
**CallbackEvents** | Pointer to **string** | Reserved. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


