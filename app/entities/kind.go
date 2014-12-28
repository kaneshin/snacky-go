package entities

import (
	"time"

	schm "github.com/kaneshin/snacky-go/app/entities/schemata"
)

type Kind struct {
	*Base
	schm.Kind
}

func NewKind(id int32, name string) *Kind {
	kind := Kind{
		Base: NewBase(),
	}
	kind.InitBase()
	kind.Id = id
	kind.Name = name
	kind.CreatedAt = time.Now()
	kind.UpdatedAt = time.Now()
	return &kind
}

func (e *Kind) Scheme() *schm.Kind {
	return &e.Kind
}
