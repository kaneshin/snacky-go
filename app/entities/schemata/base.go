package schemata

import "time"

type Base struct {
	CreatedAt int64 `db:"created_at" json:"created_at"`
	UpdatedAt int64 `db:"updated_at" json:"updated_at"`
}

func (self *Base) Init() {
	self.CreatedAt = time.Now().Unix()
	self.UpdatedAt = time.Now().Unix()
}
