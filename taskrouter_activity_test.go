package twilio

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTaskrouterActivity_Create(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/Activities", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{"FriendlyName": "TaskRouterActivity"})
		response := `{"friendly_name":"TaskRouterActivity"}`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.Activities.Create("WS123", &TaskRouterActivityParams{FriendlyName: String("TaskRouterActivity")})

	if err != nil {
		t.Errorf("TaskRouterActivity.Create returned error: %v", err)
	}

	want := &TaskRouterActivity{FriendlyName: String("TaskRouterActivity")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TaskRouterActivity.Create returned %+v, want %+v", got, want)
	}

}

func TestTaskRouterActivity_Fetch(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/Activities/WA123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"sid":"WA123"}`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.Activities.Fetch("WS123", "WA123")

	if err != nil {
		t.Errorf("TaskRouterActivity.Fetch returned error: %v", err)
	}

	want := &TaskRouterActivity{SID: String("WA123")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TaskRouterActivity.Fetch returned %+v, want %+v", got, want)
	}
}

func TestTaskRouterActivity_Read(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/Activities", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		response := `{"activities":[{"sid": "WA123"}]}`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.Activities.Read("WS123")

	if err != nil {
		t.Errorf("TaskRouterActivity.Read returned error: %v", err)
	}

	activity := &TaskRouterActivity{SID: String("WA123")}
	want := &TaskRouterActivityList{Activities: []*TaskRouterActivity{activity}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TaskRouterActivity.Read returned %+v, want %+v", got, want)
	}
}

func TestTaskRouterActivity_Update(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/Activities/WA123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		response := `{"sid":"WA123","friendly_name":"NewName"}`

		fmt.Fprint(w, response)
	})

	got, err := client.TaskRouter.Activities.Update("WS123", "WA123", &TaskRouterActivityParams{
		FriendlyName: String("NewName"),
	})

	if err != nil {
		t.Errorf("TaskRouterActivity.Update returned error: %v", err)
	}

	want := &TaskRouterActivity{SID: String("WA123"), FriendlyName: String("NewName")}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TaskRouterActivity.Update returned %+v, want %+v", got, want)
	}

}

func TestTaskRouterActivity_Delete(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/Workspaces/WS123/Activities/WA123", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	err := client.TaskRouter.Activities.Delete("WS123", "WA123")

	if err != nil {
		t.Errorf("TaskRouterActivity.Delete returned error: %v", err)
	}
}
