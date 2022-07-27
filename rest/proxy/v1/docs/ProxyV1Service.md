# ProxyV1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ChatInstanceSid** | Pointer to **string** | The SID of the Chat Service Instance |
**CallbackUrl** | Pointer to **string** | The URL we call when the interaction status changes |
**DefaultTtl** | Pointer to **int** | Default TTL for a Session, in seconds |
**NumberSelectionBehavior** | Pointer to [**string**](ServiceEnumNumberSelectionBehavior.md) |  |
**GeoMatchLevel** | Pointer to [**string**](ServiceEnumGeoMatchLevel.md) |  |
**InterceptCallbackUrl** | Pointer to **string** | The URL we call on each interaction |
**OutOfSessionCallbackUrl** | Pointer to **string** | The URL we call when an inbound call or SMS action occurs on a closed or non-existent Session |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**Url** | Pointer to **string** | The absolute URL of the Service resource |
**Links** | Pointer to **map[string]interface{}** | The URLs of resources related to the Service |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


