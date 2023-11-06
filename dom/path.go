package dom

// Support Vector Graphic
type PathEl struct {
	*Element
}

func Path(hp Props, nodes ...Node) *PathEl {
	return &PathEl{
		Element: NewElement(hp, nodes...),
	}
}

func (p *PathEl) Tag() string { return "path" }
func (p *PathEl) IsNil() bool {
	return nil == p
}
