package htmlevent

import (
	"testing"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	preventDefaultCalled := 0
	stopImmediatePropagationCalled := 0

	appEvent := app.Event{
		Value: app.Window().Get("Object").New(),
	}
	appEvent.Set("preventDefault", app.FuncOf(func(this app.Value, args []app.Value) any {
		preventDefaultCalled++
		return nil
	}))
	appEvent.Set("stopImmediatePropagation", app.FuncOf(func(this app.Value, args []app.Value) any {
		stopImmediatePropagationCalled++
		return nil
	}))
	appEvent.Set("altKey", true)
	appEvent.Set("clientX", 100)
	appEvent.Set("clientY", 200)
	appEvent.Set("shiftKey", true)

	var e MouseEvent
	err := Parse(appEvent, &e)
	require.NoError(t, err)

	require.True(t, e.AltKey)
	require.Equal(t, 100, e.ClientX)
	require.Equal(t, 200, e.ClientY)
	require.True(t, e.ShiftKey)

	e.PreventDefault()
	require.Equal(t, 1, preventDefaultCalled)
	e.StopImmediatePropagation()
	require.Equal(t, 1, stopImmediatePropagationCalled)
}
