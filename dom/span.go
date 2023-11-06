package dom

type SpanEl struct {
	*Element
}

func Span(hp Props, nodes ...Node) *SpanEl {
	return &SpanEl{
		Element: NewElement(hp, nodes...),
	}
}

func (b *SpanEl) Tag() string { return "span" }
func (b *SpanEl) IsNil() bool {
	return nil == b
}
