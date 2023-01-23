# TrusthubV1CustomerProfileEvaluation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the Evaluation resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the customer_profile resource. |
**PolicySid** | Pointer to **string** | The unique string of a policy that is associated to the customer_profile resource. |
**CustomerProfileSid** | Pointer to **string** | The unique string that we created to identify the customer_profile resource. |
**Status** | Pointer to [**string**](CustomerProfileEvaluationEnumStatus.md) |  |
**Results** | Pointer to **[]interface{}** | The results of the Evaluation which includes the valid and invalid attributes. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) |  |
**Url** | Pointer to **string** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


