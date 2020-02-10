package twilio

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// TaskRouterActivity describes the current status of a Worker,
// which determines whether they are eligible to receive task assignments.
// refer: https://www.twilio.com/docs/taskrouter/api/activity
type TaskRouterActivity struct {
	AccountSid   *string    `json:"account_sid"`
	Available    *bool      `json:"available"`
	DateCreated  *time.Time `json:"date_created"`
	DateUpdated  *time.Time `json:"date_updated"`
	FriendlyName *string    `json:"friendly_name"`
	Sid          *string    `json:"sid"`
	WorkspaceSid *string    `json:"workspace_sid"`
	URI          *string    `json:"url"`
}

// TaskRouterActivityParams activity params to create/update activity.
type TaskRouterActivityParams struct {
	Available    *bool   `form:",omitempty"`
	FriendlyName *string `form:",omitempty"`
}

// TaskRouterActivityList struct to parse response of activity read.
type TaskRouterActivityList struct {
	Activities []*TaskRouterActivity `json:"activities"`
	Meta       *Meta                 `json:"meta,omitempty"`
}

// TaskRouterActivityQueryParams query params to read workspaces.
type TaskRouterActivityQueryParams struct {
	FriendlyName *string `form:",omitempty"`
	Available    *string `form:",omitempty"`
	PageSize     *int    `form:",omitempty"`
}

// TaskRouterActivityClient is the entrypoint for activity CRUD.
type TaskRouterActivityClient struct {
	baseURL string
	client  *Twilio
}

// NewTaskRouterActivityClient constructs a new Activity Client.
func NewTaskRouterActivityClient(client *Twilio) *TaskRouterActivityClient {
	c := new(TaskRouterActivityClient)
	c.client = client
	c.baseURL = fmt.Sprintf("https://taskrouter.%s/v1/", client.BaseURL)

	return c
}

// Create creates activity with the given the config.
func (c *TaskRouterActivityClient) Create(workspaceSID string, activityParams *TaskRouterActivityParams) (*TaskRouterActivity, error) {
	url := c.url(fmt.Sprintf("/Workspaces/%s/Activities", workspaceSID))

	if len(*activityParams.FriendlyName) == 0 {
		return nil, errors.New("friendlyname is required in activityparams")
	}

	resp, err := c.client.Post(url, activityParams)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	activity := &TaskRouterActivity{}

	if err = json.NewDecoder(resp.Body).Decode(activity); err != nil {
		return nil, err
	}

	return activity, nil
}

// Fetch fetches activity for the given activity SID.
func (c *TaskRouterActivityClient) Fetch(workspaceSID string, SID string) (*TaskRouterActivity, error) {
	url := c.url(fmt.Sprintf("/Workspaces/%s/Activities/%s", workspaceSID, SID))
	resp, err := c.client.Get(url, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	activity := &TaskRouterActivity{}

	if err = json.NewDecoder(resp.Body).Decode(activity); err != nil {
		return nil, err
	}

	return activity, nil
}

// Read returns all existing activities for a workspace.
func (c *TaskRouterActivityClient) Read(workspaceSID string) (*TaskRouterActivityList, error) {
	url := c.url(fmt.Sprintf("/Workspaces/%s/Activities", workspaceSID))

	resp, err := c.client.Get(url, nil)

	if err != nil {
		return nil, err
	}

	activities := &TaskRouterActivityList{}

	if err = json.NewDecoder(resp.Body).Decode(&activities); err != nil {
		return nil, err
	}

	return activities, nil
}

// Update updates activity with given config.
func (c *TaskRouterActivityClient) Update(
	workspaceSID string,
	SID string,
	activityParams *TaskRouterActivityParams,
) (*TaskRouterActivity, error) {
	url := c.url(fmt.Sprintf("/Workspaces/%s/Activities/%s", workspaceSID, SID))

	resp, err := c.client.Post(url, activityParams)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	activity := &TaskRouterActivity{}

	if err = json.NewDecoder(resp.Body).Decode(activity); err != nil {
		return nil, err
	}

	return activity, nil
}

// Delete deletes workflow for given SID.
func (c *TaskRouterActivityClient) Delete(workspaceSID string, SID string) error {
	url := c.url(fmt.Sprintf("/Workspaces/%s/Activities/%s", workspaceSID, SID))

	resp, err := c.client.Delete(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}

func (c *TaskRouterActivityClient) url(path string) string {
	if c.client.defaultbaseURL != nil {
		return *c.client.defaultbaseURL + path
	}

	return c.baseURL + path
}
