package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
)

var currentId int
var todos Todos

var holidays Holidays

func init() {
  file, err := ioutil.ReadFile("holidays.json")
  if err != nil {
    log.Fatal(err)
  }

  err = json.Unmarshal(file, &holidays)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(holidays)
}

func RepoFindTodo(id int) Todo {
  for _, t := range todos {
    if t.Id == id {
      return t
    }
  }

  return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
  currentId += 1
  t.Id = currentId
  todos = append(todos, t)
  return t
}
