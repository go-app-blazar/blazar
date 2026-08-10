package demo

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-app-blazar/blazar/blazar"
	"github.com/go-app-blazar/blazar/htmlevent"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type CopyPage struct {
	app.Compo

	output string
}

var _ app.Initializer = (*CopyPage)(nil)

func (c *CopyPage) OnInit() {
	slog.DebugContext(context.TODO(), "CopyPage: OnInit")
}

func (c *CopyPage) OnNav(ctx app.Context) {
	slog.DebugContext(ctx.Context, "CopyPage: OnNav")
}

func (c *CopyPage) onClick(ctx app.Context, event htmlevent.PointerEvent) {
	slog.InfoContext(ctx.Context, fmt.Sprintf("CopyPage: onClick: event: %+v", event))
}

func (c *CopyPage) onCopy(ctx app.Context, value string) {
	app.Window().Get("navigator").Get("clipboard").Call("readText").Then(func(result app.Value) {
		c.output = result.String()
	})
}

func (c *CopyPage) Render() app.UI {
	return blazar.Page().
		Body(
			app.FieldSet().
				Body(
					app.Legend().Text("Copy"),
					blazar.Copy().
						Text(`Copy "Hello, World!" to clipboard`).
						Value("Hello, World!").
						OnCopy(c.onCopy).
						OnClick(c.onClick),
					blazar.Copy().
						Text(`Copy "Goodbye, World!" to clipboard (disabled)`).
						Value("Goodbye, World!").
						Disabled(true).
						OnCopy(c.onCopy).
						OnClick(c.onClick),
					blazar.Copy().
						Icon("").
						Text(`This one has no icon`).
						Value("Hidden icon copy").
						OnCopy(c.onCopy).
						OnClick(c.onClick),
					app.Span().
						Body(
							app.Text("Only the "),
							blazar.Copy().
								Style("display", "inline-block").
								Text("middle").
								Value("Magical middle text").
								OnCopy(c.onCopy).
								OnClick(c.onClick),
							app.Text(" can be copied."),
						),
					blazar.Copy().
						Value("Copied from one with a body").
						Body(
							app.Text("This one has a body with a button."),
							blazar.Button().
								Label("Button"),
						).
						OnCopy(c.onCopy).
						OnClick(c.onClick),
					blazar.Copy().
						Text(`Copy the empty string`).
						Value("").
						OnCopy(c.onCopy).
						OnClick(c.onClick),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("Output"),
					blazar.Input[string]().
						Label("Output").
						Value(c.output).
						Disabled(true),
				),
		)
}
