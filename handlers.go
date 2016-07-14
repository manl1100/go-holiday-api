package main

import (
  "net/http"
  "encoding/json"
  "strconv"

  "github.com/gorilla/mux"
)

func HolidayIndex(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(holidays); err != nil {
    panic(err)
  }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  var todoId int
  var err error
  if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
    panic(err)
  }
  todo := RepoFindTodo(todoId)
  if len(todo) > 0 {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todo); err != nil {
      panic(err)
    }
    return
  }

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusNotFound)
  if err := json.NewEncoder(w).Encode(JsonError { Code: http.StatusNotFound, Text: "Not found"}); err != nil {
    panic(err)
  }
}
