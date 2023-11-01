package dom

type Element struct {
	Ns        []Node
	innerHTML string
}

func (e *Element) IsNil() bool {
	return nil == e
}

func (e *Element) InnerHTML() string {
	if e == nil {
		return ""
	}

	return e.innerHTML
}

func (e *Element) Nodes() []Node {
	if e == nil {
		return nil
	}

	return e.Ns
}
