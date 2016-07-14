package main

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "strings"
  "strconv"
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
}

func RepoFindTodo(id int) Holidays {
  var output Holidays
  for _, t := range holidays {
    dt := strings.Split(t.Date, "/")
    var mon int
    var err error
    if mon, err = strconv.Atoi(dt[0]); err != nil {
      panic(err)
    }
    if mon == id {
      output = append(output, t)
    }
  }

  return output
}

func RepoCreateTodo(t Todo) Todo {
  currentId += 1
  t.Id = currentId
  todos = append(todos, t)
  return t
}
