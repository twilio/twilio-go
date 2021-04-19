# ApiV2010AccountUsageUsageRecordUsageRecordThisMonth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **NullableString** | The SID of the Account accrued the usage | [optional] 
**ApiVersion** | Pointer to **NullableString** | The API version used to create the resource | [optional] 
**AsOf** | Pointer to **NullableString** | Usage records up to date as of this timestamp | [optional] 
**Category** | Pointer to **NullableString** | The category of usage | [optional] 
**Count** | Pointer to **NullableString** | The number of usage events | [optional] 
**CountUnit** | Pointer to **NullableString** | The units in which count is measured | [optional] 
**Description** | Pointer to **NullableString** | A plain-language description of the usage category | [optional] 
**EndDate** | Pointer to **NullableString** | The last date for which usage is included in the UsageRecord | [optional] 
**Price** | Pointer to **NullableFloat32** | The total price of the usage | [optional] 
**PriceUnit** | Pointer to **NullableString** | The currency in which &#x60;price&#x60; is measured | [optional] 
**StartDate** | Pointer to **NullableString** | The first date for which usage is included in this UsageRecord | [optional] 
**SubresourceUris** | Pointer to **map[string]interface{}** | A list of related resources identified by their relative URIs | [optional] 
**Uri** | Pointer to **NullableString** | The URI of the resource, relative to &#x60;https://api.twilio.com&#x60; | [optional] 
**Usage** | Pointer to **NullableString** | The amount of usage | [optional] 
**UsageUnit** | Pointer to **NullableString** | The units in which usage is measured | [optional] 

## Methods

### NewApiV2010AccountUsageUsageRecordUsageRecordThisMonth

`func NewApiV2010AccountUsageUsageRecordUsageRecordThisMonth() *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth`

NewApiV2010AccountUsageUsageRecordUsageRecordThisMonth instantiates a new ApiV2010AccountUsageUsageRecordUsageRecordThisMonth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV2010AccountUsageUsageRecordUsageRecordThisMonthWithDefaults

`func NewApiV2010AccountUsageUsageRecordUsageRecordThisMonthWithDefaults() *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth`

NewApiV2010AccountUsageUsageRecordUsageRecordThisMonthWithDefaults instantiates a new ApiV2010AccountUsageUsageRecordUsageRecordThisMonth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### SetAccountSidNil

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetAccountSidNil(b bool)`

 SetAccountSidNil sets the value for AccountSid to be an explicit nil

### UnsetAccountSid
`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) UnsetAccountSid()`

UnsetAccountSid ensures that no value is present for AccountSid, not even an explicit nil
### GetApiVersion

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### SetApiVersionNil

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetApiVersionNil(b bool)`

 SetApiVersionNil sets the value for ApiVersion to be an explicit nil

### UnsetApiVersion
`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) UnsetApiVersion()`

UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
### GetAsOf

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### SetAsOfNil

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetAsOfNil(b bool)`

 SetAsOfNil sets the value for AsOf to be an explicit nil

### UnsetAsOf
`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) UnsetAsOf()`

UnsetAsOf ensures that no value is present for AsOf, not even an explicit nil
### GetCategory

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetCount

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetCount() string`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetCountOk() (*string, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetCount(v string)`

SetCount sets Count field to given value.

### HasCount

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetCountUnit

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetCountUnit() string`

GetCountUnit returns the CountUnit field if non-nil, zero value otherwise.

### GetCountUnitOk

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetCountUnitOk() (*string, bool)`

GetCountUnitOk returns a tuple with the CountUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountUnit

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetCountUnit(v string)`

SetCountUnit sets CountUnit field to given value.

### HasCountUnit

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) HasCountUnit() bool`

HasCountUnit returns a boolean if a field has been set.

### SetCountUnitNil

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetCountUnitNil(b bool)`

 SetCountUnitNil sets the value for CountUnit to be an explicit nil

### UnsetCountUnit
`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) UnsetCountUnit()`

UnsetCountUnit ensures that no value is present for CountUnit, not even an explicit nil
### GetDescription

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndDate

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetPrice

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetPriceUnit

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### SetPriceUnitNil

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetPriceUnitNil(b bool)`

 SetPriceUnitNil sets the value for PriceUnit to be an explicit nil

### UnsetPriceUnit
`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) UnsetPriceUnit()`

UnsetPriceUnit ensures that no value is present for PriceUnit, not even an explicit nil
### GetStartDate

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetSubresourceUris

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetSubresourceUris() map[string]interface{}`

GetSubresourceUris returns the SubresourceUris field if non-nil, zero value otherwise.

### GetSubresourceUrisOk

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetSubresourceUrisOk() (*map[string]interface{}, bool)`

GetSubresourceUrisOk returns a tuple with the SubresourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresourceUris

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetSubresourceUris(v map[string]interface{})`

SetSubresourceUris sets SubresourceUris field to given value.

### HasSubresourceUris

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) HasSubresourceUris() bool`

HasSubresourceUris returns a boolean if a field has been set.

### SetSubresourceUrisNil

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetSubresourceUrisNil(b bool)`

 SetSubresourceUrisNil sets the value for SubresourceUris to be an explicit nil

### UnsetSubresourceUris
`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) UnsetSubresourceUris()`

UnsetSubresourceUris ensures that no value is present for SubresourceUris, not even an explicit nil
### GetUri

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetUsage

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil
### GetUsageUnit

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetUsageUnit() string`

GetUsageUnit returns the UsageUnit field if non-nil, zero value otherwise.

### GetUsageUnitOk

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) GetUsageUnitOk() (*string, bool)`

GetUsageUnitOk returns a tuple with the UsageUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageUnit

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetUsageUnit(v string)`

SetUsageUnit sets UsageUnit field to given value.

### HasUsageUnit

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) HasUsageUnit() bool`

HasUsageUnit returns a boolean if a field has been set.

### SetUsageUnitNil

`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) SetUsageUnitNil(b bool)`

 SetUsageUnitNil sets the value for UsageUnit to be an explicit nil

### UnsetUsageUnit
`func (o *ApiV2010AccountUsageUsageRecordUsageRecordThisMonth) UnsetUsageUnit()`

UnsetUsageUnit ensures that no value is present for UsageUnit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


