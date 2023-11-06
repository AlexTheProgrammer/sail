package dom

// Support Vector Graphic
type SVGEl struct {
	*Element
}

func SVG(hp Props, nodes ...Node) *SVGEl {
	return &SVGEl{
		Element: NewElement(hp, nodes...),
	}
}

func (b *SVGEl) Tag() string { return "svg" }
func (b *SVGEl) IsNil() bool {
	return nil == b
}
