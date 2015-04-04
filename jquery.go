package jquery

import "github.com/gopherjs/gopherjs/js"

// to get jquery-gen just
// go get github.com/ericaro/apigen/jquery-gen
// If you haven't already checkout api.jquery.com xml files.
// there is a .gitmodules file
//go:generate git submodule update
//go:generate jquery-gen -i ./api.jquery.com/entries/ -o api.go

const (
	//keys
	BLUR     = "blur"
	CHANGE   = "change"
	CLICK    = "click"
	DBLCLICK = "dblclick"
	FOCUS    = "focus"
	FOCUSIN  = "focusin"
	FOCUSOUT = "focusout"
	HOVER    = "hover"
	KEYDOWN  = "keydown"
	KEYPRESS = "keypress"
	KEYUP    = "keyup"
	//form
	SUBMIT = "submit"
	LOAD   = "load"
	UNLOAD = "unload"
	RESIZE = "resize"
	//mouse
	MOUSEDOWN  = "mousedown"
	MOUSEENTER = "mouseenter"
	MOUSELEAVE = "mouseleave"
	MOUSEMOVE  = "mousemove"
	MOUSEOUT   = "mouseout"
	MOUSEOVER  = "mouseover"
	MOUSEUP    = "mouseup"
	//touch
	TOUCHSTART  = "touchstart"
	TOUCHMOVE   = "touchmove"
	TOUCHEND    = "touchend"
	TOUCHENTER  = "touchenter"
	TOUCHLEAVE  = "touchleave"
	TOUCHCANCEL = "touchcancel"
	//ajax Events
	AJAXSTART    = "ajaxStart"
	BEFORESEND   = "beforeSend"
	AJAXSEND     = "ajaxSend"
	SUCCESS      = "success"
	AJAXSUCESS   = "ajaxSuccess"
	ERROR        = "error"
	AJAXERROR    = "ajaxError"
	COMPLETE     = "complete"
	AJAXCOMPLETE = "ajaxComplete"
	AJAXSTOP     = "ajaxStop"
)

var (
	JQ             = js.Global.Get("jQuery")
	BrowserVersion = js.Global.Get("jQuery.browser.version")
	FxInterval     = js.Global.Get("jQuery.fx.interval")
	FxOff          = js.Global.Get("jQuery.fx.off")
)

//JQuery constructor
func NewJQuery(args ...interface{}) JQuery {
	return newJQuery(JQ.New(args...))
}
