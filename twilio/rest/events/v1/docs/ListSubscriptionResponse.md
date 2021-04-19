# ListSubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListVersionResponseMeta**](ListVersionResponse_meta.md) |  | [optional] 
**Subscriptions** | Pointer to [**[]EventsV1Subscription**](EventsV1Subscription.md) |  | [optional] 

## Methods

### NewListSubscriptionResponse

`func NewListSubscriptionResponse() *ListSubscriptionResponse`

NewListSubscriptionResponse instantiates a new ListSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubscriptionResponseWithDefaults

`func NewListSubscriptionResponseWithDefaults() *ListSubscriptionResponse`

NewListSubscriptionResponseWithDefaults instantiates a new ListSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListSubscriptionResponse) GetMeta() ListVersionResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSubscriptionResponse) GetMetaOk() (*ListVersionResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSubscriptionResponse) SetMeta(v ListVersionResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSubscriptionResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSubscriptions

`func (o *ListSubscriptionResponse) GetSubscriptions() []EventsV1Subscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *ListSubscriptionResponse) GetSubscriptionsOk() (*[]EventsV1Subscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *ListSubscriptionResponse) SetSubscriptions(v []EventsV1Subscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *ListSubscriptionResponse) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


