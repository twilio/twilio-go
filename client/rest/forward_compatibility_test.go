package rest

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	messagingApi "github.com/twilio/twilio-go/rest/messaging/v2"
	memoryApi "github.com/twilio/twilio-go/rest/memory/v1"
	v2010Api "github.com/twilio/twilio-go/rest/api/v2010"
)

// ─────────────────────────────────────────────────────────────────────
// MessageResource (flat)
// ─────────────────────────────────────────────────────────────────────

func baseMessageJSON() string {
	return `{
		"body": "Hello, World!",
		"num_segments": "1",
		"direction": "outbound-api",
		"from": "+15558881111",
		"to": "+15559992222",
		"date_updated": "Thu, 30 Jul 2015 20:00:00 +0000",
		"price": "-0.00750",
		"error_message": null,
		"uri": "/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
		"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"num_media": "0",
		"status": "sent",
		"messaging_service_sid": "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"sid": "SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"date_sent": "Thu, 30 Jul 2015 20:00:00 +0000",
		"date_created": "Thu, 30 Jul 2015 20:00:00 +0000",
		"error_code": null,
		"price_unit": "USD",
		"api_version": "2010-04-01",
		"subresource_uris": {
			"media": "/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media.json"
		}
	}`
}

func TestMessage_AddFieldAtParentLevel(t *testing.T) {
	jsonStr := `{
		"body": "Hello, World!",
		"num_segments": "1",
		"direction": "outbound-api",
		"from": "+15558881111",
		"to": "+15559992222",
		"date_updated": "Thu, 30 Jul 2015 20:00:00 +0000",
		"price": "-0.00750",
		"error_message": null,
		"uri": "/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Messages/SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
		"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"num_media": "0",
		"status": "sent",
		"messaging_service_sid": "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"sid": "SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"date_sent": "Thu, 30 Jul 2015 20:00:00 +0000",
		"date_created": "Thu, 30 Jul 2015 20:00:00 +0000",
		"error_code": null,
		"price_unit": "USD",
		"api_version": "2010-04-01",
		"subresource_uris": {"media": "/media.json"},
		"new_unknown_field": "some-value",
		"another_future_field": 12345
	}`

	var msg v2010Api.ApiV2010Message
	err := json.Unmarshal([]byte(jsonStr), &msg)

	assert.NoError(t, err)
	assert.Equal(t, "Hello, World!", *msg.Body)
	assert.Equal(t, "SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", *msg.Sid)
	assert.Equal(t, "sent", *msg.Status)
	assert.Equal(t, "outbound-api", *msg.Direction)
	assert.Equal(t, "+15558881111", *msg.From)
	assert.Equal(t, "+15559992222", *msg.To)
	assert.Equal(t, "1", *msg.NumSegments)
	assert.Equal(t, "-0.00750", *msg.Price)
	assert.Equal(t, "USD", *msg.PriceUnit)
}

func TestMessage_RemoveFieldAtParentLevel(t *testing.T) {
	jsonStr := `{
		"body": "Hello, World!",
		"num_segments": "1",
		"direction": "outbound-api",
		"from": "+15558881111",
		"to": "+15559992222",
		"date_updated": "Thu, 30 Jul 2015 20:00:00 +0000",
		"uri": "/message.json",
		"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"status": "sent",
		"messaging_service_sid": "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"sid": "SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"date_sent": "Thu, 30 Jul 2015 20:00:00 +0000",
		"date_created": "Thu, 30 Jul 2015 20:00:00 +0000",
		"price_unit": "USD",
		"api_version": "2010-04-01",
		"subresource_uris": {}
	}`

	var msg v2010Api.ApiV2010Message
	err := json.Unmarshal([]byte(jsonStr), &msg)

	assert.NoError(t, err)
	assert.Equal(t, "Hello, World!", *msg.Body)
	assert.Equal(t, "SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", *msg.Sid)
	assert.Nil(t, msg.ErrorMessage)
	assert.Nil(t, msg.Price)
	assert.Nil(t, msg.NumMedia)
	assert.Nil(t, msg.ErrorCode)
	assert.Equal(t, "sent", *msg.Status)
}

func TestMessage_UnknownEnumValueAtParentLevel(t *testing.T) {
	jsonStr := `{
		"body": "Hello, World!",
		"num_segments": "1",
		"direction": "outbound-future",
		"from": "+15558881111",
		"to": "+15559992222",
		"date_updated": "Thu, 30 Jul 2015 20:00:00 +0000",
		"price": "-0.00750",
		"error_message": null,
		"uri": "/message.json",
		"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"num_media": "0",
		"status": "future_unknown_status",
		"messaging_service_sid": "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"sid": "SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"date_sent": "Thu, 30 Jul 2015 20:00:00 +0000",
		"date_created": "Thu, 30 Jul 2015 20:00:00 +0000",
		"error_code": null,
		"price_unit": "USD",
		"api_version": "2010-04-01",
		"subresource_uris": {}
	}`

	var msg v2010Api.ApiV2010Message
	err := json.Unmarshal([]byte(jsonStr), &msg)

	assert.NoError(t, err)
	assert.Equal(t, "future_unknown_status", *msg.Status)
	assert.Equal(t, "outbound-future", *msg.Direction)
	assert.Equal(t, "Hello, World!", *msg.Body)
	assert.Equal(t, "SMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", *msg.Sid)
}

func TestMessage_EmptyResponseBody(t *testing.T) {
	jsonStr := `{}`

	var msg v2010Api.ApiV2010Message
	err := json.Unmarshal([]byte(jsonStr), &msg)

	assert.NoError(t, err)
	assert.Nil(t, msg.Body)
	assert.Nil(t, msg.Sid)
	assert.Nil(t, msg.Status)
	assert.Nil(t, msg.Direction)
	assert.Nil(t, msg.From)
	assert.Nil(t, msg.To)
}

// ─────────────────────────────────────────────────────────────────────
// ChannelsSenderResource (nested, top-level)
// ─────────────────────────────────────────────────────────────────────

func baseChannelsSenderJSON() string {
	return `{
		"sid": "XEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"status": "ONLINE",
		"sender_id": "whatsapp:+15558881111",
		"friendly_name": "My WhatsApp Sender",
		"configuration": {
			"waba_id": "WABA123",
			"verification_method": "sms",
			"verification_code": null,
			"voice_application_sid": null,
			"account_type": null
		},
		"webhook": {
			"callback_url": "https://example.com/callback",
			"callback_method": "POST",
			"fallback_url": "https://example.com/fallback",
			"fallback_method": "POST",
			"status_callback_url": "https://example.com/status",
			"status_callback_method": "POST"
		},
		"profile": {
			"name": "My Business",
			"about": "We sell things",
			"address": "123 Main St",
			"description": "A business",
			"logo_url": "https://example.com/logo.png",
			"banner_url": null,
			"privacy_url": "https://example.com/privacy",
			"terms_of_service_url": null,
			"accent_color": "#FF0000",
			"use_case": "TRANSACTIONAL",
			"vertical": "Shopping and Retail",
			"websites": [{"website": "https://example.com", "label": "Main"}],
			"emails": [{"email": "info@example.com", "label": "Info"}],
			"phone_numbers": [{"phone_number": "+15558881111", "label": "Support"}]
		},
		"properties": {
			"quality_rating": "GREEN",
			"messaging_limit": "1000"
		},
		"offline_reasons": [
			{"code": "12345", "message": "Some offline reason", "more_info": "https://example.com/error"}
		],
		"compliance": {
			"registration_sid": "CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			"countries": [
				{
					"country": "US",
					"registration_sid": "CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
					"status": "APPROVED",
					"carriers": [{"name": "Verizon", "status": "APPROVED"}]
				}
			]
		},
		"url": "https://messaging.twilio.com/v2/Channels/Senders/XEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	}`
}

func TestChannelsSender_AddFieldAtNestedLevel(t *testing.T) {
	jsonStr := `{
		"sid": "XEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"status": "ONLINE",
		"sender_id": "whatsapp:+15558881111",
		"friendly_name": "My WhatsApp Sender",
		"configuration": {
			"waba_id": "WABA123",
			"verification_method": "sms",
			"future_config_field": "new-value"
		},
		"webhook": {
			"callback_url": "https://example.com/callback",
			"callback_method": "POST",
			"future_webhook_field": true
		},
		"profile": {
			"name": "My Business",
			"about": "We sell things",
			"future_profile_field": "new-profile-value",
			"websites": [{"website": "https://example.com", "label": "Main"}],
			"emails": [],
			"phone_numbers": []
		},
		"properties": {
			"quality_rating": "GREEN",
			"messaging_limit": "1000",
			"future_properties_field": 999
		},
		"offline_reasons": [
			{"code": "12345", "message": "Some offline reason", "more_info": "https://example.com/error", "future_offline_field": "extra"}
		],
		"compliance": {
			"registration_sid": "CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			"future_compliance_field": "extra",
			"countries": [
				{
					"country": "US",
					"registration_sid": "CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
					"status": "APPROVED",
					"future_country_field": "extra",
					"carriers": [{"name": "Verizon", "status": "APPROVED"}]
				}
			]
		},
		"url": "https://example.com"
	}`

	var sender messagingApi.MessagingV2ChannelsSenderResponse
	err := json.Unmarshal([]byte(jsonStr), &sender)

	assert.NoError(t, err)
	assert.Equal(t, "XEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", *sender.Sid)
	assert.Equal(t, "ONLINE", *sender.Status)
	assert.NotNil(t, sender.Configuration)
	assert.Equal(t, "WABA123", *sender.Configuration.WabaId)
	assert.NotNil(t, sender.Webhook)
	assert.Equal(t, "https://example.com/callback", *sender.Webhook.CallbackUrl)
	assert.NotNil(t, sender.Profile)
	assert.Equal(t, "My Business", *sender.Profile.Name)
	assert.NotNil(t, sender.Properties)
	assert.Equal(t, "GREEN", *sender.Properties.QualityRating)
	assert.Len(t, *sender.OfflineReasons, 1)
	assert.Equal(t, "12345", *(*sender.OfflineReasons)[0].Code)
}

func TestChannelsSender_RemoveFieldInsideNestedObject(t *testing.T) {
	jsonStr := `{
		"sid": "XEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"status": "ONLINE",
		"sender_id": "whatsapp:+15558881111",
		"friendly_name": "My WhatsApp Sender",
		"configuration": {
			"verification_code": null,
			"voice_application_sid": null,
			"account_type": null
		},
		"webhook": {
			"callback_url": "https://example.com/callback",
			"callback_method": "POST",
			"status_callback_url": "https://example.com/status",
			"status_callback_method": "POST"
		},
		"profile": {
			"name": "My Business",
			"about": "We sell things",
			"address": "123 Main St",
			"description": "A business",
			"websites": [],
			"emails": [],
			"phone_numbers": []
		},
		"properties": {
			"quality_rating": "GREEN",
			"messaging_limit": "1000"
		},
		"offline_reasons": [],
		"compliance": {
			"registration_sid": "CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			"countries": []
		},
		"url": "https://example.com"
	}`

	var sender messagingApi.MessagingV2ChannelsSenderResponse
	err := json.Unmarshal([]byte(jsonStr), &sender)

	assert.NoError(t, err)
	assert.Equal(t, "XEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", *sender.Sid)
	assert.NotNil(t, sender.Configuration)
	assert.Nil(t, sender.Configuration.WabaId)
	assert.Nil(t, sender.Configuration.VerificationMethod)
	assert.NotNil(t, sender.Profile)
	assert.Nil(t, sender.Profile.LogoUrl)
	assert.Nil(t, sender.Profile.AccentColor)
	assert.NotNil(t, sender.Webhook)
	assert.Nil(t, sender.Webhook.FallbackUrl)
	assert.Nil(t, sender.Webhook.FallbackMethod)
	assert.Equal(t, "https://example.com/callback", *sender.Webhook.CallbackUrl)
}

func TestChannelsSender_RemoveWholeNestedObject(t *testing.T) {
	jsonStr := `{
		"sid": "XEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"status": "ONLINE",
		"sender_id": "whatsapp:+15558881111",
		"friendly_name": "My WhatsApp Sender",
		"configuration": null,
		"webhook": null,
		"profile": null,
		"properties": null,
		"offline_reasons": null,
		"compliance": null,
		"url": "https://example.com"
	}`

	var sender messagingApi.MessagingV2ChannelsSenderResponse
	err := json.Unmarshal([]byte(jsonStr), &sender)

	assert.NoError(t, err)
	assert.Equal(t, "XEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", *sender.Sid)
	assert.Equal(t, "ONLINE", *sender.Status)
	assert.Nil(t, sender.Configuration)
	assert.Nil(t, sender.Webhook)
	assert.Nil(t, sender.Profile)
	assert.Nil(t, sender.Properties)
	assert.Nil(t, sender.OfflineReasons)
	assert.Nil(t, sender.Compliance)
}

func TestChannelsSender_NewItemAppendedToNestedList(t *testing.T) {
	jsonStr := `{
		"sid": "XEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"status": "ONLINE",
		"sender_id": "whatsapp:+15558881111",
		"friendly_name": "My WhatsApp Sender",
		"configuration": {"waba_id": "WABA123"},
		"webhook": {"callback_url": "https://example.com/callback"},
		"profile": {
			"name": "My Business",
			"websites": [
				{"website": "https://example.com", "label": "Main"},
				{"website": "https://other.example.com", "label": "Other", "future_field": "unexpected"}
			],
			"emails": [{"email": "info@example.com", "label": "Info"}],
			"phone_numbers": [{"phone_number": "+15558881111", "label": "Support"}]
		},
		"properties": {"quality_rating": "GREEN"},
		"offline_reasons": [
			{"code": "12345", "message": "Some offline reason", "more_info": "https://example.com/error"},
			{"code": "67890", "message": "Another reason", "more_info": "https://example.com/error2", "brand_new_field": "surprise"}
		],
		"compliance": {
			"registration_sid": "CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			"countries": [
				{"country": "US", "registration_sid": "CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", "status": "APPROVED", "carriers": [{"name": "Verizon", "status": "APPROVED"}]},
				{"country": "GB", "registration_sid": "CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", "status": "APPROVED", "carriers": [{"name": "EE", "status": "APPROVED"}, {"name": "Three", "status": "CARRIER_REVIEW"}]}
			]
		},
		"url": "https://example.com"
	}`

	var sender messagingApi.MessagingV2ChannelsSenderResponse
	err := json.Unmarshal([]byte(jsonStr), &sender)

	assert.NoError(t, err)
	assert.Len(t, *sender.OfflineReasons, 2)
	assert.Equal(t, "12345", *(*sender.OfflineReasons)[0].Code)
	assert.Equal(t, "67890", *(*sender.OfflineReasons)[1].Code)
	assert.Equal(t, "Another reason", *(*sender.OfflineReasons)[1].Message)

	assert.Len(t, *sender.Profile.Websites, 2)
	assert.Equal(t, "https://example.com", (*sender.Profile.Websites)[0].Website)
	assert.Equal(t, "https://other.example.com", (*sender.Profile.Websites)[1].Website)

	assert.Len(t, sender.Compliance.Countries, 2)
	assert.Equal(t, "GB", sender.Compliance.Countries[1].Country)
	assert.Len(t, sender.Compliance.Countries[1].Carriers, 2)
}

func TestChannelsSender_UnknownEnumValueAtNestedLevelAndInsideListItem(t *testing.T) {
	jsonStr := `{
		"sid": "XEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"status": "FUTURE_UNKNOWN_STATUS",
		"sender_id": "whatsapp:+15558881111",
		"friendly_name": "My WhatsApp Sender",
		"configuration": {"waba_id": "WABA123"},
		"webhook": {"callback_url": "https://example.com/callback"},
		"profile": {"name": "My Business", "websites": [], "emails": [], "phone_numbers": []},
		"properties": {"quality_rating": "GREEN"},
		"offline_reasons": [],
		"compliance": {
			"registration_sid": "CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			"countries": [
				{
					"country": "US",
					"registration_sid": "CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
					"status": "FUTURE_COUNTRY_STATUS",
					"carriers": [{"name": "Verizon", "status": "FUTURE_CARRIER_STATUS"}]
				}
			]
		},
		"url": "https://example.com"
	}`

	var sender messagingApi.MessagingV2ChannelsSenderResponse
	err := json.Unmarshal([]byte(jsonStr), &sender)

	assert.NoError(t, err)
	assert.Equal(t, "FUTURE_UNKNOWN_STATUS", *sender.Status)
	assert.Equal(t, "XEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", *sender.Sid)
	assert.Equal(t, "whatsapp:+15558881111", *sender.SenderId)
	assert.Equal(t, "WABA123", *sender.Configuration.WabaId)
	assert.Equal(t, messagingApi.MessagingV2RcsCountryStatus("FUTURE_COUNTRY_STATUS"), sender.Compliance.Countries[0].Status)
	assert.Equal(t, messagingApi.MessagingV2RcsCarrierStatus("FUTURE_CARRIER_STATUS"), sender.Compliance.Countries[0].Carriers[0].Status)
}

// ─────────────────────────────────────────────────────────────────────
// RecallCreateResource (nested, operation-specific)
// ─────────────────────────────────────────────────────────────────────

func TestRecall_AddFieldAtParentAndNestedLevel(t *testing.T) {
	jsonStr := `{
		"observations": [
			{
				"content": "User prefers email communication",
				"occurredAt": "2025-01-15T10:30:00Z",
				"source": "conversation-analysis",
				"conversationIds": ["conv_123"],
				"id": "obs_001",
				"createdAt": "2025-01-15T10:30:00Z",
				"updatedAt": "2025-01-15T10:30:00Z",
				"score": 0.95,
				"future_obs_field": "extra-obs"
			}
		],
		"summaries": [
			{
				"source": "auto-summarizer",
				"content": "Customer discussed billing concerns",
				"occurredAt": "2025-01-14T09:00:00Z",
				"conversationId": "conv_456",
				"id": "sum_001",
				"createdAt": "2025-01-14T09:00:00Z",
				"updatedAt": "2025-01-14T09:00:00Z",
				"score": 0.88,
				"future_sum_field": "extra-sum"
			}
		],
		"communications": [
			{
				"id": "comm_001",
				"channelId": "CH123",
				"content": {"text": "Hello, how can I help?", "future_content_field": "extra-content"},
				"createdAt": "2025-01-15T10:00:00Z",
				"updatedAt": "2025-01-15T10:00:00Z",
				"author": {
					"id": "agent_001",
					"name": "Agent Smith",
					"type": "HUMAN_AGENT",
					"profileId": "prof_001",
					"address": "+15558881111",
					"channel": "sms",
					"future_author_field": "extra-author"
				},
				"recipients": [
					{
						"id": "cust_001",
						"name": "John Doe",
						"type": "CUSTOMER",
						"profileId": "prof_002",
						"address": "+15559992222",
						"channel": "sms",
						"deliveryStatus": "delivered",
						"future_recipient_field": "extra-recipient"
					}
				],
				"future_comm_field": "extra-comm"
			}
		],
		"meta": {
			"queryTime": 42,
			"future_meta_field": 999
		},
		"future_top_level_field": "extra-top"
	}`

	var recall memoryApi.MemoryRetrievalResponse
	err := json.Unmarshal([]byte(jsonStr), &recall)

	assert.NoError(t, err)
	assert.Len(t, recall.Observations, 1)
	assert.Equal(t, "User prefers email communication", recall.Observations[0].Content)
	assert.Equal(t, 0.95, recall.Observations[0].Score)

	assert.Len(t, recall.Summaries, 1)
	assert.Equal(t, "Customer discussed billing concerns", recall.Summaries[0].Content)

	assert.Len(t, recall.Communications, 1)
	assert.Equal(t, "comm_001", recall.Communications[0].Id)
	assert.Equal(t, "Agent Smith", recall.Communications[0].Author.Name)
	assert.Equal(t, "John Doe", recall.Communications[0].Recipients[0].Name)

	assert.Equal(t, 42, recall.Meta.QueryTime)
}

func TestRecall_RemoveFieldAtParentAndInsideNestedObject(t *testing.T) {
	jsonStr := `{
		"observations": [
			{
				"content": "User prefers email communication",
				"occurredAt": "2025-01-15T10:30:00Z",
				"source": "conversation-analysis",
				"id": "obs_001",
				"createdAt": "2025-01-15T10:30:00Z",
				"updatedAt": "2025-01-15T10:30:00Z"
			}
		],
		"summaries": [
			{
				"content": "Customer discussed billing concerns",
				"occurredAt": "2025-01-14T09:00:00Z",
				"conversationId": "conv_456",
				"id": "sum_001",
				"createdAt": "2025-01-14T09:00:00Z",
				"updatedAt": "2025-01-14T09:00:00Z"
			}
		],
		"communications": [
			{
				"id": "comm_001",
				"content": {},
				"author": {
					"id": "agent_001",
					"name": "Agent Smith",
					"address": "+15558881111",
					"channel": "sms"
				},
				"recipients": [
					{
						"id": "cust_001",
						"name": "John Doe",
						"address": "+15559992222",
						"channel": "sms"
					}
				]
			}
		],
		"meta": {
			"queryTime": 42
		}
	}`

	var recall memoryApi.MemoryRetrievalResponse
	err := json.Unmarshal([]byte(jsonStr), &recall)

	assert.NoError(t, err)
	assert.Len(t, recall.Observations, 1)
	assert.Equal(t, "User prefers email communication", recall.Observations[0].Content)
	assert.Equal(t, float64(0), recall.Observations[0].Score)
	assert.Nil(t, recall.Observations[0].ConversationIds)

	assert.Len(t, recall.Summaries, 1)
	assert.Equal(t, "Customer discussed billing concerns", recall.Summaries[0].Content)
	assert.Empty(t, recall.Summaries[0].Source)
	assert.Equal(t, float64(0), recall.Summaries[0].Score)

	assert.Len(t, recall.Communications, 1)
	assert.Empty(t, recall.Communications[0].ChannelId)
	assert.Empty(t, recall.Communications[0].Author.ProfileId)
	assert.Empty(t, recall.Communications[0].Recipients[0].DeliveryStatus)
	assert.Empty(t, recall.Communications[0].Recipients[0].ProfileId)
	assert.Empty(t, recall.Communications[0].Content.Text)
}

func TestRecall_UnknownEnumValueInsideNestedListItem(t *testing.T) {
	jsonStr := `{
		"observations": [],
		"summaries": [],
		"communications": [
			{
				"id": "comm_001",
				"content": {"text": "Hello"},
				"author": {
					"id": "agent_001",
					"name": "Agent Smith",
					"type": "FUTURE_AGENT_TYPE",
					"address": "+15558881111",
					"channel": "sms"
				},
				"recipients": [
					{
						"id": "cust_001",
						"name": "John Doe",
						"type": "FUTURE_RECIPIENT_TYPE",
						"address": "+15559992222",
						"channel": "sms"
					}
				]
			}
		],
		"meta": {"queryTime": 42}
	}`

	var recall memoryApi.MemoryRetrievalResponse
	err := json.Unmarshal([]byte(jsonStr), &recall)

	assert.NoError(t, err)
	assert.Len(t, recall.Communications, 1)
	assert.Equal(t, memoryApi.ParticipantType("FUTURE_AGENT_TYPE"), recall.Communications[0].Author.Type)
	assert.Equal(t, memoryApi.ParticipantType("FUTURE_RECIPIENT_TYPE"), recall.Communications[0].Recipients[0].Type)
	assert.Equal(t, "Agent Smith", recall.Communications[0].Author.Name)
	assert.Equal(t, "John Doe", recall.Communications[0].Recipients[0].Name)
	assert.Equal(t, "comm_001", recall.Communications[0].Id)
}
