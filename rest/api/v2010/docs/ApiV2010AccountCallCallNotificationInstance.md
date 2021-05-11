# ApiV2010AccountCallCallNotificationInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The SID of the Account that created the resource |
**ApiVersion** | Pointer to **string** | The API version used to create the Call Notification resource |
**CallSid** | Pointer to **string** | The SID of the Call the resource is associated with |
**DateCreated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was created |
**DateUpdated** | Pointer to **string** | The RFC 2822 date and time in GMT that the resource was last updated |
**ErrorCode** | Pointer to **string** | A unique error code corresponding to the notification |
**Log** | Pointer to **string** | An integer log level |
**MessageDate** | Pointer to **string** | The date the notification was generated |
**MessageText** | Pointer to **string** | The text of the notification |
**MoreInfo** | Pointer to **string** | A URL for more information about the error code |
**RequestMethod** | Pointer to **string** | HTTP method used with the request url |
**RequestUrl** | Pointer to **string** | URL of the resource that generated the notification |
**RequestVariables** | Pointer to **string** | Twilio-generated HTTP variables sent to the server |
**ResponseBody** | Pointer to **string** | The HTTP body returned by your server |
**ResponseHeaders** | Pointer to **string** | The HTTP headers returned by your server |
**Sid** | Pointer to **string** | The unique string that identifies the resource |
**Uri** | Pointer to **string** | The URI of the resource, relative to `https://api.twilio.com` |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


