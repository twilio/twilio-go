/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Chat
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi
import (
	"encoding/json"
	"github.com/twilio/twilio-go/client"
	"time"
)
// ChatV1Message struct for ChatV1Message
type ChatV1Message struct {
		// The unique string that we created to identify the Message resource.
	Sid *string `json:"sid,omitempty"`
		// The SID of the [Account](https://www.twilio.com/docs/api/rest/account) that created the Message resource.
	AccountSid *string `json:"account_sid,omitempty"`
		// The JSON string that stores application-specific data. **Note** If this property has been assigned a value, it's only  displayed in a FETCH action that returns a single resource; otherwise, it's null. If the attributes have not been set, `{}` is returned.
	Attributes *string `json:"attributes,omitempty"`
		// The SID of the [Service](https://www.twilio.com/docs/api/chat/rest/services) the resource is associated with.
	ServiceSid *string `json:"service_sid,omitempty"`
		// The SID of the [Channel](https://www.twilio.com/docs/chat/api/channels) that the message was sent to.
	To *string `json:"to,omitempty"`
		// The unique ID of the [Channel](https://www.twilio.com/docs/api/chat/rest/channels) the Message resource belongs to.
	ChannelSid *string `json:"channel_sid,omitempty"`
		// The date and time in GMT when the resource was created specified in [RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
		// The date and time in GMT when the resource was last updated specified in [RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
		// Whether the message has been edited since it was created.
	WasEdited *bool `json:"was_edited,omitempty"`
		// The [identity](https://www.twilio.com/docs/api/chat/guides/identity) of the message's author. The default value is `system`.
	From *string `json:"from,omitempty"`
		// The content of the message.
	Body *string `json:"body,omitempty"`
		// The index of the message within the [Channel](https://www.twilio.com/docs/chat/api/channels).
	Index *int `json:"index,omitempty"`
		// The absolute URL of the Message resource.
	Url *string `json:"url,omitempty"`
}


