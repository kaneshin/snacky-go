package schemata

import "time"

type Snack struct {
	Id        int32     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Kana      string    `db:"kana" json:"kana"`
	MakerId   int32     `db:"maker_id" json:"maker_id"`
	Price     string    `db:"price" json:"price"`
	KindId    int32     `db:"kind_id" json:"kind_id"`
	MakedAt   time.Time `db:"maked_at" json:"maked_at"`
	AreaId    int32     `db:"area_id" json:"area_id"`
	ImageUrl  string    `db:"image_url" json:"image_url"`
	Comment   string    `db:"comment" json:"comment"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
