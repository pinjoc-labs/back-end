// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlc

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Maturity struct {
	ID        int32
	Month     int32
	MonthName pgtype.Text
	Year      int32
}

type Order struct {
	ID                int32
	CollateralTokenID int32
	DebtTokenID       int32
	MaturityID        int32
	Rate              pgtype.Numeric
	AvailableToken    int32
	OrderType         string
}

type Token struct {
	ID      int32
	Name    string
	Symbol  string
	Address string
	Icon    pgtype.Text
}

type Tokenized struct {
	ID           int32
	QuoteTokenID int32
	BaseTokenID  int32
	Price        pgtype.Numeric
	MaturityID   int32
	Rate         pgtype.Numeric
	Amount       int32
	Volume       int32
	OrderType    string
}
