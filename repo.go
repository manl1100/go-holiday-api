package main

var currentId int
var todos Todos

func init() {
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
