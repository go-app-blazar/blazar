package htmlevent

import "github.com/maxence-charriere/go-app/v11/pkg/app"

const (
	EventClick       = "click"
	EventDoubleClick = "dblclick"
	EventMouseDown   = "mousedown"
	EventMouseUp     = "mouseup"
	EventInput       = "input"
	EventKeydown     = "keydown"
	EventKeypress    = "keypress"
	EventKeyup       = "keyup"
	EventToggle      = "toggle"
)

// Event is the base event type.
// See: https://developer.mozilla.org/en-US/docs/Web/API/Event
type Event struct {
	Bubbles          bool      `javascript:"bubbles"`
	Cancelable       bool      `javascript:"cancelable"`
	Composed         bool      `javascript:"composed"`
	CurrentTarget    app.Value `javascript:"currentTarget"`
	DefaultPrevented bool      `javascript:"defaultPrevented"`
	EventPhase       int       `javascript:"eventPhase"`
	IsTrusted        bool      `javascript:"isTrusted"`
	Target           app.Value `javascript:"target"`
	TimeStamp        int64     `javascript:"timeStamp"`
	Type             string    `javascript:"type"`
}

// UIEvent is the base event type for UI events.
// See: https://developer.mozilla.org/en-US/docs/Web/API/UIEvent
type UIEvent struct {
	Event
	Detail int       `javascript:"detail"`
	View   app.Value `javascript:"view"`
}

// InputEvent is the event type for input events.
// See: https://developer.mozilla.org/en-US/docs/Web/API/InputEvent
type InputEvent struct {
	UIEvent
	Data         string    `javascript:"data"`
	DataTransfer app.Value `javascript:"dataTransfer"` // TODO: Build out the DataTransfer struct.
	InputType    string    `javascript:"inputType"`
	IsComposing  bool      `javascript:"isComposing"`
}

// KeyboardEvent is the event type for keyboard events.
// See: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent
type KeyboardEvent struct {
	UIEvent
	Code        string `javascript:"code"`
	CtrlKey     bool   `javascript:"ctrlKey"`
	IsComposing bool   `javascript:"isComposing"`
	Key         string `javascript:"key"`
	Location    int    `javascript:"location"`
	MetaKey     bool   `javascript:"metaKey"`
	Repeat      bool   `javascript:"repeat"`
	ShiftKey    bool   `javascript:"shiftKey"`
}

// MouseEvent is the event type for mouse events.
// See: https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent
type MouseEvent struct {
	UIEvent
	AltKey        bool      `javascript:"altKey"`
	Button        int       `javascript:"button"`
	Buttons       int       `javascript:"buttons"`
	ClientX       int       `javascript:"clientX"`
	ClientY       int       `javascript:"clientY"`
	CtrlKey       bool      `javascript:"ctrlKey"`
	MetaKey       bool      `javascript:"metaKey"`
	MovementX     int       `javascript:"movementX"`
	MovementY     int       `javascript:"movementY"`
	OffsetX       int       `javascript:"offsetX"`
	OffsetY       int       `javascript:"offsetY"`
	PageX         int       `javascript:"pageX"`
	PageY         int       `javascript:"pageY"`
	RelatedTarget app.Value `javascript:"relatedTarget"`
	ScreenX       int       `javascript:"screenX"`
	ScreenY       int       `javascript:"screenY"`
	ShiftKey      bool      `javascript:"shiftKey"`
}

// Toggle is the event type for toggle events.
// See: https://developer.mozilla.org/en-US/docs/Web/API/ToggleEvent
type Toggle struct {
	Event
	OldState string    `javascript:"oldState"`
	NewState string    `javascript:"newState"`
	Source   app.Value `javascript:"source"`
}
