package main

import (
  // "time"
  // "encoding/json"
)

type Todo struct {
  Id int `json:"id"`
  Name string `json:"name"`
}

type Todos []Todo

type Holidays []Holiday

type Holiday struct {
  Date string `json:"date"`
  Holiday string `json:"holiday"`
  Day string `json:"day"`
}
