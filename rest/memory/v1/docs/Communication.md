# Communication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique communication identifier. |
**ChannelId** | **string** | Channel-specific ID (optional). |[optional] 
**Content** | [**CommunicationContent**](CommunicationContent.md) |  |
**CreatedAt** | [**time.Time**](time.Time.md) | When communication was created. |[optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | When communication was last updated. |[optional] 
**Author** | [**Participant**](Participant.md) |  |
**Recipients** | [**[]CommunicationRecipients**](CommunicationRecipients.md) | Communication recipients |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


