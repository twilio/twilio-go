# ListSmsCommandResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListCommandResponseMeta**](ListCommandResponse_meta.md) |  | [optional] 
**SmsCommands** | Pointer to [**[]SupersimV1SmsCommand**](SupersimV1SmsCommand.md) |  | [optional] 

## Methods

### NewListSmsCommandResponse

`func NewListSmsCommandResponse() *ListSmsCommandResponse`

NewListSmsCommandResponse instantiates a new ListSmsCommandResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSmsCommandResponseWithDefaults

`func NewListSmsCommandResponseWithDefaults() *ListSmsCommandResponse`

NewListSmsCommandResponseWithDefaults instantiates a new ListSmsCommandResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListSmsCommandResponse) GetMeta() ListCommandResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSmsCommandResponse) GetMetaOk() (*ListCommandResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSmsCommandResponse) SetMeta(v ListCommandResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSmsCommandResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSmsCommands

`func (o *ListSmsCommandResponse) GetSmsCommands() []SupersimV1SmsCommand`

GetSmsCommands returns the SmsCommands field if non-nil, zero value otherwise.

### GetSmsCommandsOk

`func (o *ListSmsCommandResponse) GetSmsCommandsOk() (*[]SupersimV1SmsCommand, bool)`

GetSmsCommandsOk returns a tuple with the SmsCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsCommands

`func (o *ListSmsCommandResponse) SetSmsCommands(v []SupersimV1SmsCommand)`

SetSmsCommands sets SmsCommands field to given value.

### HasSmsCommands

`func (o *ListSmsCommandResponse) HasSmsCommands() bool`

HasSmsCommands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


