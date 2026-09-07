package demo

import (
	"context"
	"log/slog"
	"time"

	"github.com/go-app-blazar/blazar/blazar"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type FormPage struct {
	app.Compo

	name string
}

var _ app.Initializer = (*FormPage)(nil)

func (c *FormPage) OnInit() {
	slog.DebugContext(context.TODO(), "FormPage: OnInit")

	c.name = "Monkey D. Luffy"
}

func (c *FormPage) OnNav(ctx app.Context) {
	slog.DebugContext(ctx.Context, "FormPage: OnNav")
}

func (c *FormPage) Render() app.UI {
	cancelFunction := func(ctx app.Context) {
		app.Window().Call("alert", "Cancel called")
	}
	submitFunction := func(ctx app.Context) {
		app.Window().Call("alert", "Submit called")
	}

	action1Function := func(ctx app.Context) {
		app.Window().Call("alert", "Action 1 called")
	}
	action2Function := func(ctx app.Context) {
		app.Window().Call("alert", "Action 2 called")
	}
	sleepFunction := func(ctx app.Context) {
		time.Sleep(5 * time.Second)
	}

	return blazar.Page().
		Body(
			app.FieldSet().
				Body(
					app.Legend().Text("Default"),
					blazar.Form().
						Body(
							blazar.Input[string]().
								Label("Name").
								Bind(&c.name),
						),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("With cancel"),
					blazar.Form().
						Body(
							blazar.Input[string]().
								Label("Name").
								Bind(&c.name),
						).
						Action(
							blazar.FormAction{
								Function: cancelFunction,
								Cancel:   true,
							},
						),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("With submit"),
					blazar.Form().
						Body(
							blazar.Input[string]().
								Label("Name").
								Bind(&c.name),
						).
						Action(
							blazar.FormAction{
								Function: submitFunction,
								Submit:   true,
							},
						),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("With both submit and cancel"),
					blazar.Form().
						Body(
							blazar.Input[string]().
								Label("Name").
								Bind(&c.name),
						).
						Action(
							blazar.FormAction{
								Function: cancelFunction,
								Cancel:   true,
							},
							blazar.FormAction{
								Function: submitFunction,
								Submit:   true,
							},
						),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("With custom labels"),
					blazar.Form().
						Body(
							blazar.Input[string]().
								Label("Name").
								Bind(&c.name),
						).
						Action(
							blazar.FormAction{
								Name:     "Please cancel this",
								Icon:     "trash",
								Function: cancelFunction,
								Cancel:   true,
							},
							blazar.FormAction{
								Name:     "Please submit this",
								Icon:     "save",
								Function: submitFunction,
								Submit:   true,
							},
						),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("With simple custom actions only"),
					blazar.Form().
						Body(
							blazar.Input[string]().
								Label("Name").
								Bind(&c.name),
						).
						Action(
							blazar.FormAction{Name: "Action 1", Icon: "person", Function: action1Function},
							blazar.FormAction{Name: "Action 2", Icon: "gear", Function: action2Function},
							blazar.FormAction{Name: "Sleep", Icon: "clock", Function: sleepFunction},
						),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("With simple custom actions only and no auto submit"),
					blazar.Form().
						AutoSubmit(false).
						Body(
							blazar.Input[string]().
								Label("Name").
								Bind(&c.name),
						).
						Action(
							blazar.FormAction{Name: "Action 1", Icon: "person", Function: action1Function},
							blazar.FormAction{Name: "Action 2", Icon: "gear", Function: action2Function},
							blazar.FormAction{Name: "Sleep", Icon: "clock", Function: sleepFunction},
						),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("With custom actions only"),
					blazar.Form().
						Body(
							blazar.Input[string]().
								Label("Name").
								Bind(&c.name),
						).
						Action(
							blazar.FormAction{Name: "Action 1", Icon: "person", Function: action1Function, Flat: true},
							blazar.FormAction{Name: "Action 2", Icon: "gear", Function: action2Function, Color: "red"},
							blazar.FormAction{Name: "Sleep", Icon: "clock", Function: sleepFunction, Color: "black"},
						),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("With custom actions and default actions"),
					blazar.Form().
						Body(
							blazar.Input[string]().
								Label("Name").
								Bind(&c.name),
						).
						Action(
							blazar.FormAction{
								Name:     "Please cancel this",
								Icon:     "trash",
								Function: cancelFunction,
								Cancel:   true,
							},
							blazar.FormAction{
								Name:     "Please submit this",
								Icon:     "save",
								Function: submitFunction,
								Submit:   true,
							},
							blazar.FormAction{Name: "Action 1", Icon: "person", Function: action1Function},
							blazar.FormAction{Name: "Action 2", Icon: "gear", Function: action2Function},
							blazar.FormAction{Name: "Sleep", Icon: "clock", Function: sleepFunction},
						),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("With styled custom actions and default actions"),
					blazar.Form().
						Body(
							blazar.Input[string]().
								Label("Name").
								Bind(&c.name),
						).
						Action(
							blazar.FormAction{
								Name:     "Please cancel this",
								Icon:     "trash",
								Function: cancelFunction,
								Cancel:   true,
							},
							blazar.FormAction{
								Name:     "Please submit this",
								Icon:     "save",
								Function: submitFunction,
								Submit:   true,
							},
							blazar.FormAction{Name: "Action 1", Icon: "person", Function: action1Function, Flat: true},
							blazar.FormAction{Name: "Action 2", Icon: "gear", Function: action2Function, Color: "red"},
							blazar.FormAction{Name: "Sleep", Icon: "clock", Function: sleepFunction, Color: "black"},
						),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("With outlined custom actions and default actions"),
					blazar.Form().
						Body(
							blazar.Input[string]().
								Label("Name").
								Bind(&c.name),
						).
						Action(
							blazar.FormAction{
								Name:     "Please cancel this",
								Icon:     "trash",
								Function: cancelFunction,
								Cancel:   true,
							},
							blazar.FormAction{
								Name:     "Please submit this",
								Icon:     "save",
								Function: submitFunction,
								Submit:   true,
							},
							blazar.FormAction{Name: "Action 1", Icon: "person", Function: action1Function, Outline: true},
							blazar.FormAction{Name: "Action 2", Icon: "gear", Function: action2Function, Outline: true, Color: "red"},
							blazar.FormAction{Name: "Sleep", Icon: "clock", Function: sleepFunction, Outline: true, Color: "black"},
						),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("With actions to the left"),
					blazar.Form().
						Spacer(false).
						Body(
							blazar.Input[string]().
								Label("Name").
								Bind(&c.name),
						).
						Action(
							blazar.FormAction{
								Name:     "Please submit this",
								Icon:     "save",
								Function: submitFunction,
								Submit:   true,
							},
							blazar.FormAction{Name: "Action 1", Icon: "person", Function: action1Function, Flat: true},
							blazar.FormAction{Name: "Action 2", Icon: "gear", Function: action2Function, Color: "red"},
							blazar.FormAction{Name: "Sleep", Icon: "clock", Function: sleepFunction, Color: "black"},
						),
				),
		)
}
