package qmk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	defTimeOut = time.Second * 2
)

// Json/XML-tagged model struct
type FakeModel struct {
	Text          string  `json:"text,omitempty"`
	FavoriteCount int64   `json:"favorite_count,omitempty"`
	Temperature   float64 `json:"temperature,omitempty"`
}

func createNewTestDefClient() *Client {
	defClient := &http.Client{
		Timeout: time.Second * 2,
	}
	baseURL, _ := url.Parse(defaultBaseURL + "/" + apiVersion)
	return &Client{
		defClient,
		baseURL,
	}
}

func createNewCustomClient(customClient *http.Client) *Client {
	baseURL, _ := url.Parse(defaultBaseURL + "/" + apiVersion)
	return &Client{
		customClient,
		baseURL,
	}
}

// testServer returns an http Client, ServeMux, and Server. The client proxies
// requests to the server and handlers can be registered on the mux to handle
// requests. The caller must close the test server.
func testServer() (*http.Client, *http.ServeMux, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}
	client := &http.Client{Transport: transport}
	return client, mux, server
}

func TestNew(t *testing.T) {
	defhttpClient := createNewTestDefClient()
	testingDefClient := New(nil)
	if reflect.DeepEqual(testingDefClient, defhttpClient) {
		t.Errorf("expected %v, got %v", defhttpClient, testingDefClient)
	}
	customhttpClient := &http.Client{
		Timeout: time.Second * 4,
	}
	customClient := createNewCustomClient(customhttpClient)
	testingCustomClient := New(customhttpClient)
	if testingCustomClient != customClient {
		t.Errorf("expected %v, got %v", customClient, testingCustomClient)
	}
}

func TestClient_newRequest(t *testing.T) {
	rawPlaidKeyboard := []byte(`{"last_updated":"2019-12-07 12:26:33 UTC","git_hash":"f275ffbdfc1cbd1965cd3546b45a7838012321da","keyboards":{"plaid":{"height":4,"width":12,"description":"12x4 ortholinear keyboard with through hole components","keyboard_folder":"plaid","device_ver":"0x0002","manufacturer":"dm9records","processor":"atmega328p","identifier":"0x16C0:0x27DB:0x0002","product_id":"0x27DB","readme":true,"maintainer":"hsgw","vendor_id":"0x16C0","platform":"unknown","bootloader":"USBasp","keyboard_name":"Plaid // Through Hole","processor_type":"avr","keymaps":["thehalfdeafchef","brickbots","default"],"url":"https://github.com/hsgw/plaid","layouts":{"LAYOUT_plaid_grid":{"key_count":48,"layout":[{"x":0,"w":1,"y":0},{"x":1,"w":1,"y":0},{"x":2,"w":1,"y":0},{"x":3,"w":1,"y":0},{"x":4,"w":1,"y":0},{"x":5,"w":1,"y":0},{"x":6,"w":1,"y":0},{"x":7,"w":1,"y":0},{"x":8,"w":1,"y":0},{"x":9,"w":1,"y":0},{"x":10,"w":1,"y":0},{"x":11,"w":1,"y":0},{"x":0,"w":1,"y":1},{"x":1,"w":1,"y":1},{"x":2,"w":1,"y":1},{"x":3,"w":1,"y":1},{"x":4,"w":1,"y":1},{"x":5,"w":1,"y":1},{"x":6,"w":1,"y":1},{"x":7,"w":1,"y":1},{"x":8,"w":1,"y":1},{"x":9,"w":1,"y":1},{"x":10,"w":1,"y":1},{"x":11,"w":1,"y":1},{"x":0,"w":1,"y":2},{"x":1,"w":1,"y":2},{"x":2,"w":1,"y":2},{"x":3,"w":1,"y":2},{"x":4,"w":1,"y":2},{"x":5,"w":1,"y":2},{"x":6,"w":1,"y":2},{"x":7,"w":1,"y":2},{"x":8,"w":1,"y":2},{"x":9,"w":1,"y":2},{"x":10,"w":1,"y":2},{"x":11,"w":1,"y":2},{"x":0,"w":1,"y":3},{"x":1,"w":1,"y":3},{"x":2,"w":1,"y":3},{"x":3,"w":1,"y":3},{"x":4,"w":1,"y":3},{"x":5,"w":1,"y":3},{"x":6,"w":1,"y":3},{"x":7,"w":1,"y":3},{"x":8,"w":1,"y":3},{"x":9,"w":1,"y":3},{"x":10,"w":1,"y":3},{"x":11,"w":1,"y":3}]},"LAYOUT_ortho_4x12":{"key_count":48,"layout":[{"x":0,"w":1,"y":0},{"x":1,"w":1,"y":0},{"x":2,"w":1,"y":0},{"x":3,"w":1,"y":0},{"x":4,"w":1,"y":0},{"x":5,"w":1,"y":0},{"x":6,"w":1,"y":0},{"x":7,"w":1,"y":0},{"x":8,"w":1,"y":0},{"x":9,"w":1,"y":0},{"x":10,"w":1,"y":0},{"x":11,"w":1,"y":0},{"x":0,"w":1,"y":1},{"x":1,"w":1,"y":1},{"x":2,"w":1,"y":1},{"x":3,"w":1,"y":1},{"x":4,"w":1,"y":1},{"x":5,"w":1,"y":1},{"x":6,"w":1,"y":1},{"x":7,"w":1,"y":1},{"x":8,"w":1,"y":1},{"x":9,"w":1,"y":1},{"x":10,"w":1,"y":1},{"x":11,"w":1,"y":1},{"x":0,"w":1,"y":2},{"x":1,"w":1,"y":2},{"x":2,"w":1,"y":2},{"x":3,"w":1,"y":2},{"x":4,"w":1,"y":2},{"x":5,"w":1,"y":2},{"x":6,"w":1,"y":2},{"x":7,"w":1,"y":2},{"x":8,"w":1,"y":2},{"x":9,"w":1,"y":2},{"x":10,"w":1,"y":2},{"x":11,"w":1,"y":2},{"x":0,"w":1,"y":3},{"x":1,"w":1,"y":3},{"x":2,"w":1,"y":3},{"x":3,"w":1,"y":3},{"x":4,"w":1,"y":3},{"x":5,"w":1,"y":3},{"x":6,"w":1,"y":3},{"x":7,"w":1,"y":3},{"x":8,"w":1,"y":3},{"x":9,"w":1,"y":3},{"x":10,"w":1,"y":3},{"x":11,"w":1,"y":3}]},"LAYOUT_plaid_mit":{"key_count":47,"layout":[{"x":0,"w":1,"y":0},{"x":1,"w":1,"y":0},{"x":2,"w":1,"y":0},{"x":3,"w":1,"y":0},{"x":4,"w":1,"y":0},{"x":5,"w":1,"y":0},{"x":6,"w":1,"y":0},{"x":7,"w":1,"y":0},{"x":8,"w":1,"y":0},{"x":9,"w":1,"y":0},{"x":10,"w":1,"y":0},{"x":11,"w":1,"y":0},{"x":0,"w":1,"y":1},{"x":1,"w":1,"y":1},{"x":2,"w":1,"y":1},{"x":3,"w":1,"y":1},{"x":4,"w":1,"y":1},{"x":5,"w":1,"y":1},{"x":6,"w":1,"y":1},{"x":7,"w":1,"y":1},{"x":8,"w":1,"y":1},{"x":9,"w":1,"y":1},{"x":10,"w":1,"y":1},{"x":11,"w":1,"y":1},{"x":0,"w":1,"y":2},{"x":1,"w":1,"y":2},{"x":2,"w":1,"y":2},{"x":3,"w":1,"y":2},{"x":4,"w":1,"y":2},{"x":5,"w":1,"y":2},{"x":6,"w":1,"y":2},{"x":7,"w":1,"y":2},{"x":8,"w":1,"y":2},{"x":9,"w":1,"y":2},{"x":10,"w":1,"y":2},{"x":11,"w":1,"y":2},{"x":0,"w":1,"y":3},{"x":1,"w":1,"y":3},{"x":2,"w":1,"y":3},{"x":3,"w":1,"y":3},{"x":4,"w":1,"y":3},{"x":5,"w":2,"y":3},{"x":7,"w":1,"y":3},{"x":8,"w":1,"y":3},{"x":9,"w":1,"y":3},{"x":10,"w":1,"y":3},{"x":11,"w":1,"y":3}]},"LAYOUT_planck_mit":{"key_count":47,"layout":[{"x":0,"w":1,"y":0},{"x":1,"w":1,"y":0},{"x":2,"w":1,"y":0},{"x":3,"w":1,"y":0},{"x":4,"w":1,"y":0},{"x":5,"w":1,"y":0},{"x":6,"w":1,"y":0},{"x":7,"w":1,"y":0},{"x":8,"w":1,"y":0},{"x":9,"w":1,"y":0},{"x":10,"w":1,"y":0},{"x":11,"w":1,"y":0},{"x":0,"w":1,"y":1},{"x":1,"w":1,"y":1},{"x":2,"w":1,"y":1},{"x":3,"w":1,"y":1},{"x":4,"w":1,"y":1},{"x":5,"w":1,"y":1},{"x":6,"w":1,"y":1},{"x":7,"w":1,"y":1},{"x":8,"w":1,"y":1},{"x":9,"w":1,"y":1},{"x":10,"w":1,"y":1},{"x":11,"w":1,"y":1},{"x":0,"w":1,"y":2},{"x":1,"w":1,"y":2},{"x":2,"w":1,"y":2},{"x":3,"w":1,"y":2},{"x":4,"w":1,"y":2},{"x":5,"w":1,"y":2},{"x":6,"w":1,"y":2},{"x":7,"w":1,"y":2},{"x":8,"w":1,"y":2},{"x":9,"w":1,"y":2},{"x":10,"w":1,"y":2},{"x":11,"w":1,"y":2},{"x":0,"w":1,"y":3},{"x":1,"w":1,"y":3},{"x":2,"w":1,"y":3},{"x":3,"w":1,"y":3},{"x":4,"w":1,"y":3},{"x":5,"w":2,"y":3},{"x":7,"w":1,"y":3},{"x":8,"w":1,"y":3},{"x":9,"w":1,"y":3},{"x":10,"w":1,"y":3},{"x":11,"w":1,"y":3}]},"KEYMAP":{"key_count":48,"layout":[{"x":0,"w":1,"y":0},{"x":1,"w":1,"y":0},{"x":2,"w":1,"y":0},{"x":3,"w":1,"y":0},{"x":4,"w":1,"y":0},{"x":5,"w":1,"y":0},{"x":6,"w":1,"y":0},{"x":7,"w":1,"y":0},{"x":8,"w":1,"y":0},{"x":9,"w":1,"y":0},{"x":10,"w":1,"y":0},{"x":11,"w":1,"y":0},{"x":0,"w":1,"y":1},{"x":1,"w":1,"y":1},{"x":2,"w":1,"y":1},{"x":3,"w":1,"y":1},{"x":4,"w":1,"y":1},{"x":5,"w":1,"y":1},{"x":6,"w":1,"y":1},{"x":7,"w":1,"y":1},{"x":8,"w":1,"y":1},{"x":9,"w":1,"y":1},{"x":10,"w":1,"y":1},{"x":11,"w":1,"y":1},{"x":0,"w":1,"y":2},{"x":1,"w":1,"y":2},{"x":2,"w":1,"y":2},{"x":3,"w":1,"y":2},{"x":4,"w":1,"y":2},{"x":5,"w":1,"y":2},{"x":6,"w":1,"y":2},{"x":7,"w":1,"y":2},{"x":8,"w":1,"y":2},{"x":9,"w":1,"y":2},{"x":10,"w":1,"y":2},{"x":11,"w":1,"y":2},{"x":0,"w":1,"y":3},{"x":1,"w":1,"y":3},{"x":2,"w":1,"y":3},{"x":3,"w":1,"y":3},{"x":4,"w":1,"y":3},{"x":5,"w":1,"y":3},{"x":6,"w":1,"y":3},{"x":7,"w":1,"y":3},{"x":8,"w":1,"y":3},{"x":9,"w":1,"y":3},{"x":10,"w":1,"y":3},{"x":11,"w":1,"y":3}]}}}}}`)
	var plaidKeyboard, testingKeyboard Keyboard
	_ = json.Unmarshal(rawPlaidKeyboard, &plaidKeyboard)
	dclient := createNewTestDefClient()
	testingResponse, err := dclient.newRequest("GET", "keyboards/plaid", &testingKeyboard, nil)
	if err != nil {
		t.Errorf("got %s", err)
	}
	if !cmp.Equal(testingKeyboard, plaidKeyboard) {
		t.Errorf("expected %v, got %v", plaidKeyboard, testingKeyboard)
	}
	fmt.Print(testingResponse.Body)
}

func TestClient_do(t *testing.T) {
	const expectedText = "text"
	const expectedFavoriteCount int64 = 24

	client, mux, server := testServer()
	defer server.Close()
	mux.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"text": "text", "favorite_count": 24}`)
	})

	testClient := createNewCustomClient(client)
	req, _ := http.NewRequest("GET", "http://example.com/success", nil)

	model := new(FakeModel)

	resp, err := testClient.do(req, model)

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected %d, got %d", 200, resp.StatusCode)
	}
	if model.Text != expectedText {
		t.Errorf("expected %s, got %s", expectedText, model.Text)
	}
	if model.FavoriteCount != expectedFavoriteCount {
		t.Errorf("expected %d, got %d", expectedFavoriteCount, model.FavoriteCount)
	}
}

func TestClient_do_noContent(t *testing.T) {
	client, mux, server := testServer()
	defer server.Close()
	mux.HandleFunc("/nocontent", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	})

	testClient := createNewCustomClient(client)
	req, _ := http.NewRequest("DELETE", "http://example.com/nocontent", nil)

	model := new(FakeModel)

	resp, err := testClient.do(req, model)

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	if resp.StatusCode != 204 {
		t.Errorf("expected %d, got %d", 204, resp.StatusCode)
	}
	expectedModel := &FakeModel{}
	if !reflect.DeepEqual(expectedModel, model) {
		t.Errorf("successV should not be populated, exepcted %v, got %v", expectedModel, model)
	}
}
