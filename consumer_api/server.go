package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

type Thing struct {
	Name   string  `json:"name"`
	Value  int     `json:"value"`
	Weight float32 `json:"weight"`
}

type thingHandlers struct {
	sync.Mutex
	store map[string]Thing
}

func newThingHandlers() *thingHandlers {
	return &thingHandlers{
		store: map[string]Thing{},
	}
}
func (h *thingHandlers) things(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
	}
}
func (h *thingHandlers) get(w http.ResponseWriter, r *http.Request) {
	things := make([]Thing, len(h.store))

	h.Lock()
	i := 0
	for _, thing := range h.store {
		things[i] = thing
		i++
	}
	h.Unlock()

	jsonBytes, err := json.Marshal(things)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *thingHandlers) post(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(([]byte(err.Error())))
		return
	}
	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write(([]byte(fmt.Sprintf("need application-type application/json, but got '%s'", ct))))
		return
	}
	var thing Thing
	err2 := json.Unmarshal(bodyBytes, &thing)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	h.Lock()
	h.store[thing.Name] = thing
	defer h.Unlock()
}
func (h *thingHandlers) runAlgo(w http.ResponseWriter, r *http.Request) {
	host := os.Getenv("HOST")
	if host == "" {
		host = "http://localhost:8083"
	}
	resp, err := http.Get(fmt.Sprintf("%s/run_algo", host))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func main() {
	thingHandlers := newThingHandlers()
	http.HandleFunc("/things", thingHandlers.things)
	http.HandleFunc("/things/algo", thingHandlers.runAlgo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
