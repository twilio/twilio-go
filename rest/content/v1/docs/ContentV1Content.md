# ContentV1Content

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format. |
**Sid** | Pointer to **string** | The unique string that that we created to identify the Content resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/usage/api/account) that created Content resource. |
**FriendlyName** | Pointer to **string** | A string name used to describe the Content resource. Not visible to the end recipient. |
**Language** | Pointer to **string** | Two-letter (ISO 639-1) language code (e.g., en) identifying the language the Content resource is in. |
**Variables** | Pointer to **interface{}** | Defines the default placeholder values for variables included in the Content resource. e.g. {\"1\": \"Customer_Name\"}. |
**Types** | Pointer to **interface{}** | The [Content types](https://www.twilio.com/docs/content-api/content-types-overview) (e.g. twilio/text) for this Content resource. |
**Url** | Pointer to **string** | The URL of the resource, relative to `https://content.twilio.com`. |
**Links** | Pointer to **map[string]interface{}** | A list of links related to the Content resource, such as approval_fetch and approval_create |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


