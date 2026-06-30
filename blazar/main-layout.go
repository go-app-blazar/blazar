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

	IHeadline         string
	IHeadlineFunction func() app.UI
	ISubtitle         string
	ISubtitleFunction func() app.UI
	ITrailerFunction  func() app.UI
	IDrawerFunction   func() app.UI
	IResponsiveWidth  string

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

func (c *blazarMainLayout) HeadlineFunction(function func() app.UI) *blazarMainLayout {
	c.IHeadlineFunction = function
	return c
}

func (c *blazarMainLayout) SubtitleText(text string) *blazarMainLayout {
	c.ISubtitle = text
	return c
}

func (c *blazarMainLayout) SubtitleFunction(function func() app.UI) *blazarMainLayout {
	c.ISubtitleFunction = function
	return c
}

func (c *blazarMainLayout) TrailerFunction(function func() app.UI) *blazarMainLayout {
	c.ITrailerFunction = function
	return c
}

func (c *blazarMainLayout) DrawerFunction(function func() app.UI) *blazarMainLayout {
	c.IDrawerFunction = function
	return c
}

func (c *blazarMainLayout) OnMount(ctx app.Context) {
	if debugMainLayout {
		slog.DebugContext(ctx.Context, "MainLayout: OnMount", "responsiveWidth", c.IResponsiveWidth)
	}

	c.matchMedia = matchmedia.New(ctx, "screen and (max-width: "+c.IResponsiveWidth+")")
	c.matchMedia.SetOnChange(func(ctx app.Context, value bool) {
		if debugMainLayout {
			slog.DebugContext(ctx.Context, "MediaPage: MatchMedia: OnChange", "value", value)
		}
		c.narrow = value

		ctx.Update()
	})
}

func (c *blazarMainLayout) toggleDrawer() {
	if debugMainLayout {
		slog.DebugContext(context.TODO(), "MainLayout: toggleDrawer")
	}

	c.drawerVisible = !c.drawerVisible
}

func (c *blazarMainLayout) Render() app.UI {
	if debugMainLayout {
		slog.DebugContext(context.TODO(), "MainLayout: Render")
	}

	iconVisible := false
	drawerVisibleClass := ""
	if c.IDrawerFunction != nil {
		drawerVisibleClass = "visible"

		if c.narrow {
			iconVisible = true
			drawerVisibleClass = ""
			if c.drawerVisible {
				drawerVisibleClass = "visible"
			}
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
				HeadlineFunction(c.IHeadlineFunction).
				SubtitleText(c.ISubtitle).
				SubtitleFunction(c.ISubtitleFunction).
				TrailerFunction(c.ITrailerFunction),
			app.Div().
				Class("blazar-main-layout__body").
				Body(
					app.If(c.IDrawerFunction != nil, func() app.UI {
						return app.Div().
							Class("blazar-main-layout__drawer", drawerVisibleClass).
							Body(c.IDrawerFunction())
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
