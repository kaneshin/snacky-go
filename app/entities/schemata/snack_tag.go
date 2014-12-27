package schemata

type SnackTag struct {
	Id        int32 `db:"id"`
	SnackId   int32 `db:"snack_id"`
	TagId     int32 `db:"tag_id"`
	CreatedAt int64 `db:"created_at"`
	UpdatedAt int64 `db:"updated_at"`
}
