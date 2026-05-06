# Action

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of Action to be performed after the Rule is triggered. Supported Actions are: - `WEBHOOK`: A webhook Action sends an HTTP request to a specified URL with Rule execution results.  |
**Method** | **string** | The HTTP method to be used when performing the Action. Must be set to `POST`. |
**Url** | **string** | The URL endpoint where the Action will send the HTTP request containing the Rule execution results. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


