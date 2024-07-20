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
