package blazar

import (
	"context"
	"log/slog"
	"reflect"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type MixHTML struct {
	IAttributes map[string]any
	IClasses    []string
	IData       map[string]any
	IStyles     map[string]string
	ITitle      string
}

func (c *MixHTML) mixAttr(name string, value any) *MixHTML {
	if c.IAttributes == nil {
		c.IAttributes = map[string]any{}
	}
	c.IAttributes[name] = value
	return c
}

func (c *MixHTML) mixClass(class ...string) *MixHTML {
	c.IClasses = append(c.IClasses, class...)
	return c
}

func (c *MixHTML) mixDataSet(name string, value any) *MixHTML {
	if c.IData == nil {
		c.IData = map[string]any{}
	}
	c.IData[name] = value
	return c
}

func (c *MixHTML) mixStyle(name, value string) *MixHTML {
	if c.IStyles == nil {
		c.IStyles = map[string]string{}
	}
	c.IStyles[name] = value
	return c
}

func (c *MixHTML) mixTitle(title string) *MixHTML {
	c.ITitle = title
	return c
}

func (c *MixHTML) Wrap(element app.UI) app.UI {
	//slog.Debug("Wrap", "element", element)

	elementValue := reflect.ValueOf(element)
	//slog.Debug("Wrap", "elementValue", elementValue)
	if elementValue.IsZero() {
		slog.WarnContext(context.TODO(), "mixHTML: Wrap: Element is zero", "element", element)
		return element
	}

	// TODO: Can we remove this block?
	/*for elementValue.Kind() == reflect.Pointer {
		elementValue = elementValue.Elem()
	}
	slog.Info("Wrap", "elementValue", elementValue)
	if elementValue.Kind() != reflect.Struct {
		return
	}
	*/

	// Attr
	func() {
		methodValue := elementValue.MethodByName("Attr")
		if !methodValue.IsValid() {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'Attr' is not valid", "element", element)
			return
		}
		if methodValue.IsZero() {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'Attr' not found", "element", element)
			return
		}
		//slog.Debug("Wrap", "methodValue", methodValue)
		methodType := methodValue.Type()
		if methodType.NumIn() != 2 {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'Attr' has wrong number of arguments", "element", element, "methodType", methodType)
			return
		}
		in1 := methodType.In(0)
		if in1.Kind() != reflect.String {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'Attr' has wrong type for first argument", "element", element, "in1", in1.String())
			return
		}
		in2 := methodType.In(1)
		if in2 != reflect.TypeFor[any]() {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'Attr' has wrong type for second argument", "element", element, "in2", in2.String())
			return
		}

		for name, value := range c.IAttributes {
			reflectOptions := []reflect.Value{
				reflect.ValueOf(name),
				reflect.ValueOf(value),
			}
			methodValue.Call(reflectOptions)
		}
	}()

	// Class
	func() {
		methodValue := elementValue.MethodByName("Class")
		if !methodValue.IsValid() {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'Class' is not valid", "element", element)
			return
		}
		if methodValue.IsZero() {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'Class' not found", "element", element)
			return
		}
		//slog.Debug("Wrap", "methodValue", methodValue)
		methodType := methodValue.Type()
		if methodType.NumIn() != 1 {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'Class' has wrong number of arguments", "element", element, "methodType", methodType)
			return
		}
		in1 := methodType.In(0)
		if in1 != reflect.TypeFor[[]string]() {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'Class' has wrong type for first argument", "element", element, "in1", in1.String())
			return
		}

		reflectOptions := []reflect.Value{}
		for _, class := range c.IClasses {
			reflectOptions = append(reflectOptions, reflect.ValueOf(class))
		}
		methodValue.Call(reflectOptions)
	}()

	// DataSet
	func() {
		methodValue := elementValue.MethodByName("DataSet")
		if !methodValue.IsValid() {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'DataSet' is not valid", "element", element)
			return
		}
		if methodValue.IsZero() {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'DataSet' not found", "element", element)
			return
		}
		//slog.Debug("Wrap", "methodValue", methodValue)
		methodType := methodValue.Type()
		if methodType.NumIn() != 2 {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'DataSet' has wrong number of arguments", "element", element, "methodType", methodType)
			return
		}
		in1 := methodType.In(0)
		if in1.Kind() != reflect.String {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'DataSet' has wrong type for first argument", "element", element, "in1", in1.String())
			return
		}
		in2 := methodType.In(1)
		if in2 != reflect.TypeFor[any]() {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'DataSet' has wrong type for second argument", "element", element, "in2", in2.String())
			return
		}

		for name, value := range c.IData {
			reflectOptions := []reflect.Value{
				reflect.ValueOf(name),
				reflect.ValueOf(value),
			}
			methodValue.Call(reflectOptions)
		}
	}()

	// Style
	func() {
		methodValue := elementValue.MethodByName("Style")
		if !methodValue.IsValid() {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'Style' is not valid", "element", element)
			return
		}
		if methodValue.IsZero() {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'Style' not found", "element", element)
			return
		}
		//slog.Debug("Wrap", "methodValue", methodValue)
		methodType := methodValue.Type()
		if methodType.NumIn() != 2 {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'Style' has wrong number of arguments", "element", element, "methodType", methodType)
			return
		}
		in1 := methodType.In(0)
		if in1.Kind() != reflect.String {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'Style' has wrong type for first argument", "element", element, "in1", in1.String())
			return
		}
		in2 := methodType.In(1)
		if in2.Kind() != reflect.String {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'Style' has wrong type for second argument", "element", element, "in2", in2.String())
			return
		}

		for name, value := range c.IStyles {
			reflectOptions := []reflect.Value{
				reflect.ValueOf(name),
				reflect.ValueOf(value),
			}
			methodValue.Call(reflectOptions)
		}
	}()

	// Title
	func() {
		methodValue := elementValue.MethodByName("Title")
		if !methodValue.IsValid() {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'Title' is not valid", "element", element)
			return
		}
		if methodValue.IsZero() {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'Title' not found", "element", element)
			return
		}
		//slog.Debug("Wrap", "methodValue", methodValue)
		methodType := methodValue.Type()
		if methodType.NumIn() != 1 {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'Title' has wrong number of arguments", "element", element, "methodType", methodType)
			return
		}
		in1 := methodType.In(0)
		if in1 != reflect.TypeFor[string]() {
			slog.WarnContext(context.TODO(), "mixHTML: Wrap: Method 'Title' has wrong type for first argument", "element", element, "in1", in1.String())
			return
		}

		reflectOptions := []reflect.Value{
			reflect.ValueOf(c.ITitle),
		}
		methodValue.Call(reflectOptions)
	}()

	return element
}
