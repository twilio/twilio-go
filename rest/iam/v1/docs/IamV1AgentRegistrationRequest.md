# IamV1AgentRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientUrl** | **string** | The agent's CIMD document URL, which is its client id. Must start with https:// and have a host and a path, with no credentials, fragment or dot segments. |
**Enabled** | **bool** | Whether users in the organization can connect the agent. Defaults to false. |[optional] 
**Type** | **string** | The CIMD category to register the agent for. A URL a category lists may be registered in that category only, and a URL no category lists only as MCP. Defaults to the category that lists the URL, else MCP. A URL is registered in one category per organization. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


