package htmlevent

import (
	"fmt"
	"reflect"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

// MustParse parses the app.Event into the target struct and panics if there is an error.
func MustParse(appEvent app.Event, target any) {
	err := Parse(appEvent, target)
	if err != nil {
		panic(fmt.Errorf("htmlevent: MustParse: %w", err))
	}
}

// Parse the app.Event into the target struct.
func Parse(appEvent app.Event, target any) error {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Pointer {
		return fmt.Errorf("target must be a pointer: %s", targetValue.Kind())
	}
	targetValue = targetValue.Elem()
	if targetValue.Kind() != reflect.Struct {
		return fmt.Errorf("target must be a pointer to a struct: %s", targetValue.Kind())
	}

	// This is the type of the app.Value interface.
	appValueType := reflect.TypeOf((*app.Value)(nil)).Elem()

	targetType := targetValue.Type()
	for _, field := range reflect.VisibleFields(targetType) {
		javascriptTag := field.Tag.Get("javascript")
		if javascriptTag == "" {
			continue
		}

		javascriptValue := appEvent.Get(javascriptTag)

		fieldValue := targetValue.FieldByName(field.Name)
		if !fieldValue.IsValid() {
			continue
		}
		if !fieldValue.CanSet() {
			continue
		}

		// If the field is an app.Value, we can just set it directly.
		if fieldValue.Type().Implements(appValueType) {
			fieldValue.Set(reflect.ValueOf(javascriptValue))
			continue
		}

		// If the event value is null, we can just skip it.
		if javascriptValue.IsNull() {
			continue
		}

		switch fieldValue.Kind() {
		case reflect.String:
			if javascriptValue.Type() != app.TypeString {
				return fmt.Errorf("field %s is a string, but the event value is not a string: %s", field.Name, javascriptValue.Type())
			}
			fieldValue.SetString(javascriptValue.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if javascriptValue.Type() != app.TypeNumber {
				return fmt.Errorf("field %s is a number, but the event value is not a number: %s", field.Name, javascriptValue.Type())
			}
			fieldValue.SetInt(int64(javascriptValue.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if javascriptValue.Type() != app.TypeNumber {
				return fmt.Errorf("field %s is a number, but the event value is not a number: %s", field.Name, javascriptValue.Type())
			}
			fieldValue.SetUint(uint64(javascriptValue.Int()))
		case reflect.Float32, reflect.Float64:
			if javascriptValue.Type() != app.TypeNumber {
				return fmt.Errorf("field %s is a number, but the event value is not a number: %s", field.Name, javascriptValue.Type())
			}
			fieldValue.SetFloat(javascriptValue.Float())
		case reflect.Bool:
			if javascriptValue.Type() != app.TypeBoolean {
				return fmt.Errorf("field %s is a boolean, but the event value is not a boolean: %s", field.Name, javascriptValue.Type())
			}
			fieldValue.SetBool(javascriptValue.Bool())
		case reflect.Struct:
			return fmt.Errorf("struct fields are not supported")
		case reflect.Slice:
			return fmt.Errorf("slice fields are not supported")
		case reflect.Map:
			return fmt.Errorf("map fields are not supported")
		}
	}
	return nil
}
