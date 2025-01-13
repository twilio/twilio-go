# MarketplaceV1ModuleDataManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | URL to query the subresource. |
**Sid** | Pointer to **string** | ModuleSid that identifies this Listing. |
**Description** | Pointer to **map[string]interface{}** | A JSON object describing the module and is displayed under the Description tab of the Module detail page. You can define the main body of the description, highlight key features or aspects of the module and if applicable, provide code samples for developers |
**Support** | Pointer to **map[string]interface{}** | A JSON object containing information on how customers can obtain support for the module. Use this parameter to provide details such as contact information and support description. |
**Policies** | Pointer to **map[string]interface{}** | A JSON object describing the module's privacy and legal policies and is displayed under the Policies tab of the Module detail page. The maximum file size for Policies is 5MB |
**ModuleInfo** | Pointer to **map[string]interface{}** | A JSON object containing essential attributes that define a module. This information is presented on the Module detail page in the Twilio Marketplace Catalog. You can pass the following attributes in the JSON object |
**Documentation** | Pointer to **map[string]interface{}** | A JSON object for providing comprehensive information, instructions, and resources related to the module |
**Configuration** | Pointer to **map[string]interface{}** | A JSON object for providing listing specific configuration. Contains button setup, notification url, among others. |
**Pricing** | Pointer to **map[string]interface{}** | A JSON object for providing Listing specific pricing information. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


