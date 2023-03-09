# MicrovisorV1App

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A 34-character string that uniquely identifies this App. |
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**Hash** | Pointer to **string** | App manifest hash represented as `hash_algorithm:hash_value`. |
**UniqueName** | Pointer to **string** | A developer-defined string that uniquely identifies the App. This value must be unique for all Apps on this Account. The `unique_name` value may be used as an alternative to the `sid` in the URL path to address the resource. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this App was created, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this App was last updated, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The URL of this resource. |
**Links** | Pointer to **map[string]interface{}** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


