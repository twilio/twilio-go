// Package taskrouter provides CRUD library for taskrouter subresources.
package taskrouter

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	twilio "github.com/twilio/twilio-go/internal"
)

// refer: https://www.twilio.com/docs/taskrouter/api/activity

// Activity activity.
type Activity struct {
	AccountSid   string    `json:"account_sid"`
	Available    string    `json:"available"`
	DateCreated  time.Time `json:"date_created"`
	DateUpdated  time.Time `json:"date_updated"`
	FriendlyName string    `json:"friendly_name"`
	Sid          string    `json:"sid"`
	WorkspaceSid string    `json:"workspace_sid"`
	URI          string    `json:"url"`
}

// ActivityParams activity params to create/update activity.
type ActivityParams struct {
	Available    string `url:"Available,omitempty"`
	FriendlyName string `url:"FriendlyName,omitempty"`
	WorkspaceSid string `url:"WorkspaceSid,omitempty"`
}

// ActivityClient is the entrypoint for activity CRUD.
type ActivityClient struct {
	ServiceURL string
	Client     *twilio.Client
}

// NewActivityClient constructs a new Activity Client.
func NewActivityClient(twilioClient *twilio.Client) *ActivityClient {
	c := new(ActivityClient)
	c.Client = twilioClient
	c.ServiceURL = fmt.Sprintf("https://taskrouter.%s/v1/Workspaces", twilioClient.BaseURL)

	return c
}

// Create creates activity with the given the config.
func (ac *ActivityClient) Create(activityParams ActivityParams, workspaceSID string) (*Activity, error) {
	url := fmt.Sprintf("%s/%s/%s", ac.ServiceURL, workspaceSID, "Activities")

	resp, err := ac.Client.Post(url, activityParams)

	if err != nil {
		log.Printf("error creating activity: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	activity := &Activity{}

	if decodeErr := json.NewDecoder(resp.Body).Decode(activity); decodeErr != nil {
		log.Printf("error decoding the output of activity create: %s", decodeErr)
		return nil, decodeErr
	}

	return activity, nil
}

// Fetch fetches activity for the given activity SID.
func (ac *ActivityClient) Fetch(workspaceSID string, activitySID string) (*Activity, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", ac.ServiceURL, workspaceSID, "Activities", activitySID)
	resp, err := ac.Client.Get(url, nil)

	if err != nil {
		log.Printf("error fetching activity: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	activity := &Activity{}

	if decodeErr := json.NewDecoder(resp.Body).Decode(activity); decodeErr != nil {
		log.Printf("error decoding the output of activity fetch: %s", decodeErr)
		return nil, decodeErr
	}

	return activity, nil
}

// Read returns all existing activities for a workspace.
func (ac *ActivityClient) Read(workspaceSID string) (*[]Activity, error) {
	url := fmt.Sprintf("%s/%s/%s", ac.ServiceURL, workspaceSID, "Activities")

	resp, err := ac.Client.Get(url, nil)

	if err != nil {
		log.Printf("error reading activities: %s", err)
		return nil, err
	}

	activities := make([]Activity, 0)

	if decodeErr := json.NewDecoder(resp.Body).Decode(&activities); decodeErr != nil {
		log.Printf("error decoding the output of activity read: %s", decodeErr)
		return nil, decodeErr
	}

	return &activities, nil
}

// Update updates activity with given config.
func (ac *ActivityClient) Update(activityParams ActivityParams, workspaceSID string,
	activitySID string) (*Activity, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", ac.ServiceURL, workspaceSID, "Activities", activitySID)

	resp, err := ac.Client.Post(url, activityParams)

	if err != nil {
		log.Printf("error updating activity: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	activity := &Activity{}

	if decodeErr := json.NewDecoder(resp.Body).Decode(activity); decodeErr != nil {
		log.Printf("error decoding the output of activity update: %s", decodeErr)
		return nil, decodeErr
	}

	return activity, nil
}

// Delete deletes workflow for given SID.
func (ac *ActivityClient) Delete(workspaceSID string, activitySID string) error {
	url := fmt.Sprintf("%s/%s/%s/%s", ac.ServiceURL, workspaceSID, "Activities", activitySID)

	_, err := ac.Client.Delete(url)

	return err
}
