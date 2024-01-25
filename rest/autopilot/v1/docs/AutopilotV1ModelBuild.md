# AutopilotV1ModelBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ModelBuild resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**AssistantSid** | Pointer to **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the resource. |
**Sid** | Pointer to **string** | The unique string that we created to identify the ModelBuild resource. |
**Status** | Pointer to [**string**](ModelBuildEnumStatus.md) |  |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource. It can be used as an alternative to the `sid` in the URL path to address the resource. |
**Url** | Pointer to **string** | The absolute URL of the ModelBuild resource. |
**BuildDuration** | Pointer to **int** | The time in seconds it took to build the model. |
**ErrorCode** | Pointer to **int** | If the `status` for the model build is `failed`, this value is a code to more information about the failure. This value will be null for all other statuses. See [error code dictionary](https://www.twilio.com/docs/api/errors) for a description of the error. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


