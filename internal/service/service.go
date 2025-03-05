package service

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

var MonthToInt = map[string]int32{
	"JAN": 1,
	"FEB": 2,
	"MAR": 3,
	"APR": 4,
	"MAY": 5,
	"JUN": 6,
	"JUL": 7,
	"AUG": 8,
	"SEP": 9,
	"OCT": 10,
	"NOV": 11,
	"DEC": 12,
}

type Service struct {
}

func NewService(db *pgxpool.Pool) Service {
	return Service{}
}
