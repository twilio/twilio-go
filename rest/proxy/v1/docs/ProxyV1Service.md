# ProxyV1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | **string** | The unique string that identifies the resource |[optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource |[optional] 
**AccountSid** | **string** | The SID of the Account that created the resource |[optional] 
**ChatInstanceSid** | **string** | The SID of the Chat Service Instance |[optional] 
**CallbackUrl** | **string** | The URL we call when the interaction status changes |[optional] 
**DefaultTtl** | **int** | Default TTL for a Session, in seconds |[optional] 
**NumberSelectionBehavior** | Pointer to [**string**](ServiceEnumNumberSelectionBehavior.md) |  |
**GeoMatchLevel** | Pointer to [**string**](ServiceEnumGeoMatchLevel.md) |  |
**InterceptCallbackUrl** | **string** | The URL we call on each interaction |[optional] 
**OutOfSessionCallbackUrl** | **string** | The URL we call when an inbound call or SMS action occurs on a closed or non-existent Session |[optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was created |[optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The ISO 8601 date and time in GMT when the resource was last updated |[optional] 
**Url** | **string** | The absolute URL of the Service resource |[optional] 
**Links** | **map[string]interface{}** | The URLs of resources related to the Service |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


