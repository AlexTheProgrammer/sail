package dom

// Support Vector Graphic
type P struct {
	*Element
}

func NewP(hp Props, nodes ...Node) *P {
	return &P{
		Element: NewElement(hp, nodes...),
	}
}

func (b *P) Tag() string { return "svg" }
func (b *P) IsNil() bool {
	return nil == b
}
