# FlexV1PluginVersionArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Flex Plugin Version resource. |
**PluginSid** | Pointer to **string** | The SID of the Flex Plugin resource this Flex Plugin Version belongs to. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Flex Plugin Version resource and owns this resource. |
**Version** | Pointer to **string** | The unique version of this Flex Plugin Version. |
**PluginUrl** | Pointer to **string** | The URL of where the Flex Plugin Version JavaScript bundle is hosted on. |
**Changelog** | Pointer to **string** | A changelog that describes the changes this Flex Plugin Version brings. |
**Private** | Pointer to **bool** | Whether to inject credentials while accessing this Plugin Version. The default value is false. |
**Archived** | Pointer to **bool** | Whether the Flex Plugin Version is archived. The default value is false. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Flex Plugin Version was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the Flex Plugin Version resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


