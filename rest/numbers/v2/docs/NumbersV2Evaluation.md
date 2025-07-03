# NumbersV2Evaluation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sid** | Pointer to **string** | The unique string that identifies the Evaluation resource. |
**AccountSid** | Pointer to **string** | The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Bundle resource. |
**RegulationSid** | Pointer to **string** | The unique string of a regulation that is associated to the Bundle resource. |
**BundleSid** | Pointer to **string** | The unique string that we created to identify the Bundle resource. |
**Status** | Pointer to [**string**](EvaluationEnumStatus.md) |  |
**Results** | Pointer to **[]interface{}** | The results of the Evaluation which includes the valid and invalid attributes. |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) |  |
**Url** | Pointer to **string** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


