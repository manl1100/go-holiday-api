package main

import (
  "fmt"
  "html"
  "net/http"
  "encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex (w http.ResponseWriter, r *http.Request) {
  todos := Todos {
    Todo { Name: "Something 1" },
    Todo { Name: "Somethingelse" },
  }

  json.NewEncoder(w).Encode(todos)
}
