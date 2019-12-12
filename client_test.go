package qmk

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// Json/XML-tagged model struct
type FakeModel struct {
	Text          string  `json:"text,omitempty"`
	FavoriteCount int64   `json:"favorite_count,omitempty"`
	Temperature   float64 `json:"temperature,omitempty"`
}

func createNewTestDefClient() *Client {
	baseURL, _ := url.Parse("https://api.qmk.fm")
	httpClient := &http.Client{}
	defClient := &Client{
		httpClient,
		baseURL,
	}
	return defClient
}

func createNewCustomClient(customClient *http.Client) *Client {
	baseURL, _ := url.Parse(`http://www.example.com`)
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

// createEndpoint createsthbetest server with the specified endpoint and expected response
func createEndpoint(endpoint, response string) (*http.Client, *httptest.Server) {
	client, mux, server := testServer()
	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, response)
	})
	return client, server
}

func TestNew(t *testing.T) {
	wantDefClient := createNewTestDefClient()
	gotDefClient := New(nil)
	if !cmp.Equal(wantDefClient.Client, gotDefClient.Client) {
		t.Errorf("expected default %v to equal default %v", wantDefClient, gotDefClient)
	}

	wantHttpClient := &http.Client{
		Timeout: time.Second * 4,
	}

	gotCustomClient := New(wantHttpClient)

	if gotCustomClient.Timeout != wantHttpClient.Timeout {
		t.Errorf("expected %d, got %d", wantHttpClient.Timeout, gotCustomClient.Timeout)
	}
}

func TestClient_newRequest(t *testing.T) {
	client, server := createEndpoint("/goodEndpoint", `{"text": "text", "favorite_count": 22}`)
	defer server.Close()

	testClient := createNewCustomClient(client)

	testModel := &FakeModel{}

	_, err := testClient.newRequest("GET", "/goodEndpoint", &testModel, nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}
	if testModel.FavoriteCount != 22 {
		t.Errorf("got %d, expected true", testModel.FavoriteCount)
	}
	_, err = testClient.newRequest("GET", "/badEndpoint", &testModel, nil)
	if err == nil {
		t.Error("expected error")

	}

}

func TestClient_do(t *testing.T) {
	const expectedText = "text"
	const expectedFavoriteCount int64 = 24

	client, server := createEndpoint("/success", `{"text": "text", "favorite_count": 24}`)
	defer server.Close()

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
