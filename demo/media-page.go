package demo

import (
	"context"
	"log/slog"

	"github.com/go-app-blazar/blazar/blazar"
	"github.com/go-app-blazar/blazar/matchmedia"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type MediaPage struct {
	app.Compo

	mediaQuery1 string
	result1     bool

	matchMedia *matchmedia.MatchMedia
}

func (c *MediaPage) OnMount(ctx app.Context) {
	slog.InfoContext(ctx.Context, "MediaPage: OnMount")
	c.mediaQuery1 = "screen and (max-width: 900px)"
	c.matchMedia = matchmedia.New(c.mediaQuery1)
	c.matchMedia.OnChange(func(ctx app.Context, value bool) {
		slog.InfoContext(ctx.Context, "MediaPage: MatchMedia: OnChange", "value", value)
	})
}

func (c *MediaPage) OnNav(ctx app.Context) {
	slog.InfoContext(ctx.Context, "MediaPage: OnNav")
}

func (c *MediaPage) OnUpdate(ctx app.Context) {
	slog.InfoContext(ctx.Context, "MediaPage: OnUpdate", "mediaQuery1", c.mediaQuery1, "result1", c.result1)
	c.matchMedia.SetQuery(c.mediaQuery1)
}

func (c *MediaPage) Render() app.UI {
	slog.InfoContext(context.TODO(), "MediaPage: Render", "mediaQuery1", c.mediaQuery1, "result1", c.result1)

	return blazar.Page().
		Body(
			app.FieldSet().
				Body(
					app.Legend().Text("Media Query 1"),
					blazar.Input[string]().
						Label("string").
						Bind(&c.mediaQuery1).
						On("change", func(ctx app.Context, e app.Event) {
							slog.InfoContext(ctx.Context, "MediaPage: OnChange", "mediaQuery1", c.mediaQuery1)
						}),
				),
			app.FieldSet().
				Body(
					app.Legend().Text("Result"),
					app.Div().Text("Media Query 1"),
					app.Pre().Text(c.result1),
				),
		)
}
