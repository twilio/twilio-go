# ConversationsV2Capability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique Capability identifier. |
**Name** | **string** | Human-readable Capability name. |
**DomainId** | **string** | Identifier of the domain that owns the Capability. |
**Description** | **string** | Human-readable description of the Capability. |[optional] 
**Status** | **string** | Lifecycle status of the Capability. |
**Visibility** | **string** | Whether the Capability is available publicly or only to internal consumers. |
**PolicyRef** | **string** | Identifier for the Cedar policy that controls this Capability. |
**SchemaRef** | **string** | Identifier for the Cedar schema that defines this Capability. |
**Version** | **int** | Optimistic-locking version. |
**CreatedAt** | [**time.Time**](time.Time.md) | Timestamp when the Capability was created. |
**UpdatedAt** | [**time.Time**](time.Time.md) | Timestamp when the Capability was last updated. |
**SubscriptionPrincipalSchema** | **map[string]interface{}** | JSON Schema describing the principal accepted by subscriptions to this Capability. |[optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


