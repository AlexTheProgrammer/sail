package dom

type BodyEl struct {
	Element
}

func Body(nodes ...Node) *BodyEl {
	return &BodyEl{
		Element{
			Ns: nodes,
		},
	}
}

func (b *BodyEl) Tag() string { return "body" }
func (b *BodyEl) IsNil() bool {
	return nil == b
}

// TODO: implement
func (b *BodyEl) HTMLProps() Props {
	return Props{}
}
