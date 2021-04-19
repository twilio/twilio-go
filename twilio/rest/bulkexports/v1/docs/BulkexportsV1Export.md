# BulkexportsV1Export

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Links** | Pointer to **map[string]interface{}** | Nested resource URLs. | [optional] 
**ResourceType** | Pointer to **NullableString** | The type of communication – Messages, Calls, Conferences, and Participants | [optional] 
**Url** | Pointer to **NullableString** | The URL of this resource. | [optional] 

## Methods

### NewBulkexportsV1Export

`func NewBulkexportsV1Export() *BulkexportsV1Export`

NewBulkexportsV1Export instantiates a new BulkexportsV1Export object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkexportsV1ExportWithDefaults

`func NewBulkexportsV1ExportWithDefaults() *BulkexportsV1Export`

NewBulkexportsV1ExportWithDefaults instantiates a new BulkexportsV1Export object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *BulkexportsV1Export) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BulkexportsV1Export) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BulkexportsV1Export) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BulkexportsV1Export) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *BulkexportsV1Export) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *BulkexportsV1Export) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetResourceType

`func (o *BulkexportsV1Export) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BulkexportsV1Export) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BulkexportsV1Export) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BulkexportsV1Export) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *BulkexportsV1Export) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *BulkexportsV1Export) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetUrl

`func (o *BulkexportsV1Export) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BulkexportsV1Export) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BulkexportsV1Export) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *BulkexportsV1Export) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *BulkexportsV1Export) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *BulkexportsV1Export) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


