// Package taskrouter provides CRUD library for taskrouter subresources.
package taskrouter

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	twilio "github.com/twilio/twilio-go/internal"
)

// refer: https://www.twilio.com/docs/taskrouter/api/task-queue

// TaskQueue taskqueue struct.
type TaskQueue struct {
	AccountSid              string            `json:"account_sid"`
	AssignmentActivitySid   string            `json:"assignment_activity_sid"`
	AssignmentActivityName  string            `json:"assignment_activity_name"`
	DateCreated             time.Time         `json:"date_created"`
	DateUpdated             time.Time         `json:"date_updated"`
	FriendlyName            string            `json:"friendly_name"`
	MaxReservedWorkers      int               `json:"max_reserved_workers"`
	ReservationActivitySid  string            `json:"reservation_activity_sid"`
	ReservationActivityName string            `json:"reservation_activity_name"`
	Sid                     string            `json:"sid"`
	TargetWorkers           string            `json:"target_workers"`
	TaskOrder               string            `json:"task_order"`
	URI                     string            `json:"url"`
	WorkspaceSid            string            `json:"workspace_sid"`
	Links                   map[string]string `json:"links"`
}

// TaskQueueParams taskQueue parameters.
type TaskQueueParams struct {
	FriendlyName           string `url:"FriendlyName,omitempty"`
	AssignmentActivitySid  string `url:"AssignmentActivitySid,omitempty"`
	MaxReservedWorkers     int    `url:"MaxReservedWorkers,omitempty"`
	TargetWorkers          string `url:"TargetWorkers,omitempty"`
	TaskOrder              string `url:"TaskOrder,omitempty"`
	ReservationActivitySid string `url:"ReservationActivitySid,omitempty"`
}

// TaskQueueClient is the entrypoint for taskqueue CRUD.
type TaskQueueClient struct {
	ServiceURL string
	Client     *twilio.Client
}

// NewTaskQueueClient constructs a new TaskQueue Client.
func NewTaskQueueClient(twilioClient *twilio.Client) *TaskQueueClient {
	c := new(TaskQueueClient)
	c.Client = twilioClient
	c.ServiceURL = fmt.Sprintf("https://taskrouter.%s/v1/Workspaces", twilioClient.BaseURL)

	return c
}

// Create creates taskqueue with the given the config.
func (ws *TaskQueueClient) Create(workspaceSID string, taskqueueparams TaskQueueParams) (*TaskQueue, error) {
	url := fmt.Sprintf("%s/%s/%s", ws.ServiceURL, workspaceSID, "TaskQueues")

	resp, err := ws.Client.Post(url, taskqueueparams)

	if err != nil {
		log.Printf("error creating taskqueue: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	taskQueue := &TaskQueue{}

	if decodeErr := json.NewDecoder(resp.Body).Decode(taskQueue); decodeErr != nil {
		log.Printf("error decoding the output of taskqueue create: %s", decodeErr)
		return nil, decodeErr
	}

	return taskQueue, nil
}

// Fetch fetches taskqueue for the given SID.
func (ws *TaskQueueClient) Fetch(workspaceSID string, taskQueueSID string) (*TaskQueue, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", ws.ServiceURL, workspaceSID, "TaskQueues", taskQueueSID)
	resp, err := ws.Client.Get(url, nil)

	if err != nil {
		log.Printf("error fetching the requested taskqueue. workflowSID:%s,taskQueue:%s error:%s",
			workspaceSID, taskQueueSID, err)
		return nil, err
	}

	defer resp.Body.Close()

	taskqueue := &TaskQueue{}

	if decodeErr := json.NewDecoder(resp.Body).Decode(taskqueue); decodeErr != nil {
		log.Printf("error decoding the output of taskqueue fetch: %s", decodeErr)
		return nil, decodeErr
	}

	return taskqueue, nil
}

// Read returns all existing taskqueues.
func (ws *TaskQueueClient) Read(workspaceSID string) (*[]TaskQueue, error) {
	url := fmt.Sprintf("%s/%s/%s", ws.ServiceURL, workspaceSID, "TaskQueues")

	resp, err := ws.Client.Get(url, nil)

	if err != nil {
		log.Printf("error fetching taskQueues: %s", err)
		return nil, err
	}

	taskqueues := make([]TaskQueue, 0)

	if decodeErr := json.NewDecoder(resp.Body).Decode(&taskqueues); decodeErr != nil {
		log.Printf("error decoding the output of taskQueue fetch: %s", decodeErr)
		return nil, decodeErr
	}

	return &taskqueues, nil
}

// Update updates taskqueue with given config.
func (ws *TaskQueueClient) Update(taskQueueParams TaskQueueParams, workspaceSID string,
	taskQueueSID string) (*TaskQueue, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", ws.ServiceURL, workspaceSID, "TaskQueues", taskQueueSID)

	resp, err := ws.Client.Post(url, taskQueueParams)

	if err != nil {
		log.Printf("error updating taskQueue: %s", err)
		return nil, err
	}

	defer resp.Body.Close()

	taskQueue := &TaskQueue{}

	if decodeErr := json.NewDecoder(resp.Body).Decode(taskQueue); decodeErr != nil {
		log.Printf("error decoding the output of taskqueue fetch: %s", decodeErr)
		return nil, decodeErr
	}

	return taskQueue, nil
}

// Delete deletes taskQueue with the given SID.
func (ws *TaskQueueClient) Delete(workspaceSID string, taskqueueSID string) error {
	url := fmt.Sprintf("%s/%s/%s/%s", ws.ServiceURL, workspaceSID, "TaskQueues", taskqueueSID)

	_, err := ws.Client.Delete(url)

	return err
}
