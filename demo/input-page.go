package demo

import (
	"context"
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
	outline   bool

	intMin  int
	intMax  int
	intStep int

	floatMin  float64
	floatMax  float64
	floatStep float64

	uintMin  uint
	uintMax  uint
	uintStep uint

	stringValue   string
	intValue      int
	floatValue    float64
	uintValue     uint
	boolValue     bool
	checkboxValue bool
}

var _ app.Initializer = (*InputPage)(nil)

func (c *InputPage) OnInit() {
	slog.DebugContext(context.TODO(), "InputPage: OnInit")

	c.disabled = false
	c.clearable = false
	c.prefix = ""
	c.suffix = ""
	c.outline = true

	c.intMin = 0
	c.intMax = 1000
	c.intStep = 1

	c.uintMin = 0
	c.uintMax = 1000
	c.uintStep = 1

	c.floatMin = 0.0
	c.floatMax = 1000.0
	c.floatStep = 0.1

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
					app.Div().
						Style("display", "flex").
						Style("flex-direction", "row").
						Style("gap", "1em").
						Body(
							blazar.Input[bool]().
								Label("Disabled").
								Bind(&c.disabled),
							blazar.Input[bool]().
								Label("Clearable").
								Bind(&c.clearable),
						),
					app.Div().
						Style("display", "flex").
						Style("flex-direction", "row").
						Style("gap", "1em").
						Body(
							blazar.Input[string]().
								Label("Prefix").
								Bind(&c.prefix),
							blazar.Input[string]().
								Label("Suffix").
								Bind(&c.suffix),
						),
					app.Div().
						Style("display", "flex").
						Style("flex-direction", "row").
						Style("gap", "1em").
						Body(
							blazar.Input[bool]().
								Label("Outline").
								Bind(&c.outline),
						),
					app.Div().
						Style("display", "flex").
						Style("flex-direction", "row").
						Style("gap", "1em").
						Body(
							blazar.Input[int]().
								Label("Int Min").
								Bind(&c.intMin),
							blazar.Input[int]().
								Label("Int Max").
								Bind(&c.intMax),
							blazar.Input[int]().
								Label("Int Step").
								Bind(&c.intStep),
						),
					app.Div().
						Style("display", "flex").
						Style("flex-direction", "row").
						Style("gap", "1em").
						Body(
							blazar.Input[uint]().
								Label("Uint Min").
								Bind(&c.uintMin),
							blazar.Input[uint]().
								Label("Uint Max").
								Bind(&c.uintMax),
							blazar.Input[uint]().
								Label("Uint Step").
								Bind(&c.uintStep),
						),
					app.Div().
						Style("display", "flex").
						Style("flex-direction", "row").
						Style("gap", "1em").
						Body(
							blazar.Input[float64]().
								Label("Float Min").
								Bind(&c.floatMin),
							blazar.Input[float64]().
								Label("Float Max").
								Bind(&c.floatMax),
							blazar.Input[float64]().
								Label("Float Step").
								Step(0.1).
								Bind(&c.floatStep),
						),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("Input"),
					blazar.Input[string]().
						Label("string").
						Outline(c.outline).
						Prefix(c.prefix).
						Suffix(c.suffix).
						Disabled(c.disabled).
						Clearable(c.clearable).
						Bind(&c.stringValue),
					blazar.Input[int]().
						Label("int").
						Outline(c.outline).
						Prefix(c.prefix).
						Suffix(c.suffix).
						Disabled(c.disabled).
						Clearable(c.clearable).
						Bind(&c.intValue),
					blazar.Input[int]().
						Label("int (bounded)").
						Outline(c.outline).
						Prefix(c.prefix).
						Suffix(c.suffix).
						Disabled(c.disabled).
						Clearable(c.clearable).
						Min(c.intMin).
						Max(c.intMax).
						Step(c.intStep).
						Bind(&c.intValue),
					blazar.Input[uint]().
						Label("uint").
						Outline(c.outline).
						Prefix(c.prefix).
						Suffix(c.suffix).
						Disabled(c.disabled).
						Clearable(c.clearable).
						Bind(&c.uintValue),
					blazar.Input[uint]().
						Label("uint (bounded)").
						Outline(c.outline).
						Prefix(c.prefix).
						Suffix(c.suffix).
						Disabled(c.disabled).
						Clearable(c.clearable).
						Min(c.uintMin).
						Max(c.uintMax).
						Step(c.uintStep).
						Bind(&c.uintValue),
					blazar.Input[float64]().
						Label("float").
						Outline(c.outline).
						Prefix(c.prefix).
						Suffix(c.suffix).
						Disabled(c.disabled).
						Clearable(c.clearable).
						Bind(&c.floatValue),
					blazar.Input[float64]().
						Label("float (bounded)").
						Outline(c.outline).
						Prefix(c.prefix).
						Suffix(c.suffix).
						Disabled(c.disabled).
						Clearable(c.clearable).
						Min(c.floatMin).
						Max(c.floatMax).
						Step(c.floatStep).
						Bind(&c.floatValue),
					blazar.Input[bool]().
						Label("bool").
						Outline(c.outline).
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
