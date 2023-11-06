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
	Body   *BodyEl
}

type Node interface {
	Tager          // Tager provides open and close tags for the node
	HTMLPropser    // HTMLPropser provides the html properties for the node
	Nodes() []Node // Nodes returns child nodes for the node
	IsNil() bool
}

type HTMLPropser interface {
	HTMLProps() Props
}

type Tager interface {
	Tag() string
}

func emptyRender(t Tager) string {
	return "<" + t.Tag() + "></" + t.Tag() + ">"
}

func addProp(label, prop string) string {
	if prop == "" {
		return ""
	}

	return label + "=\"" + prop + "\"" + " "
}

// Render traverses a node tree and generates a string of that node
// including all renderings of child nodes
func Render(node Node) string {
	if node.IsNil() {
		return emptyRender(node)
	}

	props := node.HTMLProps()
	nr := "<" + node.Tag() + " "

	nr += addProp("id", props.ID)
	nr += addProp("href", props.HRef)
	nr += addProp("rel", props.Rel)
	nr += addProp("content", props.Content)
	nr += addProp("charset", props.CharSet)
	nr += addProp("arialabel", props.AriaLabel)
	nr += addProp("d", props.D)
	nr += addProp("ariahidden", props.AriaHidden)
	nr += addProp("xmlns", props.XMLNS)
	nr += addProp("fill", props.Fill)
	nr += addProp("viewbox", props.ViewBox)

	if !props.Class.Nil() {
		nr += "class=\"" + props.Class.Parse() + "\"" + " "
	}
	nr += ">"

	// innerHTML is only the text up to the first child node
	nr += props.InnerHTML

	for _, n := range node.Nodes() {
		nr += Render(n)
	}

	nr += "</" + node.Tag() + ">"

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
