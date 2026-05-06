# Store

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | Provides a unique and addressable name to be assigned to this Store. This name is assigned by the developer and can be used in addition to the ID. It is intended to be human-readable and unique within the account. |
**Description** | **string** | A human readable description of this resource, up to 128 characters. |[optional] 
**Id** | **string** | The unique identifier for the Memory Store |
**Status** | **string** | The current status of the Memory Store.  A store begins in the QUEUED state as it is scheduled for processing.  It then moves to PROVISIONING at the beginning of processing. It transitions to ACTIVE once all dependent resources are provisioned, including Conversational Intelligence capabilities.  If there is an issue provisioning resources, the store will move to the FAILED state. |
**IntelligenceServiceId** | Pointer to **string** | The ID of the associated intelligence service that was provisioned for memory extraction. |
**Version** | **int** | The current version number of the Memory Store. Incremented on each successful update. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


