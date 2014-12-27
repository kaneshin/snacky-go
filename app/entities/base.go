package entities

type Base struct {
}

func NewBase() *Base {
	base := Base{}
	return &base
}

func (b *Base) InitBase() {
}
