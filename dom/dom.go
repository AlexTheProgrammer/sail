// package DOM
//
// provides an abstraction to the Document Object Model to enable
// programmatic definitions of the document object model using golang
package dom

import (
	"syscall/js"
)

type DOM struct {
	Header *Header
	Body   *Body
}

type Node interface {
	Tager          // Tager provides open and close tags for the node
	InnerHTMLer    // InnerHTMLer provides the innerHTML for the node
	Nodes() []Node // Nodes returns child nodes for the node
	IsNil() bool
}

type InnerHTMLer interface {
	InnerHTML() string
}

type Tager interface {
	CloseTag() string
	OpenTag() string
}

func emptyRender(t Tager) string {
	return t.OpenTag() + t.CloseTag()
}

// Render traverses a node tree and generates a string of that node
// including all renderings of child nodes
func Render(node Node) string {
	if node.IsNil() {
		return emptyRender(node)
	}

	nr := node.OpenTag()

	// innerHTML is only the text up to the first child node
	nr += node.InnerHTML()

	for _, n := range node.Nodes() {
		nr += Render(n)
	}

	nr += node.CloseTag()

	return nr
}

// Render injects the defined DOM into the browser
// this relies on the target html document calling
// body.innerHTML = getHTML()
// this is the default behaviour for raft.
func (d *DOM) Render() {
	htmlString := ""

	// TODO: currently we are updating only the body through the assets
	// fix this so that we can also set the headers
	htmlString += Render(d.Header)
	htmlString += Render(d.Body)

	js.Global().Set("getHtml", getHtml(htmlString))
}

// getHtml returns  syscall/js Func that can bind wasm to js API
func getHtml(htmlString string) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		return htmlString
	})
}
