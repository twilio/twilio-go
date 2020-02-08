package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// An Activity describes the current status of a Worker,
// which determines whether they are eligible to receive task assignments.
// refer: https://www.twilio.com/docs/taskrouter/api/activity
type Activity struct {
	AccountSid   string    `json:"account_sid"`
	Available    bool      `json:"available"`
	DateCreated  time.Time `json:"date_created"`
	DateUpdated  time.Time `json:"date_updated"`
	FriendlyName string    `json:"friendly_name"`
	Sid          string    `json:"sid"`
	WorkspaceSid string    `json:"workspace_sid"`
	URI          string    `json:"url"`
}

// ActivityParams activity params to create/update activity.
type ActivityParams struct {
	Available    *string `url:"Available,omitempty"`
	FriendlyName string  `url:"FriendlyName,omitempty"`
}

// ActivityList struct to parse response of activity read.
type ActivityList struct {
	Activities *[]Activity `json:"activities"`
	Meta       *Meta       `json:"meta,omitempty"`
}

// ActivityQueryParams query params to read workspaces.
type ActivityQueryParams struct {
	FriendlyName *string `url:"FriendlyName,omitempty"`
	Available    *string `url:"Available,omitempty"`
	PageSize     *int    `url:"PageSize,omitempty"`
}

// ActivityClient is the entrypoint for activity CRUD.
type ActivityClient struct {
	ServiceURL string
	client     *Twilio
}

// NewActivityClient constructs a new Activity Client.
func NewActivityClient(client *Twilio) *ActivityClient {
	c := new(ActivityClient)
	c.client = client
	c.ServiceURL = fmt.Sprintf("https://taskrouter.%s/v1/Workspaces", client.BaseURL)

	return c
}

// Create creates activity with the given the config.
func (c *ActivityClient) Create(workspaceSID string, activityParams *ActivityParams) (*Activity, error) {
	url := fmt.Sprintf("%s/%s/%s", c.ServiceURL, workspaceSID, "Activities")

	if len(activityParams.FriendlyName) == 0 {
		return nil, errors.New("friendlyname is required in activityparams")
	}

	resp, err := c.client.Post(url, activityParams)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	activity := &Activity{}

	if err = json.NewDecoder(resp.Body).Decode(activity); err != nil {
		return nil, err
	}

	return activity, nil
}

// Fetch fetches activity for the given activity SID.
func (c *ActivityClient) Fetch(workspaceSID string, activitySID string) (*Activity, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", c.ServiceURL, workspaceSID, "Activities", activitySID)
	resp, err := c.client.Get(url, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	activity := &Activity{}

	if err = json.NewDecoder(resp.Body).Decode(activity); err != nil {
		return nil, err
	}

	return activity, nil
}

// Read returns all existing activities for a workspace.
func (c *ActivityClient) Read(workspaceSID string) (*ActivityList, error) {
	url := fmt.Sprintf("%s/%s/%s", c.ServiceURL, workspaceSID, "Activities")

	resp, err := c.client.Get(url, nil)

	if err != nil {
		return nil, err
	}

	activities := &ActivityList{}

	if err = json.NewDecoder(resp.Body).Decode(&activities); err != nil {
		return nil, err
	}

	return activities, nil
}

// Update updates activity with given config.
func (c *ActivityClient) Update(workspaceSID string,
	activitySID string, activityParams *ActivityParams) (*Activity, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", c.ServiceURL, workspaceSID, "Activities", activitySID)

	resp, err := c.client.Post(url, activityParams)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	activity := &Activity{}

	if err = json.NewDecoder(resp.Body).Decode(activity); err != nil {
		return nil, err
	}

	return activity, nil
}

// Delete deletes workflow for given SID.
func (c *ActivityClient) Delete(workspaceSID string, activitySID string) error {
	url := fmt.Sprintf("%s/%s/%s/%s", c.ServiceURL, workspaceSID, "Activities", activitySID)

	resp, err := c.client.Delete(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
