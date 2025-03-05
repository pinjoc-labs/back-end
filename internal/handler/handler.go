package handler

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
}

func NewHandler(db *pgxpool.Pool) Handler {
	return Handler{}
}
