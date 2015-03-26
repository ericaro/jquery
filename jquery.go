package jquery

//go:generate go run ./apigenerator/jquery-gen/main.go -i ./api.jquery.com/entries/ -o api.go

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

//JQuery constructor
func NewJQuery(args ...interface{}) JQuery {
	return newJQuery(JQ.New(args...))
}
