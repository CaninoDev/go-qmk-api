package qmk

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestClient_KeyboardsList(t *testing.T) {
	wantResponse := [2]string{"keyboard1", "eyboard2"}
	client, server := createEndpoint("/keyboards", `["keyboard1", "eyboard2"]`)

	testClient := createNewCustomClient(client)
	gotResponse, err := testClient.KeyboardsList()

	defer server.Close()

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if wantResponse[0] != gotResponse[0] {
		t.Errorf("expected %v, got %v", wantResponse[0], gotResponse[0])
	}

}

func TestClient_KeyboardData(t *testing.T) {
	const expectedKeyboardMaintainer = "exampleMaintainer"

	client, server := createEndpoint("/keyboards/exampleKeyboard", `{"maintainer": "exampleMaintainer", "readme": true}`)
	defer server.Close()

	testClient := createNewCustomClient(client)
	keyboardResponse, err := testClient.KeyboardData("exampleKeyboard")

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if !keyboardResponse.Readme {
		t.Error("expected true, got false")
	}
	if keyboardResponse.Maintainer != expectedKeyboardMaintainer {
		t.Errorf("expected %s, got %s", expectedKeyboardMaintainer, keyboardResponse.Maintainer)
	}
}

func TestClient_KeyboardReadme(t *testing.T) {
	var expectedReadme = `This is an example readme text`

	client, mux, server := testServer()
	defer server.Close()
	mux.HandleFunc("/keyboards/exampleKeyboard/readme", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/markdown")
		fmt.Fprintf(w, `This is an example readme text`)
	})

	testClient := createNewCustomClient(client)
	readmeResponse, err := testClient.KeyboardReadme("exampleKeyboard")

	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}
	if !cmp.Equal(readmeResponse, expectedReadme) {
		t.Errorf("expected %v, got %v", expectedReadme, readmeResponse)
	}

	_, err = testClient.KeyboardReadme("noExampleKeyboard")

	if err == nil {
		t.Errorf("expected error, got %s", err)
	}

}

func TestClient_KeymapData(t *testing.T) {
	var wantKeymapData = &Keymap{
		Name: "exampleKeymap",
	}
	client, server := createEndpoint("/keyboards/exampleKeyboard/keymaps/exampleKeymap", `{"keymap_name": "exampleKeymap"}`)
	defer server.Close()

	testClient := createNewCustomClient(client)
	gotResponse, err := testClient.KeymapData("exampleKeyboard", "exampleKeymap")

	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}
	if wantKeymapData.Name != gotResponse.Name {
		t.Errorf("expected %v, got %v", wantKeymapData, gotResponse)
	}

	_, err = testClient.KeymapData("noExampleKeyboard", "noExampleKeymap")

	if err == nil {
		t.Errorf("expected error, got %s", err)
	}

}

func TestClient_KeymapReadme(t *testing.T) {
	var expectedReadme = `##This is an example readme text`

	client, mux, server := testServer()
	defer server.Close()
	mux.HandleFunc("/keyboards/exampleKeyboard/keymaps/exampleKeymap/readme", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/markdown")
		fmt.Fprintf(w, `##This is an example readme text`)
	})

	testClient := createNewCustomClient(client)
	readmeResponse, err := testClient.KeymapReadme("exampleKeyboard", "exampleKeymap")

	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}
	if !cmp.Equal(readmeResponse, expectedReadme) {
		t.Errorf("expected %v, got %v", expectedReadme, readmeResponse)
	}

	_, err = testClient.KeymapReadme("noExampleKeyboard", "noExampleKeymap")

	if err == nil {
		t.Errorf("expected error, got %s", err)
	}
}
