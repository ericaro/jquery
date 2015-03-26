package apigenerator

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// this command parses the api.jquery.com

type Entry struct {
	Type    string `xml:"type,attr"`
	RawName string `xml:"name,attr"`
	Return  string `xml:"return,attr"`
	Title   string `xml:"entry>title"`
	Desc    string `xml:"desc"`
}

//Receiver computes the expected receiver
//
// jquery has the following convention
// Deferred.foo for a foo method on the Deferred Object
//
// But it has the following exception:
//
// foo -> receiver JQuery
// JQuery.foo -> receiver none (static function)
func (e Entry) Receiver() string {

	parts := strings.Split(e.RawName, ".")
	switch {
	case len(parts) <= 1:
		return "JQuery"
	case len(parts) > 1:
		obj := strings.Join(parts[0:len(parts)-1], ".")
		if obj == "jQuery" {
			return ""
		}
		return Title(obj)
	default:
		return Title(parts[0])
	}
}
func (e Entry) ReturnVoid() bool {
	receiver, _ := TypeConv(e.Return)
	return receiver == ""
}
func (e Entry) Name() string {

	parts := strings.Split(e.RawName, ".")
	return parts[len(parts)-1]
}

type JType struct {
	Receiver   string
	Methods    map[string]*Entry
	Properties map[string]*Entry
}

func NewJType(receiver string) *JType {
	return &JType{
		Receiver:   receiver,
		Methods:    make(map[string]*Entry),
		Properties: make(map[string]*Entry),
	}
}

func Title(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[n:]
}

//TypeConv converts a jquery declared type into a go type
func TypeConv(s string) (string, error) {
	switch s {
	case "jQuery":
		return "JQuery", nil

	case "Boolean", "boolean":
		return "bool", nil

	case "Number":
		return "float64", nil

	case "Integer":
		return "int", nil

	case "String":
		return "string", nil

	case "undefined", "":
		return "", nil

	//unsupported objects
	case "Object", "jqXHR", "Function", "Promise", "Array", "XMLDocument", "Element":
		return "*js.Object", nil

	case "Event", "Callbacks", "Deferred": //supported objects
		return s, nil

	default:
		fmt.Printf("unknown type %q\n", s)
		return s, fmt.Errorf("unknown type %s", s)
	}
}

//JsConv return the string to add to a *js.Object "j" to convert it to the right type
func JsConv(s string) (string, error) {
	switch s {
	case "Boolean", "boolean":
		return "j.Bool()", nil

	case "Number":
		return "j.Float()", nil

	case "Integer":
		return "j.Int()", nil

	case "String":
		return "j.String()", nil

	case "undefined", "":
		return "", nil

	//unsupported
	case "Object", "jqXHR", "Function", "Promise", "Array", "XMLDocument", "Element":
		return "j", nil

	case "jQuery", "Event", "Callbacks", "Deferred": //supported objects
		return fmt.Sprintf("new%s(j)", Title(s)), nil

	default:
		fmt.Printf("unknown type %s\n", s)
		return s, fmt.Errorf("unknown type %s", s)
	}
}

//text/template below uses Entry as host and the following function map
var funcs = map[string]interface{}{
	"titled":   Title, //add the title function as a pipe
	"typeconv": TypeConv,
	"jsconv":   JsConv,
}

const EntryTmpl = `
{{define "JQueryStatic"}}
// {{.Name|titled}} {{.Desc}}
//
// see http://api.jquery.com/{{.RawName}}
//{{if .ReturnVoid}}
func {{.Name|titled}}(i ...interface{}) {	JQ.Call("{{.Name}}", i...) }
{{else}}
func {{.Name|titled}}(i ...interface{}) {{.Return| typeconv}}  {
	j:= JQ.Call("{{.Name}}", i...)
	return {{.Return |jsconv}}
}
{{end}}
{{end}}

{{define "JType"}}
type {{.Receiver}} struct {
	*js.Object
	{{range .Properties}}
	{{.Name|titled}} {{.Return | typeconv}} ` + "`js:\"{{.Name}}\"`" + `{{end}}
}

func new{{.Receiver}}(j *js.Object) {{.Receiver}} {return {{.Receiver}}{Object: j} }

	{{range .Methods}}
	// {{.Name|titled}} {{.Desc}}
	//
	// see http://api.jquery.com/{{.RawName}}
	//{{if .ReturnVoid}}
	func (x {{.Receiver}}) {{.Name|titled}}(i ...interface{}) {{.Return| typeconv}} {x.Call("{{.Name}}", i...)}
	{{else}}
	func (x {{.Receiver}}) {{.Name|titled}}(i ...interface{}) {{.Return| typeconv}} {
		j:=x.Call("{{.Name}}", i...)
		return {{.Return |jsconv}}
	}
	{{end}}
	{{end}}
{{end}}
`
