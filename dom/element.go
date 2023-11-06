package dom

type Element struct {
	Ns        []Node
	htmlProps Props
}

func NewElement(hp Props, nodes ...Node) *Element {
	return &Element{
		htmlProps: hp,
		Ns:        nodes,
	}
}

func (e *Element) HTMLProps() Props {
	return e.htmlProps
}

func (e *Element) El(n Node) {
	e.Ns = append(e.Ns, n)
}

func (e *Element) IsNil() bool {
	return nil == e
}

func (e *Element) Props() string {
	if e == nil {
		return ""
	}

	return e.htmlProps.InnerHTML
}

func (e *Element) Nodes() []Node {
	if e == nil {
		return nil
	}

	return e.Ns
}
