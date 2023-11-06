package dom

// Unordered List
type LIEl struct {
	*Element
}

func LI(hp Props, nodes ...Node) *LIEl {
	return &LIEl{
		Element: NewElement(hp, nodes...),
	}
}

func (u *LIEl) Tag() string { return "li" }
func (u *LIEl) IsNil() bool {
	return nil == u
}
