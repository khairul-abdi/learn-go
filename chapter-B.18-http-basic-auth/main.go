package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	log.Println("TYPE RESULT ==> ", reflect.TypeOf(res))
	log.Println("RESULT ==> ", string(res))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {
		return
	}
	if !AllowOnlyGET(w, r) {
		return
	}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":9000"

	fmt.Println("server started at localhost:9000")
	server.ListenAndServe()
}
