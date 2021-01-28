# UpdateChannelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | **string** | A valid JSON string that contains application-specific data. | [optional] 
**CreatedBy** | **string** | The &#x60;identity&#x60; of the User that created the channel. Default is: &#x60;system&#x60;. | [optional] 
**DateCreated** | [**time.Time**](time.Time.md) | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was created. The default value is the current time set by the Chat service.  Note that this should only be used in cases where a Channel is being recreated from a backup/separate source. | [optional] 
**DateUpdated** | [**time.Time**](time.Time.md) | The date, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format, to assign to the resource as the date it was last updated. | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It can be up to 256 characters long. | [optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used to address the resource in place of the resource&#39;s &#x60;sid&#x60; in the URL. This value must be 256 characters or less in length and unique within the Service. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


