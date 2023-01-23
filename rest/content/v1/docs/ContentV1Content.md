# ContentV1Content

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateCreated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT that the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The RFC 2822 date and time in GMT that the resource was last updated |[optional] 
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**FriendlyName** | **string** | A string name used to describe the Content resource |[optional] 
**Language** | **string** | Two-letter language code identifying the language the Content resource is in. |[optional] 
**Variables** | Pointer to **interface{}** | Defines the default placeholder values for variables included in the Content resource |
**Types** | Pointer to **interface{}** | The Content types (e.g. twilio/text) for this Content resource |
**Url** | **string** | The URL of the resource, relative to `https://content.twilio.com` |[optional] 
**Links** | **map[string]interface{}** | A list of links related to the Content resource |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


