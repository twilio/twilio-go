# MessagingV1CircuitBreakerUpsert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | **string** | The SID of the Account that owns the Messaging Service. |
**MessagingServiceSid** | **string** | The SID of the Messaging Service. |
**SenderSid** | **string** | The SID of the sender the circuit breaker throttles. |
**Mcc** | Pointer to **string** | Mobile country code, or null for every network. |
**Mnc** | Pointer to **string** | Mobile network code, or null for every network. |
**Throttle** | **float64** | Share of traffic, from 0 to 1, the sender is throttled to. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


