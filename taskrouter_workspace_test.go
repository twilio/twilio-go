package twilio

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func TestTaskRouterWorkspace_marshall(t *testing.T) {
	testJSONMarshal(t, TaskRouterWorkspace{}, "{}")

	got := &TaskRouterWorkspace{
		AccountSID:          String("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		DateCreated:         &time.Time{},
		DateUpdated:         &time.Time{},
		DefaultActivityName: String("Offline"),
		DefaultActivitySID:  String("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		EventCallbackURL:    String("https://workspace-example.free.beeceptor.com"),
		EventsFilter:        nil,
		FriendlyName:        String("NewWorkspace"),
		Links: map[string]*string{
			"events": String("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Events"),
		},
		MultitaskEnabled:     Bool(false),
		PrioritizeQueueOrder: String("FIFO"),
		SID:                  String("WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		TimeoutActivityName:  String("Offline"),
		TimeoutActivitySID:   String("WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
		URL:                  String("https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"),
	}

	want := `{
		"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"date_created": "0001-01-01T00:00:00Z",
		"date_updated": "0001-01-01T00:00:00Z",
		"default_activity_name": "Offline",
		"default_activity_sid": "WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"event_callback_url": "https://workspace-example.free.beeceptor.com",
		"events_filter": null,
		"friendly_name": "NewWorkspace",
		"links": {
			"events": "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Events"
		},
		"multi_task_enabled": false,
		"prioritize_queue_order": "FIFO",
		"sid": "WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"timeout_activity_name": "Offline",
		"timeout_activity_sid": "WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
		"url": "https://taskrouter.twilio.com/v1/Workspaces/WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	  }`

	testJSONMarshal(t, got, want)
}

func TestTaskrouterWorkspace_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		f := url.Values{}
		f.Add("FriendlyName", "a")
		f.Add("EventCallbackUrl", "b")
		f.Add("EventsFilter", "c")
		f.Add("MultitaskEnabled", "true")
		f.Add("Template", "d")
		f.Add("PrioritizeQueueOrder", "e")
		testFormValues(t, r, f)
		response := `{"friendly_name":"TaskRouterWorkspace"}`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.Workspaces.Create(&TaskRouterWorkspaceParams{
		FriendlyName:         String("a"),
		EventCallbackURL:     String("b"),
		EventsFilter:         String("c"),
		MultitaskEnabled:     Bool(true),
		Template:             String("d"),
		PrioritizeQueueOrder: String("e"),
	})

	if err != nil {
		t.Errorf("TaskRouterWorkspace.Create returned error: %v", err)
	}

	want := &TaskRouterWorkspace{FriendlyName: String("TaskRouterWorkspace")}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TaskRouterWorkspace.Create returned %+v, want %+v", got, want)
	}

}

func TestTaskRouterWorkspace_Fetch(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"WS123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.Workspaces.Fetch("WS123")

	if err != nil {
		t.Errorf("TaskRouterWorkspace.Fetch returned error: %v", err)
	}

	want := &TaskRouterWorkspace{SID: String("WS123")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TaskRouterWorkspace.Fetch returned %+v, want %+v", got, want)
	}
}

func TestTaskRouterWorkspace_Read(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"workspaces":[{"sid": "WS123"}]}`
		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.Workspaces.Read(nil)
	if err != nil {
		t.Errorf("TaskRouterWorkspace.Read returned error: %v", err)
	}

	workspace := &TaskRouterWorkspace{SID: String("WS123")}
	want := &TaskRouterWorkspaceList{Workspaces: []*TaskRouterWorkspace{workspace}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TaskRouterWorkspace.Read returned %+v, want %+v", got, want)
	}
}

func TestTaskRouterWorkspace_Update(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"sid":"WS123","friendly_name":"NewName"}`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.Workspaces.Update("WS123", &TaskRouterWorkspaceParams{
		FriendlyName: String("NewName"),
	})

	if err != nil {
		t.Errorf("TaskRouterWorkspace.Update returned error: %v", err)
	}

	want := &TaskRouterWorkspace{SID: String("WS123"), FriendlyName: String("NewName")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TaskRouterWorkspace.Update returned %+v, want %+v", got, want)
	}

}

func TestTaskRouterWorkspace_Delete(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.TaskRouter.Workspaces.Delete("WS123")

	if err != nil {
		t.Errorf("TaskRouterWorkspace.Delete returned error: %v", err)
	}
}
