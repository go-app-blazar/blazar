package demo

import (
	"fmt"
	"log/slog"

	"github.com/go-app-blazar/blazar/blazar"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type InputPage struct {
	app.Compo

	disabled  bool
	clearable bool
	prefix    string
	suffix    string

	stringValue   string
	intValue      int
	floatValue    float64
	uintValue     uint
	boolValue     bool
	checkboxValue bool
}

func (c *InputPage) OnMount(ctx app.Context) {
	slog.DebugContext(ctx.Context, "InputPage: OnMount")

	c.disabled = false
	c.clearable = false

	c.stringValue = "Hello, World!"
	c.intValue = 123
	c.floatValue = 123.456
	c.uintValue = 123
	c.boolValue = true
	c.checkboxValue = true
}

func (c *InputPage) OnNav(ctx app.Context) {
	slog.DebugContext(ctx.Context, "InputPage: OnNav")
}

func (c *InputPage) Render() app.UI {
	type Row struct {
		Name  string
		Value string
	}

	rows := []Row{
		{
			Name:  "string",
			Value: c.stringValue,
		},
		{
			Name:  "int",
			Value: fmt.Sprintf("%d", c.intValue),
		},
		{
			Name:  "float",
			Value: fmt.Sprintf("%f", c.floatValue),
		},
		{
			Name:  "uint",
			Value: fmt.Sprintf("%d", c.uintValue),
		},
		{
			Name:  "bool",
			Value: fmt.Sprintf("%t", c.boolValue),
		},
		{
			Name:  "checkbox",
			Value: fmt.Sprintf("%t", c.checkboxValue),
		},
	}

	return blazar.Page().
		Body(
			app.FieldSet().
				Body(
					app.Legend().Text("Config"),
					blazar.Input[bool]().
						Label("Disabled").
						Bind(&c.disabled),
					blazar.Input[bool]().
						Label("Clearable").
						Bind(&c.clearable),
					blazar.Input[string]().
						Label("Prefix").
						Bind(&c.prefix),
					blazar.Input[string]().
						Label("Suffix").
						Bind(&c.suffix),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("Input"),
					blazar.Input[string]().
						Label("string").
						Prefix(c.prefix).
						Suffix(c.suffix).
						Disabled(c.disabled).
						Clearable(c.clearable).
						Bind(&c.stringValue),
					blazar.Input[int]().
						Label("int").
						Prefix(c.prefix).
						Suffix(c.suffix).
						Disabled(c.disabled).
						Clearable(c.clearable).
						Bind(&c.intValue),
					blazar.Input[float64]().
						Label("float").
						Prefix(c.prefix).
						Suffix(c.suffix).
						Disabled(c.disabled).
						Clearable(c.clearable).
						Bind(&c.floatValue),
					blazar.Input[uint]().
						Label("uint").
						Prefix(c.prefix).
						Suffix(c.suffix).
						Disabled(c.disabled).
						Clearable(c.clearable).
						Bind(&c.uintValue),
					blazar.Input[bool]().
						Label("bool").
						Prefix(c.prefix).
						Suffix(c.suffix).
						Disabled(c.disabled).
						Clearable(c.clearable).
						Bind(&c.boolValue),
					blazar.Checkbox().
						Label("checkbox").
						Disabled(c.disabled).
						Bind(&c.checkboxValue),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("Output"),
					blazar.Table[Row]().
						Interactive(false).
						Rows(rows).
						Columns([]blazar.TableColumn[Row]{
							{
								Name: "Name",
								Type: blazar.TableColumnTypeString,
								Value: func(row Row) any {
									return row.Name
								},
							},
							{
								Name: "Value",
								Type: blazar.TableColumnTypeString,
								Value: func(row Row) any {
									return row.Value
								},
							},
						}),
				),
		)
}
