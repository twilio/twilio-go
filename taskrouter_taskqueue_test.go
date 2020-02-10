package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTaskrouterTaskQueue_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/TaskRouterTaskQueues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"FriendlyName": "TaskRouterTaskQueue"})
		response := `{"friendly_name":"TaskRouterTaskQueue"}`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.TaskQueues.Create("WS123", &TaskRouterTaskQueueParams{FriendlyName: String("TaskRouterTaskQueue")})

	if err != nil {
		t.Errorf("TaskRouterTaskQueue.Create returned error: %v", err)
	}

	expected := &TaskRouterTaskQueue{FriendlyName: String("TaskRouterTaskQueue")}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("TaskRouterTaskQueue.Create returned %+v, expected %+v", got, expected)
	}

}

func TestTaskRouterTaskQueue_Fetch(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/TaskRouterTaskQueues/WA123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"WA123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.TaskQueues.Fetch("WS123", "WA123")

	if err != nil {
		t.Errorf("TaskRouterTaskQueue.Fetch returned error: %v", err)
	}

	expected := &TaskRouterTaskQueue{Sid: String("WA123")}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("TaskRouterTaskQueue.Fetch returned %+v, expected %+v", got, expected)
	}
}

func TestTaskRouterTaskQueue_Read(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/TaskRouterTaskQueues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"task_queues":[{"sid": "WA123"}]}`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.TaskQueues.Read("WS123", nil)

	if err != nil {
		t.Errorf("TaskRouterTaskQueue.Read returned error: %v", err)
	}

	TaskQueue := &TaskRouterTaskQueue{Sid: String("WA123")}
	expected := &TaskRouterTaskQueueList{TaskRouterTaskQueues: []*TaskRouterTaskQueue{TaskQueue}}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("TaskRouterTaskQueue.Read returned %+v, expected %+v", got, expected)
	}
}

func TestTaskRouterTaskQueue_Update(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/TaskRouterTaskQueues/WA123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"sid":"WA123","friendly_name":"NewName"}`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.TaskQueues.Update("WS123", "WA123", &TaskRouterTaskQueueParams{
		FriendlyName: String("NewName"),
	})

	if err != nil {
		t.Errorf("TaskRouterTaskQueue.Update returned error: %v", err)
	}

	expected := &TaskRouterTaskQueue{Sid: String("WA123"), FriendlyName: String("NewName")}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("TaskRouterTaskQueue.Update returned %+v, expected %+v", got, expected)
	}

}

func TestTaskRouterTaskQueue_Delete(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/TaskRouterTaskQueues/WA123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.TaskRouter.TaskQueues.Delete("WS123", "WA123")

	if err != nil {
		t.Errorf("TaskRouterTaskQueue.Delete returned error: %v", err)
	}
}
