package schemata

type Maker struct {
	Id        int32  `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Kana      string `db:"kana" json:"kana"`
	MakerId   int32  `db:"maker_id" json:"maker_id"`
	Price     string `db:"price" json:"price"`
	KindId    int32  `db:"kind_id" json:"kind_id"`
	MakedAt   int64  `db:"maked_at" json:"maked_at"`
	AreaId    int32  `db:"area_id" json:"area_id"`
	ImageUrl  string `db:"image_url" json:"image_url"`
	Comment   string `db:"comment" json:"comment"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
}
