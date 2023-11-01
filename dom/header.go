package dom

type Header struct {
	Element
}

func (h *Header) OpenTag() string  { return "<head>" }
func (h *Header) CloseTag() string { return "</head>" }
func (h *Header) IsNil() bool {
	return nil == h
}
