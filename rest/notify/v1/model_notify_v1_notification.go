/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Notify
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

// NotifyV1Notification struct for NotifyV1Notification
type NotifyV1Notification struct {
	// The unique string that we created to identify the Notification resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Notification resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the [Service](https://www.twilio.com/docs/notify/api/service-resource) the resource is associated with.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The list of `identity` values of the Users to notify. We will attempt to deliver notifications only to Bindings with an identity in this list.
	Identities *[]string `json:"identities,omitempty"`
	// The tags that select the Bindings to notify. Notifications will be attempted only to Bindings that have all of the tags listed in this property.
	Tags *[]string `json:"tags,omitempty"`
	// The list of Segments to notify. The [Segment](https://www.twilio.com/docs/notify/api/segment-resource) resource is deprecated. Use the `tags` property, instead.
	Segments *[]string `json:"segments,omitempty"`
	Priority *string   `json:"priority,omitempty"`
	// How long, in seconds, the notification is valid. Can be an integer between 0 and 2,419,200, which is 4 weeks, the default and the maximum supported time to live (TTL). Delivery should be attempted if the device is offline until the TTL elapses. Zero means that the notification delivery is attempted immediately, only once, and is not stored for future delivery. SMS does not support this property.
	Ttl int `json:"ttl,omitempty"`
	// The notification title. For FCM and GCM, this translates to the `data.twi_title` value. For APNS, this translates to the `aps.alert.title` value. SMS does not support this property. This field is not visible on iOS phones and tablets but appears on Apple Watch and Android devices.
	Title *string `json:"title,omitempty"`
	// The notification text. For FCM and GCM, translates to `data.twi_body`. For APNS, translates to `aps.alert.body`. For SMS, translates to `body`. SMS requires either this `body` value, or `media_urls` attribute defined in the `sms` parameter of the notification.
	Body *string `json:"body,omitempty"`
	// The name of the sound to be played for the notification. For FCM and GCM, this Translates to `data.twi_sound`.  For APNS, this translates to `aps.sound`.  SMS does not support this property.
	Sound *string `json:"sound,omitempty"`
	// The actions to display for the notification. For APNS, translates to the `aps.category` value. For GCM, translates to the `data.twi_action` value. For SMS, this parameter is not supported and is omitted from deliveries to those channels.
	Action *string `json:"action,omitempty"`
	// The custom key-value pairs of the notification's payload. For FCM and GCM, this value translates to `data` in the FCM and GCM payloads. FCM and GCM [reserve certain keys](https://firebase.google.com/docs/cloud-messaging/http-server-ref) that cannot be used in those channels. For APNS, attributes of `data` are inserted into the APNS payload as custom properties outside of the `aps` dictionary. In all channels, we reserve keys that start with `twi_` for future use. Custom keys that start with `twi_` are not allowed and are rejected as 400 Bad request with no delivery attempted. For SMS, this parameter is not supported and is omitted from deliveries to those channels.
	Data *map[string]interface{} `json:"data,omitempty"`
	// The APNS-specific payload that overrides corresponding attributes in the generic payload for APNS Bindings. This property maps to the APNS `Payload` item, therefore the `aps` key must be used to change standard attributes. Adds custom key-value pairs to the root of the dictionary. See the [APNS documentation](https://developer.apple.com/library/content/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/CommunicatingwithAPNs.html) for more details. We reserve keys that start with `twi_` for future use. Custom keys that start with `twi_` are not allowed.
	Apn *map[string]interface{} `json:"apn,omitempty"`
	// The GCM-specific payload that overrides corresponding attributes in the generic payload for GCM Bindings.  This property maps to the root JSON dictionary. Target parameters `to`, `registration_ids`, and `notification_key` are not allowed. We reserve keys that start with `twi_` for future use. Custom keys that start with `twi_` are not allowed.
	Gcm *map[string]interface{} `json:"gcm,omitempty"`
	// The FCM-specific payload that overrides corresponding attributes in the generic payload for FCM Bindings. This property maps to the root JSON dictionary. See the [FCM documentation](https://firebase.google.com/docs/cloud-messaging/http-server-ref#downstream) for more details. Target parameters `to`, `registration_ids`, `condition`, and `notification_key` are not allowed in this parameter. We reserve keys that start with `twi_` for future use. Custom keys that start with `twi_` are not allowed. FCM also [reserves certain keys](https://firebase.google.com/docs/cloud-messaging/http-server-ref), which cannot be used in that channel.
	Fcm *map[string]interface{} `json:"fcm,omitempty"`
	// The SMS-specific payload that overrides corresponding attributes in the generic payload for SMS Bindings.  Each attribute in this value maps to the corresponding `form` parameter of the Twilio [Message](https://www.twilio.com/docs/sms/api/message-resource) resource.  These parameters of the Message resource are supported in snake case format: `body`, `media_urls`, `status_callback`, and `max_price`.  The `status_callback` parameter overrides the corresponding parameter in the messaging service, if configured. The `media_urls` property expects a JSON array.
	Sms *map[string]interface{} `json:"sms,omitempty"`
	// Deprecated.
	FacebookMessenger *map[string]interface{} `json:"facebook_messenger,omitempty"`
	// Deprecated.
	Alexa *map[string]interface{} `json:"alexa,omitempty"`
}
