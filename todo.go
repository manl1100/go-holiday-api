package main

import (
)

type Todo struct {
  Id int `json:"id"`
  Name string `json:"name"`
}

type Todos []Todo
