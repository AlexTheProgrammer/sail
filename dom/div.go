package dom

type Div struct {
	Element
}

func NewDiv(innerHTML string, nodes ...Node) *Div {
	return &Div{
		Element: Element{
			innerHTML: innerHTML,
			Ns:        nodes,
		},
	}
}

func (d *Div) OpenTag() string  { return "<div>" }
func (d *Div) CloseTag() string { return "</div>" }
func (d *Div) IsNil() bool {
	return nil == d
}
