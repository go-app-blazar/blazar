package htmlevent

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

// WrapInputEvent returns a InputEvent handler suitable for use with app.UI.On().
func WrapInputEvent(function func(ctx app.Context, e InputEvent)) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		var inputEvent InputEvent
		MustParse(e, &inputEvent)
		function(ctx, inputEvent)
	}
}

// WrapKeyboardEvent returns a KeyboardEvent handler suitable for use with app.UI.On().
func WrapKeyboardEvent(function func(ctx app.Context, e KeyboardEvent)) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		var keyboardEvent KeyboardEvent
		MustParse(e, &keyboardEvent)
		function(ctx, keyboardEvent)
	}
}

// WrapMouseEvent returns a MouseEvent handler suitable for use with app.UI.On().
func WrapMouseEvent(function func(ctx app.Context, e MouseEvent)) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		var mouseEvent MouseEvent
		MustParse(e, &mouseEvent)
		function(ctx, mouseEvent)
	}
}
