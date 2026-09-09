# IamV1PostLoginActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | The version of the Whitney webhook payload contract |[optional] 
**Event** | **string** | The Whitney event name (e.g. auth_flow.checkpoint) |[optional] 
**AccountSid** | **string** | The Twilio account SID associated with the login |[optional] 
**NodeId** | **string** | The Whitney auth flow node identifier for this action |[optional] 
**ActionId** | **string** | Whitney's identifier for this action invocation |[optional] 
**Timestamp** | **int64** | Unix timestamp (seconds) of the login event |[optional] 
**Oauth** | Pointer to [**IamV1PostLoginActionOAuthContext**](IamV1PostLoginActionOAuthContext.md) |  |
**Acr** | **string** | The authentication context class reference the login satisfied |[optional] 
**CompletedFactors** | **[]string** | The authentication factors the user completed, as RFC 8176 authentication method reference values (e.g. pwd, otp, mfa) |[optional] 
**UserContext** | Pointer to [**IamV1PostLoginActionUserContext**](IamV1PostLoginActionUserContext.md) |  |
**RequestContext** | Pointer to [**IamV1PostLoginActionRequestContext**](IamV1PostLoginActionRequestContext.md) |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


