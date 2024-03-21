# FlexV1Plugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that we created to identify the Flex Plugin resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Flex Plugin resource and owns this resource. |
**UniqueName** | Pointer to **string** | The name that uniquely identifies this Flex Plugin resource. |
**FriendlyName** | Pointer to **string** | The friendly name this Flex Plugin resource. |
**Description** | Pointer to **string** | A descriptive string that you create to describe the plugin resource. It can be up to 500 characters long |
**Archived** | Pointer to **bool** | Whether the Flex Plugin is archived. The default value is false. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT-7 when the Flex Plugin was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date and time in GMT-7 when the Flex Plugin was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The absolute URL of the Flex Plugin resource. |
**Links** | Pointer to **map[string]interface{}** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


