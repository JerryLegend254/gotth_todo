package handlers

import "github.com/JerryLegend254/gotth/types"

type Handler struct {
	store types.TodoStore
}

func New(s types.TodoStore) *Handler {
	return &Handler{store: s}
}
