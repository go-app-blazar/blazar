package blazar

import (
	"context"
	"log/slog"

	"github.com/go-app-blazar/blazar/matchmedia"
	"github.com/go-app-blazar/router"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type blazarMainLayout struct {
	app.Compo
	router.RouterViewComponent

	IHeadline        string
	IHeadlineUI      app.UI
	ISubtitle        string
	ISubtitleUI      app.UI
	ITrailer         app.UI
	IDrawer          app.UI
	IResponsiveWidth string

	matchMedia    *matchmedia.MatchMedia
	narrow        bool
	drawerVisible bool
}

var _ router.RouterViewInterface = (*blazarMainLayout)(nil)
var _ app.Mounter = (*blazarMainLayout)(nil)

func MainLayout() *blazarMainLayout {
	return &blazarMainLayout{
		IResponsiveWidth: "900px",
	}
}

func (c *blazarMainLayout) ResponsiveWidth(width string) *blazarMainLayout {
	c.IResponsiveWidth = width
	if c.matchMedia != nil {
		c.matchMedia.SetQuery("screen and (max-width: " + c.IResponsiveWidth + ")")
	}
	return c
}

func (c *blazarMainLayout) HeadlineText(text string) *blazarMainLayout {
	c.IHeadline = text
	return c
}

func (c *blazarMainLayout) Headline(ui app.UI) *blazarMainLayout {
	c.IHeadlineUI = ui
	return c
}

func (c *blazarMainLayout) SubtitleText(text string) *blazarMainLayout {
	c.ISubtitle = text
	return c
}

func (c *blazarMainLayout) Subtitle(ui app.UI) *blazarMainLayout {
	c.ISubtitleUI = ui
	return c
}

func (c *blazarMainLayout) Trailer(ui app.UI) *blazarMainLayout {
	c.ITrailer = ui
	return c
}

func (c *blazarMainLayout) Drawer(drawer app.UI) *blazarMainLayout {
	c.IDrawer = drawer
	return c
}

func (c *blazarMainLayout) OnMount(ctx app.Context) {
	slog.InfoContext(ctx.Context, "MainLayout: OnMount", "responsiveWidth", c.IResponsiveWidth)

	c.matchMedia = matchmedia.New(ctx, "screen and (max-width: "+c.IResponsiveWidth+")")
	c.matchMedia.SetOnChange(func(ctx app.Context, value bool) {
		slog.InfoContext(ctx.Context, "MediaPage: MatchMedia: OnChange", "value", value)
		c.narrow = value

		ctx.Update()
	})
}

func (c *blazarMainLayout) toggleDrawer() {
	slog.InfoContext(context.TODO(), "MainLayout: toggleDrawer")

	c.drawerVisible = !c.drawerVisible
}

func (c *blazarMainLayout) Render() app.UI {
	slog.InfoContext(context.TODO(), "MainLayout: Render")

	iconVisible := false
	drawerVisibleClass := "visible"
	if c.narrow {
		iconVisible = true
		drawerVisibleClass = ""
		if c.drawerVisible {
			drawerVisibleClass = "visible"
		}
	}

	return app.Div().
		Class("blazar-main-layout").
		Body(
			AppBar().
				Class("blazar-main-layout__header").
				NoIcon(!iconVisible).
				Icon("bars").
				IconFunction(func(ctx app.Context, e app.Event) {
					c.toggleDrawer()
				}).
				HeadlineText(c.IHeadline).
				Headline(c.IHeadlineUI).
				SubtitleText(c.ISubtitle).
				Subtitle(c.ISubtitleUI).
				Trailer(c.ITrailer),
			app.Div().
				Class("blazar-main-layout__body").
				Body(
					app.If(c.IDrawer != nil, func() app.UI {
						return app.Div().
							Class("blazar-main-layout__drawer", drawerVisibleClass).
							Body(c.IDrawer)
					}),
					app.Div().
						Class("blazar-main-layout__content").
						Body(
							app.If(c.RouterViewComponent.RouterView() != nil, func() app.UI {
								return c.RouterViewComponent.RouterView()
							}),
						),
				),
		)
}
