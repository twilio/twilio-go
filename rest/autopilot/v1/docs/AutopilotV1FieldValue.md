# AutopilotV1FieldValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the FieldValue resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**FieldTypeSid** | Pointer to **string** | The SID of the Field Type associated with the Field Value. |
**Language** | Pointer to **string** | The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) tag that specifies the language of the value. Currently supported tags: `en-US` |
**AssistantSid** | Pointer to **string** | The SID of the [Assistant](https://www.twilio.com/docs/autopilot/api/assistant) that is the parent of the FieldType associated with the resource. |
**Sid** | Pointer to **string** | The unique string that we created to identify the FieldValue resource. |
**Value** | Pointer to **string** | The Field Value data. |
**Url** | Pointer to **string** | The absolute URL of the FieldValue resource. |
**SynonymOf** | Pointer to **string** | The word for which the field value is a synonym of. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


