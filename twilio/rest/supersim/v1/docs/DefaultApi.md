# DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommand**](DefaultApi.md#CreateCommand) | **Post** /v1/Commands | 
[**CreateFleet**](DefaultApi.md#CreateFleet) | **Post** /v1/Fleets | 
[**CreateNetworkAccessProfile**](DefaultApi.md#CreateNetworkAccessProfile) | **Post** /v1/NetworkAccessProfiles | 
[**CreateNetworkAccessProfileNetwork**](DefaultApi.md#CreateNetworkAccessProfileNetwork) | **Post** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks | 
[**CreateSmsCommand**](DefaultApi.md#CreateSmsCommand) | **Post** /v1/SmsCommands | 
[**DeleteNetworkAccessProfileNetwork**](DefaultApi.md#DeleteNetworkAccessProfileNetwork) | **Delete** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks/{Sid} | 
[**FetchCommand**](DefaultApi.md#FetchCommand) | **Get** /v1/Commands/{Sid} | 
[**FetchFleet**](DefaultApi.md#FetchFleet) | **Get** /v1/Fleets/{Sid} | 
[**FetchNetwork**](DefaultApi.md#FetchNetwork) | **Get** /v1/Networks/{Sid} | 
[**FetchNetworkAccessProfile**](DefaultApi.md#FetchNetworkAccessProfile) | **Get** /v1/NetworkAccessProfiles/{Sid} | 
[**FetchNetworkAccessProfileNetwork**](DefaultApi.md#FetchNetworkAccessProfileNetwork) | **Get** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks/{Sid} | 
[**FetchSim**](DefaultApi.md#FetchSim) | **Get** /v1/Sims/{Sid} | 
[**FetchSmsCommand**](DefaultApi.md#FetchSmsCommand) | **Get** /v1/SmsCommands/{Sid} | 
[**ListCommand**](DefaultApi.md#ListCommand) | **Get** /v1/Commands | 
[**ListFleet**](DefaultApi.md#ListFleet) | **Get** /v1/Fleets | 
[**ListNetwork**](DefaultApi.md#ListNetwork) | **Get** /v1/Networks | 
[**ListNetworkAccessProfile**](DefaultApi.md#ListNetworkAccessProfile) | **Get** /v1/NetworkAccessProfiles | 
[**ListNetworkAccessProfileNetwork**](DefaultApi.md#ListNetworkAccessProfileNetwork) | **Get** /v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks | 
[**ListSim**](DefaultApi.md#ListSim) | **Get** /v1/Sims | 
[**ListSmsCommand**](DefaultApi.md#ListSmsCommand) | **Get** /v1/SmsCommands | 
[**ListUsageRecord**](DefaultApi.md#ListUsageRecord) | **Get** /v1/UsageRecords | 
[**UpdateFleet**](DefaultApi.md#UpdateFleet) | **Post** /v1/Fleets/{Sid} | 
[**UpdateNetworkAccessProfile**](DefaultApi.md#UpdateNetworkAccessProfile) | **Post** /v1/NetworkAccessProfiles/{Sid} | 
[**UpdateSim**](DefaultApi.md#UpdateSim) | **Post** /v1/Sims/{Sid} | 



## CreateCommand

> SupersimV1Command CreateCommand(ctx, optional)



Send a Command to a Sim.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateCommandParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CallbackMethod** | **string** | The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is POST. | 
**CallbackUrl** | **string** | The URL we should call using the &#x60;callback_method&#x60; after we have sent the command. | 
**Command** | **string** | The message body of the command. | 
**Sim** | **string** | The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the Command to. | 

### Return type

[**SupersimV1Command**](SupersimV1Command.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFleet

> SupersimV1Fleet CreateFleet(ctx, optional)



Create a Fleet

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateFleetParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CommandsEnabled** | **bool** | Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands. Defaults to &#x60;true&#x60;. | 
**CommandsMethod** | **string** | A string representing the HTTP method to use when making a request to &#x60;commands_url&#x60;. Can be one of &#x60;POST&#x60; or &#x60;GET&#x60;. Defaults to &#x60;POST&#x60;. | 
**CommandsUrl** | **string** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored. | 
**DataEnabled** | **bool** | Defines whether SIMs in the Fleet are capable of using 2G/3G/4G/LTE/CAT-M data connectivity. Defaults to &#x60;true&#x60;. | 
**DataLimit** | **int32** | The total data usage (download and upload combined) in Megabytes that each Sim resource assigned to the Fleet resource can consume during a billing period (normally one month). Value must be between 1MB (1) and 2TB (2,000,000). Defaults to 1GB (1,000). | 
**NetworkAccessProfile** | **string** | The SID or unique name of the Network Access Profile that will control which cellular networks the Fleet&#39;s SIMs can connect to. | 
**SmsCommandsEnabled** | **bool** | Defines whether SIMs in the Fleet are capable of sending and receiving machine-to-machine SMS via Commands. Defaults to &#x60;true&#x60;. | 
**SmsCommandsMethod** | **string** | A string representing the HTTP method to use when making a request to &#x60;sms_commands_url&#x60;. Can be one of &#x60;POST&#x60; or &#x60;GET&#x60;. Defaults to &#x60;POST&#x60;. | 
**SmsCommandsUrl** | **string** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the SMS Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored. | 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource. | 

### Return type

[**SupersimV1Fleet**](SupersimV1Fleet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkAccessProfile

> SupersimV1NetworkAccessProfile CreateNetworkAccessProfile(ctx, optional)



Create a new Network Access Profile

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateNetworkAccessProfileParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Networks** | **[]string** | List of Network SIDs that this Network Access Profile will allow connections to. | 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource. | 

### Return type

[**SupersimV1NetworkAccessProfile**](SupersimV1NetworkAccessProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkAccessProfileNetwork

> SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork CreateNetworkAccessProfileNetwork(ctx, NetworkAccessProfileSidoptional)



Add a Network resource to the Network Access Profile resource.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**NetworkAccessProfileSid** | **string** | The unique string that identifies the Network Access Profile resource. | 

### Other Parameters

Other parameters are passed through a pointer to a CreateNetworkAccessProfileNetworkParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Network** | **string** | The SID of the Network resource to be added to the Network Access Profile resource. | 

### Return type

[**SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork**](SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSmsCommand

> SupersimV1SmsCommand CreateSmsCommand(ctx, optional)



Send SMS Command to a Sim.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSmsCommandParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CallbackMethod** | **string** | The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is POST. | 
**CallbackUrl** | **string** | The URL we should call using the &#x60;callback_method&#x60; after we have sent the command. | 
**Payload** | **string** | The message body of the SMS Command. | 
**Sim** | **string** | The &#x60;sid&#x60; or &#x60;unique_name&#x60; of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the SMS Command to. | 

### Return type

[**SupersimV1SmsCommand**](SupersimV1SmsCommand.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkAccessProfileNetwork

> DeleteNetworkAccessProfileNetwork(ctx, NetworkAccessProfileSidSid)



Remove a Network resource from the Network Access Profile resource's.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**NetworkAccessProfileSid** | **string** | The unique string that identifies the Network Access Profile resource. | 
**Sid** | **string** | The SID of the Network resource to be removed from the Network Access Profile resource. | 

### Other Parameters

Other parameters are passed through a pointer to a DeleteNetworkAccessProfileNetworkParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Command resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchCommandParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**SupersimV1Command**](SupersimV1Command.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchFleet

> SupersimV1Fleet FetchFleet(ctx, Sid)



Fetch a Fleet instance from your account.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Fleet resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchFleetParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**SupersimV1Fleet**](SupersimV1Fleet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNetwork

> SupersimV1Network FetchNetwork(ctx, Sid)



Fetch a Network resource.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Network resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchNetworkParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**SupersimV1Network**](SupersimV1Network.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNetworkAccessProfile

> SupersimV1NetworkAccessProfile FetchNetworkAccessProfile(ctx, Sid)



Fetch a Network Access Profile instance from your account.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Network Access Profile resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchNetworkAccessProfileParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**SupersimV1NetworkAccessProfile**](SupersimV1NetworkAccessProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchNetworkAccessProfileNetwork

> SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork FetchNetworkAccessProfileNetwork(ctx, NetworkAccessProfileSidSid)



Fetch a Network Access Profile resource's Network resource.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**NetworkAccessProfileSid** | **string** | The unique string that identifies the Network Access Profile resource. | 
**Sid** | **string** | The SID of the Network resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchNetworkAccessProfileNetworkParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork**](SupersimV1NetworkAccessProfileNetworkAccessProfileNetwork.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSim

> SupersimV1Sim FetchSim(ctx, Sid)



Fetch a Super SIM instance from your account.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Sim resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchSimParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**SupersimV1Sim**](SupersimV1Sim.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSmsCommand

> SupersimV1SmsCommand FetchSmsCommand(ctx, Sid)



Fetch SMS Command instance from your account.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the SMS Command resource to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a FetchSmsCommandParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**SupersimV1SmsCommand**](SupersimV1SmsCommand.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommand

> ListCommandResponse ListCommand(ctx, optional)



Retrieve a list of Commands from your account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListCommandParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Sim** | **string** | The SID or unique name of the Sim that Command was sent to or from. | 
**Status** | **string** | The status of the Command. Can be: &#x60;queued&#x60;, &#x60;sent&#x60;, &#x60;delivered&#x60;, &#x60;received&#x60; or &#x60;failed&#x60;. See the [Command Status Values](https://www.twilio.com/docs/wireless/api/command-resource#status-values) for a description of each. | 
**Direction** | **string** | The direction of the Command. Can be &#x60;to_sim&#x60; or &#x60;from_sim&#x60;. The value of &#x60;to_sim&#x60; is synonymous with the term &#x60;mobile terminated&#x60;, and &#x60;from_sim&#x60; is synonymous with the term &#x60;mobile originated&#x60;. | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListCommandResponse**](ListCommandResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFleet

> ListFleetResponse ListFleet(ctx, optional)



Retrieve a list of Fleets from your account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListFleetParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**NetworkAccessProfile** | **string** | The SID or unique name of the Network Access Profile that controls which cellular networks the Fleet&#39;s SIMs can connect to. | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListFleetResponse**](ListFleetResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetwork

> ListNetworkResponse ListNetwork(ctx, optional)



Retrieve a list of Network resources.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListNetworkParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**IsoCountry** | **string** | The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the Network resources to read. | 
**Mcc** | **string** | The &#39;mobile country code&#39; of a country. Network resources with this &#x60;mcc&#x60; in their &#x60;identifiers&#x60; will be read. | 
**Mnc** | **string** | The &#39;mobile network code&#39; of a mobile operator network. Network resources with this &#x60;mnc&#x60; in their &#x60;identifiers&#x60; will be read. | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListNetworkResponse**](ListNetworkResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkAccessProfile

> ListNetworkAccessProfileResponse ListNetworkAccessProfile(ctx, optional)



Retrieve a list of Network Access Profiles from your account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListNetworkAccessProfileParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListNetworkAccessProfileResponse**](ListNetworkAccessProfileResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkAccessProfileNetwork

> ListNetworkAccessProfileNetworkResponse ListNetworkAccessProfileNetwork(ctx, NetworkAccessProfileSidoptional)



Retrieve a list of Network Access Profile resource's Network resource.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**NetworkAccessProfileSid** | **string** | The unique string that identifies the Network Access Profile resource. | 

### Other Parameters

Other parameters are passed through a pointer to a ListNetworkAccessProfileNetworkParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListNetworkAccessProfileNetworkResponse**](ListNetworkAccessProfileNetworkResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSim

> ListSimResponse ListSim(ctx, optional)



Retrieve a list of Super SIMs from your account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSimParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Status** | **string** | The status of the Sim resources to read. Can be &#x60;new&#x60;, &#x60;ready&#x60;, &#x60;active&#x60;, &#x60;inactive&#x60;, or &#x60;scheduled&#x60;. | 
**Fleet** | **string** | The SID or unique name of the Fleet to which a list of Sims are assigned. | 
**Iccid** | **string** | The [ICCID](https://en.wikipedia.org/wiki/Subscriber_identity_module#ICCID) associated with a Super SIM to filter the list by. Passing this parameter will always return a list containing zero or one SIMs. | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSimResponse**](ListSimResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSmsCommand

> ListSmsCommandResponse ListSmsCommand(ctx, optional)



Retrieve a list of SMS Commands from your account.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListSmsCommandParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Sim** | **string** | The SID or unique name of the Sim that SMS Command was sent to or from. | 
**Status** | **string** | The status of the SMS Command. Can be: &#x60;queued&#x60;, &#x60;sent&#x60;, &#x60;delivered&#x60;, &#x60;received&#x60; or &#x60;failed&#x60;. See the [SMS Command Status Values](https://www.twilio.com/docs/wireless/api/smscommand-resource#status-values) for a description of each. | 
**Direction** | **string** | The direction of the SMS Command. Can be &#x60;to_sim&#x60; or &#x60;from_sim&#x60;. The value of &#x60;to_sim&#x60; is synonymous with the term &#x60;mobile terminated&#x60;, and &#x60;from_sim&#x60; is synonymous with the term &#x60;mobile originated&#x60;. | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListSmsCommandResponse**](ListSmsCommandResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsageRecord

> ListUsageRecordResponse ListUsageRecord(ctx, optional)



List UsageRecords

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListUsageRecordParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**Sim** | **string** | SID or unique name of a Sim resource. Only show UsageRecords representing usage incurred by this Super SIM. | 
**Fleet** | **string** | SID or unique name of a Fleet resource. Only show UsageRecords representing usage for Super SIMs belonging to this Fleet resource at the time the usage occurred. | 
**Network** | **string** | SID of a Network resource. Only show UsageRecords representing usage on this network. | 
**IsoCountry** | **string** | Alpha-2 ISO Country Code. Only show UsageRecords representing usage in this country. | 
**Group** | **string** | Dimension over which to aggregate usage records. Can be: &#x60;sim&#x60;, &#x60;fleet&#x60;, &#x60;network&#x60;, &#x60;isoCountry&#x60;. Default is to not aggregate across any of these dimensions, UsageRecords will be aggregated into the time buckets described by the &#x60;Granularity&#x60; parameter. | 
**Granularity** | **string** | Time-based grouping that UsageRecords should be aggregated by. Can be: &#x60;hour&#x60;, &#x60;day&#x60;, or &#x60;all&#x60;. Default is &#x60;all&#x60;. &#x60;all&#x60; returns one UsageRecord that describes the usage for the entire period. | 
**StartTime** | **time.Time** | Only include usage that occurred at or after this time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Default is one month before the &#x60;end_time&#x60;. | 
**EndTime** | **time.Time** | Only include usage that occurred before this time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Default is the current time. | 
**PageSize** | **int32** | How many resources to return in each list page. The default is 50, and the maximum is 1000. | 

### Return type

[**ListUsageRecordResponse**](ListUsageRecordResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFleet

> SupersimV1Fleet UpdateFleet(ctx, Sidoptional)



Updates the given properties of a Super SIM Fleet instance from your account.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Fleet resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateFleetParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**CommandsMethod** | **string** | A string representing the HTTP method to use when making a request to &#x60;commands_url&#x60;. Can be one of &#x60;POST&#x60; or &#x60;GET&#x60;. Defaults to &#x60;POST&#x60;. | 
**CommandsUrl** | **string** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored. | 
**NetworkAccessProfile** | **string** | The SID or unique name of the Network Access Profile that will control which cellular networks the Fleet&#39;s SIMs can connect to. | 
**SmsCommandsMethod** | **string** | A string representing the HTTP method to use when making a request to &#x60;sms_commands_url&#x60;. Can be one of &#x60;POST&#x60; or &#x60;GET&#x60;. Defaults to &#x60;POST&#x60;. | 
**SmsCommandsUrl** | **string** | The URL that will receive a webhook when a Super SIM in the Fleet is used to send an SMS from your device to the SMS Commands number. Your server should respond with an HTTP status code in the 200 range; any response body will be ignored. | 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource. | 

### Return type

[**SupersimV1Fleet**](SupersimV1Fleet.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkAccessProfile

> SupersimV1NetworkAccessProfile UpdateNetworkAccessProfile(ctx, Sidoptional)



Updates the given properties of a Network Access Profile in your account.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Network Access Profile to update. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateNetworkAccessProfileParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**UniqueName** | **string** | The new unique name of the Network Access Profile. | 

### Return type

[**SupersimV1NetworkAccessProfile**](SupersimV1NetworkAccessProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSim

> SupersimV1Sim UpdateSim(ctx, Sidoptional)



Updates the given properties of a Super SIM instance from your account.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Sim resource to update. | 

### Other Parameters

Other parameters are passed through a pointer to a UpdateSimParams struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account to which the Sim resource should belong. The Account SID can only be that of the requesting Account or that of a Subaccount of the requesting Account. Only valid when the Sim resource&#39;s status is new. | 
**CallbackMethod** | **string** | The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is POST. | 
**CallbackUrl** | **string** | The URL we should call using the &#x60;callback_method&#x60; after an asynchronous update has finished. | 
**Fleet** | **string** | The SID or unique name of the Fleet to which the SIM resource should be assigned. | 
**Status** | **string** | The new status of the resource. Can be: &#x60;ready&#x60;, &#x60;active&#x60;, or &#x60;inactive&#x60;. See the [Super SIM Status Values](https://www.twilio.com/docs/iot/supersim/api/sim-resource#status-values) for more info. | 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource. | 

### Return type

[**SupersimV1Sim**](SupersimV1Sim.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, 
- **Accept**: application/json, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

