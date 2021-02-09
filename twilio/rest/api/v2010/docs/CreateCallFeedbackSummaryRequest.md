# CreateCallFeedbackSummaryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | [**time.Time**](time.Time.md) | Only include feedback given on or before this date. Format is &#x60;YYYY-MM-DD&#x60; and specified in UTC. | 
**IncludeSubaccounts** | **bool** | Whether to also include Feedback resources from all subaccounts. &#x60;true&#x60; includes feedback from all subaccounts and &#x60;false&#x60;, the default, includes feedback from only the specified account. | [optional] 
**StartDate** | [**time.Time**](time.Time.md) | Only include feedback given on or after this date. Format is &#x60;YYYY-MM-DD&#x60; and specified in UTC. | 
**StatusCallback** | **string** | The URL that we will request when the feedback summary is complete. | [optional] 
**StatusCallbackMethod** | **string** | The HTTP method (&#x60;GET&#x60; or &#x60;POST&#x60;) we use to make the request to the &#x60;StatusCallback&#x60; URL. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


