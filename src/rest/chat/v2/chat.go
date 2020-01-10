package chat

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	twilio "../../../base"
)

type Chat struct {
	Request twilio.Request
}

type ChatService struct {
	Sid                          string            `json:"sid"`
	AccountSid                   string            `json:"account_sid"`
	FriendlyName                 string            `json:"friendly_name"`
	DateCreated                  time.Time         `json:"date_created"`
	DateUpdated                  time.Time         `json:"date_updated"`
	DefaultServiceRoleSid        string            `json:"default_service_role_sid"`
	DefaultChannelRoleSid        string            `json:"default_channel_role_sid"`
	DefaultChannelCreatorRoleSid string            `json:"default_channel_creator_role_sid"`
	ReadStatusEnabled            string            `json:"read_status_enabled"`
	ReachabilityEnabled          string            `json:"reachability_enabled"`
	TypingIndicatorTimeout       string            `json:"typing_indicator_timeout"`
	ConsumptionReportInterval    string            `json:"consumption_report_interval"`
	Limits                       map[string]int    `json:"limits"`
	PreWebhookUrl                string            `json:"pre_webhook_url"`
	PostWebhookUrl               string            `json:"post_webhook_url"`
	WebhookMethod                string            `json:"webhook_method"`
	WebhookFilters               string            `json:"webhook_filters"`
	PreWebhookRetryCount         int               `json:"pre_webhook_retry_count"`
	PostWebhookRetryCount        int               `json:"post_webhook_retry_count"`
	Notifications                map[string]string `json:"notifications"`
	Media                        map[string]string `json:"media"`
	Url                          string            `json:"url"`
	Links                        map[string]string `json:"links"`
}

func (chat *Chat) Create(friendlyName string) (ChatService, error) {
	values := url.Values{}
	values.Set("FriendlyName", friendlyName)

	resp, err := chat.Request.Post("/Services", values)
	defer resp.Body.Close()
	cs := ChatService{}
	json.NewDecoder(resp.Body).Decode(cs)
	return cs, err

}

func (chat *Chat) Read(sid string) (*ChatService, error) {
	resp, err := chat.Request.Get(fmt.Sprintf("/Services/%s", sid))
	defer resp.Body.Close()
	cs := &ChatService{}
	json.NewDecoder(resp.Body).Decode(cs)
	return cs, err
}
