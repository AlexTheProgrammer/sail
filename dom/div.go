package dom

type DivEl struct {
	*Element
}

func Div(hp Props, nodes ...Node) *DivEl {
	return &DivEl{
		Element: NewElement(hp, nodes...),
	}
}

func (d *DivEl) Tag() string { return "div" }
func (d *DivEl) IsNil() bool {
	return nil == d
}
