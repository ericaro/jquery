package jquery

import (
	"github.com/gopherjs/gopherjs/js"
)

type JQuery struct {
	*js.Object
	Context  *js.Object `js:"context"`
	Jquery   string     `js:"jquery"`
	Length   int        `js:"length"`
	Selector string     `js:"selector"`
}

func newJQuery(j *js.Object) JQuery {
	return JQuery{Object: j}
}

type Callbacks struct {
	*js.Object
}

func newCallbacks(j *js.Object) Callbacks {
	return Callbacks{Object: j}
}

type Deferred struct {
	*js.Object
}

func newDeferred(j *js.Object) Deferred {
	return Deferred{Object: j}
}

type Event struct {
	*js.Object
	CurrentTarget  *js.Object `js:"currentTarget"`
	Data           *js.Object `js:"data"`
	DelegateTarget *js.Object `js:"delegateTarget"`
	MetaKey        bool       `js:"metaKey"`
	Namespace      string     `js:"namespace"`
	PageX          float64    `js:"pageX"`
	PageY          float64    `js:"pageY"`
	RelatedTarget  *js.Object `js:"relatedTarget"`
	Result         *js.Object `js:"result"`
	Target         *js.Object `js:"target"`
	TimeStamp      float64    `js:"timeStamp"`
	Type           string     `js:"type"`
	Which          float64    `js:"which"`
}

func newEvent(j *js.Object) Event {
	return Event{Object: j}
}

// Create a new jQuery object with elements added to the set of matched elements.
func (x JQuery) Add(i ...interface {
}) JQuery {
	return newJQuery(x.Call("add", i...))
}

// Add the previous set of elements on the stack to the current set, optionally filtered by a selector.
func (x JQuery) AddBack(i ...interface {
}) JQuery {
	return newJQuery(x.Call("addBack", i...))
}

// Adds the specified class(es) to each of the set of matched elements.
func (x JQuery) AddClass(i ...interface {
}) JQuery {
	return newJQuery(x.Call("addClass", i...))
}

// Insert content, specified by the parameter, after each element in the set of matched elements.
func (x JQuery) After(i ...interface {
}) JQuery {
	return newJQuery(x.Call("after", i...))
}

// Register a handler to be called when Ajax requests complete. This is an .
func (x JQuery) AjaxComplete(handler *js.Object) JQuery {
	return newJQuery(x.Call("ajaxComplete", handler))
}

// Register a handler to be called when Ajax requests complete with an error. This is an .
func (x JQuery) AjaxError(handler *js.Object) JQuery {
	return newJQuery(x.Call("ajaxError", handler))
}

// Attach a function to be executed before an Ajax request is sent. This is an .
func (x JQuery) AjaxSend(handler *js.Object) JQuery {
	return newJQuery(x.Call("ajaxSend", handler))
}

// Register a handler to be called when the first Ajax request begins. This is an .
func (x JQuery) AjaxStart(handler *js.Object) JQuery {
	return newJQuery(x.Call("ajaxStart", handler))
}

// Register a handler to be called when all Ajax requests have completed. This is an .
func (x JQuery) AjaxStop(handler *js.Object) JQuery {
	return newJQuery(x.Call("ajaxStop", handler))
}

// Attach a function to be executed whenever an Ajax request completes successfully. This is an .
func (x JQuery) AjaxSuccess(handler *js.Object) JQuery {
	return newJQuery(x.Call("ajaxSuccess", handler))
}

// Add the previous set of elements on the stack to the current set.
func (x JQuery) AndSelf() JQuery {
	return newJQuery(x.Call("andSelf"))
}

// Perform a custom animation of a set of CSS properties.
func (x JQuery) Animate(i ...interface {
}) JQuery {
	return newJQuery(x.Call("animate", i...))
}

// Insert content, specified by the parameter, to the end of each element in the set of matched elements.
func (x JQuery) Append(i ...interface {
}) JQuery {
	return newJQuery(x.Call("append", i...))
}

// Insert every element in the set of matched elements to the end of the target.
func (x JQuery) AppendTo(target interface {
}) JQuery {
	return newJQuery(x.Call("appendTo", target))
}

// Get the value of an attribute for the first element in the set of matched elements.
// OR
// Set one or more attributes for the set of matched elements.
func (x JQuery) Attr(i ...interface {
}) *js.Object {
	return x.Call("attr", i...)
}

// Insert content, specified by the parameter, before each element in the set of matched elements.
func (x JQuery) Before(i ...interface {
}) JQuery {
	return newJQuery(x.Call("before", i...))
}

// Attach a handler to an event for the elements.
func (x JQuery) Bind(i ...interface {
}) JQuery {
	return newJQuery(x.Call("bind", i...))
}

// Bind an event handler to the "blur" JavaScript event, or trigger that event on an element.
func (x JQuery) Blur(i ...interface {
}) JQuery {
	return newJQuery(x.Call("blur", i...))
}

// Bind an event handler to the "change" JavaScript event, or trigger that event on an element.
func (x JQuery) Change(i ...interface {
}) JQuery {
	return newJQuery(x.Call("change", i...))
}

// Get the children of each element in the set of matched elements, optionally filtered by a selector.
func (x JQuery) Children(i ...interface {
}) JQuery {
	return newJQuery(x.Call("children", i...))
}

// Remove from the queue all items that have not yet been run.
func (x JQuery) ClearQueue(i ...interface {
}) JQuery {
	return newJQuery(x.Call("clearQueue", i...))
}

// Bind an event handler to the "click" JavaScript event, or trigger that event on an element.
func (x JQuery) Click(i ...interface {
}) JQuery {
	return newJQuery(x.Call("click", i...))
}

// Create a deep copy of the set of matched elements.
func (x JQuery) Clone(i ...interface {
}) JQuery {
	return newJQuery(x.Call("clone", i...))
}

// For each element in the set, get the first element that matches the selector by testing the element itself and traversing up through its ancestors in the DOM tree.
// OR
// Get an array of all the elements and selectors matched against the current element up through the DOM tree.
func (x JQuery) Closest(i ...interface {
}) *js.Object {
	return x.Call("closest", i...)
}

// Get the children of each element in the set of matched elements, including text and comment nodes.
func (x JQuery) Contents() JQuery {
	return newJQuery(x.Call("contents"))
}

// Get the computed style properties for the first element in the set of matched elements.
// OR
// Set one or more CSS properties for the set of matched elements.
func (x JQuery) Css(i ...interface {
}) *js.Object {
	return x.Call("css", i...)
}

// Store arbitrary data associated with the matched elements.
// OR
// Return the value at the named data store for the first element in the jQuery collection, as set by data(name, value) or by an HTML5 data-* attribute.
func (x JQuery) Data(i ...interface {
}) *js.Object {
	return x.Call("data", i...)
}

// Bind an event handler to the "dblclick" JavaScript event, or trigger that event on an element.
func (x JQuery) Dblclick(i ...interface {
}) JQuery {
	return newJQuery(x.Call("dblclick", i...))
}

// Set a timer to delay execution of subsequent items in the queue.
func (x JQuery) Delay(i ...interface {
}) JQuery {
	return newJQuery(x.Call("delay", i...))
}

// Attach a handler to one or more events for all elements that match the selector, now or in the future, based on a specific set of root elements.
func (x JQuery) Delegate(i ...interface {
}) JQuery {
	return newJQuery(x.Call("delegate", i...))
}

// Execute the next function on the queue for the matched elements.
func (x JQuery) Dequeue(i ...interface {
}) JQuery {
	return newJQuery(x.Call("dequeue", i...))
}

// Remove the set of matched elements from the DOM.
func (x JQuery) Detach(i ...interface {
}) JQuery {
	return newJQuery(x.Call("detach", i...))
}

// Remove event handlers previously attached using  from the elements.
func (x JQuery) Die(i ...interface {
}) JQuery {
	return newJQuery(x.Call("die", i...))
}

// Iterate over a jQuery object, executing a function for each matched element.
func (x JQuery) Each(function *js.Object) JQuery {
	return newJQuery(x.Call("each", function))
}

// Remove all child nodes of the set of matched elements from the DOM.
func (x JQuery) Empty() JQuery {
	return newJQuery(x.Call("empty"))
}

// End the most recent filtering operation in the current chain and return the set of matched elements to its previous state.
func (x JQuery) End() JQuery {
	return newJQuery(x.Call("end"))
}

// Reduce the set of matched elements to the one at the specified index.
func (x JQuery) Eq(i ...interface {
}) JQuery {
	return newJQuery(x.Call("eq", i...))
}

// Bind an event handler to the "error" JavaScript event.
func (x JQuery) Error(i ...interface {
}) JQuery {
	return newJQuery(x.Call("error", i...))
}

// Display the matched elements by fading them to opaque.
func (x JQuery) FadeIn(i ...interface {
}) JQuery {
	return newJQuery(x.Call("fadeIn", i...))
}

// Hide the matched elements by fading them to transparent.
func (x JQuery) FadeOut(i ...interface {
}) JQuery {
	return newJQuery(x.Call("fadeOut", i...))
}

// Adjust the opacity of the matched elements.
func (x JQuery) FadeTo(i ...interface {
}) JQuery {
	return newJQuery(x.Call("fadeTo", i...))
}

// Display or hide the matched elements by animating their opacity.
func (x JQuery) FadeToggle(i ...interface {
}) JQuery {
	return newJQuery(x.Call("fadeToggle", i...))
}

// Reduce the set of matched elements to those that match the selector or pass the function's test.
func (x JQuery) Filter(i ...interface {
}) JQuery {
	return newJQuery(x.Call("filter", i...))
}

// Get the descendants of each element in the current set of matched elements, filtered by a selector, jQuery object, or element.
func (x JQuery) Find(i ...interface {
}) JQuery {
	return newJQuery(x.Call("find", i...))
}

// Stop the currently-running animation, remove all queued animations, and complete all animations for the matched elements.
func (x JQuery) Finish(i ...interface {
}) JQuery {
	return newJQuery(x.Call("finish", i...))
}

// Reduce the set of matched elements to the first in the set.
func (x JQuery) First() JQuery {
	return newJQuery(x.Call("first"))
}

// Bind an event handler to the "focus" JavaScript event, or trigger that event on an element.
func (x JQuery) Focus(i ...interface {
}) JQuery {
	return newJQuery(x.Call("focus", i...))
}

// Bind an event handler to the "focusin" event.
func (x JQuery) Focusin(i ...interface {
}) JQuery {
	return newJQuery(x.Call("focusin", i...))
}

// Bind an event handler to the "focusout" JavaScript event.
func (x JQuery) Focusout(i ...interface {
}) JQuery {
	return newJQuery(x.Call("focusout", i...))
}

// Retrieve one of the elements matched by the jQuery object.
// OR
// Retrieve the elements matched by the jQuery object.
func (x JQuery) Get(i ...interface {
}) *js.Object {
	return x.Call("get", i...)
}

// Reduce the set of matched elements to those that have a descendant that matches the selector or DOM element.
func (x JQuery) Has(i ...interface {
}) JQuery {
	return newJQuery(x.Call("has", i...))
}

// Determine whether any of the matched elements are assigned the given class.
func (x JQuery) HasClass(className string) bool {
	return x.Call("hasClass", className).Bool()
}

// Get the current computed height for the first element in the set of matched elements.
// OR
// Set the CSS height of every matched element.
func (x JQuery) Height(i ...interface {
}) *js.Object {
	return x.Call("height", i...)
}

// Hide the matched elements.
func (x JQuery) Hide(i ...interface {
}) JQuery {
	return newJQuery(x.Call("hide", i...))
}

// Bind two handlers to the matched elements, to be executed when the mouse pointer enters and leaves the elements.
// OR
// Bind a single handler to the matched elements, to be executed when the mouse pointer enters or leaves the elements.
func (x JQuery) Hover(i ...interface {
}) JQuery {
	return newJQuery(x.Call("hover", i...))
}

// Get the HTML contents of the first element in the set of matched elements.
// OR
// Set the HTML contents of each element in the set of matched elements.
func (x JQuery) Html(i ...interface {
}) *js.Object {
	return x.Call("html", i...)
}

// Search for a given element from among the matched elements.
func (x JQuery) Index(i ...interface {
}) float64 {
	return x.Call("index", i...).Float()
}

// Get the current computed height for the first element in the set of matched elements, including padding but not border.
// OR
// Set the CSS inner height of each element in the set of matched elements.
func (x JQuery) InnerHeight(i ...interface {
}) *js.Object {
	return x.Call("innerHeight", i...)
}

// Get the current computed inner width for the first element in the set of matched elements, including padding but not border.
// OR
// Set the CSS inner width of each element in the set of matched elements.
func (x JQuery) InnerWidth(i ...interface {
}) *js.Object {
	return x.Call("innerWidth", i...)
}

// Insert every element in the set of matched elements after the target.
func (x JQuery) InsertAfter(target interface {
}) JQuery {
	return newJQuery(x.Call("insertAfter", target))
}

// Insert every element in the set of matched elements before the target.
func (x JQuery) InsertBefore(target interface {
}) JQuery {
	return newJQuery(x.Call("insertBefore", target))
}

// Check the current matched set of elements against a selector, element, or jQuery object and return  if at least one of these elements matches the given arguments.
func (x JQuery) Is(i ...interface {
}) bool {
	return x.Call("is", i...).Bool()
}

// Accepts a string containing a CSS selector which is then used to match a set of elements.
// OR
// Creates DOM elements on the fly from the provided string of raw HTML.
// OR
// Binds a function to be executed when the DOM has finished loading.
func (x JQuery) JQuery(i ...interface {
}) JQuery {
	return newJQuery(x.Call("jQuery", i...))
}

// Bind an event handler to the "keydown" JavaScript event, or trigger that event on an element.
func (x JQuery) Keydown(i ...interface {
}) JQuery {
	return newJQuery(x.Call("keydown", i...))
}

// Bind an event handler to the "keypress" JavaScript event, or trigger that event on an element.
func (x JQuery) Keypress(i ...interface {
}) JQuery {
	return newJQuery(x.Call("keypress", i...))
}

// Bind an event handler to the "keyup" JavaScript event, or trigger that event on an element.
func (x JQuery) Keyup(i ...interface {
}) JQuery {
	return newJQuery(x.Call("keyup", i...))
}

// Reduce the set of matched elements to the final one in the set.
func (x JQuery) Last() JQuery {
	return newJQuery(x.Call("last"))
}

// Attach an event handler for all elements which match the current selector, now and in the future.
func (x JQuery) Live(i ...interface {
}) JQuery {
	return newJQuery(x.Call("live", i...))
}

// Bind an event handler to the "load" JavaScript event.
// OR
// Load data from the server and place the returned HTML into the matched element.
func (x JQuery) Load(i ...interface {
}) JQuery {
	return newJQuery(x.Call("load", i...))
}

// Pass each element in the current matched set through a function, producing a new jQuery object containing the return values.
func (x JQuery) Map(callback *js.Object) JQuery {
	return newJQuery(x.Call("map", callback))
}

// Bind an event handler to the "mousedown" JavaScript event, or trigger that event on an element.
func (x JQuery) Mousedown(i ...interface {
}) JQuery {
	return newJQuery(x.Call("mousedown", i...))
}

// Bind an event handler to be fired when the mouse enters an element, or trigger that handler on an element.
func (x JQuery) Mouseenter(i ...interface {
}) JQuery {
	return newJQuery(x.Call("mouseenter", i...))
}

// Bind an event handler to be fired when the mouse leaves an element, or trigger that handler on an element.
func (x JQuery) Mouseleave(i ...interface {
}) JQuery {
	return newJQuery(x.Call("mouseleave", i...))
}

// Bind an event handler to the "mousemove" JavaScript event, or trigger that event on an element.
func (x JQuery) Mousemove(i ...interface {
}) JQuery {
	return newJQuery(x.Call("mousemove", i...))
}

// Bind an event handler to the "mouseout" JavaScript event, or trigger that event on an element.
func (x JQuery) Mouseout(i ...interface {
}) JQuery {
	return newJQuery(x.Call("mouseout", i...))
}

// Bind an event handler to the "mouseover" JavaScript event, or trigger that event on an element.
func (x JQuery) Mouseover(i ...interface {
}) JQuery {
	return newJQuery(x.Call("mouseover", i...))
}

// Bind an event handler to the "mouseup" JavaScript event, or trigger that event on an element.
func (x JQuery) Mouseup(i ...interface {
}) JQuery {
	return newJQuery(x.Call("mouseup", i...))
}

// Get the immediately following sibling of each element in the set of matched elements. If a selector is provided, it retrieves the next sibling only if it matches that selector.
func (x JQuery) Next(i ...interface {
}) JQuery {
	return newJQuery(x.Call("next", i...))
}

// Get all following siblings of each element in the set of matched elements, optionally filtered by a selector.
func (x JQuery) NextAll(i ...interface {
}) JQuery {
	return newJQuery(x.Call("nextAll", i...))
}

// Get all following siblings of each element up to but not including the element matched by the selector, DOM node, or jQuery object passed.
func (x JQuery) NextUntil(i ...interface {
}) JQuery {
	return newJQuery(x.Call("nextUntil", i...))
}

// Remove elements from the set of matched elements.
func (x JQuery) Not(i ...interface {
}) JQuery {
	return newJQuery(x.Call("not", i...))
}

// Remove an event handler.
func (x JQuery) Off(i ...interface {
}) JQuery {
	return newJQuery(x.Call("off", i...))
}

// Get the current coordinates of the first element in the set of matched elements, relative to the document.
// OR
// Set the current coordinates of every element in the set of matched elements, relative to the document.
func (x JQuery) Offset(i ...interface {
}) *js.Object {
	return x.Call("offset", i...)
}

// Get the closest ancestor element that is positioned.
func (x JQuery) OffsetParent() JQuery {
	return newJQuery(x.Call("offsetParent"))
}

// Attach an event handler function for one or more events to the selected elements.
func (x JQuery) On(i ...interface {
}) JQuery {
	return newJQuery(x.Call("on", i...))
}

// Attach a handler to an event for the elements. The handler is executed at most once per element per event type.
func (x JQuery) One(i ...interface {
}) JQuery {
	return newJQuery(x.Call("one", i...))
}

// Get the current computed height for the first element in the set of matched elements, including padding, border, and optionally margin. Returns a number (without "px") representation of the value or null if called on an empty set of elements.
// OR
// Set the CSS outer Height of each element in the set of matched elements.
func (x JQuery) OuterHeight(i ...interface {
}) *js.Object {
	return x.Call("outerHeight", i...)
}

// Get the current computed width for the first element in the set of matched elements, including padding and border.
// OR
// Set the CSS outer width of each element in the set of matched elements.
func (x JQuery) OuterWidth(i ...interface {
}) *js.Object {
	return x.Call("outerWidth", i...)
}

// Get the parent of each element in the current set of matched elements, optionally filtered by a selector.
func (x JQuery) Parent(i ...interface {
}) JQuery {
	return newJQuery(x.Call("parent", i...))
}

// Get the ancestors of each element in the current set of matched elements, optionally filtered by a selector.
func (x JQuery) Parents(i ...interface {
}) JQuery {
	return newJQuery(x.Call("parents", i...))
}

// Get the ancestors of each element in the current set of matched elements, up to but not including the element matched by the selector, DOM node, or jQuery object.
func (x JQuery) ParentsUntil(i ...interface {
}) JQuery {
	return newJQuery(x.Call("parentsUntil", i...))
}

// Get the current coordinates of the first element in the set of matched elements, relative to the offset parent.
func (x JQuery) Position() *js.Object {
	return x.Call("position")
}

// Insert content, specified by the parameter, to the beginning of each element in the set of matched elements.
func (x JQuery) Prepend(i ...interface {
}) JQuery {
	return newJQuery(x.Call("prepend", i...))
}

// Insert every element in the set of matched elements to the beginning of the target.
func (x JQuery) PrependTo(target interface {
}) JQuery {
	return newJQuery(x.Call("prependTo", target))
}

// Get the immediately preceding sibling of each element in the set of matched elements, optionally filtered by a selector.
func (x JQuery) Prev(i ...interface {
}) JQuery {
	return newJQuery(x.Call("prev", i...))
}

// Get all preceding siblings of each element in the set of matched elements, optionally filtered by a selector.
func (x JQuery) PrevAll(i ...interface {
}) JQuery {
	return newJQuery(x.Call("prevAll", i...))
}

// Get all preceding siblings of each element up to but not including the element matched by the selector, DOM node, or jQuery object.
func (x JQuery) PrevUntil(i ...interface {
}) JQuery {
	return newJQuery(x.Call("prevUntil", i...))
}

//  Return a Promise object to observe when all actions of a certain type bound to the collection, queued or not, have finished.
func (x JQuery) Promise(i ...interface {
}) *js.Object {
	return x.Call("promise", i...)
}

// Get the value of a property for the first element in the set of matched elements.
// OR
// Set one or more properties for the set of matched elements.
func (x JQuery) Prop(i ...interface {
}) *js.Object {
	return x.Call("prop", i...)
}

// Add a collection of DOM elements onto the jQuery stack.
func (x JQuery) PushStack(i ...interface {
}) JQuery {
	return newJQuery(x.Call("pushStack", i...))
}

// Show the queue of functions to be executed on the matched elements.
// OR
// Manipulate the queue of functions to be executed, once for each matched element.
func (x JQuery) Queue(i ...interface {
}) *js.Object {
	return x.Call("queue", i...)
}

// Specify a function to execute when the DOM is fully loaded.
func (x JQuery) Ready(handler *js.Object) JQuery {
	return newJQuery(x.Call("ready", handler))
}

// Remove the set of matched elements from the DOM.
func (x JQuery) Remove(i ...interface {
}) JQuery {
	return newJQuery(x.Call("remove", i...))
}

// Remove an attribute from each element in the set of matched elements.
func (x JQuery) RemoveAttr(attributeName string) JQuery {
	return newJQuery(x.Call("removeAttr", attributeName))
}

// Remove a single class, multiple classes, or all classes from each element in the set of matched elements.
func (x JQuery) RemoveClass(i ...interface {
}) JQuery {
	return newJQuery(x.Call("removeClass", i...))
}

// Remove a previously-stored piece of data.
func (x JQuery) RemoveData(i ...interface {
}) JQuery {
	return newJQuery(x.Call("removeData", i...))
}

// Remove a property for the set of matched elements.
func (x JQuery) RemoveProp(propertyName string) JQuery {
	return newJQuery(x.Call("removeProp", propertyName))
}

// Replace each target element with the set of matched elements.
func (x JQuery) ReplaceAll(target interface {
}) JQuery {
	return newJQuery(x.Call("replaceAll", target))
}

// Replace each element in the set of matched elements with the provided new content and return the set of elements that was removed.
func (x JQuery) ReplaceWith(i ...interface {
}) JQuery {
	return newJQuery(x.Call("replaceWith", i...))
}

// Bind an event handler to the "resize" JavaScript event, or trigger that event on an element.
func (x JQuery) Resize(i ...interface {
}) JQuery {
	return newJQuery(x.Call("resize", i...))
}

// Bind an event handler to the "scroll" JavaScript event, or trigger that event on an element.
func (x JQuery) Scroll(i ...interface {
}) JQuery {
	return newJQuery(x.Call("scroll", i...))
}

// Get the current horizontal position of the scroll bar for the first element in the set of matched elements.
// OR
// Set the current horizontal position of the scroll bar for each of the set of matched elements.
func (x JQuery) ScrollLeft(i ...interface {
}) *js.Object {
	return x.Call("scrollLeft", i...)
}

// Get the current vertical position of the scroll bar for the first element in the set of matched elements or set the vertical position of the scroll bar for every matched element.
// OR
// Set the current vertical position of the scroll bar for each of the set of matched elements.
func (x JQuery) ScrollTop(i ...interface {
}) *js.Object {
	return x.Call("scrollTop", i...)
}

// Bind an event handler to the "select" JavaScript event, or trigger that event on an element.
func (x JQuery) Select(i ...interface {
}) JQuery {
	return newJQuery(x.Call("select", i...))
}

// Encode a set of form elements as a string for submission.
func (x JQuery) Serialize() string {
	return x.Call("serialize").String()
}

// Encode a set of form elements as an array of names and values.
func (x JQuery) SerializeArray() *js.Object {
	return x.Call("serializeArray")
}

// Display the matched elements.
func (x JQuery) Show(i ...interface {
}) JQuery {
	return newJQuery(x.Call("show", i...))
}

// Get the siblings of each element in the set of matched elements, optionally filtered by a selector.
func (x JQuery) Siblings(i ...interface {
}) JQuery {
	return newJQuery(x.Call("siblings", i...))
}

// Return the number of elements in the jQuery object.
func (x JQuery) Size() int {
	return x.Call("size").Int()
}

// Reduce the set of matched elements to a subset specified by a range of indices.
func (x JQuery) Slice(i ...interface {
}) JQuery {
	return newJQuery(x.Call("slice", i...))
}

// Display the matched elements with a sliding motion.
func (x JQuery) SlideDown(i ...interface {
}) JQuery {
	return newJQuery(x.Call("slideDown", i...))
}

// Display or hide the matched elements with a sliding motion.
func (x JQuery) SlideToggle(i ...interface {
}) JQuery {
	return newJQuery(x.Call("slideToggle", i...))
}

// Hide the matched elements with a sliding motion.
func (x JQuery) SlideUp(i ...interface {
}) JQuery {
	return newJQuery(x.Call("slideUp", i...))
}

// Stop the currently-running animation on the matched elements.
func (x JQuery) Stop(i ...interface {
}) JQuery {
	return newJQuery(x.Call("stop", i...))
}

// Bind an event handler to the "submit" JavaScript event, or trigger that event on an element.
func (x JQuery) Submit(i ...interface {
}) JQuery {
	return newJQuery(x.Call("submit", i...))
}

// Get the combined text contents of each element in the set of matched elements, including their descendants.
// OR
// Set the content of each element in the set of matched elements to the specified text.
func (x JQuery) Text(i ...interface {
}) *js.Object {
	return x.Call("text", i...)
}

// Retrieve all the elements contained in the jQuery set, as an array.
func (x JQuery) ToArray() *js.Object {
	return x.Call("toArray")
}

// Bind two or more handlers to the matched elements, to be executed on alternate clicks.
// OR
// Display or hide the matched elements.
func (x JQuery) Toggle(i ...interface {
}) JQuery {
	return newJQuery(x.Call("toggle", i...))
}

// Add or remove one or more classes from each element in the set of matched elements, depending on either the class's presence or the value of the state argument.
func (x JQuery) ToggleClass(i ...interface {
}) JQuery {
	return newJQuery(x.Call("toggleClass", i...))
}

// Execute all handlers and behaviors attached to the matched elements for the given event type.
func (x JQuery) Trigger(i ...interface {
}) JQuery {
	return newJQuery(x.Call("trigger", i...))
}

// Execute all handlers attached to an element for an event.
func (x JQuery) TriggerHandler(i ...interface {
}) *js.Object {
	return x.Call("triggerHandler", i...)
}

// Remove a previously-attached event handler from the elements.
func (x JQuery) Unbind(i ...interface {
}) JQuery {
	return newJQuery(x.Call("unbind", i...))
}

// Remove a handler from the event for all elements which match the current selector, based upon a specific set of root elements.
func (x JQuery) Undelegate(i ...interface {
}) JQuery {
	return newJQuery(x.Call("undelegate", i...))
}

// Bind an event handler to the "unload" JavaScript event.
func (x JQuery) Unload(i ...interface {
}) JQuery {
	return newJQuery(x.Call("unload", i...))
}

// Remove the parents of the set of matched elements from the DOM, leaving the matched elements in their place.
func (x JQuery) Unwrap() JQuery {
	return newJQuery(x.Call("unwrap"))
}

// Get the current value of the first element in the set of matched elements.
// OR
// Set the value of each element in the set of matched elements.
func (x JQuery) Val(i ...interface {
}) *js.Object {
	return x.Call("val", i...)
}

// Get the current computed width for the first element in the set of matched elements.
// OR
// Set the CSS width of each element in the set of matched elements.
func (x JQuery) Width(i ...interface {
}) *js.Object {
	return x.Call("width", i...)
}

// Wrap an HTML structure around each element in the set of matched elements.
func (x JQuery) Wrap(i ...interface {
}) JQuery {
	return newJQuery(x.Call("wrap", i...))
}

// Wrap an HTML structure around all elements in the set of matched elements.
func (x JQuery) WrapAll(i ...interface {
}) JQuery {
	return newJQuery(x.Call("wrapAll", i...))
}

// Wrap an HTML structure around the content of each element in the set of matched elements.
func (x JQuery) WrapInner(i ...interface {
}) JQuery {
	return newJQuery(x.Call("wrapInner", i...))
}

// Add a callback or a collection of callbacks to a callback list.
func (x Callbacks) Add(callbacks interface {
}) Callbacks {
	return newCallbacks(x.Call("add", callbacks))
}

// Disable a callback list from doing anything more.
func (x Callbacks) Disable() Callbacks {
	return newCallbacks(x.Call("disable"))
}

// Determine if the callbacks list has been disabled.
func (x Callbacks) Disabled() bool {
	return x.Call("disabled").Bool()
}

// Remove all of the callbacks from a list.
func (x Callbacks) Empty() Callbacks {
	return newCallbacks(x.Call("empty"))
}

// Call all of the callbacks with the given arguments
func (x Callbacks) Fire(arguments *js.Object) Callbacks {
	return newCallbacks(x.Call("fire", arguments))
}

// Call all callbacks in a list with the given context and arguments.
func (x Callbacks) FireWith(i ...interface {
}) Callbacks {
	return newCallbacks(x.Call("fireWith", i...))
}

// Determine if the callbacks have already been called at least once.
func (x Callbacks) Fired() bool {
	return x.Call("fired").Bool()
}

// Determine whether or not the list has any callbacks attached. If a callback is provided as an argument, determine whether it is in a list.
func (x Callbacks) Has(i ...interface {
}) bool {
	return x.Call("has", i...).Bool()
}

// Lock a callback list in its current state.
func (x Callbacks) Lock() Callbacks {
	return newCallbacks(x.Call("lock"))
}

// Determine if the callbacks list has been locked.
func (x Callbacks) Locked() bool {
	return x.Call("locked").Bool()
}

// Remove a callback or a collection of callbacks from a callback list.
func (x Callbacks) Remove(callbacks interface {
}) Callbacks {
	return newCallbacks(x.Call("remove", callbacks))
}

//  Add handlers to be called when the Deferred object is either resolved or rejected.
func (x Deferred) Always(i ...interface {
}) Deferred {
	return newDeferred(x.Call("always", i...))
}

//  Add handlers to be called when the Deferred object is resolved.
func (x Deferred) Done(i ...interface {
}) Deferred {
	return newDeferred(x.Call("done", i...))
}

//  Add handlers to be called when the Deferred object is rejected.
func (x Deferred) Fail(i ...interface {
}) Deferred {
	return newDeferred(x.Call("fail", i...))
}

//  Determine whether a Deferred object has been rejected.
func (x Deferred) IsRejected() bool {
	return x.Call("isRejected").Bool()
}

//  Determine whether a Deferred object has been resolved.
func (x Deferred) IsResolved() bool {
	return x.Call("isResolved").Bool()
}

//  Call the progressCallbacks on a Deferred object with the given .
func (x Deferred) Notify(args *js.Object) Deferred {
	return newDeferred(x.Call("notify", args))
}

//  Call the progressCallbacks on a Deferred object with the given context and .
func (x Deferred) NotifyWith(i ...interface {
}) Deferred {
	return newDeferred(x.Call("notifyWith", i...))
}

//  Utility method to filter and/or chain Deferreds.
func (x Deferred) Pipe(i ...interface {
}) *js.Object {
	return x.Call("pipe", i...)
}

//  Add handlers to be called when the Deferred object generates progress notifications.
func (x Deferred) Progress(progressCallbacks interface {
}, progressCallbacks interface {
}) Deferred {
	return newDeferred(x.Call("progress", progressCallbacks, progressCallbacks))
}

//  Return a Deferred's Promise object.
func (x Deferred) Promise(i ...interface {
}) *js.Object {
	return x.Call("promise", i...)
}

//  Reject a Deferred object and call any failCallbacks with the given .
func (x Deferred) Reject(i ...interface {
}) Deferred {
	return newDeferred(x.Call("reject", i...))
}

//  Reject a Deferred object and call any failCallbacks with the given  and .
func (x Deferred) RejectWith(i ...interface {
}) Deferred {
	return newDeferred(x.Call("rejectWith", i...))
}

//  Resolve a Deferred object and call any doneCallbacks with the given .
func (x Deferred) Resolve(i ...interface {
}) Deferred {
	return newDeferred(x.Call("resolve", i...))
}

//  Resolve a Deferred object and call any doneCallbacks with the given  and .
func (x Deferred) ResolveWith(i ...interface {
}) Deferred {
	return newDeferred(x.Call("resolveWith", i...))
}

// Determine the current state of a Deferred object.
func (x Deferred) State() string {
	return x.Call("state").String()
}

// Add handlers to be called when the Deferred object is resolved, rejected, or still in progress.
func (x Deferred) Then(i ...interface {
}) *js.Object {
	return x.Call("then", i...)
}

// Returns whether  was ever called on this event object.
func (x Event) IsDefaultPrevented() bool {
	return x.Call("isDefaultPrevented").Bool()
}

//   Returns whether event.stopImmediatePropagation() was ever called on this event object.
func (x Event) IsImmediatePropagationStopped() bool {
	return x.Call("isImmediatePropagationStopped").Bool()
}

//   Returns whether  was ever called on this event object.
func (x Event) IsPropagationStopped() bool {
	return x.Call("isPropagationStopped").Bool()
}

// If this method is called, the default action of the event will not be triggered.
func (x Event) PreventDefault() interface {
} {
	return x.Call("preventDefault").Interface()
}

// Keeps the rest of the handlers from being executed and prevents the event from bubbling up the DOM tree.
func (x Event) StopImmediatePropagation() interface {
} {
	return x.Call("stopImmediatePropagation").Interface()
}

// Prevents the event from bubbling up the DOM tree, preventing any parent handlers from being notified of the event.
func (x Event) StopPropagation() interface {
} {
	return x.Call("stopPropagation").Interface()
}

// Merge the contents of an object onto the jQuery prototype to provide new jQuery instance methods.
func FnExtend(object *js.Object) *js.Object {
	return JQ.Call("FnExtend", object)
}

// Perform an asynchronous HTTP (Ajax) request.
func Ajax(i ...interface {
}) *js.Object {
	return JQ.Call("ajax", i...)
}

// Handle custom Ajax options or modify existing options before each request is sent and before they are processed by .
func AjaxPrefilter(i ...interface {
}) interface {
} {
	return JQ.Call("ajaxPrefilter", i...).Interface()
}

// Set default values for future Ajax requests. Its use is not recommended.
func AjaxSetup(options *js.Object) interface {
} {
	return JQ.Call("ajaxSetup", options).Interface()
}

// Creates an object that handles the actual transmission of Ajax data.
func AjaxTransport(dataType string, handler *js.Object) interface {
} {
	return JQ.Call("ajaxTransport", dataType, handler).Interface()
}

// Check to see if a DOM element is a descendant of another DOM element.
func Contains(container *js.Object, contained *js.Object) bool {
	return JQ.Call("contains", container, contained).Bool()
}

// Store arbitrary data associated with the specified element. Returns the value that was set.
// OR
// Returns value at named data store for the element, as set by , or the full data store for the element.
func Data(i ...interface {
}) *js.Object {
	return JQ.Call("data", i...)
}

// Execute the next function on the queue for the matched element.
func Dequeue(i ...interface {
}) interface {
} {
	return JQ.Call("dequeue", i...).Interface()
}

// A generic iterator function, which can be used to seamlessly iterate over both objects and arrays. Arrays and array-like objects with a length property (such as a function's arguments object) are iterated by numeric index, from 0 to length-1. Other objects are iterated via their named properties.
func Each(i ...interface {
}) *js.Object {
	return JQ.Call("each", i...)
}

// Takes a string and throws an exception containing it.
func Error(message string) interface {
} {
	return JQ.Call("error", message).Interface()
}

// Merge the contents of two or more objects together into the first object.
func Extend(i ...interface {
}) *js.Object {
	return JQ.Call("extend", i...)
}

// Load data from the server using a HTTP GET request.
func Get(i ...interface {
}) *js.Object {
	return JQ.Call("get", i...)
}

// Load JSON-encoded data from the server using a GET HTTP request.
func GetJSON(i ...interface {
}) *js.Object {
	return JQ.Call("getJSON", i...)
}

// Load a JavaScript file from the server using a GET HTTP request, then execute it.
func GetScript(i ...interface {
}) *js.Object {
	return JQ.Call("getScript", i...)
}

// Execute some JavaScript code globally.
func GlobalEval(code string) interface {
} {
	return JQ.Call("globalEval", code).Interface()
}

// Finds the elements of an array which satisfy a filter function. The original array is not affected.
func Grep(i ...interface {
}) *js.Object {
	return JQ.Call("grep", i...)
}

// Determine whether an element has any jQuery data associated with it.
func HasData(element *js.Object) bool {
	return JQ.Call("hasData", element).Bool()
}

// Holds or releases the execution of jQuery's ready event.
func HoldReady(hold bool) interface {
} {
	return JQ.Call("holdReady", hold).Interface()
}

// Search for a specified value within an array and return its index (or -1 if not found).
func InArray(i ...interface {
}) float64 {
	return JQ.Call("inArray", i...).Float()
}

// Determine whether the argument is an array.
func IsArray(obj *js.Object) bool {
	return JQ.Call("isArray", obj).Bool()
}

// Check to see if an object is empty (contains no enumerable properties).
func IsEmptyObject(object *js.Object) bool {
	return JQ.Call("isEmptyObject", object).Bool()
}

// Determine if the argument passed is a Javascript function object.
func IsFunction(obj *js.Object) bool {
	return JQ.Call("isFunction", obj).Bool()
}

// Determines whether its argument is a number.
func IsNumeric(value *js.Object) bool {
	return JQ.Call("isNumeric", value).Bool()
}

// Check to see if an object is a plain object (created using "{}" or "new Object").
func IsPlainObject(object *js.Object) bool {
	return JQ.Call("isPlainObject", object).Bool()
}

// Determine whether the argument is a window.
func IsWindow(obj *js.Object) bool {
	return JQ.Call("isWindow", obj).Bool()
}

// Check to see if a DOM node is within an XML document (or is an XML document).
func IsXMLDoc(node *js.Object) bool {
	return JQ.Call("isXMLDoc", node).Bool()
}

// Convert an array-like object into a true JavaScript array.
func MakeArray(obj *js.Object) *js.Object {
	return JQ.Call("makeArray", obj)
}

// Translate all items in an array or object to new array of items.
func Map(i ...interface {
}) *js.Object {
	return JQ.Call("map", i...)
}

// Merge the contents of two arrays together into the first array.
func Merge(first *js.Object, second *js.Object) *js.Object {
	return JQ.Call("merge", first, second)
}

// A multi-purpose callbacks list object that provides a powerful way to manage callback lists.
func NewCallbacks(flags string) Callbacks {
	return newCallbacks(JQ.Call("newCallbacks", flags))
}

//  A factory function that returns a chainable utility object with methods to register multiple callbacks into callback queues, invoke callback queues, and relay the success or failure state of any synchronous or asynchronous function.
func NewDeferred(i ...interface {
}) Deferred {
	return newDeferred(JQ.Call("newDeferred", i...))
}

// Relinquish jQuery's control of the  variable.
func NoConflict(i ...interface {
}) *js.Object {
	return JQ.Call("noConflict", i...)
}

// An empty function.
func Noop() interface {
} {
	return JQ.Call("noop").Interface()
}

// Return a number representing the current time.
func Now() float64 {
	return JQ.Call("now").Float()
}

// Create a serialized representation of an array, a plain object, or a jQuery object suitable for use in a URL query string or Ajax request. In case a jQuery object is passed, it should contain input elements with name/value properties.
func Param(i ...interface {
}) string {
	return JQ.Call("param", i...).String()
}

// Parses a string into an array of DOM nodes.
func ParseHTML(i ...interface {
}) *js.Object {
	return JQ.Call("parseHTML", i...)
}

// Takes a well-formed JSON string and returns the resulting JavaScript value.
func ParseJSON(json string) interface {
} {
	return JQ.Call("parseJSON", json).Interface()
}

// Parses a string into an XML document.
func ParseXML(data string) *js.Object {
	return JQ.Call("parseXML", data)
}

// Load data from the server using a HTTP POST request.
func Post(i ...interface {
}) *js.Object {
	return JQ.Call("post", i...)
}

// Takes a function and returns a new one that will always have a particular context.
func Proxy(i ...interface {
}) *js.Object {
	return JQ.Call("proxy", i...)
}

// Show the queue of functions to be executed on the matched element.
// OR
// Manipulate the queue of functions to be executed on the matched element.
func Queue(i ...interface {
}) *js.Object {
	return JQ.Call("queue", i...)
}

// Remove a previously-stored piece of data.
func RemoveData(i ...interface {
}) JQuery {
	return newJQuery(JQ.Call("removeData", i...))
}

// Creates a new copy of jQuery whose properties and methods can be modified without affecting the original jQuery object.
func Sub() JQuery {
	return newJQuery(JQ.Call("sub"))
}

// Remove the whitespace from the beginning and end of a string.
func Trim(str string) string {
	return JQ.Call("trim", str).String()
}

// Determine the internal JavaScript [[Class]] of an object.
func Type(obj *js.Object) string {
	return JQ.Call("type", obj).String()
}

// Sorts an array of DOM elements, in place, with the duplicates removed. Note that this only works on arrays of DOM elements, not strings or numbers.
func Unique(array *js.Object) *js.Object {
	return JQ.Call("unique", array)
}

// Provides a way to execute callback functions based on one or more objects, usually  objects that represent asynchronous events.
func When(deferreds Deferred) *js.Object {
	return JQ.Call("when", deferreds)
}
