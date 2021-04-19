# ProxyV1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account that created the resource | [optional] 
**CallbackUrl** | Pointer to **NullableString** | The URL we call when the interaction status changes | [optional] 
**ChatInstanceSid** | Pointer to **NullableString** | The SID of the Chat Service Instance | [optional] 
**DateCreated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was created | [optional] 
**DateUpdated** | Pointer to **NullableTime** | The ISO 8601 date and time in GMT when the resource was last updated | [optional] 
**DefaultTtl** | Pointer to **NullableInt32** | Default TTL for a Session, in seconds | [optional] 
**GeoMatchLevel** | Pointer to **NullableString** | Where a proxy number must be located relative to the participant identifier | [optional] 
**InterceptCallbackUrl** | Pointer to **NullableString** | The URL we call on each interaction | [optional] 
**Links** | Pointer to **map[string]interface{}** | The URLs of resources related to the Service | [optional] 
**NumberSelectionBehavior** | Pointer to **NullableString** | The preference for Proxy Number selection for the Service instance | [optional] 
**OutOfSessionCallbackUrl** | Pointer to **NullableString** | The URL we call when an inbound call or SMS action occurs on a closed or non-existent Session | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 
**UniqueName** | Pointer to **NullableString** | An application-defined string that uniquely identifies the resource | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL of the Service resource | [optional] 

## Methods

### NewProxyV1Service

`func NewProxyV1Service() *ProxyV1Service`

NewProxyV1Service instantiates a new ProxyV1Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyV1ServiceWithDefaults

`func NewProxyV1ServiceWithDefaults() *ProxyV1Service`

NewProxyV1ServiceWithDefaults instantiates a new ProxyV1Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ProxyV1Service) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ProxyV1Service) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ProxyV1Service) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ProxyV1Service) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ProxyV1Service) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ProxyV1Service) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetCallbackUrl

`func (o *ProxyV1Service) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *ProxyV1Service) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *ProxyV1Service) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *ProxyV1Service) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### SetCallbackUrlNil

`func (o *ProxyV1Service) SetCallbackUrlNil(b bool)`

 SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil

### UnsetCallbackUrl
`func (o *ProxyV1Service) UnsetCallbackUrl()`

UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
### GetChatInstanceSid

`func (o *ProxyV1Service) GetChatInstanceSid() string`

GetChatInstanceSid returns the ChatInstanceSid field if non-nil, zero value otherwise.

### GetChatInstanceSidOk

`func (o *ProxyV1Service) GetChatInstanceSidOk() (*string, bool)`

GetChatInstanceSidOk returns a tuple with the ChatInstanceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatInstanceSid

`func (o *ProxyV1Service) SetChatInstanceSid(v string)`

SetChatInstanceSid sets ChatInstanceSid field to given value.

### HasChatInstanceSid

`func (o *ProxyV1Service) HasChatInstanceSid() bool`

HasChatInstanceSid returns a boolean if a field has been set.

### SetChatInstanceSidNil

`func (o *ProxyV1Service) SetChatInstanceSidNil(b bool)`

 SetChatInstanceSidNil sets the value for ChatInstanceSid to be an explicit nil

### UnsetChatInstanceSid
`func (o *ProxyV1Service) UnsetChatInstanceSid()`

UnsetChatInstanceSid ensures that no value is present for ChatInstanceSid, not even an explicit nil
### GetDateCreated

`func (o *ProxyV1Service) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ProxyV1Service) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ProxyV1Service) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ProxyV1Service) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ProxyV1Service) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ProxyV1Service) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *ProxyV1Service) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *ProxyV1Service) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *ProxyV1Service) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *ProxyV1Service) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *ProxyV1Service) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *ProxyV1Service) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetDefaultTtl

`func (o *ProxyV1Service) GetDefaultTtl() int32`

GetDefaultTtl returns the DefaultTtl field if non-nil, zero value otherwise.

### GetDefaultTtlOk

`func (o *ProxyV1Service) GetDefaultTtlOk() (*int32, bool)`

GetDefaultTtlOk returns a tuple with the DefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtl

`func (o *ProxyV1Service) SetDefaultTtl(v int32)`

SetDefaultTtl sets DefaultTtl field to given value.

### HasDefaultTtl

`func (o *ProxyV1Service) HasDefaultTtl() bool`

HasDefaultTtl returns a boolean if a field has been set.

### SetDefaultTtlNil

`func (o *ProxyV1Service) SetDefaultTtlNil(b bool)`

 SetDefaultTtlNil sets the value for DefaultTtl to be an explicit nil

### UnsetDefaultTtl
`func (o *ProxyV1Service) UnsetDefaultTtl()`

UnsetDefaultTtl ensures that no value is present for DefaultTtl, not even an explicit nil
### GetGeoMatchLevel

`func (o *ProxyV1Service) GetGeoMatchLevel() string`

GetGeoMatchLevel returns the GeoMatchLevel field if non-nil, zero value otherwise.

### GetGeoMatchLevelOk

`func (o *ProxyV1Service) GetGeoMatchLevelOk() (*string, bool)`

GetGeoMatchLevelOk returns a tuple with the GeoMatchLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoMatchLevel

`func (o *ProxyV1Service) SetGeoMatchLevel(v string)`

SetGeoMatchLevel sets GeoMatchLevel field to given value.

### HasGeoMatchLevel

`func (o *ProxyV1Service) HasGeoMatchLevel() bool`

HasGeoMatchLevel returns a boolean if a field has been set.

### SetGeoMatchLevelNil

`func (o *ProxyV1Service) SetGeoMatchLevelNil(b bool)`

 SetGeoMatchLevelNil sets the value for GeoMatchLevel to be an explicit nil

### UnsetGeoMatchLevel
`func (o *ProxyV1Service) UnsetGeoMatchLevel()`

UnsetGeoMatchLevel ensures that no value is present for GeoMatchLevel, not even an explicit nil
### GetInterceptCallbackUrl

`func (o *ProxyV1Service) GetInterceptCallbackUrl() string`

GetInterceptCallbackUrl returns the InterceptCallbackUrl field if non-nil, zero value otherwise.

### GetInterceptCallbackUrlOk

`func (o *ProxyV1Service) GetInterceptCallbackUrlOk() (*string, bool)`

GetInterceptCallbackUrlOk returns a tuple with the InterceptCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterceptCallbackUrl

`func (o *ProxyV1Service) SetInterceptCallbackUrl(v string)`

SetInterceptCallbackUrl sets InterceptCallbackUrl field to given value.

### HasInterceptCallbackUrl

`func (o *ProxyV1Service) HasInterceptCallbackUrl() bool`

HasInterceptCallbackUrl returns a boolean if a field has been set.

### SetInterceptCallbackUrlNil

`func (o *ProxyV1Service) SetInterceptCallbackUrlNil(b bool)`

 SetInterceptCallbackUrlNil sets the value for InterceptCallbackUrl to be an explicit nil

### UnsetInterceptCallbackUrl
`func (o *ProxyV1Service) UnsetInterceptCallbackUrl()`

UnsetInterceptCallbackUrl ensures that no value is present for InterceptCallbackUrl, not even an explicit nil
### GetLinks

`func (o *ProxyV1Service) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProxyV1Service) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProxyV1Service) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProxyV1Service) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ProxyV1Service) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ProxyV1Service) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetNumberSelectionBehavior

`func (o *ProxyV1Service) GetNumberSelectionBehavior() string`

GetNumberSelectionBehavior returns the NumberSelectionBehavior field if non-nil, zero value otherwise.

### GetNumberSelectionBehaviorOk

`func (o *ProxyV1Service) GetNumberSelectionBehaviorOk() (*string, bool)`

GetNumberSelectionBehaviorOk returns a tuple with the NumberSelectionBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberSelectionBehavior

`func (o *ProxyV1Service) SetNumberSelectionBehavior(v string)`

SetNumberSelectionBehavior sets NumberSelectionBehavior field to given value.

### HasNumberSelectionBehavior

`func (o *ProxyV1Service) HasNumberSelectionBehavior() bool`

HasNumberSelectionBehavior returns a boolean if a field has been set.

### SetNumberSelectionBehaviorNil

`func (o *ProxyV1Service) SetNumberSelectionBehaviorNil(b bool)`

 SetNumberSelectionBehaviorNil sets the value for NumberSelectionBehavior to be an explicit nil

### UnsetNumberSelectionBehavior
`func (o *ProxyV1Service) UnsetNumberSelectionBehavior()`

UnsetNumberSelectionBehavior ensures that no value is present for NumberSelectionBehavior, not even an explicit nil
### GetOutOfSessionCallbackUrl

`func (o *ProxyV1Service) GetOutOfSessionCallbackUrl() string`

GetOutOfSessionCallbackUrl returns the OutOfSessionCallbackUrl field if non-nil, zero value otherwise.

### GetOutOfSessionCallbackUrlOk

`func (o *ProxyV1Service) GetOutOfSessionCallbackUrlOk() (*string, bool)`

GetOutOfSessionCallbackUrlOk returns a tuple with the OutOfSessionCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfSessionCallbackUrl

`func (o *ProxyV1Service) SetOutOfSessionCallbackUrl(v string)`

SetOutOfSessionCallbackUrl sets OutOfSessionCallbackUrl field to given value.

### HasOutOfSessionCallbackUrl

`func (o *ProxyV1Service) HasOutOfSessionCallbackUrl() bool`

HasOutOfSessionCallbackUrl returns a boolean if a field has been set.

### SetOutOfSessionCallbackUrlNil

`func (o *ProxyV1Service) SetOutOfSessionCallbackUrlNil(b bool)`

 SetOutOfSessionCallbackUrlNil sets the value for OutOfSessionCallbackUrl to be an explicit nil

### UnsetOutOfSessionCallbackUrl
`func (o *ProxyV1Service) UnsetOutOfSessionCallbackUrl()`

UnsetOutOfSessionCallbackUrl ensures that no value is present for OutOfSessionCallbackUrl, not even an explicit nil
### GetSid

`func (o *ProxyV1Service) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *ProxyV1Service) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *ProxyV1Service) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *ProxyV1Service) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *ProxyV1Service) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *ProxyV1Service) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetUniqueName

`func (o *ProxyV1Service) GetUniqueName() string`

GetUniqueName returns the UniqueName field if non-nil, zero value otherwise.

### GetUniqueNameOk

`func (o *ProxyV1Service) GetUniqueNameOk() (*string, bool)`

GetUniqueNameOk returns a tuple with the UniqueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueName

`func (o *ProxyV1Service) SetUniqueName(v string)`

SetUniqueName sets UniqueName field to given value.

### HasUniqueName

`func (o *ProxyV1Service) HasUniqueName() bool`

HasUniqueName returns a boolean if a field has been set.

### SetUniqueNameNil

`func (o *ProxyV1Service) SetUniqueNameNil(b bool)`

 SetUniqueNameNil sets the value for UniqueName to be an explicit nil

### UnsetUniqueName
`func (o *ProxyV1Service) UnsetUniqueName()`

UnsetUniqueName ensures that no value is present for UniqueName, not even an explicit nil
### GetUrl

`func (o *ProxyV1Service) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProxyV1Service) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProxyV1Service) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProxyV1Service) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ProxyV1Service) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ProxyV1Service) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


