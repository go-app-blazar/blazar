package demo

import (
	"fmt"
	"log/slog"

	"github.com/go-app-blazar/blazar/blazar"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type InputPage struct {
	app.Compo

	stringValue   string
	intValue      int
	floatValue    float64
	uintValue     uint
	boolValue     bool
	checkboxValue bool
}

func (c *InputPage) OnMount(ctx app.Context) {
	slog.InfoContext(ctx.Context, "InputPage: OnMount")

	c.stringValue = "Hello, World!"
	c.intValue = 123
	c.floatValue = 123.456
	c.uintValue = 123
	c.boolValue = true
	c.checkboxValue = true
}

func (c *InputPage) OnNav(ctx app.Context) {
	slog.InfoContext(ctx.Context, "InputPage: OnNav")
}

func (c *InputPage) Render() app.UI {
	return blazar.Page().
		Body(
			app.FieldSet().
				Body(
					app.Legend().Text("Input"),
					blazar.Input[string]().
						Label("string").
						Bind(&c.stringValue),
					blazar.Input[int]().
						Label("int").
						Bind(&c.intValue),
					blazar.Input[float64]().
						Label("float").
						Bind(&c.floatValue),
					blazar.Input[uint]().
						Label("uint").
						Bind(&c.uintValue),
					blazar.Input[bool]().
						Label("bool").
						Bind(&c.boolValue),
					blazar.Checkbox().
						Label("checkbox").
						Bind(&c.checkboxValue),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("Output"),
					app.Div().Text("string"),
					app.Pre().Text(c.stringValue),
					app.Div().Text("int"),
					app.Pre().Text(fmt.Sprintf("%d", c.intValue)),
					app.Div().Text("float"),
					app.Pre().Text(fmt.Sprintf("%f", c.floatValue)),
					app.Div().Text("uint"),
					app.Pre().Text(fmt.Sprintf("%d", c.uintValue)),
					app.Div().Text("bool"),
					app.Pre().Text(fmt.Sprintf("%t", c.boolValue)),
					app.Div().Text("checkbox"),
					app.Pre().Text(fmt.Sprintf("%t", c.checkboxValue)),
				),
		)
}
