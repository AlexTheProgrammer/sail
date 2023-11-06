package dom

type Header struct {
	Element
}

func (h *Header) Tag() string { return "head" }
func (h *Header) IsNil() bool {
	return nil == h
}

// TODO: implement
func (h *Header) HTMLProps() Props {
	return Props{}
}
