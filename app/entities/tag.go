package entities

import (
	"time"

	schm "github.com/kaneshin/snacky-go/app/entities/schemata"
)

type Tag struct {
	*Base
	schm.Tag
}

func NewTag(id int32, name string) *Tag {
	tag := Tag{
		Base: NewBase(),
	}
	tag.InitBase()
	tag.Id = id
	tag.Name = name
	tag.CreatedAt = time.Now().Unix()
	tag.UpdatedAt = time.Now().Unix()
	return &tag
}
