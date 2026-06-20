package demo

import (
	"log/slog"

	"github.com/go-app-blazar/blazar/blazar"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type AppBarPage struct {
	app.Compo

	defaultOpen   bool
	defaultClosed bool
}

func (c *AppBarPage) OnMount(ctx app.Context) {
	slog.InfoContext(ctx.Context, "AppBarPage: OnMount")

	c.defaultOpen = true
	c.defaultClosed = false
}

func (c *AppBarPage) OnNav(ctx app.Context) {
	slog.InfoContext(ctx.Context, "AppBarPage: OnNav")
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
				Trailer(app.Div().Text("Trailer")),
			blazar.AppBar().
				Icon("bars").
				IconFunction(func(ctx app.Context, e app.Event) {
					app.Window().Call("alert", "Icon clicked")
				}).
				HeadlineText("App Bar with default icon, subtitle, and trailer").
				SubtitleText("Subtitle text").
				Trailer(app.Div().Text("Trailer")),
		)
}
