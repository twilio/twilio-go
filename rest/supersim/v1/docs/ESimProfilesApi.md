# ESimProfilesApi

All URIs are relative to *https://supersim.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEsimProfile**](ESimProfilesApi.md#CreateEsimProfile) | **Post** /v1/ESimProfiles | 
[**FetchEsimProfile**](ESimProfilesApi.md#FetchEsimProfile) | **Get** /v1/ESimProfiles/{Sid} | 
[**ListEsimProfile**](ESimProfilesApi.md#ListEsimProfile) | **Get** /v1/ESimProfiles | 



## CreateEsimProfile

> SupersimV1EsimProfile CreateEsimProfile(ctx, optional)



Order an eSIM Profile.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a CreateEsimProfileParams struct


Name | Type | Description
------------- | ------------- | -------------
**CallbackUrl** | **string** | The URL we should call using the `callback_method` when the status of the eSIM Profile changes. At this stage of the eSIM Profile pilot, the a request to the URL will only be called when the ESimProfile resource changes from `reserving` to `available`.
**CallbackMethod** | **string** | The HTTP method we should use to call `callback_url`. Can be: `GET` or `POST` and the default is POST.
**GenerateMatchingId** | **bool** | When set to `true`, a value for `Eid` does not need to be provided. Instead, when the eSIM profile is reserved, a matching ID will be generated and returned via the `matching_id` property. This identifies the specific eSIM profile that can be used by any capable device to claim and download the profile.
**Eid** | **string** | Identifier of the eUICC that will claim the eSIM Profile.

### Return type

[**SupersimV1EsimProfile**](SupersimV1EsimProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchEsimProfile

> SupersimV1EsimProfile FetchEsimProfile(ctx, Sid)



Fetch an eSIM Profile.

### Path Parameters


Name | Type | Description
------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**Sid** | **string** | The SID of the eSIM Profile resource to fetch.

### Other Parameters

Other parameters are passed through a pointer to a FetchEsimProfileParams struct


Name | Type | Description
------------- | ------------- | -------------

### Return type

[**SupersimV1EsimProfile**](SupersimV1EsimProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEsimProfile

> []SupersimV1EsimProfile ListEsimProfile(ctx, optional)



Retrieve a list of eSIM Profiles.

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a ListEsimProfileParams struct


Name | Type | Description
------------- | ------------- | -------------
**Eid** | **string** | List the eSIM Profiles that have been associated with an EId.
**SimSid** | **string** | Find the eSIM Profile resource related to a [Sim](https://www.twilio.com/docs/wireless/api/sim-resource) resource by providing the SIM SID. Will always return an array with either 1 or 0 records.
**Status** | **string** | List the eSIM Profiles that are in a given status.
**PageSize** | **int** | How many resources to return in each list page. The default is 50, and the maximum is 1000.
**Limit** | **int** | Max number of records to return.

### Return type

[**[]SupersimV1EsimProfile**](SupersimV1EsimProfile.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

