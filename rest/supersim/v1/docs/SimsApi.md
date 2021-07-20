# SimsApi

All URIs are relative to *https://supersim.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSim**](SimsApi.md#CreateSim) | **Post** /v1/Sims | 
[**FetchSim**](SimsApi.md#FetchSim) | **Get** /v1/Sims/{Sid} | 
[**ListSim**](SimsApi.md#ListSim) | **Get** /v1/Sims | 
[**UpdateSim**](SimsApi.md#UpdateSim) | **Post** /v1/Sims/{Sid} | 



## CreateSim

> SupersimV1Sim CreateSim(ctx, optional)



Register a Super SIM to your Account

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateSimParams struct


Name | Type | Description
------------- | ------------- | -------------
**Iccid** | **string** | The [ICCID](https://en.wikipedia.org/wiki/Subscriber_identity_module#ICCID) of the Super SIM to be added to your Account.
**RegistrationCode** | **string** | The 10-digit code required to claim the Super SIM for your Account.

### Return type

[**SupersimV1Sim**](SupersimV1Sim.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchSim

> SupersimV1Sim FetchSim(ctx, Sid)



Fetch a Super SIM instance from your account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Sim resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchSimParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SupersimV1Sim**](SupersimV1Sim.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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


Name | Type | Description
------------- | ------------- | -------------
**Status** | **string** | The status of the Sim resources to read. Can be &#x60;new&#x60;, &#x60;ready&#x60;, &#x60;active&#x60;, &#x60;inactive&#x60;, or &#x60;scheduled&#x60;.
**Fleet** | **string** | The SID or unique name of the Fleet to which a list of Sims are assigned.
**Iccid** | **string** | The [ICCID](https://en.wikipedia.org/wiki/Subscriber_identity_module#ICCID) associated with a Super SIM to filter the list by. Passing this parameter will always return a list containing zero or one SIMs.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**ListSimResponse**](ListSimResponse.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSim

> SupersimV1Sim UpdateSim(ctx, Sidoptional)



Updates the given properties of a Super SIM instance from your account.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the Sim resource to update.

### Other Parameters

Other parameters are passed through a pointer to a UpdateSimParams struct


Name | Type | Description
------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account to which the Sim resource should belong. The Account SID can only be that of the requesting Account or that of a Subaccount of the requesting Account. Only valid when the Sim resource&#39;s status is new.
**CallbackMethod** | **string** | The HTTP method we should use to call &#x60;callback_url&#x60;. Can be: &#x60;GET&#x60; or &#x60;POST&#x60; and the default is POST.
**CallbackUrl** | **string** | The URL we should call using the &#x60;callback_method&#x60; after an asynchronous update has finished.
**Fleet** | **string** | The SID or unique name of the Fleet to which the SIM resource should be assigned.
**Status** | **string** | The new status of the resource. Can be: &#x60;ready&#x60;, &#x60;active&#x60;, or &#x60;inactive&#x60;. See the [Super SIM Status Values](https://www.twilio.com/docs/iot/supersim/api/sim-resource#status-values) for more info.
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. It can be used in place of the resource&#39;s &#x60;sid&#x60; in the URL to address the resource.

### Return type

[**SupersimV1Sim**](SupersimV1Sim.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

