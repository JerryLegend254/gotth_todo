package types

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

type TodoStore interface {
	GetTodos() ([]Todo, error)
	DeleteTodos(id int) error
	GetStoreBySearchQuery(query string) ([]Todo, error)
	GetTodoByID(id int) (Todo, error)
	EditTodo(todo Todo) (Todo, error)
}
