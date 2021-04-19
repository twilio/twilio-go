# SyncV1ServiceSyncStreamStreamMessage

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** | Stream Message body | [optional] 
**Sid** | Pointer to **NullableString** | The unique string that identifies the resource | [optional] 

## Methods

### NewSyncV1ServiceSyncStreamStreamMessage

`func NewSyncV1ServiceSyncStreamStreamMessage() *SyncV1ServiceSyncStreamStreamMessage`

NewSyncV1ServiceSyncStreamStreamMessage instantiates a new SyncV1ServiceSyncStreamStreamMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncV1ServiceSyncStreamStreamMessageWithDefaults

`func NewSyncV1ServiceSyncStreamStreamMessageWithDefaults() *SyncV1ServiceSyncStreamStreamMessage`

NewSyncV1ServiceSyncStreamStreamMessageWithDefaults instantiates a new SyncV1ServiceSyncStreamStreamMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SyncV1ServiceSyncStreamStreamMessage) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SyncV1ServiceSyncStreamStreamMessage) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SyncV1ServiceSyncStreamStreamMessage) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *SyncV1ServiceSyncStreamStreamMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *SyncV1ServiceSyncStreamStreamMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SyncV1ServiceSyncStreamStreamMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetSid

`func (o *SyncV1ServiceSyncStreamStreamMessage) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *SyncV1ServiceSyncStreamStreamMessage) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *SyncV1ServiceSyncStreamStreamMessage) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *SyncV1ServiceSyncStreamStreamMessage) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *SyncV1ServiceSyncStreamStreamMessage) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *SyncV1ServiceSyncStreamStreamMessage) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


