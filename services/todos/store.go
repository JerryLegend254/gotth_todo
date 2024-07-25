package todos

import (
	"database/sql"

	"github.com/JerryLegend254/gotth/types"
)

type Store struct {
	repo *sql.DB
}

func NewStore(repo *sql.DB) *Store {
	return &Store{repo: repo}
}

func (s *Store) GetTodos() ([]types.Todo, error) {
	rows, err := s.repo.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []types.Todo{}
	for rows.Next() {
		var todo types.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (s *Store) DeleteTodos(id int) error {
	_, err := s.repo.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetStoreBySearchQuery(query string) ([]types.Todo, error) {
	rows, err := s.repo.Query("SELECT * FROM todos WHERE title LIKE ?", "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []types.Todo{}
	for rows.Next() {
		var todo types.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (s *Store) GetTodoByID(id int) (types.Todo, error) {
	var todo types.Todo
	err := s.repo.QueryRow("SELECT * FROM todos WHERE id = ?", id).Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		return types.Todo{}, err
	}
	return todo, nil
}

func (s *Store) EditTodo(todo types.Todo) (types.Todo, error) {
	_, err := s.repo.Exec("UPDATE todos SET title = ?, completed = ? WHERE id = ?", todo.Title, todo.Completed, todo.ID)
	if err != nil {
		return types.Todo{}, err
	}
	editTodo, err := s.GetTodoByID(todo.ID)
	if err != nil {
		return types.Todo{}, err
	}
	return editTodo, nil
}

func (s *Store) CreateNewTodo(todo types.Todo) (types.Todo, error) {
	res, err := s.repo.Exec("INSERT INTO todos (title, completed) VALUES (?, ?)", todo.Title, todo.Completed)
	if err != nil {
		return types.Todo{}, err
	}

	id, _ := res.LastInsertId()

	constructedTodo := types.Todo{
		ID:        int(id),
		Title:     todo.Title,
		Completed: todo.Completed,
	}
	return constructedTodo, nil
}
