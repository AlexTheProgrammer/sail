package dom

type Header struct {
	Element
}

func (h *Header) OpenTag() string  { return "<header>" }
func (h *Header) CloseTag() string { return "</header>" }
func (h *Header) IsNil() bool {
	return nil == h
}
