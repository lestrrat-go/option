package option

type Interface interface {
	Ident() interface{}
	Value() interface{}
}

type pair struct {
	ident interface{}
	value interface{}
}

func New(ident, value interface{}) Interface {
	return &pair{
		ident: ident,
		value: value,
	}
}

func (p *pair) Ident() interface{} {
	return p.ident
}

func (p *pair) Value() interface{} {
	return p.value
}
