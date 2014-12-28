package schemata

import "time"

type SnackTag struct {
	Id        int32     `db:"id"`
	SnackId   int32     `db:"snack_id"`
	TagId     int32     `db:"tag_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
