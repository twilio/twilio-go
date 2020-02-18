package twilio

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestTaskRouterTaskQueue_marshall(t *testing.T) {
	testJSONMarshal(t, TaskRouterTaskQueue{}, "{}")

	got := &TaskRouterTaskQueue{
		AccountSID:             String("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		AssignmentActivityName: String("817ca1c5-3a05-11e5-9292-98e0d9a1eb73"),
		AssignmentActivitySID:  String("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		DateCreated:            &time.Time{},
		DateUpdated:            &time.Time{},
		FriendlyName:           String("English"),
		MaxReservedWorkers:     Int(1),
		Links: map[string]*string{
			"assignment_activity": String("/Activities/WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		},
		ReservationActivityName: String("80fa2beb-3a05-11e5-8fc8-98e0d9a1eb73"),
		ReservationActivitySID:  String("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		SID:                     String("WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		TargetWorkers:           String("languages HAS \"english\""),
		TaskOrder:               String("FIFO"),
		URL:                     String("/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		WorkspaceSID:            String("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
	}

	want := `{
		"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"assignment_activity_name": "817ca1c5-3a05-11e5-9292-98e0d9a1eb73",
		"assignment_activity_sid": "WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"date_created": "0001-01-01T00:00:00Z",
		"date_updated": "0001-01-01T00:00:00Z",
		"friendly_name": "English",
		"max_reserved_workers": 1,
		"links": {
		  "assignment_activity": "/Activities/WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
		},
		"reservation_activity_name": "80fa2beb-3a05-11e5-8fc8-98e0d9a1eb73",
		"reservation_activity_sid": "WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"sid": "WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"target_workers": "languages HAS \"english\"",
		"task_order": "FIFO",
		"url": "/TaskQueues/WQXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"workspace_sid": "WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	  }`

	testJSONMarshal(t, got, want)
}

func TestTaskrouterTaskQueue_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/TaskQueues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		f := url.Values{}
		f.Add("FriendlyName", "a")
		f.Add("AssignmentActivitySid", "b")
		f.Add("MaxReservedWorkers", "10")
		f.Add("TargetWorkers", "c")
		f.Add("TaskOrder", "d")
		f.Add("ReservationActivitySid", "e")

		testFormValues(t, r, f)
		response := `{"friendly_name":"TaskRouterTaskQueue"}`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.TaskQueues.Create("WS123", &TaskRouterTaskQueueParams{
		FriendlyName:           String("a"),
		AssignmentActivitySID:  String("b"),
		MaxReservedWorkers:     Int(10),
		TargetWorkers:          String("c"),
		TaskOrder:              String("d"),
		ReservationActivitySID: String("e"),
	})

	if err != nil {
		t.Errorf("TaskRouterTaskQueue.Create returned error: %v", err)
	}

	want := &TaskRouterTaskQueue{FriendlyName: String("TaskRouterTaskQueue")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TaskRouterTaskQueue.Create returned %+v, want %+v", got, want)
	}
}

func TestTaskRouterTaskQueue_Fetch(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/TaskQueues/WA123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"WA123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.TaskQueues.Fetch("WS123", "WA123")

	if err != nil {
		t.Errorf("TaskRouterTaskQueue.Fetch returned error: %v", err)
	}

	want := &TaskRouterTaskQueue{SID: String("WA123")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TaskRouterTaskQueue.Fetch returned %+v, want %+v", got, want)
	}
}

func TestTaskRouterTaskQueue_Read(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/TaskQueues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"task_queues":[{"sid": "WA123"}]}`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.TaskQueues.Read("WS123", nil)

	if err != nil {
		t.Errorf("TaskRouterTaskQueue.Read returned error: %v", err)
	}

	TaskQueue := &TaskRouterTaskQueue{SID: String("WA123")}
	want := &TaskRouterTaskQueueList{TaskQueues: []*TaskRouterTaskQueue{TaskQueue}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TaskRouterTaskQueue.Read returned %+v, want %+v", got, want)
	}
}

func TestTaskRouterTaskQueue_Update(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/TaskQueues/WA123", func(w http.ResponseWriter, r *http.Request) {
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

	want := &TaskRouterTaskQueue{SID: String("WA123"), FriendlyName: String("NewName")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TaskRouterTaskQueue.Update returned %+v, want %+v", got, want)
	}
}

func TestTaskRouterTaskQueue_Delete(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/TaskQueues/WA123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.TaskRouter.TaskQueues.Delete("WS123", "WA123")

	if err != nil {
		t.Errorf("TaskRouterTaskQueue.Delete returned error: %v", err)
	}
}
