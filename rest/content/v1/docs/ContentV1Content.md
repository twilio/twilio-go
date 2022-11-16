# ContentV1Content

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT that the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT that the resource was last updated |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**FriendlyName** | Pointer to **string** | A string name used to describe the Content resource |
**Language** | Pointer to **string** | Two-letter language code identifying the language the Content resource is in. |
**Variables** | Pointer to **interface{}** | Defines the default placeholder values for variables included in the Content resource |
**Types** | Pointer to **interface{}** | The Content types (e.g. twilio/text) for this Content resource |
**Url** | Pointer to **string** | The URL of the resource, relative to `https://content.twilio.com` |
**Links** | Pointer to **map[string]interface{}** | A list of links related to the Content resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


