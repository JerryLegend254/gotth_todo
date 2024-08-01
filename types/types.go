package types

type Todo struct {
	ID        int    `form:"id"`
	Title     string `form:"title" validate:"required"`
	Completed bool   `form:"completed"`
}

type TodoStore interface {
	GetTodos() ([]Todo, error)
	DeleteTodos(id int) error
	GetStoreBySearchQuery(query string) ([]Todo, error)
	GetTodoByID(id int) (Todo, error)
	EditTodo(todo Todo) (Todo, error)
	CreateNewTodo(todo Todo) (Todo, error)
}
