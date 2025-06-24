package main

import (
	"encoding/json"
	"log"
	"mime"
	"net/http"
	"os"
	"time"

	"stdlib-basic/internal/taskstore"
)

type taskServer struct {
	store *taskstore.TaskStore
}

func NewTaskServer() *taskServer {
	store := taskstore.New()
	return &taskServer{store}
}

func (ts *taskServer) testRoute(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling ping at %s\n", r.URL.Path)

	w.Write([]byte("pong"))
}

func (ts *taskServer) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling task create at %s\n", r.URL.Path)

	type RequestTask struct {
		Text string    `json:"text"`
		Tags []string  `json:"tags"`
		Due  time.Time `json:"due"`
	}

	type ResponseId struct {
		Id int `json:"id"`
	}

	ct := r.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(ct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if mediatype != "application/json" {
		msg := "Content-Type must be application/json"
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var rt RequestTask
	if err := dec.Decode(&rt); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := ts.store.CreateTask(rt.Text, rt.Tags, rt.Due)
	res, err := json.Marshal(ResponseId{Id: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(res)
}

func (ts *taskServer) getAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling getting all tasks at %s\n", r.URL.Path)

	// verify stuff
	// get all tasks from store
	allTasks := ts.store.GetAllTasks()
	res, err := json.Marshal(allTasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// send to user
	w.Header().Add("Content-Type", "application/json")
	w.Write(res)
}

func (ts *taskServer) deleteAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling deleting all tasks at %s\n", r.URL.Path)
	ts.store.DeleteAllTasks()
	w.WriteHeader(http.StatusNoContent)
}

func setupServer() *http.ServeMux {
	mux := http.NewServeMux()
	server := NewTaskServer()

	mux.HandleFunc("GET /ping/", server.testRoute)
	mux.HandleFunc("POST /task/", server.createTaskHandler)
	mux.HandleFunc("GET /task/", server.getAllTasksHandler)
	mux.HandleFunc("DELETE /task/", server.deleteAllTasksHandler)
	// mux.HandleFunc("GET /task/{id}/", server.getTaskHandler)
	// mux.HandleFunc("DELETE /task/{id}/", server.deleteTaskHandler)
	// mux.HandleFunc("GET /tag/{tag}/", server.tagHandler)
	// mux.HandleFunc("GET /due/{year}/{month}/{day}/", server.dueHandler)

	return mux
}

func main() {
	mux := setupServer()
	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("SERVERPORT"), mux))
}
