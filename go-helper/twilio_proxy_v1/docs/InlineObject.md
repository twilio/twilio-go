# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | **string** | The URL we should call when the interaction status changes. | [optional] 
**ChatInstanceSid** | **string** | The SID of the Chat Service Instance managed by Proxy Service. The Chat Service enables Proxy to forward SMS and channel messages to this chat instance. This is a one-to-one relationship. | [optional] 
**DefaultTtl** | **int32** | The default &#x60;ttl&#x60; value to set for Sessions created in the Service. The TTL (time to live) is measured in seconds after the Session&#39;s last create or last Interaction. The default value of &#x60;0&#x60; indicates an unlimited Session length. You can override a Session&#39;s default TTL value by setting its &#x60;ttl&#x60; value. | [optional] 
**GeoMatchLevel** | **string** | Where a proxy number must be located relative to the participant identifier. Can be: &#x60;country&#x60;, &#x60;area-code&#x60;, or &#x60;extended-area-code&#x60;. The default value is &#x60;country&#x60; and more specific areas than &#x60;country&#x60; are only available in North America. | [optional] 
**InterceptCallbackUrl** | **string** | The URL we call on each interaction. If we receive a 403 status, we block the interaction; otherwise the interaction continues. | [optional] 
**NumberSelectionBehavior** | **string** | The preference for Proxy Number selection in the Service instance. Can be: &#x60;prefer-sticky&#x60; or &#x60;avoid-sticky&#x60; and the default is &#x60;prefer-sticky&#x60;. &#x60;prefer-sticky&#x60; means that we will try and select the same Proxy Number for a given participant if they have previous [Sessions](https://www.twilio.com/docs/proxy/api/session), but we will not fail if that Proxy Number cannot be used.  &#x60;avoid-sticky&#x60; means that we will try to use different Proxy Numbers as long as that is possible within a given pool rather than try and use a previously assigned number. | [optional] 
**OutOfSessionCallbackUrl** | **string** | The URL we should call when an inbound call or SMS action occurs on a closed or non-existent Session. If your server (or a Twilio [function](https://www.twilio.com/functions)) responds with valid [TwiML](https://www.twilio.com/docs/voice/twiml), we will process it. This means it is possible, for example, to play a message for a call, send an automated text message response, or redirect a call to another Phone Number. See [Out-of-Session Callback Response Guide](https://www.twilio.com/docs/proxy/out-session-callback-response-guide) for more information. | [optional] 
**UniqueName** | **string** | An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.** | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


