package entities

import (
	"time"

	schm "github.com/kaneshin/snacky-go/app/entities/schemata"
)

type SnackTag struct {
	*Base
	schm.SnackTag
}

func NewSnackTag(snackId int32, tagId int32) *SnackTag {
	st := SnackTag{
		Base: NewBase(),
	}
	st.InitBase()
	st.SnackId = snackId
	st.TagId = tagId
	st.CreatedAt = time.Now().Unix()
	st.UpdatedAt = time.Now().Unix()
	return &st
}
