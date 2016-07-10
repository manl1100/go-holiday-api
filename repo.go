package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
)

var currentId int
var todos Todos

func init() {
  var data Holidays
  file, err := ioutil.ReadFile("holidays.json")
  if err != nil {
    log.Fatal(err)
  }

  err = json.Unmarshal(file, &data)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(data)

  RepoCreateTodo(Todo { Name: "something" })
  RepoCreateTodo(Todo { Name: "something else"})
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
