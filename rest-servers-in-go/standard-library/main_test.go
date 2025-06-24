package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	mux := setupServer()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping/", nil)
	mux.ServeHTTP(w, req)

	if want, got := 200, w.Code; want != got {
		t.Errorf("got %d want %d", got, want)
	}
	if want, got := "pong", w.Body.String(); want != got {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestCreateTaskHandler(t *testing.T) {
	type InvalidRequestTask struct {
		abc string `json:abc"`
	}

	type RequestTask struct {
		Text string    `json:"text"`
		Tags []string  `json:"tags"`
		Due  time.Time `json:"due"`
	}

	type ResponseId struct {
		Id int `json:"id"`
	}

	rt := RequestTask{
		Text: "first task",
		Tags: []string{"tag1"},
		Due:  time.Now(),
	}

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(rt)
	if err != nil {
		t.Fatalf("%v", err)
	}

	invalidRt := InvalidRequestTask{
		abc: "lasdkjf",
	}

	var invalidBody bytes.Buffer
	err = json.NewEncoder(&invalidBody).Encode(invalidRt)
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Run("Creates task given proper content type and proper data fields", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/task/", &body)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		req.Header.Add("Content-Type", "application/json")

		mux := setupServer()
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()

		var resBody ResponseId
		err = json.NewDecoder(res.Body).Decode(&resBody)
		if err != nil {
			t.Fatalf("decoding response body failed: %v", err)
		}

		if want, got := resBody.Id, 0; want != got {
			t.Fatalf("expected %d, got %d", want, got)
		}
	})

	t.Run("Returns 400 bad request when given bad body", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/task/", &body)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		req.Header.Add("Content-Type", "application/json")

		mux := setupServer()
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()

		if want, got := 400, res.StatusCode; want != got {
			t.Fatalf("expected %d, got %d", want, got)
		}
	})

	t.Run("Returns 400 when content-type is not given", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/task/", &body)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		mux := setupServer()
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()

		if want, got := 400, res.StatusCode; want != got {
			t.Fatalf("expected %d, got %d", want, got)
		}
	})

	t.Run("Returns 415 when invalid content-type is given", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/task/", &body)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		req.Header.Add("Content-Type", "invalid")

		mux := setupServer()
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()

		if want, got := 415, res.StatusCode; want != got {
			t.Fatalf("expected %d, got %d", want, got)
		}
	})
}
