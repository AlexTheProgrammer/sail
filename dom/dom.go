// package DOM
//
// provides an abstraction to the Document Object Model to enable
// programmatic definitions of the document object model using golang
package dom

import (
	"log"
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
	log.Printf("this should be fine t: %v %T", t, t)
	if t == nil {
		return "<div></div>"
	}

	return t.OpenTag() + t.CloseTag()
}

// Render traverses a node tree and generates a string of that node
// including all renderings of child nodes
func Render(node Node) string {
	log.Printf("processing node: %v is it nil %v", node, node == nil)
	if node.IsNil() {
		log.Printf("empty node early exit: %v", node)
		return emptyRender(node)
	}

	if h, ok := node.(*Header); ok {
		if h == nil {
			log.Printf("I can see it's nil when I convert its type")
		} else {
			log.Printf("h is %#v", h)
		}
	}
	log.Printf("node is not nil it is: %v %T", node, node)

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
	log.Printf("start render process")
	htmlString := ""

	// TODO: currently we are updating only the body through the assets
	// fix this so that we can also set the headers
	log.Printf("processing header: %v", d.Header)
	log.Printf("is header nil: %v", d.Header == nil)
	htmlString += Render(d.Header)

	log.Printf("processing body: %v", d.Body)
	htmlString += Render(d.Body)

	log.Printf("about to render htmlString %v", htmlString)
	js.Global().Set("getHtml", getHtml(htmlString))
}

// getHtml returns  syscall/js Func that can bind wasm to js API
func getHtml(htmlString string) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		return htmlString
	})
}
