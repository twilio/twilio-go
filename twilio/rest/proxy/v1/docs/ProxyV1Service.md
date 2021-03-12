# ProxyV1Service

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**CallbackUrl** | Pointer to **string** | The URL we call when the interaction status changes |
**ChatInstanceSid** | Pointer to **string** | The SID of the Chat Service Instance |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |
**DefaultTtl** | Pointer to **int32** | Default TTL for a Session, in seconds |
**GeoMatchLevel** | Pointer to **string** | Where a proxy number must be located relative to the participant identifier |
**InterceptCallbackUrl** | Pointer to **string** | The URL we call on each interaction |
**Links** | Pointer to **map[string]interface{}** | The URLs of resources related to the Service |
**NumberSelectionBehavior** | Pointer to **string** | The preference for Proxy Number selection for the Service instance |
**OutOfSessionCallbackUrl** | Pointer to **string** | The URL we call when an inbound call or SMS action occurs on a closed or non-existent Session |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**UniqueName** | Pointer to **string** | An application-defined string that uniquely identifies the resource |
**Url** | Pointer to **string** | The absolute URL of the Service resource |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


