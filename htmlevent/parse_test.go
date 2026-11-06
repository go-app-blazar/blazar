package htmlevent

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
	"github.com/stretchr/testify/require"
)

type testValue struct {
	app.Value

	propertyMap map[string]any
}

func newTestValue() *testValue {
	return &testValue{
		propertyMap: make(map[string]any),
	}
}

func (t *testValue) String() string {
	return fmt.Sprintf("%+v", *t)
}

func (t *testValue) Get(m string) app.Value {
	property, ok := t.propertyMap[m]
	if !ok {
		return app.ValueOf(nil)
	}
	return app.ValueOf(property)
}

func (t *testValue) Set(m string, x any) {
	t.propertyMap[m] = x
}

func (t *testValue) Call(m string, args ...any) app.Value {
	property, ok := t.propertyMap[m]
	if !ok {
		return app.ValueOf(nil)
	}

	propertyValue := reflect.ValueOf(property)
	if propertyValue.Kind() != reflect.Func {
		return app.ValueOf(nil)
	}
	if propertyValue.IsZero() {
		return app.ValueOf(nil)
	}
	var argumentValues []reflect.Value
	{
		argumentValues = append(argumentValues, reflect.ValueOf(t))

		var appValueArgs []app.Value
		for _, arg := range args {
			appValueArgs = append(appValueArgs, app.ValueOf(arg))
		}
		argumentValues = append(argumentValues, reflect.ValueOf(appValueArgs))
	}
	_ = propertyValue.Call(argumentValues)
	return app.ValueOf(nil)
}

func TestParse(t *testing.T) {
	preventDefaultCalled := 0
	stopImmediatePropagationCalled := 0

	appEvent := app.Event{
		Value: newTestValue(),
	}
	appEvent.Set("preventDefault", func(this app.Value, args []app.Value) any {
		preventDefaultCalled++
		return nil
	})
	appEvent.Set("stopImmediatePropagation", func(this app.Value, args []app.Value) any {
		stopImmediatePropagationCalled++
		return nil
	})

	var e MouseEvent
	err := Parse(appEvent, &e)
	require.NoError(t, err)
	e.PreventDefault()
	require.Equal(t, 1, preventDefaultCalled)
	e.StopImmediatePropagation()
	require.Equal(t, 1, stopImmediatePropagationCalled)
}
