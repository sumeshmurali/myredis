package main

import (
	"encoding/json"
	"log"
	"myredis/internal/datastore"
	"net/http"
	"strings"
)

type Command struct {
	Op   string   `json:"command"`
	Key  string   `json:"key"`
	Args []string `json:"args"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodPost:
		{
			ct := r.Header.Get("Content-Type")
			if ct == "" {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}

			mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))
			if mediaType != "application/json" {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}

			r.Body = http.MaxBytesReader(w, r.Body, 1024*1024*1024)

			dec := json.NewDecoder(r.Body)
			dec.DisallowUnknownFields()

			var c Command
			err := dec.Decode(&c)
			if err != nil {
				log.Printf("here %v", err.Error())
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			ds := datastore.GetDataStore()
			res, err := ds.Command(c.Op, c.Key, c.Args...)
			if err != nil {
				log.Printf("Error %v occured while attempting process command %q on key %q with args %v \n", err.Error(), c.Op, c.Key, c.Args)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			enc := json.NewEncoder(w)
			w.Header().Set("Content-Type", "application/json")
			err = enc.Encode(res)
			if err != nil {
				log.Printf("Error %v occured while converting the result %+v to JSON", err.Error(), res)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

		}
	default:
		{
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}

func GetRouter() *http.ServeMux {
	mux := http.ServeMux{}

	mux.HandleFunc("/", handler)
	return &mux
}
