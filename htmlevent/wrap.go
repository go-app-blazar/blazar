package htmlevent

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

// WrapMouseEvent returns a MouseEvent handler suitable for use with app.UI.On().
func WrapMouseEvent(function func(ctx app.Context, e MouseEvent)) app.EventHandler {
	return func(ctx app.Context, e app.Event) {
		var mouseEvent MouseEvent
		MustParse(e, &mouseEvent)
		function(ctx, mouseEvent)
	}
}
