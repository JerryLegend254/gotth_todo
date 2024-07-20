package types

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

type TodoStore interface {
	GetTodos() ([]Todo, error)
	DeleteTodos(id int) error
}
