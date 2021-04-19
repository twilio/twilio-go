# ListCommandResponse

## Properties

Name | Type | Description
------------ | ------------- | -------------
**Commands** | Pointer to [**[]WirelessV1Command**](WirelessV1Command.md) |  | [optional] 
**Meta** | Pointer to [**ListCommandResponseMeta**](ListCommandResponse_meta.md) |  | [optional] 

## Methods

### NewListCommandResponse

`func NewListCommandResponse() *ListCommandResponse`

NewListCommandResponse instantiates a new ListCommandResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCommandResponseWithDefaults

`func NewListCommandResponseWithDefaults() *ListCommandResponse`

NewListCommandResponseWithDefaults instantiates a new ListCommandResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommands

`func (o *ListCommandResponse) GetCommands() []WirelessV1Command`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *ListCommandResponse) GetCommandsOk() (*[]WirelessV1Command, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *ListCommandResponse) SetCommands(v []WirelessV1Command)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *ListCommandResponse) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### GetMeta

`func (o *ListCommandResponse) GetMeta() ListCommandResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCommandResponse) GetMetaOk() (*ListCommandResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCommandResponse) SetMeta(v ListCommandResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCommandResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


