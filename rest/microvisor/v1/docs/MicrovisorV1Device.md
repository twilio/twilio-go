# MicrovisorV1Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | A 34-character string that uniquely identifies this Device. |
**UniqueName** | Pointer to **string** | A developer-defined string that uniquely identifies the Device. This value must be unique for all Devices on this Account. The `unique_name` value may be used as an alternative to the `sid` in the URL path to address the resource. |
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**App** | Pointer to **interface{}** | Information about the target App and the App reported by this Device. Contains the properties `target_sid`, `date_targeted`, `update_status` (one of `up-to-date`, `pending` and `error`), `update_error_code`, `reported_sid` and `date_reported`. |
**Logging** | Pointer to **interface{}** | Object specifying whether application logging is enabled for this Device. Contains the properties `enabled` and `date_expires`. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this Device was created, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this Device was last updated, given in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. |
**Url** | Pointer to **string** | The URL of this resource. |
**Links** | Pointer to **map[string]interface{}** | The absolute URLs of related resources. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


