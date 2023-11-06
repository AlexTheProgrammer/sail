package dom

// Unordered List
type ULEl struct {
	*Element
}

func UL(hp Props, nodes ...Node) *ULEl {
	return &ULEl{
		Element: NewElement(hp, nodes...),
	}
}

func (u *ULEl) Tag() string { return "ul" }
func (u *ULEl) IsNil() bool {
	return nil == u
}
