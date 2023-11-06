package dom

// a href
type AEl struct {
	*Element
}

func A(hp Props, nodes ...Node) *AEl {
	return &AEl{
		Element: NewElement(hp, nodes...),
	}
}

func (u *AEl) Tag() string { return "a" }
func (u *AEl) IsNil() bool {
	return nil == u
}
