package dom

type Body struct {
	Element
}

func NewBody(nodes ...Node) *Body {
	return &Body{
		Element{
			Ns: nodes,
		},
	}
}

func (b *Body) OpenTag() string  { return "<body>" }
func (b *Body) CloseTag() string { return "</body>" }
func (b *Body) IsNil() bool {
	return nil == b
}
