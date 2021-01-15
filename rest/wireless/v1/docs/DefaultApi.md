# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommand**](DefaultApi.md#CreateCommand) | **Post** /v1/Commands | 
[**CreateRatePlan**](DefaultApi.md#CreateRatePlan) | **Post** /v1/RatePlans | 
[**DeleteCommand**](DefaultApi.md#DeleteCommand) | **Delete** /v1/Commands/{Sid} | 
[**DeleteRatePlan**](DefaultApi.md#DeleteRatePlan) | **Delete** /v1/RatePlans/{Sid} | 
[**DeleteSim**](DefaultApi.md#DeleteSim) | **Delete** /v1/Sims/{Sid} | 
[**FetchCommand**](DefaultApi.md#FetchCommand) | **Get** /v1/Commands/{Sid} | 
[**FetchRatePlan**](DefaultApi.md#FetchRatePlan) | **Get** /v1/RatePlans/{Sid} | 
[**FetchSim**](DefaultApi.md#FetchSim) | **Get** /v1/Sims/{Sid} | 
[**ListAccountUsageRecord**](DefaultApi.md#ListAccountUsageRecord) | **Get** /v1/UsageRecords | 
[**ListCommand**](DefaultApi.md#ListCommand) | **Get** /v1/Commands | 
[**ListDataSession**](DefaultApi.md#ListDataSession) | **Get** /v1/Sims/{SimSid}/DataSessions | 
[**ListRatePlan**](DefaultApi.md#ListRatePlan) | **Get** /v1/RatePlans | 
[**ListSim**](DefaultApi.md#ListSim) | **Get** /v1/Sims | 
[**ListUsageRecord**](DefaultApi.md#ListUsageRecord) | **Get** /v1/Sims/{SimSid}/UsageRecords | 
[**UpdateRatePlan**](DefaultApi.md#UpdateRatePlan) | **Post** /v1/RatePlans/{Sid} | 
[**UpdateSim**](DefaultApi.md#UpdateSim) | **Post** /v1/Sims/{Sid} | 



## CreateCommand

> WirelessV1Command CreateCommand(ctx, optional)



Send a Command to a Sim.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateCommandOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **callbackMethod** | **optional.String**| The HTTP method we use to call &#x60;callback_url&#x60;. Can be: &#x60;POST&#x60; or &#x60;GET&#x60;, and the default is &#x60;POST&#x60;. | 
 **callbackUrl** | **optional.String**| The URL we call using the &#x60;callback_url&#x60; when the Command has finished sending, whether the command was delivered or it failed. | 
 **command** | **optional.String**| The message body of the Command. Can be plain text in text mode or a Base64 encoded byte string in binary mode. | 
 **commandMode** | **optional.String**| The mode to use when sending the SMS message. Can be: &#x60;text&#x60; or &#x60;binary&#x60;. The default SMS mode is &#x60;text&#x60;. | 
 **deliveryReceiptRequested** | **optional.Bool**| Whether to request delivery receipt from the recipient. For Commands that request delivery receipt, the Command state transitions to &#39;delivered&#39; once the server has received a delivery receipt from the device. The default value is &#x60;true&#x60;. | 
 **includeSid** | **optional.String**| Whether to include the SID of the command in the message body. Can be: &#x60;none&#x60;, &#x60;start&#x60;, or &#x60;end&#x60;, and the default behavior is &#x60;none&#x60;. When sending a Command to a SIM in text mode, we can automatically include the SID of the Command in the message body, which could be used to ensure that the device does not process the same Command more than once.  A value of &#x60;start&#x60; will prepend the message with the Command SID, and &#x60;end&#x60; will append it to the end, separating the Command SID from the message body with a space. The length of the Command SID is included in the 160 character limit so the SMS body must be 128 characters or less before the Command SID is included. | 
 **sim** | **optional.String**| The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the Command to. | 

### Return type

[**WirelessV1Command**](wireless.v1.command.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRatePlan

> WirelessV1RatePlan CreateRatePlan(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateRatePlanOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateRatePlanOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataEnabled** | **optional.Bool**| Whether SIMs can use GPRS/3G/4G/LTE data connectivity. | 
 **dataLimit** | **optional.Int32**| The total data usage (download and upload combined) in Megabytes that the Network allows during one month on the home network (T-Mobile USA). The metering period begins the day of activation and ends on the same day in the following month. Can be up to 2TB and the default value is &#x60;1000&#x60;. | 
 **dataMetering** | **optional.String**| The model used to meter data usage. Can be: &#x60;payg&#x60; and &#x60;quota-1&#x60;, &#x60;quota-10&#x60;, and &#x60;quota-50&#x60;. Learn more about the available [data metering models](https://www.twilio.com/docs/wireless/api/rateplan-resource#payg-vs-quota-data-plans). | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It does not have to be unique. | 
 **internationalRoaming** | [**optional.Interface of []string**](string.md)| The list of services that SIMs capable of using GPRS/3G/4G/LTE data connectivity can use outside of the United States. Can be: &#x60;data&#x60;, &#x60;voice&#x60;, and &#x60;messaging&#x60;. | 
 **internationalRoamingDataLimit** | **optional.Int32**| The total data usage (download and upload combined) in Megabytes that the Network allows during one month when roaming outside the United States. Can be up to 2TB. | 
 **messagingEnabled** | **optional.Bool**| Whether SIMs can make, send, and receive SMS using [Commands](https://www.twilio.com/docs/wireless/api/command-resource). | 
 **nationalRoamingDataLimit** | **optional.Int32**| The total data usage (download and upload combined) in Megabytes that the Network allows during one month on non-home networks in the United States. The metering period begins the day of activation and ends on the same day in the following month. Can be up to 2TB. See [national roaming](https://www.twilio.com/docs/wireless/api/rateplan-resource#national-roaming) for more info. | 
 **nationalRoamingEnabled** | **optional.Bool**| Whether SIMs can roam on networks other than the home network (T-Mobile USA) in the United States. See [national roaming](https://www.twilio.com/docs/wireless/api/rateplan-resource#national-roaming). | 
 **uniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource. | 
 **voiceEnabled** | **optional.Bool**| Whether SIMs can make and receive voice calls. | 

### Return type

[**WirelessV1RatePlan**](wireless.v1.rate_plan.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommand

> DeleteCommand(ctx, sid)



Delete a Command instance from your account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Command resource to delete. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRatePlan

> DeleteRatePlan(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the RatePlan resource to delete. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSim

> DeleteSim(ctx, sid)



Delete a Sim resource on your Account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID or the &#x60;unique_name&#x60; of the Sim resource to delete. | 

### Return type

 (empty response body)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchCommand

> WirelessV1Command FetchCommand(ctx, sid)



Fetch a Command instance from your account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the Command resource to fetch. | 

### Return type

[**WirelessV1Command**](wireless.v1.command.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchRatePlan

> WirelessV1RatePlan FetchRatePlan(ctx, sid)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the RatePlan resource to fetch. | 

### Return type

[**WirelessV1RatePlan**](wireless.v1.rate_plan.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSim

> WirelessV1Sim FetchSim(ctx, sid)



Fetch a Sim resource on your Account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID or the &#x60;unique_name&#x60; of the Sim resource to fetch. | 

### Return type

[**WirelessV1Sim**](wireless.v1.sim.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountUsageRecord

> WirelessV1AccountUsageRecordReadResponse ListAccountUsageRecord(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListAccountUsageRecordOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAccountUsageRecordOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **end** | **optional.Time**| Only include usage that has occurred on or before this date. Format is [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html). | 
 **start** | **optional.Time**| Only include usage that has occurred on or after this date. Format is [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html). | 
 **granularity** | **optional.String**| How to summarize the usage by time. Can be: &#x60;daily&#x60;, &#x60;hourly&#x60;, or &#x60;all&#x60;. A value of &#x60;all&#x60; returns one Usage Record that describes the usage for the entire period. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**WirelessV1AccountUsageRecordReadResponse**](wireless_v1_account_usage_recordReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommand

> WirelessV1CommandReadResponse ListCommand(ctx, optional)



Retrieve a list of Commands from your account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListCommandOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCommandOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sim** | **optional.String**| The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the [Sim resources](https://www.twilio.com/docs/wireless/api/sim-resource) to read. | 
 **status** | **optional.String**| The status of the resources to read. Can be: &#x60;queued&#x60;, &#x60;sent&#x60;, &#x60;delivered&#x60;, &#x60;received&#x60;, or &#x60;failed&#x60;. | 
 **direction** | **optional.String**| Only return Commands with this direction value. | 
 **transport** | **optional.String**| Only return Commands with this transport value. Can be: &#x60;sms&#x60; or &#x60;ip&#x60;. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**WirelessV1CommandReadResponse**](wireless_v1_commandReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataSession

> WirelessV1SimDataSessionReadResponse ListDataSession(ctx, simSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simSid** | **string**| The SID of the [Sim resource](https://www.twilio.com/docs/wireless/api/sim-resource) with the Data Sessions to read. | 
 **optional** | ***ListDataSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDataSessionOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**WirelessV1SimDataSessionReadResponse**](wireless_v1_sim_data_sessionReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRatePlan

> WirelessV1RatePlanReadResponse ListRatePlan(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListRatePlanOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListRatePlanOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**WirelessV1RatePlanReadResponse**](wireless_v1_rate_planReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSim

> WirelessV1SimReadResponse ListSim(ctx, optional)



Retrieve a list of Sim resources on your Account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSimOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSimOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **optional.String**| Only return Sim resources with this status. | 
 **iccid** | **optional.String**| Only return Sim resources with this ICCID. This will return a list with a maximum size of 1. | 
 **ratePlan** | **optional.String**| The SID or unique name of a [RatePlan resource](https://www.twilio.com/docs/wireless/api/rateplan-resource). Only return Sim resources assigned to this RatePlan resource. | 
 **eId** | **optional.String**| Deprecated. | 
 **simRegistrationCode** | **optional.String**| Only return Sim resources with this registration code. This will return a list with a maximum size of 1. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**WirelessV1SimReadResponse**](wireless_v1_simReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecord

> WirelessV1SimUsageRecordReadResponse ListUsageRecord(ctx, simSid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simSid** | **string**| The SID of the [Sim resource](https://www.twilio.com/docs/wireless/api/sim-resource)  to read the usage from. | 
 **optional** | ***ListUsageRecordOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsageRecordOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **end** | **optional.Time**| Only include usage that occurred on or before this date, specified in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html). The default is the current time. | 
 **start** | **optional.Time**| Only include usage that has occurred on or after this date, specified in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html). The default is one month before the &#x60;end&#x60; parameter value. | 
 **granularity** | **optional.String**| How to summarize the usage by time. Can be: &#x60;daily&#x60;, &#x60;hourly&#x60;, or &#x60;all&#x60;. The default is &#x60;all&#x60;. A value of &#x60;all&#x60; returns one Usage Record that describes the usage for the entire period. | 
 **pageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**WirelessV1SimUsageRecordReadResponse**](wireless_v1_sim_usage_recordReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRatePlan

> WirelessV1RatePlan UpdateRatePlan(ctx, sid, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID of the RatePlan resource to update. | 
 **optional** | ***UpdateRatePlanOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateRatePlanOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **friendlyName** | **optional.String**| A descriptive string that you create to describe the resource. It does not have to be unique. | 
 **uniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource. | 

### Return type

[**WirelessV1RatePlan**](wireless.v1.rate_plan.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSim

> WirelessV1Sim UpdateSim(ctx, sid, optional)



Updates the given properties of a Sim resource on your Account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string**| The SID or the &#x60;unique_name&#x60; of the Sim resource to update. | 
 **optional** | ***UpdateSimOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSimOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountSid** | **optional.String**| The SID of the [Account](https://www.twilio.com/docs/iam/api/account) to which the Sim resource should belong. The Account SID can only be that of the requesting Account or that of a [Subaccount](https://www.twilio.com/docs/iam/api/subaccounts) of the requesting Account. Only valid when the Sim resource&#39;s status is &#x60;new&#x60;. For more information, see the [Move SIMs between Subaccounts documentation](https://www.twilio.com/docs/wireless/api/sim-resource#move-sims-between-subaccounts). | 
 **callbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;POST&#x60; or &#x60;GET&#x60;. The default is &#x60;POST&#x60;. | 
 **callbackUrl** | **optional.String**| The URL we should call using the &#x60;callback_url&#x60; when the SIM has finished updating. When the SIM transitions from &#x60;new&#x60; to &#x60;ready&#x60; or from any status to &#x60;deactivated&#x60;, we call this URL when the status changes to an intermediate status (&#x60;ready&#x60; or &#x60;deactivated&#x60;) and again when the status changes to its final status (&#x60;active&#x60; or &#x60;canceled&#x60;). | 
 **commandsCallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;commands_callback_url&#x60;. Can be: &#x60;POST&#x60; or &#x60;GET&#x60;. The default is &#x60;POST&#x60;. | 
 **commandsCallbackUrl** | **optional.String**| The URL we should call using the &#x60;commands_callback_method&#x60; when the SIM sends a [Command](https://www.twilio.com/docs/wireless/api/command-resource). Your server should respond with an HTTP status code in the 200 range; any response body is ignored. | 
 **friendlyName** | **optional.String**| A descriptive string that you create to describe the Sim resource. It does not need to be unique. | 
 **ratePlan** | **optional.String**| The SID or unique name of the [RatePlan resource](https://www.twilio.com/docs/wireless/api/rateplan-resource) to which the Sim resource should be assigned. | 
 **resetStatus** | **optional.String**| Initiate a connectivity reset on the SIM. Set to &#x60;resetting&#x60; to initiate a connectivity reset on the SIM. No other value is valid. | 
 **smsFallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;sms_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. Default is &#x60;POST&#x60;. | 
 **smsFallbackUrl** | **optional.String**| The URL we should call using the &#x60;sms_fallback_method&#x60; when an error occurs while retrieving or executing the TwiML requested from &#x60;sms_url&#x60;. | 
 **smsMethod** | **optional.String**| The HTTP method we should use to call &#x60;sms_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. Default is &#x60;POST&#x60;. | 
 **smsUrl** | **optional.String**| The URL we should call using the &#x60;sms_method&#x60; when the SIM-connected device sends an SMS message that is not a [Command](https://www.twilio.com/docs/wireless/api/command-resource). | 
 **status** | **optional.String**| The new status of the Sim resource. Can be: &#x60;ready&#x60;, &#x60;active&#x60;, &#x60;suspended&#x60;, or &#x60;deactivated&#x60;. | 
 **uniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. It can be used in place of the &#x60;sid&#x60; in the URL path to address the resource. | 
 **voiceFallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;voice_fallback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **voiceFallbackUrl** | **optional.String**| The URL we should call using the &#x60;voice_fallback_method&#x60; when an error occurs while retrieving or executing the TwiML requested from &#x60;voice_url&#x60;. | 
 **voiceMethod** | **optional.String**| The HTTP method we should use when we call &#x60;voice_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60;. | 
 **voiceUrl** | **optional.String**| The URL we should call using the &#x60;voice_method&#x60; when the SIM-connected device makes a voice call. | 

### Return type

[**WirelessV1Sim**](wireless.v1.sim.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

