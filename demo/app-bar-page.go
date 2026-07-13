package demo

import (
	"context"
	"log/slog"

	"github.com/go-app-blazar/blazar/blazar"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type AppBarPage struct {
	app.Compo

	defaultOpen   bool
	defaultClosed bool
}

var _ app.Initializer = (*AppBarPage)(nil)

func (c *AppBarPage) OnInit() {
	slog.DebugContext(context.TODO(), "AppBarPage: OnInit")

	c.defaultOpen = true
	c.defaultClosed = false
}

func (c *AppBarPage) OnNav(ctx app.Context) {
	slog.DebugContext(ctx.Context, "AppBarPage: OnNav")
}

func (c *AppBarPage) Render() app.UI {
	return blazar.Page().
		Body(
			blazar.AppBar(),
			blazar.AppBar().
				NoIcon(true).
				HeadlineText("App Bar with no icon"),
			blazar.AppBar().
				Icon("bars").
				IconFunction(func(ctx app.Context, e app.Event) {
					app.Window().Call("alert", "Icon clicked")
				}).
				HeadlineText("App Bar with default icon"),
			blazar.AppBar().
				Icon("bars").
				IconFunction(func(ctx app.Context, e app.Event) {
					app.Window().Call("alert", "Icon clicked")
				}).
				HeadlineText("App Bar with default icon and trailer").
				TrailerFunction(func() app.UI {
					return app.Div().Text("Trailer")
				}),
			blazar.AppBar().
				Icon("bars").
				IconFunction(func(ctx app.Context, e app.Event) {
					app.Window().Call("alert", "Icon clicked")
				}).
				HeadlineText("App Bar with default icon, subtitle, and trailer").
				SubtitleText("Subtitle text").
				TrailerFunction(func() app.UI {
					return app.Div().Text("Trailer")
				}),
		)
}
