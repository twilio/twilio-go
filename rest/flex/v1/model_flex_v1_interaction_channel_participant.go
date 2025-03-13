/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
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
)

// FlexV1InteractionChannelParticipant struct for FlexV1InteractionChannelParticipant
type FlexV1InteractionChannelParticipant struct {
	// The unique string created by Twilio to identify an Interaction Channel Participant resource.
	Sid  *string `json:"sid,omitempty"`
	Type *string `json:"type,omitempty"`
	// The Interaction Sid for this channel.
	InteractionSid *string `json:"interaction_sid,omitempty"`
	// The Channel Sid for this Participant.
	ChannelSid *string `json:"channel_sid,omitempty"`
	Url        *string `json:"url,omitempty"`
	// The Participant's routing properties.
	RoutingProperties *map[string]interface{} `json:"routing_properties,omitempty"`
}
