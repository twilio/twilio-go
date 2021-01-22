# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommand**](DefaultApi.md#CreateCommand) | **Post** /v1/Commands | 
[**CreateFleet**](DefaultApi.md#CreateFleet) | **Post** /v1/Fleets | 
[**CreateNetworkAccessProfile**](DefaultApi.md#CreateNetworkAccessProfile) | **Post** /v1/NetworkAccessProfiles | 
[**CreateNetworkAccessProfileNetwork**](DefaultApi.md#CreateNetworkAccessProfileNetwork) | **Post** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks | 
[**DeleteNetworkAccessProfileNetwork**](DefaultApi.md#DeleteNetworkAccessProfileNetwork) | **Delete** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks/{Sid} | 
[**FetchCommand**](DefaultApi.md#FetchCommand) | **Get** /v1/Commands/{Sid} | 
[**FetchFleet**](DefaultApi.md#FetchFleet) | **Get** /v1/Fleets/{Sid} | 
[**FetchNetwork**](DefaultApi.md#FetchNetwork) | **Get** /v1/Networks/{Sid} | 
[**FetchNetworkAccessProfile**](DefaultApi.md#FetchNetworkAccessProfile) | **Get** /v1/NetworkAccessProfiles/{Sid} | 
[**FetchNetworkAccessProfileNetwork**](DefaultApi.md#FetchNetworkAccessProfileNetwork) | **Get** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks/{Sid} | 
[**FetchSim**](DefaultApi.md#FetchSim) | **Get** /v1/Sims/{Sid} | 
[**ListCommand**](DefaultApi.md#ListCommand) | **Get** /v1/Commands | 
[**ListFleet**](DefaultApi.md#ListFleet) | **Get** /v1/Fleets | 
[**ListNetwork**](DefaultApi.md#ListNetwork) | **Get** /v1/Networks | 
[**ListNetworkAccessProfile**](DefaultApi.md#ListNetworkAccessProfile) | **Get** /v1/NetworkAccessProfiles | 
[**ListNetworkAccessProfileNetwork**](DefaultApi.md#ListNetworkAccessProfileNetwork) | **Get** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks | 
[**ListSim**](DefaultApi.md#ListSim) | **Get** /v1/Sims | 
[**ListUsageRecord**](DefaultApi.md#ListUsageRecord) | **Get** /v1/UsageRecords | 
[**UpdateFleet**](DefaultApi.md#UpdateFleet) | **Post** /v1/Fleets/{Sid} | 
[**UpdateNetworkAccessProfile**](DefaultApi.md#UpdateNetworkAccessProfile) | **Post** /v1/NetworkAccessProfiles/{Sid} | 
[**UpdateSim**](DefaultApi.md#UpdateSim) | **Post** /v1/Sims/{Sid} | 



## CreateCommand

> SupersimV1Command CreateCommand(ctx, optional)



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
 **CallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is POST. | 
 **CallbackUrl** | **optional.String**| The URL we should call using the &#x60;callback_method&#x60; after we have sent the command. | 
 **Command** | **optional.String**| The message body of the command. | 
 **Sim** | **optional.String**| The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the Command to. | 

### Return type

[**SupersimV1Command**](supersim.v1.command.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFleet

> SupersimV1Fleet CreateFleet(ctx, optional)



Create a Fleet

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateFleetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateFleetOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **CommandsEnabled** | **optional.Bool**| Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands. Defaults to &#x60;true&#x60;. | 
 **CommandsMethod** | **optional.String**| A string representing the HTTP method to use when making a request to &#x60;commands_url&#x60;. Can be one of &#x60;POST&#x60; or &#x60;GET&#x60;. Defaults to &#x60;POST&#x60;. | 
 **CommandsUrl** | **optional.String**| The URL that will receive a webhook when a SIM in the Fleet originates a machine-to-machine SMS via Commands. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored. | 
 **DataEnabled** | **optional.Bool**| Defines whether SIMs in the Fleet are capable of using 2G/3G/4G/LTE/CAT-M data connectivity. Defaults to &#x60;true&#x60;. | 
 **DataLimit** | **optional.Int32**| The total data usage (download and upload combined) in Megabytes that each Sim resource assigned to the Fleet resource can consume during a billing period (normally one month). Value must be between 1MB (1) and 2TB (2,000,000). Defaults to 1GB (1,000). | 
 **NetworkAccessProfile** | **optional.String**| The SID or unique name of the Network Access Profile that will control which cellular networks the Fleet&#39;s SIMs can connect to. | 
 **UniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource. | 

### Return type

[**SupersimV1Fleet**](supersim.v1.fleet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkAccessProfile

> SupersimV1NetworkAccessProfile CreateNetworkAccessProfile(ctx, optional)



Create a new Network Access Profile

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateNetworkAccessProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateNetworkAccessProfileOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Networks** | [**optional.Interface of []string**](string.md)| List of Network SIDs that this Network Access Profile will allow connections to. | 
 **UniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource. | 

### Return type

[**SupersimV1NetworkAccessProfile**](supersim.v1.network_access_profile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkAccessProfileNetwork

> SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork CreateNetworkAccessProfileNetwork(ctx, NetworkAccessProfileSid, optional)



Add a Network resource to the Network Access Profile resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**NetworkAccessProfileSid** | **string**| The unique string that identifies the Network Access Profile resource. | 
 **optional** | ***CreateNetworkAccessProfileNetworkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateNetworkAccessProfileNetworkOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **Network** | **optional.String**| The SID of the Network resource to be added to the Network Access Profile resource. | 

### Return type

[**SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork**](supersim.v1.network_access_profile.network_access_profile_network.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkAccessProfileNetwork

> DeleteNetworkAccessProfileNetwork(ctx, NetworkAccessProfileSid, Sid)



Remove a Network resource from the Network Access Profile resource's.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**NetworkAccessProfileSid** | **string**| The unique string that identifies the Network Access Profile resource. | 
**Sid** | **string**|  | 

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

> SupersimV1Command FetchCommand(ctx, Sid)



Fetch a Command instance from your account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Command resource to fetch. | 

### Return type

[**SupersimV1Command**](supersim.v1.command.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFleet

> SupersimV1Fleet FetchFleet(ctx, Sid)



Fetch a Fleet instance from your account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Fleet resource to fetch. | 

### Return type

[**SupersimV1Fleet**](supersim.v1.fleet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNetwork

> SupersimV1Network FetchNetwork(ctx, Sid)



Fetch a Network resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Network resource to fetch. | 

### Return type

[**SupersimV1Network**](supersim.v1.network.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNetworkAccessProfile

> SupersimV1NetworkAccessProfile FetchNetworkAccessProfile(ctx, Sid)



Fetch a Network Access Profile instance from your account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Network Access Profile resource to fetch. | 

### Return type

[**SupersimV1NetworkAccessProfile**](supersim.v1.network_access_profile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNetworkAccessProfileNetwork

> SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork FetchNetworkAccessProfileNetwork(ctx, NetworkAccessProfileSid, Sid)



Fetch a Network Access Profile resource's Network resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**NetworkAccessProfileSid** | **string**| The unique string that identifies the Network Access Profile resource. | 
**Sid** | **string**| The SID of the Network resource to fetch. | 

### Return type

[**SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork**](supersim.v1.network_access_profile.network_access_profile_network.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSim

> SupersimV1Sim FetchSim(ctx, Sid)



Fetch a Super SIM instance from your account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Sim resource to fetch. | 

### Return type

[**SupersimV1Sim**](supersim.v1.sim.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommand

> SupersimV1CommandReadResponse ListCommand(ctx, optional)



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
 **Sim** | **optional.String**| The SID or unique name of the Sim that Command was sent to or from. | 
 **Status** | **optional.String**| The status of the Command. Can be: &#x60;queued&#x60;, &#x60;sent&#x60;, &#x60;delivered&#x60;, &#x60;received&#x60; or &#x60;failed&#x60;. See the [Command Status Values](https://www.twilio.com/docs/wireless/api/command-resource#status-values) for a description of each. | 
 **Direction** | **optional.String**| The direction of the Command. Can be &#x60;to_sim&#x60; or &#x60;from_sim&#x60;. The value of &#x60;to_sim&#x60; is synonymous with the term &#x60;mobile terminated&#x60;, and &#x60;from_sim&#x60; is synonymous with the term &#x60;mobile originated&#x60;. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**SupersimV1CommandReadResponse**](supersim_v1_commandReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFleet

> SupersimV1FleetReadResponse ListFleet(ctx, optional)



Retrieve a list of Fleets from your account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListFleetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFleetOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **NetworkAccessProfile** | **optional.String**| The SID or unique name of the Network Access Profile that controls which cellular networks the Fleet&#39;s SIMs can connect to. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**SupersimV1FleetReadResponse**](supersim_v1_fleetReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetwork

> SupersimV1NetworkReadResponse ListNetwork(ctx, optional)



Retrieve a list of Network resources.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListNetworkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListNetworkOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **IsoCountry** | **optional.String**| The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the Network resources to read. | 
 **Mcc** | **optional.String**| The &#39;mobile country code&#39; of a country. Network resources with this &#x60;mcc&#x60; in their &#x60;identifiers&#x60; will be read. | 
 **Mnc** | **optional.String**| The &#39;mobile network code&#39; of a mobile operator network. Network resources with this &#x60;mnc&#x60; in their &#x60;identifiers&#x60; will be read. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**SupersimV1NetworkReadResponse**](supersim_v1_networkReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkAccessProfile

> SupersimV1NetworkAccessProfileReadResponse ListNetworkAccessProfile(ctx, optional)



Retrieve a list of Network Access Profiles from your account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListNetworkAccessProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListNetworkAccessProfileOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**SupersimV1NetworkAccessProfileReadResponse**](supersim_v1_network_access_profileReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkAccessProfileNetwork

> SupersimV1NetworkAccessProfileNetworkAccessProfileNetworkReadResponse ListNetworkAccessProfileNetwork(ctx, NetworkAccessProfileSid, optional)



Retrieve a list of Network Access Profile resource's Network resource.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**NetworkAccessProfileSid** | **string**| The unique string that identifies the Network Access Profile resource. | 
 **optional** | ***ListNetworkAccessProfileNetworkOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListNetworkAccessProfileNetworkOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**SupersimV1NetworkAccessProfileNetworkAccessProfileNetworkReadResponse**](supersim_v1_network_access_profile_network_access_profile_networkReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSim

> SupersimV1SimReadResponse ListSim(ctx, optional)



Retrieve a list of Super SIMs from your account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListSimOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSimOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Status** | **optional.String**| The status of the Sim resources to read. Can be &#x60;new&#x60;, &#x60;ready&#x60;, &#x60;active&#x60;, &#x60;inactive&#x60;, or &#x60;scheduled&#x60;. | 
 **Fleet** | **optional.String**| The SID or unique name of the Fleet to which a list of Sims are assigned. | 
 **Iccid** | **optional.String**| The [ICCID](https://en.wikipedia.org/wiki/Subscriber_identity_module#ICCID) associated with a Super SIM to filter the list by. Passing this parameter will always return a list containing zero or one SIMs. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**SupersimV1SimReadResponse**](supersim_v1_simReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecord

> SupersimV1UsageRecordReadResponse ListUsageRecord(ctx, optional)



List UsageRecords

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListUsageRecordOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsageRecordOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **Sim** | **optional.String**| SID or unique name of a Sim resource. Only show UsageRecords representing usage incurred by this Super SIM. | 
 **Fleet** | **optional.String**| SID or unique name of a Fleet resource. Only show UsageRecords representing usage for Super SIMs belonging to this Fleet resource at the time the usage occurred. | 
 **Network** | **optional.String**| SID of a Network resource. Only show UsageRecords representing usage on this network. | 
 **IsoCountry** | **optional.String**| Alpha-2 ISO Country Code. Only show UsageRecords representing usage in this country. | 
 **Group** | **optional.String**| Dimension over which to aggregate usage records. Can be: &#x60;sim&#x60;, &#x60;fleet&#x60;, &#x60;network&#x60;, &#x60;isoCountry&#x60;. Default is to not aggregate across any of these dimensions, UsageRecords will be aggregated into the time buckets described by the &#x60;Granularity&#x60; parameter. | 
 **Granularity** | **optional.String**| Time-based grouping that UsageRecords should be aggregated by. Can be: &#x60;hour&#x60;, &#x60;day&#x60;, or &#x60;all&#x60;. Default is &#x60;all&#x60;. &#x60;all&#x60; returns one UsageRecord that describes the usage for the entire period. | 
 **StartTime** | **optional.Time**| Only include usage that occurred at or after this time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Default is one month before the &#x60;end_time&#x60;. | 
 **EndTime** | **optional.Time**| Only include usage that occurred before this time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Default is the current time. | 
 **PageSize** | **optional.Int32**| How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**SupersimV1UsageRecordReadResponse**](supersim_v1_usage_recordReadResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFleet

> SupersimV1Fleet UpdateFleet(ctx, Sid, optional)



Updates the given properties of a Super SIM Fleet instance from your account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Fleet resource to update. | 
 **optional** | ***UpdateFleetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateFleetOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **NetworkAccessProfile** | **optional.String**| The SID or unique name of the Network Access Profile that will control which cellular networks the Fleet&#39;s SIMs can connect to. | 
 **UniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource. | 

### Return type

[**SupersimV1Fleet**](supersim.v1.fleet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkAccessProfile

> SupersimV1NetworkAccessProfile UpdateNetworkAccessProfile(ctx, Sid, optional)



Updates the given properties of a Network Access Profile in your account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Network Access Profile to update. | 
 **optional** | ***UpdateNetworkAccessProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateNetworkAccessProfileOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **UniqueName** | **optional.String**| The new unique name of the Network Access Profile. | 

### Return type

[**SupersimV1NetworkAccessProfile**](supersim.v1.network_access_profile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSim

> SupersimV1Sim UpdateSim(ctx, Sid, optional)



Updates the given properties of a Super SIM instance from your account.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string**| The SID of the Sim resource to update. | 
 **optional** | ***UpdateSimOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateSimOpts struct
 

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **AccountSid** | **optional.String**| The SID of the Account to which the Sim resource should belong. The Account SID can only be that of the requesting Account or that of a Subaccount of the requesting Account. Only valid when the Sim resource&#39;s status is new. | 
 **CallbackMethod** | **optional.String**| The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is POST. | 
 **CallbackUrl** | **optional.String**| The URL we should call using the &#x60;callback_method&#x60; after an asynchronous update has finished. | 
 **Fleet** | **optional.String**| The SID or unique name of the Fleet to which the SIM resource should be assigned. | 
 **Status** | **optional.String**| The new status of the resource. Can be: &#x60;ready&#x60;, &#x60;active&#x60;, or &#x60;inactive&#x60;. See the [Super SIM Status Values](https://www.twilio.com/docs/iot/supersim/api/sim-resource#status-values) for more info. | 
 **UniqueName** | **optional.String**| An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource. | 

### Return type

[**SupersimV1Sim**](supersim.v1.sim.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

