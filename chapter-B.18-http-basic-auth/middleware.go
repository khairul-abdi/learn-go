package main

import (
	"log"
	"net/http"
)

const USERNAME = "batman"
const PASSWORD = "secret"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	log.Println("r.BasicAuth ==> ", ok)
	if !ok {
		w.Write([]byte(`something went wrong`))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		if username != USERNAME && password != PASSWORD {
			w.Write([]byte(`wrong username and password`))
			return false
		} else if username != USERNAME {
			w.Write([]byte(`wrong username`))
			return false
		} else if password != PASSWORD {
			w.Write([]byte(`wrong password`))
			return false
		}
	}

	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" && r.Method != "get" {
		w.Write([]byte("Only GET is allowed"))
		return false
	}

	return true
}
