package dom

import (
	"syscall/js"
)

// DOM represents a document object model that will be rendered in the browser
type DOM struct {
	Header any
	Body   any
}

// hard coded for now will change shortly
var htmlString = "<h2>here is my new site</h2>"

func getHtml() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		return htmlString
	})
}

// Render will inject the defined DOM into the browser
func (d *DOM) Render() {
	js.Global().Set("getHtml", getHtml())
}
