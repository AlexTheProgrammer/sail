// package DOM
//
// provides an abstraction to the Document Object Model to enable
// programmatic definitions of the document object model using golang
package dom

import (
	"syscall/js"
)

// DOM represents a document object model that will be rendered in the browser
type DOM struct {
	Header []Node
	Body   []Node
}

type Node interface {
	Render() string
}

type Div struct {
	innerHTML string
}

func (d *Div) Render() string {
	if d == nil {
		return "<div></div>"
	}

	return "<div>" + d.innerHTML + "</div>"
}

func getHtml(htmlString string) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		return htmlString
	})
}

// Render will inject the defined DOM into the browser
func (d *DOM) Render() {
	htmlString := ""

	// TODO: currently we are updating the body through the assets
	// fix this so that we can also set the headers
	for _, v := range d.Header {
		htmlString += v.Render()
	}

	for _, v := range d.Body {
		htmlString += v.Render()
	}

	js.Global().Set("getHtml", getHtml(htmlString))
}
