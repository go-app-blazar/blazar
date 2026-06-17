package blazar

import (
	"context"
	"log/slog"
	"reflect"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type HasData[T any] interface {
	DataSet(name string, value any) T
}

type UseData struct {
	Data map[string]any
}

func (c *UseData) DataSet(name string, value any) *UseData {
	if c.Data == nil {
		c.Data = map[string]any{}
	}
	c.Data[name] = value
	return c
}

func (c *UseData) Wrap(element app.UI) app.UI {
	func() {
		//slog.Debug("Wrap", "element", element)

		elementValue := reflect.ValueOf(element)
		//slog.Debug("Wrap", "elementValue", elementValue)

		// TODO: Can we remove this block?
		/*for elementValue.Kind() == reflect.Pointer {
			elementValue = elementValue.Elem()
		}
		slog.Info("Wrap", "elementValue", elementValue)
		if elementValue.Kind() != reflect.Struct {
			return
		}
		*/

		methodValue := elementValue.MethodByName("DataSet")
		if methodValue.IsZero() {
			slog.WarnContext(context.TODO(), "UseData: Wrap: Method 'DataSet' not found", "element", element)
			return
		}
		//slog.Debug("Wrap", "methodValue", methodValue)
		methodType := methodValue.Type()
		if methodType.NumIn() != 2 {
			slog.WarnContext(context.TODO(), "UseData: Wrap: Method 'DataSet' has wrong number of arguments", "element", element, "methodType", methodType)
			return
		}
		in1 := methodType.In(0)
		if in1.Kind() != reflect.String {
			slog.WarnContext(context.TODO(), "UseData: Wrap: Method 'DataSet' has wrong type for first argument", "element", element, "in1", in1.String())
			return
		}
		in2 := methodType.In(1)
		if in2 != reflect.TypeFor[any]() {
			slog.WarnContext(context.TODO(), "UseData: Wrap: Method 'DataSet' has wrong type for second argument", "element", element, "in2", in2.String())
			return
		}

		for name, value := range c.Data {
			reflectOptions := []reflect.Value{
				reflect.ValueOf(name),
				reflect.ValueOf(value),
			}
			methodValue.Call(reflectOptions)
		}
	}()
	return element
}
