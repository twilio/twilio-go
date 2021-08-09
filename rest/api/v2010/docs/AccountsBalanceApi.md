# AccountsBalanceApi

All URIs are relative to *https://api.twilio.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchBalance**](AccountsBalanceApi.md#FetchBalance) | **Get** /2010-04-01/Accounts/{AccountSid}/Balance.json | 



## FetchBalance

> ApiV2010Balance FetchBalance(ctx, optional)



Fetch the balance for an Account based on Account Sid. Balance changes may not be reflected immediately. Child accounts do not contain balance information

### Path Parameters

This endpoint does not need any path parameter.

### Other Parameters

Other parameters are passed through a pointer to a FetchBalanceParams struct


Name | Type | Description
------------- | ------------- | -------------
**PathAccountSid** | **string** | The unique SID identifier of the Account.

### Return type

[**ApiV2010Balance**](ApiV2010Balance.md)

### Authorization

[accountSid_authToken](../README.md#accountSid_authToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

