# ListItemAssignmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ListBundleResponseMeta**](ListBundleResponse_meta.md) |  | [optional] 
**Results** | Pointer to [**[]NumbersV2RegulatoryComplianceBundleItemAssignment**](NumbersV2RegulatoryComplianceBundleItemAssignment.md) |  | [optional] 

## Methods

### NewListItemAssignmentResponse

`func NewListItemAssignmentResponse() *ListItemAssignmentResponse`

NewListItemAssignmentResponse instantiates a new ListItemAssignmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListItemAssignmentResponseWithDefaults

`func NewListItemAssignmentResponseWithDefaults() *ListItemAssignmentResponse`

NewListItemAssignmentResponseWithDefaults instantiates a new ListItemAssignmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ListItemAssignmentResponse) GetMeta() ListBundleResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListItemAssignmentResponse) GetMetaOk() (*ListBundleResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListItemAssignmentResponse) SetMeta(v ListBundleResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListItemAssignmentResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetResults

`func (o *ListItemAssignmentResponse) GetResults() []NumbersV2RegulatoryComplianceBundleItemAssignment`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListItemAssignmentResponse) GetResultsOk() (*[]NumbersV2RegulatoryComplianceBundleItemAssignment, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListItemAssignmentResponse) SetResults(v []NumbersV2RegulatoryComplianceBundleItemAssignment)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListItemAssignmentResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


