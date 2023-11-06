package dom

type Props struct {
	InnerHTML  string
	Class      *Class
	ID         string
	HRef       string
	Rel        string
	Content    string
	CharSet    string
	AriaLabel  string
	D          string
	AriaHidden string
	XMLNS      string
	Fill       string
	ViewBox    string
}

type Class struct {
	raw string
}

func NewClass(className ...string) *Class {

	c := &Class{}
	for i, v := range className {
		if i == 0 {
			c.raw = v
			continue
		}
		c.raw += " " + v
	}

	return c

}

// Add adds the className to the target class
func (c *Class) Add(className string) *Class {

	if len(c.raw) == 0 {
		c.raw = className
		return c
	}

	c.raw += " " + className

	return c
}

func (c *Class) Nil() bool {
	return c == nil || len(c.raw) == 0
}

func (c *Class) Parse() string {
	return c.raw
}
