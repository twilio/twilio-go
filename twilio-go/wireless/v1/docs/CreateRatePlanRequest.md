# CreateRatePlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataEnabled** | **bool** | Whether SIMs can use GPRS/3G/4G/LTE data connectivity. | [optional] 
**DataLimit** | **int32** | The total data usage (download and upload combined) in Megabytes that the Network allows during one month on the home network (T-Mobile USA). The metering period begins the day of activation and ends on the same day in the following month. Can be up to 2TB and the default value is &#x60;1000&#x60;. | [optional] 
**DataMetering** | **string** | The model used to meter data usage. Can be: &#x60;payg&#x60; and &#x60;quota-1&#x60;, &#x60;quota-10&#x60;, and &#x60;quota-50&#x60;. Learn more about the available [data metering models](https://www.twilio.com/docs/wireless/api/rateplan-resource#payg-vs-quota-data-plans). | [optional] 
**FriendlyName** | **string** | A descriptive string that you create to describe the resource. It does not have to be unique. | [optional] 
**InternationalRoaming** | **[]string** | The list of services that SIMs capable of using GPRS/3G/4G/LTE data connectivity can use outside of the United States. Can be: &#x60;data&#x60;, &#x60;voice&#x60;, and &#x60;messaging&#x60;. | [optional] 
**InternationalRoamingDataLimit** | **int32** | The total data usage (download and upload combined) in Megabytes that the Network allows during one month when roaming outside the United States. Can be up to 2TB. | [optional] 
**MessagingEnabled** | **bool** | Whether SIMs can make, send, and receive SMS using [Commands](https://www.twilio.com/docs/wireless/api/command-resource). | [optional] 
**NationalRoamingDataLimit** | **int32** | The total data usage (download and upload combined) in Megabytes that the Network allows during one month on non-home networks in the United States. The metering period begins the day of activation and ends on the same day in the following month. Can be up to 2TB. See [national roaming](https://www.twilio.com/docs/wireless/api/rateplan-resource#national-roaming) for more info. | [optional] 
**NationalRoamingEnabled** | **bool** | Whether SIMs can roam on networks other than the home network (T-Mobile USA) in the United States. See [national roaming](https://www.twilio.com/docs/wireless/api/rateplan-resource#national-roaming). | [optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource. | [optional] 
**VoiceEnabled** | **bool** | Whether SIMs can make and receive voice calls. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


