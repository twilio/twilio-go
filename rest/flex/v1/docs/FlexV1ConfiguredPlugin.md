# FlexV1ConfiguredPlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that the Flex Plugin resource is installed for. |
**ConfigurationSid** | Pointer to **string** | The SID of the Flex Plugin Configuration that this Flex Plugin belongs to. |
**PluginSid** | Pointer to **string** | The SID of the Flex Plugin. |
**PluginVersionSid** | Pointer to **string** | The SID of the Flex Plugin Version. |
**Phase** | **int** | The phase this Flex Plugin would initialize at runtime. |[optional] [default to 0]
**PluginUrl** | Pointer to **string** | The URL of where the Flex Plugin Version JavaScript bundle is hosted on. |
**UniqueName** | Pointer to **string** | The name that uniquely identifies this Flex Plugin resource. |
**FriendlyName** | Pointer to **string** | The friendly name of this Flex Plugin resource. |
**Description** | Pointer to **string** | A descriptive string that you create to describe the plugin resource. It can be up to 500 characters long |
**PluginArchived** | Pointer to **bool** | Whether the Flex Plugin is archived. The default value is false. |
**Version** | Pointer to **string** | The latest version of this Flex Plugin Version. |
**Changelog** | Pointer to **string** | A changelog that describes the changes this Flex Plugin Version brings. |
**PluginVersionArchived** | Pointer to **bool** | Whether the Flex Plugin Version is archived. The default value is false. |
**Private** | Pointer to **bool** | Whether to validate the request is authorized to access the Flex Plugin Version. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT when the Flex Plugin was installed specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the Flex Plugin resource. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


