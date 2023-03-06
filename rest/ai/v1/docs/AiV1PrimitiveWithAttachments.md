# AiV1PrimitiveWithAttachments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The unique SID identifier of the Account. |
**Sid** | Pointer to **string** | The unique SID identifier of the Primitive. |
**UniqueName** | Pointer to **string** | Provides a unique and addressable name to be assigned to this Primitive, assigned by the developer, to be optionally used in addition to SID. |
**ModelUrn** | Pointer to **string** | Defines the underlying model for the primitive. |
**FriendlyName** | Pointer to **string** | A human readable description of this resource, up to 64 characters. |
**Type** | Pointer to [**string**](PrimitiveWithAttachmentsEnumType.md) |  |
**PrimitiveStatus** | Pointer to [**string**](PrimitiveWithAttachmentsEnumStatus.md) |  |
**DateCreated** | Pointer to [**time.Time**](time.Time.md) | The date that this Primitive was created, given in ISO 8601 format. |
**DateUpdated** | Pointer to [**time.Time**](time.Time.md) | The date that this Primitive was updated, given in ISO 8601 format. |
**LastBuilt** | Pointer to [**time.Time**](time.Time.md) | The date that this Primitive was last built, given in ISO 8601 format. |
**Definition** | Pointer to **interface{}** | Defines the configuration and examples of either the Literal Phrase Set or Semantic Phrase Set. |
**Attachments** | Pointer to **interface{}** | List of Services attached to the Primitive, either active or inactive. |
**Config** | Pointer to **interface{}** | Configuration of the primitive. |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


