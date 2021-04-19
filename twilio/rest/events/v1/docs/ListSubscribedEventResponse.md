# ListSubscribedEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListVersionResponseMeta**](ListVersionResponse_meta.md) |  | [optional] 
**Types** | Pointer to [**[]EventsV1SubscriptionSubscribedEvent**](EventsV1SubscriptionSubscribedEvent.md) |  | [optional] 

## Methods

### NewListSubscribedEventResponse

`func NewListSubscribedEventResponse() *ListSubscribedEventResponse`

NewListSubscribedEventResponse instantiates a new ListSubscribedEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubscribedEventResponseWithDefaults

`func NewListSubscribedEventResponseWithDefaults() *ListSubscribedEventResponse`

NewListSubscribedEventResponseWithDefaults instantiates a new ListSubscribedEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListSubscribedEventResponse) GetMeta() ListVersionResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSubscribedEventResponse) GetMetaOk() (*ListVersionResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSubscribedEventResponse) SetMeta(v ListVersionResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSubscribedEventResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetTypes

`func (o *ListSubscribedEventResponse) GetTypes() []EventsV1SubscriptionSubscribedEvent`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *ListSubscribedEventResponse) GetTypesOk() (*[]EventsV1SubscriptionSubscribedEvent, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *ListSubscribedEventResponse) SetTypes(v []EventsV1SubscriptionSubscribedEvent)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *ListSubscribedEventResponse) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


